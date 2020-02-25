package nsgflowlogs

import (
	"bytes"
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/elastic/beats/libbeat/logp"
)

// StorageReader - Responsible for interacting with blob storage
type StorageReader struct {
	container        *azblob.ContainerURL
	checkpointsTable *CheckpointsTable
	ReaderQueue      chan ReaderQueueItem
	ProcessorQueue   chan string
}

// NewStorageReader - Creates new instance of StorageReader
func NewStorageReader(accountName, accountKey, containerName string, checkpointsTable *CheckpointsTable, readerQueue chan ReaderQueueItem, processorQueue chan string) (*StorageReader, error) {

	logp.Info("Initializing storage reader on container %s in account %s", containerName, accountName)

	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		return nil, err
	}
	p := azblob.NewPipeline(credential, azblob.PipelineOptions{})

	URL, _ := url.Parse(fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerName))
	containerURL := azblob.NewContainerURL(*URL, p)

	sr := &StorageReader{
		container:        &containerURL,
		checkpointsTable: checkpointsTable,
		ReaderQueue:      readerQueue,
		ProcessorQueue:   processorQueue,
	}

	return sr, nil
}

func (sr *StorageReader) Run(workerIndex int) {
	logp.Info("Starting storage reader worker %v", workerIndex)
	for {
		r, more := <-sr.ReaderQueue
		if more {
			go sr.DownloadAndPublish(r)
		} else {
			return
		}
	}
}

// ListBlobsModifiedBetween - Get list of blobs modified between two specified timestamps
func (sr *StorageReader) ListBlobsModifiedBetween(startTime, endTime int64) []BlobDetails {

	ctx := context.Background()

	var blobItems []BlobDetails

	for marker := (azblob.Marker{}); marker.NotDone(); {

		listBlob, err := sr.container.ListBlobsFlatSegment(ctx, marker, azblob.ListBlobsSegmentOptions{})
		marker = listBlob.NextMarker
		if err != nil {
			logp.Error(err)
			continue
		}

		for _, blobInfo := range listBlob.Segment.BlobItems {

			lastModified := blobInfo.Properties.LastModified.UTC().Unix()
			if lastModified > startTime && lastModified < endTime {
				logp.Debug("storage_reader", "Found matching blob: %v", blobInfo.Name)
				length := blobInfo.Properties.ContentLength
				blobItems = append(blobItems, NewBlobDetails(blobInfo.Name, string(blobInfo.Properties.Etag), *length, lastModified))
			}
		}
	}

	return blobItems
}

func (sr *StorageReader) DownloadAndPublish(queueItem ReaderQueueItem) {

	data := sr.ReadBlobData(queueItem.BlobDetails.Name, queueItem.Checkpoint.Index, queueItem.BlobDetails.Length)
	if len(data) < 10 {
		return
	}

	length := int64(len(data))

	if queueItem.Checkpoint.Index == 0 {
		// Starting from beginning of file
		data = data[11 : len(data)-1]
	} else {
		// Starting from saved position
		data = strings.Replace(data, ",", "[", 1)
		data = data[0 : len(data)-1]
	}

	queueItem.Checkpoint.Index = queueItem.Checkpoint.Index + length - 2
	sr.checkpointsTable.CreateOrUpdateCheckpoint(&queueItem.Checkpoint)

	logp.Info("Data after trimming: %s", data[0:15])
	logp.Info("Data after trimming: %s", data[len(data)-15:len(data)])
	sr.ProcessorQueue <- data
}

// ReadBlobData - Reads blob from specified starting location
func (sr *StorageReader) ReadBlobData(path string, startIndex int64, endIndex int64) string {

	ctx := context.Background()

	blobURL := sr.container.NewBlockBlobURL(path)
	downloadResponse, err := blobURL.Download(ctx, startIndex, endIndex, azblob.BlobAccessConditions{}, false)

	logp.Debug("storage_reader", "Attempting to download blob %s at %v", path, startIndex)

	bodyStream := downloadResponse.Body(azblob.RetryReaderOptions{MaxRetryRequests: 10})

	downloadedData := bytes.Buffer{}
	_, err = downloadedData.ReadFrom(bodyStream)
	if err != nil {
		panic(err)
	}

	return downloadedData.String()
}
