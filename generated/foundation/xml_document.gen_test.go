// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewXMLDocument

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
