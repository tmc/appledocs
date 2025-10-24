// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// stringProtocol is the string protocol.
//
// Availability:
//   - macOS +
//
// Use this protocol when registering custom classes that conform to string.
var stringProtocol *objc.Protocol

func init() {
	stringProtocol = objc.GetProtocol("string")
}
