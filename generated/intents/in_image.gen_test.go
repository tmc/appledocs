// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents_test

import (
	"github.com/tmc/appledocs/generated/intents"
)

// Suppress unused import errors
var _ = intents.NewINImage

// ExampleNewINImageNamed demonstrates how to create a INImage instance using NewINImageNamed.
// Creates an image object from an image file in the extension’s bundle.
func ExampleNewINImageNamed() {
	_ = intents.NewINImageNamed(
		"name", // name string
	)
	// Output:
}
// ExampleNewINImageWithUIImage demonstrates how to create a INImage instance using NewINImageWithUIImage.
// Creates an image object from the specified UIKit image.
func ExampleNewINImageWithUIImage() {
	_ = intents.NewINImageWithUIImage(
		intents.Image{}, // image Image
	)
	// Output:
}
