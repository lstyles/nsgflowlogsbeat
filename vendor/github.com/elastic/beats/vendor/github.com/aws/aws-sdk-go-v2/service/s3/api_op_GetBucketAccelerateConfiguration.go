// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetBucketAccelerateConfigurationRequest
type GetBucketAccelerateConfigurationInput struct {
	_ struct{} `type:"structure"`

	// Name of the bucket for which the accelerate configuration is retrieved.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`
}

// String returns the string representation
func (s GetBucketAccelerateConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBucketAccelerateConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBucketAccelerateConfigurationInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *GetBucketAccelerateConfigurationInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBucketAccelerateConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetBucketAccelerateConfigurationOutput
type GetBucketAccelerateConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The accelerate configuration of the bucket.
	Status BucketAccelerateStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s GetBucketAccelerateConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBucketAccelerateConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Status", v, metadata)
	}
	return nil
}

const opGetBucketAccelerateConfiguration = "GetBucketAccelerateConfiguration"

// GetBucketAccelerateConfigurationRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Returns the accelerate configuration of a bucket.
//
//    // Example sending a request using GetBucketAccelerateConfigurationRequest.
//    req := client.GetBucketAccelerateConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetBucketAccelerateConfiguration
func (c *Client) GetBucketAccelerateConfigurationRequest(input *GetBucketAccelerateConfigurationInput) GetBucketAccelerateConfigurationRequest {
	op := &aws.Operation{
		Name:       opGetBucketAccelerateConfiguration,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?accelerate",
	}

	if input == nil {
		input = &GetBucketAccelerateConfigurationInput{}
	}

	req := c.newRequest(op, input, &GetBucketAccelerateConfigurationOutput{})
	return GetBucketAccelerateConfigurationRequest{Request: req, Input: input, Copy: c.GetBucketAccelerateConfigurationRequest}
}

// GetBucketAccelerateConfigurationRequest is the request type for the
// GetBucketAccelerateConfiguration API operation.
type GetBucketAccelerateConfigurationRequest struct {
	*aws.Request
	Input *GetBucketAccelerateConfigurationInput
	Copy  func(*GetBucketAccelerateConfigurationInput) GetBucketAccelerateConfigurationRequest
}

// Send marshals and sends the GetBucketAccelerateConfiguration API request.
func (r GetBucketAccelerateConfigurationRequest) Send(ctx context.Context) (*GetBucketAccelerateConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBucketAccelerateConfigurationResponse{
		GetBucketAccelerateConfigurationOutput: r.Request.Data.(*GetBucketAccelerateConfigurationOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBucketAccelerateConfigurationResponse is the response type for the
// GetBucketAccelerateConfiguration API operation.
type GetBucketAccelerateConfigurationResponse struct {
	*GetBucketAccelerateConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBucketAccelerateConfiguration request.
func (r *GetBucketAccelerateConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
