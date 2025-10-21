// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// menuProtocol is the menu: protocol.
//
// Availability:
//   - macOS 10.5+
//
// Use this protocol when registering custom classes that conform to menu:.
var menuProtocol *objc.Protocol

func init() {
	menuProtocol = objc.GetProtocol("menu:")
}
