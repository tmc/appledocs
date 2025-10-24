// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewTextLayoutFragment

// ExampleTextLayoutFragment_InvalidateLayout demonstrates using InvalidateLayout on a TextLayoutFragment instance.
// Invalidates any layout information associated with the text layout fragment.
func ExampleTextLayoutFragment_InvalidateLayout() {
	obj := appkit.NewTextLayoutFragment()
	obj.InvalidateLayout()
	// Output:
	}

