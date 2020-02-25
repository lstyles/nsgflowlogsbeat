// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

// Config - all configuration settings
type Config struct {
	ScanFrequency           time.Duration `config:"scan_frequency"`
	StorageAccountName      string        `config:"storage_account_name"`
	StorageAccountKey       string        `config:"storage_account_key"`
	StorageReaderWorkers    int           `config:"storage_reader_workers"`
	ContainerName           string        `config:"container_name"`
	CheckpointsTableName    string        `config:"checkpoints_table_name"`
	CheckpointsTableTimeout uint          `config:"checkpoints_table_timeout"`
	IgnoreOlder             time.Duration `config:"ignore_older"`
	MessageProcessorWorkers int           `config:"message_processor_workers"`
}

// DefaultConfig - default configuration settings
var DefaultConfig = Config{
	ScanFrequency:           30 * time.Second,
	StorageReaderWorkers:    2,
	ContainerName:           "insights-logs-networksecuritygroupflowevent",
	CheckpointsTableName:    "nsgflowlogsbeat_checkpoints",
	CheckpointsTableTimeout: 15,
	IgnoreOlder:             10 * time.Second,
	MessageProcessorWorkers: 4,
}
