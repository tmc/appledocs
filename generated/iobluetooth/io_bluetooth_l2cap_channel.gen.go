// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class IOBluetoothL2CAPChannel */


/* debug [class_header]: Header for IOBluetoothL2CAPChannel */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BluetoothL2CAPChannel */
// An interface definition for the [BluetoothL2CAPChannel] class.
type IBluetoothL2CAPChannel interface {
	IBluetoothObject
	
/* debug [class_interface_properties]: Properties for BluetoothL2CAPChannel */
	// properties:
	Device() IOBluetoothDevice
	IncomingMTU() BluetoothL2CAPMTU /* typedef */
	LocalChannelID() BluetoothL2CAPChannelID /* typedef */
	ObjectID() BluetoothObjectID /* typedef */
	OutgoingMTU() BluetoothL2CAPMTU /* typedef */
	PSM() BluetoothL2CAPPSM /* typedef */
	RemoteChannelID() BluetoothL2CAPChannelID /* typedef */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BluetoothL2CAPChannel */
	// methods:
	CloseChannel() int
	Delegate() objc.ID
	IsIncoming() bool
	RegisterForChannelCloseNotificationSelector(observer objc.IObject, inSelector objc.SEL) IBluetoothUserNotification
	RequestRemoteMTU(remoteMTU BluetoothL2CAPMTU /* typedef */) int
	SetDelegate(channelDelegate objc.IObject) int
	SetDelegateWithConfiguration(channelDelegate objc.IObject, channelConfiguration objc.IObject /* cross-framework: NSDictionary */) int
	WriteAsyncLengthRefcon(data unsafe.Pointer, length unsafe.Pointer, refcon unsafe.Pointer) int
	WriteAsyncTrapLengthRefcon(data unsafe.Pointer, length unsafe.Pointer, refcon unsafe.Pointer) int
	WriteSyncLength(data unsafe.Pointer, length unsafe.Pointer) int
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BluetoothL2CAPChannel */
// Alloc allocates a new instance without initialization.
func (bc _BluetoothL2CAPChannelClass) Alloc() BluetoothL2CAPChannel {
	rv := objc.Send[BluetoothL2CAPChannel](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BluetoothL2CAPChannel */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BluetoothL2CAPChannel *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BluetoothL2CAPChannel */

// Allows a client to register for L2CAP channel open notifications for any L2CAP channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/register(forChannelOpenNotifications:selector:)
func (bc _BluetoothL2CAPChannelClass) RegisterForChannelOpenNotificationsSelector(object objc.IObject, selector objc.SEL) IBluetoothUserNotification {
	rv := objc.Send[BluetoothUserNotification](objc.ID(bc.class), objc.Sel("registerForChannelOpenNotifications:selector:"), object, selector)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RegisterForChannelOpenNotificationsSelector) */


// Allows a client to register for L2CAP channel open notifications for certain types of L2CAP channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/register(forChannelOpenNotifications:selector:withPSM:direction:)
func (bc _BluetoothL2CAPChannelClass) RegisterForChannelOpenNotificationsSelectorWithPSMDirection(object objc.IObject, selector objc.SEL, psm BluetoothL2CAPPSM /* typedef */, inDirection BluetoothUserNotificationChannelDirection) IBluetoothUserNotification {
	rv := objc.Send[BluetoothUserNotification](objc.ID(bc.class), objc.Sel("registerForChannelOpenNotifications:selector:withPSM:direction:"), object, selector, psm, inDirection)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RegisterForChannelOpenNotificationsSelectorWithPSMDirection) */


// Returns the IObluetoothL2CAPChannel with the given IOBluetoothObjectID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/withObjectID(_:)
func (bc _BluetoothL2CAPChannelClass) WithObjectID(objectID BluetoothObjectID /* typedef */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("withObjectID:"), objectID)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WithObjectID) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BluetoothL2CAPChannel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BluetoothL2CAPChannel */

// Initiates the close process on an open L2CAP channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/close()
func (b_ BluetoothL2CAPChannel) CloseChannel() int {
	rv := objc.Send[int](b_.ID, objc.Sel("closeChannel"))
	return rv
}/* debug [instance_methods/method]: CloseChannel */


// Returns the currently assigned delegate
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/delegate()
func (b_ BluetoothL2CAPChannel) Delegate() objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_methods/method]: Delegate */


// Returns TRUE if the channel is an incoming channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/isIncoming()
func (b_ BluetoothL2CAPChannel) IsIncoming() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isIncoming"))
	return rv
}/* debug [instance_methods/method]: IsIncoming */


// Allows a client to register for a channel close notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/register(forChannelCloseNotification:selector:)
func (b_ BluetoothL2CAPChannel) RegisterForChannelCloseNotificationSelector(observer objc.IObject, inSelector objc.SEL) IBluetoothUserNotification {
	rv := objc.Send[BluetoothUserNotification](b_.ID, objc.Sel("registerForChannelCloseNotification:selector:"), observer, inSelector)
	return rv
}/* debug [instance_methods/method]: RegisterForChannelCloseNotificationSelector */


