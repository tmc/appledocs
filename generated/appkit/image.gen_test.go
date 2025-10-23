// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewImage

// ExampleNewImage demonstrates how to create a Image instance.
func ExampleNewImage() {
	_ = appkit.NewImage()
	// Output:
}
// ExampleNewImageByReferencingFile demonstrates how to create a Image instance using NewImageByReferencingFile.
// Initializes and returns an image object using the specified file.
func ExampleNewImageByReferencingFile() {
	_ = appkit.NewImageByReferencingFile(
		"fileName", // fileName string
	)
	// Output:
}
// ExampleNewImageNamed demonstrates how to create a Image instance using NewImageNamed.
// Returns the image object associated with the specified name.
func ExampleNewImageNamed() {
	_ = appkit.NewImageNamed(
		appkit.ImageName{}, // name ImageName
	)
	// Output:
}
// ExampleNewImageWithCoder demonstrates how to create a Image instance using NewImageWithCoder.
// Initializes and returns an image object from data in an unarchiver.
func ExampleNewImageWithCoder() {
	_ = appkit.NewImageWithCoder(
		appkit.Coder{}, // coder Coder
	)
	// Output:
}
// ExampleNewImageWithContentsOfFile demonstrates how to create a Image instance using NewImageWithContentsOfFile.
// Initializes and returns an image object with the contents of the specified file.
func ExampleNewImageWithContentsOfFile() {
	_ = appkit.NewImageWithContentsOfFile(
		"fileName", // fileName string
	)
	// Output:
}
// ExampleNewImageWithPasteboard demonstrates how to create a Image instance using NewImageWithPasteboard.
// Initializes and returns an image object with data from the specified pasteboard.
func ExampleNewImageWithPasteboard() {
	_ = appkit.NewImageWithPasteboard(
		appkit.NSPasteboard{}, // pasteboard NSPasteboard
	)
	// Output:
}
// ExampleNewImageWithSymbolNameBundleVariableValue demonstrates how to create a Image instance using NewImageWithSymbolNameBundleVariableValue.
func ExampleNewImageWithSymbolNameBundleVariableValue() {
	_ = appkit.NewImageWithSymbolNameBundleVariableValue(
		"name", // name string
		appkit.Bundle{}, // bundle Bundle
		0.0, // value float64
	)
	// Output:
}
// ExampleNewImageWithSymbolNameVariableValue demonstrates how to create a Image instance using NewImageWithSymbolNameVariableValue.
// Creates a symbol image with the symbol name and variable value you specify.
func ExampleNewImageWithSymbolNameVariableValue() {
	_ = appkit.NewImageWithSymbolNameVariableValue(
		"name", // name string
		0.0, // value float64
	)
	// Output:
}
// ExampleNewImageWithSystemSymbolNameAccessibilityDescription demonstrates how to create a Image instance using NewImageWithSystemSymbolNameAccessibilityDescription.
// Creates a symbol image with the system symbol name and accessibility description you specify.
func ExampleNewImageWithSystemSymbolNameAccessibilityDescription() {
	_ = appkit.NewImageWithSystemSymbolNameAccessibilityDescription(
		"name", // name string
		"description", // description string
	)
	// Output:
}
// ExampleNewImageWithSystemSymbolNameVariableValueAccessibilityDescription demonstrates how to create a Image instance using NewImageWithSystemSymbolNameVariableValueAccessibilityDescription.
// Creates a symbol image with the system symbol name and variable value you specify.
func ExampleNewImageWithSystemSymbolNameVariableValueAccessibilityDescription() {
	_ = appkit.NewImageWithSystemSymbolNameVariableValueAccessibilityDescription(
		"name", // name string
		0.0, // value float64
		"description", // description string
	)
	// Output:
}
