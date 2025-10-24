// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewException

// ExampleException_Raise demonstrates using Raise on a Exception instance.
// Raises the receiver, causing program flow to jump to the local exception handler.
func ExampleException_Raise() {
	obj := foundation.NewException()
	obj.Raise()
	// Output:
}
