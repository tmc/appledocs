// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewMutableFontCollection

// ExampleNewMutableFontCollectionWithName demonstrates how to create a MutableFontCollection instance using NewMutableFontCollectionWithName.
// Creates a mutable named font collection object.
func ExampleNewMutableFontCollectionWithName() {
	_ = appkit.NewMutableFontCollectionWithName(
		appkit.FontCollectionName /* not a class type */ {}, // name FontCollectionName /* not a class type */
	)
	// Output:
}

// ExampleNewMutableFontCollectionWithNameVisibility demonstrates how to create a MutableFontCollection instance using NewMutableFontCollectionWithNameVisibility.
// Creates a mutable font collection with the specified name and font visibility.
func ExampleNewMutableFontCollectionWithNameVisibility() {
	_ = appkit.NewMutableFontCollectionWithNameVisibility(
		appkit.FontCollectionName /* not a class type */ {},       // name FontCollectionName /* not a class type */
		appkit.FontCollectionVisibility /* not a class type */ {}, // visibility FontCollectionVisibility /* not a class type */
	)
	// Output:
}
