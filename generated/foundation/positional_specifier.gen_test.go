// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewPositionalSpecifier

// ExampleNewPositionalSpecifierWithPositionObjectSpecifier demonstrates how to create a PositionalSpecifier instance using NewPositionalSpecifierWithPositionObjectSpecifier.
// Initializes a positional specifier with a given position relative to another given specifier.
func ExampleNewPositionalSpecifierWithPositionObjectSpecifier() {
	_ = foundation.NewPositionalSpecifierWithPositionObjectSpecifier(
		foundation.InsertionPosition{}, // position InsertionPosition
		foundation.NSScriptObjectSpecifier{}, // specifier NSScriptObjectSpecifier
	)
	// Output:
}
