// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeElasticGpusRequest
type DescribeElasticGpusInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The Elastic Graphics accelerator IDs.
	ElasticGpuIds []string `locationName:"ElasticGpuId" locationNameList:"item" type:"list"`

	// The filters.
	//
	//    * availability-zone - The Availability Zone in which the Elastic Graphics
	//    accelerator resides.
	//
	//    * elastic-gpu-health - The status of the Elastic Graphics accelerator
	//    (OK | IMPAIRED).
	//
	//    * elastic-gpu-state - The state of the Elastic Graphics accelerator (ATTACHED).
	//
	//    * elastic-gpu-type - The type of Elastic Graphics accelerator; for example,
	//    eg1.medium.
	//
	//    * instance-id - The ID of the instance to which the Elastic Graphics accelerator
	//    is associated.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value. This
	// value can be between 5 and 1000.
	MaxResults *int64 `min:"10" type:"integer"`

	// The token to request the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeElasticGpusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeElasticGpusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeElasticGpusInput"}
	if s.MaxResults != nil && *s.MaxResults < 10 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 10))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeElasticGpusResult
type DescribeElasticGpusOutput struct {
	_ struct{} `type:"structure"`

	// Information about the Elastic Graphics accelerators.
	ElasticGpuSet []ElasticGpus `locationName:"elasticGpuSet" locationNameList:"item" type:"list"`

	// The total number of items to return. If the total number of items available
	// is more than the value specified in max-items then a Next-Token will be provided
	// in the output that you can use to resume pagination.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeElasticGpusOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeElasticGpus = "DescribeElasticGpus"

// DescribeElasticGpusRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the Elastic Graphics accelerator associated with your instances.
// For more information about Elastic Graphics, see Amazon Elastic Graphics
// (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/elastic-graphics.html).
//
//    // Example sending a request using DescribeElasticGpusRequest.
//    req := client.DescribeElasticGpusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeElasticGpus
func (c *Client) DescribeElasticGpusRequest(input *DescribeElasticGpusInput) DescribeElasticGpusRequest {
	op := &aws.Operation{
		Name:       opDescribeElasticGpus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeElasticGpusInput{}
	}

	req := c.newRequest(op, input, &DescribeElasticGpusOutput{})
	return DescribeElasticGpusRequest{Request: req, Input: input, Copy: c.DescribeElasticGpusRequest}
}

// DescribeElasticGpusRequest is the request type for the
// DescribeElasticGpus API operation.
type DescribeElasticGpusRequest struct {
	*aws.Request
	Input *DescribeElasticGpusInput
	Copy  func(*DescribeElasticGpusInput) DescribeElasticGpusRequest
}

// Send marshals and sends the DescribeElasticGpus API request.
func (r DescribeElasticGpusRequest) Send(ctx context.Context) (*DescribeElasticGpusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeElasticGpusResponse{
		DescribeElasticGpusOutput: r.Request.Data.(*DescribeElasticGpusOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeElasticGpusResponse is the response type for the
// DescribeElasticGpus API operation.
type DescribeElasticGpusResponse struct {
	*DescribeElasticGpusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeElasticGpus request.
func (r *DescribeElasticGpusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
