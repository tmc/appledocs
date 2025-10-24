// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import "github.com/ebitengine/purego/objc"

// excludedCredentialsProtocol is the excludedCredentials protocol.
//
// Availability:
//   - Mac Catalyst 16.6+
//   - iOS 17.4+
//   - iPadOS 17.4+
//   - macOS 13.5+
//
// Use this protocol when registering custom classes that conform to excludedCredentials.
var excludedCredentialsProtocol *objc.Protocol

func init() {
	excludedCredentialsProtocol = objc.GetProtocol("excludedCredentials")
}

