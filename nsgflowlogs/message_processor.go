package nsgflowlogs

import (
	"encoding/json"
	"strings"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
)

// MessageProcessor - used to process JSON arrays coming from blobs
type MessageProcessor struct {
	ProcessorQueue chan string
	client         beat.Client
}

// NewMessageProcessor - Creates new instance of MessageProcessor
func NewMessageProcessor(processorQueue chan string, client beat.Client) (*MessageProcessor, error) {

	mp := &MessageProcessor{
		ProcessorQueue: processorQueue,
		client:         client,
	}

	return mp, nil
}

func (mp *MessageProcessor) Run(workerIndex int) {
	logp.Info("Starting message processor worker %v", workerIndex)
	for {
		m, more := <-mp.ProcessorQueue
		if more {
			go mp.ProcessMessages(m)
		} else {
			return
		}
	}
}

func (mp *MessageProcessor) ProcessMessages(data string) {

	var messages []NsgMessage
	var events []beat.Event

	json.Unmarshal([]byte(data), &messages)

	logp.Info("Got: %v messages", len(messages))

	for _, message := range messages {
		for _, outerFlow := range message.Properties.Flows {
			for _, innerFlow := range outerFlow.Flows {
				for _, flowTuple := range innerFlow.FlowTuples {
					tuple := strings.Split(flowTuple, ",")
					event := beat.Event{
						Timestamp: message.Time,
						Fields: common.MapStr{
							"systemId":           message.SystemId,
							"macAddress":         message.MacAddress,
							"category":           message.Category,
							"resourceId":         message.ResourceId,
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
					}

					events = append(events, event)
				}
			}
		}
	}
	mp.client.PublishAll(events)
}
