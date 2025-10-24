// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation_test

import (
	"github.com/tmc/appledocs/generated/avfoundation"
)

// Suppress unused import errors
var _ = avfoundation.NewAssetWriter

// ExampleAssetWriter_CancelWriting demonstrates using CancelWriting on a AssetWriter instance.
// Cancels the creation of the output file.
func ExampleAssetWriter_CancelWriting() {
	obj := avfoundation.NewAssetWriter()
	obj.CancelWriting()
	// Output:
	}

// ExampleAssetWriter_FlushSegment demonstrates using FlushSegment on a AssetWriter instance.
// Closes the current segment and outputs it to a delegate method.
func ExampleAssetWriter_FlushSegment() {
	obj := avfoundation.NewAssetWriter()
	obj.FlushSegment()
	// Output:
	}

