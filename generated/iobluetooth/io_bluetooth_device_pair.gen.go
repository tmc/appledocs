// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [BluetoothDevicePair] class.
var (
	BluetoothDevicePairClass     _BluetoothDevicePairClass
	BluetoothDevicePairClassOnce sync.Once
)

func getBluetoothDevicePairClass() _BluetoothDevicePairClass {
	BluetoothDevicePairClassOnce.Do(func() {
		BluetoothDevicePairClass = _BluetoothDevicePairClass{objc.GetClass("IOBluetoothDevicePair")}
	})
	return BluetoothDevicePairClass
}

type _BluetoothDevicePairClass struct {
	class objc.Class
}

// An interface definition for the [BluetoothDevicePair] class.
type IBluetoothDevicePair interface {
	objectivec.IObject
	Device() unsafe.Pointer
	ReplyPINCodePINCode(PINCodeSize unsafe.Pointer, PINCode unsafe.Pointer)
	ReplyUserConfirmation(reply bool)
	SetDevice(inDevice unsafe.Pointer)
	Start() unsafe.Pointer
	Stop()
}

// An instance of IOBluetoothDevicePair represents a pairing attempt to a remote Bluetooth device.
//
// Use the IOBluetoothDevicePair object to attempt to pair with any Bluetooth device. Once -start is invoked on it, progress is returned to the delegate via the messages defined below. This object enables you to pair with devices within your application without having to use the standard panels provided by the IOBluetoothUI framework, allowing you to write custom UI to select devices, and still handle the ability to perform device pairings. Of note is that this object MAY attempt to perform two low-level pairings, depending on the type of device you are attempting to pair. This is inconsequential to your code, however, as it occurs automatically and does not change the messaging. Once started, the pairing can be stopped. This will set the delegate to nil and then attempt to disconnect from the device if already connected.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePair
type BluetoothDevicePair struct {
	objectivec.Object
}

// BluetoothDevicePairFrom constructs a [BluetoothDevicePair] from an unsafe.Pointer.
//
// An instance of IOBluetoothDevicePair represents a pairing attempt to a remote Bluetooth device.
func BluetoothDevicePairFrom(ptr unsafe.Pointer) BluetoothDevicePair {
	return BluetoothDevicePair{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (bc _BluetoothDevicePairClass) Alloc() BluetoothDevicePair {
	rv := objc.Send[BluetoothDevicePair](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BluetoothDevicePairClass) New() BluetoothDevicePair {
	rv := objc.Send[BluetoothDevicePair](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BluetoothDevicePair) Init() BluetoothDevicePair {
	rv := objc.Send[BluetoothDevicePair](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BluetoothDevicePair) Autorelease() BluetoothDevicePair {
	rv := objc.Send[BluetoothDevicePair](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBluetoothDevicePair creates a new BluetoothDevicePair instance.
func NewBluetoothDevicePair() BluetoothDevicePair {
	return getBluetoothDevicePairClass().New()
}




// Creates an autorelease IOBluetoothDevicePair object with a device as the pairing target.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePair/init(device:)
func NewBluetoothDevicePairWithDevice(device unsafe.Pointer) BluetoothDevicePair {
	rv := objc.Send[BluetoothDevicePair](objc.ID(getBluetoothDevicePairClass().class), objc.Sel("pairWithDevice:"), device)
	return rv
}


// Creates an autorelease IOBluetoothDevicePair object with a device as the pairing target.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePair/init(device:)
func (bc _BluetoothDevicePairClass) PairWithDevice(device unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("pairWithDevice:"), device)
	return rv
}

// Get the IOBluetoothDevice being used by the object.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePair/device()
func (b_ BluetoothDevicePair) Device() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("device"))
	return rv
}

// This is the required reply to the devicePairingPINCodeRequest delegate message. Set the PIN code to use during pairing if required.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePair/replyPINCode(_:pinCode:)
func (b_ BluetoothDevicePair) ReplyPINCodePINCode(PINCodeSize unsafe.Pointer, PINCode unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("replyPINCode:PINCode:"), PINCodeSize, PINCode)
}

// This is the required reply to the devicePairingUserConfirmationRequest delegate message.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePair/replyUserConfirmation(_:)
func (b_ BluetoothDevicePair) ReplyUserConfirmation(reply bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("replyUserConfirmation:"), reply)
}

// Set the device object to pair with. It is retained by the object.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePair/setDevice(_:)
func (b_ BluetoothDevicePair) SetDevice(inDevice unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setDevice:"), inDevice)
}

// Kicks off the pairing with the device.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePair/start()
func (b_ BluetoothDevicePair) Start() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("start"))
	return rv
}

// Stops the current pairing. Removes the delegate and disconnects if device was connected.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePair/stop()
func (b_ BluetoothDevicePair) Stop() {
	objc.Send[objc.ID](b_.ID, objc.Sel("stop"))
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePair/delegate
func (b_ BluetoothDevicePair) Delegate() objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePair/delegate
func (b_ BluetoothDevicePair) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setDelegate:"), value)
}


