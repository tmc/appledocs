// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewXMLDTD

// ExampleNewXMLDTD demonstrates how to create a XMLDTD instance.
func ExampleNewXMLDTD() {
	_ = foundation.NewXMLDTD()
	// Output:
}
// ExampleNewXMLDTDWithContentsOfURLOptionsError demonstrates how to create a XMLDTD instance using NewXMLDTDWithContentsOfURLOptionsError.
// Initializes and returns an   object created from the DTD declarations in a URL-referenced source.
func ExampleNewXMLDTDWithContentsOfURLOptionsError() {
	_ = foundation.NewXMLDTDWithContentsOfURLOptionsError(
		foundation.URL{}, // url URL
		foundation.XMLNodeOptions{}, // mask XMLNodeOptions
		foundation.NSError{}, // error NSError
	)
	// Output:
}
// ExampleNewXMLDTDWithDataOptionsError demonstrates how to create a XMLDTD instance using NewXMLDTDWithDataOptionsError.
// Initializes and returns an   object created from the DTD declarations encapsulated in an   object
func ExampleNewXMLDTDWithDataOptionsError() {
	_ = foundation.NewXMLDTDWithDataOptionsError(
		foundation.NSData{}, // data NSData
		foundation.XMLNodeOptions{}, // mask XMLNodeOptions
		foundation.NSError{}, // error NSError
	)
	// Output:
}
// ExampleNewXMLDTDWithKindOptions demonstrates how to create a XMLDTD instance using NewXMLDTDWithKindOptions.
func ExampleNewXMLDTDWithKindOptions() {
	_ = foundation.NewXMLDTDWithKindOptions(
		foundation.XMLNodeKind{}, // kind XMLNodeKind
		foundation.XMLNodeOptions{}, // options XMLNodeOptions
	)
	// Output:
}
