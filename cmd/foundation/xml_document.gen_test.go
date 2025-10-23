// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewXMLDocument

// ExampleNewXMLDocumentWithContentsOfURLOptionsError demonstrates how to create a XMLDocument instance using NewXMLDocumentWithContentsOfURLOptionsError.
// Initializes and returns an NSXMLDocument object created from the XML or HTML contents of a URL-referenced source
func ExampleNewXMLDocumentWithContentsOfURLOptionsError() {
	_ = foundation.NewXMLDocumentWithContentsOfURLOptionsError(
		foundation.URL{}, // url URL
		foundation.XMLNodeOptions{}, // mask XMLNodeOptions
		foundation.NSError{}, // error NSError
	)
	// Output:
}
// ExampleNewXMLDocumentWithDataOptionsError demonstrates how to create a XMLDocument instance using NewXMLDocumentWithDataOptionsError.
// Initializes and returns an   object created from an   object.
func ExampleNewXMLDocumentWithDataOptionsError() {
	_ = foundation.NewXMLDocumentWithDataOptionsError(
		foundation.NSData{}, // data NSData
		foundation.XMLNodeOptions{}, // mask XMLNodeOptions
		foundation.NSError{}, // error NSError
	)
	// Output:
}
// ExampleNewXMLDocumentWithRootElement demonstrates how to create a XMLDocument instance using NewXMLDocumentWithRootElement.
// Returns an   object initialized with a single child, the root element.
func ExampleNewXMLDocumentWithRootElement() {
	_ = foundation.NewXMLDocumentWithRootElement(
		foundation.NSXMLElement{}, // element NSXMLElement
	)
	// Output:
}
// ExampleNewXMLDocumentWithXMLStringOptionsError demonstrates how to create a XMLDocument instance using NewXMLDocumentWithXMLStringOptionsError.
// Initializes and returns an   object created from a string containing XML markup text.
func ExampleNewXMLDocumentWithXMLStringOptionsError() {
	_ = foundation.NewXMLDocumentWithXMLStringOptionsError(
		"string", // string string
		foundation.XMLNodeOptions{}, // mask XMLNodeOptions
		foundation.NSError{}, // error NSError
	)
	// Output:
}
