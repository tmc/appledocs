// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// menuWillOpenProtocol is the menuWillOpen: protocol.
//
// Availability:
//   - macOS 10.5+
//
// Use this protocol when registering custom classes that conform to menuWillOpen:.
var menuWillOpenProtocol *objc.Protocol

func init() {
	menuWillOpenProtocol = objc.GetProtocol("menuWillOpen:")
}
