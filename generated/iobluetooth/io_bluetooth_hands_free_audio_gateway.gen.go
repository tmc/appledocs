// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [BluetoothHandsFreeAudioGateway] class.
type IBluetoothHandsFreeAudioGateway interface {
	IBluetoothHandsFree
	// properties:
	// methods:
	CreateIndicatorMinMaxCurrentValue(indicatorName string, minValue int, maxValue int, currentValue int)
	ProcessATCommand(atCommand string)
	SendOKResponse()
	SendResponse(response string)
	SendResponseWithOK(response string, withOK bool)
}

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

// Alloc allocates a new instance without initialization.
func (bc _BluetoothHandsFreeAudioGatewayClass) Alloc() BluetoothHandsFreeAudioGateway {
	rv := objc.Send[BluetoothHandsFreeAudioGateway](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Creates an object that controls a connected Bluetooth hands-free phone or headset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeAudioGateway/init(device:delegate:)
func NewBluetoothHandsFreeAudioGatewayWithDeviceDelegate(device IOBluetoothDevice, inDelegate objectivec.IObject) BluetoothHandsFreeAudioGateway {
	instance := getBluetoothHandsFreeAudioGatewayClass().Alloc()
	rv := objc.Send[BluetoothHandsFreeAudioGateway](instance.ID, objc.Sel("initWithDevice:delegate:"), device, inDelegate)
	rv.Autorelease()
	return rv
}



// Sends a request to the Bluetooth device to show or update a status indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeAudioGateway/createIndicator(_:min:max:currentValue:)
func (b_ BluetoothHandsFreeAudioGateway) CreateIndicatorMinMaxCurrentValue(indicatorName string, minValue int, maxValue int, currentValue int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("createIndicator:min:max:currentValue:"), objc.String(indicatorName), minValue, maxValue, currentValue)
}


// Processes a command from a connected Bluetooth hands-free phone or headset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeAudioGateway/process(atCommand:)
func (b_ BluetoothHandsFreeAudioGateway) ProcessATCommand(atCommand string) {
	objc.Send[objc.ID](b_.ID, objc.Sel("processATCommand:"), objc.String(atCommand))
}


// Sends a success message to a connected Bluetooth hands-free phone or headset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeAudioGateway/sendOKResponse()
func (b_ BluetoothHandsFreeAudioGateway) SendOKResponse() {
	objc.Send[objc.ID](b_.ID, objc.Sel("sendOKResponse"))
}


// Sends data followed by a success message to a connected Bluetooth hands-free phone or headset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeAudioGateway/sendResponse(_:)
func (b_ BluetoothHandsFreeAudioGateway) SendResponse(response string) {
	objc.Send[objc.ID](b_.ID, objc.Sel("sendResponse:"), objc.String(response))
}


// Sends data followed by an optional success message to a connected Bluetooth hands-free phone or headset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeAudioGateway/sendResponse(_:withOK:)
func (b_ BluetoothHandsFreeAudioGateway) SendResponseWithOK(response string, withOK bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("sendResponse:withOK:"), objc.String(response), withOK)
}


