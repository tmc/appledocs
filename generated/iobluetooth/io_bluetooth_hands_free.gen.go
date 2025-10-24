// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class IOBluetoothHandsFree */


/* debug [class_header]: Header for IOBluetoothHandsFree */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BluetoothHandsFree */
// An interface definition for the [BluetoothHandsFree] class.
type IBluetoothHandsFree interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for BluetoothHandsFree */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	Device() IOBluetoothDevice
	DeviceCallHoldModes() uint32 /* not a class type */
	DeviceSupportedFeatures() uint32 /* not a class type */
	DeviceSupportedSMSServices() uint32 /* not a class type */
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
	SupportedFeatures() uint32 /* not a class type */
	SetSupportedFeatures(value uint32 /* not a class type */)
	IsConnected() bool
	SetIsConnected(value bool)
	IsInputMuted() bool
	SetIsInputMuted(value bool)
	IsOutputMuted() bool
	SetIsOutputMuted(value bool)
	IsSMSEnabled() bool
	SetIsSMSEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BluetoothHandsFree */
	// methods:
	Connect()
	ConnectSCO()
	Disconnect()
	DisconnectSCO()
	Indicator(indicatorName objc.IObject /* cross-framework: NSString */) int
	IsSCOConnected() bool
	SetIndicatorValue(indicatorName objc.IObject /* cross-framework: NSString */, indicatorValue int)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BluetoothHandsFree */
// Alloc allocates a new instance without initialization.
func (bc _BluetoothHandsFreeClass) Alloc() BluetoothHandsFree {
	rv := objc.Send[BluetoothHandsFree](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BluetoothHandsFree */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BluetoothHandsFree */

// Create a new IOBluetoothHandsFree object
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/init(device:delegate:)
func NewBluetoothHandsFreeWithDeviceDelegate(device IOBluetoothDevice, inDelegate unsafe.Pointer) BluetoothHandsFree {
	instance := getBluetoothHandsFreeClass().Alloc()
	rv := objc.Send[BluetoothHandsFree](instance.ID, objc.Sel("initWithDevice:delegate:"), device, inDelegate)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewBluetoothHandsFreeWithDeviceDelegate */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BluetoothHandsFree */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BluetoothHandsFree */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BluetoothHandsFree */

// Connect to the device
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/connect()
func (b_ BluetoothHandsFree) Connect() {
	objc.Send[objc.ID](b_.ID, objc.Sel("connect"))
}/* debug [instance_methods/method]: Connect */


// Open a SCO connection with the device
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/connectSCO()
func (b_ BluetoothHandsFree) ConnectSCO() {
	objc.Send[objc.ID](b_.ID, objc.Sel("connectSCO"))
}/* debug [instance_methods/method]: ConnectSCO */


// Disconnect from the device
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/disconnect()
func (b_ BluetoothHandsFree) Disconnect() {
	objc.Send[objc.ID](b_.ID, objc.Sel("disconnect"))
}/* debug [instance_methods/method]: Disconnect */


// Disconnect the SCO connection with the device
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/disconnectSCO()
func (b_ BluetoothHandsFree) DisconnectSCO() {
	objc.Send[objc.ID](b_.ID, objc.Sel("disconnectSCO"))
}/* debug [instance_methods/method]: DisconnectSCO */


// Return an indicator’s value
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/indicator(_:)
func (b_ BluetoothHandsFree) Indicator(indicatorName objc.IObject /* cross-framework: NSString */) int {
	rv := objc.Send[int](b_.ID, objc.Sel("indicator:"), indicatorName)
	return rv
}/* debug [instance_methods/method]: Indicator */


// Determine if there is a SCO connection to the device
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/isSCOConnected()
func (b_ BluetoothHandsFree) IsSCOConnected() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isSCOConnected"))
	return rv
}/* debug [instance_methods/method]: IsSCOConnected */


// Set an indicator’s value
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/setIndicator(_:value:)
func (b_ BluetoothHandsFree) SetIndicatorValue(indicatorName objc.IObject /* cross-framework: NSString */, indicatorValue int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIndicator:value:"), indicatorName, indicatorValue)
}/* debug [instance_methods/method]: SetIndicatorValue */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BluetoothHandsFree */

// Return the delegate
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/delegate
func (b_ BluetoothHandsFree) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// Return the delegate
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/delegate
func (b_ BluetoothHandsFree) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// Return the IOBluetoothDevice.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/device
func (b_ BluetoothHandsFree) Device() IOBluetoothDevice {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("device"))
	return rv
}/* debug [instance_properties/getter]: device */


// Return the device’s supported call hold modes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/deviceCallHoldModes
func (b_ BluetoothHandsFree) DeviceCallHoldModes() uint32 /* not a class type */ {
	rv := objc.Send[uint32](b_.ID, objc.Sel("deviceCallHoldModes"))
	return rv
}/* debug [instance_properties/getter]: deviceCallHoldModes */


// Return the device’s supported features.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/deviceSupportedFeatures
func (b_ BluetoothHandsFree) DeviceSupportedFeatures() uint32 /* not a class type */ {
	rv := objc.Send[uint32](b_.ID, objc.Sel("deviceSupportedFeatures"))
	return rv
}/* debug [instance_properties/getter]: deviceSupportedFeatures */


// Return the device’s supported SMS services.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/deviceSupportedSMSServices
func (b_ BluetoothHandsFree) DeviceSupportedSMSServices() uint32 /* not a class type */ {
	rv := objc.Send[uint32](b_.ID, objc.Sel("deviceSupportedSMSServices"))
	return rv
}/* debug [instance_properties/getter]: deviceSupportedSMSServices */


// Return the input volume
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/inputVolume
func (b_ BluetoothHandsFree) InputVolume() float32 {
	rv := objc.Send[float32](b_.ID, objc.Sel("inputVolume"))
	return rv
}/* debug [instance_properties/getter]: inputVolume */


// Return the input volume
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/inputVolume
func (b_ BluetoothHandsFree) SetInputVolume(value float32) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setInputVolume:"), value)
}/* debug [instance_properties/setter]: inputVolume */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/isConnected
func (b_ BluetoothHandsFree) Connected() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("connected"))
	return rv
}/* debug [instance_properties/getter]: connected */


