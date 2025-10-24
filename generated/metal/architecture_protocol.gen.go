// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import "github.com/ebitengine/purego/objc"

// architectureProtocol is the architecture protocol.
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to architecture.
var architectureProtocol *objc.Protocol

func init() {
	architectureProtocol = objc.GetProtocol("architecture")
}

