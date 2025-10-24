// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation_test

import (
	"github.com/tmc/appledocs/generated/avfoundation"
)

// Suppress unused import errors
var _ = avfoundation.NewAssetWriterInput

// ExampleAssetWriterInput_MarkAsFinished demonstrates using MarkAsFinished on a AssetWriterInput instance.
// Marks the input as finished to indicate that you’re done appending samples to it.
func ExampleAssetWriterInput_MarkAsFinished() {
	obj := avfoundation.NewAssetWriterInput()
	obj.MarkAsFinished()
	// Output:
	}

// ExampleAssetWriterInput_MarkCurrentPassAsFinished demonstrates using MarkCurrentPassAsFinished on a AssetWriterInput instance.
// Tells the input to analyze the appended media to determine whether it can improve the results by reencoding certain segments.
func ExampleAssetWriterInput_MarkCurrentPassAsFinished() {
	obj := avfoundation.NewAssetWriterInput()
	obj.MarkCurrentPassAsFinished()
	// Output:
	}

