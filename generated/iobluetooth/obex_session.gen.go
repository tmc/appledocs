// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OBEXSession] class.
var (
	OBEXSessionClass     _OBEXSessionClass
	OBEXSessionClassOnce sync.Once
)

func getOBEXSessionClass() _OBEXSessionClass {
	OBEXSessionClassOnce.Do(func() {
		OBEXSessionClass = _OBEXSessionClass{objc.GetClass("OBEXSession")}
	})
	return OBEXSessionClass
}

type _OBEXSessionClass struct {
	class objc.Class
}

// An interface definition for the [OBEXSession] class.
type IOBEXSession interface {
	objectivec.IObject
	// properties:
	// methods:
	ClientHandleIncomingData(event unsafe.Pointer)
	CloseTransportConnection() OBEXError /* typedef */
	GetAvailableCommandPayloadLength(inOpCode OBEXOpCode /* typedef */) OBEXMaxPacketLength /* typedef */
	GetAvailableCommandResponsePayloadLength(inOpCode OBEXOpCode /* typedef */) OBEXMaxPacketLength /* typedef */
	GetMaxPacketLength() OBEXMaxPacketLength /* typedef */
	HasOpenOBEXConnection() bool /* primitive/slice/pointer. */
	HasOpenTransportConnection() unsafe.Pointer
	OBEXAbortOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr /* not a class type */, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon unsafe.Pointer) OBEXError /* typedef */
	OBEXAbortResponseOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inResponseOpCode OBEXOpCode /* typedef */, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr /* not a class type */, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon unsafe.Pointer) OBEXError /* typedef */
	OBEXConnectMaxPacketLengthOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inFlags OBEXFlags /* typedef */, inMaxPacketLength OBEXMaxPacketLength /* typedef */, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr /* not a class type */, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon unsafe.Pointer) OBEXError /* typedef */
	OBEXConnectResponseFlagsMaxPacketLengthOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inResponseOpCode OBEXOpCode /* typedef */, inFlags OBEXFlags /* typedef */, inMaxPacketLength OBEXMaxPacketLength /* typedef */, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr /* not a class type */, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon unsafe.Pointer) OBEXError /* typedef */
	OBEXDisconnectOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr /* not a class type */, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon unsafe.Pointer) OBEXError /* typedef */
	OBEXDisconnectResponseOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inResponseOpCode OBEXOpCode /* typedef */, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr /* not a class type */, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon unsafe.Pointer) OBEXError /* typedef */
	OBEXGetHeadersHeadersLengthEventSelectorSelectorTargetRefCon(isFinalChunk unsafe.Pointer, inHeaders unsafe.Pointer, inHeadersLength uintptr /* not a class type */, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon unsafe.Pointer) OBEXError /* typedef */
	OBEXGetResponseOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inResponseOpCode OBEXOpCode /* typedef */, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr /* not a class type */, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon unsafe.Pointer) OBEXError /* typedef */
	OBEXPutHeadersDataHeadersDataLengthBodyDataBodyDataLengthEventSelectorSelectorTargetRefCon(isFinalChunk unsafe.Pointer, inHeadersData unsafe.Pointer, inHeadersDataLength uintptr /* not a class type */, inBodyData unsafe.Pointer, inBodyDataLength uintptr /* not a class type */, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon unsafe.Pointer) OBEXError /* typedef */
	OBEXPutResponseOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inResponseOpCode OBEXOpCode /* typedef */, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr /* not a class type */, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon unsafe.Pointer) OBEXError /* typedef */
	OBEXSetPathConstantsOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inFlags OBEXFlags /* typedef */, inConstants OBEXConstants /* typedef */, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr /* not a class type */, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon unsafe.Pointer) OBEXError /* typedef */
	OBEXSetPathResponseOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inResponseOpCode OBEXOpCode /* typedef */, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr /* not a class type */, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon unsafe.Pointer) OBEXError /* typedef */
	OpenTransportConnectionSelectorTargetRefCon(inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon unsafe.Pointer) OBEXError /* typedef */
	SendDataToTransportDataLength(inDataToSend unsafe.Pointer, inDataLength uintptr /* not a class type */) OBEXError /* typedef */
	ServerHandleIncomingData(event unsafe.Pointer)
	SetEventCallback(inEventCallback OBEXSessionEventCallback /* typedef */)
	SetEventRefCon(inRefCon unsafe.Pointer)
	SetEventSelectorTargetRefCon(inEventSelector objc.SEL, inEventSelectorTarget objectivec.IObject, inUserRefCon unsafe.Pointer)
}

