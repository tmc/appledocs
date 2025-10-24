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
