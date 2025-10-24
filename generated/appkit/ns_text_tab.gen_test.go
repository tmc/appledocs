// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewTextTab

// ExampleNewTextTabWithTypeLocation demonstrates how to create a TextTab instance using NewTextTabWithTypeLocation.
// Initializes a newly allocated text tab with the specified alignment and location.
func ExampleNewTextTabWithTypeLocation() {
	_ = appkit.NewTextTabWithTypeLocation(
		appkit.TextTabType{}, // type TextTabType
		0.0, // loc float64
	)
	// Output:
}
