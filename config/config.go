// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type Config struct {
	ScanFrequency           time.Duration `config:"scan_frequency"`
	StorageAccountName      string        `config:"storage_account_name"`
	StorageAccountKey       string        `config:"storage_account_key"`
	ContainerName           string        `config:"container_name"`
	CheckpointsTableName    string        `config:"checkpoints_table_name"`
	CheckpointsTableTimeout uint          `config:"checkpoints_table_timeout"`
	IgnoreOlder             time.Duration `config:"ignore_older"`
}

var DefaultConfig = Config{
	ScanFrequency:           30 * time.Second,
	ContainerName:           "insights-logs-networksecuritygroupflowevent",
	CheckpointsTableName:    "nsgflowlogsbeat_checkpoints",
	CheckpointsTableTimeout: 15,
	IgnoreOlder:             0,
}
