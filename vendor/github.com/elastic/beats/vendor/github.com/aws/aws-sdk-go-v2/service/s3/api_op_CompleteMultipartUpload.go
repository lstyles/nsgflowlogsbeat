// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/CompleteMultipartUploadRequest
type CompleteMultipartUploadInput struct {
	_ struct{} `type:"structure" payload:"MultipartUpload"`

	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// Key is a required field
	Key *string `location:"uri" locationName:"Key" min:"1" type:"string" required:"true"`

	MultipartUpload *CompletedMultipartUpload `locationName:"CompleteMultipartUpload" type:"structure" xmlURI:"http://s3.amazonaws.com/doc/2006-03-01/"`

	// Confirms that the requester knows that she or he will be charged for the
	// request. Bucket owners need not specify this parameter in their requests.
	// Documentation on downloading objects from requester pays buckets can be found
	// at http://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html
	RequestPayer RequestPayer `location:"header" locationName:"x-amz-request-payer" type:"string" enum:"true"`

	// UploadId is a required field
	UploadId *string `location:"querystring" locationName:"uploadId" type:"string" required:"true"`
}

// String returns the string representation
func (s CompleteMultipartUploadInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CompleteMultipartUploadInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CompleteMultipartUploadInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if s.UploadId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UploadId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *CompleteMultipartUploadInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CompleteMultipartUploadInput) MarshalFields(e protocol.FieldEncoder) error {

	if len(s.RequestPayer) > 0 {
		v := s.RequestPayer

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-payer", v, metadata)
	}
	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.Key != nil {
		v := *s.Key

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Key", protocol.StringValue(v), metadata)
	}
	if s.MultipartUpload != nil {
		v := s.MultipartUpload

		metadata := protocol.Metadata{XMLNamespaceURI: "http://s3.amazonaws.com/doc/2006-03-01/"}
		e.SetFields(protocol.PayloadTarget, "CompleteMultipartUpload", v, metadata)
	}
	if s.UploadId != nil {
		v := *s.UploadId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "uploadId", protocol.StringValue(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/CompleteMultipartUploadOutput
type CompleteMultipartUploadOutput struct {
	_ struct{} `type:"structure"`

	Bucket *string `type:"string"`

	// Entity tag of the object.
	ETag *string `type:"string"`

	// If the object expiration is configured, this will contain the expiration
	// date (expiry-date) and rule ID (rule-id). The value of rule-id is URL encoded.
	Expiration *string `location:"header" locationName:"x-amz-expiration" type:"string"`

	Key *string `min:"1" type:"string"`

	Location *string `type:"string"`

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged RequestCharged `location:"header" locationName:"x-amz-request-charged" type:"string" enum:"true"`

	// If present, specifies the ID of the AWS Key Management Service (KMS) master
	// encryption key that was used for the object.
	SSEKMSKeyId *string `location:"header" locationName:"x-amz-server-side-encryption-aws-kms-key-id" type:"string"`

	// The Server-side encryption algorithm used when storing this object in S3
	// (e.g., AES256, aws:kms).
	ServerSideEncryption ServerSideEncryption `location:"header" locationName:"x-amz-server-side-encryption" type:"string" enum:"true"`

	// Version of the object.
	VersionId *string `location:"header" locationName:"x-amz-version-id" type:"string"`
}

// String returns the string representation
func (s CompleteMultipartUploadOutput) String() string {
	return awsutil.Prettify(s)
}

func (s *CompleteMultipartUploadOutput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CompleteMultipartUploadOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.ETag != nil {
		v := *s.ETag

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ETag", protocol.StringValue(v), metadata)
	}
	if s.Key != nil {
		v := *s.Key

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Key", protocol.StringValue(v), metadata)
	}
	if s.Location != nil {
		v := *s.Location

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Location", protocol.StringValue(v), metadata)
	}
	if s.Expiration != nil {
		v := *s.Expiration

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-expiration", protocol.StringValue(v), metadata)
	}
	if len(s.RequestCharged) > 0 {
		v := s.RequestCharged

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-charged", v, metadata)
	}
	if s.SSEKMSKeyId != nil {
		v := *s.SSEKMSKeyId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-aws-kms-key-id", protocol.StringValue(v), metadata)
	}
	if len(s.ServerSideEncryption) > 0 {
		v := s.ServerSideEncryption

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption", v, metadata)
	}
	if s.VersionId != nil {
		v := *s.VersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-version-id", protocol.StringValue(v), metadata)
	}
	return nil
}

const opCompleteMultipartUpload = "CompleteMultipartUpload"

// CompleteMultipartUploadRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Completes a multipart upload by assembling previously uploaded parts.
//
//    // Example sending a request using CompleteMultipartUploadRequest.
//    req := client.CompleteMultipartUploadRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/CompleteMultipartUpload
func (c *Client) CompleteMultipartUploadRequest(input *CompleteMultipartUploadInput) CompleteMultipartUploadRequest {
	op := &aws.Operation{
		Name:       opCompleteMultipartUpload,
		HTTPMethod: "POST",
		HTTPPath:   "/{Bucket}/{Key+}",
	}

	if input == nil {
		input = &CompleteMultipartUploadInput{}
	}

	req := c.newRequest(op, input, &CompleteMultipartUploadOutput{})
	return CompleteMultipartUploadRequest{Request: req, Input: input, Copy: c.CompleteMultipartUploadRequest}
}

// CompleteMultipartUploadRequest is the request type for the
// CompleteMultipartUpload API operation.
type CompleteMultipartUploadRequest struct {
	*aws.Request
	Input *CompleteMultipartUploadInput
	Copy  func(*CompleteMultipartUploadInput) CompleteMultipartUploadRequest
}

// Send marshals and sends the CompleteMultipartUpload API request.
func (r CompleteMultipartUploadRequest) Send(ctx context.Context) (*CompleteMultipartUploadResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CompleteMultipartUploadResponse{
		CompleteMultipartUploadOutput: r.Request.Data.(*CompleteMultipartUploadOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CompleteMultipartUploadResponse is the response type for the
// CompleteMultipartUpload API operation.
type CompleteMultipartUploadResponse struct {
	*CompleteMultipartUploadOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CompleteMultipartUpload request.
func (r *CompleteMultipartUploadResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
