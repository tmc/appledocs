// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class IOBluetoothHandsFreeAudioGateway */


/* debug [class_header]: Header for IOBluetoothHandsFreeAudioGateway */
// The class instance for the [BluetoothHandsFreeAudioGateway] class.
var (
	BluetoothHandsFreeAudioGatewayClass     _BluetoothHandsFreeAudioGatewayClass
	BluetoothHandsFreeAudioGatewayClassOnce sync.Once
)

func getBluetoothHandsFreeAudioGatewayClass() _BluetoothHandsFreeAudioGatewayClass {
	BluetoothHandsFreeAudioGatewayClassOnce.Do(func() {
		BluetoothHandsFreeAudioGatewayClass = _BluetoothHandsFreeAudioGatewayClass{objc.GetClass("IOBluetoothHandsFreeAudioGateway")}
	})
	return BluetoothHandsFreeAudioGatewayClass
}

type _BluetoothHandsFreeAudioGatewayClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BluetoothHandsFreeAudioGateway */
// An interface definition for the [BluetoothHandsFreeAudioGateway] class.
type IBluetoothHandsFreeAudioGateway interface {
	IBluetoothHandsFree
	
/* debug [class_interface_properties]: Properties for BluetoothHandsFreeAudioGateway */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BluetoothHandsFreeAudioGateway */
	// methods:
	CreateIndicatorMinMaxCurrentValue(indicatorName objc.IObject /* cross-framework: NSString */, minValue int, maxValue int, currentValue int)
	ProcessATCommand(atCommand objc.IObject /* cross-framework: NSString */)
	SendOKResponse()
	SendResponse(response objc.IObject /* cross-framework: NSString */)
	SendResponseWithOK(response objc.IObject /* cross-framework: NSString */, withOK bool)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BluetoothHandsFreeAudioGateway */
// Alloc allocates a new instance without initialization.
func (bc _BluetoothHandsFreeAudioGatewayClass) Alloc() BluetoothHandsFreeAudioGateway {
	rv := objc.Send[BluetoothHandsFreeAudioGateway](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BluetoothHandsFreeAudioGatewayClass) New() BluetoothHandsFreeAudioGateway {
	rv := objc.Send[BluetoothHandsFreeAudioGateway](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BluetoothHandsFreeAudioGateway) Init() BluetoothHandsFreeAudioGateway {
	rv := objc.Send[BluetoothHandsFreeAudioGateway](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BluetoothHandsFreeAudioGateway) Autorelease() BluetoothHandsFreeAudioGateway {
	rv := objc.Send[BluetoothHandsFreeAudioGateway](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBluetoothHandsFreeAudioGateway creates a new BluetoothHandsFreeAudioGateway instance.
func NewBluetoothHandsFreeAudioGateway() BluetoothHandsFreeAudioGateway {
	return getBluetoothHandsFreeAudioGatewayClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BluetoothHandsFreeAudioGateway */
// An object that sends data to a connected Bluetooth hands-free phone or headset and processes commands from it.
//
// This class represents the audio gateway portion of a Bluetooth audio profile.


// An object that sends data to a connected Bluetooth hands-free phone or headset and processes commands from it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeAudioGateway
type BluetoothHandsFreeAudioGateway struct {
	BluetoothHandsFree
}

// BluetoothHandsFreeAudioGatewayFrom constructs a [BluetoothHandsFreeAudioGateway] from an unsafe.Pointer.
//
// An object that sends data to a connected Bluetooth hands-free phone or headset and processes commands from it.
func BluetoothHandsFreeAudioGatewayFrom(ptr unsafe.Pointer) BluetoothHandsFreeAudioGateway {
	return BluetoothHandsFreeAudioGateway{
		BluetoothHandsFree: BluetoothHandsFreeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BluetoothHandsFreeAudioGateway */

// Creates an object that controls a connected Bluetooth hands-free phone or headset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeAudioGateway/init(device:delegate:)
func NewBluetoothHandsFreeAudioGatewayWithDeviceDelegate(device IOBluetoothDevice, inDelegate objc.IObject) BluetoothHandsFreeAudioGateway {
	instance := getBluetoothHandsFreeAudioGatewayClass().Alloc()
	rv := objc.Send[BluetoothHandsFreeAudioGateway](instance.ID, objc.Sel("initWithDevice:delegate:"), device, inDelegate)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewBluetoothHandsFreeAudioGatewayWithDeviceDelegate */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BluetoothHandsFreeAudioGateway */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BluetoothHandsFreeAudioGateway */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BluetoothHandsFreeAudioGateway */

// Sends a request to the Bluetooth device to show or update a status indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeAudioGateway/createIndicator(_:min:max:currentValue:)
func (b_ BluetoothHandsFreeAudioGateway) CreateIndicatorMinMaxCurrentValue(indicatorName objc.IObject /* cross-framework: NSString */, minValue int, maxValue int, currentValue int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("createIndicator:min:max:currentValue:"), indicatorName, minValue, maxValue, currentValue)
}/* debug [instance_methods/method]: CreateIndicatorMinMaxCurrentValue */


// Processes a command from a connected Bluetooth hands-free phone or headset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeAudioGateway/process(atCommand:)
func (b_ BluetoothHandsFreeAudioGateway) ProcessATCommand(atCommand objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("processATCommand:"), atCommand)
}/* debug [instance_methods/method]: ProcessATCommand */


// Sends a success message to a connected Bluetooth hands-free phone or headset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeAudioGateway/sendOKResponse()
func (b_ BluetoothHandsFreeAudioGateway) SendOKResponse() {
	objc.Send[objc.ID](b_.ID, objc.Sel("sendOKResponse"))
}/* debug [instance_methods/method]: SendOKResponse */


// Sends data followed by a success message to a connected Bluetooth hands-free phone or headset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeAudioGateway/sendResponse(_:)
func (b_ BluetoothHandsFreeAudioGateway) SendResponse(response objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("sendResponse:"), response)
}/* debug [instance_methods/method]: SendResponse */


// Sends data followed by an optional success message to a connected Bluetooth hands-free phone or headset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeAudioGateway/sendResponse(_:withOK:)
func (b_ BluetoothHandsFreeAudioGateway) SendResponseWithOK(response objc.IObject /* cross-framework: NSString */, withOK bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("sendResponse:withOK:"), response, withOK)
}/* debug [instance_methods/method]: SendResponseWithOK */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BluetoothHandsFreeAudioGateway */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IOBluetoothHandsFreeAudioGateway */


