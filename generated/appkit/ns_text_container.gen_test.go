// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewTextContainer

// ExampleNewTextContainerWithContainerSize demonstrates how to create a TextContainer instance using NewTextContainerWithContainerSize.
// Initializes a text container with a specified bounding rectangle.
func ExampleNewTextContainerWithContainerSize() {
	_ = appkit.NewTextContainerWithContainerSize(
		appkit.Size /* not a class type */{}, // aContainerSize Size /* not a class type */
	)
	// Output:
}
