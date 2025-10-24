// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import "github.com/ebitengine/purego/objc"

// compressionProtocol is the compression protocol.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - tvOS +
//   - visionOS +
//
// Use this protocol when registering custom classes that conform to compression.
var compressionProtocol *objc.Protocol

func init() {
	compressionProtocol = objc.GetProtocol("compression")
}

