// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewMenuItemBadge

// ExampleNewMenuItemBadgeWithCount demonstrates how to create a MenuItemBadge instance using NewMenuItemBadgeWithCount.
// Creates a badge with a count and an empty string.
func ExampleNewMenuItemBadgeWithCount() {
	_ = appkit.NewMenuItemBadgeWithCount(
		0, // itemCount int
	)
	// Output:
}

// ExampleNewMenuItemBadgeWithCountType demonstrates how to create a MenuItemBadge instance using NewMenuItemBadgeWithCountType.
func ExampleNewMenuItemBadgeWithCountType() {
	_ = appkit.NewMenuItemBadgeWithCountType(
		0,                          // itemCount int
		appkit.MenuItemBadgeType{}, // type MenuItemBadgeType
	)
	// Output:
}

// ExampleNewMenuItemBadgeWithString demonstrates how to create a MenuItemBadge instance using NewMenuItemBadgeWithString.
// Creates a badge with the provided custom string.
func ExampleNewMenuItemBadgeWithString() {
	_ = appkit.NewMenuItemBadgeWithString(
		"string", // string string
	)
	// Output:
}
