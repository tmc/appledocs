// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation_test

import (
	"github.com/tmc/appledocs/generated/avfoundation"
)

// Suppress unused import errors
var _ = avfoundation.NewCaptureDataOutputSynchronizer

// ExampleNewCaptureDataOutputSynchronizerWithDataOutputs demonstrates how to create a CaptureDataOutputSynchronizer instance using NewCaptureDataOutputSynchronizerWithDataOutputs.
// Creates a capture output synchronizer for the specified capture outputs.
func ExampleNewCaptureDataOutputSynchronizerWithDataOutputs() {
	_ = avfoundation.NewCaptureDataOutputSynchronizerWithDataOutputs(
		[]avfoundation.CaptureOutput{}, // dataOutputs []CaptureOutput
	)
	// Output:
}
