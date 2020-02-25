package nsgflowlogs

import (
	"time"

	"github.com/elastic/beats/libbeat/logp"
	"github.com/lstyles/nsgflowlogsbeat/config"
)

// LogHarvester - responsible for retrieving blobs and managing state
type LogHarvester struct {
	config           *config.Config
	checkpointsTable *CheckpointsTable
	storageReader    *StorageReader
	ReaderQueue      chan *ReaderQueueItem
	ProcessorQueue   chan string
}

type ReaderQueueItem struct {
	BlobDetails *BlobDetails
	Checkpoint  *Checkpoint
}

// NewLogHarvester - Creates a new instance of LogHarvester
func NewLogHarvester(config *config.Config, readerQueue chan *ReaderQueueItem, processorQueue chan string) (*LogHarvester, error) {

	logp.Info("Initializing Log Harvester")

	ct, err := NewCheckpointsTable(config.StorageAccountName, config.StorageAccountKey, config.CheckpointsTableName, config.CheckpointsTableTimeout)
	if err != nil {
		return nil, err
	}

	logp.Info("Initialized checkpoints table")

	sr, serr := NewStorageReader(config.StorageAccountName, config.StorageAccountKey, config.ContainerName)
	if serr != nil {
		return nil, serr
	}

	logp.Info("Initialized storage reader")

	lh := &LogHarvester{
		config:           config,
		checkpointsTable: ct,
		storageReader:    sr,
		ReaderQueue:      readerQueue,
		ProcessorQueue:   processorQueue,
	}

	return lh, nil
}

func (lh *LogHarvester) HarvestLogs() {
	for {
		r, more := <-lh.ReaderQueue
		if more {
			go lh.DownloadAndPublish(r)
		} else {
			return
		}
	}
}

func (lh *LogHarvester) DownloadAndPublish(queueItem *ReaderQueueItem) {

	data := lh.storageReader.ReadBlobData(queueItem.BlobDetails.Name, queueItem.Checkpoint.Index)

	queueItem.Checkpoint.Index = int64(len(data) - 1)
	lh.checkpointsTable.CreateOrUpdateCheckpoint(queueItem.Checkpoint)

	if queueItem.Checkpoint.Index == 0 {
		// Starting from beginning of file
		data = data[11 : len(data)-1]
	} else {
		// Starting from saved position
		data = data[1 : len(data)-1]
	}

	lh.ProcessorQueue <- data
}

// ScanForChanges - Scans storage account for blobs that need to be parsed
func (lh *LogHarvester) ScanForChanges() {

	logp.Info("Scanning storage account for changes")

	timeNow := int64(time.Now().UTC().Unix())
	ignoreOlder := int64(lh.config.IgnoreOlder.Seconds())
	minStart := timeNow - ignoreOlder

	lc, err := lh.checkpointsTable.GetCheckpoint("System", "LastScanTS")
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
	err = lh.checkpointsTable.CreateOrUpdateCheckpoint(lc)
	if err != nil {
		logp.Error(err)
	}

	blobsToProcess := lh.storageReader.ListBlobsModifiedBetween(startTime, endTime)
	blobsCount := len(blobsToProcess)
	logp.Info("Found %v blobs", blobsCount)

	for i, blob := range blobsToProcess {
		i++

		c, err := lh.checkpointsTable.GetCheckpoint(blob.PartitionKey, blob.RowKey)
		if err != nil {
			logp.Error(err)
			continue
		}

		if c != nil {
			if blob.ETag == c.ETag {
				logp.Info("Blob ETag hasn't changed. Skipping.")
				continue
			}
		} else {
			c = &Checkpoint{
				PartitionKey: blob.PartitionKey,
				RowKey:       blob.RowKey,
				ETag:         "", // Only set ETag when updating Index after reading the blob
				Length:       blob.Length,
				Index:        0,
			}
		}

		q := &ReaderQueueItem{
			BlobDetails: &blob,
			Checkpoint:  c,
		}

		lh.ReaderQueue <- q
	}

	logp.Info("Finished scanning storage account for changes.")
}
