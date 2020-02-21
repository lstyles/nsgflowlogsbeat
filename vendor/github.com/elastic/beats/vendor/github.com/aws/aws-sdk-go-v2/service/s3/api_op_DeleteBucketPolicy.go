// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/DeleteBucketPolicyRequest
type DeleteBucketPolicyInput struct {
	_ struct{} `type:"structure"`

	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteBucketPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteBucketPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteBucketPolicyInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *DeleteBucketPolicyInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBucketPolicyInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/DeleteBucketPolicyOutput
type DeleteBucketPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteBucketPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBucketPolicyOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteBucketPolicy = "DeleteBucketPolicy"

// DeleteBucketPolicyRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Deletes the policy from the bucket.
//
//    // Example sending a request using DeleteBucketPolicyRequest.
//    req := client.DeleteBucketPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/DeleteBucketPolicy
func (c *Client) DeleteBucketPolicyRequest(input *DeleteBucketPolicyInput) DeleteBucketPolicyRequest {
	op := &aws.Operation{
		Name:       opDeleteBucketPolicy,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}?policy",
	}

	if input == nil {
		input = &DeleteBucketPolicyInput{}
	}

	req := c.newRequest(op, input, &DeleteBucketPolicyOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteBucketPolicyRequest{Request: req, Input: input, Copy: c.DeleteBucketPolicyRequest}
}

// DeleteBucketPolicyRequest is the request type for the
// DeleteBucketPolicy API operation.
type DeleteBucketPolicyRequest struct {
	*aws.Request
	Input *DeleteBucketPolicyInput
	Copy  func(*DeleteBucketPolicyInput) DeleteBucketPolicyRequest
}

// Send marshals and sends the DeleteBucketPolicy API request.
func (r DeleteBucketPolicyRequest) Send(ctx context.Context) (*DeleteBucketPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteBucketPolicyResponse{
		DeleteBucketPolicyOutput: r.Request.Data.(*DeleteBucketPolicyOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteBucketPolicyResponse is the response type for the
// DeleteBucketPolicy API operation.
type DeleteBucketPolicyResponse struct {
	*DeleteBucketPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteBucketPolicy request.
func (r *DeleteBucketPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
