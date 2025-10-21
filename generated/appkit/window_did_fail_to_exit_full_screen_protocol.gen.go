// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// windowDidFailToExitFullScreenProtocol is the windowDidFailToExitFullScreen: protocol.
//
// Availability:
//   - macOS 10.7+
//
// Use this protocol when registering custom classes that conform to windowDidFailToExitFullScreen:.
var windowDidFailToExitFullScreenProtocol *objc.Protocol

func init() {
	windowDidFailToExitFullScreenProtocol = objc.GetProtocol("windowDidFailToExitFullScreen:")
}
