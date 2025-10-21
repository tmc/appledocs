// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import "github.com/ebitengine/purego/objc"

// centralManagerProtocol is the centralManager: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+
//
// Use this protocol when registering custom classes that conform to centralManager:.
var centralManagerProtocol *objc.Protocol

func init() {
	centralManagerProtocol = objc.GetProtocol("centralManager:")
}