// Return the input mute state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/isInputMuted
func (b_ BluetoothHandsFree) InputMuted() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("inputMuted"))
	return rv
}/* debug [instance_properties/getter]: inputMuted */


// Return the input mute state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/isInputMuted
func (b_ BluetoothHandsFree) SetInputMuted(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setInputMuted:"), value)
}/* debug [instance_properties/setter]: inputMuted */


// Return the output mute state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/isOutputMuted
func (b_ BluetoothHandsFree) OutputMuted() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("outputMuted"))
	return rv
}/* debug [instance_properties/getter]: outputMuted */


// Return the output mute state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/isOutputMuted
func (b_ BluetoothHandsFree) SetOutputMuted(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setOutputMuted:"), value)
}/* debug [instance_properties/setter]: outputMuted */


// Return YES if the device has SMS enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/isSMSEnabled
func (b_ BluetoothHandsFree) SMSEnabled() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("SMSEnabled"))
	return rv
}/* debug [instance_properties/getter]: SMSEnabled */


// Return the output volume
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/outputVolume
func (b_ BluetoothHandsFree) OutputVolume() float32 {
	rv := objc.Send[float32](b_.ID, objc.Sel("outputVolume"))
	return rv
}/* debug [instance_properties/getter]: outputVolume */


// Return the output volume
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/outputVolume
func (b_ BluetoothHandsFree) SetOutputVolume(value float32) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setOutputVolume:"), value)
}/* debug [instance_properties/setter]: outputVolume */


// Return the device’s SMS mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/smsMode
func (b_ BluetoothHandsFree) SMSMode() BluetoothSMSMode {
	rv := objc.Send[BluetoothSMSMode](b_.ID, objc.Sel("SMSMode"))
	return rv
}/* debug [instance_properties/getter]: SMSMode */


// Set the supported features
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/supportedFeatures
func (b_ BluetoothHandsFree) SupportedFeatures() uint32 /* not a class type */ {
	rv := objc.Send[uint32](b_.ID, objc.Sel("supportedFeatures"))
	return rv
}/* debug [instance_properties/getter]: supportedFeatures */


// Set the supported features
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/supportedFeatures
func (b_ BluetoothHandsFree) SetSupportedFeatures(value uint32 /* not a class type */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSupportedFeatures:"), value)
}/* debug [instance_properties/setter]: supportedFeatures */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothhandsfree/isconnected
func (b_ BluetoothHandsFree) IsConnected() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isConnected"))
	return rv
}/* debug [instance_properties/getter]: isConnected */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothhandsfree/isconnected
func (b_ BluetoothHandsFree) SetIsConnected(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsConnected:"), value)
}/* debug [instance_properties/setter]: isConnected */


// Return the input mute state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothhandsfree/isinputmuted
func (b_ BluetoothHandsFree) IsInputMuted() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isInputMuted"))
	return rv
}/* debug [instance_properties/getter]: isInputMuted */


// Return the input mute state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothhandsfree/isinputmuted
func (b_ BluetoothHandsFree) SetIsInputMuted(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsInputMuted:"), value)
}/* debug [instance_properties/setter]: isInputMuted */


// Return the output mute state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothhandsfree/isoutputmuted
func (b_ BluetoothHandsFree) IsOutputMuted() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isOutputMuted"))
	return rv
}/* debug [instance_properties/getter]: isOutputMuted */


// Return the output mute state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothhandsfree/isoutputmuted
func (b_ BluetoothHandsFree) SetIsOutputMuted(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsOutputMuted:"), value)
}/* debug [instance_properties/setter]: isOutputMuted */


// Return YES if the device has SMS enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothhandsfree/issmsenabled
func (b_ BluetoothHandsFree) IsSMSEnabled() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isSMSEnabled"))
	return rv
}/* debug [instance_properties/getter]: isSMSEnabled */


// Return YES if the device has SMS enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/iobluetoothhandsfree/issmsenabled
func (b_ BluetoothHandsFree) SetIsSMSEnabled(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsSMSEnabled:"), value)
}/* debug [instance_properties/setter]: isSMSEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IOBluetoothHandsFree */


