// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import "github.com/ebitengine/purego/objc"

// CredentialRequestProtocol is the ASCredentialRequest protocol.
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to ASCredentialRequest.
var CredentialRequestProtocol *objc.Protocol

func init() {
	CredentialRequestProtocol = objc.GetProtocol("ASCredentialRequest")
}

