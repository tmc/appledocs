// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// windowWillEnterFullScreenProtocol is the windowWillEnterFullScreen: protocol.
//
// Availability:
//   - macOS 10.7+
//
// Use this protocol when registering custom classes that conform to windowWillEnterFullScreen:.
var windowWillEnterFullScreenProtocol *objc.Protocol

func init() {
	windowWillEnterFullScreenProtocol = objc.GetProtocol("windowWillEnterFullScreen:")
}

