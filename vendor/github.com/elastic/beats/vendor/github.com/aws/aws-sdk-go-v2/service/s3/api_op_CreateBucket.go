// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/CreateBucketRequest
type CreateBucketInput struct {
	_ struct{} `type:"structure" payload:"CreateBucketConfiguration"`

	// The canned ACL to apply to the bucket.
	ACL BucketCannedACL `location:"header" locationName:"x-amz-acl" type:"string" enum:"true"`

	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	CreateBucketConfiguration *CreateBucketConfiguration `locationName:"CreateBucketConfiguration" type:"structure" xmlURI:"http://s3.amazonaws.com/doc/2006-03-01/"`

	// Allows grantee the read, write, read ACP, and write ACP permissions on the
	// bucket.
	GrantFullControl *string `location:"header" locationName:"x-amz-grant-full-control" type:"string"`

	// Allows grantee to list the objects in the bucket.
	GrantRead *string `location:"header" locationName:"x-amz-grant-read" type:"string"`

	// Allows grantee to read the bucket ACL.
	GrantReadACP *string `location:"header" locationName:"x-amz-grant-read-acp" type:"string"`

	// Allows grantee to create, overwrite, and delete any object in the bucket.
	GrantWrite *string `location:"header" locationName:"x-amz-grant-write" type:"string"`

	// Allows grantee to write the ACL for the applicable bucket.
	GrantWriteACP *string `location:"header" locationName:"x-amz-grant-write-acp" type:"string"`

	// Specifies whether you want S3 Object Lock to be enabled for the new bucket.
	ObjectLockEnabledForBucket *bool `location:"header" locationName:"x-amz-bucket-object-lock-enabled" type:"boolean"`
}

// String returns the string representation
func (s CreateBucketInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateBucketInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateBucketInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *CreateBucketInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateBucketInput) MarshalFields(e protocol.FieldEncoder) error {

	if len(s.ACL) > 0 {
		v := s.ACL

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-acl", v, metadata)
	}
	if s.GrantFullControl != nil {
		v := *s.GrantFullControl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-grant-full-control", protocol.StringValue(v), metadata)
	}
	if s.GrantRead != nil {
		v := *s.GrantRead

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-grant-read", protocol.StringValue(v), metadata)
	}
	if s.GrantReadACP != nil {
		v := *s.GrantReadACP

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-grant-read-acp", protocol.StringValue(v), metadata)
	}
	if s.GrantWrite != nil {
		v := *s.GrantWrite

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-grant-write", protocol.StringValue(v), metadata)
	}
	if s.GrantWriteACP != nil {
		v := *s.GrantWriteACP

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-grant-write-acp", protocol.StringValue(v), metadata)
	}
	if s.ObjectLockEnabledForBucket != nil {
		v := *s.ObjectLockEnabledForBucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-bucket-object-lock-enabled", protocol.BoolValue(v), metadata)
	}
	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.CreateBucketConfiguration != nil {
		v := s.CreateBucketConfiguration

		metadata := protocol.Metadata{XMLNamespaceURI: "http://s3.amazonaws.com/doc/2006-03-01/"}
		e.SetFields(protocol.PayloadTarget, "CreateBucketConfiguration", v, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/CreateBucketOutput
type CreateBucketOutput struct {
	_ struct{} `type:"structure"`

	Location *string `location:"header" locationName:"Location" type:"string"`
}

// String returns the string representation
func (s CreateBucketOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateBucketOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Location != nil {
		v := *s.Location

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Location", protocol.StringValue(v), metadata)
	}
	return nil
}

const opCreateBucket = "CreateBucket"

// CreateBucketRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Creates a new bucket.
//
//    // Example sending a request using CreateBucketRequest.
//    req := client.CreateBucketRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/CreateBucket
func (c *Client) CreateBucketRequest(input *CreateBucketInput) CreateBucketRequest {
	op := &aws.Operation{
		Name:       opCreateBucket,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}",
	}

	if input == nil {
		input = &CreateBucketInput{}
	}

	req := c.newRequest(op, input, &CreateBucketOutput{})
	return CreateBucketRequest{Request: req, Input: input, Copy: c.CreateBucketRequest}
}

// CreateBucketRequest is the request type for the
// CreateBucket API operation.
type CreateBucketRequest struct {
	*aws.Request
	Input *CreateBucketInput
	Copy  func(*CreateBucketInput) CreateBucketRequest
}

// Send marshals and sends the CreateBucket API request.
func (r CreateBucketRequest) Send(ctx context.Context) (*CreateBucketResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateBucketResponse{
		CreateBucketOutput: r.Request.Data.(*CreateBucketOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateBucketResponse is the response type for the
// CreateBucket API operation.
type CreateBucketResponse struct {
	*CreateBucketOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateBucket request.
func (r *CreateBucketResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
