package workers

type ReaderQueueItem struct {
	Name   string
	ETag   string
	Index  int64
	Length int64
}
