// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml_test

import (
	"github.com/tmc/appledocs/generated/coreml"
)

// Suppress unused import errors
var _ = coreml.NewMultiArray

// ExampleNewMultiArrayByConcatenatingMultiArraysAlongAxisDataType demonstrates how to create a MultiArray instance using NewMultiArrayByConcatenatingMultiArraysAlongAxisDataType.
// Merges an array of multiarrays into one multiarray along an axis.
func ExampleNewMultiArrayByConcatenatingMultiArraysAlongAxisDataType() {
	_ = coreml.NewMultiArrayByConcatenatingMultiArraysAlongAxisDataType(
		[]coreml.MultiArray{}, // multiArrays []MultiArray
		0, // axis int
		coreml.MultiArrayDataType{}, // dataType MultiArrayDataType
	)
	// Output:
}
