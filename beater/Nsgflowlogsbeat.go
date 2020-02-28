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

// Nsgflowlogsbeat is used to conform to the beat interface.
type Nsgflowlogsbeat struct {
	config config.Config
	done   chan struct{}
}

// New creates an instance of nsgflowlogsbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {

	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading configuration file: %v", err)
	}

	if c.ScanFrequency.Seconds() < 30 {
		logp.Warn("Chosen interval of %s is not valid. Changing to default 30s", c.ScanFrequency.String())
		c.ScanFrequency = 30 * time.Second
	}

	logp.Debug("nsgflowlogsbeat", "Validating configuration")
	if err := c.Validate(); err != nil {
		panic(err)
	}

	bt := &Nsgflowlogsbeat{
		done:   make(chan struct{}),
		config: c,
	}

	return bt, nil
}

// Run starts nsgflowlogsbeat.
func (bt *Nsgflowlogsbeat) Run(b *beat.Beat) error {
	logp.Info("nsgflowlogsbeat is running! Hit CTRL-C to stop it.")

	ticker := time.NewTicker(bt.config.ScanFrequency)
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		lp, err := nsgflowlogs.NewLogProcessor(b, &bt.config, bt.done)
		if err != nil {
			panic(err)
		}

		lp.Process(bt.done)
	}
}

// Stop stops nsgflowlogsbeat.
func (bt *Nsgflowlogsbeat) Stop() {
	logp.Info("Stopping nsgflowlogsbeat...")
	if bt.done != nil {
		close(bt.done)
	}
}
