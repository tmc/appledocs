// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import "github.com/ebitengine/purego/objc"

// shouldEvaluateTrustForConnectionProtocol is the shouldEvaluateTrustForConnection: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 18.0)
//   - iOS 9.0+ (Deprecated in 18.0)
//   - iPadOS 9.0+ (Deprecated in 18.0)
//   - macOS 10.11+ (Deprecated in 15.0)
//   - tvOS 17.0+ (Deprecated in 18.0)
//   - visionOS 1.0+ (Deprecated in 2.0)
//
// Use this protocol when registering custom classes that conform to shouldEvaluateTrustForConnection:.
var shouldEvaluateTrustForConnectionProtocol *objc.Protocol

func init() {
	shouldEvaluateTrustForConnectionProtocol = objc.GetProtocol("shouldEvaluateTrustForConnection:")
}

