// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// readableTypesForPasteboardProtocol is the readableTypesForPasteboard: protocol.
//
// Availability:
//   - macOS +
//
// Use this protocol when registering custom classes that conform to readableTypesForPasteboard:.
var readableTypesForPasteboardProtocol *objc.Protocol

func init() {
	readableTypesForPasteboardProtocol = objc.GetProtocol("readableTypesForPasteboard:")
}
