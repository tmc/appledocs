// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import "github.com/ebitengine/purego/objc"

// devicePairingConnectedProtocol is the devicePairingConnected: protocol.
//
// Availability:
//   - macOS +
//
// Use this protocol when registering custom classes that conform to devicePairingConnected:.
var devicePairingConnectedProtocol *objc.Protocol

func init() {
	devicePairingConnectedProtocol = objc.GetProtocol("devicePairingConnected:")
}

