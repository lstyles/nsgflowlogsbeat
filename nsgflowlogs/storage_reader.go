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

type StorageReader struct {
	container *azblob.ContainerURL
}

func NewStorageReader(accountName, accountKey, containerName string) *StorageReader {

	logp.Info("Connecting to storage account %s.", accountName)
	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		logp.Error(err)
	}
	p := azblob.NewPipeline(credential, azblob.PipelineOptions{})

	URL, _ := url.Parse(fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerName))
	containerURL := azblob.NewContainerURL(*URL, p)

	sr := &StorageReader{
		container: &containerURL,
	}

	return sr
}

func (sr *StorageReader) ListBlobsModifiedBetween(startTime, endTime int64) []BlobDetails {

	logp.Info("Listing blobs modified between %v and %v", startTime, endTime)

	ctx := context.Background()

	var blobItems []BlobDetails

	for marker := (azblob.Marker{}); marker.NotDone(); {

		listBlob, err := sr.container.ListBlobsFlatSegment(ctx, marker, azblob.ListBlobsSegmentOptions{})
		if err != nil {
			panic(err)
		}

		marker = listBlob.NextMarker

		for _, blobInfo := range listBlob.Segment.BlobItems {

			lastModified := blobInfo.Properties.LastModified.UTC().Unix()
			if lastModified > startTime && lastModified < endTime {
				logp.Debug("Found matching blob: %v", blobInfo.Name)
				blobItems = append(blobItems, NewBlobDetails(blobInfo.Name, blobInfo.Properties.Etag, lastModified))
			}
		}
	}

	return blobItems
}

func (sr *StorageReader) ReadBlobData(path string, startIndex int64) string {

	logp.Info("Reading blob data")

	ctx := context.Background()

	blobUrl := sr.container.NewBlockBlobURL(path)
	downloadResponse, err := blobUrl.Download(ctx, startIndex, azblob.CountToEnd, azblob.BlobAccessConditions{}, false)
	bodyStream := downloadResponse.Body(azblob.RetryReaderOptions{MaxRetryRequests: 10})

	downloadedData := bytes.Buffer{}
	_, err = downloadedData.ReadFrom(bodyStream)
	if err != nil {
		panic(err)
	}

	return downloadedData.String()
}

type BlobDetails struct {
	Name         string
	PartitionKey string
	RowKey       string
	ETag         azblob.ETag
	LastModified int64
}

func NewBlobDetails(name string, etag azblob.ETag, lastModified int64) BlobDetails {

	parts := strings.Split(name, "/")

	subscriptionId := parts[2]
	rgName := parts[4]
	nsgName := parts[8]
	year := strings.Split(parts[9], "=")[1]
	month := strings.Split(parts[10], "=")[1]
	day := strings.Split(parts[11], "=")[1]
	hour := strings.Split(parts[12], "=")[1]
	minute := strings.Split(parts[13], "=")[1]
	mac := strings.Split(parts[14], "=")[1]

	bd := BlobDetails{
		Name:         name,
		PartitionKey: fmt.Sprintf("%s_%s_%s_%s", strings.Replace(subscriptionId, "-", "_", 0), rgName, nsgName, mac),
		RowKey:       fmt.Sprintf("%s_%s_%s_%s_%s", year, month, day, hour, minute),
		ETag:         etag,
		LastModified: lastModified,
	}

	return bd
}
