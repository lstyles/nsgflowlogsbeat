package image // import "github.com/docker/docker/api/types/image"

// ----------------------------------------------------------------------------
// DO NOT EDIT THIS FILE
// This file was generated by `swagger generate operation`
//
// See hack/generate-swagger-api.sh
// ----------------------------------------------------------------------------

// HistoryResponseItem individual image layer information in response to ImageHistory operation
// swagger:model HistoryResponseItem
type HistoryResponseItem struct {

	// comment
	// Required: true
	Comment string `json:"Comment"`

	// created
	// Required: true
	Created int64 `json:"Created"`

	// created by
	// Required: true
	CreatedBy string `json:"CreatedBy"`

	// Id
	// Required: true
	ID string `json:"Id"`

	// size
	// Required: true
	Size int64 `json:"Size"`

	// tags
	// Required: true
	Tags []string `json:"Tags"`
}
