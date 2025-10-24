// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import "github.com/ebitengine/purego/objc"

// rankProtocol is the rank protocol.
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to rank.
var rankProtocol *objc.Protocol

func init() {
	rankProtocol = objc.GetProtocol("rank")
}

