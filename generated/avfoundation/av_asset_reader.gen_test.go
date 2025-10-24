// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation_test

import (
	"github.com/tmc/appledocs/generated/avfoundation"
)

// Suppress unused import errors
var _ = avfoundation.NewAssetReader

// ExampleAssetReader_CancelReading demonstrates using CancelReading on a AssetReader instance.
// Cancels any background work and stops the reader’s outputs from reading more samples.
func ExampleAssetReader_CancelReading() {
	obj := avfoundation.NewAssetReader()
	obj.CancelReading()
	// Output:
	}

