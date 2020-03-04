package nsgflowlogs

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/lstyles/nsgflowlogsbeat/config"
	"github.com/lstyles/nsgflowlogsbeat/storagereader"
)

type Worker struct {
	config        *config.Config
	client        beat.Client
	storageReader *storagereader.StorageReader
	readerQueue   chan ReaderQueueItem
	acker         *EventACKer
}

func NewWorker(config *config.Config, pipeline beat.Pipeline, readerQueue chan ReaderQueueItem) (*Worker, error) {

	a, ackerr := NewEventACKer(config)
	if ackerr != nil {
		return nil, ackerr
	}

	client, err := pipeline.ConnectWith(beat.ClientConfig{
		ACKCount:    a.ACKCount,
		PublishMode: beat.GuaranteedSend,
	})
	if err != nil {
		return nil, err
	}

	sr, serr := storagereader.NewStorageReader(
		config.StorageAccountName,
		config.StorageAccountKey,
		config.ContainerName,
	)
	if serr != nil {
		return nil, serr
	}

	w := &Worker{
		config:        config,
		client:        client,
		storageReader: sr,
		readerQueue:   readerQueue,
		acker:         a,
	}

	return w, nil
}

func (w *Worker) Run(i int, wg *sync.WaitGroup, done chan struct{}) {

	logp.Info("Starting worker %d", i)

	defer wg.Done()
	defer w.client.Close()
	go func() {
		<-done
		w.client.Close()
	}()

	defer func() {
		logp.Info("Stopped processing")
	}()

	for stop := false; !stop; {
		select {
		case <-done:
			return
		default:
		}

		item, more := <-w.readerQueue
		if !more {
			logp.Info("No more items in the queue. Stopping worker")
			return
		}

		logp.Info("Processing reader queue item.")

		fromStart := item.Index == 0
		if !fromStart {
			item.Index -= 2
		}
		data := w.storageReader.ReadBlobData(item.Name, item.Index, item.Length-item.Index)
		length := int64(len(data))
		if length == 2 {
			logp.Warn("Downloaded only 2 bytes")
			continue
		}

		logp.Info("Downloaded %d bytes", length)

		if fromStart {
			// Reading from beginning of the file
			data = data[11 : len(data)-1]
		} else {
			data[0] = byte('[')
			data = data[0 : len(data)-1]
		}

		dLen := len(data)

		// if the blob has been written to since we scanned, we have to replace the last character
		if data[dLen-1] == byte(',') {
			data[dLen-1] = byte(']')
		}

		logp.Info("Data after trimming: %s", data[0:15])
		logp.Info("Data after trimming: %s", data[dLen-15:dLen])

		var messages []NsgMessage
		var events []beat.Event

		json.Unmarshal(data, &messages)

		logp.Info("Got: %v messages", len(messages))

		for _, message := range messages {
			for _, outerFlow := range message.Properties.Flows {
				for _, innerFlow := range outerFlow.Flows {
					for _, flowTuple := range innerFlow.FlowTuples {
						tuple := strings.Split(flowTuple, ",")
						i, err := strconv.ParseInt(tuple[0], 10, 64)
						if err != nil {
							logp.Error(err)
							continue
						}

						event := beat.Event{
							Timestamp: time.Unix(i, 0),
							Fields: common.MapStr{
								"systemId":           message.SystemID,
								"macAddress":         message.MacAddress,
								"category":           message.Category,
								"resourceId":         message.ResourceID,
								"operationName":      message.OperationName,
								"version":            message.Properties.Version,
								"nsgRuleName":        outerFlow.Rule,
								"startTime":          tuple[0],
								"sourceAddress":      tuple[1],
								"destinationAddress": tuple[2],
								"sourcePort":         tuple[3],
								"destinationPort":    tuple[4],
								"transportProtocol":  tuple[5],
								"deviceDirection":    tuple[6],
								"deviceAction":       tuple[7],
								"flowState":          tuple[8],
								"packetsStoD":        tuple[9],
								"bytesStoD":          tuple[10],
								"packetsDtoS":        tuple[11],
								"bytesDtoS":          tuple[12],
							},
							Private: common.MapStr{
								"partitionKey": item.PartitionKey,
								"rowKey":       item.RowKey,
								"index":        item.Index + length,
								"etag":         item.ETag,
							},
						}

						events = append(events, event)
					}
				}
			}
		}

		if len(events) > 0 {
			// Add batch_complete to last event
			logp.Info("Processing last message from the batch.")
			levt := events[len(events)-1]
			levt.Private.(common.MapStr).Update(common.MapStr{"batch_complete": true})

			w.acker.Add(len(events))
			w.client.PublishAll(events)

			logp.Info("Published events. Waiting for %d acks", w.acker.Active())

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			w.acker.Wait(ctx)
		} else {
			logp.Info("No events to publish")
		}
	}

	return
}
