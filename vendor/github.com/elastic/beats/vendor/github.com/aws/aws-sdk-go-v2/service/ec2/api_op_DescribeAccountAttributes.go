// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeAccountAttributesRequest
type DescribeAccountAttributesInput struct {
	_ struct{} `type:"structure"`

	// The account attribute names.
	AttributeNames []AccountAttributeName `locationName:"attributeName" locationNameList:"attributeName" type:"list"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`
}

// String returns the string representation
func (s DescribeAccountAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeAccountAttributesResult
type DescribeAccountAttributesOutput struct {
	_ struct{} `type:"structure"`

	// Information about the account attributes.
	AccountAttributes []AccountAttribute `locationName:"accountAttributeSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeAccountAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeAccountAttributes = "DescribeAccountAttributes"

// DescribeAccountAttributesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes attributes of your AWS account. The following are the supported
// account attributes:
//
//    * supported-platforms: Indicates whether your account can launch instances
//    into EC2-Classic and EC2-VPC, or only into EC2-VPC.
//
//    * default-vpc: The ID of the default VPC for your account, or none.
//
//    * max-instances: The maximum number of On-Demand Instances that you can
//    run.
//
//    * vpc-max-security-groups-per-interface: The maximum number of security
//    groups that you can assign to a network interface.
//
//    * max-elastic-ips: The maximum number of Elastic IP addresses that you
//    can allocate for use with EC2-Classic.
//
//    * vpc-max-elastic-ips: The maximum number of Elastic IP addresses that
//    you can allocate for use with EC2-VPC.
//
//    // Example sending a request using DescribeAccountAttributesRequest.
//    req := client.DescribeAccountAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeAccountAttributes
func (c *Client) DescribeAccountAttributesRequest(input *DescribeAccountAttributesInput) DescribeAccountAttributesRequest {
	op := &aws.Operation{
		Name:       opDescribeAccountAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeAccountAttributesInput{}
	}

	req := c.newRequest(op, input, &DescribeAccountAttributesOutput{})
	return DescribeAccountAttributesRequest{Request: req, Input: input, Copy: c.DescribeAccountAttributesRequest}
}

// DescribeAccountAttributesRequest is the request type for the
// DescribeAccountAttributes API operation.
type DescribeAccountAttributesRequest struct {
	*aws.Request
	Input *DescribeAccountAttributesInput
	Copy  func(*DescribeAccountAttributesInput) DescribeAccountAttributesRequest
}

// Send marshals and sends the DescribeAccountAttributes API request.
func (r DescribeAccountAttributesRequest) Send(ctx context.Context) (*DescribeAccountAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAccountAttributesResponse{
		DescribeAccountAttributesOutput: r.Request.Data.(*DescribeAccountAttributesOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAccountAttributesResponse is the response type for the
// DescribeAccountAttributes API operation.
type DescribeAccountAttributesResponse struct {
	*DescribeAccountAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAccountAttributes request.
func (r *DescribeAccountAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
