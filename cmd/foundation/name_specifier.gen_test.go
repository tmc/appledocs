// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewNameSpecifier

// ExampleNewNameSpecifierWithContainerClassDescriptionContainerSpecifierKeyName demonstrates how to create a NameSpecifier instance using NewNameSpecifierWithContainerClassDescriptionContainerSpecifierKeyName.
// Invokes the super class’s   method and then sets the name instance variable to  .
func ExampleNewNameSpecifierWithContainerClassDescriptionContainerSpecifierKeyName() {
	_ = foundation.NewNameSpecifierWithContainerClassDescriptionContainerSpecifierKeyName(
		foundation.NSScriptClassDescription{}, // classDesc NSScriptClassDescription
		foundation.NSScriptObjectSpecifier{}, // container NSScriptObjectSpecifier
		"property", // property string
		"name", // name string
	)
	// Output:
}
