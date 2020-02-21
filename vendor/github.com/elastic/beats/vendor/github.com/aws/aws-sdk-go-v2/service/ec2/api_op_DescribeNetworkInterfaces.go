// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for DescribeNetworkInterfaces.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeNetworkInterfacesRequest
type DescribeNetworkInterfacesInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// One or more filters.
	//
	//    * addresses.private-ip-address - The private IPv4 addresses associated
	//    with the network interface.
	//
	//    * addresses.primary - Whether the private IPv4 address is the primary
	//    IP address associated with the network interface.
	//
	//    * addresses.association.public-ip - The association ID returned when the
	//    network interface was associated with the Elastic IP address (IPv4).
	//
	//    * addresses.association.owner-id - The owner ID of the addresses associated
	//    with the network interface.
	//
	//    * association.association-id - The association ID returned when the network
	//    interface was associated with an IPv4 address.
	//
	//    * association.allocation-id - The allocation ID returned when you allocated
	//    the Elastic IP address (IPv4) for your network interface.
	//
	//    * association.ip-owner-id - The owner of the Elastic IP address (IPv4)
	//    associated with the network interface.
	//
	//    * association.public-ip - The address of the Elastic IP address (IPv4)
	//    bound to the network interface.
	//
	//    * association.public-dns-name - The public DNS name for the network interface
	//    (IPv4).
	//
	//    * attachment.attachment-id - The ID of the interface attachment.
	//
	//    * attachment.attach.time - The time that the network interface was attached
	//    to an instance.
	//
	//    * attachment.delete-on-termination - Indicates whether the attachment
	//    is deleted when an instance is terminated.
	//
	//    * attachment.device-index - The device index to which the network interface
	//    is attached.
	//
	//    * attachment.instance-id - The ID of the instance to which the network
	//    interface is attached.
	//
	//    * attachment.instance-owner-id - The owner ID of the instance to which
	//    the network interface is attached.
	//
	//    * attachment.nat-gateway-id - The ID of the NAT gateway to which the network
	//    interface is attached.
	//
	//    * attachment.status - The status of the attachment (attaching | attached
	//    | detaching | detached).
	//
	//    * availability-zone - The Availability Zone of the network interface.
	//
	//    * description - The description of the network interface.
	//
	//    * group-id - The ID of a security group associated with the network interface.
	//
	//    * group-name - The name of a security group associated with the network
	//    interface.
	//
	//    * ipv6-addresses.ipv6-address - An IPv6 address associated with the network
	//    interface.
	//
	//    * mac-address - The MAC address of the network interface.
	//
	//    * network-interface-id - The ID of the network interface.
	//
	//    * owner-id - The AWS account ID of the network interface owner.
	//
	//    * private-ip-address - The private IPv4 address or addresses of the network
	//    interface.
	//
	//    * private-dns-name - The private DNS name of the network interface (IPv4).
	//
	//    * requester-id - The ID of the entity that launched the instance on your
	//    behalf (for example, AWS Management Console, Auto Scaling, and so on).
	//
	//    * requester-managed - Indicates whether the network interface is being
	//    managed by an AWS service (for example, AWS Management Console, Auto Scaling,
	//    and so on).
	//
	//    * source-dest-check - Indicates whether the network interface performs
	//    source/destination checking. A value of true means checking is enabled,
	//    and false means checking is disabled. The value must be false for the
	//    network interface to perform network address translation (NAT) in your
	//    VPC.
	//
	//    * status - The status of the network interface. If the network interface
	//    is not attached to an instance, the status is available; if a network
	//    interface is attached to an instance the status is in-use.
	//
	//    * subnet-id - The ID of the subnet for the network interface.
	//
	//    * tag:<key> - The key/value combination of a tag assigned to the resource.
	//    Use the tag key in the filter name and the tag value as the filter value.
	//    For example, to find all resources that have a tag with the key Owner
	//    and the value TeamA, specify tag:Owner for the filter name and TeamA for
	//    the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. Use this filter
	//    to find all resources assigned a tag with a specific key, regardless of
	//    the tag value.
	//
	//    * vpc-id - The ID of the VPC for the network interface.
	Filters []Filter `locationName:"filter" locationNameList:"Filter" type:"list"`

	// The maximum number of items to return for this request. The request returns
	// a token that you can specify in a subsequent call to get the next set of
	// results.
	MaxResults *int64 `min:"5" type:"integer"`

	// One or more network interface IDs.
	//
	// Default: Describes all your network interfaces.
	NetworkInterfaceIds []string `locationName:"NetworkInterfaceId" locationNameList:"item" type:"list"`

	// The token to retrieve the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeNetworkInterfacesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeNetworkInterfacesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeNetworkInterfacesInput"}
	if s.MaxResults != nil && *s.MaxResults < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 5))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of DescribeNetworkInterfaces.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeNetworkInterfacesResult
type DescribeNetworkInterfacesOutput struct {
	_ struct{} `type:"structure"`

	// Information about one or more network interfaces.
	NetworkInterfaces []NetworkInterface `locationName:"networkInterfaceSet" locationNameList:"item" type:"list"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeNetworkInterfacesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeNetworkInterfaces = "DescribeNetworkInterfaces"

// DescribeNetworkInterfacesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes one or more of your network interfaces.
//
//    // Example sending a request using DescribeNetworkInterfacesRequest.
//    req := client.DescribeNetworkInterfacesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeNetworkInterfaces
func (c *Client) DescribeNetworkInterfacesRequest(input *DescribeNetworkInterfacesInput) DescribeNetworkInterfacesRequest {
	op := &aws.Operation{
		Name:       opDescribeNetworkInterfaces,
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
		input = &DescribeNetworkInterfacesInput{}
	}

	req := c.newRequest(op, input, &DescribeNetworkInterfacesOutput{})
	return DescribeNetworkInterfacesRequest{Request: req, Input: input, Copy: c.DescribeNetworkInterfacesRequest}
}

// DescribeNetworkInterfacesRequest is the request type for the
// DescribeNetworkInterfaces API operation.
type DescribeNetworkInterfacesRequest struct {
	*aws.Request
	Input *DescribeNetworkInterfacesInput
	Copy  func(*DescribeNetworkInterfacesInput) DescribeNetworkInterfacesRequest
}

// Send marshals and sends the DescribeNetworkInterfaces API request.
func (r DescribeNetworkInterfacesRequest) Send(ctx context.Context) (*DescribeNetworkInterfacesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeNetworkInterfacesResponse{
		DescribeNetworkInterfacesOutput: r.Request.Data.(*DescribeNetworkInterfacesOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeNetworkInterfacesRequestPaginator returns a paginator for DescribeNetworkInterfaces.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeNetworkInterfacesRequest(input)
//   p := ec2.NewDescribeNetworkInterfacesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeNetworkInterfacesPaginator(req DescribeNetworkInterfacesRequest) DescribeNetworkInterfacesPaginator {
	return DescribeNetworkInterfacesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeNetworkInterfacesInput
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

// DescribeNetworkInterfacesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeNetworkInterfacesPaginator struct {
	aws.Pager
}

func (p *DescribeNetworkInterfacesPaginator) CurrentPage() *DescribeNetworkInterfacesOutput {
	return p.Pager.CurrentPage().(*DescribeNetworkInterfacesOutput)
}

// DescribeNetworkInterfacesResponse is the response type for the
// DescribeNetworkInterfaces API operation.
type DescribeNetworkInterfacesResponse struct {
	*DescribeNetworkInterfacesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeNetworkInterfaces request.
func (r *DescribeNetworkInterfacesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
