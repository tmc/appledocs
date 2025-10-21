// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices_test

import (
	"github.com/tmc/appledocs/generated/authenticationservices"
)

// Suppress unused import errors
var _ = authenticationservices.NewAuthorizationPublicKeyCredentialPRFRegistrationInput

// ExampleNewAuthorizationPublicKeyCredentialPRFRegistrationInputWithInputValues demonstrates how to create a AuthorizationPublicKeyCredentialPRFRegistrationInput instance using NewAuthorizationPublicKeyCredentialPRFRegistrationInputWithInputValues.
func ExampleNewAuthorizationPublicKeyCredentialPRFRegistrationInputWithInputValues() {
	_ = authenticationservices.NewAuthorizationPublicKeyCredentialPRFRegistrationInputWithInputValues(
		authenticationservices.ASAuthorizationPublicKeyCredentialPRFAssertionInputValues{}, // inputValues ASAuthorizationPublicKeyCredentialPRFAssertionInputValues
	)
	// Output:
}
