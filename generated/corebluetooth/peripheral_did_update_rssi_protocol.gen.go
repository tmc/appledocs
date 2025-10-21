// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import "github.com/ebitengine/purego/objc"

// peripheralDidUpdateRSSIProtocol is the peripheralDidUpdateRSSI: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 5.0+ (Deprecated in 8.0)
//   - iPadOS 5.0+ (Deprecated in 8.0)
//   - macOS 10.7+ (Deprecated in 10.13)
//   - tvOS 9.0+
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 2.0)
//
// Use this protocol when registering custom classes that conform to peripheralDidUpdateRSSI:.
var peripheralDidUpdateRSSIProtocol *objc.Protocol

func init() {
	peripheralDidUpdateRSSIProtocol = objc.GetProtocol("peripheralDidUpdateRSSI:")
}
