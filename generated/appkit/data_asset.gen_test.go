// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewDataAsset

// ExampleNewDataAssetWithNameBundle demonstrates how to create a DataAsset instance using NewDataAssetWithNameBundle.
// Initializes and returns an object with a reference to the named data asset that’s in an asset catalog in the specified bundle.
func ExampleNewDataAssetWithNameBundle() {
	_ = appkit.NewDataAssetWithNameBundle(
		appkit.DataAssetName{}, // name DataAssetName
		appkit.Bundle{}, // bundle Bundle
	)
	// Output:
}
