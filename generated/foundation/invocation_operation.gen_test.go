// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewInvocationOperation

// ExampleNewInvocationOperationWithInvocation demonstrates how to create a InvocationOperation instance using NewInvocationOperationWithInvocation.
// Returns an   object initialized with the specified invocation object.
func ExampleNewInvocationOperationWithInvocation() {
	_ = foundation.NewInvocationOperationWithInvocation(
		foundation.NSInvocation{}, // inv NSInvocation
	)
	// Output:
}
