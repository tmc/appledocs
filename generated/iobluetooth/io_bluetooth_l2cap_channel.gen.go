// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BluetoothL2CAPChannel] class.
var (
	BluetoothL2CAPChannelClass     _BluetoothL2CAPChannelClass
	BluetoothL2CAPChannelClassOnce sync.Once
)

func getBluetoothL2CAPChannelClass() _BluetoothL2CAPChannelClass {
	BluetoothL2CAPChannelClassOnce.Do(func() {
		BluetoothL2CAPChannelClass = _BluetoothL2CAPChannelClass{objc.GetClass("IOBluetoothL2CAPChannel")}
	})
	return BluetoothL2CAPChannelClass
}

type _BluetoothL2CAPChannelClass struct {
	class objc.Class
}

// An interface definition for the [BluetoothL2CAPChannel] class.
type IBluetoothL2CAPChannel interface {
	IBluetoothObject
	// properties:
	Device() IOBluetoothDevice /* already interface */
	IncomingMTU() BluetoothL2CAPMTU /* typedef */
	LocalChannelID() BluetoothL2CAPChannelID /* typedef */
	ObjectID() objc.IObject /* cross-framework: BluetoothObjectID */
	OutgoingMTU() BluetoothL2CAPMTU /* typedef */
	PSM() BluetoothL2CAPPSM /* typedef */
	RemoteChannelID() BluetoothL2CAPChannelID /* typedef */
	// methods:
	CloseChannel() Return /* not a class type */
	Delegate() objc.ID
	IsIncoming() bool /* primitive/slice/pointer. */
	RegisterForChannelCloseNotificationSelector(observer objectivec.IObject, inSelector objc.SEL) IBluetoothUserNotification
	RequestRemoteMTU(remoteMTU BluetoothL2CAPMTU /* typedef */) Return /* not a class type */
	SetDelegate(channelDelegate objectivec.IObject) Return /* not a class type */
	SetDelegateWithConfiguration(channelDelegate objectivec.IObject, channelConfiguration objectivec.IObject) Return /* not a class type */
	WriteAsyncLengthRefcon(data unsafe.Pointer, length unsafe.Pointer, refcon unsafe.Pointer) Return /* not a class type */
	WriteAsyncTrapLengthRefcon(data unsafe.Pointer, length unsafe.Pointer, refcon unsafe.Pointer) Return /* not a class type */
	WriteSyncLength(data unsafe.Pointer, length unsafe.Pointer) Return /* not a class type */
}

// An instance of IOBluetoothL2CAPChannel represents a single open L2CAP channel.
//
// A client won’t create IOBluetoothL2CAPChannel objects directly. Instead, the IOBluetoothDevice’s L2CAP channel open API is responsible for opening a new L2CAP channel and returning an IOBluetoothL2CAPChannel instance representing that newly opened channel. Additionally, the IOBluetooth notification system will send notifications when new L2CAP channels are open (if requested). After a new L2CAP channel is opened, the L2CAP configuration process will not be completed until an incoming data listener is registered with the IOBluetoothL2CAPChannel object. The reason for this is to due to the limited buffering done of incoming L2CAP data. This way, we avoid the situation where incoming data is received before the client is ready for it. Once a client is done with an IOBluetoothL2CAPChannel that it opened, it should call -closeChannel. Additionally, if the client does not intend to use the connection to the remote device any further, it should call -closeConnection on the IOBluetoothDevice object.


// An instance of IOBluetoothL2CAPChannel represents a single open L2CAP channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel
type BluetoothL2CAPChannel struct {
	BluetoothObject
}

