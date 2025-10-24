// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// actionProtocol is the action protocol.
//
// Availability:
//   - macOS 10.0+
//
// Use this protocol when registering custom classes that conform to action.
var actionProtocol *objc.Protocol

func init() {
	actionProtocol = objc.GetProtocol("action")
}
