// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import "github.com/ebitengine/purego/objc"

// setAddressProtocol is the setAddress: protocol.
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//
// Use this protocol when registering custom classes that conform to setAddress:.
var setAddressProtocol *objc.Protocol

func init() {
	setAddressProtocol = objc.GetProtocol("setAddress:")
}

