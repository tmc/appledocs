// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// windowWillMoveProtocol is the windowWillMove: protocol.
//
// Availability:
//   - macOS 10.10+
//
// Use this protocol when registering custom classes that conform to windowWillMove:.
var windowWillMoveProtocol *objc.Protocol

func init() {
	windowWillMoveProtocol = objc.GetProtocol("windowWillMove:")
}
