// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewXMLDTDNode

// ExampleNewXMLDTDNode demonstrates how to create a XMLDTDNode instance.
func ExampleNewXMLDTDNode() {
	_ = foundation.NewXMLDTDNode()
	// Output:
}
// ExampleNewXMLDTDNodeWithKindOptions demonstrates how to create a XMLDTDNode instance using NewXMLDTDNodeWithKindOptions.
func ExampleNewXMLDTDNodeWithKindOptions() {
	_ = foundation.NewXMLDTDNodeWithKindOptions(
		foundation.XMLNodeKind{}, // kind XMLNodeKind
		foundation.XMLNodeOptions{}, // options XMLNodeOptions
	)
	// Output:
}
// ExampleNewXMLDTDNodeWithXMLString demonstrates how to create a XMLDTDNode instance using NewXMLDTDNodeWithXMLString.
// Returns an   object initialized with the DTD declaration in a given string.
func ExampleNewXMLDTDNodeWithXMLString() {
	_ = foundation.NewXMLDTDNodeWithXMLString(
		"string", // string string
	)
	// Output:
}
