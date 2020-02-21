// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/ModifyDBSnapshotMessage
type ModifyDBSnapshotInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the DB snapshot to modify.
	//
	// DBSnapshotIdentifier is a required field
	DBSnapshotIdentifier *string `type:"string" required:"true"`

	// The engine version to upgrade the DB snapshot to.
	//
	// The following are the database engines and engine versions that are available
	// when you upgrade a DB snapshot.
	//
	// MySQL
	//
	//    * 5.5.46 (supported for 5.1 DB snapshots)
	//
	// Oracle
	//
	//    * 12.1.0.2.v8 (supported for 12.1.0.1 DB snapshots)
	//
	//    * 11.2.0.4.v12 (supported for 11.2.0.2 DB snapshots)
	//
	//    * 11.2.0.4.v11 (supported for 11.2.0.3 DB snapshots)
	EngineVersion *string `type:"string"`

	// The option group to identify with the upgraded DB snapshot.
	//
	// You can specify this parameter when you upgrade an Oracle DB snapshot. The
	// same option group considerations apply when upgrading a DB snapshot as when
	// upgrading a DB instance. For more information, see Option Group Considerations
	// (http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.Oracle.html#USER_UpgradeDBInstance.Oracle.OGPG.OG)
	// in the Amazon RDS User Guide.
	OptionGroupName *string `type:"string"`
}

// String returns the string representation
func (s ModifyDBSnapshotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyDBSnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyDBSnapshotInput"}

	if s.DBSnapshotIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBSnapshotIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/ModifyDBSnapshotResult
type ModifyDBSnapshotOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details of an Amazon RDS DB snapshot.
	//
	// This data type is used as a response element in the DescribeDBSnapshots action.
	DBSnapshot *DBSnapshot `type:"structure"`
}

// String returns the string representation
func (s ModifyDBSnapshotOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyDBSnapshot = "ModifyDBSnapshot"

// ModifyDBSnapshotRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Updates a manual DB snapshot, which can be encrypted or not encrypted, with
// a new engine version.
//
// Amazon RDS supports upgrading DB snapshots for MySQL and Oracle.
//
//    // Example sending a request using ModifyDBSnapshotRequest.
//    req := client.ModifyDBSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/ModifyDBSnapshot
func (c *Client) ModifyDBSnapshotRequest(input *ModifyDBSnapshotInput) ModifyDBSnapshotRequest {
	op := &aws.Operation{
		Name:       opModifyDBSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyDBSnapshotInput{}
	}

	req := c.newRequest(op, input, &ModifyDBSnapshotOutput{})
	return ModifyDBSnapshotRequest{Request: req, Input: input, Copy: c.ModifyDBSnapshotRequest}
}

// ModifyDBSnapshotRequest is the request type for the
// ModifyDBSnapshot API operation.
type ModifyDBSnapshotRequest struct {
	*aws.Request
	Input *ModifyDBSnapshotInput
	Copy  func(*ModifyDBSnapshotInput) ModifyDBSnapshotRequest
}

// Send marshals and sends the ModifyDBSnapshot API request.
func (r ModifyDBSnapshotRequest) Send(ctx context.Context) (*ModifyDBSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyDBSnapshotResponse{
		ModifyDBSnapshotOutput: r.Request.Data.(*ModifyDBSnapshotOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyDBSnapshotResponse is the response type for the
// ModifyDBSnapshot API operation.
type ModifyDBSnapshotResponse struct {
	*ModifyDBSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyDBSnapshot request.
func (r *ModifyDBSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
