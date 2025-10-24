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
// ExampleNewXMLDTDWithKindOptions demonstrates how to create a XMLDTD instance using NewXMLDTDWithKindOptions.
func ExampleNewXMLDTDWithKindOptions() {
	_ = foundation.NewXMLDTDWithKindOptions(
		foundation.XMLNodeKind /* not a class type */{}, // kind XMLNodeKind /* not a class type */
		foundation.XMLNodeOptions{}, // options XMLNodeOptions
	)
	// Output:
}
