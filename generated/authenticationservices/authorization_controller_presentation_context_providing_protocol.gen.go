// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import "github.com/ebitengine/purego/objc"

// AuthorizationControllerPresentationContextProvidingProtocol is the ASAuthorizationControllerPresentationContextProviding protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to ASAuthorizationControllerPresentationContextProviding.
var AuthorizationControllerPresentationContextProvidingProtocol *objc.Protocol

func init() {
	AuthorizationControllerPresentationContextProvidingProtocol = objc.GetProtocol("ASAuthorizationControllerPresentationContextProviding")
}

