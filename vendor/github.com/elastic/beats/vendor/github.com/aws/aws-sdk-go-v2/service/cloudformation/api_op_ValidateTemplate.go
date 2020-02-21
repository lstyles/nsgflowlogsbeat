// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The input for ValidateTemplate action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/ValidateTemplateInput
type ValidateTemplateInput struct {
	_ struct{} `type:"structure"`

	// Structure containing the template body with a minimum length of 1 byte and
	// a maximum length of 51,200 bytes. For more information, go to Template Anatomy
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html)
	// in the AWS CloudFormation User Guide.
	//
	// Conditional: You must pass TemplateURL or TemplateBody. If both are passed,
	// only TemplateBody is used.
	TemplateBody *string `min:"1" type:"string"`

	// Location of file containing the template body. The URL must point to a template
	// (max size: 460,800 bytes) that is located in an Amazon S3 bucket. For more
	// information, go to Template Anatomy (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html)
	// in the AWS CloudFormation User Guide.
	//
	// Conditional: You must pass TemplateURL or TemplateBody. If both are passed,
	// only TemplateBody is used.
	TemplateURL *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ValidateTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ValidateTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ValidateTemplateInput"}
	if s.TemplateBody != nil && len(*s.TemplateBody) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TemplateBody", 1))
	}
	if s.TemplateURL != nil && len(*s.TemplateURL) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TemplateURL", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The output for ValidateTemplate action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/ValidateTemplateOutput
type ValidateTemplateOutput struct {
	_ struct{} `type:"structure"`

	// The capabilities found within the template. If your template contains IAM
	// resources, you must specify the CAPABILITY_IAM or CAPABILITY_NAMED_IAM value
	// for this parameter when you use the CreateStack or UpdateStack actions with
	// your template; otherwise, those actions return an InsufficientCapabilities
	// error.
	//
	// For more information, see Acknowledging IAM Resources in AWS CloudFormation
	// Templates (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#capabilities).
	Capabilities []Capability `type:"list"`

	// The list of resources that generated the values in the Capabilities response
	// element.
	CapabilitiesReason *string `type:"string"`

	// A list of the transforms that are declared in the template.
	DeclaredTransforms []string `type:"list"`

	// The description found within the template.
	Description *string `min:"1" type:"string"`

	// A list of TemplateParameter structures.
	Parameters []TemplateParameter `type:"list"`
}

// String returns the string representation
func (s ValidateTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

const opValidateTemplate = "ValidateTemplate"

// ValidateTemplateRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Validates a specified template. AWS CloudFormation first checks if the template
// is valid JSON. If it isn't, AWS CloudFormation checks if the template is
// valid YAML. If both these checks fail, AWS CloudFormation returns a template
// validation error.
//
//    // Example sending a request using ValidateTemplateRequest.
//    req := client.ValidateTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/ValidateTemplate
func (c *Client) ValidateTemplateRequest(input *ValidateTemplateInput) ValidateTemplateRequest {
	op := &aws.Operation{
		Name:       opValidateTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ValidateTemplateInput{}
	}

	req := c.newRequest(op, input, &ValidateTemplateOutput{})
	return ValidateTemplateRequest{Request: req, Input: input, Copy: c.ValidateTemplateRequest}
}

// ValidateTemplateRequest is the request type for the
// ValidateTemplate API operation.
type ValidateTemplateRequest struct {
	*aws.Request
	Input *ValidateTemplateInput
	Copy  func(*ValidateTemplateInput) ValidateTemplateRequest
}

// Send marshals and sends the ValidateTemplate API request.
func (r ValidateTemplateRequest) Send(ctx context.Context) (*ValidateTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ValidateTemplateResponse{
		ValidateTemplateOutput: r.Request.Data.(*ValidateTemplateOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ValidateTemplateResponse is the response type for the
// ValidateTemplate API operation.
type ValidateTemplateResponse struct {
	*ValidateTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ValidateTemplate request.
func (r *ValidateTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
