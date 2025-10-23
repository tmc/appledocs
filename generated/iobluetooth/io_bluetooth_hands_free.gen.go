// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BluetoothHandsFree] class.
var (
	BluetoothHandsFreeClass     _BluetoothHandsFreeClass
	BluetoothHandsFreeClassOnce sync.Once
)

func getBluetoothHandsFreeClass() _BluetoothHandsFreeClass {
	BluetoothHandsFreeClassOnce.Do(func() {
		BluetoothHandsFreeClass = _BluetoothHandsFreeClass{objc.GetClass("IOBluetoothHandsFree")}
	})
	return BluetoothHandsFreeClass
}

type _BluetoothHandsFreeClass struct {
	class objc.Class
}

// An interface definition for the [BluetoothHandsFree] class.
type IBluetoothHandsFree interface {
	objectivec.IObject
	// properties:
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	Device() IOBluetoothDevice /* already interface */
	DeviceCallHoldModes() uint32 /* not a class type */
	DeviceSupportedFeatures() uint32 /* not a class type */
	DeviceSupportedSMSServices() uint32 /* not a class type */
	InputVolume() float32 /* primitive/slice/pointer. */
	SetInputVolume(value float32 /* primitive/slice/pointer. */)
	Connected() bool /* primitive/slice/pointer. */
	InputMuted() bool /* primitive/slice/pointer. */
	SetInputMuted(value bool /* primitive/slice/pointer. */)
	OutputMuted() bool /* primitive/slice/pointer. */
	SetOutputMuted(value bool /* primitive/slice/pointer. */)
	SMSEnabled() bool /* primitive/slice/pointer. */
	OutputVolume() float32 /* primitive/slice/pointer. */
	SetOutputVolume(value float32 /* primitive/slice/pointer. */)
	SMSMode() BluetoothSMSMode
	SupportedFeatures() uint32 /* not a class type */
	SetSupportedFeatures(value uint32 /* not a class type */)
	IsConnected() bool /* primitive/slice/pointer. */
	SetIsConnected(value bool /* primitive/slice/pointer. */)
	IsInputMuted() bool /* primitive/slice/pointer. */
	SetIsInputMuted(value bool /* primitive/slice/pointer. */)
	IsOutputMuted() bool /* primitive/slice/pointer. */
	SetIsOutputMuted(value bool /* primitive/slice/pointer. */)
	IsSMSEnabled() bool /* primitive/slice/pointer. */
	SetIsSMSEnabled(value bool /* primitive/slice/pointer. */)
	// methods:
	Connect()
	ConnectSCO()
	Disconnect()
	DisconnectSCO()
	Indicator(indicatorName string /* primitive/slice/pointer. */) int /* primitive/slice/pointer. */
	IsSCOConnected() bool /* primitive/slice/pointer. */
	SetIndicatorValue(indicatorName string /* primitive/slice/pointer. */, indicatorValue int /* primitive/slice/pointer. */)
}

// Hands free profile class.
//
// Superclass of IOBluetoothHandsFreeDevice and IOBluetoothHandsFreeAudioGateway classes. Contains the common code used to support the bluetoooth hands free profile.


// Hands free profile class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree
type BluetoothHandsFree struct {
	objectivec.Object
}

// BluetoothHandsFreeFrom constructs a [BluetoothHandsFree] from an unsafe.Pointer.
//
// Hands free profile class.
func BluetoothHandsFreeFrom(ptr unsafe.Pointer) BluetoothHandsFree {
	return BluetoothHandsFree{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (bc _BluetoothHandsFreeClass) Alloc() BluetoothHandsFree {
	rv := objc.Send[BluetoothHandsFree](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BluetoothHandsFreeClass) New() BluetoothHandsFree {
	rv := objc.Send[BluetoothHandsFree](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BluetoothHandsFree) Init() BluetoothHandsFree {
	rv := objc.Send[BluetoothHandsFree](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BluetoothHandsFree) Autorelease() BluetoothHandsFree {
	rv := objc.Send[BluetoothHandsFree](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBluetoothHandsFree creates a new BluetoothHandsFree instance.
func NewBluetoothHandsFree() BluetoothHandsFree {
	return getBluetoothHandsFreeClass().New()
}



// Create a new IOBluetoothHandsFree object
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/init(device:delegate:)
func NewBluetoothHandsFreeWithDeviceDelegate(device BluetoothDevice /* already interface */, inDelegate objectivec.IObject) BluetoothHandsFree {
	instance := getBluetoothHandsFreeClass().Alloc()
	rv := objc.Send[BluetoothHandsFree](instance.ID, objc.Sel("initWithDevice:delegate:"), device, inDelegate)
	rv.Autorelease()
	return rv
}



// Connect to the device
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/connect()
func (b_ BluetoothHandsFree) Connect() {
	objc.Send[objc.ID](b_.ID, objc.Sel("connect"))
}


// Open a SCO connection with the device
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/connectSCO()
func (b_ BluetoothHandsFree) ConnectSCO() {
	objc.Send[objc.ID](b_.ID, objc.Sel("connectSCO"))
}


// Disconnect from the device
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/disconnect()
func (b_ BluetoothHandsFree) Disconnect() {
	objc.Send[objc.ID](b_.ID, objc.Sel("disconnect"))
}


// Disconnect the SCO connection with the device
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/disconnectSCO()
func (b_ BluetoothHandsFree) DisconnectSCO() {
	objc.Send[objc.ID](b_.ID, objc.Sel("disconnectSCO"))
}


// Return an indicator’s value
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/indicator(_:)
func (b_ BluetoothHandsFree) Indicator(indicatorName string /* primitive/slice/pointer. */) int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](b_.ID, objc.Sel("indicator:"), objc.String(indicatorName))
	return rv
}


// Determine if there is a SCO connection to the device
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/isSCOConnected()
func (b_ BluetoothHandsFree) IsSCOConnected() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("isSCOConnected"))
	return rv
}


// Set an indicator’s value
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/setIndicator(_:value:)
func (b_ BluetoothHandsFree) SetIndicatorValue(indicatorName string /* primitive/slice/pointer. */, indicatorValue int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIndicator:value:"), objc.String(indicatorName), indicatorValue)
}


// Return the delegate
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/delegate
func (b_ BluetoothHandsFree) Delegate() objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("delegate"))
	return rv
}


// Return the delegate
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/delegate
func (b_ BluetoothHandsFree) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setDelegate:"), value)
}


