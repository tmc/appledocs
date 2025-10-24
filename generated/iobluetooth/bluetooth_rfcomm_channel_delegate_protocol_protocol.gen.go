// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// PBluetoothRFCOMMChannelDelegate is the IOBluetoothRFCOMMChannelDelegate protocol interface.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.iobluetooth/documentation/IOBluetooth/IOBluetoothRFCOMMChannelDelegate
type PBluetoothRFCOMMChannelDelegate interface {
	// Optional methods
	RfcommChannelClosed(rfcommChannel IOBluetoothRFCOMMChannel)
	HasRfcommChannelClosed() bool
	RfcommChannelControlSignalsChanged(rfcommChannel IOBluetoothRFCOMMChannel)
	HasRfcommChannelControlSignalsChanged() bool
	RfcommChannelDataDataLength(rfcommChannel IOBluetoothRFCOMMChannel, dataPointer unsafe.Pointer, dataLength uintptr /* not a class type */)
	HasRfcommChannelDataDataLength() bool
	RfcommChannelFlowControlChanged(rfcommChannel IOBluetoothRFCOMMChannel)
	HasRfcommChannelFlowControlChanged() bool
	RfcommChannelOpenCompleteStatus(rfcommChannel IOBluetoothRFCOMMChannel, error_ int)
	HasRfcommChannelOpenCompleteStatus() bool
	RfcommChannelQueueSpaceAvailable(rfcommChannel IOBluetoothRFCOMMChannel)
	HasRfcommChannelQueueSpaceAvailable() bool
	RfcommChannelWriteCompleteRefconStatus(rfcommChannel IOBluetoothRFCOMMChannel, refcon unsafe.Pointer, error_ int)
	HasRfcommChannelWriteCompleteRefconStatus() bool
	RfcommChannelWriteCompleteRefconStatusBytesWritten(rfcommChannel IOBluetoothRFCOMMChannel, refcon unsafe.Pointer, error_ int, length uintptr /* not a class type */)
	HasRfcommChannelWriteCompleteRefconStatusBytesWritten() bool
}

// BluetoothRFCOMMChannelDelegate is a delegate implementation builder for the PBluetoothRFCOMMChannelDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type BluetoothRFCOMMChannelDelegate struct {
	_RfcommChannelClosed func(rfcommChannel IOBluetoothRFCOMMChannel)
	_RfcommChannelControlSignalsChanged func(rfcommChannel IOBluetoothRFCOMMChannel)
	_RfcommChannelDataDataLength func(rfcommChannel IOBluetoothRFCOMMChannel, dataPointer unsafe.Pointer, dataLength uintptr /* not a class type */)
	_RfcommChannelFlowControlChanged func(rfcommChannel IOBluetoothRFCOMMChannel)
	_RfcommChannelOpenCompleteStatus func(rfcommChannel IOBluetoothRFCOMMChannel, error_ int)
	_RfcommChannelQueueSpaceAvailable func(rfcommChannel IOBluetoothRFCOMMChannel)
	_RfcommChannelWriteCompleteRefconStatus func(rfcommChannel IOBluetoothRFCOMMChannel, refcon unsafe.Pointer, error_ int)
	_RfcommChannelWriteCompleteRefconStatusBytesWritten func(rfcommChannel IOBluetoothRFCOMMChannel, refcon unsafe.Pointer, error_ int, length uintptr /* not a class type */)
}

// SetRfcommChannelClosed sets the handler for the RfcommChannelClosed delegate method.
func (d *BluetoothRFCOMMChannelDelegate) SetRfcommChannelClosed(f func(rfcommChannel IOBluetoothRFCOMMChannel)) {
	d._RfcommChannelClosed = f
}

// SetRfcommChannelControlSignalsChanged sets the handler for the RfcommChannelControlSignalsChanged delegate method.
func (d *BluetoothRFCOMMChannelDelegate) SetRfcommChannelControlSignalsChanged(f func(rfcommChannel IOBluetoothRFCOMMChannel)) {
	d._RfcommChannelControlSignalsChanged = f
}

// SetRfcommChannelDataDataLength sets the handler for the RfcommChannelDataDataLength delegate method.
func (d *BluetoothRFCOMMChannelDelegate) SetRfcommChannelDataDataLength(f func(rfcommChannel IOBluetoothRFCOMMChannel, dataPointer unsafe.Pointer, dataLength uintptr /* not a class type */)) {
	d._RfcommChannelDataDataLength = f
}

