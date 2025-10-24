// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// textShouldEndEditingProtocol is the textShouldEndEditing: protocol.
//
// Availability:
//   - macOS +
//
// Use this protocol when registering custom classes that conform to textShouldEndEditing:.
var textShouldEndEditingProtocol *objc.Protocol

func init() {
	textShouldEndEditingProtocol = objc.GetProtocol("textShouldEndEditing:")
}
