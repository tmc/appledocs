// Code generated from Apple documentation for Collaboration. DO NOT EDIT.

package collaboration_test

import (
	"github.com/tmc/appledocs/generated/collaboration"
)

// Suppress unused import errors
var _ = collaboration.NewCBIdentity

// ExampleNewCBIdentityWithNameAuthority demonstrates how to create a CBIdentity instance using NewCBIdentityWithNameAuthority.
// Returns the identity object with the given name from the specified identity authority.
func ExampleNewCBIdentityWithNameAuthority() {
	_ = collaboration.NewCBIdentityWithNameAuthority(
		"name", // name string
		collaboration.CBIdentityAuthority{}, // authority CBIdentityAuthority
	)
	// Output:
}
// ExampleNewCBIdentityWithUUIDStringAuthority demonstrates how to create a CBIdentity instance using NewCBIdentityWithUUIDStringAuthority.
// Returns the identity object with the given UUID from the specified identity authority.
func ExampleNewCBIdentityWithUUIDStringAuthority() {
	_ = collaboration.NewCBIdentityWithUUIDStringAuthority(
		"uuid", // uuid string
		collaboration.CBIdentityAuthority{}, // authority CBIdentityAuthority
	)
	// Output:
}
