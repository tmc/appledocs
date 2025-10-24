// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// frameProtocol is the frame protocol.
//
// Availability:
//   - macOS 10.15+
//
// Use this protocol when registering custom classes that conform to frame.
var frameProtocol *objc.Protocol

func init() {
	frameProtocol = objc.GetProtocol("frame")
}
