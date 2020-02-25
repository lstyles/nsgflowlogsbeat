package nsgflowlogs

import (
	"time"

	"github.com/elastic/beats/libbeat/logp"
	"github.com/lstyles/nsgflowlogsbeat/config"
)

// LogHarvester - responsible for retrieving blobs and managing state
type LogHarvester struct {
	config           *config.Config
	CheckpointsTable *CheckpointsTable
	ReaderQueue      chan ReaderQueueItem
	ProcessorQueue   chan string
}

type ReaderQueueItem struct {
	BlobDetails BlobDetails
	Checkpoint  Checkpoint
}

// NewLogHarvester - Creates a new instance of LogHarvester
func NewLogHarvester(config *config.Config) (*LogHarvester, error) {

	logp.Info("Initializing Log Harvester")

	ct, err := NewCheckpointsTable(config.StorageAccountName, config.StorageAccountKey, config.CheckpointsTableName, config.CheckpointsTableTimeout)
	if err != nil {
		return nil, err
	}

	logp.Info("Initialized checkpoints table")

	lh := &LogHarvester{
		config:           config,
		CheckpointsTable: ct,
		ReaderQueue:      make(chan ReaderQueueItem),
		ProcessorQueue:   make(chan string),
	}

	return lh, nil
}

// ScanForChanges - Scans storage account for blobs that need to be parsed
func (lh *LogHarvester) ScanForChanges() {

	logp.Info("Scanning storage account for changes")

	timeNow := int64(time.Now().UTC().Unix())
	ignoreOlder := int64(lh.config.IgnoreOlder.Seconds())
	minStart := timeNow - ignoreOlder

	lc, err := lh.CheckpointsTable.GetCheckpoint("System", "LastScanTS")
	if err != nil {
		logp.Error(err)
		return
	}

	if lc == nil {
		lc = &Checkpoint{
			Index:        0,
			PartitionKey: "System",
			RowKey:       "LastScanTS",
		}
	} else {
		logp.Info("Retrieved LastScanMS checkpoint with index: %v.", lc.Index)
	}

	startTime := int64(0)
	if lc.Index == 0 || lc.Index < minStart {
		startTime = minStart
	} else {
		startTime = lc.Index
	}

	endTime := timeNow

	lc.Index = time.Now().UTC().Unix()
	err = lh.CheckpointsTable.CreateOrUpdateCheckpoint(lc)
	if err != nil {
		logp.Error(err)
	}

	sr, serr := NewStorageReader(lh.config.StorageAccountName, lh.config.StorageAccountKey, lh.config.ContainerName, lh.CheckpointsTable, lh.ReaderQueue, lh.ProcessorQueue)
	if serr != nil {
		panic(serr)
	}

	blobsToProcess := sr.ListBlobsModifiedBetween(startTime, endTime)
	blobsCount := len(blobsToProcess)
	logp.Info("Found %v blobs", blobsCount)

	for i, blob := range blobsToProcess {
		i++

		c, err := lh.CheckpointsTable.GetCheckpoint(blob.PartitionKey, blob.RowKey)
		if err != nil {
			logp.Error(err)
			continue
		}

		var checkPoint Checkpoint
		if c != nil {
			checkPoint = *c
			if blob.ETag == checkPoint.ETag {
				logp.Info("Blob ETag hasn't changed. Skipping.")
				continue
			}
			c.Length = blob.Length
		} else {
			checkPoint = Checkpoint{
				PartitionKey: blob.PartitionKey,
				RowKey:       blob.RowKey,
				ETag:         "", // Only set ETag when updating Index after reading the blob
				Length:       blob.Length,
				Index:        0,
			}
		}

		q := ReaderQueueItem{
			BlobDetails: blob,
			Checkpoint:  checkPoint,
		}

		lh.ReaderQueue <- q
	}

	logp.Info("Finished scanning storage account for changes.")
}