// Object representing an OBEX connection to a remote target.
//
// You will have no need for a obtaining/using a raw OBEXSession, since it requires an underlying transport to do anything useful. However, once you have an object that is a subclass of this class, you can use the functions herein to manipulate that OBEXSession. First off, you will want to use OBEXConnect (if you are a client session) to actually cause the transport to open a connection to a remote target and establish an OBEX connection over it. From there you can issue more commands based on the responses from a server. If you are a server session, the first thing you should receive is an OBEXConnect command packet, and you will want to issue an OBEXConnectResponse packet, with your reesponse to that command (success, denied, bad request, etc.). You can use the session accessors to access certain information, such as the negotiated max packet length. If you wish to implement your own OBEXSession over a transport such as ethernet, you will need to see the end of the file to determine which functions to override, and what to pass to those functions. No timeout mechanism has been implemented so far for an OBEXSessions. If you need timeouts, you will need to implement them yourself. This is being explored for a future revision. However, be aware that the OBEX Specification does not explicitly require timeouts, so be sure you allow ample time for commands to complete, as some devices may be slow when sending large amounts of data.


// Object representing an OBEX connection to a remote target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession
type OBEXSession struct {
	objectivec.Object
}

// OBEXSessionFrom constructs a [OBEXSession] from an unsafe.Pointer.
//
// Object representing an OBEX connection to a remote target.
func OBEXSessionFrom(ptr unsafe.Pointer) OBEXSession {
	return OBEXSession{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _OBEXSessionClass) Alloc() OBEXSession {
	rv := objc.Send[OBEXSession](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OBEXSessionClass) New() OBEXSession {
	rv := objc.Send[OBEXSession](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OBEXSession) Init() OBEXSession {
	rv := objc.Send[OBEXSession](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OBEXSession) Autorelease() OBEXSession {
	rv := objc.Send[OBEXSession](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOBEXSession creates a new OBEXSession instance.
func NewOBEXSession() OBEXSession {
	return getOBEXSessionClass().New()
}



// Tranport subclasses need to invoke this from their own data-receive handlers. For example, when data is received over a Bluetooth RFCOMM channel in the IOBluetoothOBEXSession, it in turn calls this to dispatch the data. If you do not handle this case, your server session will not work, guaranteed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/clientHandleIncomingData(_:)
func (o_ OBEXSession) ClientHandleIncomingData(event unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("clientHandleIncomingData:"), event)
}


// You must override this - it will be called when the transport connection should be shutdown.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/closeTransportConnection()
func (o_ OBEXSession) CloseTransportConnection() OBEXError /* typedef */ {
	rv := objc.Send[OBEXError](o_.ID, objc.Sel("closeTransportConnection"))
	return rv
}


// Determine the maximum amount of data you can send in a particular command as an OBEX client session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/getAvailableCommandPayloadLength(_:)
func (o_ OBEXSession) GetAvailableCommandPayloadLength(inOpCode OBEXOpCode /* typedef */) OBEXMaxPacketLength /* typedef */ {
	rv := objc.Send[OBEXMaxPacketLength](o_.ID, objc.Sel("getAvailableCommandPayloadLength:"), inOpCode)
	return rv
}


// Determine the maximum amount of data you can send in a particular command response as an OBEX server session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/getAvailableCommandResponsePayloadLength(_:)
func (o_ OBEXSession) GetAvailableCommandResponsePayloadLength(inOpCode OBEXOpCode /* typedef */) OBEXMaxPacketLength /* typedef */ {
	rv := objc.Send[OBEXMaxPacketLength](o_.ID, objc.Sel("getAvailableCommandResponsePayloadLength:"), inOpCode)
	return rv
}


// Gets current max packet length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/getMaxPacketLength()
func (o_ OBEXSession) GetMaxPacketLength() OBEXMaxPacketLength /* typedef */ {
	rv := objc.Send[OBEXMaxPacketLength](o_.ID, objc.Sel("getMaxPacketLength"))
	return rv
}


// Has a successful connect packet been sent and received? This API tells you so.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/hasOpenOBEXConnection()
func (o_ OBEXSession) HasOpenOBEXConnection() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("hasOpenOBEXConnection"))
	return rv
}


// You must override this - it will be called periodically to determine if a transport connection is open or not.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/hasOpenTransportConnection()
func (o_ OBEXSession) HasOpenTransportConnection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("hasOpenTransportConnection"))
	return rv
}


// Send an OBEX Abort command to the session’s target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/obexAbort(_:optionalHeadersLength:eventSelector:selectorTarget:refCon:)
func (o_ OBEXSession) OBEXAbortOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr /* not a class type */, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon unsafe.Pointer) OBEXError /* typedef */ {
	rv := objc.Send[OBEXError](o_.ID, objc.Sel("OBEXAbort:optionalHeadersLength:eventSelector:selectorTarget:refCon:"), inOptionalHeaders, inOptionalHeadersLength, inSelector, inTarget, inUserRefCon)
	return rv
}