// SetRfcommChannelFlowControlChanged sets the handler for the RfcommChannelFlowControlChanged delegate method.
func (d *BluetoothRFCOMMChannelDelegate) SetRfcommChannelFlowControlChanged(f func(rfcommChannel IOBluetoothRFCOMMChannel)) {
	d._RfcommChannelFlowControlChanged = f
}

// SetRfcommChannelOpenCompleteStatus sets the handler for the RfcommChannelOpenCompleteStatus delegate method.
func (d *BluetoothRFCOMMChannelDelegate) SetRfcommChannelOpenCompleteStatus(f func(rfcommChannel IOBluetoothRFCOMMChannel, error_ int)) {
	d._RfcommChannelOpenCompleteStatus = f
}

// SetRfcommChannelQueueSpaceAvailable sets the handler for the RfcommChannelQueueSpaceAvailable delegate method.
func (d *BluetoothRFCOMMChannelDelegate) SetRfcommChannelQueueSpaceAvailable(f func(rfcommChannel IOBluetoothRFCOMMChannel)) {
	d._RfcommChannelQueueSpaceAvailable = f
}

// SetRfcommChannelWriteCompleteRefconStatus sets the handler for the RfcommChannelWriteCompleteRefconStatus delegate method.
func (d *BluetoothRFCOMMChannelDelegate) SetRfcommChannelWriteCompleteRefconStatus(f func(rfcommChannel IOBluetoothRFCOMMChannel, refcon unsafe.Pointer, error_ int)) {
	d._RfcommChannelWriteCompleteRefconStatus = f
}

// SetRfcommChannelWriteCompleteRefconStatusBytesWritten sets the handler for the RfcommChannelWriteCompleteRefconStatusBytesWritten delegate method.
func (d *BluetoothRFCOMMChannelDelegate) SetRfcommChannelWriteCompleteRefconStatusBytesWritten(f func(rfcommChannel IOBluetoothRFCOMMChannel, refcon unsafe.Pointer, error_ int, length uintptr /* not a class type */)) {
	d._RfcommChannelWriteCompleteRefconStatusBytesWritten = f
}

// RfcommChannelClosed implements the PBluetoothRFCOMMChannelDelegate interface.
func (d *BluetoothRFCOMMChannelDelegate) RfcommChannelClosed(rfcommChannel IOBluetoothRFCOMMChannel) {
	if d._RfcommChannelClosed != nil {
		d._RfcommChannelClosed(rfcommChannel)
	}
}

// HasRfcommChannelClosed returns true if a handler for RfcommChannelClosed has been set.
func (d *BluetoothRFCOMMChannelDelegate) HasRfcommChannelClosed() bool {
	return d._RfcommChannelClosed != nil
}

// RfcommChannelControlSignalsChanged implements the PBluetoothRFCOMMChannelDelegate interface.
func (d *BluetoothRFCOMMChannelDelegate) RfcommChannelControlSignalsChanged(rfcommChannel IOBluetoothRFCOMMChannel) {
	if d._RfcommChannelControlSignalsChanged != nil {
		d._RfcommChannelControlSignalsChanged(rfcommChannel)
	}
}

// HasRfcommChannelControlSignalsChanged returns true if a handler for RfcommChannelControlSignalsChanged has been set.
func (d *BluetoothRFCOMMChannelDelegate) HasRfcommChannelControlSignalsChanged() bool {
	return d._RfcommChannelControlSignalsChanged != nil
}

// RfcommChannelDataDataLength implements the PBluetoothRFCOMMChannelDelegate interface.
func (d *BluetoothRFCOMMChannelDelegate) RfcommChannelDataDataLength(rfcommChannel IOBluetoothRFCOMMChannel, dataPointer unsafe.Pointer, dataLength uintptr /* not a class type */) {
	if d._RfcommChannelDataDataLength != nil {
		d._RfcommChannelDataDataLength(rfcommChannel, dataPointer, dataLength)
	}
}

// HasRfcommChannelDataDataLength returns true if a handler for RfcommChannelDataDataLength has been set.
func (d *BluetoothRFCOMMChannelDelegate) HasRfcommChannelDataDataLength() bool {
	return d._RfcommChannelDataDataLength != nil
}

