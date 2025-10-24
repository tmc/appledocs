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
// ExampleNewFeatureValueWithInt64 demonstrates how to create a FeatureValue instance using NewFeatureValueWithInt64.
// Creates a feature value that contains an integer.
func ExampleNewFeatureValueWithInt64() {
	_ = coreml.NewFeatureValueWithInt64(
		0, // value int64
	)
	// Output:
}
// ExampleNewFeatureValueWithPixelBuffer demonstrates how to create a FeatureValue instance using NewFeatureValueWithPixelBuffer.
// Creates a feature value that contains an image from a pixel buffer.
func ExampleNewFeatureValueWithPixelBuffer() {
	_ = coreml.NewFeatureValueWithPixelBuffer(
		coreml.PixelBufferRef /* not a class type */{}, // value PixelBufferRef /* not a class type */
	)
	// Output:
}