// Send an abort response to a session’s target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/obexAbortResponse(_:optionalHeaders:optionalHeadersLength:eventSelector:selectorTarget:refCon:)
func (o_ OBEXSession) OBEXAbortResponseOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inResponseOpCode OBEXOpCode /* typedef */, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr /* not a class type */, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon unsafe.Pointer) OBEXError /* typedef */ {
	rv := objc.Send[OBEXError](o_.ID, objc.Sel("OBEXAbortResponse:optionalHeaders:optionalHeadersLength:eventSelector:selectorTarget:refCon:"), inResponseOpCode, inOptionalHeaders, inOptionalHeadersLength, inSelector, inTarget, inUserRefCon)
	return rv
}


// Initiate an OBEX connection to a device. Causes underlying transport (Bluetooth, et al) to attempt to connect to a remote device. After success, an OBEX connect packet is sent to establish the OBEX Connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/obexConnect(_:maxPacketLength:optionalHeaders:optionalHeadersLength:eventSelector:selectorTarget:refCon:)
func (o_ OBEXSession) OBEXConnectMaxPacketLengthOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inFlags OBEXFlags /* typedef */, inMaxPacketLength OBEXMaxPacketLength /* typedef */, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr /* not a class type */, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon unsafe.Pointer) OBEXError /* typedef */ {
	rv := objc.Send[OBEXError](o_.ID, objc.Sel("OBEXConnect:maxPacketLength:optionalHeaders:optionalHeadersLength:eventSelector:selectorTarget:refCon:"), inFlags, inMaxPacketLength, inOptionalHeaders, inOptionalHeadersLength, inSelector, inTarget, inUserRefCon)
	return rv
}


// Send a connect response to a session’s target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/obexConnectResponse(_:flags:maxPacketLength:optionalHeaders:optionalHeadersLength:eventSelector:selectorTarget:refCon:)
func (o_ OBEXSession) OBEXConnectResponseFlagsMaxPacketLengthOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inResponseOpCode OBEXOpCode /* typedef */, inFlags OBEXFlags /* typedef */, inMaxPacketLength OBEXMaxPacketLength /* typedef */, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr /* not a class type */, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon unsafe.Pointer) OBEXError /* typedef */ {
	rv := objc.Send[OBEXError](o_.ID, objc.Sel("OBEXConnectResponse:flags:maxPacketLength:optionalHeaders:optionalHeadersLength:eventSelector:selectorTarget:refCon:"), inResponseOpCode, inFlags, inMaxPacketLength, inOptionalHeaders, inOptionalHeadersLength, inSelector, inTarget, inUserRefCon)
	return rv
}


