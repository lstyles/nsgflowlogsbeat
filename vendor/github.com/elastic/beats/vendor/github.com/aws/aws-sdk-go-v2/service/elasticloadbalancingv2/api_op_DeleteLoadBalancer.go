// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/DeleteLoadBalancerInput
type DeleteLoadBalancerInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the load balancer.
	//
	// LoadBalancerArn is a required field
	LoadBalancerArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteLoadBalancerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteLoadBalancerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteLoadBalancerInput"}

	if s.LoadBalancerArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/DeleteLoadBalancerOutput
type DeleteLoadBalancerOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteLoadBalancerOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteLoadBalancer = "DeleteLoadBalancer"

// DeleteLoadBalancerRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Deletes the specified Application Load Balancer or Network Load Balancer
// and its attached listeners.
//
// You can't delete a load balancer if deletion protection is enabled. If the
// load balancer does not exist or has already been deleted, the call succeeds.
//
// Deleting a load balancer does not affect its registered targets. For example,
// your EC2 instances continue to run and are still registered to their target
// groups. If you no longer need these EC2 instances, you can stop or terminate
// them.
//
//    // Example sending a request using DeleteLoadBalancerRequest.
//    req := client.DeleteLoadBalancerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/DeleteLoadBalancer
func (c *Client) DeleteLoadBalancerRequest(input *DeleteLoadBalancerInput) DeleteLoadBalancerRequest {
	op := &aws.Operation{
		Name:       opDeleteLoadBalancer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteLoadBalancerInput{}
	}

	req := c.newRequest(op, input, &DeleteLoadBalancerOutput{})
	return DeleteLoadBalancerRequest{Request: req, Input: input, Copy: c.DeleteLoadBalancerRequest}
}

// DeleteLoadBalancerRequest is the request type for the
// DeleteLoadBalancer API operation.
type DeleteLoadBalancerRequest struct {
	*aws.Request
	Input *DeleteLoadBalancerInput
	Copy  func(*DeleteLoadBalancerInput) DeleteLoadBalancerRequest
}

// Send marshals and sends the DeleteLoadBalancer API request.
func (r DeleteLoadBalancerRequest) Send(ctx context.Context) (*DeleteLoadBalancerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteLoadBalancerResponse{
		DeleteLoadBalancerOutput: r.Request.Data.(*DeleteLoadBalancerOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteLoadBalancerResponse is the response type for the
// DeleteLoadBalancer API operation.
type DeleteLoadBalancerResponse struct {
	*DeleteLoadBalancerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteLoadBalancer request.
func (r *DeleteLoadBalancerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
