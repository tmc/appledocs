// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import "github.com/ebitengine/purego/objc"

// imageProtocol is the image protocol.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 10.13+
//   - tvOS 11.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to image.
var imageProtocol *objc.Protocol

func init() {
	imageProtocol = objc.GetProtocol("image")
}

