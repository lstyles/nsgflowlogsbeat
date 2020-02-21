// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/SetAlarmStateInput
type SetAlarmStateInput struct {
	_ struct{} `type:"structure"`

	// The name for the alarm. This name must be unique within the AWS account.
	// The maximum length is 255 characters.
	//
	// AlarmName is a required field
	AlarmName *string `min:"1" type:"string" required:"true"`

	// The reason that this alarm is set to this specific state, in text format.
	//
	// StateReason is a required field
	StateReason *string `type:"string" required:"true"`

	// The reason that this alarm is set to this specific state, in JSON format.
	StateReasonData *string `type:"string"`

	// The value of the state.
	//
	// StateValue is a required field
	StateValue StateValue `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s SetAlarmStateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetAlarmStateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SetAlarmStateInput"}

	if s.AlarmName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AlarmName"))
	}
	if s.AlarmName != nil && len(*s.AlarmName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AlarmName", 1))
	}

	if s.StateReason == nil {
		invalidParams.Add(aws.NewErrParamRequired("StateReason"))
	}
	if len(s.StateValue) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("StateValue"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/SetAlarmStateOutput
type SetAlarmStateOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SetAlarmStateOutput) String() string {
	return awsutil.Prettify(s)
}

const opSetAlarmState = "SetAlarmState"

// SetAlarmStateRequest returns a request value for making API operation for
// Amazon CloudWatch.
//
// Temporarily sets the state of an alarm for testing purposes. When the updated
// state differs from the previous value, the action configured for the appropriate
// state is invoked. For example, if your alarm is configured to send an Amazon
// SNS message when an alarm is triggered, temporarily changing the alarm state
// to ALARM sends an SNS message. The alarm returns to its actual state (often
// within seconds). Because the alarm state change happens quickly, it is typically
// only visible in the alarm's History tab in the Amazon CloudWatch console
// or through DescribeAlarmHistory.
//
//    // Example sending a request using SetAlarmStateRequest.
//    req := client.SetAlarmStateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/SetAlarmState
func (c *Client) SetAlarmStateRequest(input *SetAlarmStateInput) SetAlarmStateRequest {
	op := &aws.Operation{
		Name:       opSetAlarmState,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SetAlarmStateInput{}
	}

	req := c.newRequest(op, input, &SetAlarmStateOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return SetAlarmStateRequest{Request: req, Input: input, Copy: c.SetAlarmStateRequest}
}

// SetAlarmStateRequest is the request type for the
// SetAlarmState API operation.
type SetAlarmStateRequest struct {
	*aws.Request
	Input *SetAlarmStateInput
	Copy  func(*SetAlarmStateInput) SetAlarmStateRequest
}

// Send marshals and sends the SetAlarmState API request.
func (r SetAlarmStateRequest) Send(ctx context.Context) (*SetAlarmStateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetAlarmStateResponse{
		SetAlarmStateOutput: r.Request.Data.(*SetAlarmStateOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetAlarmStateResponse is the response type for the
// SetAlarmState API operation.
type SetAlarmStateResponse struct {
	*SetAlarmStateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetAlarmState request.
func (r *SetAlarmStateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
