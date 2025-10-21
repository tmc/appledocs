// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// pasteboardProtocol is the pasteboard: protocol.
//
// Use this protocol when registering custom classes that conform to pasteboard:.
var pasteboardProtocol *objc.Protocol

func init() {
	pasteboardProtocol = objc.GetProtocol("pasteboard:")
}
