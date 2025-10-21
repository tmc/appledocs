// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// windowDidFailToEnterFullScreenProtocol is the windowDidFailToEnterFullScreen: protocol.
//
// Availability:
//   - macOS 10.7+
//
// Use this protocol when registering custom classes that conform to windowDidFailToEnterFullScreen:.
var windowDidFailToEnterFullScreenProtocol *objc.Protocol

func init() {
	windowDidFailToEnterFullScreenProtocol = objc.GetProtocol("windowDidFailToEnterFullScreen:")
}
