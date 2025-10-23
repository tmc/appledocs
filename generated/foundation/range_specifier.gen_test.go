// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewRangeSpecifier

// ExampleNewRangeSpecifierWithCoder demonstrates how to create a RangeSpecifier instance using NewRangeSpecifierWithCoder.
func ExampleNewRangeSpecifierWithCoder() {
	_ = foundation.NewRangeSpecifierWithCoder(
		foundation.NSCoder{}, // inCoder NSCoder
	)
	// Output:
}
// ExampleNewRangeSpecifierWithContainerClassDescriptionContainerSpecifierKeyStartSpecifierEndSpecifier demonstrates how to create a RangeSpecifier instance using NewRangeSpecifierWithContainerClassDescriptionContainerSpecifierKeyStartSpecifierEndSpecifier.
// Returns a range specifier initialized with the given properties.
func ExampleNewRangeSpecifierWithContainerClassDescriptionContainerSpecifierKeyStartSpecifierEndSpecifier() {
	_ = foundation.NewRangeSpecifierWithContainerClassDescriptionContainerSpecifierKeyStartSpecifierEndSpecifier(
		foundation.NSScriptClassDescription{}, // classDesc NSScriptClassDescription
		foundation.NSScriptObjectSpecifier{}, // container NSScriptObjectSpecifier
		"property", // property string
		foundation.NSScriptObjectSpecifier{}, // startSpec NSScriptObjectSpecifier
		foundation.NSScriptObjectSpecifier{}, // endSpec NSScriptObjectSpecifier
	)
	// Output:
}
