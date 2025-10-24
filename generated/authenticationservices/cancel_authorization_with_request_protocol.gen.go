// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import "github.com/ebitengine/purego/objc"

// cancelAuthorizationWithRequestProtocol is the cancelAuthorizationWithRequest: protocol.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to cancelAuthorizationWithRequest:.
var cancelAuthorizationWithRequestProtocol *objc.Protocol

func init() {
	cancelAuthorizationWithRequestProtocol = objc.GetProtocol("cancelAuthorizationWithRequest:")
}

