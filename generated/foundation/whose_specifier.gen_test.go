// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewWhoseSpecifier

// ExampleNewWhoseSpecifierWithCoder demonstrates how to create a WhoseSpecifier instance using NewWhoseSpecifierWithCoder.
func ExampleNewWhoseSpecifierWithCoder() {
	_ = foundation.NewWhoseSpecifierWithCoder(
		foundation.NSCoder{}, // inCoder NSCoder
	)
	// Output:
}
// ExampleNewWhoseSpecifierWithContainerClassDescriptionContainerSpecifierKeyTest demonstrates how to create a WhoseSpecifier instance using NewWhoseSpecifierWithContainerClassDescriptionContainerSpecifierKeyTest.
// Returns an   object initialized with the given attributes.
func ExampleNewWhoseSpecifierWithContainerClassDescriptionContainerSpecifierKeyTest() {
	_ = foundation.NewWhoseSpecifierWithContainerClassDescriptionContainerSpecifierKeyTest(
		foundation.NSScriptClassDescription{}, // classDesc NSScriptClassDescription
		foundation.NSScriptObjectSpecifier{}, // container NSScriptObjectSpecifier
		"property", // property string
		foundation.NSScriptWhoseTest{}, // test NSScriptWhoseTest
	)
	// Output:
}
