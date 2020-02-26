package checkpoint

// Checkpoint - represents a row in Azure Table
type Checkpoint struct {
	PartitionKey string
	RowKey       string
	ETag         string
	Index        int64
}
