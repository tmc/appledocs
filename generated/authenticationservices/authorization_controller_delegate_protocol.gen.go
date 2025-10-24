// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import "github.com/ebitengine/purego/objc"

// AuthorizationControllerDelegateProtocol is the ASAuthorizationControllerDelegate protocol.
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
// Use this protocol when registering custom classes that conform to ASAuthorizationControllerDelegate.
var AuthorizationControllerDelegateProtocol *objc.Protocol

func init() {
	AuthorizationControllerDelegateProtocol = objc.GetProtocol("ASAuthorizationControllerDelegate")
}

