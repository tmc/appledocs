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
