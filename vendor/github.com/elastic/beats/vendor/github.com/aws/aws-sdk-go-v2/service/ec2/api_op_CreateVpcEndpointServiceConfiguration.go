// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateVpcEndpointServiceConfigurationRequest
type CreateVpcEndpointServiceConfigurationInput struct {
	_ struct{} `type:"structure"`

	// Indicate whether requests from service consumers to create an endpoint to
	// your service must be accepted. To accept a request, use AcceptVpcEndpointConnections.
	AcceptanceRequired *bool `type:"boolean"`

	// Unique, case-sensitive identifier you provide to ensure the idempotency of
	// the request. For more information, see How to Ensure Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Run_Instance_Idempotency.html).
	ClientToken *string `type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The Amazon Resource Names (ARNs) of one or more Network Load Balancers for
	// your service.
	//
	// NetworkLoadBalancerArns is a required field
	NetworkLoadBalancerArns []string `locationName:"NetworkLoadBalancerArn" locationNameList:"item" type:"list" required:"true"`
}

// String returns the string representation
func (s CreateVpcEndpointServiceConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateVpcEndpointServiceConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateVpcEndpointServiceConfigurationInput"}

	if s.NetworkLoadBalancerArns == nil {
		invalidParams.Add(aws.NewErrParamRequired("NetworkLoadBalancerArns"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateVpcEndpointServiceConfigurationResult
type CreateVpcEndpointServiceConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier you provide to ensure the idempotency of
	// the request.
	ClientToken *string `locationName:"clientToken" type:"string"`

	// Information about the service configuration.
	ServiceConfiguration *ServiceConfiguration `locationName:"serviceConfiguration" type:"structure"`
}

// String returns the string representation
func (s CreateVpcEndpointServiceConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateVpcEndpointServiceConfiguration = "CreateVpcEndpointServiceConfiguration"

// CreateVpcEndpointServiceConfigurationRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates a VPC endpoint service configuration to which service consumers (AWS
// accounts, IAM users, and IAM roles) can connect. Service consumers can create
// an interface VPC endpoint to connect to your service.
//
// To create an endpoint service configuration, you must first create a Network
// Load Balancer for your service. For more information, see VPC Endpoint Services
// (https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/endpoint-service.html)
// in the Amazon Virtual Private Cloud User Guide.
//
//    // Example sending a request using CreateVpcEndpointServiceConfigurationRequest.
//    req := client.CreateVpcEndpointServiceConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateVpcEndpointServiceConfiguration
func (c *Client) CreateVpcEndpointServiceConfigurationRequest(input *CreateVpcEndpointServiceConfigurationInput) CreateVpcEndpointServiceConfigurationRequest {
	op := &aws.Operation{
		Name:       opCreateVpcEndpointServiceConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateVpcEndpointServiceConfigurationInput{}
	}

	req := c.newRequest(op, input, &CreateVpcEndpointServiceConfigurationOutput{})
	return CreateVpcEndpointServiceConfigurationRequest{Request: req, Input: input, Copy: c.CreateVpcEndpointServiceConfigurationRequest}
}

// CreateVpcEndpointServiceConfigurationRequest is the request type for the
// CreateVpcEndpointServiceConfiguration API operation.
type CreateVpcEndpointServiceConfigurationRequest struct {
	*aws.Request
	Input *CreateVpcEndpointServiceConfigurationInput
	Copy  func(*CreateVpcEndpointServiceConfigurationInput) CreateVpcEndpointServiceConfigurationRequest
}

// Send marshals and sends the CreateVpcEndpointServiceConfiguration API request.
func (r CreateVpcEndpointServiceConfigurationRequest) Send(ctx context.Context) (*CreateVpcEndpointServiceConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateVpcEndpointServiceConfigurationResponse{
		CreateVpcEndpointServiceConfigurationOutput: r.Request.Data.(*CreateVpcEndpointServiceConfigurationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateVpcEndpointServiceConfigurationResponse is the response type for the
// CreateVpcEndpointServiceConfiguration API operation.
type CreateVpcEndpointServiceConfigurationResponse struct {
	*CreateVpcEndpointServiceConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateVpcEndpointServiceConfiguration request.
func (r *CreateVpcEndpointServiceConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
