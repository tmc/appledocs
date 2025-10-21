// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// menuDidCloseProtocol is the menuDidClose: protocol.
//
// Availability:
//   - macOS 10.5+
//
// Use this protocol when registering custom classes that conform to menuDidClose:.
var menuDidCloseProtocol *objc.Protocol

func init() {
	menuDidCloseProtocol = objc.GetProtocol("menuDidClose:")
}
