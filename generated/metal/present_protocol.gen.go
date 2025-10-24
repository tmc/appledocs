// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import "github.com/ebitengine/purego/objc"

// presentProtocol is the present protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.11+
//   - tvOS +
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to present.
var presentProtocol *objc.Protocol

func init() {
	presentProtocol = objc.GetProtocol("present")
}

