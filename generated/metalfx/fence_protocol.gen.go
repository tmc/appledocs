// Code generated from Apple documentation for MetalFX. DO NOT EDIT.

package metalfx

import "github.com/ebitengine/purego/objc"

// fenceProtocol is the fence protocol.
//
// Availability:
//   - Mac Catalyst 10.0+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - macOS 10.13+
//   - tvOS 10.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to fence.
var fenceProtocol *objc.Protocol

func init() {
	fenceProtocol = objc.GetProtocol("fence")
}


