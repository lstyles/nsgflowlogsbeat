// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyInstancePlacementRequest
type ModifyInstancePlacementInput struct {
	_ struct{} `type:"structure"`

	// The affinity setting for the instance.
	Affinity Affinity `locationName:"affinity" type:"string" enum:"true"`

	// The name of the placement group in which to place the instance. For spread
	// placement groups, the instance must have a tenancy of default. For cluster
	// and partition placement groups, the instance must have a tenancy of default
	// or dedicated.
	//
	// To remove an instance from a placement group, specify an empty string ("").
	GroupName *string `type:"string"`

	// The ID of the Dedicated Host with which to associate the instance.
	HostId *string `locationName:"hostId" type:"string"`

	// The ID of the instance that you are modifying.
	//
	// InstanceId is a required field
	InstanceId *string `locationName:"instanceId" type:"string" required:"true"`

	// Reserved for future use.
	PartitionNumber *int64 `type:"integer"`

	// The tenancy for the instance.
	Tenancy HostTenancy `locationName:"tenancy" type:"string" enum:"true"`
}

// String returns the string representation
func (s ModifyInstancePlacementInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyInstancePlacementInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyInstancePlacementInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyInstancePlacementResult
type ModifyInstancePlacementOutput struct {
	_ struct{} `type:"structure"`

	// Is true if the request succeeds, and an error otherwise.
	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s ModifyInstancePlacementOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyInstancePlacement = "ModifyInstancePlacement"

// ModifyInstancePlacementRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Modifies the placement attributes for a specified instance. You can do the
// following:
//
//    * Modify the affinity between an instance and a Dedicated Host (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-hosts-overview.html).
//    When affinity is set to host and the instance is not associated with a
//    specific Dedicated Host, the next time the instance is launched, it is
//    automatically associated with the host on which it lands. If the instance
//    is restarted or rebooted, this relationship persists.
//
//    * Change the Dedicated Host with which an instance is associated.
//
//    * Change the instance tenancy of an instance from host to dedicated, or
//    from dedicated to host.
//
//    * Move an instance to or from a placement group (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html).
//
// At least one attribute for affinity, host ID, tenancy, or placement group
// name must be specified in the request. Affinity and tenancy can be modified
// in the same request.
//
// To modify the host ID, tenancy, placement group, or partition for an instance,
// the instance must be in the stopped state.
//
//    // Example sending a request using ModifyInstancePlacementRequest.
//    req := client.ModifyInstancePlacementRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyInstancePlacement
func (c *Client) ModifyInstancePlacementRequest(input *ModifyInstancePlacementInput) ModifyInstancePlacementRequest {
	op := &aws.Operation{
		Name:       opModifyInstancePlacement,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyInstancePlacementInput{}
	}

	req := c.newRequest(op, input, &ModifyInstancePlacementOutput{})
	return ModifyInstancePlacementRequest{Request: req, Input: input, Copy: c.ModifyInstancePlacementRequest}
}

// ModifyInstancePlacementRequest is the request type for the
// ModifyInstancePlacement API operation.
type ModifyInstancePlacementRequest struct {
	*aws.Request
	Input *ModifyInstancePlacementInput
	Copy  func(*ModifyInstancePlacementInput) ModifyInstancePlacementRequest
}

// Send marshals and sends the ModifyInstancePlacement API request.
func (r ModifyInstancePlacementRequest) Send(ctx context.Context) (*ModifyInstancePlacementResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyInstancePlacementResponse{
		ModifyInstancePlacementOutput: r.Request.Data.(*ModifyInstancePlacementOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyInstancePlacementResponse is the response type for the
// ModifyInstancePlacement API operation.
type ModifyInstancePlacementResponse struct {
	*ModifyInstancePlacementOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyInstancePlacement request.
func (r *ModifyInstancePlacementResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
