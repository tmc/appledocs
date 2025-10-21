// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import "github.com/ebitengine/purego/objc"

// AuthorizationProviderExtensionAuthorizationRequestHandlerProtocol is the ASAuthorizationProviderExtensionAuthorizationRequestHandler protocol.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to ASAuthorizationProviderExtensionAuthorizationRequestHandler.
var AuthorizationProviderExtensionAuthorizationRequestHandlerProtocol *objc.Protocol

func init() {
	AuthorizationProviderExtensionAuthorizationRequestHandlerProtocol = objc.GetProtocol("ASAuthorizationProviderExtensionAuthorizationRequestHandler")
}
