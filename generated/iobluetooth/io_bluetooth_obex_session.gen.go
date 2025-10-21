// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [BluetoothOBEXSession] class.
type IBluetoothOBEXSession interface {
	IOBEXSession
	CloseTransportConnection() OBEXError
	GetDevice() BluetoothDevice
	GetRFCOMMChannel() BluetoothRFCOMMChannel
	HasOpenTransportConnection() unsafe.Pointer
	IsSessionTargetAMac() bool
	OpenTransportConnectionSelectorTargetRefCon(inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon unsafe.Pointer) OBEXError
	RestartTransmission()
	SendBufferTroughChannel() unsafe.Pointer
	SendDataToTransportDataLength(inDataToSend unsafe.Pointer, inDataLength unsafe.Pointer) OBEXError
	SetOBEXSessionOpenConnectionCallbackRefCon(inCallback IBluetoothOBEXSessionOpenConnectionCallback, inUserRefCon unsafe.Pointer)
	SetOpenTransportConnectionAsyncSelectorTargetRefCon(inSelector objc.SEL, inSelectorTarget objectivec.IObject, inUserRefCon unsafe.Pointer)
}

// An OBEX Session with a Bluetooth RFCOMM channel as the transport.
//
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

// Alloc allocates a new instance without initialization.
func (bc _BluetoothOBEXSessionClass) Alloc() BluetoothOBEXSession {
	rv := objc.Send[BluetoothOBEXSession](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Initializes a Bluetooth-based OBEX Session using a Bluetooth device.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/init(device:channelID:)
func NewBluetoothOBEXSessionWithDeviceChannelID(inDevice IOBluetoothDevice, inChannelID IBluetoothRFCOMMChannelID) BluetoothOBEXSession {
	instance := getBluetoothOBEXSessionClass().Alloc()
	rv := objc.Send[BluetoothOBEXSession](instance.ID, objc.Sel("initWithDevice:channelID:"), inDevice, inChannelID)
	rv.Autorelease()
	return rv
}



// Initializes a Bluetooth-based OBEX Session using an incoming RFCOMM channel.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/init(incomingRFCOMMChannel:eventSelector:selectorTarget:refCon:)
func NewBluetoothOBEXSessionWithIncomingRFCOMMChannelEventSelectorSelectorTargetRefCon(inChannel IOBluetoothRFCOMMChannel, inEventSelector objc.SEL, inEventSelectorTarget objectivec.IObject, inUserRefCon unsafe.Pointer) BluetoothOBEXSession {
	instance := getBluetoothOBEXSessionClass().Alloc()
	rv := objc.Send[BluetoothOBEXSession](instance.ID, objc.Sel("initWithIncomingRFCOMMChannel:eventSelector:selectorTarget:refCon:"), inChannel, inEventSelector, inEventSelectorTarget, inUserRefCon)
	rv.Autorelease()
	return rv
}



// Initializes a Bluetooth-based OBEX Session using an SDP service record.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/init(sdpServiceRecord:)
func NewBluetoothOBEXSessionWithSDPServiceRecord(inSDPServiceRecord IOBluetoothSDPServiceRecord) BluetoothOBEXSession {
	instance := getBluetoothOBEXSessionClass().Alloc()
	rv := objc.Send[BluetoothOBEXSession](instance.ID, objc.Sel("initWithSDPServiceRecord:"), inSDPServiceRecord)
	rv.Autorelease()
	return rv
}


// Creates a Bluetooth-based OBEX Session using a Bluetooth device and a Bluetooth RFCOMM channel ID.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/withDevice(_:channelID:)
func (bc _BluetoothOBEXSessionClass) WithDeviceChannelID(inDevice IOBluetoothDevice, inRFCOMMChannelID IBluetoothRFCOMMChannelID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("withDevice:channelID:"), inDevice, inRFCOMMChannelID)
	return rv
}

// Creates a Bluetooth-based OBEX Session using an incoming RFCOMM channel.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/withIncomingRFCOMMChannel(_:eventSelector:selectorTarget:refCon:)
func (bc _BluetoothOBEXSessionClass) WithIncomingRFCOMMChannelEventSelectorSelectorTargetRefCon(inChannel IOBluetoothRFCOMMChannel, inEventSelector objc.SEL, inEventSelectorTarget objectivec.IObject, inUserRefCon unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("withIncomingRFCOMMChannel:eventSelector:selectorTarget:refCon:"), inChannel, inEventSelector, inEventSelectorTarget, inUserRefCon)
	return rv
}