// RfcommChannelFlowControlChanged implements the PBluetoothRFCOMMChannelDelegate interface.
func (d *BluetoothRFCOMMChannelDelegate) RfcommChannelFlowControlChanged(rfcommChannel IOBluetoothRFCOMMChannel) {
	if d._RfcommChannelFlowControlChanged != nil {
		d._RfcommChannelFlowControlChanged(rfcommChannel)
	}
}

// HasRfcommChannelFlowControlChanged returns true if a handler for RfcommChannelFlowControlChanged has been set.
func (d *BluetoothRFCOMMChannelDelegate) HasRfcommChannelFlowControlChanged() bool {
	return d._RfcommChannelFlowControlChanged != nil
}

// RfcommChannelOpenCompleteStatus implements the PBluetoothRFCOMMChannelDelegate interface.
func (d *BluetoothRFCOMMChannelDelegate) RfcommChannelOpenCompleteStatus(rfcommChannel IOBluetoothRFCOMMChannel, error_ int) {
	if d._RfcommChannelOpenCompleteStatus != nil {
		d._RfcommChannelOpenCompleteStatus(rfcommChannel, error_)
	}
}

// HasRfcommChannelOpenCompleteStatus returns true if a handler for RfcommChannelOpenCompleteStatus has been set.
func (d *BluetoothRFCOMMChannelDelegate) HasRfcommChannelOpenCompleteStatus() bool {
	return d._RfcommChannelOpenCompleteStatus != nil
}

// RfcommChannelQueueSpaceAvailable implements the PBluetoothRFCOMMChannelDelegate interface.
func (d *BluetoothRFCOMMChannelDelegate) RfcommChannelQueueSpaceAvailable(rfcommChannel IOBluetoothRFCOMMChannel) {
	if d._RfcommChannelQueueSpaceAvailable != nil {
		d._RfcommChannelQueueSpaceAvailable(rfcommChannel)
	}
}

// HasRfcommChannelQueueSpaceAvailable returns true if a handler for RfcommChannelQueueSpaceAvailable has been set.
func (d *BluetoothRFCOMMChannelDelegate) HasRfcommChannelQueueSpaceAvailable() bool {
	return d._RfcommChannelQueueSpaceAvailable != nil
}

// RfcommChannelWriteCompleteRefconStatus implements the PBluetoothRFCOMMChannelDelegate interface.
func (d *BluetoothRFCOMMChannelDelegate) RfcommChannelWriteCompleteRefconStatus(rfcommChannel IOBluetoothRFCOMMChannel, refcon unsafe.Pointer, error_ int) {
	if d._RfcommChannelWriteCompleteRefconStatus != nil {
		d._RfcommChannelWriteCompleteRefconStatus(rfcommChannel, refcon, error_)
	}
}

// HasRfcommChannelWriteCompleteRefconStatus returns true if a handler for RfcommChannelWriteCompleteRefconStatus has been set.
func (d *BluetoothRFCOMMChannelDelegate) HasRfcommChannelWriteCompleteRefconStatus() bool {
	return d._RfcommChannelWriteCompleteRefconStatus != nil
}

// RfcommChannelWriteCompleteRefconStatusBytesWritten implements the PBluetoothRFCOMMChannelDelegate interface.
func (d *BluetoothRFCOMMChannelDelegate) RfcommChannelWriteCompleteRefconStatusBytesWritten(rfcommChannel IOBluetoothRFCOMMChannel, refcon unsafe.Pointer, error_ int, length uintptr /* not a class type */) {
	if d._RfcommChannelWriteCompleteRefconStatusBytesWritten != nil {
		d._RfcommChannelWriteCompleteRefconStatusBytesWritten(rfcommChannel, refcon, error_, length)
	}
}

// HasRfcommChannelWriteCompleteRefconStatusBytesWritten returns true if a handler for RfcommChannelWriteCompleteRefconStatusBytesWritten has been set.
func (d *BluetoothRFCOMMChannelDelegate) HasRfcommChannelWriteCompleteRefconStatusBytesWritten() bool {
	return d._RfcommChannelWriteCompleteRefconStatusBytesWritten != nil
}
