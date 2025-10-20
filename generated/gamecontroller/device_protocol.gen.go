// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import "github.com/ebitengine/purego/objc"

// deviceProtocol is the device protocol.
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to device.
var deviceProtocol *objc.Protocol

func init() {
	deviceProtocol = objc.GetProtocol("device")
}


