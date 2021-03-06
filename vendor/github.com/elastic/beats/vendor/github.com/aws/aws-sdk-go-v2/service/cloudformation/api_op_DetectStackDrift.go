// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/DetectStackDriftInput
type DetectStackDriftInput struct {
	_ struct{} `type:"structure"`

	// The logical names of any resources you want to use as filters.
	LogicalResourceIds []string `min:"1" type:"list"`

	// The name of the stack for which you want to detect drift.
	//
	// StackName is a required field
	StackName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DetectStackDriftInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DetectStackDriftInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DetectStackDriftInput"}
	if s.LogicalResourceIds != nil && len(s.LogicalResourceIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LogicalResourceIds", 1))
	}

	if s.StackName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackName"))
	}
	if s.StackName != nil && len(*s.StackName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StackName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/DetectStackDriftOutput
type DetectStackDriftOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the drift detection results of this operation.
	//
	// AWS CloudFormation generates new results, with a new drift detection ID,
	// each time this operation is run. However, the number of drift results AWS
	// CloudFormation retains for any given stack, and for how long, may vary.
	//
	// StackDriftDetectionId is a required field
	StackDriftDetectionId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DetectStackDriftOutput) String() string {
	return awsutil.Prettify(s)
}

const opDetectStackDrift = "DetectStackDrift"

// DetectStackDriftRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Detects whether a stack's actual configuration differs, or has drifted, from
// it's expected configuration, as defined in the stack template and any values
// specified as template parameters. For each resource in the stack that supports
// drift detection, AWS CloudFormation compares the actual configuration of
// the resource with its expected template configuration. Only resource properties
// explicitly defined in the stack template are checked for drift. A stack is
// considered to have drifted if one or more of its resources differ from their
// expected template configurations. For more information, see Detecting Unregulated
// Configuration Changes to Stacks and Resources (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-drift.html).
//
// Use DetectStackDrift to detect drift on all supported resources for a given
// stack, or DetectStackResourceDrift to detect drift on individual resources.
//
// For a list of stack resources that currently support drift detection, see
// Resources that Support Drift Detection (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-drift-resource-list.html).
//
// DetectStackDrift can take up to several minutes, depending on the number
// of resources contained within the stack. Use DescribeStackDriftDetectionStatus
// to monitor the progress of a detect stack drift operation. Once the drift
// detection operation has completed, use DescribeStackResourceDrifts to return
// drift information about the stack and its resources.
//
// When detecting drift on a stack, AWS CloudFormation does not detect drift
// on any nested stacks belonging to that stack. Perform DetectStackDrift directly
// on the nested stack itself.
//
//    // Example sending a request using DetectStackDriftRequest.
//    req := client.DetectStackDriftRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/DetectStackDrift
func (c *Client) DetectStackDriftRequest(input *DetectStackDriftInput) DetectStackDriftRequest {
	op := &aws.Operation{
		Name:       opDetectStackDrift,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DetectStackDriftInput{}
	}

	req := c.newRequest(op, input, &DetectStackDriftOutput{})
	return DetectStackDriftRequest{Request: req, Input: input, Copy: c.DetectStackDriftRequest}
}

// DetectStackDriftRequest is the request type for the
// DetectStackDrift API operation.
type DetectStackDriftRequest struct {
	*aws.Request
	Input *DetectStackDriftInput
	Copy  func(*DetectStackDriftInput) DetectStackDriftRequest
}

// Send marshals and sends the DetectStackDrift API request.
func (r DetectStackDriftRequest) Send(ctx context.Context) (*DetectStackDriftResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DetectStackDriftResponse{
		DetectStackDriftOutput: r.Request.Data.(*DetectStackDriftOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DetectStackDriftResponse is the response type for the
// DetectStackDrift API operation.
type DetectStackDriftResponse struct {
	*DetectStackDriftOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DetectStackDrift request.
func (r *DetectStackDriftResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
