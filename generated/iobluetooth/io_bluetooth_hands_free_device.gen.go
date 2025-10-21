// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [BluetoothHandsFreeDevice] class.
type IBluetoothHandsFreeDevice interface {
	IBluetoothHandsFree
	AcceptCall()
	AcceptCallOnPhone()
	AddHeldCall()
	CallTransfer()
	CurrentCallList()
	DialNumber(aNumber appkit.string)
	EndCall()
	HoldCall()
	MemoryDial(memoryLocation unsafe.Pointer)
	PlaceAllOthersOnHold(index unsafe.Pointer)
	Redial()
	ReleaseActiveCalls()
	ReleaseCall(index unsafe.Pointer)
	ReleaseHeldCalls()
	SendATCommand(atCommand appkit.string)
	SendATCommandTimeoutSelectorTarget(atCommand appkit.string, timeout unsafe.Pointer, selector objc.SEL, target objectivec.IObject)
	SendDTMF(character appkit.string)
	SendSMSMessage(aNumber appkit.string, aMessage appkit.string)
	SubscriberNumber()
	TransferAudioToComputer()
	TransferAudioToPhone()
}

// An object you use to manage phone calls on a connected Bluetooth hands-free phone or headset.
//
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

// Alloc allocates a new instance without initialization.
func (bc _BluetoothHandsFreeDeviceClass) Alloc() BluetoothHandsFreeDevice {
	rv := objc.Send[BluetoothHandsFreeDevice](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates an object to manage phone calls on a hands-free Bluetooth device.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/init(device:delegate:)
func NewBluetoothHandsFreeDeviceWithDeviceDelegate(device IOBluetoothDevice, delegate objectivec.IObject) BluetoothHandsFreeDevice {
	instance := getBluetoothHandsFreeDeviceClass().Alloc()
	rv := objc.Send[BluetoothHandsFreeDevice](instance.ID, objc.Sel("initWithDevice:delegate:"), device, delegate)
	rv.Autorelease()
	return rv
}


// Accepts an incoming call.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/acceptCall()
func (b_ BluetoothHandsFreeDevice) AcceptCall() {
	objc.Send[objc.ID](b_.ID, objc.Sel("acceptCall"))
}

// Accepts an incoming call and transfers the audio to the managed hands-free phone or headset.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/acceptCallOnPhone()
func (b_ BluetoothHandsFreeDevice) AcceptCallOnPhone() {
	objc.Send[objc.ID](b_.ID, objc.Sel("acceptCallOnPhone"))
}

// Adds held calls to the current conversation.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/addHeldCall()
func (b_ BluetoothHandsFreeDevice) AddHeldCall() {
	objc.Send[objc.ID](b_.ID, objc.Sel("addHeldCall"))
}

// Ends all calls that are active or on hold, and accepts any waiting calls.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/callTransfer()
func (b_ BluetoothHandsFreeDevice) CallTransfer() {
	objc.Send[objc.ID](b_.ID, objc.Sel("callTransfer"))
}

// Requests that the Bluetooth audio gateway send the delegate a list of calls that are active, on hold, or being set up.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/currentCallList()
func (b_ BluetoothHandsFreeDevice) CurrentCallList() {
	objc.Send[objc.ID](b_.ID, objc.Sel("currentCallList"))
}

// Calls the phone number on a hands-free phone or headset.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/dialNumber(_:)
func (b_ BluetoothHandsFreeDevice) DialNumber(aNumber appkit.string) {
	objc.Send[objc.ID](b_.ID, objc.Sel("dialNumber:"), aNumber)
}

// Ends the current call or refuses an incoming call.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/endCall()
func (b_ BluetoothHandsFreeDevice) EndCall() {
	objc.Send[objc.ID](b_.ID, objc.Sel("endCall"))
}

// Places all active calls on hold and accepts a held or waiting call.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/holdCall()
func (b_ BluetoothHandsFreeDevice) HoldCall() {
	objc.Send[objc.ID](b_.ID, objc.Sel("holdCall"))
}

// Calls the phone number stored in a speed dial or memory slot of the hands-free phone or headset.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/memoryDial(_:)
func (b_ BluetoothHandsFreeDevice) MemoryDial(memoryLocation unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("memoryDial:"), memoryLocation)
}

// Places all calls except the call with the specified index on hold.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/placeAllOthers(onHold:)
func (b_ BluetoothHandsFreeDevice) PlaceAllOthersOnHold(index unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("placeAllOthersOnHold:"), index)
}

// Calls the number stored on the hands-free phone or headset again.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/redial()
func (b_ BluetoothHandsFreeDevice) Redial() {
	objc.Send[objc.ID](b_.ID, objc.Sel("redial"))
}

// Ends all active calls and accepts a held or waiting call.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/releaseActiveCalls()
func (b_ BluetoothHandsFreeDevice) ReleaseActiveCalls() {
	objc.Send[objc.ID](b_.ID, objc.Sel("releaseActiveCalls"))
}

// Ends the call with the specified index.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/releaseCall(_:)
func (b_ BluetoothHandsFreeDevice) ReleaseCall(index unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("releaseCall:"), index)
}

// Ends all calls that are on hold or returns a busy signal for a waiting call.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/releaseHeldCalls()
func (b_ BluetoothHandsFreeDevice) ReleaseHeldCalls() {
	objc.Send[objc.ID](b_.ID, objc.Sel("releaseHeldCalls"))
}

// Sends an AT command to the Bluetooth audio gateway.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/send(atCommand:)
func (b_ BluetoothHandsFreeDevice) SendATCommand(atCommand appkit.string) {
	objc.Send[objc.ID](b_.ID, objc.Sel("sendATCommand:"), atCommand)
}

// Send an AT command to the Bluetooth audio gateway and performs a selector on completion or timeout.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/send(atCommand:timeout:selector:target:)
func (b_ BluetoothHandsFreeDevice) SendATCommandTimeoutSelectorTarget(atCommand appkit.string, timeout unsafe.Pointer, selector objc.SEL, target objectivec.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("sendATCommand:timeout:selector:target:"), atCommand, timeout, selector, target)
}

// Sends the tone associated with a phone key to the hands-free Bluetooth device.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/sendDTMF(_:)
func (b_ BluetoothHandsFreeDevice) SendDTMF(character appkit.string) {
	objc.Send[objc.ID](b_.ID, objc.Sel("sendDTMF:"), character)
}

// Sends a text message to a phone number.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/sendSMS(_:message:)
func (b_ BluetoothHandsFreeDevice) SendSMSMessage(aNumber appkit.string, aMessage appkit.string) {
	objc.Send[objc.ID](b_.ID, objc.Sel("sendSMS:message:"), aNumber, aMessage)
}

// Requests that the Bluetooth audio gateway send the subscriber number to the delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/subscriberNumber()
func (b_ BluetoothHandsFreeDevice) SubscriberNumber() {
	objc.Send[objc.ID](b_.ID, objc.Sel("subscriberNumber"))
}

// Moves the audio for current and future calls to a Mac.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/transferAudioToComputer()
func (b_ BluetoothHandsFreeDevice) TransferAudioToComputer() {
	objc.Send[objc.ID](b_.ID, objc.Sel("transferAudioToComputer"))
}

// Moves the audio for current or future calls to a phone.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/transferAudioToPhone()
func (b_ BluetoothHandsFreeDevice) TransferAudioToPhone() {
	objc.Send[objc.ID](b_.ID, objc.Sel("transferAudioToPhone"))
}


