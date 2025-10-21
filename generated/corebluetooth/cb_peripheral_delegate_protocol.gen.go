// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import "github.com/ebitengine/purego/objc"

// CBPeripheralDelegateProtocol is the CBPeripheralDelegate protocol.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 4.0+
//
// Use this protocol when registering custom classes that conform to CBPeripheralDelegate.
var CBPeripheralDelegateProtocol *objc.Protocol

func init() {
	CBPeripheralDelegateProtocol = objc.GetProtocol("CBPeripheralDelegate")
}
