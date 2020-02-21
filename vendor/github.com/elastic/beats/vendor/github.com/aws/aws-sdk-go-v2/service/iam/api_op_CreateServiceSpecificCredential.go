// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/CreateServiceSpecificCredentialRequest
type CreateServiceSpecificCredentialInput struct {
	_ struct{} `type:"structure"`

	// The name of the AWS service that is to be associated with the credentials.
	// The service you specify here is the only service that can be accessed using
	// these credentials.
	//
	// ServiceName is a required field
	ServiceName *string `type:"string" required:"true"`

	// The name of the IAM user that is to be associated with the credentials. The
	// new service-specific credentials have the same permissions as the associated
	// user except that they can be used only to access the specified service.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// UserName is a required field
	UserName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateServiceSpecificCredentialInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateServiceSpecificCredentialInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateServiceSpecificCredentialInput"}

	if s.ServiceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServiceName"))
	}

	if s.UserName == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserName"))
	}
	if s.UserName != nil && len(*s.UserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/CreateServiceSpecificCredentialResponse
type CreateServiceSpecificCredentialOutput struct {
	_ struct{} `type:"structure"`

	// A structure that contains information about the newly created service-specific
	// credential.
	//
	// This is the only time that the password for this credential set is available.
	// It cannot be recovered later. Instead, you must reset the password with ResetServiceSpecificCredential.
	ServiceSpecificCredential *ServiceSpecificCredential `type:"structure"`
}

// String returns the string representation
func (s CreateServiceSpecificCredentialOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateServiceSpecificCredential = "CreateServiceSpecificCredential"

// CreateServiceSpecificCredentialRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Generates a set of credentials consisting of a user name and password that
// can be used to access the service specified in the request. These credentials
// are generated by IAM, and can be used only for the specified service.
//
// You can have a maximum of two sets of service-specific credentials for each
// supported service per user.
//
// The only supported service at this time is AWS CodeCommit.
//
// You can reset the password to a new service-generated value by calling ResetServiceSpecificCredential.
//
// For more information about service-specific credentials, see Using IAM with
// AWS CodeCommit: Git Credentials, SSH Keys, and AWS Access Keys (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_ssh-keys.html)
// in the IAM User Guide.
//
//    // Example sending a request using CreateServiceSpecificCredentialRequest.
//    req := client.CreateServiceSpecificCredentialRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/CreateServiceSpecificCredential
func (c *Client) CreateServiceSpecificCredentialRequest(input *CreateServiceSpecificCredentialInput) CreateServiceSpecificCredentialRequest {
	op := &aws.Operation{
		Name:       opCreateServiceSpecificCredential,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateServiceSpecificCredentialInput{}
	}

	req := c.newRequest(op, input, &CreateServiceSpecificCredentialOutput{})
	return CreateServiceSpecificCredentialRequest{Request: req, Input: input, Copy: c.CreateServiceSpecificCredentialRequest}
}

// CreateServiceSpecificCredentialRequest is the request type for the
// CreateServiceSpecificCredential API operation.
type CreateServiceSpecificCredentialRequest struct {
	*aws.Request
	Input *CreateServiceSpecificCredentialInput
	Copy  func(*CreateServiceSpecificCredentialInput) CreateServiceSpecificCredentialRequest
}

// Send marshals and sends the CreateServiceSpecificCredential API request.
func (r CreateServiceSpecificCredentialRequest) Send(ctx context.Context) (*CreateServiceSpecificCredentialResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateServiceSpecificCredentialResponse{
		CreateServiceSpecificCredentialOutput: r.Request.Data.(*CreateServiceSpecificCredentialOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateServiceSpecificCredentialResponse is the response type for the
// CreateServiceSpecificCredential API operation.
type CreateServiceSpecificCredentialResponse struct {
	*CreateServiceSpecificCredentialOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateServiceSpecificCredential request.
func (r *CreateServiceSpecificCredentialResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
