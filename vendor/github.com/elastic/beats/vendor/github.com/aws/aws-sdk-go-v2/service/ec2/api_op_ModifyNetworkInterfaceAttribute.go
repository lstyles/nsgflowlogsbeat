// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
)

// Contains the parameters for ModifyNetworkInterfaceAttribute.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyNetworkInterfaceAttributeRequest
type ModifyNetworkInterfaceAttributeInput struct {
	_ struct{} `type:"structure"`

	// Information about the interface attachment. If modifying the 'delete on termination'
	// attribute, you must specify the ID of the interface attachment.
	Attachment *NetworkInterfaceAttachmentChanges `locationName:"attachment" type:"structure"`

	// A description for the network interface.
	Description *AttributeValue `locationName:"description" type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// Changes the security groups for the network interface. The new set of groups
	// you specify replaces the current set. You must specify at least one group,
	// even if it's just the default security group in the VPC. You must specify
	// the ID of the security group, not the name.
	Groups []string `locationName:"SecurityGroupId" locationNameList:"SecurityGroupId" type:"list"`

	// The ID of the network interface.
	//
	// NetworkInterfaceId is a required field
	NetworkInterfaceId *string `locationName:"networkInterfaceId" type:"string" required:"true"`

	// Indicates whether source/destination checking is enabled. A value of true
	// means checking is enabled, and false means checking is disabled. This value
	// must be false for a NAT instance to perform NAT. For more information, see
	// NAT Instances (https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_NAT_Instance.html)
	// in the Amazon Virtual Private Cloud User Guide.
	SourceDestCheck *AttributeBooleanValue `locationName:"sourceDestCheck" type:"structure"`
}

// String returns the string representation
func (s ModifyNetworkInterfaceAttributeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyNetworkInterfaceAttributeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyNetworkInterfaceAttributeInput"}

	if s.NetworkInterfaceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("NetworkInterfaceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyNetworkInterfaceAttributeOutput
type ModifyNetworkInterfaceAttributeOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ModifyNetworkInterfaceAttributeOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyNetworkInterfaceAttribute = "ModifyNetworkInterfaceAttribute"

// ModifyNetworkInterfaceAttributeRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Modifies the specified network interface attribute. You can specify only
// one attribute at a time. You can use this action to attach and detach security
// groups from an existing EC2 instance.
//
//    // Example sending a request using ModifyNetworkInterfaceAttributeRequest.
//    req := client.ModifyNetworkInterfaceAttributeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyNetworkInterfaceAttribute
func (c *Client) ModifyNetworkInterfaceAttributeRequest(input *ModifyNetworkInterfaceAttributeInput) ModifyNetworkInterfaceAttributeRequest {
	op := &aws.Operation{
		Name:       opModifyNetworkInterfaceAttribute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyNetworkInterfaceAttributeInput{}
	}

	req := c.newRequest(op, input, &ModifyNetworkInterfaceAttributeOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return ModifyNetworkInterfaceAttributeRequest{Request: req, Input: input, Copy: c.ModifyNetworkInterfaceAttributeRequest}
}

// ModifyNetworkInterfaceAttributeRequest is the request type for the
// ModifyNetworkInterfaceAttribute API operation.
type ModifyNetworkInterfaceAttributeRequest struct {
	*aws.Request
	Input *ModifyNetworkInterfaceAttributeInput
	Copy  func(*ModifyNetworkInterfaceAttributeInput) ModifyNetworkInterfaceAttributeRequest
}

// Send marshals and sends the ModifyNetworkInterfaceAttribute API request.
func (r ModifyNetworkInterfaceAttributeRequest) Send(ctx context.Context) (*ModifyNetworkInterfaceAttributeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyNetworkInterfaceAttributeResponse{
		ModifyNetworkInterfaceAttributeOutput: r.Request.Data.(*ModifyNetworkInterfaceAttributeOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyNetworkInterfaceAttributeResponse is the response type for the
// ModifyNetworkInterfaceAttribute API operation.
type ModifyNetworkInterfaceAttributeResponse struct {
	*ModifyNetworkInterfaceAttributeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyNetworkInterfaceAttribute request.
func (r *ModifyNetworkInterfaceAttributeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
