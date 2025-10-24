// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewXMLNode

// ExampleNewXMLNode demonstrates how to create a XMLNode instance.
func ExampleNewXMLNode() {
	_ = foundation.NewXMLNode()
	// Output:
}
// ExampleXMLNode_Detach demonstrates using Detach on a XMLNode instance.
// Detaches the receiver from its parent node.
func ExampleXMLNode_Detach() {
	obj := foundation.NewXMLNode()
	obj.Detach()
	// Output:
	}

