// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewLocale

// ExampleNewLocaleWithCoder demonstrates how to create a Locale instance using NewLocaleWithCoder.
// Returns a locale initialized from data in the given unarchiver.
func ExampleNewLocaleWithCoder() {
	_ = foundation.NewLocaleWithCoder(
		foundation.NSCoder{}, // coder NSCoder
	)
	// Output:
}
// ExampleNewLocaleWithLocaleIdentifier demonstrates how to create a Locale instance using NewLocaleWithLocaleIdentifier.
// Initializes a locale using a given locale identifier.
func ExampleNewLocaleWithLocaleIdentifier() {
	_ = foundation.NewLocaleWithLocaleIdentifier(
		"string", // string string
	)
	// Output:
}
