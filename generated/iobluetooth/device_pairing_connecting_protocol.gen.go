// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import "github.com/ebitengine/purego/objc"

// devicePairingConnectingProtocol is the devicePairingConnecting: protocol.
//
// Availability:
//   - macOS +
//
// Use this protocol when registering custom classes that conform to devicePairingConnecting:.
var devicePairingConnectingProtocol *objc.Protocol

func init() {
	devicePairingConnectingProtocol = objc.GetProtocol("devicePairingConnecting:")
}

