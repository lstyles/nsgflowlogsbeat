// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
)

// Contains the parameters for ResetImageAttribute.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ResetImageAttributeRequest
type ResetImageAttributeInput struct {
	_ struct{} `type:"structure"`

	// The attribute to reset (currently you can only reset the launch permission
	// attribute).
	//
	// Attribute is a required field
	Attribute ResetImageAttributeName `type:"string" required:"true" enum:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the AMI.
	//
	// ImageId is a required field
	ImageId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ResetImageAttributeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ResetImageAttributeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ResetImageAttributeInput"}
	if len(s.Attribute) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Attribute"))
	}

	if s.ImageId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ImageId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ResetImageAttributeOutput
type ResetImageAttributeOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ResetImageAttributeOutput) String() string {
	return awsutil.Prettify(s)
}

const opResetImageAttribute = "ResetImageAttribute"

// ResetImageAttributeRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Resets an attribute of an AMI to its default value.
//
// The productCodes attribute can't be reset.
//
//    // Example sending a request using ResetImageAttributeRequest.
//    req := client.ResetImageAttributeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ResetImageAttribute
func (c *Client) ResetImageAttributeRequest(input *ResetImageAttributeInput) ResetImageAttributeRequest {
	op := &aws.Operation{
		Name:       opResetImageAttribute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ResetImageAttributeInput{}
	}

	req := c.newRequest(op, input, &ResetImageAttributeOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return ResetImageAttributeRequest{Request: req, Input: input, Copy: c.ResetImageAttributeRequest}
}

// ResetImageAttributeRequest is the request type for the
// ResetImageAttribute API operation.
type ResetImageAttributeRequest struct {
	*aws.Request
	Input *ResetImageAttributeInput
	Copy  func(*ResetImageAttributeInput) ResetImageAttributeRequest
}

// Send marshals and sends the ResetImageAttribute API request.
func (r ResetImageAttributeRequest) Send(ctx context.Context) (*ResetImageAttributeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ResetImageAttributeResponse{
		ResetImageAttributeOutput: r.Request.Data.(*ResetImageAttributeOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ResetImageAttributeResponse is the response type for the
// ResetImageAttribute API operation.
type ResetImageAttributeResponse struct {
	*ResetImageAttributeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ResetImageAttribute request.
func (r *ResetImageAttributeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
