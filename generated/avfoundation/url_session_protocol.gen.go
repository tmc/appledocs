// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import "github.com/ebitengine/purego/objc"

// URLSessionProtocol is the URLSession: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 10.15+ (Deprecated in 11.0)
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to URLSession:.
var URLSessionProtocol *objc.Protocol

func init() {
	URLSessionProtocol = objc.GetProtocol("URLSession:")
}

