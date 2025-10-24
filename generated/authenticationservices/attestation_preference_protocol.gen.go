// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import "github.com/ebitengine/purego/objc"

// attestationPreferenceProtocol is the attestationPreference protocol.
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to attestationPreference.
var attestationPreferenceProtocol *objc.Protocol

func init() {
	attestationPreferenceProtocol = objc.GetProtocol("attestationPreference")
}

