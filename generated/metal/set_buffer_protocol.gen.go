// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import "github.com/ebitengine/purego/objc"

// setBufferProtocol is the setBuffer: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 10.13+
//   - tvOS 11.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to setBuffer:.
var setBufferProtocol *objc.Protocol

func init() {
	setBufferProtocol = objc.GetProtocol("setBuffer:")
}