// Send an OBEX Disconnect command to the session’s target. THIS DOES NOT necessarily close the underlying transport connection. Deleting the session will ensure that closure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/obexDisconnect(_:optionalHeadersLength:eventSelector:selectorTarget:refCon:)
func (o_ OBEXSession) OBEXDisconnectOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr /* not a class type */, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon unsafe.Pointer) OBEXError /* typedef */ {
	rv := objc.Send[OBEXError](o_.ID, objc.Sel("OBEXDisconnect:optionalHeadersLength:eventSelector:selectorTarget:refCon:"), inOptionalHeaders, inOptionalHeadersLength, inSelector, inTarget, inUserRefCon)
	return rv
}


// Send a disconnect response to a session’s target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/obexDisconnectResponse(_:optionalHeaders:optionalHeadersLength:eventSelector:selectorTarget:refCon:)
func (o_ OBEXSession) OBEXDisconnectResponseOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inResponseOpCode OBEXOpCode /* typedef */, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr /* not a class type */, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon unsafe.Pointer) OBEXError /* typedef */ {
	rv := objc.Send[OBEXError](o_.ID, objc.Sel("OBEXDisconnectResponse:optionalHeaders:optionalHeadersLength:eventSelector:selectorTarget:refCon:"), inResponseOpCode, inOptionalHeaders, inOptionalHeadersLength, inSelector, inTarget, inUserRefCon)
	return rv
}


// Send an OBEX Get command to the session’s target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/obexGet(_:headers:headersLength:eventSelector:selectorTarget:refCon:)
func (o_ OBEXSession) OBEXGetHeadersHeadersLengthEventSelectorSelectorTargetRefCon(isFinalChunk unsafe.Pointer, inHeaders unsafe.Pointer, inHeadersLength uintptr /* not a class type */, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon unsafe.Pointer) OBEXError /* typedef */ {
	rv := objc.Send[OBEXError](o_.ID, objc.Sel("OBEXGet:headers:headersLength:eventSelector:selectorTarget:refCon:"), isFinalChunk, inHeaders, inHeadersLength, inSelector, inTarget, inUserRefCon)
	return rv
}


// Send a get response to a session’s target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/obexGetResponse(_:optionalHeaders:optionalHeadersLength:eventSelector:selectorTarget:refCon:)
func (o_ OBEXSession) OBEXGetResponseOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inResponseOpCode OBEXOpCode /* typedef */, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr /* not a class type */, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon unsafe.Pointer) OBEXError /* typedef */ {
	rv := objc.Send[OBEXError](o_.ID, objc.Sel("OBEXGetResponse:optionalHeaders:optionalHeadersLength:eventSelector:selectorTarget:refCon:"), inResponseOpCode, inOptionalHeaders, inOptionalHeadersLength, inSelector, inTarget, inUserRefCon)
	return rv
}


// Send an OBEX Put command to the session’s target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/obexPut(_:headersData:headersDataLength:bodyData:bodyDataLength:eventSelector:selectorTarget:refCon:)
func (o_ OBEXSession) OBEXPutHeadersDataHeadersDataLengthBodyDataBodyDataLengthEventSelectorSelectorTargetRefCon(isFinalChunk unsafe.Pointer, inHeadersData unsafe.Pointer, inHeadersDataLength uintptr /* not a class type */, inBodyData unsafe.Pointer, inBodyDataLength uintptr /* not a class type */, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon unsafe.Pointer) OBEXError /* typedef */ {
	rv := objc.Send[OBEXError](o_.ID, objc.Sel("OBEXPut:headersData:headersDataLength:bodyData:bodyDataLength:eventSelector:selectorTarget:refCon:"), isFinalChunk, inHeadersData, inHeadersDataLength, inBodyData, inBodyDataLength, inSelector, inTarget, inUserRefCon)
	return rv
}


