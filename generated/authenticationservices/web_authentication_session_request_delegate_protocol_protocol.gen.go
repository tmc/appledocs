// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"

	"github.com/tmc/appledocs/generated/foundation"
)

// PWebAuthenticationSessionRequestDelegate is the ASWebAuthenticationSessionRequestDelegate protocol interface.
//
// An interface through which the session request can inform its delegate, which is typically a browser, about the outcome of the authentication attempt.
//
// Availability:
//   - macOS 10.15+
//
// See: doc://com.apple.authenticationservices/documentation/AuthenticationServices/ASWebAuthenticationSessionRequestDelegate
type PWebAuthenticationSessionRequestDelegate interface {
	// Optional methods
	AuthenticationSessionRequestDidCancelWithError(authenticationSessionRequest IASWebAuthenticationSessionRequest, error_ objc.IObject /* cross-framework: Error */)
	HasAuthenticationSessionRequestDidCancelWithError() bool
	AuthenticationSessionRequestDidCompleteWithCallbackURL(authenticationSessionRequest IASWebAuthenticationSessionRequest, callbackURL objc.IObject /* cross-framework: NSURL */)
	HasAuthenticationSessionRequestDidCompleteWithCallbackURL() bool
}

// WebAuthenticationSessionRequestDelegate is a delegate implementation builder for the PWebAuthenticationSessionRequestDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type WebAuthenticationSessionRequestDelegate struct {
	_AuthenticationSessionRequestDidCancelWithError func(authenticationSessionRequest IASWebAuthenticationSessionRequest, error_ objc.IObject /* cross-framework: Error */)
	_AuthenticationSessionRequestDidCompleteWithCallbackURL func(authenticationSessionRequest IASWebAuthenticationSessionRequest, callbackURL objc.IObject /* cross-framework: NSURL */)
}

// SetAuthenticationSessionRequestDidCancelWithError sets the handler for the AuthenticationSessionRequestDidCancelWithError delegate method.
//
// Tells the delegate, typically a browser, that the authentication was canceled.
func (d *WebAuthenticationSessionRequestDelegate) SetAuthenticationSessionRequestDidCancelWithError(f func(authenticationSessionRequest IASWebAuthenticationSessionRequest, error_ objc.IObject /* cross-framework: Error */)) {
	d._AuthenticationSessionRequestDidCancelWithError = f
}

// SetAuthenticationSessionRequestDidCompleteWithCallbackURL sets the handler for the AuthenticationSessionRequestDidCompleteWithCallbackURL delegate method.
//
// Tells the delegate, typically a browser, that the authentication completed successfully.
func (d *WebAuthenticationSessionRequestDelegate) SetAuthenticationSessionRequestDidCompleteWithCallbackURL(f func(authenticationSessionRequest IASWebAuthenticationSessionRequest, callbackURL objc.IObject /* cross-framework: NSURL */)) {
	d._AuthenticationSessionRequestDidCompleteWithCallbackURL = f
}

// AuthenticationSessionRequestDidCancelWithError implements the PWebAuthenticationSessionRequestDelegate interface.
func (d *WebAuthenticationSessionRequestDelegate) AuthenticationSessionRequestDidCancelWithError(authenticationSessionRequest IASWebAuthenticationSessionRequest, error_ objc.IObject /* cross-framework: Error */) {
	if d._AuthenticationSessionRequestDidCancelWithError != nil {
		d._AuthenticationSessionRequestDidCancelWithError(authenticationSessionRequest, error_)
	}
}

// HasAuthenticationSessionRequestDidCancelWithError returns true if a handler for AuthenticationSessionRequestDidCancelWithError has been set.
func (d *WebAuthenticationSessionRequestDelegate) HasAuthenticationSessionRequestDidCancelWithError() bool {
	return d._AuthenticationSessionRequestDidCancelWithError != nil
}

// AuthenticationSessionRequestDidCompleteWithCallbackURL implements the PWebAuthenticationSessionRequestDelegate interface.
func (d *WebAuthenticationSessionRequestDelegate) AuthenticationSessionRequestDidCompleteWithCallbackURL(authenticationSessionRequest IASWebAuthenticationSessionRequest, callbackURL objc.IObject /* cross-framework: NSURL */) {
	if d._AuthenticationSessionRequestDidCompleteWithCallbackURL != nil {
		d._AuthenticationSessionRequestDidCompleteWithCallbackURL(authenticationSessionRequest, callbackURL)
	}
}

// HasAuthenticationSessionRequestDidCompleteWithCallbackURL returns true if a handler for AuthenticationSessionRequestDidCompleteWithCallbackURL has been set.
func (d *WebAuthenticationSessionRequestDelegate) HasAuthenticationSessionRequestDidCompleteWithCallbackURL() bool {
	return d._AuthenticationSessionRequestDidCompleteWithCallbackURL != nil
}
