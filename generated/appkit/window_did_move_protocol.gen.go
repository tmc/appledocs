// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// windowDidMoveProtocol is the windowDidMove: protocol.
//
// Availability:
//   - macOS 10.10+
//
// Use this protocol when registering custom classes that conform to windowDidMove:.
var windowDidMoveProtocol *objc.Protocol

func init() {
	windowDidMoveProtocol = objc.GetProtocol("windowDidMove:")
}
