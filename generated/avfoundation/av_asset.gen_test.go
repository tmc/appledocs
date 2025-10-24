// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation_test

import (
	"github.com/tmc/appledocs/generated/avfoundation"
)

// Suppress unused import errors
var _ = avfoundation.NewAsset

// ExampleAsset_CancelLoading demonstrates using CancelLoading on a Asset instance.
// Cancels all pending requests to asynchronously load property values.
func ExampleAsset_CancelLoading() {
	obj := avfoundation.NewAsset()
	obj.CancelLoading()
	// Output:
	}

