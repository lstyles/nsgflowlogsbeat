package wafv2

import (
	"github.com/awslabs/goformation/v4/cloudformation/policies"
)

// RuleGroup_OrStatementTwo AWS CloudFormation Resource (AWS::WAFv2::RuleGroup.OrStatementTwo)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-orstatementtwo.html
type RuleGroup_OrStatementTwo struct {

	// Statements AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-orstatementtwo.html#cfn-wafv2-rulegroup-orstatementtwo-statements
	Statements *RuleGroup_StatementThrees `json:"Statements,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *RuleGroup_OrStatementTwo) AWSCloudFormationType() string {
	return "AWS::WAFv2::RuleGroup.OrStatementTwo"
}
