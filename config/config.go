// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import (
	"fmt"
	"time"

	"github.com/joeshaw/multierror"
)

// Config - all configuration settings
type Config struct {
	ScanFrequency           time.Duration `config:"scan_frequency"`
	StorageAccountName      string        `config:"storage_account_name"`
	StorageAccountKey       string        `config:"storage_account_key"`
	ContainerName           string        `config:"container_name"`
	CheckpointsTableName    string        `config:"checkpoints_table_name"`
	CheckpointsTableTimeout uint          `config:"checkpoints_table_timeout"`
	IgnoreOlder             time.Duration `config:"ignore_older"`
	Workers                 int           `config:"workers"`
}

// DefaultConfig - default configuration settings
var DefaultConfig = Config{
	ScanFrequency:           30 * time.Second,
	ContainerName:           "insights-logs-networksecuritygroupflowevent",
	CheckpointsTableName:    "nsgflowlogsbeat_checkpoints",
	CheckpointsTableTimeout: 15,
	IgnoreOlder:             10 * time.Minute,
	Workers:                 4,
}

// Validate validates the configuration and returns an error describing all problems or nil if there are none
func (cfg Config) Validate() error {
	var errs multierror.Errors

	if len(cfg.StorageAccountName) == 0 {
		errs = append(errs, fmt.Errorf("account name is required"))
	}

	if len(cfg.StorageAccountKey) == 0 {
		errs = append(errs, fmt.Errorf("account key is required"))
	}

	if len(cfg.CheckpointsTableName) == 0 {
		errs = append(errs, fmt.Errorf("checkpoints table name is required"))
	}

	if len(cfg.ContainerName) == 0 {
		errs = append(errs, fmt.Errorf("container name is required"))
	}

	return errs.Err()
}
