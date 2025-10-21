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
// ExampleNewFeatureValueWithMultiArray demonstrates how to create a FeatureValue instance using NewFeatureValueWithMultiArray.
// Creates a feature value that contains a multidimensional array.
func ExampleNewFeatureValueWithMultiArray() {
	_ = coreml.NewFeatureValueWithMultiArray(
		coreml.MLMultiArray{}, // value MLMultiArray
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
