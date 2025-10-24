// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// viewProtocol is the view: protocol.
//
// Availability:
//   - macOS +
//
// Use this protocol when registering custom classes that conform to view:.
var viewProtocol *objc.Protocol

func init() {
	viewProtocol = objc.GetProtocol("view:")
}
