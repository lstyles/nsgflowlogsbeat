package nsgflowlogs

type ReaderQueueItem struct {
	Name         string
	PartitionKey string
	RowKey       string
	ETag         string
	Index        int64
	Length       int64
}
