package nsgflowlogs

import (
	"context"
	"sync"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/common/atomic"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/lstyles/nsgflowlogsbeat/checkpoint"
	"github.com/lstyles/nsgflowlogsbeat/config"
)

type EventACKer struct {
	active          *atomic.Int
	wg              *sync.WaitGroup
	config          *config.Config
	checkpointTable *checkpoint.Table
}

func NewEventACKer(config *config.Config) (*EventACKer, error) {

	ct, err := checkpoint.NewCheckpointTable(config.StorageAccountName, config.StorageAccountKey, config.CheckpointsTableName, config.CheckpointsTableTimeout)
	if err != nil {
		return nil, err
	}

	return &EventACKer{
		active:          atomic.NewInt(0),
		wg:              &sync.WaitGroup{},
		config:          config,
		checkpointTable: ct,
	}, nil
}

// ACKEvents receives callbacks from the publisher for every event that is
// published. It persists the record number of the last event in each
func (a *EventACKer) ACKLastEvent(data []interface{}) {

	for _, d := range data {

		logp.Info("Acking event %v", d)
		msg := d.(common.MapStr)
		batch_complete, err := msg.GetValue("batch_complete")
		if err != nil {
			logp.Error(err)
			continue
		}

		if batch_complete.(bool) {
			logp.Info("Batch complete: %v", batch_complete)

			index, _ := msg.GetValue("index")
			partitionKey, _ := msg.GetValue("partitionKey")
			rowKey, _ := msg.GetValue("rowKey")
			etag, _ := msg.GetValue("etag")

			a.checkpointTable.UpdateCheckpoint(partitionKey.(string), rowKey.(string), etag.(string), index.(int64))
		}

	}
}

func (a *EventACKer) ACKCount(i int) {
	logp.Info("Acking %d out of %d active events.", i, a.Active())
	a.active.Add(-1 * i)
	a.wg.Add(-1 * i)
}

// Wait waits for all events to be ACKed or for the context to be done.
func (a *EventACKer) Wait(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		defer cancel()
		a.wg.Wait()
	}()
	<-ctx.Done()
}

// Add adds to the number of active events.
func (a *EventACKer) Add(delta int) {
	a.active.Add(delta)
	a.wg.Add(delta)
}

// Active returns the number of active events (published but not yet ACKed).
func (a *EventACKer) Active() int {
	return a.active.Load()
}
