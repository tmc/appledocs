// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml_test

import (
	"github.com/tmc/appledocs/generated/coreml"
)

// Suppress unused import errors
var _ = coreml.NewFeatureValue

// ExampleNewFeatureValueUndefinedFeatureValueWithType demonstrates how to create a FeatureValue instance using NewFeatureValueUndefinedFeatureValueWithType.
// Creates a feature value with a type that represents an undefined or missing value.
func ExampleNewFeatureValueUndefinedFeatureValueWithType() {
	_ = coreml.NewFeatureValueUndefinedFeatureValueWithType(
		coreml.FeatureType{}, // type FeatureType
	)
	// Output:
}
// ExampleNewFeatureValueWithDouble demonstrates how to create a FeatureValue instance using NewFeatureValueWithDouble.
// Creates a feature value that contains a double.
func ExampleNewFeatureValueWithDouble() {
	_ = coreml.NewFeatureValueWithDouble(
		0.0, // value float64
	)
	// Output:
}
// ExampleNewFeatureValueWithMultiArray demonstrates how to create a FeatureValue instance using NewFeatureValueWithMultiArray.
// Creates a feature value that contains a multidimensional array.
func ExampleNewFeatureValueWithMultiArray() {
	_ = coreml.NewFeatureValueWithMultiArray(
		coreml.MLMultiArray{}, // value MLMultiArray
	)
	// Output:
}
// ExampleNewFeatureValueWithPixelBuffer demonstrates how to create a FeatureValue instance using NewFeatureValueWithPixelBuffer.
// Creates a feature value that contains an image from a pixel buffer.
func ExampleNewFeatureValueWithPixelBuffer() {
	_ = coreml.NewFeatureValueWithPixelBuffer(
		coreml.PixelBufferRef{}, // value PixelBufferRef
	)
	// Output:
}
// ExampleNewFeatureValueWithSequence demonstrates how to create a FeatureValue instance using NewFeatureValueWithSequence.
// Creates a feature value that contains a sequence.
func ExampleNewFeatureValueWithSequence() {
	_ = coreml.NewFeatureValueWithSequence(
		coreml.MLSequence{}, // sequence MLSequence
	)
	// Output:
}
// ExampleNewFeatureValueWithString demonstrates how to create a FeatureValue instance using NewFeatureValueWithString.
// Creates a feature value that contains a string.
func ExampleNewFeatureValueWithString() {
	_ = coreml.NewFeatureValueWithString(
		"value", // value string
	)
	// Output:
}
