// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewDraggingImageComponent

// ExampleNewDraggingImageComponentWithKey demonstrates how to create a DraggingImageComponent instance using NewDraggingImageComponentWithKey.
// Initializes and returns a dragging image component with the specified key.
func ExampleNewDraggingImageComponentWithKey() {
	_ = appkit.NewDraggingImageComponentWithKey(
		appkit.DraggingImageComponentKey{}, // key DraggingImageComponentKey
	)
	// Output:
}
