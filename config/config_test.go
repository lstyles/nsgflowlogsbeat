// +build !integration

package config

import (
	"testing"

	"github.com/elastic/beats/libbeat/common"
	"github.com/stretchr/testify/assert"
)

type validationTestCase struct {
	config Config
	errMsg string
}

func (v validationTestCase) run(t *testing.T) {
	if v.errMsg == "" {
		assert.NoError(t, v.config.Validate())
	} else {
		err := v.config.Validate()
		if err != nil {
			assert.Contains(t, err.Error(), v.errMsg)
		} else {
			t.Errorf("expected error with '%s'", v.errMsg)
		}
	}
}

func TestConfigValidate(t *testing.T) {
	testCases := []validationTestCase{
		// Top-level config
		{
			Config{
				StorageAccountName:   "<storage_account_name>",
				StorageAccountKey:    "<storage_account_key>",
				ContainerName:        "<container_name>",
				CheckpointsTableName: "<checkpoints_table_name>",
			},
			"", // No Error
		},
		{
			Config{},
			"4 errors: account name is required; account key is required; checkpoints table name is required; container name is required",
		},
	}

	for _, test := range testCases {
		test.run(t)
	}
}

func newConfig(from map[string]interface{}) *common.Config {
	cfg, err := common.NewConfigFrom(from)
	if err != nil {
		panic(err)
	}
	return cfg
}
