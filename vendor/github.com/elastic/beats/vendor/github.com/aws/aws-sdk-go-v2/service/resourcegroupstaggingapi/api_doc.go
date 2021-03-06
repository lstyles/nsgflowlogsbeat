// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package resourcegroupstaggingapi provides the client and types for making API
// requests to AWS Resource Groups Tagging API.
//
// This guide describes the API operations for the resource groups tagging.
//
// A tag is a label that you assign to an AWS resource. A tag consists of a
// key and a value, both of which you define. For example, if you have two Amazon
// EC2 instances, you might assign both a tag key of "Stack." But the value
// of "Stack" might be "Testing" for one and "Production" for the other.
//
// Tagging can help you organize your resources and enables you to simplify
// resource management, access management and cost allocation. For more information
// about tagging, see Working with Tag Editor (http://docs.aws.amazon.com/awsconsolehelpdocs/latest/gsg/tag-editor.html)
// and Working with Resource Groups (http://docs.aws.amazon.com/awsconsolehelpdocs/latest/gsg/resource-groups.html).
// For more information about permissions you need to use the resource groups
// tagging APIs, see Obtaining Permissions for Resource Groups (http://docs.aws.amazon.com/awsconsolehelpdocs/latest/gsg/obtaining-permissions-for-resource-groups.html)
// and Obtaining Permissions for Tagging (http://docs.aws.amazon.com/awsconsolehelpdocs/latest/gsg/obtaining-permissions-for-tagging.html).
//
// You can use the resource groups tagging APIs to complete the following tasks:
//
//    * Tag and untag supported resources located in the specified region for
//    the AWS account
//
//    * Use tag-based filters to search for resources located in the specified
//    region for the AWS account
//
//    * List all existing tag keys in the specified region for the AWS account
//
//    * List all existing values for the specified key in the specified region
//    for the AWS account
//
// Not all resources can have tags. For a lists of resources that you can tag,
// see Supported Resources (http://docs.aws.amazon.com/awsconsolehelpdocs/latest/gsg/supported-resources.html)
// in the AWS Resource Groups and Tag Editor User Guide.
//
// To make full use of the resource groups tagging APIs, you might need additional
// IAM permissions, including permission to access the resources of individual
// services as well as permission to view and apply tags to those resources.
// For more information, see Obtaining Permissions for Tagging (http://docs.aws.amazon.com/awsconsolehelpdocs/latest/gsg/obtaining-permissions-for-tagging.html)
// in the AWS Resource Groups and Tag Editor User Guide.
//
// See https://docs.aws.amazon.com/goto/WebAPI/resourcegroupstaggingapi-2017-01-26 for more information on this service.
//
// See resourcegroupstaggingapi package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/resourcegroupstaggingapi/
//
// Using the Client
//
// To use AWS Resource Groups Tagging API with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS Resource Groups Tagging API client for more information on
// creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/resourcegroupstaggingapi/#New
package resourcegroupstaggingapi
