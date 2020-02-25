package nsgflowlogs

// Checkpoint - represents a row in Azure Table
type Checkpoint struct {
	PartitionKey string
	RowKey       string
	ETag         string
	Length       int64
	Index        int64
}

// NewCheckpoint - Creates a new instance of Checkpoint
func NewCheckpoint(partitionKey, rowKey, etag string, length, index int64) Checkpoint {

	c := Checkpoint{
		PartitionKey: partitionKey,
		RowKey:       rowKey,
		ETag:         etag,
		Length:       length,
		Index:        index,
	}

	return c
}
