// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewXMLElement

// ExampleNewXMLElementWithKindOptions demonstrates how to create a XMLElement instance using NewXMLElementWithKindOptions.
func ExampleNewXMLElementWithKindOptions() {
	_ = foundation.NewXMLElementWithKindOptions(
		foundation.XMLNodeKind{}, // kind XMLNodeKind
		foundation.XMLNodeOptions{}, // options XMLNodeOptions
	)
	// Output:
}
// ExampleNewXMLElementWithName demonstrates how to create a XMLElement instance using NewXMLElementWithName.
// Returns an   object initialized with the specified name.
func ExampleNewXMLElementWithName() {
	_ = foundation.NewXMLElementWithName(
		"name", // name string
	)
	// Output:
}
// ExampleNewXMLElementWithNameStringValue demonstrates how to create a XMLElement instance using NewXMLElementWithNameStringValue.
// Returns an   object initialized with a specified name and a single text-node child containing a specified value.
func ExampleNewXMLElementWithNameStringValue() {
	_ = foundation.NewXMLElementWithNameStringValue(
		"name", // name string
		"string", // string string
	)
	// Output:
}
// ExampleNewXMLElementWithNameURI demonstrates how to create a XMLElement instance using NewXMLElementWithNameURI.
// Returns an   object initialized with the specified name and URI.
func ExampleNewXMLElementWithNameURI() {
	_ = foundation.NewXMLElementWithNameURI(
		"name", // name string
		"URI", // URI string
	)
	// Output:
}
