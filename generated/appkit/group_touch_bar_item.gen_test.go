// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewGroupTouchBarItem

// ExampleNewGroupTouchBarItemGroupItemWithIdentifierItems demonstrates how to create a GroupTouchBarItem instance using NewGroupTouchBarItemGroupItemWithIdentifierItems.
// Initializes and returns a group item whose bar is constructed from the supplied items.
func ExampleNewGroupTouchBarItemGroupItemWithIdentifierItems() {
	_ = appkit.NewGroupTouchBarItemGroupItemWithIdentifierItems(
		appkit.TouchBarItemIdentifier{}, // identifier TouchBarItemIdentifier
		[]appkit.TouchBarItem{}, // items []TouchBarItem
	)
	// Output:
}
// ExampleNewGroupTouchBarItemGroupItemWithIdentifierItemsAllowedCompressionOptions demonstrates how to create a GroupTouchBarItem instance using NewGroupTouchBarItemGroupItemWithIdentifierItemsAllowedCompressionOptions.
// Initializes and returns a group item whose bar is constructed from the supplied items, and with the specified compression options.
func ExampleNewGroupTouchBarItemGroupItemWithIdentifierItemsAllowedCompressionOptions() {
	_ = appkit.NewGroupTouchBarItemGroupItemWithIdentifierItemsAllowedCompressionOptions(
		appkit.TouchBarItemIdentifier{}, // identifier TouchBarItemIdentifier
		[]appkit.TouchBarItem{}, // items []TouchBarItem
		appkit.NSUserInterfaceCompressionOptions{}, // allowedCompressionOptions NSUserInterfaceCompressionOptions
	)
	// Output:
}
