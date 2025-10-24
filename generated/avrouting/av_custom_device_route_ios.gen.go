//go:build darwin && ios

// Code generated from Apple documentation for AVRouting. DO NOT EDIT.

package avrouting

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CustomDeviceRoute


// iOS-only properties

// An identifier to use to establish a connection to a Bluetooth device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomDeviceRoute/bluetoothIdentifier
func (c_ CustomDeviceRoute) BluetoothIdentifier() foundation.UUID {
	rv := objc.Send[foundation.UUID](c_.ID, objc.Sel("bluetoothIdentifier"))
	return rv
}

// A local or remote endpoint to connect to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomDeviceRoute/networkEndpoint
func (c_ CustomDeviceRoute) NetworkEndpoint() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("networkEndpoint"))
	return rv
}





