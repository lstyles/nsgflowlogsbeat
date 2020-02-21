// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/CopyDBSnapshotMessage
type CopyDBSnapshotInput struct {
	_ struct{} `type:"structure"`

	// True to copy all tags from the source DB snapshot to the target DB snapshot,
	// and otherwise false. The default is false.
	CopyTags *bool `type:"boolean"`

	// DestinationRegion is used for presigning the request to a given region.
	DestinationRegion *string `type:"string"`

	// The AWS KMS key ID for an encrypted DB snapshot. The KMS key ID is the Amazon
	// Resource Name (ARN), KMS key identifier, or the KMS key alias for the KMS
	// encryption key.
	//
	// If you copy an encrypted DB snapshot from your AWS account, you can specify
	// a value for this parameter to encrypt the copy with a new KMS encryption
	// key. If you don't specify a value for this parameter, then the copy of the
	// DB snapshot is encrypted with the same KMS key as the source DB snapshot.
	//
	// If you copy an encrypted DB snapshot that is shared from another AWS account,
	// then you must specify a value for this parameter.
	//
	// If you specify this parameter when you copy an unencrypted snapshot, the
	// copy is encrypted.
	//
	// If you copy an encrypted snapshot to a different AWS Region, then you must
	// specify a KMS key for the destination AWS Region. KMS encryption keys are
	// specific to the AWS Region that they are created in, and you can't use encryption
	// keys from one AWS Region in another AWS Region.
	KmsKeyId *string `type:"string"`

	// The name of an option group to associate with the copy of the snapshot.
	//
	// Specify this option if you are copying a snapshot from one AWS Region to
	// another, and your DB instance uses a nondefault option group. If your source
	// DB instance uses Transparent Data Encryption for Oracle or Microsoft SQL
	// Server, you must specify this option when copying across AWS Regions. For
	// more information, see Option Group Considerations (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_CopySnapshot.html#USER_CopySnapshot.Options)
	// in the Amazon RDS User Guide.
	OptionGroupName *string `type:"string"`

	// The URL that contains a Signature Version 4 signed request for the CopyDBSnapshot
	// API action in the source AWS Region that contains the source DB snapshot
	// to copy.
	//
	// You must specify this parameter when you copy an encrypted DB snapshot from
	// another AWS Region by using the Amazon RDS API. You can specify the --source-region
	// option instead of this parameter when you copy an encrypted DB snapshot from
	// another AWS Region by using the AWS CLI.
	//
	// The presigned URL must be a valid request for the CopyDBSnapshot API action
	// that can be executed in the source AWS Region that contains the encrypted
	// DB snapshot to be copied. The presigned URL request must contain the following
	// parameter values:
	//
	//    * DestinationRegion - The AWS Region that the encrypted DB snapshot is
	//    copied to. This AWS Region is the same one where the CopyDBSnapshot action
	//    is called that contains this presigned URL. For example, if you copy an
	//    encrypted DB snapshot from the us-west-2 AWS Region to the us-east-1 AWS
	//    Region, then you call the CopyDBSnapshot action in the us-east-1 AWS Region
	//    and provide a presigned URL that contains a call to the CopyDBSnapshot
	//    action in the us-west-2 AWS Region. For this example, the DestinationRegion
	//    in the presigned URL must be set to the us-east-1 AWS Region.
	//
	//    * KmsKeyId - The AWS KMS key identifier for the key to use to encrypt
	//    the copy of the DB snapshot in the destination AWS Region. This is the
	//    same identifier for both the CopyDBSnapshot action that is called in the
	//    destination AWS Region, and the action contained in the presigned URL.
	//
	//    * SourceDBSnapshotIdentifier - The DB snapshot identifier for the encrypted
	//    snapshot to be copied. This identifier must be in the Amazon Resource
	//    Name (ARN) format for the source AWS Region. For example, if you are copying
	//    an encrypted DB snapshot from the us-west-2 AWS Region, then your SourceDBSnapshotIdentifier
	//    looks like the following example: arn:aws:rds:us-west-2:123456789012:snapshot:mysql-instance1-snapshot-20161115.
	//
	// To learn how to generate a Signature Version 4 signed request, see Authenticating
	// Requests: Using Query Parameters (AWS Signature Version 4) (https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-query-string-auth.html)
	// and Signature Version 4 Signing Process (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html).
	PreSignedUrl *string `type:"string"`

	// The identifier for the source DB snapshot.
	//
	// If the source snapshot is in the same AWS Region as the copy, specify a valid
	// DB snapshot identifier. For example, you might specify rds:mysql-instance1-snapshot-20130805.
	//
	// If the source snapshot is in a different AWS Region than the copy, specify
	// a valid DB snapshot ARN. For example, you might specify arn:aws:rds:us-west-2:123456789012:snapshot:mysql-instance1-snapshot-20130805.
	//
	// If you are copying from a shared manual DB snapshot, this parameter must
	// be the Amazon Resource Name (ARN) of the shared DB snapshot.
	//
	// If you are copying an encrypted snapshot this parameter must be in the ARN
	// format for the source AWS Region, and must match the SourceDBSnapshotIdentifier
	// in the PreSignedUrl parameter.
	//
	// Constraints:
	//
	//    * Must specify a valid system snapshot in the "available" state.
	//
	// Example: rds:mydb-2012-04-02-00-01
	//
	// Example: arn:aws:rds:us-west-2:123456789012:snapshot:mysql-instance1-snapshot-20130805
	//
	// SourceDBSnapshotIdentifier is a required field
	SourceDBSnapshotIdentifier *string `type:"string" required:"true"`

	// SourceRegion is the source region where the resource exists. This is not
	// sent over the wire and is only used for presigning. This value should always
	// have the same region as the source ARN.
	SourceRegion *string `type:"string" ignore:"true"`

	// A list of tags. For more information, see Tagging Amazon RDS Resources (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html)
	// in the Amazon RDS User Guide.
	Tags []Tag `locationNameList:"Tag" type:"list"`

	// The identifier for the copy of the snapshot.
	//
	// Constraints:
	//
	//    * Can't be null, empty, or blank
	//
	//    * Must contain from 1 to 255 letters, numbers, or hyphens
	//
	//    * First character must be a letter
	//
	//    * Can't end with a hyphen or contain two consecutive hyphens
	//
	// Example: my-db-snapshot
	//
	// TargetDBSnapshotIdentifier is a required field
	TargetDBSnapshotIdentifier *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CopyDBSnapshotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CopyDBSnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CopyDBSnapshotInput"}

	if s.SourceDBSnapshotIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceDBSnapshotIdentifier"))
	}

	if s.TargetDBSnapshotIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetDBSnapshotIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/CopyDBSnapshotResult
type CopyDBSnapshotOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details of an Amazon RDS DB snapshot.
	//
	// This data type is used as a response element in the DescribeDBSnapshots action.
	DBSnapshot *DBSnapshot `type:"structure"`
}

// String returns the string representation
func (s CopyDBSnapshotOutput) String() string {
	return awsutil.Prettify(s)
}

const opCopyDBSnapshot = "CopyDBSnapshot"

// CopyDBSnapshotRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Copies the specified DB snapshot. The source DB snapshot must be in the "available"
// state.
//
// You can copy a snapshot from one AWS Region to another. In that case, the
// AWS Region where you call the CopyDBSnapshot action is the destination AWS
// Region for the DB snapshot copy.
//
// For more information about copying snapshots, see Copying a DB Snapshot (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_CopyDBSnapshot.html)
// in the Amazon RDS User Guide.
//
//    // Example sending a request using CopyDBSnapshotRequest.
//    req := client.CopyDBSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/CopyDBSnapshot
func (c *Client) CopyDBSnapshotRequest(input *CopyDBSnapshotInput) CopyDBSnapshotRequest {
	op := &aws.Operation{
		Name:       opCopyDBSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CopyDBSnapshotInput{}
	}

	req := c.newRequest(op, input, &CopyDBSnapshotOutput{})
	return CopyDBSnapshotRequest{Request: req, Input: input, Copy: c.CopyDBSnapshotRequest}
}

// CopyDBSnapshotRequest is the request type for the
// CopyDBSnapshot API operation.
type CopyDBSnapshotRequest struct {
	*aws.Request
	Input *CopyDBSnapshotInput
	Copy  func(*CopyDBSnapshotInput) CopyDBSnapshotRequest
}

// Send marshals and sends the CopyDBSnapshot API request.
func (r CopyDBSnapshotRequest) Send(ctx context.Context) (*CopyDBSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CopyDBSnapshotResponse{
		CopyDBSnapshotOutput: r.Request.Data.(*CopyDBSnapshotOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CopyDBSnapshotResponse is the response type for the
// CopyDBSnapshot API operation.
type CopyDBSnapshotResponse struct {
	*CopyDBSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CopyDBSnapshot request.
func (r *CopyDBSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
