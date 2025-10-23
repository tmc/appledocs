// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewScriptObjectSpecifier

// ExampleNewScriptObjectSpecifierWithCoder demonstrates how to create a ScriptObjectSpecifier instance using NewScriptObjectSpecifierWithCoder.
func ExampleNewScriptObjectSpecifierWithCoder() {
	_ = foundation.NewScriptObjectSpecifierWithCoder(
		foundation.NSCoder{}, // inCoder NSCoder
	)
	// Output:
}
// ExampleNewScriptObjectSpecifierWithContainerClassDescriptionContainerSpecifierKey demonstrates how to create a ScriptObjectSpecifier instance using NewScriptObjectSpecifierWithContainerClassDescriptionContainerSpecifierKey.
// Returns an   object initialized with the given attributes.
func ExampleNewScriptObjectSpecifierWithContainerClassDescriptionContainerSpecifierKey() {
	_ = foundation.NewScriptObjectSpecifierWithContainerClassDescriptionContainerSpecifierKey(
		foundation.NSScriptClassDescription{}, // classDesc NSScriptClassDescription
		foundation.NSScriptObjectSpecifier{}, // container NSScriptObjectSpecifier
		"property", // property string
	)
	// Output:
}
// ExampleNewScriptObjectSpecifierWithContainerSpecifierKey demonstrates how to create a ScriptObjectSpecifier instance using NewScriptObjectSpecifierWithContainerSpecifierKey.
// Returns an   object initialized with a given container specifier  and key.
func ExampleNewScriptObjectSpecifierWithContainerSpecifierKey() {
	_ = foundation.NewScriptObjectSpecifierWithContainerSpecifierKey(
		foundation.NSScriptObjectSpecifier{}, // container NSScriptObjectSpecifier
		"property", // property string
	)
	// Output:
}
// ExampleNewScriptObjectSpecifierWithDescriptor demonstrates how to create a ScriptObjectSpecifier instance using NewScriptObjectSpecifierWithDescriptor.
// Returns a new object specifier for an Apple event descriptor.
func ExampleNewScriptObjectSpecifierWithDescriptor() {
	_ = foundation.NewScriptObjectSpecifierWithDescriptor(
		foundation.NSAppleEventDescriptor{}, // descriptor NSAppleEventDescriptor
	)
	// Output:
}
