// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import "github.com/ebitengine/purego/objc"

// WebAuthenticationSessionRequestDelegateProtocol is the ASWebAuthenticationSessionRequestDelegate protocol.
//
// Availability:
//   - macOS 10.15+
//
// Use this protocol when registering custom classes that conform to ASWebAuthenticationSessionRequestDelegate.
var WebAuthenticationSessionRequestDelegateProtocol *objc.Protocol

func init() {
	WebAuthenticationSessionRequestDelegateProtocol = objc.GetProtocol("ASWebAuthenticationSessionRequestDelegate")
}

