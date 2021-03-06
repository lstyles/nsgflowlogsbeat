// Package ec2metadata provides the client for making API calls to the
// EC2 Instance Metadata service.
//
// This package's client can be disabled completely by setting the environment
// variable "AWS_EC2_METADATA_DISABLED=true". This environment variable set to
// true instructs the SDK to disable the EC2 Metadata client. The client cannot
// be used while the environemnt variable is set to true, (case insensitive).
package ec2metadata

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/aws/defaults"
)

const disableServiceEnvVar = "AWS_EC2_METADATA_DISABLED"

// A Client is an EC2 Instance Metadata service Client.
type Client struct {
	*aws.Client
}

// New creates a new instance of the Client client with a Config.
// This client is safe to use across multiple goroutines.
//
// Example:
//     // Create a Client client from just a config.
//     svc := ec2metadata.New(cfg)
func New(config aws.Config) *Client {
	svc := &Client{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName: "EC2 Instance Metadata",
				ServiceID:   "EC2InstanceMetadata",
				EndpointsID: "ec2metadata",
				APIVersion:  "latest",
			},
		),
	}

	svc.Handlers.Unmarshal.PushBack(unmarshalHandler)
	svc.Handlers.UnmarshalError.PushBack(unmarshalError)
	svc.Handlers.Validate.Clear()
	svc.Handlers.Validate.PushBack(validateEndpointHandler)

	// Disable the EC2 Instance Metadata service if the environment variable is
	// set. This shortcirctes the service's functionality to always fail to
	// send requests.
	if strings.ToLower(os.Getenv(disableServiceEnvVar)) == "true" {
		svc.Handlers.Send.SwapNamed(aws.NamedHandler{
			Name: defaults.SendHandler.Name,
			Fn: func(r *aws.Request) {
				r.HTTPResponse = &http.Response{
					Header: http.Header{},
				}
				r.Error = awserr.New(
					aws.ErrCodeRequestCanceled,
					"EC2 IMDS access disabled via "+disableServiceEnvVar+" env var",
					nil)
			},
		})
	}

	return svc
}

func httpClientZero(c *http.Client) bool {
	return c == nil || (c.Transport == nil && c.CheckRedirect == nil && c.Jar == nil && c.Timeout == 0)
}

type metadataOutput struct {
	Content string
}

func unmarshalHandler(r *aws.Request) {
	defer r.HTTPResponse.Body.Close()
	b := &bytes.Buffer{}
	if _, err := io.Copy(b, r.HTTPResponse.Body); err != nil {
		r.Error = awserr.New("SerializationError", "unable to unmarshal EC2 metadata respose", err)
		return
	}

	if data, ok := r.Data.(*metadataOutput); ok {
		data.Content = b.String()
	}
}

func unmarshalError(r *aws.Request) {
	defer r.HTTPResponse.Body.Close()
	b := &bytes.Buffer{}
	if _, err := io.Copy(b, r.HTTPResponse.Body); err != nil {
		r.Error = awserr.New("SerializationError", "unable to unmarshal EC2 metadata error respose", err)
		return
	}

	// Response body format is not consistent between metadata endpoints.
	// Grab the error message as a string and include that as the source error
	r.Error = awserr.New("EC2MetadataError", "failed to make Client request", errors.New(b.String()))
}

func validateEndpointHandler(r *aws.Request) {
	if r.Metadata.Endpoint == "" {
		r.Error = aws.ErrMissingEndpoint
	}
}
