// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import "github.com/ebitengine/purego/objc"

// memoryBarrierWithResourcesProtocol is the memoryBarrierWithResources: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 10.14+
//   - tvOS 16.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to memoryBarrierWithResources:.
var memoryBarrierWithResourcesProtocol *objc.Protocol

func init() {
	memoryBarrierWithResourcesProtocol = objc.GetProtocol("memoryBarrierWithResources:")
}
