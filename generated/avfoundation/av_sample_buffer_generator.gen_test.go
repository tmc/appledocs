// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation_test

import (
	"github.com/tmc/appledocs/generated/avfoundation"
)

// Suppress unused import errors
var _ = avfoundation.NewSampleBufferGenerator

// ExampleSampleBufferGenerator_MakeBatch demonstrates using MakeBatch on a SampleBufferGenerator instance.
// Creates a batch object to handle generating multiple sample buffers.
func ExampleSampleBufferGenerator_MakeBatch() {
	obj := avfoundation.NewSampleBufferGenerator()
	_ = obj.MakeBatch()
	// Output:
	}

