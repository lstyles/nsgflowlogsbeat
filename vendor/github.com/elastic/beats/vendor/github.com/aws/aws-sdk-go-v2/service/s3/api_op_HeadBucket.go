// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/HeadBucketRequest
type HeadBucketInput struct {
	_ struct{} `type:"structure"`

	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`
}

// String returns the string representation
func (s HeadBucketInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *HeadBucketInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "HeadBucketInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *HeadBucketInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s HeadBucketInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/HeadBucketOutput
type HeadBucketOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s HeadBucketOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s HeadBucketOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opHeadBucket = "HeadBucket"

// HeadBucketRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// This operation is useful to determine if a bucket exists and you have permission
// to access it.
//
//    // Example sending a request using HeadBucketRequest.
//    req := client.HeadBucketRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/HeadBucket
func (c *Client) HeadBucketRequest(input *HeadBucketInput) HeadBucketRequest {
	op := &aws.Operation{
		Name:       opHeadBucket,
		HTTPMethod: "HEAD",
		HTTPPath:   "/{Bucket}",
	}

	if input == nil {
		input = &HeadBucketInput{}
	}

	req := c.newRequest(op, input, &HeadBucketOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return HeadBucketRequest{Request: req, Input: input, Copy: c.HeadBucketRequest}
}

// HeadBucketRequest is the request type for the
// HeadBucket API operation.
type HeadBucketRequest struct {
	*aws.Request
	Input *HeadBucketInput
	Copy  func(*HeadBucketInput) HeadBucketRequest
}

// Send marshals and sends the HeadBucket API request.
func (r HeadBucketRequest) Send(ctx context.Context) (*HeadBucketResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &HeadBucketResponse{
		HeadBucketOutput: r.Request.Data.(*HeadBucketOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// HeadBucketResponse is the response type for the
// HeadBucket API operation.
type HeadBucketResponse struct {
	*HeadBucketOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// HeadBucket request.
func (r *HeadBucketResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
