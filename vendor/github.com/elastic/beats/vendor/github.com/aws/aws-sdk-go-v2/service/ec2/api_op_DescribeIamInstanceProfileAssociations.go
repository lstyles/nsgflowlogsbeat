// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeIamInstanceProfileAssociationsRequest
type DescribeIamInstanceProfileAssociationsInput struct {
	_ struct{} `type:"structure"`

	// The IAM instance profile associations.
	AssociationIds []string `locationName:"AssociationId" locationNameList:"AssociationId" type:"list"`

	// The filters.
	//
	//    * instance-id - The ID of the instance.
	//
	//    * state - The state of the association (associating | associated | disassociating
	//    | disassociated).
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value.
	MaxResults *int64 `min:"5" type:"integer"`

	// The token to request the next page of results.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeIamInstanceProfileAssociationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeIamInstanceProfileAssociationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeIamInstanceProfileAssociationsInput"}
	if s.MaxResults != nil && *s.MaxResults < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 5))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeIamInstanceProfileAssociationsResult
type DescribeIamInstanceProfileAssociationsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the IAM instance profile associations.
	IamInstanceProfileAssociations []IamInstanceProfileAssociation `locationName:"iamInstanceProfileAssociationSet" locationNameList:"item" type:"list"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeIamInstanceProfileAssociationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeIamInstanceProfileAssociations = "DescribeIamInstanceProfileAssociations"

// DescribeIamInstanceProfileAssociationsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes your IAM instance profile associations.
//
//    // Example sending a request using DescribeIamInstanceProfileAssociationsRequest.
//    req := client.DescribeIamInstanceProfileAssociationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeIamInstanceProfileAssociations
func (c *Client) DescribeIamInstanceProfileAssociationsRequest(input *DescribeIamInstanceProfileAssociationsInput) DescribeIamInstanceProfileAssociationsRequest {
	op := &aws.Operation{
		Name:       opDescribeIamInstanceProfileAssociations,
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
		input = &DescribeIamInstanceProfileAssociationsInput{}
	}

	req := c.newRequest(op, input, &DescribeIamInstanceProfileAssociationsOutput{})
	return DescribeIamInstanceProfileAssociationsRequest{Request: req, Input: input, Copy: c.DescribeIamInstanceProfileAssociationsRequest}
}

// DescribeIamInstanceProfileAssociationsRequest is the request type for the
// DescribeIamInstanceProfileAssociations API operation.
type DescribeIamInstanceProfileAssociationsRequest struct {
	*aws.Request
	Input *DescribeIamInstanceProfileAssociationsInput
	Copy  func(*DescribeIamInstanceProfileAssociationsInput) DescribeIamInstanceProfileAssociationsRequest
}

// Send marshals and sends the DescribeIamInstanceProfileAssociations API request.
func (r DescribeIamInstanceProfileAssociationsRequest) Send(ctx context.Context) (*DescribeIamInstanceProfileAssociationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeIamInstanceProfileAssociationsResponse{
		DescribeIamInstanceProfileAssociationsOutput: r.Request.Data.(*DescribeIamInstanceProfileAssociationsOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeIamInstanceProfileAssociationsRequestPaginator returns a paginator for DescribeIamInstanceProfileAssociations.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeIamInstanceProfileAssociationsRequest(input)
//   p := ec2.NewDescribeIamInstanceProfileAssociationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeIamInstanceProfileAssociationsPaginator(req DescribeIamInstanceProfileAssociationsRequest) DescribeIamInstanceProfileAssociationsPaginator {
	return DescribeIamInstanceProfileAssociationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeIamInstanceProfileAssociationsInput
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

// DescribeIamInstanceProfileAssociationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeIamInstanceProfileAssociationsPaginator struct {
	aws.Pager
}

func (p *DescribeIamInstanceProfileAssociationsPaginator) CurrentPage() *DescribeIamInstanceProfileAssociationsOutput {
	return p.Pager.CurrentPage().(*DescribeIamInstanceProfileAssociationsOutput)
}

// DescribeIamInstanceProfileAssociationsResponse is the response type for the
// DescribeIamInstanceProfileAssociations API operation.
type DescribeIamInstanceProfileAssociationsResponse struct {
	*DescribeIamInstanceProfileAssociationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeIamInstanceProfileAssociations request.
func (r *DescribeIamInstanceProfileAssociationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
