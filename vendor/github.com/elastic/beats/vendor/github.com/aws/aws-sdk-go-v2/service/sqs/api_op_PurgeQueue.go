// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/PurgeQueueRequest
type PurgeQueueInput struct {
	_ struct{} `type:"structure"`

	// The URL of the queue from which the PurgeQueue action deletes messages.
	//
	// Queue URLs and names are case-sensitive.
	//
	// QueueUrl is a required field
	QueueUrl *string `type:"string" required:"true"`
}

// String returns the string representation
func (s PurgeQueueInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PurgeQueueInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PurgeQueueInput"}

	if s.QueueUrl == nil {
		invalidParams.Add(aws.NewErrParamRequired("QueueUrl"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/PurgeQueueOutput
type PurgeQueueOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PurgeQueueOutput) String() string {
	return awsutil.Prettify(s)
}

const opPurgeQueue = "PurgeQueue"

// PurgeQueueRequest returns a request value for making API operation for
// Amazon Simple Queue Service.
//
// Deletes the messages in a queue specified by the QueueURL parameter.
//
// When you use the PurgeQueue action, you can't retrieve any messages deleted
// from a queue.
//
// The message deletion process takes up to 60 seconds. We recommend waiting
// for 60 seconds regardless of your queue's size.
//
// Messages sent to the queue before you call PurgeQueue might be received but
// are deleted within the next minute.
//
// Messages sent to the queue after you call PurgeQueue might be deleted while
// the queue is being purged.
//
//    // Example sending a request using PurgeQueueRequest.
//    req := client.PurgeQueueRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/PurgeQueue
func (c *Client) PurgeQueueRequest(input *PurgeQueueInput) PurgeQueueRequest {
	op := &aws.Operation{
		Name:       opPurgeQueue,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PurgeQueueInput{}
	}

	req := c.newRequest(op, input, &PurgeQueueOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return PurgeQueueRequest{Request: req, Input: input, Copy: c.PurgeQueueRequest}
}

// PurgeQueueRequest is the request type for the
// PurgeQueue API operation.
type PurgeQueueRequest struct {
	*aws.Request
	Input *PurgeQueueInput
	Copy  func(*PurgeQueueInput) PurgeQueueRequest
}

// Send marshals and sends the PurgeQueue API request.
func (r PurgeQueueRequest) Send(ctx context.Context) (*PurgeQueueResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PurgeQueueResponse{
		PurgeQueueOutput: r.Request.Data.(*PurgeQueueOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PurgeQueueResponse is the response type for the
// PurgeQueue API operation.
type PurgeQueueResponse struct {
	*PurgeQueueOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PurgeQueue request.
func (r *PurgeQueueResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
