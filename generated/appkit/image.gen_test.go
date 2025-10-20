// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewImage

// ExampleNewImageWithSymbolNameVariableValue demonstrates how to create a Image instance using NewImageWithSymbolNameVariableValue.
// Creates a symbol image with the symbol name and variable value you specify.
func ExampleNewImageWithSymbolNameVariableValue() {
	_ = appkit.NewImageWithSymbolNameVariableValue(
		"name", // name string
		0.0,    // value float64
	)
	// Output:
}

// ExampleNewImageWithSystemSymbolNameAccessibilityDescription demonstrates how to create a Image instance using NewImageWithSystemSymbolNameAccessibilityDescription.
// Creates a symbol image with the system symbol name and accessibility description you specify.
func ExampleNewImageWithSystemSymbolNameAccessibilityDescription() {
	_ = appkit.NewImageWithSystemSymbolNameAccessibilityDescription(
		"name",        // name string
		"description", // description string
	)
	// Output:
}
