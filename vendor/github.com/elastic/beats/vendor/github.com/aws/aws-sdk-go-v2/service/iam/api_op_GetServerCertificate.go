// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/GetServerCertificateRequest
type GetServerCertificateInput struct {
	_ struct{} `type:"structure"`

	// The name of the server certificate you want to retrieve information about.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// ServerCertificateName is a required field
	ServerCertificateName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetServerCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetServerCertificateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetServerCertificateInput"}

	if s.ServerCertificateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServerCertificateName"))
	}
	if s.ServerCertificateName != nil && len(*s.ServerCertificateName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ServerCertificateName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a successful GetServerCertificate request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/GetServerCertificateResponse
type GetServerCertificateOutput struct {
	_ struct{} `type:"structure"`

	// A structure containing details about the server certificate.
	//
	// ServerCertificate is a required field
	ServerCertificate *ServerCertificate `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetServerCertificateOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetServerCertificate = "GetServerCertificate"

// GetServerCertificateRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Retrieves information about the specified server certificate stored in IAM.
//
// For more information about working with server certificates, see Working
// with Server Certificates (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_server-certs.html)
// in the IAM User Guide. This topic includes a list of AWS services that can
// use the server certificates that you manage with IAM.
//
//    // Example sending a request using GetServerCertificateRequest.
//    req := client.GetServerCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/GetServerCertificate
func (c *Client) GetServerCertificateRequest(input *GetServerCertificateInput) GetServerCertificateRequest {
	op := &aws.Operation{
		Name:       opGetServerCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetServerCertificateInput{}
	}

	req := c.newRequest(op, input, &GetServerCertificateOutput{})
	return GetServerCertificateRequest{Request: req, Input: input, Copy: c.GetServerCertificateRequest}
}

// GetServerCertificateRequest is the request type for the
// GetServerCertificate API operation.
type GetServerCertificateRequest struct {
	*aws.Request
	Input *GetServerCertificateInput
	Copy  func(*GetServerCertificateInput) GetServerCertificateRequest
}

// Send marshals and sends the GetServerCertificate API request.
func (r GetServerCertificateRequest) Send(ctx context.Context) (*GetServerCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetServerCertificateResponse{
		GetServerCertificateOutput: r.Request.Data.(*GetServerCertificateOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetServerCertificateResponse is the response type for the
// GetServerCertificate API operation.
type GetServerCertificateResponse struct {
	*GetServerCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetServerCertificate request.
func (r *GetServerCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
