// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteTransitGatewayRequest
type DeleteTransitGatewayInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the transit gateway.
	//
	// TransitGatewayId is a required field
	TransitGatewayId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteTransitGatewayInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTransitGatewayInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteTransitGatewayInput"}

	if s.TransitGatewayId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TransitGatewayId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteTransitGatewayResult
type DeleteTransitGatewayOutput struct {
	_ struct{} `type:"structure"`

	// Information about the deleted transit gateway.
	TransitGateway *TransitGateway `locationName:"transitGateway" type:"structure"`
}

// String returns the string representation
func (s DeleteTransitGatewayOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteTransitGateway = "DeleteTransitGateway"

// DeleteTransitGatewayRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes the specified transit gateway.
//
//    // Example sending a request using DeleteTransitGatewayRequest.
//    req := client.DeleteTransitGatewayRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteTransitGateway
func (c *Client) DeleteTransitGatewayRequest(input *DeleteTransitGatewayInput) DeleteTransitGatewayRequest {
	op := &aws.Operation{
		Name:       opDeleteTransitGateway,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteTransitGatewayInput{}
	}

	req := c.newRequest(op, input, &DeleteTransitGatewayOutput{})
	return DeleteTransitGatewayRequest{Request: req, Input: input, Copy: c.DeleteTransitGatewayRequest}
}

// DeleteTransitGatewayRequest is the request type for the
// DeleteTransitGateway API operation.
type DeleteTransitGatewayRequest struct {
	*aws.Request
	Input *DeleteTransitGatewayInput
	Copy  func(*DeleteTransitGatewayInput) DeleteTransitGatewayRequest
}

// Send marshals and sends the DeleteTransitGateway API request.
func (r DeleteTransitGatewayRequest) Send(ctx context.Context) (*DeleteTransitGatewayResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteTransitGatewayResponse{
		DeleteTransitGatewayOutput: r.Request.Data.(*DeleteTransitGatewayOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteTransitGatewayResponse is the response type for the
// DeleteTransitGateway API operation.
type DeleteTransitGatewayResponse struct {
	*DeleteTransitGatewayOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteTransitGateway request.
func (r *DeleteTransitGatewayResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
