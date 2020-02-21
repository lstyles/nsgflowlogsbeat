// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/UpdateAccessKeyRequest
type UpdateAccessKeyInput struct {
	_ struct{} `type:"structure"`

	// The access key ID of the secret access key you want to update.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters that can consist of any upper or lowercased letter
	// or digit.
	//
	// AccessKeyId is a required field
	AccessKeyId *string `min:"16" type:"string" required:"true"`

	// The status you want to assign to the secret access key. Active means that
	// the key can be used for API calls to AWS, while Inactive means that the key
	// cannot be used.
	//
	// Status is a required field
	Status StatusType `type:"string" required:"true" enum:"true"`

	// The name of the user whose key you want to update.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	UserName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateAccessKeyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateAccessKeyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateAccessKeyInput"}

	if s.AccessKeyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccessKeyId"))
	}
	if s.AccessKeyId != nil && len(*s.AccessKeyId) < 16 {
		invalidParams.Add(aws.NewErrParamMinLen("AccessKeyId", 16))
	}
	if len(s.Status) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Status"))
	}
	if s.UserName != nil && len(*s.UserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/UpdateAccessKeyOutput
type UpdateAccessKeyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateAccessKeyOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateAccessKey = "UpdateAccessKey"

// UpdateAccessKeyRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Changes the status of the specified access key from Active to Inactive, or
// vice versa. This operation can be used to disable a user's key as part of
// a key rotation workflow.
//
// If the UserName is not specified, the user name is determined implicitly
// based on the AWS access key ID used to sign the request. This operation works
// for access keys under the AWS account. Consequently, you can use this operation
// to manage AWS account root user credentials even if the AWS account has no
// associated users.
//
// For information about rotating keys, see Managing Keys and Certificates (https://docs.aws.amazon.com/IAM/latest/UserGuide/ManagingCredentials.html)
// in the IAM User Guide.
//
//    // Example sending a request using UpdateAccessKeyRequest.
//    req := client.UpdateAccessKeyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/UpdateAccessKey
func (c *Client) UpdateAccessKeyRequest(input *UpdateAccessKeyInput) UpdateAccessKeyRequest {
	op := &aws.Operation{
		Name:       opUpdateAccessKey,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateAccessKeyInput{}
	}

	req := c.newRequest(op, input, &UpdateAccessKeyOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UpdateAccessKeyRequest{Request: req, Input: input, Copy: c.UpdateAccessKeyRequest}
}

// UpdateAccessKeyRequest is the request type for the
// UpdateAccessKey API operation.
type UpdateAccessKeyRequest struct {
	*aws.Request
	Input *UpdateAccessKeyInput
	Copy  func(*UpdateAccessKeyInput) UpdateAccessKeyRequest
}

// Send marshals and sends the UpdateAccessKey API request.
func (r UpdateAccessKeyRequest) Send(ctx context.Context) (*UpdateAccessKeyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateAccessKeyResponse{
		UpdateAccessKeyOutput: r.Request.Data.(*UpdateAccessKeyOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateAccessKeyResponse is the response type for the
// UpdateAccessKey API operation.
type UpdateAccessKeyResponse struct {
	*UpdateAccessKeyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateAccessKey request.
func (r *UpdateAccessKeyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
