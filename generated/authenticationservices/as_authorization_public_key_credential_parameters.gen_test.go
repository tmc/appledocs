// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices_test

import (
	"github.com/tmc/appledocs/generated/authenticationservices"
)

// Suppress unused import errors
var _ = authenticationservices.NewAuthorizationPublicKeyCredentialParameters

// ExampleNewAuthorizationPublicKeyCredentialParametersWithAlgorithm demonstrates how to create a AuthorizationPublicKeyCredentialParameters instance using NewAuthorizationPublicKeyCredentialParametersWithAlgorithm.
// Creates the object with an algorithm.
func ExampleNewAuthorizationPublicKeyCredentialParametersWithAlgorithm() {
	_ = authenticationservices.NewAuthorizationPublicKeyCredentialParametersWithAlgorithm(
		authenticationservices.COSEAlgorithmIdentifier /* typedef */{}, // algorithm COSEAlgorithmIdentifier /* typedef */
	)
	// Output:
}
