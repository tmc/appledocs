// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewFileCoordinator

// ExampleFileCoordinator_Cancel demonstrates using Cancel on a FileCoordinator instance.
// Cancels any active file coordination calls.
func ExampleFileCoordinator_Cancel() {
	obj := foundation.NewFileCoordinator()
	obj.Cancel()
	// Output:
	}

