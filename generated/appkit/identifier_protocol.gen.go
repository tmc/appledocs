// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// identifierProtocol is the identifier protocol.
//
// Availability:
//   - macOS +
//
// Use this protocol when registering custom classes that conform to identifier.
var identifierProtocol *objc.Protocol

func init() {
	identifierProtocol = objc.GetProtocol("identifier")
}
