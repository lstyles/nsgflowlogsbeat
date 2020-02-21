// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteInternetGatewayRequest
type DeleteInternetGatewayInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the internet gateway.
	//
	// InternetGatewayId is a required field
	InternetGatewayId *string `locationName:"internetGatewayId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteInternetGatewayInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteInternetGatewayInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteInternetGatewayInput"}

	if s.InternetGatewayId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InternetGatewayId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteInternetGatewayOutput
type DeleteInternetGatewayOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteInternetGatewayOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteInternetGateway = "DeleteInternetGateway"

// DeleteInternetGatewayRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes the specified internet gateway. You must detach the internet gateway
// from the VPC before you can delete it.
//
//    // Example sending a request using DeleteInternetGatewayRequest.
//    req := client.DeleteInternetGatewayRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteInternetGateway
func (c *Client) DeleteInternetGatewayRequest(input *DeleteInternetGatewayInput) DeleteInternetGatewayRequest {
	op := &aws.Operation{
		Name:       opDeleteInternetGateway,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteInternetGatewayInput{}
	}

	req := c.newRequest(op, input, &DeleteInternetGatewayOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteInternetGatewayRequest{Request: req, Input: input, Copy: c.DeleteInternetGatewayRequest}
}

// DeleteInternetGatewayRequest is the request type for the
// DeleteInternetGateway API operation.
type DeleteInternetGatewayRequest struct {
	*aws.Request
	Input *DeleteInternetGatewayInput
	Copy  func(*DeleteInternetGatewayInput) DeleteInternetGatewayRequest
}

// Send marshals and sends the DeleteInternetGateway API request.
func (r DeleteInternetGatewayRequest) Send(ctx context.Context) (*DeleteInternetGatewayResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteInternetGatewayResponse{
		DeleteInternetGatewayOutput: r.Request.Data.(*DeleteInternetGatewayOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteInternetGatewayResponse is the response type for the
// DeleteInternetGateway API operation.
type DeleteInternetGatewayResponse struct {
	*DeleteInternetGatewayOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteInternetGateway request.
func (r *DeleteInternetGatewayResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