// Initiates the process to reconfigure the L2CAP channel with a new outgoing MTU.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/requestRemoteMTU(_:)
func (b_ BluetoothL2CAPChannel) RequestRemoteMTU(remoteMTU BluetoothL2CAPMTU /* typedef */) int {
	rv := objc.Send[int](b_.ID, objc.Sel("requestRemoteMTU:"), remoteMTU)
	return rv
}/* debug [instance_methods/method]: RequestRemoteMTU */


// Allows an object to register itself as client of the L2CAP channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/setDelegate(_:)
func (b_ BluetoothL2CAPChannel) SetDelegate(channelDelegate objc.IObject) int {
	rv := objc.Send[int](b_.ID, objc.Sel("setDelegate:"), channelDelegate)
	return rv
}/* debug [instance_methods/method]: SetDelegate */


// Allows an object to register itself as client of the L2CAP channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/setDelegate(_:withConfiguration:)
func (b_ BluetoothL2CAPChannel) SetDelegateWithConfiguration(channelDelegate objc.IObject, channelConfiguration objc.IObject /* cross-framework: NSDictionary */) int {
	rv := objc.Send[int](b_.ID, objc.Sel("setDelegate:withConfiguration:"), channelDelegate, channelConfiguration)
	return rv
}/* debug [instance_methods/method]: SetDelegateWithConfiguration */


// Writes the given data over the target L2CAP channel asynchronously to the remote device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/writeAsync(_:length:refcon:)
func (b_ BluetoothL2CAPChannel) WriteAsyncLengthRefcon(data unsafe.Pointer, length unsafe.Pointer, refcon unsafe.Pointer) int {
	rv := objc.Send[int](b_.ID, objc.Sel("writeAsync:length:refcon:"), data, length, refcon)
	return rv
}/* debug [instance_methods/method]: WriteAsyncLengthRefcon */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/writeAsyncTrap(_:length:refcon:)
func (b_ BluetoothL2CAPChannel) WriteAsyncTrapLengthRefcon(data unsafe.Pointer, length unsafe.Pointer, refcon unsafe.Pointer) int {
	rv := objc.Send[int](b_.ID, objc.Sel("writeAsyncTrap:length:refcon:"), data, length, refcon)
	return rv
}/* debug [instance_methods/method]: WriteAsyncTrapLengthRefcon */


// Writes the given data synchronously over the target L2CAP channel to the remote device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/writeSync(_:length:)
func (b_ BluetoothL2CAPChannel) WriteSyncLength(data unsafe.Pointer, length unsafe.Pointer) int {
	rv := objc.Send[int](b_.ID, objc.Sel("writeSync:length:"), data, length)
	return rv
}/* debug [instance_methods/method]: WriteSyncLength */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BluetoothL2CAPChannel */

// Returns the IOBluetoothDevice to which the target L2CAP channel is open.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/device
func (b_ BluetoothL2CAPChannel) Device() IOBluetoothDevice {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("device"))
	return rv
}/* debug [instance_properties/getter]: device */


// Returns the current incoming MTU for the L2CAP channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/incomingMTU
func (b_ BluetoothL2CAPChannel) IncomingMTU() BluetoothL2CAPMTU /* typedef */ {
	rv := objc.Send[uint16](b_.ID, objc.Sel("incomingMTU"))
	return rv
}/* debug [instance_properties/getter]: incomingMTU */


// Returns the local L2CAP channel ID for the target L2CAP channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/localChannelID
func (b_ BluetoothL2CAPChannel) LocalChannelID() BluetoothL2CAPChannelID /* typedef */ {
	rv := objc.Send[uint16](b_.ID, objc.Sel("localChannelID"))
	return rv
}/* debug [instance_properties/getter]: localChannelID */


// Returns the IOBluetoothObjectID of the given IOBluetoothL2CAPChannel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/objectID
func (b_ BluetoothL2CAPChannel) ObjectID() BluetoothObjectID /* typedef */ {
	rv := objc.Send[uint64](b_.ID, objc.Sel("objectID"))
	return rv
}/* debug [instance_properties/getter]: objectID */


// Returns the current outgoing MTU for the L2CAP channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/outgoingMTU
func (b_ BluetoothL2CAPChannel) OutgoingMTU() BluetoothL2CAPMTU /* typedef */ {
	rv := objc.Send[uint16](b_.ID, objc.Sel("outgoingMTU"))
	return rv
}/* debug [instance_properties/getter]: outgoingMTU */


// Returns the PSM for the target L2CAP channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/psm
func (b_ BluetoothL2CAPChannel) PSM() BluetoothL2CAPPSM /* typedef */ {
	rv := objc.Send[uint16](b_.ID, objc.Sel("PSM"))
	return rv
}/* debug [instance_properties/getter]: PSM */


// Returns the remote L2CAP channel ID for the target L2CAP channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/remoteChannelID
func (b_ BluetoothL2CAPChannel) RemoteChannelID() BluetoothL2CAPChannelID /* typedef */ {
	rv := objc.Send[uint16](b_.ID, objc.Sel("remoteChannelID"))
	return rv
}/* debug [instance_properties/getter]: remoteChannelID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IOBluetoothL2CAPChannel */



