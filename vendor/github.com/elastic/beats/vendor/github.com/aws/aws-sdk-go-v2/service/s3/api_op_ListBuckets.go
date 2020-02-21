// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/ListBucketsInput
type ListBucketsInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ListBucketsInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListBucketsInput) MarshalFields(e protocol.FieldEncoder) error {

	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/ListBucketsOutput
type ListBucketsOutput struct {
	_ struct{} `type:"structure"`

	Buckets []Bucket `locationNameList:"Bucket" type:"list"`

	Owner *Owner `type:"structure"`
}

// String returns the string representation
func (s ListBucketsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListBucketsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.Buckets) > 0 {
		v := s.Buckets

		metadata := protocol.Metadata{ListLocationName: "Bucket"}
		ls0 := e.List(protocol.BodyTarget, "Buckets", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Owner != nil {
		v := s.Owner

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Owner", v, metadata)
	}
	return nil
}

const opListBuckets = "ListBuckets"

// ListBucketsRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Returns a list of all buckets owned by the authenticated sender of the request.
//
//    // Example sending a request using ListBucketsRequest.
//    req := client.ListBucketsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/ListBuckets
func (c *Client) ListBucketsRequest(input *ListBucketsInput) ListBucketsRequest {
	op := &aws.Operation{
		Name:       opListBuckets,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListBucketsInput{}
	}

	req := c.newRequest(op, input, &ListBucketsOutput{})
	return ListBucketsRequest{Request: req, Input: input, Copy: c.ListBucketsRequest}
}

// ListBucketsRequest is the request type for the
// ListBuckets API operation.
type ListBucketsRequest struct {
	*aws.Request
	Input *ListBucketsInput
	Copy  func(*ListBucketsInput) ListBucketsRequest
}

// Send marshals and sends the ListBuckets API request.
func (r ListBucketsRequest) Send(ctx context.Context) (*ListBucketsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListBucketsResponse{
		ListBucketsOutput: r.Request.Data.(*ListBucketsOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListBucketsResponse is the response type for the
// ListBuckets API operation.
type ListBucketsResponse struct {
	*ListBucketsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListBuckets request.
func (r *ListBucketsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
