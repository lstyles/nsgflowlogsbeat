// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for CreateSpotDatafeedSubscription.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateSpotDatafeedSubscriptionRequest
type CreateSpotDatafeedSubscriptionInput struct {
	_ struct{} `type:"structure"`

	// The Amazon S3 bucket in which to store the Spot Instance data feed.
	//
	// Bucket is a required field
	Bucket *string `locationName:"bucket" type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// A prefix for the data feed file names.
	Prefix *string `locationName:"prefix" type:"string"`
}

// String returns the string representation
func (s CreateSpotDatafeedSubscriptionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSpotDatafeedSubscriptionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateSpotDatafeedSubscriptionInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of CreateSpotDatafeedSubscription.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateSpotDatafeedSubscriptionResult
type CreateSpotDatafeedSubscriptionOutput struct {
	_ struct{} `type:"structure"`

	// The Spot Instance data feed subscription.
	SpotDatafeedSubscription *SpotDatafeedSubscription `locationName:"spotDatafeedSubscription" type:"structure"`
}

// String returns the string representation
func (s CreateSpotDatafeedSubscriptionOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateSpotDatafeedSubscription = "CreateSpotDatafeedSubscription"

// CreateSpotDatafeedSubscriptionRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates a data feed for Spot Instances, enabling you to view Spot Instance
// usage logs. You can create one data feed per AWS account. For more information,
// see Spot Instance Data Feed (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-data-feeds.html)
// in the Amazon EC2 User Guide for Linux Instances.
//
//    // Example sending a request using CreateSpotDatafeedSubscriptionRequest.
//    req := client.CreateSpotDatafeedSubscriptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateSpotDatafeedSubscription
func (c *Client) CreateSpotDatafeedSubscriptionRequest(input *CreateSpotDatafeedSubscriptionInput) CreateSpotDatafeedSubscriptionRequest {
	op := &aws.Operation{
		Name:       opCreateSpotDatafeedSubscription,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateSpotDatafeedSubscriptionInput{}
	}

	req := c.newRequest(op, input, &CreateSpotDatafeedSubscriptionOutput{})
	return CreateSpotDatafeedSubscriptionRequest{Request: req, Input: input, Copy: c.CreateSpotDatafeedSubscriptionRequest}
}

// CreateSpotDatafeedSubscriptionRequest is the request type for the
// CreateSpotDatafeedSubscription API operation.
type CreateSpotDatafeedSubscriptionRequest struct {
	*aws.Request
	Input *CreateSpotDatafeedSubscriptionInput
	Copy  func(*CreateSpotDatafeedSubscriptionInput) CreateSpotDatafeedSubscriptionRequest
}

// Send marshals and sends the CreateSpotDatafeedSubscription API request.
func (r CreateSpotDatafeedSubscriptionRequest) Send(ctx context.Context) (*CreateSpotDatafeedSubscriptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateSpotDatafeedSubscriptionResponse{
		CreateSpotDatafeedSubscriptionOutput: r.Request.Data.(*CreateSpotDatafeedSubscriptionOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateSpotDatafeedSubscriptionResponse is the response type for the
// CreateSpotDatafeedSubscription API operation.
type CreateSpotDatafeedSubscriptionResponse struct {
	*CreateSpotDatafeedSubscriptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateSpotDatafeedSubscription request.
func (r *CreateSpotDatafeedSubscriptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
