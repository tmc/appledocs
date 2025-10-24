// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// PBluetoothL2CAPChannelDelegate is the IOBluetoothL2CAPChannelDelegate protocol interface.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.iobluetooth/documentation/IOBluetooth/IOBluetoothL2CAPChannelDelegate
type PBluetoothL2CAPChannelDelegate interface {
	// Optional methods
	L2capChannelClosed(l2capChannel IOBluetoothL2CAPChannel)
	HasL2capChannelClosed() bool
	L2capChannelDataDataLength(l2capChannel IOBluetoothL2CAPChannel, dataPointer unsafe.Pointer, dataLength uintptr /* not a class type */)
	HasL2capChannelDataDataLength() bool
	L2capChannelOpenCompleteStatus(l2capChannel IOBluetoothL2CAPChannel, error_ int)
	HasL2capChannelOpenCompleteStatus() bool
	L2capChannelQueueSpaceAvailable(l2capChannel IOBluetoothL2CAPChannel)
	HasL2capChannelQueueSpaceAvailable() bool
	L2capChannelReconfigured(l2capChannel IOBluetoothL2CAPChannel)
	HasL2capChannelReconfigured() bool
	L2capChannelWriteCompleteRefconStatus(l2capChannel IOBluetoothL2CAPChannel, refcon unsafe.Pointer, error_ int)
	HasL2capChannelWriteCompleteRefconStatus() bool
}

// BluetoothL2CAPChannelDelegate is a delegate implementation builder for the PBluetoothL2CAPChannelDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type BluetoothL2CAPChannelDelegate struct {
	_L2capChannelClosed func(l2capChannel IOBluetoothL2CAPChannel)
	_L2capChannelDataDataLength func(l2capChannel IOBluetoothL2CAPChannel, dataPointer unsafe.Pointer, dataLength uintptr /* not a class type */)
	_L2capChannelOpenCompleteStatus func(l2capChannel IOBluetoothL2CAPChannel, error_ int)
	_L2capChannelQueueSpaceAvailable func(l2capChannel IOBluetoothL2CAPChannel)
	_L2capChannelReconfigured func(l2capChannel IOBluetoothL2CAPChannel)
	_L2capChannelWriteCompleteRefconStatus func(l2capChannel IOBluetoothL2CAPChannel, refcon unsafe.Pointer, error_ int)
}

// SetL2capChannelClosed sets the handler for the L2capChannelClosed delegate method.
func (d *BluetoothL2CAPChannelDelegate) SetL2capChannelClosed(f func(l2capChannel IOBluetoothL2CAPChannel)) {
	d._L2capChannelClosed = f
}

// SetL2capChannelDataDataLength sets the handler for the L2capChannelDataDataLength delegate method.
func (d *BluetoothL2CAPChannelDelegate) SetL2capChannelDataDataLength(f func(l2capChannel IOBluetoothL2CAPChannel, dataPointer unsafe.Pointer, dataLength uintptr /* not a class type */)) {
	d._L2capChannelDataDataLength = f
}

// SetL2capChannelOpenCompleteStatus sets the handler for the L2capChannelOpenCompleteStatus delegate method.
func (d *BluetoothL2CAPChannelDelegate) SetL2capChannelOpenCompleteStatus(f func(l2capChannel IOBluetoothL2CAPChannel, error_ int)) {
	d._L2capChannelOpenCompleteStatus = f
}

// SetL2capChannelQueueSpaceAvailable sets the handler for the L2capChannelQueueSpaceAvailable delegate method.
func (d *BluetoothL2CAPChannelDelegate) SetL2capChannelQueueSpaceAvailable(f func(l2capChannel IOBluetoothL2CAPChannel)) {
	d._L2capChannelQueueSpaceAvailable = f
}

// SetL2capChannelReconfigured sets the handler for the L2capChannelReconfigured delegate method.
func (d *BluetoothL2CAPChannelDelegate) SetL2capChannelReconfigured(f func(l2capChannel IOBluetoothL2CAPChannel)) {
	d._L2capChannelReconfigured = f
}

// SetL2capChannelWriteCompleteRefconStatus sets the handler for the L2capChannelWriteCompleteRefconStatus delegate method.
func (d *BluetoothL2CAPChannelDelegate) SetL2capChannelWriteCompleteRefconStatus(f func(l2capChannel IOBluetoothL2CAPChannel, refcon unsafe.Pointer, error_ int)) {
	d._L2capChannelWriteCompleteRefconStatus = f
}

