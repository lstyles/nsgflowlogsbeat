// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ReplaceTransitGatewayRouteRequest
type ReplaceTransitGatewayRouteInput struct {
	_ struct{} `type:"structure"`

	// Indicates whether traffic matching this route is to be dropped.
	Blackhole *bool `type:"boolean"`

	// The CIDR range used for the destination match. Routing decisions are based
	// on the most specific match.
	//
	// DestinationCidrBlock is a required field
	DestinationCidrBlock *string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the attachment.
	TransitGatewayAttachmentId *string `type:"string"`

	// The ID of the route table.
	//
	// TransitGatewayRouteTableId is a required field
	TransitGatewayRouteTableId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ReplaceTransitGatewayRouteInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ReplaceTransitGatewayRouteInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ReplaceTransitGatewayRouteInput"}

	if s.DestinationCidrBlock == nil {
		invalidParams.Add(aws.NewErrParamRequired("DestinationCidrBlock"))
	}

	if s.TransitGatewayRouteTableId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TransitGatewayRouteTableId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ReplaceTransitGatewayRouteResult
type ReplaceTransitGatewayRouteOutput struct {
	_ struct{} `type:"structure"`

	// Information about the modified route.
	Route *TransitGatewayRoute `locationName:"route" type:"structure"`
}

// String returns the string representation
func (s ReplaceTransitGatewayRouteOutput) String() string {
	return awsutil.Prettify(s)
}

const opReplaceTransitGatewayRoute = "ReplaceTransitGatewayRoute"

// ReplaceTransitGatewayRouteRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Replaces the specified route in the specified transit gateway route table.
//
//    // Example sending a request using ReplaceTransitGatewayRouteRequest.
//    req := client.ReplaceTransitGatewayRouteRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ReplaceTransitGatewayRoute
func (c *Client) ReplaceTransitGatewayRouteRequest(input *ReplaceTransitGatewayRouteInput) ReplaceTransitGatewayRouteRequest {
	op := &aws.Operation{
		Name:       opReplaceTransitGatewayRoute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ReplaceTransitGatewayRouteInput{}
	}

	req := c.newRequest(op, input, &ReplaceTransitGatewayRouteOutput{})
	return ReplaceTransitGatewayRouteRequest{Request: req, Input: input, Copy: c.ReplaceTransitGatewayRouteRequest}
}

// ReplaceTransitGatewayRouteRequest is the request type for the
// ReplaceTransitGatewayRoute API operation.
type ReplaceTransitGatewayRouteRequest struct {
	*aws.Request
	Input *ReplaceTransitGatewayRouteInput
	Copy  func(*ReplaceTransitGatewayRouteInput) ReplaceTransitGatewayRouteRequest
}

// Send marshals and sends the ReplaceTransitGatewayRoute API request.
func (r ReplaceTransitGatewayRouteRequest) Send(ctx context.Context) (*ReplaceTransitGatewayRouteResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ReplaceTransitGatewayRouteResponse{
		ReplaceTransitGatewayRouteOutput: r.Request.Data.(*ReplaceTransitGatewayRouteOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ReplaceTransitGatewayRouteResponse is the response type for the
// ReplaceTransitGatewayRoute API operation.
type ReplaceTransitGatewayRouteResponse struct {
	*ReplaceTransitGatewayRouteOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ReplaceTransitGatewayRoute request.
func (r *ReplaceTransitGatewayRouteResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
