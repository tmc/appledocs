// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewDrawer

// ExampleNewDrawerWithContentSizePreferredEdge demonstrates how to create a Drawer instance using NewDrawerWithContentSizePreferredEdge.
// Creates a new drawer with the given size on the specified edge of the parent window.
func ExampleNewDrawerWithContentSizePreferredEdge() {
	_ = appkit.NewDrawerWithContentSizePreferredEdge(
		appkit.Size /* not a class type */{}, // contentSize Size /* not a class type */
		appkit.RectEdge /* not a class type */{}, // edge RectEdge /* not a class type */
	)
	// Output:
}
