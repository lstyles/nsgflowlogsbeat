// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DescribeGlobalClustersMessage
type DescribeGlobalClustersInput struct {
	_ struct{} `type:"structure"`

	// A filter that specifies one or more global DB clusters to describe.
	//
	// Supported filters:
	//
	//    * db-cluster-id - Accepts DB cluster identifiers and DB cluster Amazon
	//    Resource Names (ARNs). The results list will only include information
	//    about the DB clusters identified by these ARNs.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// The user-supplied DB cluster identifier. If this parameter is specified,
	// information from only the specific DB cluster is returned. This parameter
	// isn't case-sensitive.
	//
	// Constraints:
	//
	//    * If supplied, must match an existing DBClusterIdentifier.
	GlobalClusterIdentifier *string `type:"string"`

	// An optional pagination token provided by a previous DescribeGlobalClusters
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeGlobalClustersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeGlobalClustersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeGlobalClustersInput"}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/GlobalClustersMessage
type DescribeGlobalClustersOutput struct {
	_ struct{} `type:"structure"`

	// The list of global clusters returned by this request.
	GlobalClusters []GlobalCluster `locationNameList:"GlobalClusterMember" type:"list"`

	// An optional pagination token provided by a previous DescribeGlobalClusters
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeGlobalClustersOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeGlobalClusters = "DescribeGlobalClusters"

// DescribeGlobalClustersRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Returns information about Aurora global database clusters. This API supports
// pagination.
//
// For more information on Amazon Aurora, see What Is Amazon Aurora? (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide.
//
// This action only applies to Aurora DB clusters.
//
//    // Example sending a request using DescribeGlobalClustersRequest.
//    req := client.DescribeGlobalClustersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DescribeGlobalClusters
func (c *Client) DescribeGlobalClustersRequest(input *DescribeGlobalClustersInput) DescribeGlobalClustersRequest {
	op := &aws.Operation{
		Name:       opDescribeGlobalClusters,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeGlobalClustersInput{}
	}

	req := c.newRequest(op, input, &DescribeGlobalClustersOutput{})
	return DescribeGlobalClustersRequest{Request: req, Input: input, Copy: c.DescribeGlobalClustersRequest}
}

// DescribeGlobalClustersRequest is the request type for the
// DescribeGlobalClusters API operation.
type DescribeGlobalClustersRequest struct {
	*aws.Request
	Input *DescribeGlobalClustersInput
	Copy  func(*DescribeGlobalClustersInput) DescribeGlobalClustersRequest
}

// Send marshals and sends the DescribeGlobalClusters API request.
func (r DescribeGlobalClustersRequest) Send(ctx context.Context) (*DescribeGlobalClustersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeGlobalClustersResponse{
		DescribeGlobalClustersOutput: r.Request.Data.(*DescribeGlobalClustersOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeGlobalClustersRequestPaginator returns a paginator for DescribeGlobalClusters.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeGlobalClustersRequest(input)
//   p := rds.NewDescribeGlobalClustersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeGlobalClustersPaginator(req DescribeGlobalClustersRequest) DescribeGlobalClustersPaginator {
	return DescribeGlobalClustersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeGlobalClustersInput
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

// DescribeGlobalClustersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeGlobalClustersPaginator struct {
	aws.Pager
}

func (p *DescribeGlobalClustersPaginator) CurrentPage() *DescribeGlobalClustersOutput {
	return p.Pager.CurrentPage().(*DescribeGlobalClustersOutput)
}

// DescribeGlobalClustersResponse is the response type for the
// DescribeGlobalClusters API operation.
type DescribeGlobalClustersResponse struct {
	*DescribeGlobalClustersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeGlobalClusters request.
func (r *DescribeGlobalClustersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
