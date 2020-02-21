// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribePlacementGroupsRequest
type DescribePlacementGroupsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The filters.
	//
	//    * group-name - The name of the placement group.
	//
	//    * state - The state of the placement group (pending | available | deleting
	//    | deleted).
	//
	//    * strategy - The strategy of the placement group (cluster | spread | partition).
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The names of the placement groups.
	//
	// Default: Describes all your placement groups, or only those otherwise specified.
	GroupNames []string `locationName:"groupName" type:"list"`
}

// String returns the string representation
func (s DescribePlacementGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribePlacementGroupsResult
type DescribePlacementGroupsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the placement groups.
	PlacementGroups []PlacementGroup `locationName:"placementGroupSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribePlacementGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribePlacementGroups = "DescribePlacementGroups"

// DescribePlacementGroupsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the specified placement groups or all of your placement groups.
// For more information, see Placement Groups (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using DescribePlacementGroupsRequest.
//    req := client.DescribePlacementGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribePlacementGroups
func (c *Client) DescribePlacementGroupsRequest(input *DescribePlacementGroupsInput) DescribePlacementGroupsRequest {
	op := &aws.Operation{
		Name:       opDescribePlacementGroups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribePlacementGroupsInput{}
	}

	req := c.newRequest(op, input, &DescribePlacementGroupsOutput{})
	return DescribePlacementGroupsRequest{Request: req, Input: input, Copy: c.DescribePlacementGroupsRequest}
}

// DescribePlacementGroupsRequest is the request type for the
// DescribePlacementGroups API operation.
type DescribePlacementGroupsRequest struct {
	*aws.Request
	Input *DescribePlacementGroupsInput
	Copy  func(*DescribePlacementGroupsInput) DescribePlacementGroupsRequest
}

// Send marshals and sends the DescribePlacementGroups API request.
func (r DescribePlacementGroupsRequest) Send(ctx context.Context) (*DescribePlacementGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribePlacementGroupsResponse{
		DescribePlacementGroupsOutput: r.Request.Data.(*DescribePlacementGroupsOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribePlacementGroupsResponse is the response type for the
// DescribePlacementGroups API operation.
type DescribePlacementGroupsResponse struct {
	*DescribePlacementGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribePlacementGroups request.
func (r *DescribePlacementGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
