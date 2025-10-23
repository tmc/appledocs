// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewXMLNode

// ExampleNewXMLNodeWithKind demonstrates how to create a XMLNode instance using NewXMLNodeWithKind.
// Returns an   instance initialized with the constant indicating node kind.
func ExampleNewXMLNodeWithKind() {
	_ = foundation.NewXMLNodeWithKind(
		foundation.XMLNodeKind{}, // kind XMLNodeKind
	)
	// Output:
}
// ExampleNewXMLNodeWithKindOptions demonstrates how to create a XMLNode instance using NewXMLNodeWithKindOptions.
// Returns an   instance initialized with the constant indicating node kind and one or more initialization options.
func ExampleNewXMLNodeWithKindOptions() {
	_ = foundation.NewXMLNodeWithKindOptions(
		foundation.XMLNodeKind{}, // kind XMLNodeKind
		foundation.XMLNodeOptions{}, // options XMLNodeOptions
	)
	// Output:
}
