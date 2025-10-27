// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewGroupTouchBarItem

// ExampleNewGroupTouchBarItemAlertStyleGroupItemWithIdentifier demonstrates how to create a GroupTouchBarItem instance using NewGroupTouchBarItemAlertStyleGroupItemWithIdentifier.
// Initializes and returns a group item configured to match system alerts.
func ExampleNewGroupTouchBarItemAlertStyleGroupItemWithIdentifier() {
	_ = appkit.NewGroupTouchBarItemAlertStyleGroupItemWithIdentifier(
		appkit.TouchBarItemIdentifier{}, // identifier TouchBarItemIdentifier
	)
	// Output:
}
// ExampleNewGroupTouchBarItemGroupItemWithIdentifierItems demonstrates how to create a GroupTouchBarItem instance using NewGroupTouchBarItemGroupItemWithIdentifierItems.
// Initializes and returns a group item whose bar is constructed from the supplied items.
func ExampleNewGroupTouchBarItemGroupItemWithIdentifierItems() {
	_ = appkit.NewGroupTouchBarItemGroupItemWithIdentifierItems(
		appkit.TouchBarItemIdentifier{}, // identifier TouchBarItemIdentifier
		[]appkit.TouchBarItem{}, // items []TouchBarItem
	)
	// Output:
}
