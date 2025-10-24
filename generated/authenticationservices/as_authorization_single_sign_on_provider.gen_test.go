// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices_test

import (
	"github.com/tmc/appledocs/generated/authenticationservices"
)

// Suppress unused import errors
var _ = authenticationservices.NewAuthorizationSingleSignOnProvider

// ExampleAuthorizationSingleSignOnProvider_CreateRequest demonstrates using CreateRequest on a AuthorizationSingleSignOnProvider instance.
// Creates a single sign-on (SSO) authorization request.
func ExampleAuthorizationSingleSignOnProvider_CreateRequest() {
	obj := authenticationservices.NewAuthorizationSingleSignOnProvider()
	_ = obj.CreateRequest()
	// Output:
	}

