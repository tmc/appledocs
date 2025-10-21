// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import "github.com/ebitengine/purego/objc"

// devicePairingStartedProtocol is the devicePairingStarted: protocol.
//
// Use this protocol when registering custom classes that conform to devicePairingStarted:.
var devicePairingStartedProtocol *objc.Protocol

func init() {
	devicePairingStartedProtocol = objc.GetProtocol("devicePairingStarted:")
}
