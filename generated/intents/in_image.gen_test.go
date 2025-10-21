// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents_test

import (
	"github.com/tmc/appledocs/generated/intents"
)

// Suppress unused import errors
var _ = intents.NewINImage


// ExampleNewINImageWithCGImage demonstrates how to create a INImage instance using NewINImageWithCGImage.
// Creates an image object from the specified Core Graphics image.
func ExampleNewINImageWithCGImage() {
	_ = intents.NewINImageWithCGImage(
		intents.CGImageRef{}, // imageRef CGImageRef
	)
	// Output:
}




// ExampleNewINImageNamed demonstrates how to create a INImage instance using NewINImageNamed.
// Creates an image object from an image file in the extension’s bundle.
func ExampleNewINImageNamed() {
	_ = intents.NewINImageNamed(
		"name", // name string
	)
	// Output:
}