// Return the IOBluetoothDevice.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/device
func (b_ BluetoothHandsFree) Device() IOBluetoothDevice /* already interface */ {
	rv := objc.Send[BluetoothDevice](b_.ID, objc.Sel("device"))
	return rv
}


// Return the device’s supported call hold modes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/deviceCallHoldModes
func (b_ BluetoothHandsFree) DeviceCallHoldModes() uint32 /* not a class type */ {
	rv := objc.Send[uint32](b_.ID, objc.Sel("deviceCallHoldModes"))
	return rv
}


// Return the device’s supported features.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/deviceSupportedFeatures
func (b_ BluetoothHandsFree) DeviceSupportedFeatures() uint32 /* not a class type */ {
	rv := objc.Send[uint32](b_.ID, objc.Sel("deviceSupportedFeatures"))
	return rv
}


// Return the device’s supported SMS services.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/deviceSupportedSMSServices
func (b_ BluetoothHandsFree) DeviceSupportedSMSServices() uint32 /* not a class type */ {
	rv := objc.Send[uint32](b_.ID, objc.Sel("deviceSupportedSMSServices"))
	return rv
}


// Return the input volume
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/inputVolume
func (b_ BluetoothHandsFree) InputVolume() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](b_.ID, objc.Sel("inputVolume"))
	return rv
}


// Return the input volume
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/inputVolume
func (b_ BluetoothHandsFree) SetInputVolume(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setInputVolume:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/isConnected
func (b_ BluetoothHandsFree) Connected() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("connected"))
	return rv
}


// Return the input mute state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/isInputMuted
func (b_ BluetoothHandsFree) InputMuted() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("inputMuted"))
	return rv
}


// Return the input mute state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/isInputMuted
func (b_ BluetoothHandsFree) SetInputMuted(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setInputMuted:"), value)
}


// Return the output mute state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/isOutputMuted
func (b_ BluetoothHandsFree) OutputMuted() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("outputMuted"))
	return rv
}


// Return the output mute state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/isOutputMuted
func (b_ BluetoothHandsFree) SetOutputMuted(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setOutputMuted:"), value)
}


// Return YES if the device has SMS enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/isSMSEnabled
func (b_ BluetoothHandsFree) SMSEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("SMSEnabled"))
	return rv
}


// Return the output volume
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/outputVolume
func (b_ BluetoothHandsFree) OutputVolume() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](b_.ID, objc.Sel("outputVolume"))
	return rv
}


// Return the output volume
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/outputVolume
func (b_ BluetoothHandsFree) SetOutputVolume(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setOutputVolume:"), value)
}


// Return the device’s SMS mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/smsMode
func (b_ BluetoothHandsFree) SMSMode() BluetoothSMSMode {
	rv := objc.Send[BluetoothSMSMode](b_.ID, objc.Sel("SMSMode"))
	return rv
}


// Set the supported features
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/supportedFeatures
func (b_ BluetoothHandsFree) SupportedFeatures() uint32 /* not a class type */ {
	rv := objc.Send[uint32](b_.ID, objc.Sel("supportedFeatures"))
	return rv
}


// Set the supported features
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/supportedFeatures
func (b_ BluetoothHandsFree) SetSupportedFeatures(value uint32 /* not a class type */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSupportedFeatures:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothhandsfree/isconnected
func (b_ BluetoothHandsFree) IsConnected() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("isConnected"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothhandsfree/isconnected
func (b_ BluetoothHandsFree) SetIsConnected(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsConnected:"), value)
}


// Return the input mute state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothhandsfree/isinputmuted
func (b_ BluetoothHandsFree) IsInputMuted() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("isInputMuted"))
	return rv
}


// Return the input mute state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothhandsfree/isinputmuted
func (b_ BluetoothHandsFree) SetIsInputMuted(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsInputMuted:"), value)
}


// Return the output mute state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothhandsfree/isoutputmuted
func (b_ BluetoothHandsFree) IsOutputMuted() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("isOutputMuted"))
	return rv
}


// Return the output mute state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothhandsfree/isoutputmuted
func (b_ BluetoothHandsFree) SetIsOutputMuted(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsOutputMuted:"), value)
}


// Return YES if the device has SMS enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothhandsfree/issmsenabled
func (b_ BluetoothHandsFree) IsSMSEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("isSMSEnabled"))
	return rv
}


// Return YES if the device has SMS enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothhandsfree/issmsenabled
func (b_ BluetoothHandsFree) SetIsSMSEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsSMSEnabled:"), value)
}


