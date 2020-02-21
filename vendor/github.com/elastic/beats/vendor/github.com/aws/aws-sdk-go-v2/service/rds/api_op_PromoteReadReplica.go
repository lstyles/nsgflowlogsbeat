// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/PromoteReadReplicaMessage
type PromoteReadReplicaInput struct {
	_ struct{} `type:"structure"`

	// The number of days to retain automated backups. Setting this parameter to
	// a positive number enables backups. Setting this parameter to 0 disables automated
	// backups.
	//
	// Default: 1
	//
	// Constraints:
	//
	//    * Must be a value from 0 to 8
	BackupRetentionPeriod *int64 `type:"integer"`

	// The DB instance identifier. This value is stored as a lowercase string.
	//
	// Constraints:
	//
	//    * Must match the identifier of an existing Read Replica DB instance.
	//
	// Example: mydbinstance
	//
	// DBInstanceIdentifier is a required field
	DBInstanceIdentifier *string `type:"string" required:"true"`

	// The daily time range during which automated backups are created if automated
	// backups are enabled, using the BackupRetentionPeriod parameter.
	//
	// The default is a 30-minute window selected at random from an 8-hour block
	// of time for each AWS Region. To see the time blocks available, see Adjusting
	// the Preferred Maintenance Window (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/AdjustingTheMaintenanceWindow.html)
	// in the Amazon RDS User Guide.
	//
	// Constraints:
	//
	//    * Must be in the format hh24:mi-hh24:mi.
	//
	//    * Must be in Universal Coordinated Time (UTC).
	//
	//    * Must not conflict with the preferred maintenance window.
	//
	//    * Must be at least 30 minutes.
	PreferredBackupWindow *string `type:"string"`
}

// String returns the string representation
func (s PromoteReadReplicaInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PromoteReadReplicaInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PromoteReadReplicaInput"}

	if s.DBInstanceIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBInstanceIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/PromoteReadReplicaResult
type PromoteReadReplicaOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details of an Amazon RDS DB instance.
	//
	// This data type is used as a response element in the DescribeDBInstances action.
	DBInstance *DBInstance `type:"structure"`
}

// String returns the string representation
func (s PromoteReadReplicaOutput) String() string {
	return awsutil.Prettify(s)
}

const opPromoteReadReplica = "PromoteReadReplica"

// PromoteReadReplicaRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Promotes a Read Replica DB instance to a standalone DB instance.
//
//    * Backup duration is a function of the amount of changes to the database
//    since the previous backup. If you plan to promote a Read Replica to a
//    standalone instance, we recommend that you enable backups and complete
//    at least one backup prior to promotion. In addition, a Read Replica cannot
//    be promoted to a standalone instance when it is in the backing-up status.
//    If you have enabled backups on your Read Replica, configure the automated
//    backup window so that daily backups do not interfere with Read Replica
//    promotion.
//
//    * This command doesn't apply to Aurora MySQL and Aurora PostgreSQL.
//
//    // Example sending a request using PromoteReadReplicaRequest.
//    req := client.PromoteReadReplicaRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/PromoteReadReplica
func (c *Client) PromoteReadReplicaRequest(input *PromoteReadReplicaInput) PromoteReadReplicaRequest {
	op := &aws.Operation{
		Name:       opPromoteReadReplica,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PromoteReadReplicaInput{}
	}

	req := c.newRequest(op, input, &PromoteReadReplicaOutput{})
	return PromoteReadReplicaRequest{Request: req, Input: input, Copy: c.PromoteReadReplicaRequest}
}

// PromoteReadReplicaRequest is the request type for the
// PromoteReadReplica API operation.
type PromoteReadReplicaRequest struct {
	*aws.Request
	Input *PromoteReadReplicaInput
	Copy  func(*PromoteReadReplicaInput) PromoteReadReplicaRequest
}

// Send marshals and sends the PromoteReadReplica API request.
func (r PromoteReadReplicaRequest) Send(ctx context.Context) (*PromoteReadReplicaResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PromoteReadReplicaResponse{
		PromoteReadReplicaOutput: r.Request.Data.(*PromoteReadReplicaOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PromoteReadReplicaResponse is the response type for the
// PromoteReadReplica API operation.
type PromoteReadReplicaResponse struct {
	*PromoteReadReplicaOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PromoteReadReplica request.
func (r *PromoteReadReplicaResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
