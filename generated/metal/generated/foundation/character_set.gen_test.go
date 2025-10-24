// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewCharacterSet

// ExampleNewCharacterSetWithRange demonstrates how to create a CharacterSet instance using NewCharacterSetWithRange.
// Returns a character set containing characters with Unicode values in a given range.
func ExampleNewCharacterSetWithRange() {
	_ = foundation.NewCharacterSetWithRange(
		foundation.NSRange /* not a class type */{}, // aRange NSRange /* not a class type */
	)
	// Output:
}
