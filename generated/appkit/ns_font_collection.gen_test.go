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