// Creates a Bluetooth-based OBEX Session using an SDP service record, typically obtained from a device/service browser window controller.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/withSDPServiceRecord(_:)
func (bc _BluetoothOBEXSessionClass) WithSDPServiceRecord(inSDPServiceRecord IOBluetoothSDPServiceRecord) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("withSDPServiceRecord:"), inSDPServiceRecord)
	return rv
}

// An OBEXSession override. When this is called by the session baseclass, we will close the transport connection if it is opened. In our case, it will be the RFCOMM channel that needs closing.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/closeTransportConnection()
func (b_ BluetoothOBEXSession) CloseTransportConnection() OBEXError {
	rv := objc.Send[OBEXError](b_.ID, objc.Sel("closeTransportConnection"))
	return rv
}

// Get the Bluetooth Device being used by the session object.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/getDevice()
func (b_ BluetoothOBEXSession) GetDevice() BluetoothDevice {
	rv := objc.Send[BluetoothDevice](b_.ID, objc.Sel("getDevice"))
	return rv
}

// Get the Bluetooth RFCOMM channel being used by the session object.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/getRFCOMMChannel()
func (b_ BluetoothOBEXSession) GetRFCOMMChannel() BluetoothRFCOMMChannel {
	rv := objc.Send[BluetoothRFCOMMChannel](b_.ID, objc.Sel("getRFCOMMChannel"))
	return rv
}

// An OBEXSession override. When this is called by the session baseclass, we will return whether or not we have a transport connection established to another OBEX server/client. In our case we will tell whether or not the RFCOMM channel to a remote device is still open.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/hasOpenTransportConnection()
func (b_ BluetoothOBEXSession) HasOpenTransportConnection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("hasOpenTransportConnection"))
	return rv
}

// Tells whether the target device is a Mac by checking its service record.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/isSessionTargetAMac()
func (b_ BluetoothOBEXSession) IsSessionTargetAMac() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isSessionTargetAMac"))
	return rv
}

// An OBEXSession override. When this is called by the session baseclass, we will attempt to open the transport connection. In our case, this would be an RFCOMM channel to another Bluetooth device.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/openTransportConnection(_:selectorTarget:refCon:)
func (b_ BluetoothOBEXSession) OpenTransportConnectionSelectorTargetRefCon(inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon unsafe.Pointer) OBEXError {
	rv := objc.Send[OBEXError](b_.ID, objc.Sel("openTransportConnection:selectorTarget:refCon:"), inSelector, inTarget, inUserRefCon)
	return rv
}

// If the transmission was stopped due to the lack of buffers this call restarts it.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/restartTransmission()
func (b_ BluetoothOBEXSession) RestartTransmission() {
	objc.Send[objc.ID](b_.ID, objc.Sel("restartTransmission"))
}

// Sends the next block of data through the rfcomm channel.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/sendBufferTroughChannel()
func (b_ BluetoothOBEXSession) SendBufferTroughChannel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("sendBufferTroughChannel"))
	return rv
}

// An OBEXSession override. When this is called by the session baseclass, we will send the data we are given over our transport connection. If none is open, we could try to open it, or just return an error. In our case, it will be sent over the RFCOMM channel.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/sendData(toTransport:dataLength:)
func (b_ BluetoothOBEXSession) SendDataToTransportDataLength(inDataToSend unsafe.Pointer, inDataLength unsafe.Pointer) OBEXError {
	rv := objc.Send[OBEXError](b_.ID, objc.Sel("sendDataToTransport:dataLength:"), inDataToSend, inDataLength)
	return rv
}

// For C API support. Allows you to set the callback to be invoked when the OBEX connection is actually opened.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/setOBEXSessionOpenConnectionCallback(_:refCon:)
func (b_ BluetoothOBEXSession) SetOBEXSessionOpenConnectionCallbackRefCon(inCallback IBluetoothOBEXSessionOpenConnectionCallback, inUserRefCon unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setOBEXSessionOpenConnectionCallback:refCon:"), inCallback, inUserRefCon)
}

// Allows you to set the selector to be used when a transport connection is opened, or fails to open.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/setOpenTransportConnectionAsyncSelector(_:target:refCon:)
func (b_ BluetoothOBEXSession) SetOpenTransportConnectionAsyncSelectorTargetRefCon(inSelector objc.SEL, inSelectorTarget objectivec.IObject, inUserRefCon unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setOpenTransportConnectionAsyncSelector:target:refCon:"), inSelector, inSelectorTarget, inUserRefCon)
}


