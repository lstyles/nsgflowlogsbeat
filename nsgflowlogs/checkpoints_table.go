package nsgflowlogs

import (
	"fmt"
	"strconv"
	"time"

	"github.com/elastic/beats/libbeat/logp"

	"github.com/Azure/azure-sdk-for-go/storage"
)

// CheckpointsTable - Azure Table storage for checkpoints
type CheckpointsTable struct {
	table   *storage.Table
	timeout uint
}

// NewCheckpointsTable - Initializes new instance of CheckpointsTable
func NewCheckpointsTable(accountName, accountKey, checkpointsTableName string, timeout uint) *CheckpointsTable {

	logp.Info("Initializing checkpoints table %s on account %s", checkpointsTableName, accountName)

	client, err := storage.NewBasicClient(accountName, accountKey)
	if err != nil {
		logp.Error(err)
		return nil
	}

	tableService := client.GetTableService()

	table := tableService.GetTableReference(checkpointsTableName)
	err = table.Get(15, storage.NoMetadata)
	if err != nil {
		switch err.(type) {
		case storage.AzureStorageServiceError:

			if e := err.(storage.AzureStorageServiceError); e.StatusCode == 404 {
				logp.Info("Checkpoints table doesn't exist. Creating...")

				ec := table.Create(timeout, storage.NoMetadata, nil)

				if ec != nil {
					panic(ec)
				}
			} else {
				panic(err)
			}
		default:
			panic(err)
		}
	}

	ct := &CheckpointsTable{
		table: table,
	}

	logp.Info("Got a reference to a Azure Table %s", ct.table.Name)

	return ct
}

// GetLastScanTS - Gets last scan timestamp from checkpoints table
func (ct *CheckpointsTable) GetLastScanTS() int64 {
	logp.Info("Getting last scan timestamp")

	result := ct.getRecord("System", "LastScanTS")
	if result != nil {
		r, _ := strconv.ParseInt(result.Properties["Index"].(string), 10, 64)
		return r
	}

	return 0
}

// SetLastScanTS - Sets last scan timestamp in checkpoints table
func (ct *CheckpointsTable) SetLastScanTS(ts int64) error {
	logp.Info("Setting last scan timestamp to %v", ts)

	return ct.createOrUpdateRecord("System", "LastScanTS", "", ts)
}

// GetCheckpoint - Gets checkpoint index for specified blob
func (ct *CheckpointsTable) GetCheckpoint(partitionKey, rowKey string) int64 {
	logp.Info("Getting last scan timestamp")

	result := ct.getRecord(partitionKey, rowKey)

	if result != nil {
		r, _ := strconv.ParseInt(result.Properties["Index"].(string), 10, 64)
		return r
	}

	return 0
}

// SetCheckpoint - Sets checkpoint for specified blob
func (ct *CheckpointsTable) SetCheckpoint(partitionKey, rowKey string, etag string, index int64) error {
	logp.Info("Setting checkpoint for %v %v to %v", partitionKey, rowKey, index)

	return ct.createOrUpdateRecord(partitionKey, rowKey, etag, index)
}

func (ct *CheckpointsTable) getRecord(partitionKey, rowKey string) *storage.Entity {

	logp.Info("Retrieving checkpoint table record. PartionKey: %s. RowKey: %s", partitionKey, rowKey)

	options := storage.QueryOptions{
		Top:    1,
		Filter: fmt.Sprintf("PartitionKey eq '%s' and RowKey eq '%s'", partitionKey, rowKey),
	}

	logp.Debug("OData Query filter: %s", options.Filter)

	result, err := ct.table.QueryEntities(ct.timeout, storage.NoMetadata, &options)
	if err != nil {
		panic(err)
		return nil
	}

	if len(result.Entities) != 0 {
		return result.Entities[0]
	}

	return nil
}

func (ct *CheckpointsTable) createOrUpdateRecord(partitionKey, rowKey string, etag string, index int64) error {

	logp.Info("Updating checkpoints table record. PartitionKey: %s. RowKey: %s, Index: %v", partitionKey, rowKey, index)

	tableBatch := ct.table.NewBatch()
	entity := &storage.Entity{
		Table: ct.table,
	}

	entity.PartitionKey = partitionKey
	entity.RowKey = rowKey
	entity.TimeStamp = time.Now()

	m := make(map[string]interface{})
	m["Index"] = int64(index)
	m["ETag"] = etag

	entity.Properties = m

	tableBatch.InsertOrReplaceEntityByForce(entity)
	err := tableBatch.ExecuteBatch()
	if err != nil {
		logp.Error(err)
		return err
	}

	return nil
}
