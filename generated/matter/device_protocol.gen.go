// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import "github.com/ebitengine/purego/objc"

// deviceProtocol is the device: protocol.
//
// Availability:
//   - Mac Catalyst 18.2+
//   - iOS 18.2+
//   - iPadOS 18.2+
//   - macOS 15.2+
//   - tvOS 18.2+
//   - visionOS 2.2+
//   - watchOS 11.2+
//
// Use this protocol when registering custom classes that conform to device:.
var deviceProtocol *objc.Protocol

func init() {
	deviceProtocol = objc.GetProtocol("device:")
}
