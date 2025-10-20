// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import "github.com/ebitengine/purego/objc"

// environmentProtocol is the environment: protocol.
//
// Availability:
//   - Mac Catalyst 18.0+
//   - iOS 18.0+
//   - iPadOS 18.0+
//   - macOS 15.0+
//   - visionOS 2.0+
//   - watchOS 11.0+
//
// Use this protocol when registering custom classes that conform to environment:.
var environmentProtocol *objc.Protocol

func init() {
	environmentProtocol = objc.GetProtocol("environment:")
}

