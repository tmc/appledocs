// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import "github.com/ebitengine/purego/objc"

// WebAuthenticationPresentationContextProvidingProtocol is the ASWebAuthenticationPresentationContextProviding protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to ASWebAuthenticationPresentationContextProviding.
var WebAuthenticationPresentationContextProvidingProtocol *objc.Protocol

func init() {
	WebAuthenticationPresentationContextProvidingProtocol = objc.GetProtocol("ASWebAuthenticationPresentationContextProviding")
}
