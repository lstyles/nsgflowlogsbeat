// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutObjectLegalHoldRequest
type PutObjectLegalHoldInput struct {
	_ struct{} `type:"structure" payload:"LegalHold"`

	// The bucket containing the object that you want to place a Legal Hold on.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// The key name for the object that you want to place a Legal Hold on.
	//
	// Key is a required field
	Key *string `location:"uri" locationName:"Key" min:"1" type:"string" required:"true"`

	// Container element for the Legal Hold configuration you want to apply to the
	// specified object.
	LegalHold *ObjectLockLegalHold `locationName:"LegalHold" type:"structure" xmlURI:"http://s3.amazonaws.com/doc/2006-03-01/"`

	// Confirms that the requester knows that she or he will be charged for the
	// request. Bucket owners need not specify this parameter in their requests.
	// Documentation on downloading objects from requester pays buckets can be found
	// at http://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html
	RequestPayer RequestPayer `location:"header" locationName:"x-amz-request-payer" type:"string" enum:"true"`

	// The version ID of the object that you want to place a Legal Hold on.
	VersionId *string `location:"querystring" locationName:"versionId" type:"string"`
}

// String returns the string representation
func (s PutObjectLegalHoldInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutObjectLegalHoldInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutObjectLegalHoldInput"}

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

func (s *PutObjectLegalHoldInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutObjectLegalHoldInput) MarshalFields(e protocol.FieldEncoder) error {

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
	if s.LegalHold != nil {
		v := s.LegalHold

		metadata := protocol.Metadata{XMLNamespaceURI: "http://s3.amazonaws.com/doc/2006-03-01/"}
		e.SetFields(protocol.PayloadTarget, "LegalHold", v, metadata)
	}
	if s.VersionId != nil {
		v := *s.VersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "versionId", protocol.StringValue(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutObjectLegalHoldOutput
type PutObjectLegalHoldOutput struct {
	_ struct{} `type:"structure"`

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged RequestCharged `location:"header" locationName:"x-amz-request-charged" type:"string" enum:"true"`
}

// String returns the string representation
func (s PutObjectLegalHoldOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutObjectLegalHoldOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.RequestCharged) > 0 {
		v := s.RequestCharged

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-charged", v, metadata)
	}
	return nil
}

const opPutObjectLegalHold = "PutObjectLegalHold"

// PutObjectLegalHoldRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Applies a Legal Hold configuration to the specified object.
//
//    // Example sending a request using PutObjectLegalHoldRequest.
//    req := client.PutObjectLegalHoldRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutObjectLegalHold
func (c *Client) PutObjectLegalHoldRequest(input *PutObjectLegalHoldInput) PutObjectLegalHoldRequest {
	op := &aws.Operation{
		Name:       opPutObjectLegalHold,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}/{Key+}?legal-hold",
	}

	if input == nil {
		input = &PutObjectLegalHoldInput{}
	}

	req := c.newRequest(op, input, &PutObjectLegalHoldOutput{})
	return PutObjectLegalHoldRequest{Request: req, Input: input, Copy: c.PutObjectLegalHoldRequest}
}

// PutObjectLegalHoldRequest is the request type for the
// PutObjectLegalHold API operation.
type PutObjectLegalHoldRequest struct {
	*aws.Request
	Input *PutObjectLegalHoldInput
	Copy  func(*PutObjectLegalHoldInput) PutObjectLegalHoldRequest
}

// Send marshals and sends the PutObjectLegalHold API request.
func (r PutObjectLegalHoldRequest) Send(ctx context.Context) (*PutObjectLegalHoldResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutObjectLegalHoldResponse{
		PutObjectLegalHoldOutput: r.Request.Data.(*PutObjectLegalHoldOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutObjectLegalHoldResponse is the response type for the
// PutObjectLegalHold API operation.
type PutObjectLegalHoldResponse struct {
	*PutObjectLegalHoldOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutObjectLegalHold request.
func (r *PutObjectLegalHoldResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
