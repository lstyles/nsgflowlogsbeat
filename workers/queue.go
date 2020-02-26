package workers

type ReaderQueueItem struct {
	Name   string
	ETag   string
	Index  int64
	Length int64
}

type ProcessorQueueItem struct {
	Name  string
	ETag  string
	Index int64
	Data  *[]byte
}
