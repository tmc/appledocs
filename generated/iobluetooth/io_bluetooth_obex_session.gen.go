// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class IOBluetoothOBEXSession */


/* debug [class_header]: Header for IOBluetoothOBEXSession */
// The class instance for the [BluetoothOBEXSession] class.
var (
	BluetoothOBEXSessionClass     _BluetoothOBEXSessionClass
	BluetoothOBEXSessionClassOnce sync.Once
)

func getBluetoothOBEXSessionClass() _BluetoothOBEXSessionClass {
	BluetoothOBEXSessionClassOnce.Do(func() {
		BluetoothOBEXSessionClass = _BluetoothOBEXSessionClass{objc.GetClass("IOBluetoothOBEXSession")}
	})
	return BluetoothOBEXSessionClass
}

type _BluetoothOBEXSessionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BluetoothOBEXSession */
// An interface definition for the [BluetoothOBEXSession] class.
type IBluetoothOBEXSession interface {
	IOBEXSession
	
/* debug [class_interface_properties]: Properties for BluetoothOBEXSession */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BluetoothOBEXSession */
	// methods:
	CloseTransportConnection() OBEXError /* typedef */
	GetDevice() IBluetoothDevice
	GetRFCOMMChannel() IBluetoothRFCOMMChannel
	HasOpenTransportConnection() unsafe.Pointer
	IsSessionTargetAMac() bool
	OpenTransportConnectionSelectorTargetRefCon(inSelector objc.SEL, inTarget objc.IObject, inUserRefCon unsafe.Pointer) OBEXError /* typedef */
	RestartTransmission()
	SendBufferTroughChannel() int
	SendDataToTransportDataLength(inDataToSend unsafe.Pointer, inDataLength uintptr /* not a class type */) OBEXError /* typedef */
	SetOBEXSessionOpenConnectionCallbackRefCon(inCallback BluetoothOBEXSessionOpenConnectionCallback /* typedef */, inUserRefCon unsafe.Pointer)
	SetOpenTransportConnectionAsyncSelectorTargetRefCon(inSelector objc.SEL, inSelectorTarget objc.IObject, inUserRefCon unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BluetoothOBEXSession */
// Alloc allocates a new instance without initialization.
func (bc _BluetoothOBEXSessionClass) Alloc() BluetoothOBEXSession {
	rv := objc.Send[BluetoothOBEXSession](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BluetoothOBEXSessionClass) New() BluetoothOBEXSession {
	rv := objc.Send[BluetoothOBEXSession](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BluetoothOBEXSession) Init() BluetoothOBEXSession {
	rv := objc.Send[BluetoothOBEXSession](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BluetoothOBEXSession) Autorelease() BluetoothOBEXSession {
	rv := objc.Send[BluetoothOBEXSession](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBluetoothOBEXSession creates a new BluetoothOBEXSession instance.
func NewBluetoothOBEXSession() BluetoothOBEXSession {
	return getBluetoothOBEXSessionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BluetoothOBEXSession */
// An OBEX Session with a Bluetooth RFCOMM channel as the transport.


// An OBEX Session with a Bluetooth RFCOMM channel as the transport.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession
type BluetoothOBEXSession struct {
	OBEXSession
}

// BluetoothOBEXSessionFrom constructs a [BluetoothOBEXSession] from an unsafe.Pointer.
//
// An OBEX Session with a Bluetooth RFCOMM channel as the transport.
func BluetoothOBEXSessionFrom(ptr unsafe.Pointer) BluetoothOBEXSession {
	return BluetoothOBEXSession{
		OBEXSession: OBEXSessionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BluetoothOBEXSession */

// Initializes a Bluetooth-based OBEX Session using a Bluetooth device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/init(device:channelID:)
func NewBluetoothOBEXSessionWithDeviceChannelID(inDevice IOBluetoothDevice, inChannelID BluetoothRFCOMMChannelID /* typedef */) BluetoothOBEXSession {
	instance := getBluetoothOBEXSessionClass().Alloc()
	rv := objc.Send[BluetoothOBEXSession](instance.ID, objc.Sel("initWithDevice:channelID:"), inDevice, inChannelID)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewBluetoothOBEXSessionWithDeviceChannelID */


// Initializes a Bluetooth-based OBEX Session using an incoming RFCOMM channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/init(incomingRFCOMMChannel:eventSelector:selectorTarget:refCon:)
func NewBluetoothOBEXSessionWithIncomingRFCOMMChannelEventSelectorSelectorTargetRefCon(inChannel IOBluetoothRFCOMMChannel, inEventSelector objc.SEL, inEventSelectorTarget objc.IObject, inUserRefCon unsafe.Pointer) BluetoothOBEXSession {
	instance := getBluetoothOBEXSessionClass().Alloc()
	rv := objc.Send[BluetoothOBEXSession](instance.ID, objc.Sel("initWithIncomingRFCOMMChannel:eventSelector:selectorTarget:refCon:"), inChannel, inEventSelector, inEventSelectorTarget, inUserRefCon)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewBluetoothOBEXSessionWithIncomingRFCOMMChannelEventSelectorSelectorTargetRefCon */


// Initializes a Bluetooth-based OBEX Session using an SDP service record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/init(sdpServiceRecord:)
func NewBluetoothOBEXSessionWithSDPServiceRecord(inSDPServiceRecord IOBluetoothSDPServiceRecord) BluetoothOBEXSession {
	instance := getBluetoothOBEXSessionClass().Alloc()
	rv := objc.Send[BluetoothOBEXSession](instance.ID, objc.Sel("initWithSDPServiceRecord:"), inSDPServiceRecord)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewBluetoothOBEXSessionWithSDPServiceRecord */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BluetoothOBEXSession */

// Creates a Bluetooth-based OBEX Session using a Bluetooth device and a Bluetooth RFCOMM channel ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/withDevice(_:channelID:)
func (bc _BluetoothOBEXSessionClass) WithDeviceChannelID(inDevice IOBluetoothDevice, inRFCOMMChannelID BluetoothRFCOMMChannelID /* typedef */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("withDevice:channelID:"), inDevice, inRFCOMMChannelID)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WithDeviceChannelID) */


// Creates a Bluetooth-based OBEX Session using an incoming RFCOMM channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/withIncomingRFCOMMChannel(_:eventSelector:selectorTarget:refCon:)
func (bc _BluetoothOBEXSessionClass) WithIncomingRFCOMMChannelEventSelectorSelectorTargetRefCon(inChannel IOBluetoothRFCOMMChannel, inEventSelector objc.SEL, inEventSelectorTarget objc.IObject, inUserRefCon unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("withIncomingRFCOMMChannel:eventSelector:selectorTarget:refCon:"), inChannel, inEventSelector, inEventSelectorTarget, inUserRefCon)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WithIncomingRFCOMMChannelEventSelectorSelectorTargetRefCon) */


// Creates a Bluetooth-based OBEX Session using an SDP service record, typically obtained from a device/service browser window controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/withSDPServiceRecord(_:)
func (bc _BluetoothOBEXSessionClass) WithSDPServiceRecord(inSDPServiceRecord IOBluetoothSDPServiceRecord) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("withSDPServiceRecord:"), inSDPServiceRecord)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WithSDPServiceRecord) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BluetoothOBEXSession */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BluetoothOBEXSession */

// An OBEXSession override. When this is called by the session baseclass, we will close the transport connection if it is opened. In our case, it will be the RFCOMM channel that needs closing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/closeTransportConnection()
func (b_ BluetoothOBEXSession) CloseTransportConnection() OBEXError /* typedef */ {
	rv := objc.Send[int32](b_.ID, objc.Sel("closeTransportConnection"))
	return rv
}/* debug [instance_methods/method]: CloseTransportConnection */


// Get the Bluetooth Device being used by the session object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/getDevice()
func (b_ BluetoothOBEXSession) GetDevice() IBluetoothDevice {
	rv := objc.Send[BluetoothDevice](b_.ID, objc.Sel("getDevice"))
	return rv
}/* debug [instance_methods/method]: GetDevice */


// Get the Bluetooth RFCOMM channel being used by the session object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/getRFCOMMChannel()
func (b_ BluetoothOBEXSession) GetRFCOMMChannel() IBluetoothRFCOMMChannel {
	rv := objc.Send[BluetoothRFCOMMChannel](b_.ID, objc.Sel("getRFCOMMChannel"))
	return rv
}/* debug [instance_methods/method]: GetRFCOMMChannel */


// An OBEXSession override. When this is called by the session baseclass, we will return whether or not we have a transport connection established to another OBEX server/client. In our case we will tell whether or not the RFCOMM channel to a remote device is still open.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/hasOpenTransportConnection()
func (b_ BluetoothOBEXSession) HasOpenTransportConnection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("hasOpenTransportConnection"))
	return rv
}/* debug [instance_methods/method]: HasOpenTransportConnection */


// Tells whether the target device is a Mac by checking its service record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/isSessionTargetAMac()
func (b_ BluetoothOBEXSession) IsSessionTargetAMac() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isSessionTargetAMac"))
	return rv
}/* debug [instance_methods/method]: IsSessionTargetAMac */


// An OBEXSession override. When this is called by the session baseclass, we will attempt to open the transport connection. In our case, this would be an RFCOMM channel to another Bluetooth device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/openTransportConnection(_:selectorTarget:refCon:)
func (b_ BluetoothOBEXSession) OpenTransportConnectionSelectorTargetRefCon(inSelector objc.SEL, inTarget objc.IObject, inUserRefCon unsafe.Pointer) OBEXError /* typedef */ {
	rv := objc.Send[int32](b_.ID, objc.Sel("openTransportConnection:selectorTarget:refCon:"), inSelector, inTarget, inUserRefCon)
	return rv
}/* debug [instance_methods/method]: OpenTransportConnectionSelectorTargetRefCon */


// If the transmission was stopped due to the lack of buffers this call restarts it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/restartTransmission()
func (b_ BluetoothOBEXSession) RestartTransmission() {
	objc.Send[objc.ID](b_.ID, objc.Sel("restartTransmission"))
}/* debug [instance_methods/method]: RestartTransmission */


// Sends the next block of data through the rfcomm channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/sendBufferTroughChannel()
func (b_ BluetoothOBEXSession) SendBufferTroughChannel() int {
	rv := objc.Send[int](b_.ID, objc.Sel("sendBufferTroughChannel"))
	return rv
}/* debug [instance_methods/method]: SendBufferTroughChannel */


// An OBEXSession override. When this is called by the session baseclass, we will send the data we are given over our transport connection. If none is open, we could try to open it, or just return an error. In our case, it will be sent over the RFCOMM channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/sendData(toTransport:dataLength:)
func (b_ BluetoothOBEXSession) SendDataToTransportDataLength(inDataToSend unsafe.Pointer, inDataLength uintptr /* not a class type */) OBEXError /* typedef */ {
	rv := objc.Send[int32](b_.ID, objc.Sel("sendDataToTransport:dataLength:"), inDataToSend, inDataLength)
	return rv
}/* debug [instance_methods/method]: SendDataToTransportDataLength */


// For C API support. Allows you to set the callback to be invoked when the OBEX connection is actually opened.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/setOBEXSessionOpenConnectionCallback(_:refCon:)
func (b_ BluetoothOBEXSession) SetOBEXSessionOpenConnectionCallbackRefCon(inCallback BluetoothOBEXSessionOpenConnectionCallback /* typedef */, inUserRefCon unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setOBEXSessionOpenConnectionCallback:refCon:"), inCallback, inUserRefCon)
}/* debug [instance_methods/method]: SetOBEXSessionOpenConnectionCallbackRefCon */


// Allows you to set the selector to be used when a transport connection is opened, or fails to open.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/setOpenTransportConnectionAsyncSelector(_:target:refCon:)
func (b_ BluetoothOBEXSession) SetOpenTransportConnectionAsyncSelectorTargetRefCon(inSelector objc.SEL, inSelectorTarget objc.IObject, inUserRefCon unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setOpenTransportConnectionAsyncSelector:target:refCon:"), inSelector, inSelectorTarget, inUserRefCon)
}/* debug [instance_methods/method]: SetOpenTransportConnectionAsyncSelectorTargetRefCon */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BluetoothOBEXSession */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IOBluetoothOBEXSession */


