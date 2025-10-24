// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class IOBluetoothHandsFreeDevice */


/* debug [class_header]: Header for IOBluetoothHandsFreeDevice */
// The class instance for the [BluetoothHandsFreeDevice] class.
var (
	BluetoothHandsFreeDeviceClass     _BluetoothHandsFreeDeviceClass
	BluetoothHandsFreeDeviceClassOnce sync.Once
)

func getBluetoothHandsFreeDeviceClass() _BluetoothHandsFreeDeviceClass {
	BluetoothHandsFreeDeviceClassOnce.Do(func() {
		BluetoothHandsFreeDeviceClass = _BluetoothHandsFreeDeviceClass{objc.GetClass("IOBluetoothHandsFreeDevice")}
	})
	return BluetoothHandsFreeDeviceClass
}

type _BluetoothHandsFreeDeviceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BluetoothHandsFreeDevice */
// An interface definition for the [BluetoothHandsFreeDevice] class.
type IBluetoothHandsFreeDevice interface {
	IBluetoothHandsFree
	
/* debug [class_interface_properties]: Properties for BluetoothHandsFreeDevice */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BluetoothHandsFreeDevice */
	// methods:
	AcceptCall()
	AcceptCallOnPhone()
	AddHeldCall()
	CallTransfer()
	CurrentCallList()
	DialNumber(aNumber objc.IObject /* cross-framework: NSString */)
	EndCall()
	HoldCall()
	MemoryDial(memoryLocation int)
	PlaceAllOthersOnHold(index int)
	Redial()
	ReleaseActiveCalls()
	ReleaseCall(index int)
	ReleaseHeldCalls()
	SendATCommand(atCommand objc.IObject /* cross-framework: NSString */)
	SendATCommandTimeoutSelectorTarget(atCommand objc.IObject /* cross-framework: NSString */, timeout float32, selector objc.SEL, target objc.IObject)
	SendDTMF(character objc.IObject /* cross-framework: NSString */)
	SendSMSMessage(aNumber objc.IObject /* cross-framework: NSString */, aMessage objc.IObject /* cross-framework: NSString */)
	SubscriberNumber()
	TransferAudioToComputer()
	TransferAudioToPhone()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BluetoothHandsFreeDevice */
// Alloc allocates a new instance without initialization.
func (bc _BluetoothHandsFreeDeviceClass) Alloc() BluetoothHandsFreeDevice {
	rv := objc.Send[BluetoothHandsFreeDevice](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BluetoothHandsFreeDeviceClass) New() BluetoothHandsFreeDevice {
	rv := objc.Send[BluetoothHandsFreeDevice](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BluetoothHandsFreeDevice) Init() BluetoothHandsFreeDevice {
	rv := objc.Send[BluetoothHandsFreeDevice](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BluetoothHandsFreeDevice) Autorelease() BluetoothHandsFreeDevice {
	rv := objc.Send[BluetoothHandsFreeDevice](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBluetoothHandsFreeDevice creates a new BluetoothHandsFreeDevice instance.
func NewBluetoothHandsFreeDevice() BluetoothHandsFreeDevice {
	return getBluetoothHandsFreeDeviceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BluetoothHandsFreeDevice */
// An object you use to manage phone calls on a connected Bluetooth hands-free phone or headset.


// An object you use to manage phone calls on a connected Bluetooth hands-free phone or headset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice
type BluetoothHandsFreeDevice struct {
	BluetoothHandsFree
}

// BluetoothHandsFreeDeviceFrom constructs a [BluetoothHandsFreeDevice] from an unsafe.Pointer.
//
// An object you use to manage phone calls on a connected Bluetooth hands-free phone or headset.
func BluetoothHandsFreeDeviceFrom(ptr unsafe.Pointer) BluetoothHandsFreeDevice {
	return BluetoothHandsFreeDevice{
		BluetoothHandsFree: BluetoothHandsFreeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BluetoothHandsFreeDevice */

// Creates an object to manage phone calls on a hands-free Bluetooth device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/init(device:delegate:)
func NewBluetoothHandsFreeDeviceWithDeviceDelegate(device IOBluetoothDevice, delegate objc.IObject) BluetoothHandsFreeDevice {
	instance := getBluetoothHandsFreeDeviceClass().Alloc()
	rv := objc.Send[BluetoothHandsFreeDevice](instance.ID, objc.Sel("initWithDevice:delegate:"), device, delegate)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewBluetoothHandsFreeDeviceWithDeviceDelegate */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BluetoothHandsFreeDevice */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BluetoothHandsFreeDevice */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BluetoothHandsFreeDevice */

// Accepts an incoming call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/acceptCall()
func (b_ BluetoothHandsFreeDevice) AcceptCall() {
	objc.Send[objc.ID](b_.ID, objc.Sel("acceptCall"))
}/* debug [instance_methods/method]: AcceptCall */


// Accepts an incoming call and transfers the audio to the managed hands-free phone or headset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/acceptCallOnPhone()
func (b_ BluetoothHandsFreeDevice) AcceptCallOnPhone() {
	objc.Send[objc.ID](b_.ID, objc.Sel("acceptCallOnPhone"))
}/* debug [instance_methods/method]: AcceptCallOnPhone */


// Adds held calls to the current conversation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/addHeldCall()
func (b_ BluetoothHandsFreeDevice) AddHeldCall() {
	objc.Send[objc.ID](b_.ID, objc.Sel("addHeldCall"))
}/* debug [instance_methods/method]: AddHeldCall */


// Ends all calls that are active or on hold, and accepts any waiting calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/callTransfer()
func (b_ BluetoothHandsFreeDevice) CallTransfer() {
	objc.Send[objc.ID](b_.ID, objc.Sel("callTransfer"))
}/* debug [instance_methods/method]: CallTransfer */


// Requests that the Bluetooth audio gateway send the delegate a list of calls that are active, on hold, or being set up.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/currentCallList()
func (b_ BluetoothHandsFreeDevice) CurrentCallList() {
	objc.Send[objc.ID](b_.ID, objc.Sel("currentCallList"))
}/* debug [instance_methods/method]: CurrentCallList */


// Calls the phone number on a hands-free phone or headset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/dialNumber(_:)
func (b_ BluetoothHandsFreeDevice) DialNumber(aNumber objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("dialNumber:"), aNumber)
}/* debug [instance_methods/method]: DialNumber */


// Ends the current call or refuses an incoming call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/endCall()
func (b_ BluetoothHandsFreeDevice) EndCall() {
	objc.Send[objc.ID](b_.ID, objc.Sel("endCall"))
}/* debug [instance_methods/method]: EndCall */


// Places all active calls on hold and accepts a held or waiting call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/holdCall()
func (b_ BluetoothHandsFreeDevice) HoldCall() {
	objc.Send[objc.ID](b_.ID, objc.Sel("holdCall"))
}/* debug [instance_methods/method]: HoldCall */


// Calls the phone number stored in a speed dial or memory slot of the hands-free phone or headset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/memoryDial(_:)
func (b_ BluetoothHandsFreeDevice) MemoryDial(memoryLocation int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("memoryDial:"), memoryLocation)
}/* debug [instance_methods/method]: MemoryDial */


// Places all calls except the call with the specified index on hold.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/placeAllOthers(onHold:)
func (b_ BluetoothHandsFreeDevice) PlaceAllOthersOnHold(index int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("placeAllOthersOnHold:"), index)
}/* debug [instance_methods/method]: PlaceAllOthersOnHold */


// Calls the number stored on the hands-free phone or headset again.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/redial()
func (b_ BluetoothHandsFreeDevice) Redial() {
	objc.Send[objc.ID](b_.ID, objc.Sel("redial"))
}/* debug [instance_methods/method]: Redial */


// Ends all active calls and accepts a held or waiting call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/releaseActiveCalls()
func (b_ BluetoothHandsFreeDevice) ReleaseActiveCalls() {
	objc.Send[objc.ID](b_.ID, objc.Sel("releaseActiveCalls"))
}/* debug [instance_methods/method]: ReleaseActiveCalls */


// Ends the call with the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/releaseCall(_:)
func (b_ BluetoothHandsFreeDevice) ReleaseCall(index int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("releaseCall:"), index)
}/* debug [instance_methods/method]: ReleaseCall */


// Ends all calls that are on hold or returns a busy signal for a waiting call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/releaseHeldCalls()
func (b_ BluetoothHandsFreeDevice) ReleaseHeldCalls() {
	objc.Send[objc.ID](b_.ID, objc.Sel("releaseHeldCalls"))
}/* debug [instance_methods/method]: ReleaseHeldCalls */


// Sends an AT command to the Bluetooth audio gateway.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/send(atCommand:)
func (b_ BluetoothHandsFreeDevice) SendATCommand(atCommand objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("sendATCommand:"), atCommand)
}/* debug [instance_methods/method]: SendATCommand */


// Send an AT command to the Bluetooth audio gateway and performs a selector on completion or timeout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/send(atCommand:timeout:selector:target:)
func (b_ BluetoothHandsFreeDevice) SendATCommandTimeoutSelectorTarget(atCommand objc.IObject /* cross-framework: NSString */, timeout float32, selector objc.SEL, target objc.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("sendATCommand:timeout:selector:target:"), atCommand, timeout, selector, target)
}/* debug [instance_methods/method]: SendATCommandTimeoutSelectorTarget */


// Sends the tone associated with a phone key to the hands-free Bluetooth device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/sendDTMF(_:)
func (b_ BluetoothHandsFreeDevice) SendDTMF(character objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("sendDTMF:"), character)
}/* debug [instance_methods/method]: SendDTMF */


// Sends a text message to a phone number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/sendSMS(_:message:)
func (b_ BluetoothHandsFreeDevice) SendSMSMessage(aNumber objc.IObject /* cross-framework: NSString */, aMessage objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("sendSMS:message:"), aNumber, aMessage)
}/* debug [instance_methods/method]: SendSMSMessage */


// Requests that the Bluetooth audio gateway send the subscriber number to the delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/subscriberNumber()
func (b_ BluetoothHandsFreeDevice) SubscriberNumber() {
	objc.Send[objc.ID](b_.ID, objc.Sel("subscriberNumber"))
}/* debug [instance_methods/method]: SubscriberNumber */


// Moves the audio for current and future calls to a Mac.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/transferAudioToComputer()
func (b_ BluetoothHandsFreeDevice) TransferAudioToComputer() {
	objc.Send[objc.ID](b_.ID, objc.Sel("transferAudioToComputer"))
}/* debug [instance_methods/method]: TransferAudioToComputer */


// Moves the audio for current or future calls to a phone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/transferAudioToPhone()
func (b_ BluetoothHandsFreeDevice) TransferAudioToPhone() {
	objc.Send[objc.ID](b_.ID, objc.Sel("transferAudioToPhone"))
}/* debug [instance_methods/method]: TransferAudioToPhone */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BluetoothHandsFreeDevice */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IOBluetoothHandsFreeDevice */


