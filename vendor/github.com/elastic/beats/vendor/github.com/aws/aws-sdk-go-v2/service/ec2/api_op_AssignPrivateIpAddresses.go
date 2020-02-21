// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
)

// Contains the parameters for AssignPrivateIpAddresses.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AssignPrivateIpAddressesRequest
type AssignPrivateIpAddressesInput struct {
	_ struct{} `type:"structure"`

	// Indicates whether to allow an IP address that is already assigned to another
	// network interface or instance to be reassigned to the specified network interface.
	AllowReassignment *bool `locationName:"allowReassignment" type:"boolean"`

	// The ID of the network interface.
	//
	// NetworkInterfaceId is a required field
	NetworkInterfaceId *string `locationName:"networkInterfaceId" type:"string" required:"true"`

	// One or more IP addresses to be assigned as a secondary private IP address
	// to the network interface. You can't specify this parameter when also specifying
	// a number of secondary IP addresses.
	//
	// If you don't specify an IP address, Amazon EC2 automatically selects an IP
	// address within the subnet range.
	PrivateIpAddresses []string `locationName:"privateIpAddress" locationNameList:"PrivateIpAddress" type:"list"`

	// The number of secondary IP addresses to assign to the network interface.
	// You can't specify this parameter when also specifying private IP addresses.
	SecondaryPrivateIpAddressCount *int64 `locationName:"secondaryPrivateIpAddressCount" type:"integer"`
}

// String returns the string representation
func (s AssignPrivateIpAddressesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssignPrivateIpAddressesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssignPrivateIpAddressesInput"}

	if s.NetworkInterfaceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("NetworkInterfaceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AssignPrivateIpAddressesOutput
type AssignPrivateIpAddressesOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AssignPrivateIpAddressesOutput) String() string {
	return awsutil.Prettify(s)
}

const opAssignPrivateIpAddresses = "AssignPrivateIpAddresses"

// AssignPrivateIpAddressesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Assigns one or more secondary private IP addresses to the specified network
// interface.
//
// You can specify one or more specific secondary IP addresses, or you can specify
// the number of secondary IP addresses to be automatically assigned within
// the subnet's CIDR block range. The number of secondary IP addresses that
// you can assign to an instance varies by instance type. For information about
// instance types, see Instance Types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html)
// in the Amazon Elastic Compute Cloud User Guide. For more information about
// Elastic IP addresses, see Elastic IP Addresses (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
// When you move a secondary private IP address to another network interface,
// any Elastic IP address that is associated with the IP address is also moved.
//
// Remapping an IP address is an asynchronous operation. When you move an IP
// address from one network interface to another, check network/interfaces/macs/mac/local-ipv4s
// in the instance metadata to confirm that the remapping is complete.
//
//    // Example sending a request using AssignPrivateIpAddressesRequest.
//    req := client.AssignPrivateIpAddressesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AssignPrivateIpAddresses
func (c *Client) AssignPrivateIpAddressesRequest(input *AssignPrivateIpAddressesInput) AssignPrivateIpAddressesRequest {
	op := &aws.Operation{
		Name:       opAssignPrivateIpAddresses,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssignPrivateIpAddressesInput{}
	}

	req := c.newRequest(op, input, &AssignPrivateIpAddressesOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return AssignPrivateIpAddressesRequest{Request: req, Input: input, Copy: c.AssignPrivateIpAddressesRequest}
}

// AssignPrivateIpAddressesRequest is the request type for the
// AssignPrivateIpAddresses API operation.
type AssignPrivateIpAddressesRequest struct {
	*aws.Request
	Input *AssignPrivateIpAddressesInput
	Copy  func(*AssignPrivateIpAddressesInput) AssignPrivateIpAddressesRequest
}

// Send marshals and sends the AssignPrivateIpAddresses API request.
func (r AssignPrivateIpAddressesRequest) Send(ctx context.Context) (*AssignPrivateIpAddressesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssignPrivateIpAddressesResponse{
		AssignPrivateIpAddressesOutput: r.Request.Data.(*AssignPrivateIpAddressesOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssignPrivateIpAddressesResponse is the response type for the
// AssignPrivateIpAddresses API operation.
type AssignPrivateIpAddressesResponse struct {
	*AssignPrivateIpAddressesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssignPrivateIpAddresses request.
func (r *AssignPrivateIpAddressesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
