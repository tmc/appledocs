// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// PasteboardWritingProtocol is the NSPasteboardWriting protocol.
//
// Use this protocol when registering custom classes that conform to NSPasteboardWriting.
var PasteboardWritingProtocol *objc.Protocol

func init() {
	PasteboardWritingProtocol = objc.GetProtocol("NSPasteboardWriting")
}
