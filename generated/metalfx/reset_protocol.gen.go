// Code generated from Apple documentation for MetalFX. DO NOT EDIT.

package metalfx

import "github.com/ebitengine/purego/objc"

// resetProtocol is the reset protocol.
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to reset.
var resetProtocol *objc.Protocol

func init() {
	resetProtocol = objc.GetProtocol("reset")
}
