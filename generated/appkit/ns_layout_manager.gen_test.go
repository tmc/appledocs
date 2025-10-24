// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewLayoutManager

// ExampleNewLayoutManager demonstrates how to create a LayoutManager instance.
// Initializes a newly created layout manager object.
func ExampleNewLayoutManager() {
	_ = appkit.NewLayoutManager()
	// Output:
}
// ExampleLayoutManager_FirstUnlaidCharacterIndex demonstrates using FirstUnlaidCharacterIndex on a LayoutManager instance.
// Returns the index for the first character in the layout manager that isn’t in the layout.
func ExampleLayoutManager_FirstUnlaidCharacterIndex() {
	obj := appkit.NewLayoutManager()
	_ = obj.FirstUnlaidCharacterIndex()
	// Output:
	}

// ExampleLayoutManager_FirstUnlaidGlyphIndex demonstrates using FirstUnlaidGlyphIndex on a LayoutManager instance.
// Returns the index for the first glyph in the layout manager that isn’t in the layout.
func ExampleLayoutManager_FirstUnlaidGlyphIndex() {
	obj := appkit.NewLayoutManager()
	_ = obj.FirstUnlaidGlyphIndex()
	// Output:
	}

