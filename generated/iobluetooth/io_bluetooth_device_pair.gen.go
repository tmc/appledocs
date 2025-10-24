// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class IOBluetoothDevicePair */


/* debug [class_header]: Header for IOBluetoothDevicePair */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BluetoothDevicePair */
// An interface definition for the [BluetoothDevicePair] class.
type IBluetoothDevicePair interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for BluetoothDevicePair */
	// properties:
	Delegate() objc.ID
	SetDelegate(value objc.ID)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BluetoothDevicePair */
	// methods:
	Device() IBluetoothDevice
	ReplyPINCodePINCode(PINCodeSize unsafe.Pointer, PINCode objc.IObject /* cross-framework: BluetoothPINCode */)
	ReplyUserConfirmation(reply bool)
	SetDevice(inDevice IOBluetoothDevice)
	Start() int
	Stop()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BluetoothDevicePair */
// Alloc allocates a new instance without initialization.
func (bc _BluetoothDevicePairClass) Alloc() BluetoothDevicePair {
	rv := objc.Send[BluetoothDevicePair](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BluetoothDevicePair */
// An instance of IOBluetoothDevicePair represents a pairing attempt to a remote Bluetooth device.
//
// Use the IOBluetoothDevicePair object to attempt to pair with any Bluetooth device. Once -start is invoked on it, progress is returned to the delegate via the messages defined below. This object enables you to pair with devices within your application without having to use the standard panels provided by the IOBluetoothUI framework, allowing you to write custom UI to select devices, and still handle the ability to perform device pairings. Of note is that this object MAY attempt to perform two low-level pairings, depending on the type of device you are attempting to pair. This is inconsequential to your code, however, as it occurs automatically and does not change the messaging. Once started, the pairing can be stopped. This will set the delegate to nil and then attempt to disconnect from the device if already connected.


// An instance of IOBluetoothDevicePair represents a pairing attempt to a remote Bluetooth device.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BluetoothDevicePair */

// Creates an autorelease IOBluetoothDevicePair object with a device as the pairing target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePair/init(device:)
func NewBluetoothDevicePairWithDevice(device IOBluetoothDevice) BluetoothDevicePair {
	rv := objc.Send[BluetoothDevicePair](objc.ID(getBluetoothDevicePairClass().class), objc.Sel("pairWithDevice:"), device)
	return rv
}/* debug [class_init_methods/constructor]: NewBluetoothDevicePairWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BluetoothDevicePair */

// Creates an autorelease IOBluetoothDevicePair object with a device as the pairing target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePair/init(device:)
func (bc _BluetoothDevicePairClass) PairWithDevice(device IOBluetoothDevice) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("pairWithDevice:"), device)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PairWithDevice) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BluetoothDevicePair */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BluetoothDevicePair */

// Get the IOBluetoothDevice being used by the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePair/device()
func (b_ BluetoothDevicePair) Device() IBluetoothDevice {
	rv := objc.Send[BluetoothDevice](b_.ID, objc.Sel("device"))
	return rv
}/* debug [instance_methods/method]: Device */


// This is the required reply to the devicePairingPINCodeRequest delegate message. Set the PIN code to use during pairing if required.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePair/replyPINCode(_:pinCode:)
func (b_ BluetoothDevicePair) ReplyPINCodePINCode(PINCodeSize unsafe.Pointer, PINCode objc.IObject /* cross-framework: BluetoothPINCode */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("replyPINCode:PINCode:"), PINCodeSize, PINCode)
}/* debug [instance_methods/method]: ReplyPINCodePINCode */


// This is the required reply to the devicePairingUserConfirmationRequest delegate message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePair/replyUserConfirmation(_:)
func (b_ BluetoothDevicePair) ReplyUserConfirmation(reply bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("replyUserConfirmation:"), reply)
}/* debug [instance_methods/method]: ReplyUserConfirmation */


// Set the device object to pair with. It is retained by the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePair/setDevice(_:)
func (b_ BluetoothDevicePair) SetDevice(inDevice IOBluetoothDevice) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setDevice:"), inDevice)
}/* debug [instance_methods/method]: SetDevice */


// Kicks off the pairing with the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePair/start()
func (b_ BluetoothDevicePair) Start() int {
	rv := objc.Send[int](b_.ID, objc.Sel("start"))
	return rv
}/* debug [instance_methods/method]: Start */


// Stops the current pairing. Removes the delegate and disconnects if device was connected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePair/stop()
func (b_ BluetoothDevicePair) Stop() {
	objc.Send[objc.ID](b_.ID, objc.Sel("stop"))
}/* debug [instance_methods/method]: Stop */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BluetoothDevicePair */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePair/delegate
func (b_ BluetoothDevicePair) Delegate() objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePair/delegate
func (b_ BluetoothDevicePair) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IOBluetoothDevicePair */


