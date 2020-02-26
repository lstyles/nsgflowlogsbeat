package workers

import (
	"encoding/json"
	"strings"
	"sync"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
)

// Processor worker is responsible for parsing received byte array
// into individual messages and passing them to beat client for publishing
type Processor struct {
	processorQueue chan ProcessorQueueItem
	client         beat.Client
}

// NewProcessor creates a new instance of Processor worker
func NewProcessor(processorQueue chan ProcessorQueueItem, publisher beat.Pipeline) (*Processor, error) {

	client, err := publisher.Connect()
	if err != nil {
		return nil, err
	}

	p := &Processor{
		processorQueue: processorQueue,
		client:         client,
	}

	return p, nil
}

// StartWorker starts the reader worker
func (p *Processor) StartWorker(workerIndex int, wg *sync.WaitGroup) {

	logp.Info("Starting message processor worker #%v", workerIndex)

	defer wg.Done()

	for {
		item, more := <-p.processorQueue
		if more {
			go p.processMessages(item)
		} else {
			logp.Info("Stopping message processor worker #%v", workerIndex)
			return
		}
	}
}

func (p *Processor) processMessages(queueItem ProcessorQueueItem) {

	var messages []NsgMessage
	var events []beat.Event

	json.Unmarshal(*queueItem.Data, &messages)

	logp.Info("Got: %v messages", len(messages))

	for _, message := range messages {
		for _, outerFlow := range message.Properties.Flows {
			for _, innerFlow := range outerFlow.Flows {
				for _, flowTuple := range innerFlow.FlowTuples {
					tuple := strings.Split(flowTuple, ",")
					event := beat.Event{
						Timestamp: message.Time,
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
							"blobName":  queueItem.Name,
							"blobIndex": queueItem.Index,
							"blobETag":  queueItem.ETag,
						},
					}

					events = append(events, event)
				}
			}
		}
	}
	p.client.PublishAll(events)
}
