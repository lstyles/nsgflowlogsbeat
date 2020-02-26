package checkpoint

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/Azure/azure-sdk-for-go/storage"
	"github.com/elastic/beats/libbeat/logp"
)

// Table interacts with Azure Table to manage state checkpoints.
type Table struct {
	table  *storage.Table
	config *Config
}

// Config contains information required to connect to Azure Table
type Config struct {
	accountName string
	accountKey  string
	tableName   string
	timeout     uint
}

// NewCheckpointTable creates a new instance of CheckpointsTable.
// During initialization the table will be created if it doesn't already exist.
func NewCheckpointTable(accountName, accountKey, checkpointsTableName string, timeout uint) (*Table, error) {

	logp.Debug(
		"checkpoint_table",
		"Creating new instance of checkpoint table",
	)

	if len(accountName) == 0 {
		return nil, errors.New("account name is required")
	}

	if len(accountKey) == 0 {
		return nil, errors.New("account key is required")
	}

	if len(checkpointsTableName) == 0 {
		return nil, errors.New("checkpoints table name is required")
	}

	c := &Config{
		accountName: accountName,
		accountKey:  accountKey,
		tableName:   checkpointsTableName,
		timeout:     timeout,
	}

	t := &Table{
		config: c,
	}

	err := t.initialize()
	if err != nil {
		return nil, err
	}

	return t, nil
}

// GetCheckpoint gets checkpoint from checkpoints table
func (ct *Table) GetCheckpoint(partitionKey, rowKey string) (*Checkpoint, error) {

	logp.Debug(
		"checkpoint_table", "Retrieving checkpoint. PartitionKey: %s, RowKey: %s.",
		partitionKey,
		rowKey,
	)

	options := storage.QueryOptions{
		Top:    1,
		Filter: ct.getQueryFilter(partitionKey, rowKey),
	}

	result, err := ct.table.QueryEntities(ct.config.timeout, storage.NoMetadata, &options)
	if err != nil {
		return nil, err
	}

	if len(result.Entities) != 0 {
		c := result.Entities[0]

		index, _ := strconv.ParseInt(c.Properties["Index"].(string), 10, 64)

		r := &Checkpoint{
			PartitionKey: c.PartitionKey,
			RowKey:       c.RowKey,
			ETag:         c.Properties["ETag"].(string),
			Index:        index,
		}
		return r, nil
	}

	return nil, nil
}

// CreateOrUpdateCheckpoint creates or updates checkpoint in checkpoints table
func (ct *Table) CreateOrUpdateCheckpoint(checkpoint *Checkpoint) error {

	logp.Debug(
		"checkpoint_table", "Creating or updating checkpoint. PartitionKey: %s, RowKey: %s, ETag: %s, Index: %v.",
		checkpoint.PartitionKey,
		checkpoint.RowKey,
		checkpoint.ETag,
		checkpoint.Index,
	)

	tableBatch := ct.table.NewBatch()
	entity := &storage.Entity{
		Table:        ct.table,
		PartitionKey: checkpoint.PartitionKey,
		RowKey:       checkpoint.RowKey,
		TimeStamp:    time.Now(),
	}

	m := make(map[string]interface{})
	m["ETag"] = checkpoint.ETag
	m["Index"] = checkpoint.Index

	entity.Properties = m

	tableBatch.InsertOrReplaceEntityByForce(entity)
	err := tableBatch.ExecuteBatch()
	if err != nil {
		return err
	}

	return nil
}

func (ct *Table) initialize() error {

	config := *ct.config

	logp.Debug(
		"checkpoint_table",
		"Initializing checkpoint table %s in account %s. Timeout set to %vs.",
		config.tableName,
		config.accountName,
		config.timeout,
	)

	client, err := storage.NewBasicClient(config.accountName, config.accountKey)
	if err != nil {
		return err
	}

	tableService := client.GetTableService()

	ct.table = tableService.GetTableReference(config.tableName)
	err = ct.ensureCreated()
	if err != nil {
		return err
	}

	return nil
}

func (ct *Table) ensureCreated() error {

	logp.Debug(
		"checkpoint_table",
		"Checking whether checkpoints table exists in the storage account and creating it if not.",
	)

	exists, err := ct.tableExists()
	if err != nil {
		logp.Error(err)
		return err
	}

	if !exists {
		err := ct.createTable()
		if err != nil {
			return err
		}
	}

	return nil
}

func (ct *Table) tableExists() (bool, error) {

	logp.Debug(
		"checkpoint_table",
		"Checking if checkpoints table already exists in the storage account.",
	)

	err := ct.table.Get(ct.config.timeout, storage.NoMetadata)
	if err != nil {
		switch err.(type) {
		case storage.AzureStorageServiceError:

			if e := err.(storage.AzureStorageServiceError); e.StatusCode == 404 {
				logp.Warn("Service returned 404. Checkpoints table doesn't exist. Will create.")
				return false, nil
			}
		default:
			logp.Error(err)
			return true, err
		}
	}

	return true, nil
}

func (ct *Table) createTable() error {

	logp.Debug(
		"checkpoint_table",
		"Creating new table to store checkpoints.",
	)
	err := ct.table.Create(ct.config.timeout, storage.NoMetadata, nil)
	if err != nil {
		return err
	}

	return nil
}

func (ct *Table) getQueryFilter(partitionKey, rowKey string) string {

	return fmt.Sprintf("PartitionKey eq '%s' and RowKey eq '%s'", partitionKey, rowKey)
}
