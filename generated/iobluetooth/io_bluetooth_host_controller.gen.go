// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BluetoothHostController] class.
var (
	BluetoothHostControllerClass     _BluetoothHostControllerClass
	BluetoothHostControllerClassOnce sync.Once
)

func getBluetoothHostControllerClass() _BluetoothHostControllerClass {
	BluetoothHostControllerClassOnce.Do(func() {
		BluetoothHostControllerClass = _BluetoothHostControllerClass{objc.GetClass("IOBluetoothHostController")}
	})
	return BluetoothHostControllerClass
}

type _BluetoothHostControllerClass struct {
	class objc.Class
}

// An interface definition for the [BluetoothHostController] class.
type IBluetoothHostController interface {
	objectivec.IObject
	AddressAsString() string
	ClassOfDevice() unsafe.Pointer
	NameAsString() string
	SetClassOfDeviceForTimeInterval(classOfDevice unsafe.Pointer, seconds foundation.TimeInterval) unsafe.Pointer
}

// This class is a representation of a Bluetooth Host Controller Interface that is present on the local computer (either plugged in externally or available internally).
//
// This object can be used to ask a Bluetooth HCI for certain pieces of information, and be used to make it perform certain functions.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHostController
type BluetoothHostController struct {
	objectivec.Object
}

// BluetoothHostControllerFrom constructs a [BluetoothHostController] from an unsafe.Pointer.
//
// This class is a representation of a Bluetooth Host Controller Interface that is present on the local computer (either plugged in externally or available internally).
func BluetoothHostControllerFrom(ptr unsafe.Pointer) BluetoothHostController {
	return BluetoothHostController{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (bc _BluetoothHostControllerClass) Alloc() BluetoothHostController {
	rv := objc.Send[BluetoothHostController](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BluetoothHostControllerClass) New() BluetoothHostController {
	rv := objc.Send[BluetoothHostController](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BluetoothHostController) Init() BluetoothHostController {
	rv := objc.Send[BluetoothHostController](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BluetoothHostController) Autorelease() BluetoothHostController {
	rv := objc.Send[BluetoothHostController](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBluetoothHostController creates a new BluetoothHostController instance.
func NewBluetoothHostController() BluetoothHostController {
	return getBluetoothHostControllerClass().New()
}


// Gets the default HCI controller object.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHostController/default()
func (bc _BluetoothHostControllerClass) DefaultController() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("defaultController"))
	return rv
}

// Convience routine to get the HCI controller’s Bluetooth address as an NSString object.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHostController/addressAsString()
func (b_ BluetoothHostController) AddressAsString() string {
	rv := objc.Send[string](b_.ID, objc.Sel("addressAsString"))
	return rv
}

// Gets the current class of device value.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHostController/classOfDevice()
func (b_ BluetoothHostController) ClassOfDevice() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("classOfDevice"))
	return rv
}

// Gets the “friendly” name of HCI controller.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHostController/nameAsString()
func (b_ BluetoothHostController) NameAsString() string {
	rv := objc.Send[string](b_.ID, objc.Sel("nameAsString"))
	return rv
}

// Sets the current class of device value, for the specified amount of time. Note that the time interval be set and valid. The range of acceptable values is 30-120 seconds. Anything above or below will be rounded up, or down, as appropriate.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHostController/setClassOfDevice(_:forTimeInterval:)
func (b_ BluetoothHostController) SetClassOfDeviceForTimeInterval(classOfDevice unsafe.Pointer, seconds foundation.TimeInterval) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("setClassOfDevice:forTimeInterval:"), classOfDevice, seconds)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHostController/delegate
func (b_ BluetoothHostController) Delegate() objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHostController/delegate
func (b_ BluetoothHostController) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setDelegate:"), value)
}

// Gets the controller power state
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHostController/powerState
func (b_ BluetoothHostController) PowerState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("powerState"))
	return rv
}



