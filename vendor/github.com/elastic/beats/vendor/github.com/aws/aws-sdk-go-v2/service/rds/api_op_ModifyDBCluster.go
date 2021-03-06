// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/ModifyDBClusterMessage
type ModifyDBClusterInput struct {
	_ struct{} `type:"structure"`

	// A value that specifies whether the modifications in this request and any
	// pending modifications are asynchronously applied as soon as possible, regardless
	// of the PreferredMaintenanceWindow setting for the DB cluster. If this parameter
	// is set to false, changes to the DB cluster are applied during the next maintenance
	// window.
	//
	// The ApplyImmediately parameter only affects the EnableIAMDatabaseAuthentication,
	// MasterUserPassword, and NewDBClusterIdentifier values. If you set the ApplyImmediately
	// parameter value to false, then changes to the EnableIAMDatabaseAuthentication,
	// MasterUserPassword, and NewDBClusterIdentifier values are applied during
	// the next maintenance window. All other changes are applied immediately, regardless
	// of the value of the ApplyImmediately parameter.
	//
	// Default: false
	ApplyImmediately *bool `type:"boolean"`

	// The target backtrack window, in seconds. To disable backtracking, set this
	// value to 0.
	//
	// Default: 0
	//
	// Constraints:
	//
	//    * If specified, this value must be set to a number from 0 to 259,200 (72
	//    hours).
	BacktrackWindow *int64 `type:"long"`

	// The number of days for which automated backups are retained. You must specify
	// a minimum value of 1.
	//
	// Default: 1
	//
	// Constraints:
	//
	//    * Must be a value from 1 to 35
	BackupRetentionPeriod *int64 `type:"integer"`

	// The configuration setting for the log types to be enabled for export to CloudWatch
	// Logs for a specific DB cluster.
	CloudwatchLogsExportConfiguration *CloudwatchLogsExportConfiguration `type:"structure"`

	// True to copy all tags from the DB cluster to snapshots of the DB cluster,
	// and otherwise false. The default is false.
	CopyTagsToSnapshot *bool `type:"boolean"`

	// The DB cluster identifier for the cluster being modified. This parameter
	// is not case-sensitive.
	//
	// Constraints:
	//
	//    * Must match the identifier of an existing DBCluster.
	//
	// DBClusterIdentifier is a required field
	DBClusterIdentifier *string `type:"string" required:"true"`

	// The name of the DB cluster parameter group to use for the DB cluster.
	DBClusterParameterGroupName *string `type:"string"`

	// Indicates if the DB cluster has deletion protection enabled. The database
	// can't be deleted when this value is set to true.
	DeletionProtection *bool `type:"boolean"`

	//
	// HTTP endpoint functionality is in beta for Aurora Serverless and is subject
	// to change.
	//
	// A value that indicates whether to enable the HTTP endpoint for an Aurora
	// Serverless DB cluster. By default, the HTTP endpoint is disabled.
	//
	// When enabled, the HTTP endpoint provides a connectionless web service API
	// for running SQL queries on the Aurora Serverless DB cluster. You can also
	// query your database from inside the RDS console with the query editor.
	//
	// For more information about Aurora Serverless, see Using Amazon Aurora Serverless
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.html)
	// in the Amazon Aurora User Guide.
	EnableHttpEndpoint *bool `type:"boolean"`

	// True to enable mapping of AWS Identity and Access Management (IAM) accounts
	// to database accounts, and otherwise false.
	//
	// Default: false
	EnableIAMDatabaseAuthentication *bool `type:"boolean"`

	// The version number of the database engine to which you want to upgrade. Changing
	// this parameter results in an outage. The change is applied during the next
	// maintenance window unless the ApplyImmediately parameter is set to true.
	//
	// For a list of valid engine versions, use the DescribeDBEngineVersions action.
	EngineVersion *string `type:"string"`

	// The new password for the master database user. This password can contain
	// any printable ASCII character except "/", """, or "@".
	//
	// Constraints: Must contain from 8 to 41 characters.
	MasterUserPassword *string `type:"string"`

	// The new DB cluster identifier for the DB cluster when renaming a DB cluster.
	// This value is stored as a lowercase string.
	//
	// Constraints:
	//
	//    * Must contain from 1 to 63 letters, numbers, or hyphens
	//
	//    * The first character must be a letter
	//
	//    * Can't end with a hyphen or contain two consecutive hyphens
	//
	// Example: my-cluster2
	NewDBClusterIdentifier *string `type:"string"`

	// A value that indicates that the DB cluster should be associated with the
	// specified option group. Changing this parameter doesn't result in an outage
	// except in the following case, and the change is applied during the next maintenance
	// window unless the ApplyImmediately parameter is set to true for this request.
	// If the parameter change results in an option group that enables OEM, this
	// change can cause a brief (sub-second) period during which new connections
	// are rejected but existing connections are not interrupted.
	//
	// Permanent options can't be removed from an option group. The option group
	// can't be removed from a DB cluster once it is associated with a DB cluster.
	OptionGroupName *string `type:"string"`

	// The port number on which the DB cluster accepts connections.
	//
	// Constraints: Value must be 1150-65535
	//
	// Default: The same port as the original DB cluster.
	Port *int64 `type:"integer"`

	// The daily time range during which automated backups are created if automated
	// backups are enabled, using the BackupRetentionPeriod parameter.
	//
	// The default is a 30-minute window selected at random from an 8-hour block
	// of time for each AWS Region. To see the time blocks available, see Adjusting
	// the Preferred DB Cluster Maintenance Window (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_UpgradeDBInstance.Maintenance.html#AdjustingTheMaintenanceWindow.Aurora)
	// in the Amazon Aurora User Guide.
	//
	// Constraints:
	//
	//    * Must be in the format hh24:mi-hh24:mi.
	//
	//    * Must be in Universal Coordinated Time (UTC).
	//
	//    * Must not conflict with the preferred maintenance window.
	//
	//    * Must be at least 30 minutes.
	PreferredBackupWindow *string `type:"string"`

	// The weekly time range during which system maintenance can occur, in Universal
	// Coordinated Time (UTC).
	//
	// Format: ddd:hh24:mi-ddd:hh24:mi
	//
	// The default is a 30-minute window selected at random from an 8-hour block
	// of time for each AWS Region, occurring on a random day of the week. To see
	// the time blocks available, see Adjusting the Preferred DB Cluster Maintenance
	// Window (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_UpgradeDBInstance.Maintenance.html#AdjustingTheMaintenanceWindow.Aurora)
	// in the Amazon Aurora User Guide.
	//
	// Valid Days: Mon, Tue, Wed, Thu, Fri, Sat, Sun.
	//
	// Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string `type:"string"`

	// The scaling properties of the DB cluster. You can only modify scaling properties
	// for DB clusters in serverless DB engine mode.
	ScalingConfiguration *ScalingConfiguration `type:"structure"`

	// A list of VPC security groups that the DB cluster will belong to.
	VpcSecurityGroupIds []string `locationNameList:"VpcSecurityGroupId" type:"list"`
}

