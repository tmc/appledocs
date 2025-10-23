// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewUUID

// ExampleNewUUID demonstrates how to create a UUID instance.
// Initializes a new UUID with RFC 4122 version 4 random bytes.
func ExampleNewUUID() {
	_ = foundation.NewUUID()
	// Output:
}
// ExampleNewUUIDWithUUIDString demonstrates how to create a UUID instance using NewUUIDWithUUIDString.
// Initializes a new UUID with the formatted string.
func ExampleNewUUIDWithUUIDString() {
	_ = foundation.NewUUIDWithUUIDString(
		"string", // string string
	)
	// Output:
}
