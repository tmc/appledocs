// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// writableTypesForPasteboardProtocol is the writableTypesForPasteboard: protocol.
//
// Availability:
//   - macOS +
//
// Use this protocol when registering custom classes that conform to writableTypesForPasteboard:.
var writableTypesForPasteboardProtocol *objc.Protocol

func init() {
	writableTypesForPasteboardProtocol = objc.GetProtocol("writableTypesForPasteboard:")
}
