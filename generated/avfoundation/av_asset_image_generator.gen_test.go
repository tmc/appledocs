// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation_test

import (
	"github.com/tmc/appledocs/generated/avfoundation"
)

// Suppress unused import errors
var _ = avfoundation.NewAssetImageGenerator

// ExampleAssetImageGenerator_CancelAllCGImageGeneration demonstrates using CancelAllCGImageGeneration on a AssetImageGenerator instance.
// Cancels all pending image generation requests.
func ExampleAssetImageGenerator_CancelAllCGImageGeneration() {
	obj := avfoundation.NewAssetImageGenerator()
	obj.CancelAllCGImageGeneration()
	// Output:
	}

