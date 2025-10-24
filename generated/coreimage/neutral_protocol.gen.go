// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import "github.com/ebitengine/purego/objc"

// neutralProtocol is the neutral protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.4+
//   - tvOS +
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to neutral.
var neutralProtocol *objc.Protocol

func init() {
	neutralProtocol = objc.GetProtocol("neutral")
}

