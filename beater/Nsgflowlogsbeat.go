package beater

import (
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
	done         chan struct{}
	config       config.Config
	client       beat.Client
	LogHarvester *nsgflowlogs.LogHarvester
}

// New creates an instance of nsgflowlogsbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	if c.ScanFrequency.Seconds() < 30 {
		logp.Warn("Chosen interval of %s is not valid. Changing to default 30s", c.ScanFrequency.String())
		c.ScanFrequency = 1 * time.Minute
	}

	if len(c.StorageAccountName) == 0 || len(c.StorageAccountKey) == 0 {
		panic("Storage account name and key are both required.")
	}

	logp.Info("Storage account name %s is %d characters long", c.StorageAccountName, len(c.StorageAccountName))

	lh, lherr := nsgflowlogs.NewLogHarvester(&c, make(chan *nsgflowlogs.ReaderQueueItem), make(chan string))
	if lherr != nil {
		panic(lherr)
	}

	bt := &nsgflowlogsbeat{
		done:         make(chan struct{}),
		config:       c,
		LogHarvester: lh,
	}
	return bt, nil
}

// Run starts nsgflowlogsbeat.
func (bt *nsgflowlogsbeat) Run(b *beat.Beat) error {
	logp.Info("nsgflowlogsbeat is running! Hit CTRL-C to stop it.")

	_, err := b.Publisher.Connect()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(bt.config.ScanFrequency)
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}
		// Start Message Processor workers

		// Start Storage Reader workers

		// Scan for changes
		bt.LogHarvester.ScanForChanges()

		// Wait for workers
	}
}

// Stop stops nsgflowlogsbeat.
func (bt *nsgflowlogsbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
