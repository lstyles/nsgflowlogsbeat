package container // import "github.com/docker/docker/api/types/container"

// ----------------------------------------------------------------------------
// DO NOT EDIT THIS FILE
// This file was generated by `swagger generate operation`
//
// See hack/generate-swagger-api.sh
// ----------------------------------------------------------------------------

// ContainerUpdateOKBody OK response to ContainerUpdate operation
// swagger:model ContainerUpdateOKBody
type ContainerUpdateOKBody struct {

	// warnings
	// Required: true
	Warnings []string `json:"Warnings"`
}
