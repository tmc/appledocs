// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import "github.com/ebitengine/purego/objc"

// centralManagerProtocol is the centralManager: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+
//
// Use this protocol when registering custom classes that conform to centralManager:.
var centralManagerProtocol *objc.Protocol

func init() {
	centralManagerProtocol = objc.GetProtocol("centralManager:")
}
