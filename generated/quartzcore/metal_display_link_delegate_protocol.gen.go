// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import "github.com/ebitengine/purego/objc"

// MetalDisplayLinkDelegateProtocol is the CAMetalDisplayLinkDelegate protocol.
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to CAMetalDisplayLinkDelegate.
var MetalDisplayLinkDelegateProtocol *objc.Protocol

func init() {
	MetalDisplayLinkDelegateProtocol = objc.GetProtocol("CAMetalDisplayLinkDelegate")
}
