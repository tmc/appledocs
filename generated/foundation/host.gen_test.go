// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewHost


// ExampleNewHostWithAddress demonstrates how to create a Host instance using NewHostWithAddress.
// Returns the   with the Internet address  .
func ExampleNewHostWithAddress() {
	_ = foundation.NewHostWithAddress(
		"address", // address string
	)
	// Output:
}


// ExampleNewHostWithName demonstrates how to create a Host instance using NewHostWithName.
// Returns a host with a specific name.
func ExampleNewHostWithName() {
	_ = foundation.NewHostWithName(
		"name", // name string
	)
	// Output:
}


