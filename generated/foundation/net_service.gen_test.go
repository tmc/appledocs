// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewNetService

// ExampleNewNetServiceWithDomainTypeName demonstrates how to create a NetService instance using NewNetServiceWithDomainTypeName.
// Returns the receiver, initialized as a network service of a given type and sets the initial host information.
func ExampleNewNetServiceWithDomainTypeName() {
	_ = foundation.NewNetServiceWithDomainTypeName(
		"domain", // domain string
		"type", // type string
		"name", // name string
	)
	// Output:
}
