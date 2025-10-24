// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewAppearance

// ExampleNewAppearanceNamed demonstrates how to create a Appearance instance using NewAppearanceNamed.
// Creates an appearance object based on the name of one of the standard system appearances.
func ExampleNewAppearanceNamed() {
	_ = appkit.NewAppearanceNamed(
		appkit.AppearanceName /* typedef */{}, // name AppearanceName /* typedef */
	)
	// Output:
}
