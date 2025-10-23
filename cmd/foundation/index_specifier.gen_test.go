// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewIndexSpecifier

// ExampleNewIndexSpecifierWithContainerClassDescriptionContainerSpecifierKeyIndex demonstrates how to create a IndexSpecifier instance using NewIndexSpecifierWithContainerClassDescriptionContainerSpecifierKeyIndex.
// Initializes an allocated   object with a class description, container specifier, collection key, and object index.
func ExampleNewIndexSpecifierWithContainerClassDescriptionContainerSpecifierKeyIndex() {
	_ = foundation.NewIndexSpecifierWithContainerClassDescriptionContainerSpecifierKeyIndex(
		foundation.NSScriptClassDescription{}, // classDesc NSScriptClassDescription
		foundation.NSScriptObjectSpecifier{}, // container NSScriptObjectSpecifier
		"property", // property string
		0, // index int
	)
	// Output:
}
