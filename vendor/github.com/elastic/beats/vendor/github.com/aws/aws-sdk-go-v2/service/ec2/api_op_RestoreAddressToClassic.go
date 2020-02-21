// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/RestoreAddressToClassicRequest
type RestoreAddressToClassicInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The Elastic IP address.
	//
	// PublicIp is a required field
	PublicIp *string `locationName:"publicIp" type:"string" required:"true"`
}

// String returns the string representation
func (s RestoreAddressToClassicInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RestoreAddressToClassicInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RestoreAddressToClassicInput"}

	if s.PublicIp == nil {
		invalidParams.Add(aws.NewErrParamRequired("PublicIp"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/RestoreAddressToClassicResult
type RestoreAddressToClassicOutput struct {
	_ struct{} `type:"structure"`

	// The Elastic IP address.
	PublicIp *string `locationName:"publicIp" type:"string"`

	// The move status for the IP address.
	Status Status `locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s RestoreAddressToClassicOutput) String() string {
	return awsutil.Prettify(s)
}

const opRestoreAddressToClassic = "RestoreAddressToClassic"

// RestoreAddressToClassicRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Restores an Elastic IP address that was previously moved to the EC2-VPC platform
// back to the EC2-Classic platform. You cannot move an Elastic IP address that
// was originally allocated for use in EC2-VPC. The Elastic IP address must
// not be associated with an instance or network interface.
//
//    // Example sending a request using RestoreAddressToClassicRequest.
//    req := client.RestoreAddressToClassicRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/RestoreAddressToClassic
func (c *Client) RestoreAddressToClassicRequest(input *RestoreAddressToClassicInput) RestoreAddressToClassicRequest {
	op := &aws.Operation{
		Name:       opRestoreAddressToClassic,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RestoreAddressToClassicInput{}
	}

	req := c.newRequest(op, input, &RestoreAddressToClassicOutput{})
	return RestoreAddressToClassicRequest{Request: req, Input: input, Copy: c.RestoreAddressToClassicRequest}
}

// RestoreAddressToClassicRequest is the request type for the
// RestoreAddressToClassic API operation.
type RestoreAddressToClassicRequest struct {
	*aws.Request
	Input *RestoreAddressToClassicInput
	Copy  func(*RestoreAddressToClassicInput) RestoreAddressToClassicRequest
}

// Send marshals and sends the RestoreAddressToClassic API request.
func (r RestoreAddressToClassicRequest) Send(ctx context.Context) (*RestoreAddressToClassicResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RestoreAddressToClassicResponse{
		RestoreAddressToClassicOutput: r.Request.Data.(*RestoreAddressToClassicOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RestoreAddressToClassicResponse is the response type for the
// RestoreAddressToClassic API operation.
type RestoreAddressToClassicResponse struct {
	*RestoreAddressToClassicOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RestoreAddressToClassic request.
func (r *RestoreAddressToClassicResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
