// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import "github.com/ebitengine/purego/objc"

// memoryBarrierWithScopeProtocol is the memoryBarrierWithScope: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to memoryBarrierWithScope:.
var memoryBarrierWithScopeProtocol *objc.Protocol

func init() {
	memoryBarrierWithScopeProtocol = objc.GetProtocol("memoryBarrierWithScope:")
}