// BluetoothL2CAPChannelFrom constructs a [BluetoothL2CAPChannel] from an unsafe.Pointer.
//
// An instance of IOBluetoothL2CAPChannel represents a single open L2CAP channel.
func BluetoothL2CAPChannelFrom(ptr unsafe.Pointer) BluetoothL2CAPChannel {
	return BluetoothL2CAPChannel{
		BluetoothObject: BluetoothObjectFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _BluetoothL2CAPChannelClass) Alloc() BluetoothL2CAPChannel {
	rv := objc.Send[BluetoothL2CAPChannel](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BluetoothL2CAPChannelClass) New() BluetoothL2CAPChannel {
	rv := objc.Send[BluetoothL2CAPChannel](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BluetoothL2CAPChannel) Init() BluetoothL2CAPChannel {
	rv := objc.Send[BluetoothL2CAPChannel](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BluetoothL2CAPChannel) Autorelease() BluetoothL2CAPChannel {
	rv := objc.Send[BluetoothL2CAPChannel](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBluetoothL2CAPChannel creates a new BluetoothL2CAPChannel instance.
func NewBluetoothL2CAPChannel() BluetoothL2CAPChannel {
	return getBluetoothL2CAPChannelClass().New()
}



// Allows a client to register for L2CAP channel open notifications for any L2CAP channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/register(forChannelOpenNotifications:selector:)
func (bc _BluetoothL2CAPChannelClass) RegisterForChannelOpenNotificationsSelector(object objectivec.IObject, selector objc.SEL) IBluetoothUserNotification {
	rv := objc.Send[BluetoothUserNotification](objc.ID(bc.class), objc.Sel("registerForChannelOpenNotifications:selector:"), object, selector)
	return rv
}


// Allows a client to register for L2CAP channel open notifications for certain types of L2CAP channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/register(forChannelOpenNotifications:selector:withPSM:direction:)
func (bc _BluetoothL2CAPChannelClass) RegisterForChannelOpenNotificationsSelectorWithPSMDirection(object objectivec.IObject, selector objc.SEL, psm BluetoothL2CAPPSM /* typedef */, inDirection BluetoothUserNotificationChannelDirection) IBluetoothUserNotification {
	rv := objc.Send[BluetoothUserNotification](objc.ID(bc.class), objc.Sel("registerForChannelOpenNotifications:selector:withPSM:direction:"), object, selector, psm, inDirection)
	return rv
}


// Returns the IObluetoothL2CAPChannel with the given IOBluetoothObjectID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/withObjectID(_:)
func (bc _BluetoothL2CAPChannelClass) WithObjectID(objectID objc.IObject /* cross-framework BluetoothObjectID */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("withObjectID:"), objectID)
	return rv
}


// Initiates the close process on an open L2CAP channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/close()
func (b_ BluetoothL2CAPChannel) CloseChannel() Return /* not a class type */ {
	rv := objc.Send[Return](b_.ID, objc.Sel("closeChannel"))
	return rv
}


// Returns the currently assigned delegate
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/delegate()
func (b_ BluetoothL2CAPChannel) Delegate() objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("delegate"))
	return rv
}


// Returns TRUE if the channel is an incoming channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/isIncoming()
func (b_ BluetoothL2CAPChannel) IsIncoming() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("isIncoming"))
	return rv
}


// Allows a client to register for a channel close notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/register(forChannelCloseNotification:selector:)
func (b_ BluetoothL2CAPChannel) RegisterForChannelCloseNotificationSelector(observer objectivec.IObject, inSelector objc.SEL) IBluetoothUserNotification {
	rv := objc.Send[BluetoothUserNotification](b_.ID, objc.Sel("registerForChannelCloseNotification:selector:"), observer, inSelector)
	return rv
}


// Initiates the process to reconfigure the L2CAP channel with a new outgoing MTU.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/requestRemoteMTU(_:)
func (b_ BluetoothL2CAPChannel) RequestRemoteMTU(remoteMTU BluetoothL2CAPMTU /* typedef */) Return /* not a class type */ {
	rv := objc.Send[Return](b_.ID, objc.Sel("requestRemoteMTU:"), remoteMTU)
	return rv
}


// Allows an object to register itself as client of the L2CAP channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/setDelegate(_:)
func (b_ BluetoothL2CAPChannel) SetDelegate(channelDelegate objectivec.IObject) Return /* not a class type */ {
	rv := objc.Send[Return](b_.ID, objc.Sel("setDelegate:"), channelDelegate)
	return rv
}


// Allows an object to register itself as client of the L2CAP channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/setDelegate(_:withConfiguration:)
func (b_ BluetoothL2CAPChannel) SetDelegateWithConfiguration(channelDelegate objectivec.IObject, channelConfiguration objectivec.IObject) Return /* not a class type */ {
	rv := objc.Send[Return](b_.ID, objc.Sel("setDelegate:withConfiguration:"), channelDelegate, channelConfiguration)
	return rv
}


// Writes the given data over the target L2CAP channel asynchronously to the remote device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/writeAsync(_:length:refcon:)
func (b_ BluetoothL2CAPChannel) WriteAsyncLengthRefcon(data unsafe.Pointer, length unsafe.Pointer, refcon unsafe.Pointer) Return /* not a class type */ {
	rv := objc.Send[Return](b_.ID, objc.Sel("writeAsync:length:refcon:"), data, length, refcon)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/writeAsyncTrap(_:length:refcon:)
func (b_ BluetoothL2CAPChannel) WriteAsyncTrapLengthRefcon(data unsafe.Pointer, length unsafe.Pointer, refcon unsafe.Pointer) Return /* not a class type */ {
	rv := objc.Send[Return](b_.ID, objc.Sel("writeAsyncTrap:length:refcon:"), data, length, refcon)
	return rv
}


// Writes the given data synchronously over the target L2CAP channel to the remote device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/writeSync(_:length:)
func (b_ BluetoothL2CAPChannel) WriteSyncLength(data unsafe.Pointer, length unsafe.Pointer) Return /* not a class type */ {
	rv := objc.Send[Return](b_.ID, objc.Sel("writeSync:length:"), data, length)
	return rv
}


// Returns the IOBluetoothDevice to which the target L2CAP channel is open.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/device
func (b_ BluetoothL2CAPChannel) Device() IOBluetoothDevice /* already interface */ {
	rv := objc.Send[BluetoothDevice](b_.ID, objc.Sel("device"))
	return rv
}


// Returns the current incoming MTU for the L2CAP channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/incomingMTU
func (b_ BluetoothL2CAPChannel) IncomingMTU() BluetoothL2CAPMTU /* typedef */ {
	rv := objc.Send[BluetoothL2CAPMTU](b_.ID, objc.Sel("incomingMTU"))
	return rv
}


// Returns the local L2CAP channel ID for the target L2CAP channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/localChannelID
func (b_ BluetoothL2CAPChannel) LocalChannelID() BluetoothL2CAPChannelID /* typedef */ {
	rv := objc.Send[BluetoothL2CAPChannelID](b_.ID, objc.Sel("localChannelID"))
	return rv
}


// Returns the IOBluetoothObjectID of the given IOBluetoothL2CAPChannel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/objectID
func (b_ BluetoothL2CAPChannel) ObjectID() objc.IObject /* cross-framework: BluetoothObjectID */ {
	rv := objc.Send[BluetoothObjectID](b_.ID, objc.Sel("objectID"))
	return rv
}


// Returns the current outgoing MTU for the L2CAP channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/outgoingMTU
func (b_ BluetoothL2CAPChannel) OutgoingMTU() BluetoothL2CAPMTU /* typedef */ {
	rv := objc.Send[BluetoothL2CAPMTU](b_.ID, objc.Sel("outgoingMTU"))
	return rv
}


// Returns the PSM for the target L2CAP channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/psm
func (b_ BluetoothL2CAPChannel) PSM() BluetoothL2CAPPSM /* typedef */ {
	rv := objc.Send[BluetoothL2CAPPSM](b_.ID, objc.Sel("PSM"))
	return rv
}


// Returns the remote L2CAP channel ID for the target L2CAP channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/remoteChannelID
func (b_ BluetoothL2CAPChannel) RemoteChannelID() BluetoothL2CAPChannelID /* typedef */ {
	rv := objc.Send[BluetoothL2CAPChannelID](b_.ID, objc.Sel("remoteChannelID"))
	return rv
}



