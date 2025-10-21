// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// accessibilityURLProtocol is the accessibilityURL protocol.
//
// Availability:
//   - macOS 10.10+
//
// Use this protocol when registering custom classes that conform to accessibilityURL.
var accessibilityURLProtocol *objc.Protocol

func init() {
	accessibilityURLProtocol = objc.GetProtocol("accessibilityURL")
}
