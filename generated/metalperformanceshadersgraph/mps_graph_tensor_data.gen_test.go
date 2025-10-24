// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph_test

import (
	"github.com/tmc/appledocs/generated/metalperformanceshadersgraph"
)

// Suppress unused import errors
var _ = metalperformanceshadersgraph.NewGraphTensorData

// ExampleNewGraphTensorDataWithMPSImageBatch demonstrates how to create a GraphTensorData instance using NewGraphTensorDataWithMPSImageBatch.
// Initializes a tensor data with an MPS image batch.
func ExampleNewGraphTensorDataWithMPSImageBatch() {
	_ = metalperformanceshadersgraph.NewGraphTensorDataWithMPSImageBatch(
		metalperformanceshadersgraph.ImageBatch /* not a class type */{}, // imageBatch ImageBatch /* not a class type */
	)
	// Output:
}
// ExampleGraphTensorData_Mpsndarray demonstrates using Mpsndarray on a GraphTensorData instance.
// Return an mpsndarray object will copy contents if the contents are not stored in an MPS ndarray.
func ExampleGraphTensorData_Mpsndarray() {
	obj := metalperformanceshadersgraph.NewGraphTensorData()
	_ = obj.Mpsndarray()
	// Output:
	}

