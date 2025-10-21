// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import "github.com/ebitengine/purego/objc"

// colorProtocol is the color protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.4+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to color.
var colorProtocol *objc.Protocol

func init() {
	colorProtocol = objc.GetProtocol("color")
}
