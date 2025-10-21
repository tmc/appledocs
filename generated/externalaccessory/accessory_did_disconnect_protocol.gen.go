// Code generated from Apple documentation for ExternalAccessory. DO NOT EDIT.

package externalaccessory

import "github.com/ebitengine/purego/objc"

// accessoryDidDisconnectProtocol is the accessoryDidDisconnect: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 3.0+
//   - iPadOS 3.0+
//   - macOS 10.13+
//   - tvOS 10.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to accessoryDidDisconnect:.
var accessoryDidDisconnectProtocol *objc.Protocol

func init() {
	accessoryDidDisconnectProtocol = objc.GetProtocol("accessoryDidDisconnect:")
}
