// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/ListObjectsRequest
type ListObjectsInput struct {
	_ struct{} `type:"structure"`

	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// A delimiter is a character you use to group keys.
	Delimiter *string `location:"querystring" locationName:"delimiter" type:"string"`

	// Requests Amazon S3 to encode the object keys in the response and specifies
	// the encoding method to use. An object key may contain any Unicode character;
	// however, XML 1.0 parser cannot parse some characters, such as characters
	// with an ASCII value from 0 to 10. For characters that are not supported in
	// XML 1.0, you can add this parameter to request that Amazon S3 encode the
	// keys in the response.
	EncodingType EncodingType `location:"querystring" locationName:"encoding-type" type:"string" enum:"true"`

	// Specifies the key to start with when listing objects in a bucket.
	Marker *string `location:"querystring" locationName:"marker" type:"string"`

	// Sets the maximum number of keys returned in the response. The response might
	// contain fewer keys but will never contain more.
	MaxKeys *int64 `location:"querystring" locationName:"max-keys" type:"integer"`

	// Limits the response to keys that begin with the specified prefix.
	Prefix *string `location:"querystring" locationName:"prefix" type:"string"`

	// Confirms that the requester knows that she or he will be charged for the
	// list objects request. Bucket owners need not specify this parameter in their
	// requests.
	RequestPayer RequestPayer `location:"header" locationName:"x-amz-request-payer" type:"string" enum:"true"`
}

// String returns the string representation
func (s ListObjectsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListObjectsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListObjectsInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *ListObjectsInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListObjectsInput) MarshalFields(e protocol.FieldEncoder) error {

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
	if s.Delimiter != nil {
		v := *s.Delimiter

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "delimiter", protocol.StringValue(v), metadata)
	}
	if len(s.EncodingType) > 0 {
		v := s.EncodingType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "encoding-type", v, metadata)
	}
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "marker", protocol.StringValue(v), metadata)
	}
	if s.MaxKeys != nil {
		v := *s.MaxKeys

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "max-keys", protocol.Int64Value(v), metadata)
	}
	if s.Prefix != nil {
		v := *s.Prefix

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "prefix", protocol.StringValue(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/ListObjectsOutput
type ListObjectsOutput struct {
	_ struct{} `type:"structure"`

	CommonPrefixes []CommonPrefix `type:"list" flattened:"true"`

	Contents []Object `type:"list" flattened:"true"`

	Delimiter *string `type:"string"`

	// Encoding type used by Amazon S3 to encode object keys in the response.
	EncodingType EncodingType `type:"string" enum:"true"`

	// A flag that indicates whether or not Amazon S3 returned all of the results
	// that satisfied the search criteria.
	IsTruncated *bool `type:"boolean"`

	Marker *string `type:"string"`

	MaxKeys *int64 `type:"integer"`

	Name *string `type:"string"`

	// When response is truncated (the IsTruncated element value in the response
	// is true), you can use the key name in this field as marker in the subsequent
	// request to get next set of objects. Amazon S3 lists objects in alphabetical
	// order Note: This element is returned only if you have delimiter request parameter
	// specified. If response does not include the NextMaker and it is truncated,
	// you can use the value of the last Key in the response as the marker in the
	// subsequent request to get the next set of object keys.
	NextMarker *string `type:"string"`

	Prefix *string `type:"string"`
}

// String returns the string representation
func (s ListObjectsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListObjectsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.CommonPrefixes) > 0 {
		v := s.CommonPrefixes

		metadata := protocol.Metadata{Flatten: true}
		ls0 := e.List(protocol.BodyTarget, "CommonPrefixes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if len(s.Contents) > 0 {
		v := s.Contents

		metadata := protocol.Metadata{Flatten: true}
		ls0 := e.List(protocol.BodyTarget, "Contents", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Delimiter != nil {
		v := *s.Delimiter

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Delimiter", protocol.StringValue(v), metadata)
	}
	if len(s.EncodingType) > 0 {
		v := s.EncodingType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "EncodingType", v, metadata)
	}
	if s.IsTruncated != nil {
		v := *s.IsTruncated

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IsTruncated", protocol.BoolValue(v), metadata)
	}
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Marker", protocol.StringValue(v), metadata)
	}
	if s.MaxKeys != nil {
		v := *s.MaxKeys

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaxKeys", protocol.Int64Value(v), metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.StringValue(v), metadata)
	}
	if s.NextMarker != nil {
		v := *s.NextMarker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextMarker", protocol.StringValue(v), metadata)
	}
	if s.Prefix != nil {
		v := *s.Prefix

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Prefix", protocol.StringValue(v), metadata)
	}
	return nil
}

const opListObjects = "ListObjects"

// ListObjectsRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Returns some or all (up to 1000) of the objects in a bucket. You can use
// the request parameters as selection criteria to return a subset of the objects
// in a bucket.
//
//    // Example sending a request using ListObjectsRequest.
//    req := client.ListObjectsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/ListObjects
func (c *Client) ListObjectsRequest(input *ListObjectsInput) ListObjectsRequest {
	op := &aws.Operation{
		Name:       opListObjects,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"NextMarker || Contents[-1].Key"},
			LimitToken:      "MaxKeys",
			TruncationToken: "IsTruncated",
		},
	}

	if input == nil {
		input = &ListObjectsInput{}
	}

	req := c.newRequest(op, input, &ListObjectsOutput{})
	return ListObjectsRequest{Request: req, Input: input, Copy: c.ListObjectsRequest}
}

// ListObjectsRequest is the request type for the
// ListObjects API operation.
type ListObjectsRequest struct {
	*aws.Request
	Input *ListObjectsInput
	Copy  func(*ListObjectsInput) ListObjectsRequest
}

// Send marshals and sends the ListObjects API request.
func (r ListObjectsRequest) Send(ctx context.Context) (*ListObjectsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListObjectsResponse{
		ListObjectsOutput: r.Request.Data.(*ListObjectsOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListObjectsRequestPaginator returns a paginator for ListObjects.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListObjectsRequest(input)
//   p := s3.NewListObjectsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListObjectsPaginator(req ListObjectsRequest) ListObjectsPaginator {
	return ListObjectsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListObjectsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListObjectsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListObjectsPaginator struct {
	aws.Pager
}

func (p *ListObjectsPaginator) CurrentPage() *ListObjectsOutput {
	return p.Pager.CurrentPage().(*ListObjectsOutput)
}

// ListObjectsResponse is the response type for the
// ListObjects API operation.
type ListObjectsResponse struct {
	*ListObjectsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListObjects request.
func (r *ListObjectsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
