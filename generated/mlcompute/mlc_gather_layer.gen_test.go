// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute_test

import (
	"github.com/tmc/appledocs/generated/mlcompute"
)

// Suppress unused import errors
var _ = mlcompute.NewCGatherLayer

// ExampleNewCGatherLayerWithDimension demonstrates how to create a CGatherLayer instance using NewCGatherLayerWithDimension.
// Creates a gather layer with the dimension you specify.
func ExampleNewCGatherLayerWithDimension() {
	_ = mlcompute.NewCGatherLayerWithDimension(
		0, // dimension uint
	)
	// Output:
}
