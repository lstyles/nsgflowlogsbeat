// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribePrefixListsRequest
type DescribePrefixListsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// One or more filters.
	//
	//    * prefix-list-id: The ID of a prefix list.
	//
	//    * prefix-list-name: The name of a prefix list.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `type:"integer"`

	// The token for the next page of results.
	NextToken *string `type:"string"`

	// One or more prefix list IDs.
	PrefixListIds []string `locationName:"PrefixListId" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribePrefixListsInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribePrefixListsResult
type DescribePrefixListsOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// All available prefix lists.
	PrefixLists []PrefixList `locationName:"prefixListSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribePrefixListsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribePrefixLists = "DescribePrefixLists"

// DescribePrefixListsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes available AWS services in a prefix list format, which includes
// the prefix list name and prefix list ID of the service and the IP address
// range for the service. A prefix list ID is required for creating an outbound
// security group rule that allows traffic from a VPC to access an AWS service
// through a gateway VPC endpoint. Currently, the services that support this
// action are Amazon S3 and Amazon DynamoDB.
//
//    // Example sending a request using DescribePrefixListsRequest.
//    req := client.DescribePrefixListsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribePrefixLists
func (c *Client) DescribePrefixListsRequest(input *DescribePrefixListsInput) DescribePrefixListsRequest {
	op := &aws.Operation{
		Name:       opDescribePrefixLists,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribePrefixListsInput{}
	}

	req := c.newRequest(op, input, &DescribePrefixListsOutput{})
	return DescribePrefixListsRequest{Request: req, Input: input, Copy: c.DescribePrefixListsRequest}
}

// DescribePrefixListsRequest is the request type for the
// DescribePrefixLists API operation.
type DescribePrefixListsRequest struct {
	*aws.Request
	Input *DescribePrefixListsInput
	Copy  func(*DescribePrefixListsInput) DescribePrefixListsRequest
}

// Send marshals and sends the DescribePrefixLists API request.
func (r DescribePrefixListsRequest) Send(ctx context.Context) (*DescribePrefixListsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribePrefixListsResponse{
		DescribePrefixListsOutput: r.Request.Data.(*DescribePrefixListsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribePrefixListsRequestPaginator returns a paginator for DescribePrefixLists.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribePrefixListsRequest(input)
//   p := ec2.NewDescribePrefixListsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribePrefixListsPaginator(req DescribePrefixListsRequest) DescribePrefixListsPaginator {
	return DescribePrefixListsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribePrefixListsInput
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

// DescribePrefixListsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribePrefixListsPaginator struct {
	aws.Pager
}

func (p *DescribePrefixListsPaginator) CurrentPage() *DescribePrefixListsOutput {
	return p.Pager.CurrentPage().(*DescribePrefixListsOutput)
}

// DescribePrefixListsResponse is the response type for the
// DescribePrefixLists API operation.
type DescribePrefixListsResponse struct {
	*DescribePrefixListsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribePrefixLists request.
func (r *DescribePrefixListsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
