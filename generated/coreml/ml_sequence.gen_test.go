// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml_test

import (
	"github.com/tmc/appledocs/generated/coreml"
)

// Suppress unused import errors
var _ = coreml.NewSequence

// ExampleNewSequenceEmptySequenceWithType demonstrates how to create a Sequence instance using NewSequenceEmptySequenceWithType.
// Creates an empty sequence of strings or integers.
func ExampleNewSequenceEmptySequenceWithType() {
	_ = coreml.NewSequenceEmptySequenceWithType(
		coreml.FeatureType{}, // type FeatureType
	)
	// Output:
}
// ExampleNewSequenceWithStringArray demonstrates how to create a Sequence instance using NewSequenceWithStringArray.
// Creates a sequence of strings from a string array.
func ExampleNewSequenceWithStringArray() {
	_ = coreml.NewSequenceWithStringArray(
		[]coreml.string{}, // stringValues []string
	)
	// Output:
}
