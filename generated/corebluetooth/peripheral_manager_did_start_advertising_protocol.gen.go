// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import "github.com/ebitengine/purego/objc"

// peripheralManagerDidStartAdvertisingProtocol is the peripheralManagerDidStartAdvertising: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.9+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+
//
// Use this protocol when registering custom classes that conform to peripheralManagerDidStartAdvertising:.
var peripheralManagerDidStartAdvertisingProtocol *objc.Protocol

func init() {
	peripheralManagerDidStartAdvertisingProtocol = objc.GetProtocol("peripheralManagerDidStartAdvertising:")
}

