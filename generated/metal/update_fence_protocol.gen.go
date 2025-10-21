// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import "github.com/ebitengine/purego/objc"

// updateFenceProtocol is the updateFence: protocol.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 11.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to updateFence:.
var updateFenceProtocol *objc.Protocol

func init() {
	updateFenceProtocol = objc.GetProtocol("updateFence:")
}
