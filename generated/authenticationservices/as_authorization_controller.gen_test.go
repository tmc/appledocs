// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices_test

import (
	"github.com/tmc/appledocs/generated/authenticationservices"
)

// Suppress unused import errors
var _ = authenticationservices.NewAuthorizationController

// ExampleNewAuthorizationControllerWithAuthorizationRequests demonstrates how to create a AuthorizationController instance using NewAuthorizationControllerWithAuthorizationRequests.
// Creates a controller from a collection of authorization requests.
func ExampleNewAuthorizationControllerWithAuthorizationRequests() {
	_ = authenticationservices.NewAuthorizationControllerWithAuthorizationRequests(
		[]authenticationservices.AuthorizationRequest{}, // authorizationRequests []AuthorizationRequest
	)
	// Output:
}
// ExampleAuthorizationController_Cancel demonstrates using Cancel on a AuthorizationController instance.
// Cancels any active authorization requests.
func ExampleAuthorizationController_Cancel() {
	obj := authenticationservices.NewAuthorizationController()
	obj.Cancel()
	// Output:
	}

// ExampleAuthorizationController_PerformRequests demonstrates using PerformRequests on a AuthorizationController instance.
// Starts the specified authorization flows during controller initialization.
func ExampleAuthorizationController_PerformRequests() {
	obj := authenticationservices.NewAuthorizationController()
	obj.PerformRequests()
	// Output:
	}

