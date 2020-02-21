// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/SetSecurityGroupsInput
type SetSecurityGroupsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the load balancer.
	//
	// LoadBalancerArn is a required field
	LoadBalancerArn *string `type:"string" required:"true"`

	// The IDs of the security groups.
	//
	// SecurityGroups is a required field
	SecurityGroups []string `type:"list" required:"true"`
}

// String returns the string representation
func (s SetSecurityGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetSecurityGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SetSecurityGroupsInput"}

	if s.LoadBalancerArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerArn"))
	}

	if s.SecurityGroups == nil {
		invalidParams.Add(aws.NewErrParamRequired("SecurityGroups"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/SetSecurityGroupsOutput
type SetSecurityGroupsOutput struct {
	_ struct{} `type:"structure"`

	// The IDs of the security groups associated with the load balancer.
	SecurityGroupIds []string `type:"list"`
}

// String returns the string representation
func (s SetSecurityGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

const opSetSecurityGroups = "SetSecurityGroups"

// SetSecurityGroupsRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Associates the specified security groups with the specified Application Load
// Balancer. The specified security groups override the previously associated
// security groups.
//
// You can't specify a security group for a Network Load Balancer.
//
//    // Example sending a request using SetSecurityGroupsRequest.
//    req := client.SetSecurityGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/SetSecurityGroups
func (c *Client) SetSecurityGroupsRequest(input *SetSecurityGroupsInput) SetSecurityGroupsRequest {
	op := &aws.Operation{
		Name:       opSetSecurityGroups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SetSecurityGroupsInput{}
	}

	req := c.newRequest(op, input, &SetSecurityGroupsOutput{})
	return SetSecurityGroupsRequest{Request: req, Input: input, Copy: c.SetSecurityGroupsRequest}
}

// SetSecurityGroupsRequest is the request type for the
// SetSecurityGroups API operation.
type SetSecurityGroupsRequest struct {
	*aws.Request
	Input *SetSecurityGroupsInput
	Copy  func(*SetSecurityGroupsInput) SetSecurityGroupsRequest
}

// Send marshals and sends the SetSecurityGroups API request.
func (r SetSecurityGroupsRequest) Send(ctx context.Context) (*SetSecurityGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetSecurityGroupsResponse{
		SetSecurityGroupsOutput: r.Request.Data.(*SetSecurityGroupsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetSecurityGroupsResponse is the response type for the
// SetSecurityGroups API operation.
type SetSecurityGroupsResponse struct {
	*SetSecurityGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetSecurityGroups request.
func (r *SetSecurityGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
