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
	Connect()
	ConnectSCO()
	Disconnect()
	DisconnectSCO()
	Indicator(indicatorName string) int
	IsSCOConnected() bool
	SetIndicatorValue(indicatorName string, indicatorValue int)
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	Device() IOBluetoothDevice
	DeviceCallHoldModes() uint32
	DeviceSupportedFeatures() uint32
	DeviceSupportedSMSServices() uint32
	InputVolume() float32
	SetInputVolume(value float32)
	Connected() bool
	InputMuted() bool
	SetInputMuted(value bool)
	OutputMuted() bool
	SetOutputMuted(value bool)
	SMSEnabled() bool
	OutputVolume() float32
	SetOutputVolume(value float32)
	SMSMode() BluetoothSMSMode
	SupportedFeatures() uint32
	SetSupportedFeatures(value Iuint32)
	IsConnected() bool
	SetIsConnected(value bool)
	IsInputMuted() bool
	SetIsInputMuted(value bool)
	IsOutputMuted() bool
	SetIsOutputMuted(value bool)
	IsSMSEnabled() bool
	SetIsSMSEnabled(value bool)
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
func NewBluetoothHandsFreeWithDeviceDelegate(device IOBluetoothDevice, inDelegate objectivec.IObject) BluetoothHandsFree {
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
func (b_ BluetoothHandsFree) Indicator(indicatorName string) int {
	rv := objc.Send[int](b_.ID, objc.Sel("indicator:"), objc.String(indicatorName))
	return rv
}


// Determine if there is a SCO connection to the device
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/isSCOConnected()
func (b_ BluetoothHandsFree) IsSCOConnected() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isSCOConnected"))
	return rv
}


// Set an indicator’s value
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/setIndicator(_:value:)
func (b_ BluetoothHandsFree) SetIndicatorValue(indicatorName string, indicatorValue int) {
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
func (b_ BluetoothHandsFree) Device() IOBluetoothDevice {
	rv := objc.Send[IOBluetoothDevice](b_.ID, objc.Sel("device"))
	return rv
}


// Return the device’s supported call hold modes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/deviceCallHoldModes
func (b_ BluetoothHandsFree) DeviceCallHoldModes() uint32 {
	rv := objc.Send[uint32](b_.ID, objc.Sel("deviceCallHoldModes"))
	return rv
}


// Return the device’s supported features.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/deviceSupportedFeatures
func (b_ BluetoothHandsFree) DeviceSupportedFeatures() uint32 {
	rv := objc.Send[uint32](b_.ID, objc.Sel("deviceSupportedFeatures"))
	return rv
}


// Return the device’s supported SMS services.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/deviceSupportedSMSServices
func (b_ BluetoothHandsFree) DeviceSupportedSMSServices() uint32 {
	rv := objc.Send[uint32](b_.ID, objc.Sel("deviceSupportedSMSServices"))
	return rv
}


// Return the input volume
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/inputVolume
func (b_ BluetoothHandsFree) InputVolume() float32 {
	rv := objc.Send[float32](b_.ID, objc.Sel("inputVolume"))
	return rv
}


// Return the input volume
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/inputVolume
func (b_ BluetoothHandsFree) SetInputVolume(value float32) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setInputVolume:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/isConnected
func (b_ BluetoothHandsFree) Connected() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("connected"))
	return rv
}


// Return the input mute state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/isInputMuted
func (b_ BluetoothHandsFree) InputMuted() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("inputMuted"))
	return rv
}


// Return the input mute state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/isInputMuted
func (b_ BluetoothHandsFree) SetInputMuted(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setInputMuted:"), value)
}


// Return the output mute state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/isOutputMuted
func (b_ BluetoothHandsFree) OutputMuted() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("outputMuted"))
	return rv
}


// Return the output mute state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/isOutputMuted
func (b_ BluetoothHandsFree) SetOutputMuted(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setOutputMuted:"), value)
}


// Return YES if the device has SMS enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/isSMSEnabled
func (b_ BluetoothHandsFree) SMSEnabled() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("SMSEnabled"))
	return rv
}


// Return the output volume
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/outputVolume
func (b_ BluetoothHandsFree) OutputVolume() float32 {
	rv := objc.Send[float32](b_.ID, objc.Sel("outputVolume"))
	return rv
}


// Return the output volume
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/outputVolume
func (b_ BluetoothHandsFree) SetOutputVolume(value float32) {
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
func (b_ BluetoothHandsFree) SupportedFeatures() uint32 {
	rv := objc.Send[uint32](b_.ID, objc.Sel("supportedFeatures"))
	return rv
}


// Set the supported features
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/supportedFeatures
func (b_ BluetoothHandsFree) SetSupportedFeatures(value Iuint32) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSupportedFeatures:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothhandsfree/isconnected
func (b_ BluetoothHandsFree) IsConnected() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isConnected"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothhandsfree/isconnected
func (b_ BluetoothHandsFree) SetIsConnected(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsConnected:"), value)
}


// Return the input mute state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothhandsfree/isinputmuted
func (b_ BluetoothHandsFree) IsInputMuted() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isInputMuted"))
	return rv
}


// Return the input mute state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothhandsfree/isinputmuted
func (b_ BluetoothHandsFree) SetIsInputMuted(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsInputMuted:"), value)
}


// Return the output mute state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothhandsfree/isoutputmuted
func (b_ BluetoothHandsFree) IsOutputMuted() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isOutputMuted"))
	return rv
}


// Return the output mute state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothhandsfree/isoutputmuted
func (b_ BluetoothHandsFree) SetIsOutputMuted(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsOutputMuted:"), value)
}


// Return YES if the device has SMS enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothhandsfree/issmsenabled
func (b_ BluetoothHandsFree) IsSMSEnabled() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isSMSEnabled"))
	return rv
}


// Return YES if the device has SMS enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothhandsfree/issmsenabled
func (b_ BluetoothHandsFree) SetIsSMSEnabled(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsSMSEnabled:"), value)
}


