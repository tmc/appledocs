// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices_test

import (
	"github.com/tmc/appledocs/generated/authenticationservices"
)

// Suppress unused import errors
var _ = authenticationservices.NewWebAuthenticationSession

// ExampleWebAuthenticationSession_Start demonstrates using Start on a WebAuthenticationSession instance.
// Starts a web authentication session.
func ExampleWebAuthenticationSession_Start() {
	obj := authenticationservices.NewWebAuthenticationSession()
	_ = obj.Start()
	// Output:
	}

