// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewMutableFontCollection

// ExampleNewMutableFontCollectionWithDescriptors demonstrates how to create a MutableFontCollection instance using NewMutableFontCollectionWithDescriptors.
// Creates a mutable font collection containing the fonts that match the specified font descriptors.
func ExampleNewMutableFontCollectionWithDescriptors() {
	_ = appkit.NewMutableFontCollectionWithDescriptors(
		[]appkit.FontDescriptor{}, // queryDescriptors []FontDescriptor
	)
	// Output:
}
// ExampleNewMutableFontCollectionWithName demonstrates how to create a MutableFontCollection instance using NewMutableFontCollectionWithName.
// Creates a mutable named font collection object.
func ExampleNewMutableFontCollectionWithName() {
	_ = appkit.NewMutableFontCollectionWithName(
		appkit.FontCollectionName /* typedef */{}, // name FontCollectionName /* typedef */
	)
	// Output:
}
// ExampleNewMutableFontCollectionWithNameVisibility demonstrates how to create a MutableFontCollection instance using NewMutableFontCollectionWithNameVisibility.
// Creates a mutable font collection with the specified name and font visibility.
func ExampleNewMutableFontCollectionWithNameVisibility() {
	_ = appkit.NewMutableFontCollectionWithNameVisibility(
		appkit.FontCollectionName /* typedef */{}, // name FontCollectionName /* typedef */
		appkit.FontCollectionVisibility{}, // visibility FontCollectionVisibility
	)
	// Output:
}