// String returns the string representation
func (s ModifyDBClusterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyDBClusterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyDBClusterInput"}

	if s.DBClusterIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBClusterIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/ModifyDBClusterResult
type ModifyDBClusterOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details of an Amazon Aurora DB cluster.
	//
	// This data type is used as a response element in the DescribeDBClusters, StopDBCluster,
	// and StartDBCluster actions.
	DBCluster *DBCluster `type:"structure"`
}

// String returns the string representation
func (s ModifyDBClusterOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyDBCluster = "ModifyDBCluster"

// ModifyDBClusterRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Modify a setting for an Amazon Aurora DB cluster. You can change one or more
// database configuration parameters by specifying these parameters and the
// new values in the request. For more information on Amazon Aurora, see What
// Is Amazon Aurora? (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide.
//
// This action only applies to Aurora DB clusters.
//
//    // Example sending a request using ModifyDBClusterRequest.
//    req := client.ModifyDBClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/ModifyDBCluster
func (c *Client) ModifyDBClusterRequest(input *ModifyDBClusterInput) ModifyDBClusterRequest {
	op := &aws.Operation{
		Name:       opModifyDBCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyDBClusterInput{}
	}

	req := c.newRequest(op, input, &ModifyDBClusterOutput{})
	return ModifyDBClusterRequest{Request: req, Input: input, Copy: c.ModifyDBClusterRequest}
}

// ModifyDBClusterRequest is the request type for the
// ModifyDBCluster API operation.
type ModifyDBClusterRequest struct {
	*aws.Request
	Input *ModifyDBClusterInput
	Copy  func(*ModifyDBClusterInput) ModifyDBClusterRequest
}

// Send marshals and sends the ModifyDBCluster API request.
func (r ModifyDBClusterRequest) Send(ctx context.Context) (*ModifyDBClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyDBClusterResponse{
		ModifyDBClusterOutput: r.Request.Data.(*ModifyDBClusterOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyDBClusterResponse is the response type for the
// ModifyDBCluster API operation.
type ModifyDBClusterResponse struct {
	*ModifyDBClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyDBCluster request.
func (r *ModifyDBClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
