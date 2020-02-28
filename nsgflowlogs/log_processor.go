package nsgflowlogs

import (
	"context"
	"sync"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/lstyles/nsgflowlogsbeat/checkpoint"
	"github.com/lstyles/nsgflowlogsbeat/config"
	"github.com/lstyles/nsgflowlogsbeat/storagereader"
)

type LogProcessor struct {
	beat            *beat.Beat
	config          *config.Config
	done            chan struct{}
	pipeline        beat.Pipeline
	acker           *EventACKer
	storageReader   *storagereader.StorageReader
	checkpointTable *checkpoint.Table
	readerQueue     chan ReaderQueueItem
}

func NewLogProcessor(b *beat.Beat, c *config.Config, done chan struct{}) (*LogProcessor, error) {

	logp.Debug("log_processor", "Initalizing new log processor.")

	ct, err := checkpoint.NewCheckpointTable(c.StorageAccountName, c.StorageAccountKey, c.CheckpointsTableName, c.CheckpointsTableTimeout)
	if err != nil {
		return nil, err
	}

	sr, serr := storagereader.NewStorageReader(
		c.StorageAccountName,
		c.StorageAccountKey,
		c.ContainerName,
	)
	if serr != nil {
		return nil, serr
	}

	acker, err := NewEventACKer(c)

	lp := &LogProcessor{
		beat:            b,
		config:          c,
		done:            done,
		pipeline:        b.Publisher,
		acker:           acker,
		storageReader:   sr,
		checkpointTable: ct,
		readerQueue:     make(chan ReaderQueueItem, 2000),
	}

	lp.pipeline.SetACKHandler(beat.PipelineACKHandler{
		ACKLastEvents: lp.acker.ACKLastEvent,
	})

	return lp, nil
}

func (lp *LogProcessor) Process(done chan struct{}) {

	logp.Info("Processing NSG flow logs")

	// Get list of blobs to process
	var wg sync.WaitGroup

	// Initalize workers
	for i := 1; i <= lp.config.Workers; i++ {

		// Run message proccessor worker
		w, err := NewWorker(lp.config, lp.pipeline, lp.readerQueue)
		if err != nil {
			panic(err)
		}

		wg.Add(1)
		go w.Run(i, &wg, done)
	}

	wg.Add(1)
	go lp.ScanForUpdatedBlobs(&wg)

	logp.Info("Waiting for all workers to finish")
	logp.Info("Waiting for acker to finish")

	wg.Wait()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	lp.acker.Wait(ctx)

}

func (lp *LogProcessor) ScanForUpdatedBlobs(wg *sync.WaitGroup) {

	logp.Info("Scanning for updated blobs")

	defer wg.Done()

	timeNow := int64(time.Now().UTC().Unix())
	ignoreOlder := int64(lp.config.IgnoreOlder.Seconds())

	var startTime int64
	if ignoreOlder == 0 {
		startTime = 0
	} else {
		startTime = timeNow - ignoreOlder
	}

	endTime := timeNow

	blobsToProcess := lp.storageReader.ListBlobsModifiedBetween(startTime, endTime)
	blobsCount := len(*blobsToProcess)

	logp.Info("Found %v blobs", blobsCount)

	for i, blob := range *blobsToProcess {
		i++

		c, err := lp.checkpointTable.GetCheckpoint(blob.PartitionKey, blob.RowKey)
		if err != nil {
			logp.Error(err)
			continue
		}

		finishFile := blob.Length-c.Index == 2

		c.Length = blob.Length
		if finishFile {
			logp.Info("This happened...")
		}

		logp.Info("Blob length is %d. Updating checkpoint", c.Length)
		err = lp.checkpointTable.CreateOrUpdateCheckpoint(c)
		if err != nil {
			logp.Error(err)
			continue
		}

		//if finishFile {
		//	continue
		//}

		if blob.ETag == c.ETag {
			logp.Info("Blob ETag hasn't changed. Skipping.")
			continue
		}

		q := ReaderQueueItem{
			Name:         blob.Name,
			Index:        c.Index,
			PartitionKey: c.PartitionKey,
			RowKey:       c.RowKey,
			ETag:         blob.ETag,
		}

		logp.Debug("trace", "Sending blob %v/%v to the reader queue", i, blobsCount)
		lp.readerQueue <- q
	}

	close(lp.readerQueue)

	logp.Info("Finished scanning storage account for changes.")
}
