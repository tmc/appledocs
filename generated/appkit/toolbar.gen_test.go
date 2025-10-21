// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewToolbar

// ExampleNewToolbar demonstrates how to create a Toolbar instance.
// Creates a new toolbar with an empty identifier string.
func ExampleNewToolbar() {
	_ = appkit.NewToolbar()
	// Output:
}
// ExampleNewToolbarWithIdentifier demonstrates how to create a Toolbar instance using NewToolbarWithIdentifier.
// Creates a newly allocated toolbar with the specified identifier.
func ExampleNewToolbarWithIdentifier() {
	_ = appkit.NewToolbarWithIdentifier(
		appkit.ToolbarIdentifier{}, // identifier ToolbarIdentifier
	)
	// Output:
}
