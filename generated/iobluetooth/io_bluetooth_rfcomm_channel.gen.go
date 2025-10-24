// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class IOBluetoothRFCOMMChannel */


/* debug [class_header]: Header for IOBluetoothRFCOMMChannel */
// The class instance for the [BluetoothRFCOMMChannel] class.
var (
	BluetoothRFCOMMChannelClass     _BluetoothRFCOMMChannelClass
	BluetoothRFCOMMChannelClassOnce sync.Once
)

func getBluetoothRFCOMMChannelClass() _BluetoothRFCOMMChannelClass {
	BluetoothRFCOMMChannelClassOnce.Do(func() {
		BluetoothRFCOMMChannelClass = _BluetoothRFCOMMChannelClass{objc.GetClass("IOBluetoothRFCOMMChannel")}
	})
	return BluetoothRFCOMMChannelClass
}

type _BluetoothRFCOMMChannelClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BluetoothRFCOMMChannel */
// An interface definition for the [BluetoothRFCOMMChannel] class.
type IBluetoothRFCOMMChannel interface {
	IBluetoothObject
	
/* debug [class_interface_properties]: Properties for BluetoothRFCOMMChannel */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BluetoothRFCOMMChannel */
	// methods:
	CloseChannel() int
	Delegate() objc.ID
	GetDevice() IBluetoothDevice
	GetChannelID() BluetoothRFCOMMChannelID /* typedef */
	GetMTU() BluetoothRFCOMMMTU /* typedef */
	GetObjectID() BluetoothObjectID /* typedef */
	GetRFCOMMChannelRef() BluetoothRFCOMMChannelRef /* typedef */
	IsIncoming() bool
	IsOpen() bool
	IsTransmissionPaused() bool
	RegisterForChannelCloseNotificationSelector(observer objc.IObject, inSelector objc.SEL) IBluetoothUserNotification
	SendRemoteLineStatus(lineStatus BluetoothRFCOMMLineStatus) int
	SetDelegate(delegate objc.IObject) int
	SetSerialParametersDataBitsParityStopBits(speed unsafe.Pointer, nBits unsafe.Pointer, parity BluetoothRFCOMMParityType, bitStop unsafe.Pointer) int
	WriteAsyncLengthRefcon(data unsafe.Pointer, length unsafe.Pointer, refcon unsafe.Pointer) int
	WriteSyncLength(data unsafe.Pointer, length unsafe.Pointer) int
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BluetoothRFCOMMChannel */
// Alloc allocates a new instance without initialization.
func (bc _BluetoothRFCOMMChannelClass) Alloc() BluetoothRFCOMMChannel {
	rv := objc.Send[BluetoothRFCOMMChannel](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BluetoothRFCOMMChannelClass) New() BluetoothRFCOMMChannel {
	rv := objc.Send[BluetoothRFCOMMChannel](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BluetoothRFCOMMChannel) Init() BluetoothRFCOMMChannel {
	rv := objc.Send[BluetoothRFCOMMChannel](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BluetoothRFCOMMChannel) Autorelease() BluetoothRFCOMMChannel {
	rv := objc.Send[BluetoothRFCOMMChannel](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBluetoothRFCOMMChannel creates a new BluetoothRFCOMMChannel instance.
func NewBluetoothRFCOMMChannel() BluetoothRFCOMMChannel {
	return getBluetoothRFCOMMChannelClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BluetoothRFCOMMChannel */
// An instance of this class represents an RFCOMM channel as defined by the Bluetooth SDP spec..
//
// An RFCOMM channel object can be obtained by opening an RFCOMM channel in a device, or by requesting a notification when a channel is created (this is commonly used to provide services).


// An instance of this class represents an RFCOMM channel as defined by the Bluetooth SDP spec..
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel
type BluetoothRFCOMMChannel struct {
	BluetoothObject
}

// BluetoothRFCOMMChannelFrom constructs a [BluetoothRFCOMMChannel] from an unsafe.Pointer.
//
// An instance of this class represents an RFCOMM channel as defined by the Bluetooth SDP spec..
func BluetoothRFCOMMChannelFrom(ptr unsafe.Pointer) BluetoothRFCOMMChannel {
	return BluetoothRFCOMMChannel{
		BluetoothObject: BluetoothObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BluetoothRFCOMMChannel *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BluetoothRFCOMMChannel */

// Allows a client to register for RFCOMM channel open notifications for any RFCOMM channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/register(forChannelOpenNotifications:selector:)
func (bc _BluetoothRFCOMMChannelClass) RegisterForChannelOpenNotificationsSelector(object objc.IObject, selector objc.SEL) IBluetoothUserNotification {
	rv := objc.Send[BluetoothUserNotification](objc.ID(bc.class), objc.Sel("registerForChannelOpenNotifications:selector:"), object, selector)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RegisterForChannelOpenNotificationsSelector) */


// Allows a client to register for RFCOMM channel open notifications for certain types of RFCOMM channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/register(forChannelOpenNotifications:selector:withChannelID:direction:)
func (bc _BluetoothRFCOMMChannelClass) RegisterForChannelOpenNotificationsSelectorWithChannelIDDirection(object objc.IObject, selector objc.SEL, channelID BluetoothRFCOMMChannelID /* typedef */, inDirection BluetoothUserNotificationChannelDirection) IBluetoothUserNotification {
	rv := objc.Send[BluetoothUserNotification](objc.ID(bc.class), objc.Sel("registerForChannelOpenNotifications:selector:withChannelID:direction:"), object, selector, channelID, inDirection)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RegisterForChannelOpenNotificationsSelectorWithChannelIDDirection) */


// Returns the IObluetoothRFCOMMChannel with the given IOBluetoothObjectID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/withObjectID(_:)
func (bc _BluetoothRFCOMMChannelClass) WithObjectID(objectID BluetoothObjectID /* typedef */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("withObjectID:"), objectID)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WithObjectID) */


// Method call to convert an IOBluetoothRFCOMMChannelRef into an IOBluetoothRFCOMMChannel *.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/withRFCOMMChannelRef(_:)
func (bc _BluetoothRFCOMMChannelClass) WithRFCOMMChannelRef(rfcommChannelRef BluetoothRFCOMMChannelRef /* typedef */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("withRFCOMMChannelRef:"), rfcommChannelRef)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WithRFCOMMChannelRef) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BluetoothRFCOMMChannel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BluetoothRFCOMMChannel */

// Close the channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/close()
func (b_ BluetoothRFCOMMChannel) CloseChannel() int {
	rv := objc.Send[int](b_.ID, objc.Sel("closeChannel"))
	return rv
}/* debug [instance_methods/method]: CloseChannel */


// Returns the object delegate
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/delegate()
func (b_ BluetoothRFCOMMChannel) Delegate() objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_methods/method]: Delegate */


// Returns the Bluetooth Device that carries the rfcomm data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/getDevice()
func (b_ BluetoothRFCOMMChannel) GetDevice() IBluetoothDevice {
	rv := objc.Send[BluetoothDevice](b_.ID, objc.Sel("getDevice"))
	return rv
}/* debug [instance_methods/method]: GetDevice */


// Returns the object rfcomm channel ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/getID()
func (b_ BluetoothRFCOMMChannel) GetChannelID() BluetoothRFCOMMChannelID /* typedef */ {
	rv := objc.Send[uint8](b_.ID, objc.Sel("getChannelID"))
	return rv
}/* debug [instance_methods/method]: GetChannelID */


// Returns the channel maximum transfer unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/getMTU()
func (b_ BluetoothRFCOMMChannel) GetMTU() BluetoothRFCOMMMTU /* typedef */ {
	rv := objc.Send[uint16](b_.ID, objc.Sel("getMTU"))
	return rv
}/* debug [instance_methods/method]: GetMTU */


// Returns the IOBluetoothObjectID of the given IOBluetoothRFCOMMChannel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/getObjectID()
func (b_ BluetoothRFCOMMChannel) GetObjectID() BluetoothObjectID /* typedef */ {
	rv := objc.Send[uint64](b_.ID, objc.Sel("getObjectID"))
	return rv
}/* debug [instance_methods/method]: GetObjectID */


// Returns an IOBluetoothRFCOMMChannelRef representation of the target IOBluetoothRFCOMMChannel object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/getRef()
func (b_ BluetoothRFCOMMChannel) GetRFCOMMChannelRef() BluetoothRFCOMMChannelRef /* typedef */ {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getRFCOMMChannelRef"))
	return rv
}/* debug [instance_methods/method]: GetRFCOMMChannelRef */


// Returns the direction of the channel. An incoming channel is one that was opened by the remote device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/isIncoming()
func (b_ BluetoothRFCOMMChannel) IsIncoming() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isIncoming"))
	return rv
}/* debug [instance_methods/method]: IsIncoming */


// Returns the state of the channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/isOpen()
func (b_ BluetoothRFCOMMChannel) IsOpen() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isOpen"))
	return rv
}/* debug [instance_methods/method]: IsOpen */


// Returns TRUE if flow control is off.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/isTransmissionPaused()
func (b_ BluetoothRFCOMMChannel) IsTransmissionPaused() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isTransmissionPaused"))
	return rv
}/* debug [instance_methods/method]: IsTransmissionPaused */


// Allows a client to register for a channel close notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/register(forChannelCloseNotification:selector:)
func (b_ BluetoothRFCOMMChannel) RegisterForChannelCloseNotificationSelector(observer objc.IObject, inSelector objc.SEL) IBluetoothUserNotification {
	rv := objc.Send[BluetoothUserNotification](b_.ID, objc.Sel("registerForChannelCloseNotification:selector:"), observer, inSelector)
	return rv
}/* debug [instance_methods/method]: RegisterForChannelCloseNotificationSelector */


// Sends an error to the remote side.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/sendRemoteLineStatus(_:)
func (b_ BluetoothRFCOMMChannel) SendRemoteLineStatus(lineStatus BluetoothRFCOMMLineStatus) int {
	rv := objc.Send[int](b_.ID, objc.Sel("sendRemoteLineStatus:"), lineStatus)
	return rv
}/* debug [instance_methods/method]: SendRemoteLineStatus */


// Allows an object to register itself as a client of the RFCOMM channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/setDelegate(_:)
func (b_ BluetoothRFCOMMChannel) SetDelegate(delegate objc.IObject) int {
	rv := objc.Send[int](b_.ID, objc.Sel("setDelegate:"), delegate)
	return rv
}/* debug [instance_methods/method]: SetDelegate */


// Changes the parameters of the serial connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/setSerialParameters(_:dataBits:parity:stopBits:)
func (b_ BluetoothRFCOMMChannel) SetSerialParametersDataBitsParityStopBits(speed unsafe.Pointer, nBits unsafe.Pointer, parity BluetoothRFCOMMParityType, bitStop unsafe.Pointer) int {
	rv := objc.Send[int](b_.ID, objc.Sel("setSerialParameters:dataBits:parity:stopBits:"), speed, nBits, parity, bitStop)
	return rv
}/* debug [instance_methods/method]: SetSerialParametersDataBitsParityStopBits */


// Sends a block of data in the channel asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/writeAsync(_:length:refcon:)
func (b_ BluetoothRFCOMMChannel) WriteAsyncLengthRefcon(data unsafe.Pointer, length unsafe.Pointer, refcon unsafe.Pointer) int {
	rv := objc.Send[int](b_.ID, objc.Sel("writeAsync:length:refcon:"), data, length, refcon)
	return rv
}/* debug [instance_methods/method]: WriteAsyncLengthRefcon */


// Sends a block of data in the channel synchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/writeSync(_:length:)
func (b_ BluetoothRFCOMMChannel) WriteSyncLength(data unsafe.Pointer, length unsafe.Pointer) int {
	rv := objc.Send[int](b_.ID, objc.Sel("writeSync:length:"), data, length)
	return rv
}/* debug [instance_methods/method]: WriteSyncLength */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BluetoothRFCOMMChannel */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IOBluetoothRFCOMMChannel */



