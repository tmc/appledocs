// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices_test

import (
	"github.com/tmc/appledocs/generated/authenticationservices"
)

// Suppress unused import errors
var _ = authenticationservices.NewPasskeyRegistrationCredentialExtensionOutput

// ExampleNewPasskeyRegistrationCredentialExtensionOutputWithLargeBlobOutput demonstrates how to create a PasskeyRegistrationCredentialExtensionOutput instance using NewPasskeyRegistrationCredentialExtensionOutputWithLargeBlobOutput.
func ExampleNewPasskeyRegistrationCredentialExtensionOutputWithLargeBlobOutput() {
	_ = authenticationservices.NewPasskeyRegistrationCredentialExtensionOutputWithLargeBlobOutput(
		authenticationservices.ASAuthorizationPublicKeyCredentialLargeBlobRegistrationOutput{}, // largeBlob ASAuthorizationPublicKeyCredentialLargeBlobRegistrationOutput
	)
	// Output:
}
