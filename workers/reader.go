package workers

import (
	"sync"

	"github.com/elastic/beats/libbeat/logp"
	"github.com/lstyles/nsgflowlogsbeat/config"
	"github.com/lstyles/nsgflowlogsbeat/storagereader"
)

// Reader worker is responsible for reading block data from Azure blob and
// passing it to the processor workers
type Reader struct {
	storageReader  *storagereader.StorageReader
	config         *config.Config
	readerQueue    chan ReaderQueueItem
	processorQueue chan ProcessorQueueItem
}

// NewReader creates a new instance of Reader worker
func NewReader(config *config.Config, readerQueue chan ReaderQueueItem, processorQueue chan ProcessorQueueItem) (*Reader, error) {

	sr, err := storagereader.NewStorageReader(
		config.StorageAccountName,
		config.StorageAccountKey,
		config.ContainerName,
	)
	if err != nil {
		return nil, err
	}

	r := &Reader{
		storageReader:  sr,
		config:         config,
		readerQueue:    readerQueue,
		processorQueue: processorQueue,
	}

	return r, nil
}

// StartWorker starts the reader worker
func (r *Reader) StartWorker(workerIndex int, wg *sync.WaitGroup) {

	logp.Info("Starting storage reader worker #%v", workerIndex)

	defer wg.Done()

	for {
		item, more := <-r.readerQueue
		if more {
			go r.downloadAndPublish(item)
		} else {
			logp.Info("Stopping storage reader worker #%v", workerIndex)
			//close(r.processorQueue)
			return
		}
	}
}

func (r *Reader) downloadAndPublish(queueItem ReaderQueueItem) {

	data := r.storageReader.ReadBlobData(queueItem.Name, queueItem.Index, queueItem.Length)
	length := int64(len(data))

	if queueItem.Index == 0 {
		// Starting from beginning of file
		data = data[11 : len(data)-1]
	} else {
		// Starting from saved position
		data[0] = byte('[')
		data = data[0 : len(data)-1]
	}

	logp.Info("Data after trimming: %s", data[0:15])
	logp.Info("Data after trimming: %s", data[len(data)-15:len(data)])

	d := ProcessorQueueItem{
		Name:  queueItem.Name,
		Index: queueItem.Index + length,
		Data:  &data,
		ETag:  queueItem.ETag,
	}

	r.processorQueue <- d
}
