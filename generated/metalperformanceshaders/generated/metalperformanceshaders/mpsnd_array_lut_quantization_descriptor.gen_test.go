// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders_test

import (
	"github.com/tmc/appledocs/generated/metalperformanceshaders"
)

// Suppress unused import errors
var _ = metalperformanceshaders.NewNDArrayLUTQuantizationDescriptor

// ExampleNewNDArrayLUTQuantizationDescriptorWithDataType demonstrates how to create a NDArrayLUTQuantizationDescriptor instance using NewNDArrayLUTQuantizationDescriptorWithDataType.
func ExampleNewNDArrayLUTQuantizationDescriptorWithDataType() {
	_ = metalperformanceshaders.NewNDArrayLUTQuantizationDescriptorWithDataType(
		metalperformanceshaders.DataType{}, // quantizationDataType DataType
	)
	// Output:
}
// ExampleNewNDArrayLUTQuantizationDescriptorWithDataTypeVectorAxis demonstrates how to create a NDArrayLUTQuantizationDescriptor instance using NewNDArrayLUTQuantizationDescriptorWithDataTypeVectorAxis.
func ExampleNewNDArrayLUTQuantizationDescriptorWithDataTypeVectorAxis() {
	_ = metalperformanceshaders.NewNDArrayLUTQuantizationDescriptorWithDataTypeVectorAxis(
		metalperformanceshaders.DataType{}, // quantizationDataType DataType
		0, // vectorAxis uint
	)
	// Output:
}
