// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// textDidChangeProtocol is the textDidChange: protocol.
//
// Availability:
//   - macOS 10.10+
//
// Use this protocol when registering custom classes that conform to textDidChange:.
var textDidChangeProtocol *objc.Protocol

func init() {
	textDidChangeProtocol = objc.GetProtocol("textDidChange:")
}
