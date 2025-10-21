// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewMenuItem

// ExampleNewMenuItemWithCoder demonstrates how to create a MenuItem instance using NewMenuItemWithCoder.
func ExampleNewMenuItemWithCoder() {
	_ = appkit.NewMenuItemWithCoder(
		appkit.Coder{}, // coder Coder
	)
	// Output:
}
// ExampleNewMenuItemWithTitleActionKeyEquivalent demonstrates how to create a MenuItem instance using NewMenuItemWithTitleActionKeyEquivalent.
// Returns an initialized instance of  .
func ExampleNewMenuItemWithTitleActionKeyEquivalent() {
	_ = appkit.NewMenuItemWithTitleActionKeyEquivalent(
		"string", // string string
		0, // selector objc.SEL
		"charCode", // charCode string
	)
	// Output:
}
