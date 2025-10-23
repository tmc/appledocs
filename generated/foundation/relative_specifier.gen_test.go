// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewRelativeSpecifier

// ExampleNewRelativeSpecifierWithCoder demonstrates how to create a RelativeSpecifier instance using NewRelativeSpecifierWithCoder.
func ExampleNewRelativeSpecifierWithCoder() {
	_ = foundation.NewRelativeSpecifierWithCoder(
		foundation.NSCoder{}, // inCoder NSCoder
	)
	// Output:
}
// ExampleNewRelativeSpecifierWithContainerClassDescriptionContainerSpecifierKeyRelativePositionBaseSpecifier demonstrates how to create a RelativeSpecifier instance using NewRelativeSpecifierWithContainerClassDescriptionContainerSpecifierKeyRelativePositionBaseSpecifier.
// Invokes the super class’s   method and initializes the relative position and base specifier to   and  .
func ExampleNewRelativeSpecifierWithContainerClassDescriptionContainerSpecifierKeyRelativePositionBaseSpecifier() {
	_ = foundation.NewRelativeSpecifierWithContainerClassDescriptionContainerSpecifierKeyRelativePositionBaseSpecifier(
		foundation.NSScriptClassDescription{}, // classDesc NSScriptClassDescription
		foundation.NSScriptObjectSpecifier{}, // container NSScriptObjectSpecifier
		"property", // property string
		foundation.RelativePosition{}, // relPos RelativePosition
		foundation.NSScriptObjectSpecifier{}, // baseSpecifier NSScriptObjectSpecifier
	)
	// Output:
}
