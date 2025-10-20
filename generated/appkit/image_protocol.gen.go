// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// imageProtocol is the image: protocol.
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.4)
//
// Use this protocol when registering custom classes that conform to image:.
var imageProtocol *objc.Protocol

func init() {
	imageProtocol = objc.GetProtocol("image:")
}
