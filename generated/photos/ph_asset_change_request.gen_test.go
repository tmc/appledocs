// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos_test

import (
	"github.com/tmc/appledocs/generated/photos"
)

// Suppress unused import errors
var _ = photos.NewPHAssetChangeRequest

// ExampleNewPHAssetChangeRequestForAsset demonstrates how to create a PHAssetChangeRequest instance using NewPHAssetChangeRequestForAsset.
// Creates a request for modifying the specified asset.
func ExampleNewPHAssetChangeRequestForAsset() {
	_ = photos.NewPHAssetChangeRequestForAsset(
		photos.PHAsset{}, // asset PHAsset
	)
	// Output:
}
