// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import "github.com/ebitengine/purego/objc"

// textureProtocol is the texture protocol.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to texture.
var textureProtocol *objc.Protocol

func init() {
	textureProtocol = objc.GetProtocol("texture")
}

