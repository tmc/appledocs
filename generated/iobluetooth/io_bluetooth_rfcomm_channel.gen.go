// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [BluetoothRFCOMMChannel] class.
type IBluetoothRFCOMMChannel interface {
	IBluetoothObject
	CloseChannel() unsafe.Pointer
	Delegate() objc.ID
	GetDevice() unsafe.Pointer
	GetChannelID() unsafe.Pointer
	GetMTU() unsafe.Pointer
	GetObjectID() unsafe.Pointer
	GetRFCOMMChannelRef() unsafe.Pointer
	IsIncoming() bool
	IsOpen() bool
	IsTransmissionPaused() bool
	RegisterForChannelCloseNotificationSelector(observer objc.ID, inSelector objc.SEL) unsafe.Pointer
	SendRemoteLineStatus(lineStatus unsafe.Pointer) unsafe.Pointer
	SetDelegate(delegate objc.ID) unsafe.Pointer
	SetSerialParametersDataBitsParityStopBits(speed unsafe.Pointer, nBits unsafe.Pointer, parity unsafe.Pointer, bitStop unsafe.Pointer) unsafe.Pointer
	WriteLengthSleep(data unsafe.Pointer, length unsafe.Pointer, sleep bool) unsafe.Pointer
	WriteAsyncLengthRefcon(data unsafe.Pointer, length unsafe.Pointer, refcon unsafe.Pointer) unsafe.Pointer
	WriteSimpleLengthSleepBytesSent(data unsafe.Pointer, length unsafe.Pointer, sleep bool, numBytesSent unsafe.Pointer) unsafe.Pointer
	WriteSyncLength(data unsafe.Pointer, length unsafe.Pointer) unsafe.Pointer
}

// An instance of this class represents an RFCOMM channel as defined by the Bluetooth SDP spec..
//
// An RFCOMM channel object can be obtained by opening an RFCOMM channel in a device, or by requesting a notification when a channel is created (this is commonly used to provide services).
//
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

// Alloc allocates a new instance without initialization.
func (bc _BluetoothRFCOMMChannelClass) Alloc() BluetoothRFCOMMChannel {
	rv := objc.Send[BluetoothRFCOMMChannel](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Allows a client to register for RFCOMM channel open notifications for any RFCOMM channel.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/register(forChannelOpenNotifications:selector:)
func (bc _BluetoothRFCOMMChannelClass) RegisterForChannelOpenNotificationsSelector(object objc.ID, selector objc.SEL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("registerForChannelOpenNotifications:selector:"), object, selector)
	return rv
}

// Allows a client to register for RFCOMM channel open notifications for certain types of RFCOMM channels.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/register(forChannelOpenNotifications:selector:withChannelID:direction:)
func (bc _BluetoothRFCOMMChannelClass) RegisterForChannelOpenNotificationsSelectorWithChannelIDDirection(object objc.ID, selector objc.SEL, channelID unsafe.Pointer, inDirection unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("registerForChannelOpenNotifications:selector:withChannelID:direction:"), object, selector, channelID, inDirection)
	return rv
}

// Returns the IObluetoothRFCOMMChannel with the given IOBluetoothObjectID.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/withObjectID(_:)
func (bc _BluetoothRFCOMMChannelClass) WithObjectID(objectID unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("withObjectID:"), objectID)
	return rv
}

// Method call to convert an IOBluetoothRFCOMMChannelRef into an IOBluetoothRFCOMMChannel *.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/withRFCOMMChannelRef(_:)
func (bc _BluetoothRFCOMMChannelClass) WithRFCOMMChannelRef(rfcommChannelRef unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("withRFCOMMChannelRef:"), rfcommChannelRef)
	return rv
}

// Close the channel.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/close()
func (b_ BluetoothRFCOMMChannel) CloseChannel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("closeChannel"))
	return rv
}

// Returns the object delegate
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/delegate()
func (b_ BluetoothRFCOMMChannel) Delegate() objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("delegate"))
	return rv
}

