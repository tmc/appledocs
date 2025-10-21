// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import "github.com/ebitengine/purego/objc"

// autoreversesProtocol is the autoreverses protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to autoreverses.
var autoreversesProtocol *objc.Protocol

func init() {
	autoreversesProtocol = objc.GetProtocol("autoreverses")
}
