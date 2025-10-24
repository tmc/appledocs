// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// drawerWillCloseProtocol is the drawerWillClose: protocol.
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.13)
//
// Use this protocol when registering custom classes that conform to drawerWillClose:.
var drawerWillCloseProtocol *objc.Protocol

func init() {
	drawerWillCloseProtocol = objc.GetProtocol("drawerWillClose:")
}
