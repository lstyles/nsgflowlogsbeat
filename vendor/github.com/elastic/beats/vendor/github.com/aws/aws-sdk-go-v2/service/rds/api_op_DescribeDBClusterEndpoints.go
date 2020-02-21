// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DescribeDBClusterEndpointsMessage
type DescribeDBClusterEndpointsInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the endpoint to describe. This parameter is stored as a
	// lowercase string.
	DBClusterEndpointIdentifier *string `type:"string"`

	// The DB cluster identifier of the DB cluster associated with the endpoint.
	// This parameter is stored as a lowercase string.
	DBClusterIdentifier *string `type:"string"`

	// A set of name-value pairs that define which endpoints to include in the output.
	// The filters are specified as name-value pairs, in the format Name=endpoint_type,Values=endpoint_type1,endpoint_type2,....
	// Name can be one of: db-cluster-endpoint-type, db-cluster-endpoint-custom-type,
	// db-cluster-endpoint-id, db-cluster-endpoint-status. Values for the db-cluster-endpoint-type
	// filter can be one or more of: reader, writer, custom. Values for the db-cluster-endpoint-custom-type
	// filter can be one or more of: reader, any. Values for the db-cluster-endpoint-status
	// filter can be one or more of: available, creating, deleting, modifying.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// An optional pagination token provided by a previous DescribeDBClusterEndpoints
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
func (s DescribeDBClusterEndpointsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDBClusterEndpointsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDBClusterEndpointsInput"}
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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DBClusterEndpointMessage
type DescribeDBClusterEndpointsOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details of the endpoints associated with the cluster and matching
	// any filter conditions.
	DBClusterEndpoints []DBClusterEndpoint `locationNameList:"DBClusterEndpointList" type:"list"`

	// An optional pagination token provided by a previous DescribeDBClusterEndpoints
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeDBClusterEndpointsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeDBClusterEndpoints = "DescribeDBClusterEndpoints"

// DescribeDBClusterEndpointsRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Returns information about endpoints for an Amazon Aurora DB cluster.
//
// This action only applies to Aurora DB clusters.
//
//    // Example sending a request using DescribeDBClusterEndpointsRequest.
//    req := client.DescribeDBClusterEndpointsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DescribeDBClusterEndpoints
func (c *Client) DescribeDBClusterEndpointsRequest(input *DescribeDBClusterEndpointsInput) DescribeDBClusterEndpointsRequest {
	op := &aws.Operation{
		Name:       opDescribeDBClusterEndpoints,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDBClusterEndpointsInput{}
	}

	req := c.newRequest(op, input, &DescribeDBClusterEndpointsOutput{})
	return DescribeDBClusterEndpointsRequest{Request: req, Input: input, Copy: c.DescribeDBClusterEndpointsRequest}
}

// DescribeDBClusterEndpointsRequest is the request type for the
// DescribeDBClusterEndpoints API operation.
type DescribeDBClusterEndpointsRequest struct {
	*aws.Request
	Input *DescribeDBClusterEndpointsInput
	Copy  func(*DescribeDBClusterEndpointsInput) DescribeDBClusterEndpointsRequest
}

// Send marshals and sends the DescribeDBClusterEndpoints API request.
func (r DescribeDBClusterEndpointsRequest) Send(ctx context.Context) (*DescribeDBClusterEndpointsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDBClusterEndpointsResponse{
		DescribeDBClusterEndpointsOutput: r.Request.Data.(*DescribeDBClusterEndpointsOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDBClusterEndpointsResponse is the response type for the
// DescribeDBClusterEndpoints API operation.
type DescribeDBClusterEndpointsResponse struct {
	*DescribeDBClusterEndpointsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDBClusterEndpoints request.
func (r *DescribeDBClusterEndpointsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
