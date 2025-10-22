// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization_test

import (
	"github.com/tmc/appledocs/generated/virtualization"
)

// Suppress unused import errors
var _ = virtualization.NewVZMACAddress

// ExampleNewVZMACAddressWithString demonstrates how to create a VZMACAddress instance using NewVZMACAddressWithString.
// Creates a MAC address object from a specially formatted string.
func ExampleNewVZMACAddressWithString() {
	_ = virtualization.NewVZMACAddressWithString(
		"string", // string string
	)
	// Output:
}
