// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import "github.com/ebitengine/purego/objc"

// clientDataProtocol is the clientData protocol.
//
// Availability:
//   - Mac Catalyst 17.4+
//   - iOS 17.4+
//   - iPadOS 17.4+
//   - macOS 14.4+
//
// Use this protocol when registering custom classes that conform to clientData.
var clientDataProtocol *objc.Protocol

func init() {
	clientDataProtocol = objc.GetProtocol("clientData")
}