// L2capChannelClosed implements the PBluetoothL2CAPChannelDelegate interface.
func (d *BluetoothL2CAPChannelDelegate) L2capChannelClosed(l2capChannel IOBluetoothL2CAPChannel) {
	if d._L2capChannelClosed != nil {
		d._L2capChannelClosed(l2capChannel)
	}
}

// HasL2capChannelClosed returns true if a handler for L2capChannelClosed has been set.
func (d *BluetoothL2CAPChannelDelegate) HasL2capChannelClosed() bool {
	return d._L2capChannelClosed != nil
}

// L2capChannelDataDataLength implements the PBluetoothL2CAPChannelDelegate interface.
func (d *BluetoothL2CAPChannelDelegate) L2capChannelDataDataLength(l2capChannel IOBluetoothL2CAPChannel, dataPointer unsafe.Pointer, dataLength uintptr /* not a class type */) {
	if d._L2capChannelDataDataLength != nil {
		d._L2capChannelDataDataLength(l2capChannel, dataPointer, dataLength)
	}
}

// HasL2capChannelDataDataLength returns true if a handler for L2capChannelDataDataLength has been set.
func (d *BluetoothL2CAPChannelDelegate) HasL2capChannelDataDataLength() bool {
	return d._L2capChannelDataDataLength != nil
}

// L2capChannelOpenCompleteStatus implements the PBluetoothL2CAPChannelDelegate interface.
func (d *BluetoothL2CAPChannelDelegate) L2capChannelOpenCompleteStatus(l2capChannel IOBluetoothL2CAPChannel, error_ int) {
	if d._L2capChannelOpenCompleteStatus != nil {
		d._L2capChannelOpenCompleteStatus(l2capChannel, error_)
	}
}

// HasL2capChannelOpenCompleteStatus returns true if a handler for L2capChannelOpenCompleteStatus has been set.
func (d *BluetoothL2CAPChannelDelegate) HasL2capChannelOpenCompleteStatus() bool {
	return d._L2capChannelOpenCompleteStatus != nil
}

// L2capChannelQueueSpaceAvailable implements the PBluetoothL2CAPChannelDelegate interface.
func (d *BluetoothL2CAPChannelDelegate) L2capChannelQueueSpaceAvailable(l2capChannel IOBluetoothL2CAPChannel) {
	if d._L2capChannelQueueSpaceAvailable != nil {
		d._L2capChannelQueueSpaceAvailable(l2capChannel)
	}
}

// HasL2capChannelQueueSpaceAvailable returns true if a handler for L2capChannelQueueSpaceAvailable has been set.
func (d *BluetoothL2CAPChannelDelegate) HasL2capChannelQueueSpaceAvailable() bool {
	return d._L2capChannelQueueSpaceAvailable != nil
}

// L2capChannelReconfigured implements the PBluetoothL2CAPChannelDelegate interface.
func (d *BluetoothL2CAPChannelDelegate) L2capChannelReconfigured(l2capChannel IOBluetoothL2CAPChannel) {
	if d._L2capChannelReconfigured != nil {
		d._L2capChannelReconfigured(l2capChannel)
	}
}

// HasL2capChannelReconfigured returns true if a handler for L2capChannelReconfigured has been set.
func (d *BluetoothL2CAPChannelDelegate) HasL2capChannelReconfigured() bool {
	return d._L2capChannelReconfigured != nil
}

// L2capChannelWriteCompleteRefconStatus implements the PBluetoothL2CAPChannelDelegate interface.
func (d *BluetoothL2CAPChannelDelegate) L2capChannelWriteCompleteRefconStatus(l2capChannel IOBluetoothL2CAPChannel, refcon unsafe.Pointer, error_ int) {
	if d._L2capChannelWriteCompleteRefconStatus != nil {
		d._L2capChannelWriteCompleteRefconStatus(l2capChannel, refcon, error_)
	}
}

// HasL2capChannelWriteCompleteRefconStatus returns true if a handler for L2capChannelWriteCompleteRefconStatus has been set.
func (d *BluetoothL2CAPChannelDelegate) HasL2capChannelWriteCompleteRefconStatus() bool {
	return d._L2capChannelWriteCompleteRefconStatus != nil
}
