// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AssociateDhcpOptionsRequest
type AssociateDhcpOptionsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the DHCP options set, or default to associate no DHCP options with
	// the VPC.
	//
	// DhcpOptionsId is a required field
	DhcpOptionsId *string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the VPC.
	//
	// VpcId is a required field
	VpcId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateDhcpOptionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateDhcpOptionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateDhcpOptionsInput"}

	if s.DhcpOptionsId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DhcpOptionsId"))
	}

	if s.VpcId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpcId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AssociateDhcpOptionsOutput
type AssociateDhcpOptionsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AssociateDhcpOptionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opAssociateDhcpOptions = "AssociateDhcpOptions"

// AssociateDhcpOptionsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Associates a set of DHCP options (that you've previously created) with the
// specified VPC, or associates no DHCP options with the VPC.
//
// After you associate the options with the VPC, any existing instances and
// all new instances that you launch in that VPC use the options. You don't
// need to restart or relaunch the instances. They automatically pick up the
// changes within a few hours, depending on how frequently the instance renews
// its DHCP lease. You can explicitly renew the lease using the operating system
// on the instance.
//
// For more information, see DHCP Options Sets (https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_DHCP_Options.html)
// in the Amazon Virtual Private Cloud User Guide.
//
//    // Example sending a request using AssociateDhcpOptionsRequest.
//    req := client.AssociateDhcpOptionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AssociateDhcpOptions
func (c *Client) AssociateDhcpOptionsRequest(input *AssociateDhcpOptionsInput) AssociateDhcpOptionsRequest {
	op := &aws.Operation{
		Name:       opAssociateDhcpOptions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociateDhcpOptionsInput{}
	}

	req := c.newRequest(op, input, &AssociateDhcpOptionsOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return AssociateDhcpOptionsRequest{Request: req, Input: input, Copy: c.AssociateDhcpOptionsRequest}
}

// AssociateDhcpOptionsRequest is the request type for the
// AssociateDhcpOptions API operation.
type AssociateDhcpOptionsRequest struct {
	*aws.Request
	Input *AssociateDhcpOptionsInput
	Copy  func(*AssociateDhcpOptionsInput) AssociateDhcpOptionsRequest
}

// Send marshals and sends the AssociateDhcpOptions API request.
func (r AssociateDhcpOptionsRequest) Send(ctx context.Context) (*AssociateDhcpOptionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateDhcpOptionsResponse{
		AssociateDhcpOptionsOutput: r.Request.Data.(*AssociateDhcpOptionsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateDhcpOptionsResponse is the response type for the
// AssociateDhcpOptions API operation.
type AssociateDhcpOptionsResponse struct {
	*AssociateDhcpOptionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateDhcpOptions request.
func (r *AssociateDhcpOptionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
