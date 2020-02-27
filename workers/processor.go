package workers

import (
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

// Processor worker is responsible for parsing received byte array
// into individual messages and passing them to beat client for publishing
type Processor struct {
	storageReader *storagereader.StorageReader
	config        *config.Config
	readerQueue   chan ReaderQueueItem
	client        beat.Client
}

// NewProcessor creates a new instance of Processor worker
func NewProcessor(publisher beat.Pipeline, config *config.Config, readerQueue chan ReaderQueueItem) (*Processor, error) {

	client, err := publisher.ConnectWith(beat.ClientConfig{
		PublishMode: beat.GuaranteedSend,
		ACKEvents: func(data []interface{}) {
			logp.Info("Successfully published %v", data)
		},
		/*	ACKLastEvent: func(data interface{}) {
				logp.Info("Successfully published last %v", data)
			},
			ACKCount: func(n int) {
				logp.Info("Successfully published %d events.", n)
			},*/
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

	p := &Processor{
		storageReader: sr,
		config:        config,
		readerQueue:   readerQueue,
		client:        client,
	}

	return p, nil
}

// StartWorker starts the reader worker
func (p *Processor) StartWorker(workerIndex int, wg *sync.WaitGroup) {

	logp.Info("Starting message processor worker #%v", workerIndex)

	defer wg.Done()
	defer p.client.Close()

	for {
		item, more := <-p.readerQueue
		if more {
			go p.processMessages(item)
		} else {
			logp.Info("Stopping message processor worker #%v", workerIndex)
			return
		}
	}
}

func (p *Processor) processMessages(queueItem ReaderQueueItem) {

	data := p.storageReader.ReadBlobData(queueItem.Name, queueItem.Index, queueItem.Length)
	length := int64(len(data))

	if queueItem.Index == 0 {
		// Starting from beginning of file
		data = data[11 : len(data)-1]
	} else {
		// Starting from saved position
		data[0] = byte('[')
		data = data[0 : len(data)-1]
	}

	// Sometimes blob has been written to before we got round to reading it
	// and the last character is going to be a comma instead of ']'.
	// We're handling it here
	if data[len(data)-1] == byte(',') {
		data[len(data)-1] = byte(']')
	}

	logp.Info("Data after trimming: %s", data[0:15])
	logp.Info("Data after trimming: %s", data[len(data)-15:len(data)])

	queueItem.Index += length

	logp.Debug("trace", "Sending data string to the processor queue")

	var messages []NsgMessage
	var events []beat.Event

	json.Unmarshal(data, &messages)

	logp.Info("Got: %v messages", len(messages))

	for _, message := range messages {
		for _, outerFlow := range message.Properties.Flows {
			for _, innerFlow := range outerFlow.Flows {
				for _, flowTuple := range innerFlow.FlowTuples {
					tuple := strings.Split(flowTuple, ",")
					i, err := strconv.ParseInt("1405544146", 10, 64)
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

	logp.Debug("trace", "Publishing %v messages to beat pipeline", len(events))

	p.client.PublishAll(events)
}
