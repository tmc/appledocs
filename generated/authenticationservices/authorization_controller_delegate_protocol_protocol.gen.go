// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"
)

// PAuthorizationControllerDelegate is the ASAuthorizationControllerDelegate protocol interface.
//
// An interface for providing information about the outcome of an authorization request.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+
//
// See: doc://com.apple.authenticationservices/documentation/AuthenticationServices/ASAuthorizationControllerDelegate
type PAuthorizationControllerDelegate interface {
	// Optional methods
	AuthorizationControllerDidCompleteWithCustomMethod(controller IASAuthorizationController, method AuthorizationCustomMethod /* typedef */)
	HasAuthorizationControllerDidCompleteWithCustomMethod() bool
	AuthorizationControllerDidCompleteWithAuthorization(controller IASAuthorizationController, authorization IASAuthorization)
	HasAuthorizationControllerDidCompleteWithAuthorization() bool
	AuthorizationControllerDidCompleteWithError(controller IASAuthorizationController, error_ objc.IObject /* cross-framework: Error */)
	HasAuthorizationControllerDidCompleteWithError() bool
}

// AuthorizationControllerDelegate is a delegate implementation builder for the PAuthorizationControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type AuthorizationControllerDelegate struct {
	_AuthorizationControllerDidCompleteWithCustomMethod func(controller IASAuthorizationController, method AuthorizationCustomMethod /* typedef */)
	_AuthorizationControllerDidCompleteWithAuthorization func(controller IASAuthorizationController, authorization IASAuthorization)
	_AuthorizationControllerDidCompleteWithError func(controller IASAuthorizationController, error_ objc.IObject /* cross-framework: Error */)
}

// SetAuthorizationControllerDidCompleteWithCustomMethod sets the handler for the AuthorizationControllerDidCompleteWithCustomMethod delegate method.
//
// Informs the delegate when authorization completes, and specifies the custom method the user selected.
func (d *AuthorizationControllerDelegate) SetAuthorizationControllerDidCompleteWithCustomMethod(f func(controller IASAuthorizationController, method AuthorizationCustomMethod /* typedef */)) {
	d._AuthorizationControllerDidCompleteWithCustomMethod = f
}

// SetAuthorizationControllerDidCompleteWithAuthorization sets the handler for the AuthorizationControllerDidCompleteWithAuthorization delegate method.
//
// Tells the delegate when authorization completes successfully.
func (d *AuthorizationControllerDelegate) SetAuthorizationControllerDidCompleteWithAuthorization(f func(controller IASAuthorizationController, authorization IASAuthorization)) {
	d._AuthorizationControllerDidCompleteWithAuthorization = f
}

// SetAuthorizationControllerDidCompleteWithError sets the handler for the AuthorizationControllerDidCompleteWithError delegate method.
//
// Tells the delegate when authorization fails, and provides an error explaining why.
func (d *AuthorizationControllerDelegate) SetAuthorizationControllerDidCompleteWithError(f func(controller IASAuthorizationController, error_ objc.IObject /* cross-framework: Error */)) {
	d._AuthorizationControllerDidCompleteWithError = f
}

// AuthorizationControllerDidCompleteWithCustomMethod implements the PAuthorizationControllerDelegate interface.
func (d *AuthorizationControllerDelegate) AuthorizationControllerDidCompleteWithCustomMethod(controller IASAuthorizationController, method AuthorizationCustomMethod /* typedef */) {
	if d._AuthorizationControllerDidCompleteWithCustomMethod != nil {
		d._AuthorizationControllerDidCompleteWithCustomMethod(controller, method)
	}
}

// HasAuthorizationControllerDidCompleteWithCustomMethod returns true if a handler for AuthorizationControllerDidCompleteWithCustomMethod has been set.
func (d *AuthorizationControllerDelegate) HasAuthorizationControllerDidCompleteWithCustomMethod() bool {
	return d._AuthorizationControllerDidCompleteWithCustomMethod != nil
}

// AuthorizationControllerDidCompleteWithAuthorization implements the PAuthorizationControllerDelegate interface.
func (d *AuthorizationControllerDelegate) AuthorizationControllerDidCompleteWithAuthorization(controller IASAuthorizationController, authorization IASAuthorization) {
	if d._AuthorizationControllerDidCompleteWithAuthorization != nil {
		d._AuthorizationControllerDidCompleteWithAuthorization(controller, authorization)
	}
}

// HasAuthorizationControllerDidCompleteWithAuthorization returns true if a handler for AuthorizationControllerDidCompleteWithAuthorization has been set.
func (d *AuthorizationControllerDelegate) HasAuthorizationControllerDidCompleteWithAuthorization() bool {
	return d._AuthorizationControllerDidCompleteWithAuthorization != nil
}

// AuthorizationControllerDidCompleteWithError implements the PAuthorizationControllerDelegate interface.
func (d *AuthorizationControllerDelegate) AuthorizationControllerDidCompleteWithError(controller IASAuthorizationController, error_ objc.IObject /* cross-framework: Error */) {
	if d._AuthorizationControllerDidCompleteWithError != nil {
		d._AuthorizationControllerDidCompleteWithError(controller, error_)
	}
}

// HasAuthorizationControllerDidCompleteWithError returns true if a handler for AuthorizationControllerDidCompleteWithError has been set.
func (d *AuthorizationControllerDelegate) HasAuthorizationControllerDidCompleteWithError() bool {
	return d._AuthorizationControllerDidCompleteWithError != nil
}
