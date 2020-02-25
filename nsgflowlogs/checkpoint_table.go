package nsgflowlogs

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Azure/azure-sdk-for-go/storage"
	"github.com/elastic/beats/libbeat/logp"
)

// CheckpointsTable - Azure Table storage for checkpoints
type CheckpointsTable struct {
	table   *storage.Table
	timeout uint
}

// NewCheckpointsTable - Initializes new instance of CheckpointsTable
func NewCheckpointsTable(accountName, accountKey, checkpointsTableName string, timeout uint) (*CheckpointsTable, error) {

	logp.Info("Initializing checkpoints table %s on account %s", checkpointsTableName, accountName)

	client, err := storage.NewBasicClient(accountName, accountKey)
	if err != nil {
		return nil, err
	}

	tableService := client.GetTableService()

	table := tableService.GetTableReference(checkpointsTableName)
	ct := &CheckpointsTable{
		table:   table,
		timeout: timeout,
	}

	exists, err := ct.tableExists()
	if err != nil {
		logp.Error(err)
		return nil, err
	}

	if !exists {
		err := ct.createTable()
		if err != nil {
			return nil, err
		}
	}

	logp.Info("Successfully initialized checkpoints table %s", ct.table.Name)

	return ct, nil
}

// CreateOrUpdateCheckpoint - Creates or updates checkpoint in checkpoints table
func (ct *CheckpointsTable) CreateOrUpdateCheckpoint(checkpoint *Checkpoint) error {

	logp.Debug(
		"checkpoint_table", "Creating or updating checkpoint. PartitionKey: %s, RowKey: %s, ETag: %s, Length: %v, Index: %v.",
		checkpoint.PartitionKey,
		checkpoint.RowKey,
		checkpoint.ETag,
		checkpoint.Length,
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
	m["Length"] = checkpoint.Length
	m["Index"] = checkpoint.Index

	entity.Properties = m

	tableBatch.InsertOrReplaceEntityByForce(entity)
	err := tableBatch.ExecuteBatch()
	if err != nil {
		return err
	}

	return nil
}

// GetCheckpoint - Gets checkpoint from checkpoints table
func (ct *CheckpointsTable) GetCheckpoint(partitionKey, rowKey string) (*Checkpoint, error) {

	logp.Debug(
		"checkpoint_table", "Retrieving checkpoint. PartitionKey: %s, RowKey: %s.",
		partitionKey,
		rowKey,
	)

	options := storage.QueryOptions{
		Top:    1,
		Filter: ct.getQueryFilter(partitionKey, rowKey),
	}

	result, err := ct.table.QueryEntities(ct.timeout, storage.NoMetadata, &options)
	if err != nil {
		return nil, err
	}

	if len(result.Entities) != 0 {
		c := result.Entities[0]

		length, _ := strconv.ParseInt(c.Properties["Length"].(string), 10, 64)
		index, _ := strconv.ParseInt(c.Properties["Index"].(string), 10, 64)

		r := &Checkpoint{
			PartitionKey: c.PartitionKey,
			RowKey:       c.RowKey,
			ETag:         c.Properties["ETag"].(string),
			Length:       length,
			Index:        index,
		}
		return r, nil
	}

	return nil, nil
}

func (ct *CheckpointsTable) getQueryFilter(partitionKey, rowKey string) string {

	return fmt.Sprintf("PartitionKey eq '%s' and RowKey eq '%s'", partitionKey, rowKey)
}

func (ct *CheckpointsTable) createTable() error {

	err := ct.table.Create(ct.timeout, storage.NoMetadata, nil)
	if err != nil {
		return err
	}

	return nil
}

func (ct *CheckpointsTable) tableExists() (bool, error) {

	logp.Debug("checkpoint_table", "Checking if checkpoints table already exists.")
	err := ct.table.Get(ct.timeout, storage.NoMetadata)
	if err != nil {
		logp.Debug("checkpoint_table", "Received error from azure servie")
		logp.Error(err)
		switch err.(type) {
		case storage.AzureStorageServiceError:

			if e := err.(storage.AzureStorageServiceError); e.StatusCode == 404 {
				logp.Debug("checkpoint_table", "Service returned 404. Checkpoints table doesn't exist")
				return false, nil
			}
		default:
			logp.Error(err)
			return true, err
		}
	}

	logp.Debug("checkpoint_table", "Checkpoints table exists")
	return true, nil
}
