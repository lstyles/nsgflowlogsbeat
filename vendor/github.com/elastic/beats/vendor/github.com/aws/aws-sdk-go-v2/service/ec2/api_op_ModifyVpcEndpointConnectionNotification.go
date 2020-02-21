// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyVpcEndpointConnectionNotificationRequest
type ModifyVpcEndpointConnectionNotificationInput struct {
	_ struct{} `type:"structure"`

	// One or more events for the endpoint. Valid values are Accept, Connect, Delete,
	// and Reject.
	ConnectionEvents []string `locationNameList:"item" type:"list"`

	// The ARN for the SNS topic for the notification.
	ConnectionNotificationArn *string `type:"string"`

	// The ID of the notification.
	//
	// ConnectionNotificationId is a required field
	ConnectionNotificationId *string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`
}

// String returns the string representation
func (s ModifyVpcEndpointConnectionNotificationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyVpcEndpointConnectionNotificationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyVpcEndpointConnectionNotificationInput"}

	if s.ConnectionNotificationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConnectionNotificationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyVpcEndpointConnectionNotificationResult
type ModifyVpcEndpointConnectionNotificationOutput struct {
	_ struct{} `type:"structure"`

	// Returns true if the request succeeds; otherwise, it returns an error.
	ReturnValue *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s ModifyVpcEndpointConnectionNotificationOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyVpcEndpointConnectionNotification = "ModifyVpcEndpointConnectionNotification"

// ModifyVpcEndpointConnectionNotificationRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Modifies a connection notification for VPC endpoint or VPC endpoint service.
// You can change the SNS topic for the notification, or the events for which
// to be notified.
//
//    // Example sending a request using ModifyVpcEndpointConnectionNotificationRequest.
//    req := client.ModifyVpcEndpointConnectionNotificationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyVpcEndpointConnectionNotification
func (c *Client) ModifyVpcEndpointConnectionNotificationRequest(input *ModifyVpcEndpointConnectionNotificationInput) ModifyVpcEndpointConnectionNotificationRequest {
	op := &aws.Operation{
		Name:       opModifyVpcEndpointConnectionNotification,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyVpcEndpointConnectionNotificationInput{}
	}

	req := c.newRequest(op, input, &ModifyVpcEndpointConnectionNotificationOutput{})
	return ModifyVpcEndpointConnectionNotificationRequest{Request: req, Input: input, Copy: c.ModifyVpcEndpointConnectionNotificationRequest}
}

// ModifyVpcEndpointConnectionNotificationRequest is the request type for the
// ModifyVpcEndpointConnectionNotification API operation.
type ModifyVpcEndpointConnectionNotificationRequest struct {
	*aws.Request
	Input *ModifyVpcEndpointConnectionNotificationInput
	Copy  func(*ModifyVpcEndpointConnectionNotificationInput) ModifyVpcEndpointConnectionNotificationRequest
}

// Send marshals and sends the ModifyVpcEndpointConnectionNotification API request.
func (r ModifyVpcEndpointConnectionNotificationRequest) Send(ctx context.Context) (*ModifyVpcEndpointConnectionNotificationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyVpcEndpointConnectionNotificationResponse{
		ModifyVpcEndpointConnectionNotificationOutput: r.Request.Data.(*ModifyVpcEndpointConnectionNotificationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyVpcEndpointConnectionNotificationResponse is the response type for the
// ModifyVpcEndpointConnectionNotification API operation.
type ModifyVpcEndpointConnectionNotificationResponse struct {
	*ModifyVpcEndpointConnectionNotificationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyVpcEndpointConnectionNotification request.
func (r *ModifyVpcEndpointConnectionNotificationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
