// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/EnableEbsEncryptionByDefaultRequest
type EnableEbsEncryptionByDefaultInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`
}

// String returns the string representation
func (s EnableEbsEncryptionByDefaultInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/EnableEbsEncryptionByDefaultResult
type EnableEbsEncryptionByDefaultOutput struct {
	_ struct{} `type:"structure"`

	// Account-level encryption status after performing the action.
	EbsEncryptionByDefault *bool `locationName:"ebsEncryptionByDefault" type:"boolean"`
}

// String returns the string representation
func (s EnableEbsEncryptionByDefaultOutput) String() string {
	return awsutil.Prettify(s)
}

const opEnableEbsEncryptionByDefault = "EnableEbsEncryptionByDefault"

// EnableEbsEncryptionByDefaultRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Enables default encryption for EBS volumes that are created in your account
// in the current region.
//
// Once encryption is enabled with this action, EBS volumes that are created
// in your account will always be encrypted even if encryption is not specified
// at launch. This setting overrides the encrypted setting to true in all API
// calls that create EBS volumes in your account. A volume will be encrypted
// even if you specify encryption to be false in the API call that creates the
// volume.
//
// If you do not specify a customer master key (CMK) in the API call that creates
// the EBS volume, then the volume is encrypted to your AWS account's default
// CMK.
//
// You can specify a default CMK of your choice using ModifyEbsDefaultKmsKeyId.
//
// Enabling default encryption for EBS volumes has no effect on existing unencrypted
// volumes in your account. Encrypting the data in these requires manual action.
// You can either create an encrypted snapshot of an unencrypted volume, or
// encrypt a copy of an unencrypted snapshot. Any volume restored from an encrypted
// snapshot is also encrypted. For more information, see Amazon EBS Snapshots
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSSnapshots.html).
//
// Once EBS encryption by default is enabled, you can no longer launch older-generation
// instance types that do not support encryption. For more information, see
// Supported Instance Types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#EBSEncryption_supported_instances).
//
//    // Example sending a request using EnableEbsEncryptionByDefaultRequest.
//    req := client.EnableEbsEncryptionByDefaultRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/EnableEbsEncryptionByDefault
func (c *Client) EnableEbsEncryptionByDefaultRequest(input *EnableEbsEncryptionByDefaultInput) EnableEbsEncryptionByDefaultRequest {
	op := &aws.Operation{
		Name:       opEnableEbsEncryptionByDefault,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &EnableEbsEncryptionByDefaultInput{}
	}

	req := c.newRequest(op, input, &EnableEbsEncryptionByDefaultOutput{})
	return EnableEbsEncryptionByDefaultRequest{Request: req, Input: input, Copy: c.EnableEbsEncryptionByDefaultRequest}
}

// EnableEbsEncryptionByDefaultRequest is the request type for the
// EnableEbsEncryptionByDefault API operation.
type EnableEbsEncryptionByDefaultRequest struct {
	*aws.Request
	Input *EnableEbsEncryptionByDefaultInput
	Copy  func(*EnableEbsEncryptionByDefaultInput) EnableEbsEncryptionByDefaultRequest
}

// Send marshals and sends the EnableEbsEncryptionByDefault API request.
func (r EnableEbsEncryptionByDefaultRequest) Send(ctx context.Context) (*EnableEbsEncryptionByDefaultResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &EnableEbsEncryptionByDefaultResponse{
		EnableEbsEncryptionByDefaultOutput: r.Request.Data.(*EnableEbsEncryptionByDefaultOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// EnableEbsEncryptionByDefaultResponse is the response type for the
// EnableEbsEncryptionByDefault API operation.
type EnableEbsEncryptionByDefaultResponse struct {
	*EnableEbsEncryptionByDefaultOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// EnableEbsEncryptionByDefault request.
func (r *EnableEbsEncryptionByDefaultResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
