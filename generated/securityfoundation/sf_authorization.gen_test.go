// Code generated from Apple documentation for SecurityFoundation. DO NOT EDIT.

package securityfoundation_test

import (
	"github.com/tmc/appledocs/generated/securityfoundation"
)

// Suppress unused import errors
var _ = securityfoundation.NewSFAuthorization

// ExampleNewSFAuthorization demonstrates how to create a SFAuthorization instance.
// Initializes an authorization object with default environment, flags, and rights.
func ExampleNewSFAuthorization() {
	_ = securityfoundation.NewSFAuthorization()
	// Output:
}
// ExampleSFAuthorization_AuthorizationRef demonstrates using AuthorizationRef on a SFAuthorization instance.
// Returns the authorization reference for this object.
func ExampleSFAuthorization_AuthorizationRef() {
	obj := securityfoundation.NewSFAuthorization()
	_ = obj.AuthorizationRef()
	// Output:
	}

// ExampleSFAuthorization_InvalidateCredentials demonstrates using InvalidateCredentials on a SFAuthorization instance.
// Prevents any rights that were obtained by this object from being preserved.
func ExampleSFAuthorization_InvalidateCredentials() {
	obj := securityfoundation.NewSFAuthorization()
	obj.InvalidateCredentials()
	// Output:
	}




