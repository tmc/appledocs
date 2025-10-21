// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// windowWillExitFullScreenProtocol is the windowWillExitFullScreen: protocol.
//
// Availability:
//   - macOS 10.7+
//
// Use this protocol when registering custom classes that conform to windowWillExitFullScreen:.
var windowWillExitFullScreenProtocol *objc.Protocol

func init() {
	windowWillExitFullScreenProtocol = objc.GetProtocol("windowWillExitFullScreen:")
}
