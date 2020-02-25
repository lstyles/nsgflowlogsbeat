package nsgflowlogs

import (
	"fmt"
	"strings"
)

// BlobDetails - represents blob list item
type BlobDetails struct {
	Name         string
	PartitionKey string
	RowKey       string
	ETag         string
	Length       int64
	LastModified int64
}

// NewBlobDetails - Creates a new instance of BlobDetails
func NewBlobDetails(name, etag string, length, lastModified int64) BlobDetails {

	parts := strings.Split(name, "/")

	subscriptionID := parts[2]
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
		PartitionKey: fmt.Sprintf("%s_%s_%s_%s_%s_%s", strings.Replace(subscriptionID, "-", "_", 0), rgName, nsgName, year, month, day),
		RowKey:       fmt.Sprintf("%s_%s_%s", hour, minute, mac),
		ETag:         etag,
		Length:       length,
		LastModified: lastModified,
	}

	return bd
}
