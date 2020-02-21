// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DescribeDBClusterSnapshotsMessage
type DescribeDBClusterSnapshotsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the DB cluster to retrieve the list of DB cluster snapshots for.
	// This parameter can't be used in conjunction with the DBClusterSnapshotIdentifier
	// parameter. This parameter is not case-sensitive.
	//
	// Constraints:
	//
	//    * If supplied, must match the identifier of an existing DBCluster.
	DBClusterIdentifier *string `type:"string"`

	// A specific DB cluster snapshot identifier to describe. This parameter can't
	// be used in conjunction with the DBClusterIdentifier parameter. This value
	// is stored as a lowercase string.
	//
	// Constraints:
	//
	//    * If supplied, must match the identifier of an existing DBClusterSnapshot.
	//
	//    * If this identifier is for an automated snapshot, the SnapshotType parameter
	//    must also be specified.
	DBClusterSnapshotIdentifier *string `type:"string"`

	// This parameter is not currently supported.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// True to include manual DB cluster snapshots that are public and can be copied
	// or restored by any AWS account, and otherwise false. The default is false.
	// The default is false.
	//
	// You can share a manual DB cluster snapshot as public by using the ModifyDBClusterSnapshotAttribute
	// API action.
	IncludePublic *bool `type:"boolean"`

	// True to include shared manual DB cluster snapshots from other AWS accounts
	// that this AWS account has been given permission to copy or restore, and otherwise
	// false. The default is false.
	//
	// You can give an AWS account permission to restore a manual DB cluster snapshot
	// from another AWS account by the ModifyDBClusterSnapshotAttribute API action.
	IncludeShared *bool `type:"boolean"`

	// An optional pagination token provided by a previous DescribeDBClusterSnapshots
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

	// The type of DB cluster snapshots to be returned. You can specify one of the
	// following values:
	//
	//    * automated - Return all DB cluster snapshots that have been automatically
	//    taken by Amazon RDS for my AWS account.
	//
	//    * manual - Return all DB cluster snapshots that have been taken by my
	//    AWS account.
	//
	//    * shared - Return all manual DB cluster snapshots that have been shared
	//    to my AWS account.
	//
	//    * public - Return all DB cluster snapshots that have been marked as public.
	//
	// If you don't specify a SnapshotType value, then both automated and manual
	// DB cluster snapshots are returned. You can include shared DB cluster snapshots
	// with these results by setting the IncludeShared parameter to true. You can
	// include public DB cluster snapshots with these results by setting the IncludePublic
	// parameter to true.
	//
	// The IncludeShared and IncludePublic parameters don't apply for SnapshotType
	// values of manual or automated. The IncludePublic parameter doesn't apply
	// when SnapshotType is set to shared. The IncludeShared parameter doesn't apply
	// when SnapshotType is set to public.
	SnapshotType *string `type:"string"`
}

// String returns the string representation
func (s DescribeDBClusterSnapshotsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDBClusterSnapshotsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDBClusterSnapshotsInput"}
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

// Provides a list of DB cluster snapshots for the user as the result of a call
// to the DescribeDBClusterSnapshots action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DBClusterSnapshotMessage
type DescribeDBClusterSnapshotsOutput struct {
	_ struct{} `type:"structure"`

	// Provides a list of DB cluster snapshots for the user.
	DBClusterSnapshots []DBClusterSnapshot `locationNameList:"DBClusterSnapshot" type:"list"`

	// An optional pagination token provided by a previous DescribeDBClusterSnapshots
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeDBClusterSnapshotsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeDBClusterSnapshots = "DescribeDBClusterSnapshots"

// DescribeDBClusterSnapshotsRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Returns information about DB cluster snapshots. This API action supports
// pagination.
//
// For more information on Amazon Aurora, see What Is Amazon Aurora? (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide.
//
// This action only applies to Aurora DB clusters.
//
//    // Example sending a request using DescribeDBClusterSnapshotsRequest.
//    req := client.DescribeDBClusterSnapshotsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DescribeDBClusterSnapshots
func (c *Client) DescribeDBClusterSnapshotsRequest(input *DescribeDBClusterSnapshotsInput) DescribeDBClusterSnapshotsRequest {
	op := &aws.Operation{
		Name:       opDescribeDBClusterSnapshots,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDBClusterSnapshotsInput{}
	}

	req := c.newRequest(op, input, &DescribeDBClusterSnapshotsOutput{})
	return DescribeDBClusterSnapshotsRequest{Request: req, Input: input, Copy: c.DescribeDBClusterSnapshotsRequest}
}

// DescribeDBClusterSnapshotsRequest is the request type for the
// DescribeDBClusterSnapshots API operation.
type DescribeDBClusterSnapshotsRequest struct {
	*aws.Request
	Input *DescribeDBClusterSnapshotsInput
	Copy  func(*DescribeDBClusterSnapshotsInput) DescribeDBClusterSnapshotsRequest
}

// Send marshals and sends the DescribeDBClusterSnapshots API request.
func (r DescribeDBClusterSnapshotsRequest) Send(ctx context.Context) (*DescribeDBClusterSnapshotsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDBClusterSnapshotsResponse{
		DescribeDBClusterSnapshotsOutput: r.Request.Data.(*DescribeDBClusterSnapshotsOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDBClusterSnapshotsResponse is the response type for the
// DescribeDBClusterSnapshots API operation.
type DescribeDBClusterSnapshotsResponse struct {
	*DescribeDBClusterSnapshotsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDBClusterSnapshots request.
func (r *DescribeDBClusterSnapshotsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
