package beater

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/lstyles/nsgflowlogsbeat/config"
	"github.com/lstyles/nsgflowlogsbeat/nsgflowlogs"
)

// nsgflowlogsbeat configuration.
type nsgflowlogsbeat struct {
	done             chan struct{}
	config           config.Config
	client           beat.Client
	checkpointsTable *nsgflowlogs.CheckpointsTable
	storageReader    *nsgflowlogs.StorageReader
}

// New creates an instance of nsgflowlogsbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	if c.ScanFrequency.Seconds() < 10 {
		logp.Warn("Chosen interval of %s is not valid. Changing to default 10s", c.ScanFrequency.String())
		c.ScanFrequency = 1 * time.Minute
	}

	if len(c.StorageAccountName) == 0 || len(c.StorageAccountKey) == 0 {
		panic("Storage account name and key are both required.")
	}

	logp.Info("Storage account name %s is %d characters long", c.StorageAccountName, len(c.StorageAccountName))

	bt := &nsgflowlogsbeat{
		done:             make(chan struct{}),
		config:           c,
		checkpointsTable: nsgflowlogs.NewCheckpointsTable(c.StorageAccountName, c.StorageAccountKey, c.CheckpointsTableName, c.CheckpointsTableTimeout),
		storageReader:    nsgflowlogs.NewStorageReader(c.StorageAccountName, c.StorageAccountKey, c.ContainerName),
	}
	return bt, nil
}

// Run starts nsgflowlogsbeat.
func (bt *nsgflowlogsbeat) Run(b *beat.Beat) error {
	logp.Info("nsgflowlogsbeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	// Catch up first, before getting into scan frequency
	timeNow := int64(time.Now().UTC().Unix())
	ignoreOlder := int64(bt.config.IgnoreOlder.Seconds())
	lastScanTS := bt.checkpointsTable.GetLastScanTS()

	logp.Info("Got: timeNow: %v, ignoreOlder: %v, lastScanTS: %v", timeNow, ignoreOlder, lastScanTS)

	// Earliest possible start taking ignore_older config setting into account
	minStart := timeNow - ignoreOlder

	startTime := int64(0)
	if lastScanTS == 0 || lastScanTS < minStart {
		startTime = minStart
	} else {
		startTime = lastScanTS
	}

	endTime := timeNow

	logp.Info("Scanning blobs modified between %v and %v", startTime, endTime)

	blobsToProcess := bt.storageReader.ListBlobsModifiedBetween(startTime, endTime)
	blobsCount := len(blobsToProcess)
	logp.Info("Found %v blobs", blobsCount)

	for i, blob := range blobsToProcess {
		i++
		logp.Info("Processing %d/%d - %s - %s", i, blobsCount, blob.PartitionKey, blob.RowKey)

		index := bt.checkpointsTable.GetCheckpoint(blob.PartitionKey, blob.RowKey)
		data := bt.storageReader.ReadBlobData(blob.Name, index)

		var messages []string

		if index == 0 {
			// Starting from beginning of file
			data = data[11 : len(data)-1]
		} else {
			// Starting from saved position
			data = data[1 : len(data)-1]
		}

		json.Unmarshal([]byte(data), &messages)

		logp.Info("Got: %v messages", len(messages))

		for _, message := range messages {
			logp.Info(message)
			event := beat.Event{
				Timestamp: time.Now(),
				Fields: common.MapStr{
					"type":    b.Info.Name,
					"message": message,
				},
			}
			bt.client.Publish(event)
			logp.Info("Event sent")
		}

	}

	bt.checkpointsTable.SetLastScanTS(endTime)

	ticker := time.NewTicker(bt.config.ScanFrequency)
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		event := beat.Event{
			Timestamp: time.Now(),
			Fields: common.MapStr{
				"type":    b.Info.Name,
				"message": "",
			},
		}
		bt.client.Publish(event)
		logp.Info("Event sent")
	}
}

// Stop stops nsgflowlogsbeat.
func (bt *nsgflowlogsbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
