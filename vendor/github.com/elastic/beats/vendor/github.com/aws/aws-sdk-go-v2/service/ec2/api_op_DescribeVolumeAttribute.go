// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for DescribeVolumeAttribute.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeVolumeAttributeRequest
type DescribeVolumeAttributeInput struct {
	_ struct{} `type:"structure"`

	// The attribute of the volume. This parameter is required.
	//
	// Attribute is a required field
	Attribute VolumeAttributeName `type:"string" required:"true" enum:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the volume.
	//
	// VolumeId is a required field
	VolumeId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeVolumeAttributeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeVolumeAttributeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeVolumeAttributeInput"}
	if len(s.Attribute) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Attribute"))
	}

	if s.VolumeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VolumeId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of DescribeVolumeAttribute.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeVolumeAttributeResult
type DescribeVolumeAttributeOutput struct {
	_ struct{} `type:"structure"`

	// The state of autoEnableIO attribute.
	AutoEnableIO *AttributeBooleanValue `locationName:"autoEnableIO" type:"structure"`

	// A list of product codes.
	ProductCodes []ProductCode `locationName:"productCodes" locationNameList:"item" type:"list"`

	// The ID of the volume.
	VolumeId *string `locationName:"volumeId" type:"string"`
}

// String returns the string representation
func (s DescribeVolumeAttributeOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeVolumeAttribute = "DescribeVolumeAttribute"

// DescribeVolumeAttributeRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the specified attribute of the specified volume. You can specify
// only one attribute at a time.
//
// For more information about EBS volumes, see Amazon EBS Volumes (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumes.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using DescribeVolumeAttributeRequest.
//    req := client.DescribeVolumeAttributeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeVolumeAttribute
func (c *Client) DescribeVolumeAttributeRequest(input *DescribeVolumeAttributeInput) DescribeVolumeAttributeRequest {
	op := &aws.Operation{
		Name:       opDescribeVolumeAttribute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeVolumeAttributeInput{}
	}

	req := c.newRequest(op, input, &DescribeVolumeAttributeOutput{})
	return DescribeVolumeAttributeRequest{Request: req, Input: input, Copy: c.DescribeVolumeAttributeRequest}
}

// DescribeVolumeAttributeRequest is the request type for the
// DescribeVolumeAttribute API operation.
type DescribeVolumeAttributeRequest struct {
	*aws.Request
	Input *DescribeVolumeAttributeInput
	Copy  func(*DescribeVolumeAttributeInput) DescribeVolumeAttributeRequest
}

// Send marshals and sends the DescribeVolumeAttribute API request.
func (r DescribeVolumeAttributeRequest) Send(ctx context.Context) (*DescribeVolumeAttributeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeVolumeAttributeResponse{
		DescribeVolumeAttributeOutput: r.Request.Data.(*DescribeVolumeAttributeOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeVolumeAttributeResponse is the response type for the
// DescribeVolumeAttribute API operation.
type DescribeVolumeAttributeResponse struct {
	*DescribeVolumeAttributeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeVolumeAttribute request.
func (r *DescribeVolumeAttributeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
