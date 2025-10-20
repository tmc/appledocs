// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// textDidEndEditingProtocol is the textDidEndEditing: protocol.
//
// Availability:
//   - macOS 10.10+
//
// Use this protocol when registering custom classes that conform to textDidEndEditing:.
var textDidEndEditingProtocol *objc.Protocol

func init() {
	textDidEndEditingProtocol = objc.GetProtocol("textDidEndEditing:")
}
