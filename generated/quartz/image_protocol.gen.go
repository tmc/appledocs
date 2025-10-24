// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import "github.com/ebitengine/purego/objc"

// imageProtocol is the image protocol.
//
// Availability:
//   - macOS 10.4+
//
// Use this protocol when registering custom classes that conform to image.
var imageProtocol *objc.Protocol

func init() {
	imageProtocol = objc.GetProtocol("image")
}

