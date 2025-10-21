// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import "github.com/ebitengine/purego/objc"

// AuthorizationProviderExtensionRegistrationHandlerProtocol is the ASAuthorizationProviderExtensionRegistrationHandler protocol.
//
// Availability:
//   - macOS 13.0+
//
// Use this protocol when registering custom classes that conform to ASAuthorizationProviderExtensionRegistrationHandler.
var AuthorizationProviderExtensionRegistrationHandlerProtocol *objc.Protocol

func init() {
	AuthorizationProviderExtensionRegistrationHandlerProtocol = objc.GetProtocol("ASAuthorizationProviderExtensionRegistrationHandler")
}
