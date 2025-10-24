// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class IOBluetoothHostController */


/* debug [class_header]: Header for IOBluetoothHostController */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BluetoothHostController */
// An interface definition for the [BluetoothHostController] class.
type IBluetoothHostController interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for BluetoothHostController */
	// properties:
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	PowerState() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BluetoothHostController */
	// methods:
	AddressAsString() foundation.String
	ClassOfDevice() BluetoothClassOfDevice /* typedef */
	NameAsString() foundation.String
	SetClassOfDeviceForTimeInterval(classOfDevice BluetoothClassOfDevice /* typedef */, seconds float64) int
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BluetoothHostController */
// Alloc allocates a new instance without initialization.
func (bc _BluetoothHostControllerClass) Alloc() BluetoothHostController {
	rv := objc.Send[BluetoothHostController](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BluetoothHostController */
// This class is a representation of a Bluetooth Host Controller Interface that is present on the local computer (either plugged in externally or available internally).
//
// This object can be used to ask a Bluetooth HCI for certain pieces of information, and be used to make it perform certain functions.


// This class is a representation of a Bluetooth Host Controller Interface that is present on the local computer (either plugged in externally or available internally).
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BluetoothHostController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BluetoothHostController */

// Gets the default HCI controller object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHostController/default()
func (bc _BluetoothHostControllerClass) DefaultController() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("defaultController"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DefaultController) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BluetoothHostController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BluetoothHostController */

// Convience routine to get the HCI controller’s Bluetooth address as an NSString object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHostController/addressAsString()
func (b_ BluetoothHostController) AddressAsString() foundation.String {
	rv := objc.Send[foundation.String](b_.ID, objc.Sel("addressAsString"))
	return rv
}/* debug [instance_methods/method]: AddressAsString */


// Gets the current class of device value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHostController/classOfDevice()
func (b_ BluetoothHostController) ClassOfDevice() BluetoothClassOfDevice /* typedef */ {
	rv := objc.Send[uint32](b_.ID, objc.Sel("classOfDevice"))
	return rv
}/* debug [instance_methods/method]: ClassOfDevice */


// Gets the “friendly” name of HCI controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHostController/nameAsString()
func (b_ BluetoothHostController) NameAsString() foundation.String {
	rv := objc.Send[foundation.String](b_.ID, objc.Sel("nameAsString"))
	return rv
}/* debug [instance_methods/method]: NameAsString */


// Sets the current class of device value, for the specified amount of time. Note that the time interval be set and valid. The range of acceptable values is 30-120 seconds. Anything above or below will be rounded up, or down, as appropriate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHostController/setClassOfDevice(_:forTimeInterval:)
func (b_ BluetoothHostController) SetClassOfDeviceForTimeInterval(classOfDevice BluetoothClassOfDevice /* typedef */, seconds float64) int {
	rv := objc.Send[int](b_.ID, objc.Sel("setClassOfDevice:forTimeInterval:"), classOfDevice, seconds)
	return rv
}/* debug [instance_methods/method]: SetClassOfDeviceForTimeInterval */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BluetoothHostController */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHostController/delegate
func (b_ BluetoothHostController) Delegate() objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHostController/delegate
func (b_ BluetoothHostController) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// Gets the controller power state
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHostController/powerState
func (b_ BluetoothHostController) PowerState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("powerState"))
	return rv
}/* debug [instance_properties/getter]: powerState */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IOBluetoothHostController */



