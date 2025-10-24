// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import "github.com/ebitengine/purego/objc"

// textProtocol is the text protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 3.2+
//   - iPadOS 3.2+
//   - macOS 10.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to text.
var textProtocol *objc.Protocol

func init() {
	textProtocol = objc.GetProtocol("text")
}

