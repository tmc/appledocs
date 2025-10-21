// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewXMLElement

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


