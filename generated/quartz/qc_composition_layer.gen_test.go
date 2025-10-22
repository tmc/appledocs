// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz_test

import (
	"github.com/tmc/appledocs/generated/quartz"
)

// Suppress unused import errors
var _ = quartz.NewQCCompositionLayer

// ExampleNewQCCompositionLayerWithComposition demonstrates how to create a QCCompositionLayer instance using NewQCCompositionLayerWithComposition.
// Initializes and returns a  composition layer using the provided Quartz Composer composition.
func ExampleNewQCCompositionLayerWithComposition() {
	_ = quartz.NewQCCompositionLayerWithComposition(
		quartz.QCComposition{}, // composition QCComposition
	)
	// Output:
}
// ExampleNewQCCompositionLayerWithFile demonstrates how to create a QCCompositionLayer instance using NewQCCompositionLayerWithFile.
// Initializes and returns a composition layer using the Quartz Composer composition in the specified file.
func ExampleNewQCCompositionLayerWithFile() {
	_ = quartz.NewQCCompositionLayerWithFile(
		"/tmp/test", // path string
	)
	// Output:
}
