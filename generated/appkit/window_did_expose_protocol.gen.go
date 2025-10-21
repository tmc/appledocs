// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// windowDidExposeProtocol is the windowDidExpose: protocol.
//
// Availability:
//   - macOS 10.10+
//
// Use this protocol when registering custom classes that conform to windowDidExpose:.
var windowDidExposeProtocol *objc.Protocol

func init() {
	windowDidExposeProtocol = objc.GetProtocol("windowDidExpose:")
}
