// Code generated from Apple documentation for MetalFX. DO NOT EDIT.

package metalfx

import "github.com/ebitengine/purego/objc"

// encodeToCommandBufferProtocol is the encodeToCommandBuffer: protocol.
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//
// Use this protocol when registering custom classes that conform to encodeToCommandBuffer:.
var encodeToCommandBufferProtocol *objc.Protocol

func init() {
	encodeToCommandBufferProtocol = objc.GetProtocol("encodeToCommandBuffer:")
}
