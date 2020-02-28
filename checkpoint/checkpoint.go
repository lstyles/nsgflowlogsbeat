package checkpoint

// Checkpoint - represents a row in Azure Table
type Checkpoint struct {
	PartitionKey string
	RowKey       string
	ETag         string
	Index        int64
	Length       int64
}

func NewCheckpoint(partitionKey, rowKey string) *Checkpoint {
	return &Checkpoint{
		PartitionKey: partitionKey,
		RowKey:       rowKey,
		ETag:         "",
		Index:        0,
		Length:       0,
	}
}
