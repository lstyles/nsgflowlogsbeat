package nsgflowlogs

import (
	"bytes"
	"context"
	"fmt"
	"net/url"

	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/elastic/beats/libbeat/logp"
)

// StorageReader - Responsible for interacting with blob storage
type StorageReader struct {
	container *azblob.ContainerURL
}

// NewStorageReader - Creates new instance of StorageReader
func NewStorageReader(accountName, accountKey, containerName string) (*StorageReader, error) {

	logp.Info("Initializing storage reader on container %s in account %s", containerName, accountName)

	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		return nil, err
	}
	p := azblob.NewPipeline(credential, azblob.PipelineOptions{})

	URL, _ := url.Parse(fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerName))
	containerURL := azblob.NewContainerURL(*URL, p)

	sr := &StorageReader{
		container: &containerURL,
	}

	return sr, nil
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
				logp.Debug("Found matching blob: %v", blobInfo.Name)
				blobItems = append(blobItems, NewBlobDetails(blobInfo.Name, string(blobInfo.Properties.Etag), *blobInfo.Properties.ContentLength, lastModified))
			}
		}
	}

	return blobItems
}

// ReadBlobData - Reads blob from specified starting location
func (sr *StorageReader) ReadBlobData(path string, startIndex int64) string {

	logp.Info("Reading blob data")

	ctx := context.Background()

	blobURL := sr.container.NewBlockBlobURL(path)
	downloadResponse, err := blobURL.Download(ctx, startIndex, azblob.CountToEnd, azblob.BlobAccessConditions{}, false)
	bodyStream := downloadResponse.Body(azblob.RetryReaderOptions{MaxRetryRequests: 10})

	downloadedData := bytes.Buffer{}
	_, err = downloadedData.ReadFrom(bodyStream)
	if err != nil {
		panic(err)
	}

	return downloadedData.String()
}
