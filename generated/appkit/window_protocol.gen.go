// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// windowProtocol is the window: protocol.
//
// Availability:
//   - macOS 10.5+
//
// Use this protocol when registering custom classes that conform to window:.
var windowProtocol *objc.Protocol

func init() {
	windowProtocol = objc.GetProtocol("window:")
}
