// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyInstanceAttributeRequest
type ModifyInstanceAttributeInput struct {
	_ struct{} `type:"structure"`

	// The name of the attribute.
	Attribute InstanceAttributeName `locationName:"attribute" type:"string" enum:"true"`

	// Modifies the DeleteOnTermination attribute for volumes that are currently
	// attached. The volume must be owned by the caller. If no value is specified
	// for DeleteOnTermination, the default is true and the volume is deleted when
	// the instance is terminated.
	//
	// To add instance store volumes to an Amazon EBS-backed instance, you must
	// add them when you launch the instance. For more information, see Updating
	// the Block Device Mapping when Launching an Instance (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html#Using_OverridingAMIBDM)
	// in the Amazon Elastic Compute Cloud User Guide.
	BlockDeviceMappings []InstanceBlockDeviceMappingSpecification `locationName:"blockDeviceMapping" locationNameList:"item" type:"list"`

	// If the value is true, you can't terminate the instance using the Amazon EC2
	// console, CLI, or API; otherwise, you can. You cannot use this parameter for
	// Spot Instances.
	DisableApiTermination *AttributeBooleanValue `locationName:"disableApiTermination" type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// Specifies whether the instance is optimized for Amazon EBS I/O. This optimization
	// provides dedicated throughput to Amazon EBS and an optimized configuration
	// stack to provide optimal EBS I/O performance. This optimization isn't available
	// with all instance types. Additional usage charges apply when using an EBS
	// Optimized instance.
	EbsOptimized *AttributeBooleanValue `locationName:"ebsOptimized" type:"structure"`

	// Set to true to enable enhanced networking with ENA for the instance.
	//
	// This option is supported only for HVM instances. Specifying this option with
	// a PV instance can make it unreachable.
	EnaSupport *AttributeBooleanValue `locationName:"enaSupport" type:"structure"`

	// [EC2-VPC] Changes the security groups of the instance. You must specify at
	// least one security group, even if it's just the default security group for
	// the VPC. You must specify the security group ID, not the security group name.
	Groups []string `locationName:"GroupId" locationNameList:"groupId" type:"list"`

	// The ID of the instance.
	//
	// InstanceId is a required field
	InstanceId *string `locationName:"instanceId" type:"string" required:"true"`

	// Specifies whether an instance stops or terminates when you initiate shutdown
	// from the instance (using the operating system command for system shutdown).
	InstanceInitiatedShutdownBehavior *AttributeValue `locationName:"instanceInitiatedShutdownBehavior" type:"structure"`

	// Changes the instance type to the specified value. For more information, see
	// Instance Types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html).
	// If the instance type is not valid, the error returned is InvalidInstanceAttributeValue.
	InstanceType *AttributeValue `locationName:"instanceType" type:"structure"`

	// Changes the instance's kernel to the specified value. We recommend that you
	// use PV-GRUB instead of kernels and RAM disks. For more information, see PV-GRUB
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedKernels.html).
	Kernel *AttributeValue `locationName:"kernel" type:"structure"`

	// Changes the instance's RAM disk to the specified value. We recommend that
	// you use PV-GRUB instead of kernels and RAM disks. For more information, see
	// PV-GRUB (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedKernels.html).
	Ramdisk *AttributeValue `locationName:"ramdisk" type:"structure"`

	// Specifies whether source/destination checking is enabled. A value of true
	// means that checking is enabled, and false means that checking is disabled.
	// This value must be false for a NAT instance to perform NAT.
	SourceDestCheck *AttributeBooleanValue `type:"structure"`

	// Set to simple to enable enhanced networking with the Intel 82599 Virtual
	// Function interface for the instance.
	//
	// There is no way to disable enhanced networking with the Intel 82599 Virtual
	// Function interface at this time.
	//
	// This option is supported only for HVM instances. Specifying this option with
	// a PV instance can make it unreachable.
	SriovNetSupport *AttributeValue `locationName:"sriovNetSupport" type:"structure"`

	// Changes the instance's user data to the specified value. If you are using
	// an AWS SDK or command line tool, base64-encoding is performed for you, and
	// you can load the text from a file. Otherwise, you must provide base64-encoded
	// text.
	UserData *BlobAttributeValue `locationName:"userData" type:"structure"`

	// A new value for the attribute. Use only with the kernel, ramdisk, userData,
	// disableApiTermination, or instanceInitiatedShutdownBehavior attribute.
	Value *string `locationName:"value" type:"string"`
}

// String returns the string representation
func (s ModifyInstanceAttributeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyInstanceAttributeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyInstanceAttributeInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyInstanceAttributeOutput
type ModifyInstanceAttributeOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ModifyInstanceAttributeOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyInstanceAttribute = "ModifyInstanceAttribute"

// ModifyInstanceAttributeRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Modifies the specified attribute of the specified instance. You can specify
// only one attribute at a time.
//
// Note: Using this action to change the security groups associated with an
// elastic network interface (ENI) attached to an instance in a VPC can result
// in an error if the instance has more than one ENI. To change the security
// groups associated with an ENI attached to an instance that has multiple ENIs,
// we recommend that you use the ModifyNetworkInterfaceAttribute action.
//
// To modify some attributes, the instance must be stopped. For more information,
// see Modifying Attributes of a Stopped Instance (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_ChangingAttributesWhileInstanceStopped.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using ModifyInstanceAttributeRequest.
//    req := client.ModifyInstanceAttributeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyInstanceAttribute
func (c *Client) ModifyInstanceAttributeRequest(input *ModifyInstanceAttributeInput) ModifyInstanceAttributeRequest {
	op := &aws.Operation{
		Name:       opModifyInstanceAttribute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyInstanceAttributeInput{}
	}

	req := c.newRequest(op, input, &ModifyInstanceAttributeOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return ModifyInstanceAttributeRequest{Request: req, Input: input, Copy: c.ModifyInstanceAttributeRequest}
}

// ModifyInstanceAttributeRequest is the request type for the
// ModifyInstanceAttribute API operation.
type ModifyInstanceAttributeRequest struct {
	*aws.Request
	Input *ModifyInstanceAttributeInput
	Copy  func(*ModifyInstanceAttributeInput) ModifyInstanceAttributeRequest
}

// Send marshals and sends the ModifyInstanceAttribute API request.
func (r ModifyInstanceAttributeRequest) Send(ctx context.Context) (*ModifyInstanceAttributeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyInstanceAttributeResponse{
		ModifyInstanceAttributeOutput: r.Request.Data.(*ModifyInstanceAttributeOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyInstanceAttributeResponse is the response type for the
// ModifyInstanceAttribute API operation.
type ModifyInstanceAttributeResponse struct {
	*ModifyInstanceAttributeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyInstanceAttribute request.
func (r *ModifyInstanceAttributeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
