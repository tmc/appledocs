// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import "github.com/ebitengine/purego/objc"

// regionProtocol is the region protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - macOS 10.12+
//   - tvOS 10.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to region.
var regionProtocol *objc.Protocol

func init() {
	regionProtocol = objc.GetProtocol("region")
}
