// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders_test

import (
	"github.com/tmc/appledocs/generated/metalperformanceshaders"
)

// Suppress unused import errors
var _ = metalperformanceshaders.NewNDArrayAffineQuantizationDescriptor

// ExampleNewNDArrayAffineQuantizationDescriptor demonstrates how to create a NDArrayAffineQuantizationDescriptor instance.
func ExampleNewNDArrayAffineQuantizationDescriptor() {
	_ = metalperformanceshaders.NewNDArrayAffineQuantizationDescriptor()
	// Output:
}
// ExampleNewNDArrayAffineQuantizationDescriptorWithDataTypeHasZeroPointHasMinValue demonstrates how to create a NDArrayAffineQuantizationDescriptor instance using NewNDArrayAffineQuantizationDescriptorWithDataTypeHasZeroPointHasMinValue.
func ExampleNewNDArrayAffineQuantizationDescriptorWithDataTypeHasZeroPointHasMinValue() {
	_ = metalperformanceshaders.NewNDArrayAffineQuantizationDescriptorWithDataTypeHasZeroPointHasMinValue(
		metalperformanceshaders.DataType /* not a class type */{}, // quantizationDataType DataType /* not a class type */
		false, // hasZeroPoint bool
		false, // hasMinValue bool
	)
	// Output:
}
