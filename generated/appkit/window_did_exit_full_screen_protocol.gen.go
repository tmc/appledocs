// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// windowDidExitFullScreenProtocol is the windowDidExitFullScreen: protocol.
//
// Availability:
//   - macOS 10.7+
//
// Use this protocol when registering custom classes that conform to windowDidExitFullScreen:.
var windowDidExitFullScreenProtocol *objc.Protocol

func init() {
	windowDidExitFullScreenProtocol = objc.GetProtocol("windowDidExitFullScreen:")
}

