// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders_test

import (
	"github.com/tmc/appledocs/generated/metalperformanceshaders"
)

// Suppress unused import errors
var _ = metalperformanceshaders.NewNDArrayAffineQuantizationDescriptor

// ExampleNewNDArrayAffineQuantizationDescriptorWithDataTypeHasZeroPointHasMinValue demonstrates how to create a NDArrayAffineQuantizationDescriptor instance using NewNDArrayAffineQuantizationDescriptorWithDataTypeHasZeroPointHasMinValue.
func ExampleNewNDArrayAffineQuantizationDescriptorWithDataTypeHasZeroPointHasMinValue() {
	_ = metalperformanceshaders.NewNDArrayAffineQuantizationDescriptorWithDataTypeHasZeroPointHasMinValue(
		metalperformanceshaders.DataType{}, // quantizationDataType DataType
		false, // hasZeroPoint bool
		false, // hasMinValue bool
	)
	// Output:
}
