// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/CreateGlobalClusterMessage
type CreateGlobalClusterInput struct {
	_ struct{} `type:"structure"`

	// The name for your database of up to 64 alpha-numeric characters. If you do
	// not provide a name, Amazon Aurora will not create a database in the global
	// database cluster you are creating.
	DatabaseName *string `type:"string"`

	// The deletion protection setting for the new global database. The global database
	// can't be deleted when this value is set to true.
	DeletionProtection *bool `type:"boolean"`

	// Provides the name of the database engine to be used for this DB cluster.
	Engine *string `type:"string"`

	// The engine version of the Aurora global database.
	EngineVersion *string `type:"string"`

	// The cluster identifier of the new global database cluster.
	GlobalClusterIdentifier *string `type:"string"`

	// The Amazon Resource Name (ARN) to use as the primary cluster of the global
	// database. This parameter is optional.
	SourceDBClusterIdentifier *string `type:"string"`

	// The storage encryption setting for the new global database cluster.
	StorageEncrypted *bool `type:"boolean"`
}

// String returns the string representation
func (s CreateGlobalClusterInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/CreateGlobalClusterResult
type CreateGlobalClusterOutput struct {
	_ struct{} `type:"structure"`

	// A data type representing an Aurora global database.
	GlobalCluster *GlobalCluster `type:"structure"`
}

// String returns the string representation
func (s CreateGlobalClusterOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateGlobalCluster = "CreateGlobalCluster"

// CreateGlobalClusterRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
//
// Creates an Aurora global database spread across multiple regions. The global
// database contains a single primary cluster with read-write capability, and
// a read-only secondary cluster that receives data from the primary cluster
// through high-speed replication performed by the Aurora storage subsystem.
//
// You can create a global database that is initially empty, and then add a
// primary cluster and a secondary cluster to it. Or you can specify an existing
// Aurora cluster during the create operation, and this cluster becomes the
// primary cluster of the global database.
//
// This action only applies to Aurora DB clusters.
//
//    // Example sending a request using CreateGlobalClusterRequest.
//    req := client.CreateGlobalClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/CreateGlobalCluster
func (c *Client) CreateGlobalClusterRequest(input *CreateGlobalClusterInput) CreateGlobalClusterRequest {
	op := &aws.Operation{
		Name:       opCreateGlobalCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateGlobalClusterInput{}
	}

	req := c.newRequest(op, input, &CreateGlobalClusterOutput{})
	return CreateGlobalClusterRequest{Request: req, Input: input, Copy: c.CreateGlobalClusterRequest}
}

// CreateGlobalClusterRequest is the request type for the
// CreateGlobalCluster API operation.
type CreateGlobalClusterRequest struct {
	*aws.Request
	Input *CreateGlobalClusterInput
	Copy  func(*CreateGlobalClusterInput) CreateGlobalClusterRequest
}

// Send marshals and sends the CreateGlobalCluster API request.
func (r CreateGlobalClusterRequest) Send(ctx context.Context) (*CreateGlobalClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateGlobalClusterResponse{
		CreateGlobalClusterOutput: r.Request.Data.(*CreateGlobalClusterOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateGlobalClusterResponse is the response type for the
// CreateGlobalCluster API operation.
type CreateGlobalClusterResponse struct {
	*CreateGlobalClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateGlobalCluster request.
func (r *CreateGlobalClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
