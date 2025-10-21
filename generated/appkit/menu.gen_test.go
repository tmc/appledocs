// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewMenu

// ExampleNewMenuWithCoder demonstrates how to create a Menu instance using NewMenuWithCoder.
func ExampleNewMenuWithCoder() {
	_ = appkit.NewMenuWithCoder(
		appkit.Coder{}, // coder Coder
	)
	// Output:
}
// ExampleNewMenuWithTitle demonstrates how to create a Menu instance using NewMenuWithTitle.
// Initializes and returns a menu having the specified title and with autoenabling of menu items turned on.
func ExampleNewMenuWithTitle() {
	_ = appkit.NewMenuWithTitle(
		"title", // title string
	)
	// Output:
}
