// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/GetPolicyVersionRequest
type GetPolicyVersionInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the managed policy that you want information
	// about.
	//
	// For more information about ARNs, see Amazon Resource Names (ARNs) and AWS
	// Service Namespaces (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the AWS General Reference.
	//
	// PolicyArn is a required field
	PolicyArn *string `min:"20" type:"string" required:"true"`

	// Identifies the policy version to retrieve.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters that consists of the lowercase letter 'v' followed
	// by one or two digits, and optionally followed by a period '.' and a string
	// of letters and digits.
	//
	// VersionId is a required field
	VersionId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetPolicyVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetPolicyVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetPolicyVersionInput"}

	if s.PolicyArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyArn"))
	}
	if s.PolicyArn != nil && len(*s.PolicyArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyArn", 20))
	}

	if s.VersionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VersionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a successful GetPolicyVersion request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/GetPolicyVersionResponse
type GetPolicyVersionOutput struct {
	_ struct{} `type:"structure"`

	// A structure containing details about the policy version.
	PolicyVersion *PolicyVersion `type:"structure"`
}

// String returns the string representation
func (s GetPolicyVersionOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetPolicyVersion = "GetPolicyVersion"

// GetPolicyVersionRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Retrieves information about the specified version of the specified managed
// policy, including the policy document.
//
// Policies returned by this API are URL-encoded compliant with RFC 3986 (https://tools.ietf.org/html/rfc3986).
// You can use a URL decoding method to convert the policy back to plain JSON
// text. For example, if you use Java, you can use the decode method of the
// java.net.URLDecoder utility class in the Java SDK. Other languages and SDKs
// provide similar functionality.
//
// To list the available versions for a policy, use ListPolicyVersions.
//
// This API retrieves information about managed policies. To retrieve information
// about an inline policy that is embedded in a user, group, or role, use the
// GetUserPolicy, GetGroupPolicy, or GetRolePolicy API.
//
// For more information about the types of policies, see Managed Policies and
// Inline Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide.
//
// For more information about managed policy versions, see Versioning for Managed
// Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-versions.html)
// in the IAM User Guide.
//
//    // Example sending a request using GetPolicyVersionRequest.
//    req := client.GetPolicyVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/GetPolicyVersion
func (c *Client) GetPolicyVersionRequest(input *GetPolicyVersionInput) GetPolicyVersionRequest {
	op := &aws.Operation{
		Name:       opGetPolicyVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetPolicyVersionInput{}
	}

	req := c.newRequest(op, input, &GetPolicyVersionOutput{})
	return GetPolicyVersionRequest{Request: req, Input: input, Copy: c.GetPolicyVersionRequest}
}

// GetPolicyVersionRequest is the request type for the
// GetPolicyVersion API operation.
type GetPolicyVersionRequest struct {
	*aws.Request
	Input *GetPolicyVersionInput
	Copy  func(*GetPolicyVersionInput) GetPolicyVersionRequest
}

// Send marshals and sends the GetPolicyVersion API request.
func (r GetPolicyVersionRequest) Send(ctx context.Context) (*GetPolicyVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetPolicyVersionResponse{
		GetPolicyVersionOutput: r.Request.Data.(*GetPolicyVersionOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetPolicyVersionResponse is the response type for the
// GetPolicyVersion API operation.
type GetPolicyVersionResponse struct {
	*GetPolicyVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetPolicyVersion request.
func (r *GetPolicyVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
