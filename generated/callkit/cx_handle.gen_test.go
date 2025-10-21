// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit_test

import (
	"github.com/tmc/appledocs/generated/callkit"
)

// Suppress unused import errors
var _ = callkit.NewCXHandle

// ExampleNewCXHandleWithTypeValue demonstrates how to create a CXHandle instance using NewCXHandleWithTypeValue.
// Initializes a new handle of a given type with the specified value.
func ExampleNewCXHandleWithTypeValue() {
	_ = callkit.NewCXHandleWithTypeValue(
		callkit.CXHandleType{}, // type CXHandleType
		"value", // value string
	)
	// Output:
}
