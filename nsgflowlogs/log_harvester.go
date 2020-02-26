package nsgflowlogs

import (
	"sync"
	"time"

	"github.com/lstyles/nsgflowlogsbeat/storagereader"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/lstyles/nsgflowlogsbeat/checkpoint"
	"github.com/lstyles/nsgflowlogsbeat/config"
	"github.com/lstyles/nsgflowlogsbeat/workers"
)

// LogHarvester - responsible for retrieving blobs and managing state
type LogHarvester struct {
	config          *config.Config
	publisher       beat.Pipeline
	CheckpointTable *checkpoint.Table
	ReaderQueue     chan workers.ReaderQueueItem
	ProcessorQueue  chan workers.ProcessorQueueItem
}

// NewLogHarvester - Creates a new instance of LogHarvester
func NewLogHarvester(config *config.Config, publisher beat.Pipeline) (*LogHarvester, error) {

	logp.Info("Initializing Log Harvester")

	ct, err := checkpoint.NewCheckpointTable(config.StorageAccountName, config.StorageAccountKey, config.CheckpointsTableName, config.CheckpointsTableTimeout)
	if err != nil {
		return nil, err
	}

	logp.Info("Initialized checkpoints table")

	lh := &LogHarvester{
		config:          config,
		publisher:       publisher,
		CheckpointTable: ct,
		ReaderQueue:     make(chan workers.ReaderQueueItem),
		ProcessorQueue:  make(chan workers.ProcessorQueueItem),
	}

	return lh, nil
}

func (lh *LogHarvester) ScanAndProcessUpdates() {

	// Start message processors
	var wg sync.WaitGroup
	for mpw := 1; mpw <= lh.config.MessageProcessorWorkers; mpw++ {
		// Run message proccessor worker
		wg.Add(1)
		mp, mperr := workers.NewProcessor(lh.ProcessorQueue, lh.publisher)
		if mperr != nil {
			panic(mperr)
		}

		go mp.StartWorker(mpw, &wg)
	}
	// Start storage readers
	for srw := 1; srw <= lh.config.StorageReaderWorkers; srw++ {
		// Run storage reader worker
		wg.Add(1)
		sr, err := workers.NewReader(lh.config, lh.ReaderQueue, lh.ProcessorQueue)
		if err != nil {
			panic(err)
		}

		go sr.StartWorker(srw, &wg)
	}

	// Scan for updated blobs
	wg.Add(1)
	go lh.ScanForChanges(&wg)

	// Wait for all workers to finish
	wg.Wait()
}

// ScanForChanges - Scans storage account for blobs that need to be parsed
func (lh *LogHarvester) ScanForChanges(wg *sync.WaitGroup) {

	logp.Info("Scanning storage account for changes")

	defer wg.Done()

	timeNow := int64(time.Now().UTC().Unix())
	ignoreOlder := int64(lh.config.IgnoreOlder.Seconds())

	var startTime int64
	if ignoreOlder == 0 {
		startTime = 0
	} else {
		startTime = timeNow - ignoreOlder
	}

	endTime := timeNow

	sr, serr := storagereader.NewStorageReader(lh.config.StorageAccountName, lh.config.StorageAccountKey, lh.config.ContainerName)
	if serr != nil {
		panic(serr)
	}

	blobsToProcess := sr.ListBlobsModifiedBetween(startTime, endTime)
	blobsCount := len(*blobsToProcess)
	logp.Info("Found %v blobs", blobsCount)

	for i, blob := range *blobsToProcess {
		i++

		c, err := lh.CheckpointTable.GetCheckpoint(blob.PartitionKey, blob.RowKey)
		if err != nil {
			logp.Error(err)
			continue
		}

		index := int64(0)

		var checkPoint checkpoint.Checkpoint
		if c != nil {
			checkPoint = *c
			if blob.ETag == checkPoint.ETag {
				logp.Info("Blob ETag hasn't changed. Skipping.")
				continue
			}
			index = c.Index
		}

		q := workers.ReaderQueueItem{
			Name:   blob.Name,
			Index:  index,
			Length: blob.Length,
		}

		lh.ReaderQueue <- q
	}

	close(lh.ReaderQueue)

	logp.Info("Finished scanning storage account for changes.")
}
