// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewThread


// ExampleNewThread demonstrates how to create a Thread instance.
// Returns an initialized   object.
func ExampleNewThread() {
	_ = foundation.NewThread()
	// Output:
}

// ExampleNewThreadWithTargetSelectorObject demonstrates how to create a Thread instance using NewThreadWithTargetSelectorObject.
// Returns an   object initialized with the given arguments.
func ExampleNewThreadWithTargetSelectorObject() {
	_ = foundation.NewThreadWithTargetSelectorObject(
		0, // target objc.ID
		0, // selector objc.SEL
		0, // argument objc.ID
	)
	// Output:
}


