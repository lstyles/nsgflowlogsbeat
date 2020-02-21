// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutObjectRetentionRequest
type PutObjectRetentionInput struct {
	_ struct{} `type:"structure" payload:"Retention"`

	// The bucket that contains the object you want to apply this Object Retention
	// configuration to.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// Indicates whether this operation should bypass Governance-mode restrictions.j
	BypassGovernanceRetention *bool `location:"header" locationName:"x-amz-bypass-governance-retention" type:"boolean"`

	// The key name for the object that you want to apply this Object Retention
	// configuration to.
	//
	// Key is a required field
	Key *string `location:"uri" locationName:"Key" min:"1" type:"string" required:"true"`

	// Confirms that the requester knows that she or he will be charged for the
	// request. Bucket owners need not specify this parameter in their requests.
	// Documentation on downloading objects from requester pays buckets can be found
	// at http://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html
	RequestPayer RequestPayer `location:"header" locationName:"x-amz-request-payer" type:"string" enum:"true"`

	// The container element for the Object Retention configuration.
	Retention *ObjectLockRetention `locationName:"Retention" type:"structure" xmlURI:"http://s3.amazonaws.com/doc/2006-03-01/"`

	// The version ID for the object that you want to apply this Object Retention
	// configuration to.
	VersionId *string `location:"querystring" locationName:"versionId" type:"string"`
}

// String returns the string representation
func (s PutObjectRetentionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutObjectRetentionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutObjectRetentionInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *PutObjectRetentionInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutObjectRetentionInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.BypassGovernanceRetention != nil {
		v := *s.BypassGovernanceRetention

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-bypass-governance-retention", protocol.BoolValue(v), metadata)
	}
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
	if s.Retention != nil {
		v := s.Retention

		metadata := protocol.Metadata{XMLNamespaceURI: "http://s3.amazonaws.com/doc/2006-03-01/"}
		e.SetFields(protocol.PayloadTarget, "Retention", v, metadata)
	}
	if s.VersionId != nil {
		v := *s.VersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "versionId", protocol.StringValue(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutObjectRetentionOutput
type PutObjectRetentionOutput struct {
	_ struct{} `type:"structure"`

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged RequestCharged `location:"header" locationName:"x-amz-request-charged" type:"string" enum:"true"`
}

// String returns the string representation
func (s PutObjectRetentionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutObjectRetentionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.RequestCharged) > 0 {
		v := s.RequestCharged

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-charged", v, metadata)
	}
	return nil
}

const opPutObjectRetention = "PutObjectRetention"

// PutObjectRetentionRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Places an Object Retention configuration on an object.
//
//    // Example sending a request using PutObjectRetentionRequest.
//    req := client.PutObjectRetentionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutObjectRetention
func (c *Client) PutObjectRetentionRequest(input *PutObjectRetentionInput) PutObjectRetentionRequest {
	op := &aws.Operation{
		Name:       opPutObjectRetention,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}/{Key+}?retention",
	}

	if input == nil {
		input = &PutObjectRetentionInput{}
	}

	req := c.newRequest(op, input, &PutObjectRetentionOutput{})
	return PutObjectRetentionRequest{Request: req, Input: input, Copy: c.PutObjectRetentionRequest}
}

// PutObjectRetentionRequest is the request type for the
// PutObjectRetention API operation.
type PutObjectRetentionRequest struct {
	*aws.Request
	Input *PutObjectRetentionInput
	Copy  func(*PutObjectRetentionInput) PutObjectRetentionRequest
}

// Send marshals and sends the PutObjectRetention API request.
func (r PutObjectRetentionRequest) Send(ctx context.Context) (*PutObjectRetentionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutObjectRetentionResponse{
		PutObjectRetentionOutput: r.Request.Data.(*PutObjectRetentionOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutObjectRetentionResponse is the response type for the
// PutObjectRetention API operation.
type PutObjectRetentionResponse struct {
	*PutObjectRetentionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutObjectRetention request.
func (r *PutObjectRetentionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
