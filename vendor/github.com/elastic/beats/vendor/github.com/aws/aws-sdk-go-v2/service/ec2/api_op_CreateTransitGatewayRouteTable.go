// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateTransitGatewayRouteTableRequest
type CreateTransitGatewayRouteTableInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The tags to apply to the transit gateway route table.
	TagSpecifications []TagSpecification `locationNameList:"item" type:"list"`

	// The ID of the transit gateway.
	//
	// TransitGatewayId is a required field
	TransitGatewayId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateTransitGatewayRouteTableInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateTransitGatewayRouteTableInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateTransitGatewayRouteTableInput"}

	if s.TransitGatewayId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TransitGatewayId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateTransitGatewayRouteTableResult
type CreateTransitGatewayRouteTableOutput struct {
	_ struct{} `type:"structure"`

	// Information about the transit gateway route table.
	TransitGatewayRouteTable *TransitGatewayRouteTable `locationName:"transitGatewayRouteTable" type:"structure"`
}

// String returns the string representation
func (s CreateTransitGatewayRouteTableOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateTransitGatewayRouteTable = "CreateTransitGatewayRouteTable"

// CreateTransitGatewayRouteTableRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates a route table for the specified transit gateway.
//
//    // Example sending a request using CreateTransitGatewayRouteTableRequest.
//    req := client.CreateTransitGatewayRouteTableRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateTransitGatewayRouteTable
func (c *Client) CreateTransitGatewayRouteTableRequest(input *CreateTransitGatewayRouteTableInput) CreateTransitGatewayRouteTableRequest {
	op := &aws.Operation{
		Name:       opCreateTransitGatewayRouteTable,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateTransitGatewayRouteTableInput{}
	}

	req := c.newRequest(op, input, &CreateTransitGatewayRouteTableOutput{})
	return CreateTransitGatewayRouteTableRequest{Request: req, Input: input, Copy: c.CreateTransitGatewayRouteTableRequest}
}

// CreateTransitGatewayRouteTableRequest is the request type for the
// CreateTransitGatewayRouteTable API operation.
type CreateTransitGatewayRouteTableRequest struct {
	*aws.Request
	Input *CreateTransitGatewayRouteTableInput
	Copy  func(*CreateTransitGatewayRouteTableInput) CreateTransitGatewayRouteTableRequest
}

// Send marshals and sends the CreateTransitGatewayRouteTable API request.
func (r CreateTransitGatewayRouteTableRequest) Send(ctx context.Context) (*CreateTransitGatewayRouteTableResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateTransitGatewayRouteTableResponse{
		CreateTransitGatewayRouteTableOutput: r.Request.Data.(*CreateTransitGatewayRouteTableOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateTransitGatewayRouteTableResponse is the response type for the
// CreateTransitGatewayRouteTable API operation.
type CreateTransitGatewayRouteTableResponse struct {
	*CreateTransitGatewayRouteTableOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateTransitGatewayRouteTable request.
func (r *CreateTransitGatewayRouteTableResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
