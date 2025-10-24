// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import "github.com/ebitengine/purego/objc"

// animationDidStopProtocol is the animationDidStop: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - macOS 10.12+
//   - tvOS 10.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to animationDidStop:.
var animationDidStopProtocol *objc.Protocol

func init() {
	animationDidStopProtocol = objc.GetProtocol("animationDidStop:")
}

