// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/UnmonitorInstancesRequest
type UnmonitorInstancesInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The IDs of the instances.
	//
	// InstanceIds is a required field
	InstanceIds []string `locationName:"InstanceId" locationNameList:"InstanceId" type:"list" required:"true"`
}

// String returns the string representation
func (s UnmonitorInstancesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UnmonitorInstancesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UnmonitorInstancesInput"}

	if s.InstanceIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/UnmonitorInstancesResult
type UnmonitorInstancesOutput struct {
	_ struct{} `type:"structure"`

	// The monitoring information.
	InstanceMonitorings []InstanceMonitoring `locationName:"instancesSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s UnmonitorInstancesOutput) String() string {
	return awsutil.Prettify(s)
}

const opUnmonitorInstances = "UnmonitorInstances"

// UnmonitorInstancesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Disables detailed monitoring for a running instance. For more information,
// see Monitoring Your Instances and Volumes (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-cloudwatch.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using UnmonitorInstancesRequest.
//    req := client.UnmonitorInstancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/UnmonitorInstances
func (c *Client) UnmonitorInstancesRequest(input *UnmonitorInstancesInput) UnmonitorInstancesRequest {
	op := &aws.Operation{
		Name:       opUnmonitorInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UnmonitorInstancesInput{}
	}

	req := c.newRequest(op, input, &UnmonitorInstancesOutput{})
	return UnmonitorInstancesRequest{Request: req, Input: input, Copy: c.UnmonitorInstancesRequest}
}

// UnmonitorInstancesRequest is the request type for the
// UnmonitorInstances API operation.
type UnmonitorInstancesRequest struct {
	*aws.Request
	Input *UnmonitorInstancesInput
	Copy  func(*UnmonitorInstancesInput) UnmonitorInstancesRequest
}

// Send marshals and sends the UnmonitorInstances API request.
func (r UnmonitorInstancesRequest) Send(ctx context.Context) (*UnmonitorInstancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UnmonitorInstancesResponse{
		UnmonitorInstancesOutput: r.Request.Data.(*UnmonitorInstancesOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UnmonitorInstancesResponse is the response type for the
// UnmonitorInstances API operation.
type UnmonitorInstancesResponse struct {
	*UnmonitorInstancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UnmonitorInstances request.
func (r *UnmonitorInstancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