// Returns the Bluetooth Device that carries the rfcomm data.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/getDevice()
func (b_ BluetoothRFCOMMChannel) GetDevice() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getDevice"))
	return rv
}

// Returns the object rfcomm channel ID.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/getID()
func (b_ BluetoothRFCOMMChannel) GetChannelID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getChannelID"))
	return rv
}

// Returns the channel maximum transfer unit.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/getMTU()
func (b_ BluetoothRFCOMMChannel) GetMTU() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getMTU"))
	return rv
}

// Returns the IOBluetoothObjectID of the given IOBluetoothRFCOMMChannel.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/getObjectID()
func (b_ BluetoothRFCOMMChannel) GetObjectID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getObjectID"))
	return rv
}

// Returns an IOBluetoothRFCOMMChannelRef representation of the target IOBluetoothRFCOMMChannel object.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/getRef()
func (b_ BluetoothRFCOMMChannel) GetRFCOMMChannelRef() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getRFCOMMChannelRef"))
	return rv
}

// Returns the direction of the channel. An incoming channel is one that was opened by the remote device.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/isIncoming()
func (b_ BluetoothRFCOMMChannel) IsIncoming() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isIncoming"))
	return rv
}

// Returns the state of the channel.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/isOpen()
func (b_ BluetoothRFCOMMChannel) IsOpen() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isOpen"))
	return rv
}

// Returns TRUE if flow control is off.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/isTransmissionPaused()
func (b_ BluetoothRFCOMMChannel) IsTransmissionPaused() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isTransmissionPaused"))
	return rv
}

// Allows a client to register for a channel close notification.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/register(forChannelCloseNotification:selector:)
func (b_ BluetoothRFCOMMChannel) RegisterForChannelCloseNotificationSelector(observer objc.ID, inSelector objc.SEL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("registerForChannelCloseNotification:selector:"), observer, inSelector)
	return rv
}

// Sends an error to the remote side.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/sendRemoteLineStatus(_:)
func (b_ BluetoothRFCOMMChannel) SendRemoteLineStatus(lineStatus unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("sendRemoteLineStatus:"), lineStatus)
	return rv
}

// Allows an object to register itself as a client of the RFCOMM channel.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/setDelegate(_:)
func (b_ BluetoothRFCOMMChannel) SetDelegate(delegate objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("setDelegate:"), delegate)
	return rv
}

// Changes the parameters of the serial connection.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/setSerialParameters(_:dataBits:parity:stopBits:)
func (b_ BluetoothRFCOMMChannel) SetSerialParametersDataBitsParityStopBits(speed unsafe.Pointer, nBits unsafe.Pointer, parity unsafe.Pointer, bitStop unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("setSerialParameters:dataBits:parity:stopBits:"), speed, nBits, parity, bitStop)
	return rv
}

// Sends a block of data in the channel syncronously.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/write:length:sleep:
func (b_ BluetoothRFCOMMChannel) WriteLengthSleep(data unsafe.Pointer, length unsafe.Pointer, sleep bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("write:length:sleep:"), data, length, sleep)
	return rv
}

// Sends a block of data in the channel asynchronously.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/writeAsync(_:length:refcon:)
func (b_ BluetoothRFCOMMChannel) WriteAsyncLengthRefcon(data unsafe.Pointer, length unsafe.Pointer, refcon unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("writeAsync:length:refcon:"), data, length, refcon)
	return rv
}

// Sends a block of data in the channel.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/writeSimple:length:sleep:bytesSent:
func (b_ BluetoothRFCOMMChannel) WriteSimpleLengthSleepBytesSent(data unsafe.Pointer, length unsafe.Pointer, sleep bool, numBytesSent unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("writeSimple:length:sleep:bytesSent:"), data, length, sleep, numBytesSent)
	return rv
}

// Sends a block of data in the channel synchronously.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/writeSync(_:length:)
func (b_ BluetoothRFCOMMChannel) WriteSyncLength(data unsafe.Pointer, length unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("writeSync:length:"), data, length)
	return rv
}



