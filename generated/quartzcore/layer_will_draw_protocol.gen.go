// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import "github.com/ebitengine/purego/objc"

// layerWillDrawProtocol is the layerWillDraw: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - macOS 10.12+
//   - tvOS 10.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to layerWillDraw:.
var layerWillDrawProtocol *objc.Protocol

func init() {
	layerWillDrawProtocol = objc.GetProtocol("layerWillDraw:")
}
