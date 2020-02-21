// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/RebootDBInstanceMessage
type RebootDBInstanceInput struct {
	_ struct{} `type:"structure"`

	// The DB instance identifier. This parameter is stored as a lowercase string.
	//
	// Constraints:
	//
	//    * Must match the identifier of an existing DBInstance.
	//
	// DBInstanceIdentifier is a required field
	DBInstanceIdentifier *string `type:"string" required:"true"`

	// When true, the reboot is conducted through a MultiAZ failover.
	//
	// Constraint: You can't specify true if the instance is not configured for
	// MultiAZ.
	ForceFailover *bool `type:"boolean"`
}

// String returns the string representation
func (s RebootDBInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RebootDBInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RebootDBInstanceInput"}

	if s.DBInstanceIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBInstanceIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/RebootDBInstanceResult
type RebootDBInstanceOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details of an Amazon RDS DB instance.
	//
	// This data type is used as a response element in the DescribeDBInstances action.
	DBInstance *DBInstance `type:"structure"`
}

// String returns the string representation
func (s RebootDBInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

const opRebootDBInstance = "RebootDBInstance"

// RebootDBInstanceRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// You might need to reboot your DB instance, usually for maintenance reasons.
// For example, if you make certain modifications, or if you change the DB parameter
// group associated with the DB instance, you must reboot the instance for the
// changes to take effect.
//
// Rebooting a DB instance restarts the database engine service. Rebooting a
// DB instance results in a momentary outage, during which the DB instance status
// is set to rebooting.
//
// For more information about rebooting, see Rebooting a DB Instance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_RebootInstance.html)
// in the Amazon RDS User Guide.
//
//    // Example sending a request using RebootDBInstanceRequest.
//    req := client.RebootDBInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/RebootDBInstance
func (c *Client) RebootDBInstanceRequest(input *RebootDBInstanceInput) RebootDBInstanceRequest {
	op := &aws.Operation{
		Name:       opRebootDBInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RebootDBInstanceInput{}
	}

	req := c.newRequest(op, input, &RebootDBInstanceOutput{})
	return RebootDBInstanceRequest{Request: req, Input: input, Copy: c.RebootDBInstanceRequest}
}

// RebootDBInstanceRequest is the request type for the
// RebootDBInstance API operation.
type RebootDBInstanceRequest struct {
	*aws.Request
	Input *RebootDBInstanceInput
	Copy  func(*RebootDBInstanceInput) RebootDBInstanceRequest
}

// Send marshals and sends the RebootDBInstance API request.
func (r RebootDBInstanceRequest) Send(ctx context.Context) (*RebootDBInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RebootDBInstanceResponse{
		RebootDBInstanceOutput: r.Request.Data.(*RebootDBInstanceOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RebootDBInstanceResponse is the response type for the
// RebootDBInstance API operation.
type RebootDBInstanceResponse struct {
	*RebootDBInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RebootDBInstance request.
func (r *RebootDBInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
