// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewFontCollection

// ExampleNewFontCollectionWithDescriptors demonstrates how to create a FontCollection instance using NewFontCollectionWithDescriptors.
// Returns a font collection matching the given descriptors.
func ExampleNewFontCollectionWithDescriptors() {
	_ = appkit.NewFontCollectionWithDescriptors(
		[]appkit.FontDescriptor{}, // queryDescriptors []FontDescriptor
	)
	// Output:
}
// ExampleNewFontCollectionWithName demonstrates how to create a FontCollection instance using NewFontCollectionWithName.
// Creates a named font collection object.
func ExampleNewFontCollectionWithName() {
	_ = appkit.NewFontCollectionWithName(
		appkit.FontCollectionName{}, // name FontCollectionName
	)
	// Output:
}
// ExampleNewFontCollectionWithNameVisibility demonstrates how to create a FontCollection instance using NewFontCollectionWithNameVisibility.
// Creates a font collection with the specified name and font visibility.
func ExampleNewFontCollectionWithNameVisibility() {
	_ = appkit.NewFontCollectionWithNameVisibility(
		appkit.FontCollectionName{}, // name FontCollectionName
		appkit.FontCollectionVisibility{}, // visibility FontCollectionVisibility
	)
	// Output:
}
