// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewToolbarItem

// ExampleNewToolbarItemWithItemIdentifier demonstrates how to create a ToolbarItem instance using NewToolbarItemWithItemIdentifier.
// Creates a toolbar item with the specified identifier.
func ExampleNewToolbarItemWithItemIdentifier() {
	_ = appkit.NewToolbarItemWithItemIdentifier(
		appkit.ToolbarItemIdentifier{}, // itemIdentifier ToolbarItemIdentifier
	)
	// Output:
}
