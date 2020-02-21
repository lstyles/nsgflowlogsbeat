// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for DescribeNetworkInterfacePermissions.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeNetworkInterfacePermissionsRequest
type DescribeNetworkInterfacePermissionsInput struct {
	_ struct{} `type:"structure"`

	// One or more filters.
	//
	//    * network-interface-permission.network-interface-permission-id - The ID
	//    of the permission.
	//
	//    * network-interface-permission.network-interface-id - The ID of the network
	//    interface.
	//
	//    * network-interface-permission.aws-account-id - The AWS account ID.
	//
	//    * network-interface-permission.aws-service - The AWS service.
	//
	//    * network-interface-permission.permission - The type of permission (INSTANCE-ATTACH
	//    | EIP-ASSOCIATE).
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value. If
	// this parameter is not specified, up to 50 results are returned by default.
	MaxResults *int64 `min:"5" type:"integer"`

	// One or more network interface permission IDs.
	NetworkInterfacePermissionIds []string `locationName:"NetworkInterfacePermissionId" type:"list"`

	// The token to request the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeNetworkInterfacePermissionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeNetworkInterfacePermissionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeNetworkInterfacePermissionsInput"}
	if s.MaxResults != nil && *s.MaxResults < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 5))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output for DescribeNetworkInterfacePermissions.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeNetworkInterfacePermissionsResult
type DescribeNetworkInterfacePermissionsOutput struct {
	_ struct{} `type:"structure"`

	// The network interface permissions.
	NetworkInterfacePermissions []NetworkInterfacePermission `locationName:"networkInterfacePermissions" locationNameList:"item" type:"list"`

	// The token to use to retrieve the next page of results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeNetworkInterfacePermissionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeNetworkInterfacePermissions = "DescribeNetworkInterfacePermissions"

// DescribeNetworkInterfacePermissionsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the permissions for your network interfaces.
//
//    // Example sending a request using DescribeNetworkInterfacePermissionsRequest.
//    req := client.DescribeNetworkInterfacePermissionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeNetworkInterfacePermissions
func (c *Client) DescribeNetworkInterfacePermissionsRequest(input *DescribeNetworkInterfacePermissionsInput) DescribeNetworkInterfacePermissionsRequest {
	op := &aws.Operation{
		Name:       opDescribeNetworkInterfacePermissions,
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
		input = &DescribeNetworkInterfacePermissionsInput{}
	}

	req := c.newRequest(op, input, &DescribeNetworkInterfacePermissionsOutput{})
	return DescribeNetworkInterfacePermissionsRequest{Request: req, Input: input, Copy: c.DescribeNetworkInterfacePermissionsRequest}
}

// DescribeNetworkInterfacePermissionsRequest is the request type for the
// DescribeNetworkInterfacePermissions API operation.
type DescribeNetworkInterfacePermissionsRequest struct {
	*aws.Request
	Input *DescribeNetworkInterfacePermissionsInput
	Copy  func(*DescribeNetworkInterfacePermissionsInput) DescribeNetworkInterfacePermissionsRequest
}

// Send marshals and sends the DescribeNetworkInterfacePermissions API request.
func (r DescribeNetworkInterfacePermissionsRequest) Send(ctx context.Context) (*DescribeNetworkInterfacePermissionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeNetworkInterfacePermissionsResponse{
		DescribeNetworkInterfacePermissionsOutput: r.Request.Data.(*DescribeNetworkInterfacePermissionsOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeNetworkInterfacePermissionsRequestPaginator returns a paginator for DescribeNetworkInterfacePermissions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeNetworkInterfacePermissionsRequest(input)
//   p := ec2.NewDescribeNetworkInterfacePermissionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeNetworkInterfacePermissionsPaginator(req DescribeNetworkInterfacePermissionsRequest) DescribeNetworkInterfacePermissionsPaginator {
	return DescribeNetworkInterfacePermissionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeNetworkInterfacePermissionsInput
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

// DescribeNetworkInterfacePermissionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeNetworkInterfacePermissionsPaginator struct {
	aws.Pager
}

func (p *DescribeNetworkInterfacePermissionsPaginator) CurrentPage() *DescribeNetworkInterfacePermissionsOutput {
	return p.Pager.CurrentPage().(*DescribeNetworkInterfacePermissionsOutput)
}

// DescribeNetworkInterfacePermissionsResponse is the response type for the
// DescribeNetworkInterfacePermissions API operation.
type DescribeNetworkInterfacePermissionsResponse struct {
	*DescribeNetworkInterfacePermissionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeNetworkInterfacePermissions request.
func (r *DescribeNetworkInterfacePermissionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
