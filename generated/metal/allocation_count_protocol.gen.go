// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import "github.com/ebitengine/purego/objc"

// allocationCountProtocol is the allocationCount protocol.
//
// Availability:
//   - Mac Catalyst 18.0+
//   - iOS 18.0+
//   - iPadOS 18.0+
//   - macOS 15.0+
//   - tvOS 18.0+
//   - visionOS 2.0+
//
// Use this protocol when registering custom classes that conform to allocationCount.
var allocationCountProtocol *objc.Protocol

func init() {
	allocationCountProtocol = objc.GetProtocol("allocationCount")
}

