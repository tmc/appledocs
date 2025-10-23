// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewValueTransformer

// ExampleNewValueTransformerForName demonstrates how to create a ValueTransformer instance using NewValueTransformerForName.
// Returns the value transformer identified by a given identifier.
func ExampleNewValueTransformerForName() {
	_ = foundation.NewValueTransformerForName(
		foundation.ValueTransformerName{}, // name ValueTransformerName
	)
	// Output:
}
