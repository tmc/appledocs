//go:build darwin && ios

// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for CBPeripheral


// iOS-only properties

// A Boolean value that indicates if the remote device has authorization to receive data over ANCS protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/ancsAuthorized
func (c_ CBPeripheral) AncsAuthorized() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("ancsAuthorized"))
	return rv
}





