// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute_test

import (
	"github.com/tmc/appledocs/generated/mlcompute"
)

// Suppress unused import errors
var _ = mlcompute.NewCInferenceGraph

// ExampleNewCInferenceGraphWithGraphObjects demonstrates how to create a CInferenceGraph instance using NewCInferenceGraphWithGraphObjects.
// Creates an inference graph with the layers from the graph objects you specify.
func ExampleNewCInferenceGraphWithGraphObjects() {
	_ = mlcompute.NewCInferenceGraphWithGraphObjects(
		[]mlcompute.CGraph{}, // graphObjects []CGraph
	)
	// Output:
}
