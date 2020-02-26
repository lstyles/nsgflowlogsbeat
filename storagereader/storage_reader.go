package storagereader

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/elastic/beats/libbeat/logp"
)

// StorageReader interacts with Azure Blob storage
type StorageReader struct {
	container *azblob.ContainerURL
	config    *Config
}

// Config holds storage reader configuration
type Config struct {
	accountName   string
	accountKey    string
	containerName string
}

// NewStorageReader creates a new instance of StorageReader
func NewStorageReader(accountName, accountKey, containerName string) (*StorageReader, error) {

	logp.Debug(
		"storage_reader",
		"Creating new instance of storage reader",
	)

	if len(accountName) == 0 {
		return nil, errors.New("account name is required")
	}

	if len(accountKey) == 0 {
		return nil, errors.New("account key is required")
	}

	if len(containerName) == 0 {
		return nil, errors.New("checkpoints table name is required")
	}

	c := &Config{
		accountName:   accountName,
		accountKey:    accountKey,
		containerName: containerName,
	}

	sr := &StorageReader{
		config: c,
	}

	err := sr.initialize()
	if err != nil {
		return nil, err
	}

	return sr, nil
}

func (sr *StorageReader) initialize() error {

	config := *sr.config

	logp.Debug(
		"checkpoint_table",
		"Initializing storage account connection to account %s, container %s.",
		config.accountName,
		config.containerName,
	)

	cred, err := azblob.NewSharedKeyCredential(config.accountName, config.accountKey)
	if err != nil {
		return err
	}

	p := azblob.NewPipeline(cred, azblob.PipelineOptions{})

	URL, _ := url.Parse(fmt.Sprintf("https://%s.blob.core.windows.net/%s", config.accountName, config.containerName))
	containerURL := azblob.NewContainerURL(*URL, p)

	sr.container = &containerURL

	return nil
}

// ListBlobsModifiedBetween - Get list of blobs modified between two specified timestamps
func (sr *StorageReader) ListBlobsModifiedBetween(startTime, endTime int64) *[]BlobDetails {

	logp.Debug("storage_reader", "Listing blobs modified between %v and %v.", startTime, endTime)
	ctx := context.Background()

	var blobItems []BlobDetails

	i := 0
	for marker := (azblob.Marker{}); marker.NotDone(); {

		listBlob, err := sr.container.ListBlobsFlatSegment(ctx, marker, azblob.ListBlobsSegmentOptions{})
		marker = listBlob.NextMarker
		if err != nil {
			logp.Error(err)
			continue
		}
		i++

		for _, blobInfo := range listBlob.Segment.BlobItems {

			lastModified := blobInfo.Properties.LastModified.UTC().Unix()
			if lastModified > startTime && lastModified < endTime {
				length := blobInfo.Properties.ContentLength
				blobItems = append(blobItems, NewBlobDetails(blobInfo.Name, string(blobInfo.Properties.Etag), *length, lastModified))
			}
		}
	}

	logp.Info("Found %v blobs in container. Found %v blobs modified between %v and %v.",
		i,
		len(blobItems),
		startTime,
		endTime,
	)

	return &blobItems
}

// ReadBlobData - Reads blob from specified starting location
func (sr *StorageReader) ReadBlobData(path string, startIndex, endIndex int64) []byte {

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

	return downloadedData.Bytes()
}