// Send a put response to a session’s target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/obexPutResponse(_:optionalHeaders:optionalHeadersLength:eventSelector:selectorTarget:refCon:)
func (o_ OBEXSession) OBEXPutResponseOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inResponseOpCode OBEXOpCode /* typedef */, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr /* not a class type */, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon unsafe.Pointer) OBEXError /* typedef */ {
	rv := objc.Send[OBEXError](o_.ID, objc.Sel("OBEXPutResponse:optionalHeaders:optionalHeadersLength:eventSelector:selectorTarget:refCon:"), inResponseOpCode, inOptionalHeaders, inOptionalHeadersLength, inSelector, inTarget, inUserRefCon)
	return rv
}


// Send an OBEX SetPath command to the session’s target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/obexSetPath(_:constants:optionalHeaders:optionalHeadersLength:eventSelector:selectorTarget:refCon:)
func (o_ OBEXSession) OBEXSetPathConstantsOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inFlags OBEXFlags /* typedef */, inConstants OBEXConstants /* typedef */, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr /* not a class type */, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon unsafe.Pointer) OBEXError /* typedef */ {
	rv := objc.Send[OBEXError](o_.ID, objc.Sel("OBEXSetPath:constants:optionalHeaders:optionalHeadersLength:eventSelector:selectorTarget:refCon:"), inFlags, inConstants, inOptionalHeaders, inOptionalHeadersLength, inSelector, inTarget, inUserRefCon)
	return rv
}


// Send a set path response to a session’s target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/obexSetPathResponse(_:optionalHeaders:optionalHeadersLength:eventSelector:selectorTarget:refCon:)
func (o_ OBEXSession) OBEXSetPathResponseOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inResponseOpCode OBEXOpCode /* typedef */, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr /* not a class type */, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon unsafe.Pointer) OBEXError /* typedef */ {
	rv := objc.Send[OBEXError](o_.ID, objc.Sel("OBEXSetPathResponse:optionalHeaders:optionalHeadersLength:eventSelector:selectorTarget:refCon:"), inResponseOpCode, inOptionalHeaders, inOptionalHeadersLength, inSelector, inTarget, inUserRefCon)
	return rv
}


// Opens a transport connection to a device. A Bluetooth connection is one example of a transport.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/openTransportConnection(_:selectorTarget:refCon:)
func (o_ OBEXSession) OpenTransportConnectionSelectorTargetRefCon(inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon unsafe.Pointer) OBEXError /* typedef */ {
	rv := objc.Send[OBEXError](o_.ID, objc.Sel("openTransportConnection:selectorTarget:refCon:"), inSelector, inTarget, inUserRefCon)
	return rv
}


// You must override this to send data over your transport. This does nothing by default, it will return a kOBEXUnsupportedError.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/sendData(toTransport:dataLength:)
func (o_ OBEXSession) SendDataToTransportDataLength(inDataToSend unsafe.Pointer, inDataLength uintptr /* not a class type */) OBEXError /* typedef */ {
	rv := objc.Send[OBEXError](o_.ID, objc.Sel("sendDataToTransport:dataLength:"), inDataToSend, inDataLength)
	return rv
}


// Tranport subclasses need to invoke this from their own data-receive handlers. For example, when data is received over a Bluetooth RFCOMM channel in the IOBluetoothOBEXSession, it in turn calls this to dispatch the data. If you do not handle this case, your server session will not work, guaranteed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/serverHandleIncomingData(_:)
func (o_ OBEXSession) ServerHandleIncomingData(event unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("serverHandleIncomingData:"), event)
}


// Sets the C-API callback used when the session recieves data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/setEventCallback(_:)
func (o_ OBEXSession) SetEventCallback(inEventCallback OBEXSessionEventCallback /* typedef */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setEventCallback:"), inEventCallback)
}


// Sets the C-API callback refCon used when the session recieves data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/setEventRefCon(_:)
func (o_ OBEXSession) SetEventRefCon(inRefCon unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setEventRefCon:"), inRefCon)
}


// Allow you to set a selector to be called when events occur on the OBEX session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/setEventSelector(_:target:refCon:)
func (o_ OBEXSession) SetEventSelectorTargetRefCon(inEventSelector objc.SEL, inEventSelectorTarget objectivec.IObject, inUserRefCon unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setEventSelector:target:refCon:"), inEventSelector, inEventSelectorTarget, inUserRefCon)
}



