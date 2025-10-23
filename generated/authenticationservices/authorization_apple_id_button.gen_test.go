// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices_test

import (
	"github.com/tmc/appledocs/generated/authenticationservices"
)

// Suppress unused import errors
var _ = authenticationservices.NewAuthorizationAppleIDButton

// ExampleNewAuthorizationAppleIDButtonWithAuthorizationButtonTypeAuthorizationButtonStyle demonstrates how to create a AuthorizationAppleIDButton instance using NewAuthorizationAppleIDButtonWithAuthorizationButtonTypeAuthorizationButtonStyle.
// Creates a new Sign In with Apple authorization button with the given type and style.
func ExampleNewAuthorizationAppleIDButtonWithAuthorizationButtonTypeAuthorizationButtonStyle() {
	_ = authenticationservices.NewAuthorizationAppleIDButtonWithAuthorizationButtonTypeAuthorizationButtonStyle(
		authenticationservices.ASAuthorizationAppleIDButtonType{}, // type ASAuthorizationAppleIDButtonType
		authenticationservices.ASAuthorizationAppleIDButtonStyle{}, // style ASAuthorizationAppleIDButtonStyle
	)
	// Output:
}
// ExampleNewAuthorizationAppleIDButtonWithTypeStyle demonstrates how to create a AuthorizationAppleIDButton instance using NewAuthorizationAppleIDButtonWithTypeStyle.
// Creates a new Sign In with Apple authorization button with the given type and style.
func ExampleNewAuthorizationAppleIDButtonWithTypeStyle() {
	_ = authenticationservices.NewAuthorizationAppleIDButtonWithTypeStyle(
		authenticationservices.ASAuthorizationAppleIDButtonType{}, // type ASAuthorizationAppleIDButtonType
		authenticationservices.ASAuthorizationAppleIDButtonStyle{}, // style ASAuthorizationAppleIDButtonStyle
	)
	// Output:
}
