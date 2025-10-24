// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// TextInputProtocol is the NSTextInput protocol.
//
// Availability:
//   - macOS +
//
// Use this protocol when registering custom classes that conform to NSTextInput.
var TextInputProtocol *objc.Protocol

func init() {
	TextInputProtocol = objc.GetProtocol("NSTextInput")
}
