// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sts

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sts-2011-06-15/GetCallerIdentityRequest
type GetCallerIdentityInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetCallerIdentityInput) String() string {
	return awsutil.Prettify(s)
}

// Contains the response to a successful GetCallerIdentity request, including
// information about the entity making the request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sts-2011-06-15/GetCallerIdentityResponse
type GetCallerIdentityOutput struct {
	_ struct{} `type:"structure"`

	// The AWS account ID number of the account that owns or contains the calling
	// entity.
	Account *string `type:"string"`

	// The AWS ARN associated with the calling entity.
	Arn *string `min:"20" type:"string"`

	// The unique identifier of the calling entity. The exact value depends on the
	// type of entity that is making the call. The values returned are those listed
	// in the aws:userid column in the Principal table (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_variables.html#principaltable)
	// found on the Policy Variables reference page in the IAM User Guide.
	UserId *string `type:"string"`
}

// String returns the string representation
func (s GetCallerIdentityOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetCallerIdentity = "GetCallerIdentity"

// GetCallerIdentityRequest returns a request value for making API operation for
// AWS Security Token Service.
//
// Returns details about the IAM identity whose credentials are used to call
// the API.
//
//    // Example sending a request using GetCallerIdentityRequest.
//    req := client.GetCallerIdentityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sts-2011-06-15/GetCallerIdentity
func (c *Client) GetCallerIdentityRequest(input *GetCallerIdentityInput) GetCallerIdentityRequest {
	op := &aws.Operation{
		Name:       opGetCallerIdentity,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetCallerIdentityInput{}
	}

	req := c.newRequest(op, input, &GetCallerIdentityOutput{})
	return GetCallerIdentityRequest{Request: req, Input: input, Copy: c.GetCallerIdentityRequest}
}

// GetCallerIdentityRequest is the request type for the
// GetCallerIdentity API operation.
type GetCallerIdentityRequest struct {
	*aws.Request
	Input *GetCallerIdentityInput
	Copy  func(*GetCallerIdentityInput) GetCallerIdentityRequest
}

// Send marshals and sends the GetCallerIdentity API request.
func (r GetCallerIdentityRequest) Send(ctx context.Context) (*GetCallerIdentityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetCallerIdentityResponse{
		GetCallerIdentityOutput: r.Request.Data.(*GetCallerIdentityOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetCallerIdentityResponse is the response type for the
// GetCallerIdentity API operation.
type GetCallerIdentityResponse struct {
	*GetCallerIdentityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetCallerIdentity request.
func (r *GetCallerIdentityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
