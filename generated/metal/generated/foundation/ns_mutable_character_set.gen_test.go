// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewMutableCharacterSet

// ExampleMutableCharacterSet_Invert demonstrates using Invert on a MutableCharacterSet instance.
// Replaces all the characters in the receiver with all the characters it didn’t previously contain.
func ExampleMutableCharacterSet_Invert() {
	obj := foundation.NewMutableCharacterSet()
	obj.Invert()
	// Output:
	}

