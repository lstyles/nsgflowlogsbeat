package checkpoint

import (
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
	} else {
		// Checkpoint doesn't exist, create and return
		c := NewCheckpoint(partitionKey, rowKey)

		return c, nil
	}
}

// CreateOrUpdateCheckpoint creates or updates checkpoint in checkpoints table
func (ct *Table) CreateOrUpdateCheckpoint(checkpoint *Checkpoint) error {

	logp.Debug(
		"checkpoint_table", "Creating or updating checkpoint. PartitionKey: %s, RowKey: %s, ETag: %s, Index: %v.",
		checkpoint.PartitionKey,
		checkpoint.RowKey,
		checkpoint.ETag,
		checkpoint.Index,
		checkpoint.Length,
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
	m["Length"] = checkpoint.Length

	entity.Properties = m

	tableBatch.InsertOrReplaceEntityByForce(entity)
	err := tableBatch.ExecuteBatch()
	if err != nil {
		return err
	}

	return nil
}

func (ct *Table) UpdateCheckpoint(partitionKey, rowKey, etag string, index int64) {

	logp.Info("Updating checkpoint Partition Key: %s, Row Key: %s, ETag: %s, Index: %d", partitionKey, rowKey, etag, index)
	c, err := ct.GetCheckpoint(partitionKey, rowKey)
	if err != nil {
		logp.Error(err)
	}
	if c.Index == 0 {
		logp.Warn("This should never happen")
	}

	c.ETag = etag
	c.Index = index
	c.Length = index

	ct.CreateOrUpdateCheckpoint(c)
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
