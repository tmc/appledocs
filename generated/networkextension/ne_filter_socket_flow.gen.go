// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NEFilterSocketFlow] class.
var (
	NEFilterSocketFlowClass     _NEFilterSocketFlowClass
	NEFilterSocketFlowClassOnce sync.Once
)

func getNEFilterSocketFlowClass() _NEFilterSocketFlowClass {
	NEFilterSocketFlowClassOnce.Do(func() {
		NEFilterSocketFlowClass = _NEFilterSocketFlowClass{objc.GetClass("NEFilterSocketFlow")}
	})
	return NEFilterSocketFlowClass
}

type _NEFilterSocketFlowClass struct {
	class objc.Class
}

// An interface definition for the [NEFilterSocketFlow] class.
type INEFilterSocketFlow interface {
	INEFilterFlow
	LocalEndpoint() NWEndpoint
	SetLocalEndpoint(value INWEndpoint)
	LocalFlowEndpoint() NWEndpoint
	SetLocalFlowEndpoint(value INWEndpoint)
	RemoteEndpoint() NWEndpoint
	SetRemoteEndpoint(value INWEndpoint)
	RemoteFlowEndpoint() NWEndpoint
	SetRemoteFlowEndpoint(value INWEndpoint)
	RemoteHostname() string
	SetRemoteHostname(value string)
	SocketFamily() unsafe.Pointer
	SetSocketFamily(value unsafe.Pointer)
	SocketProtocol() unsafe.Pointer
	SetSocketProtocol(value unsafe.Pointer)
	SocketType() unsafe.Pointer
	SetSocketType(value unsafe.Pointer)
}

// A flow of network data that the filter examines.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterSocketFlow
type NEFilterSocketFlow struct {
	NEFilterFlow
}

// NEFilterSocketFlowFrom constructs a [NEFilterSocketFlow] from an unsafe.Pointer.
//
// A flow of network data that the filter examines.
func NEFilterSocketFlowFrom(ptr unsafe.Pointer) NEFilterSocketFlow {
	return NEFilterSocketFlow{
		NEFilterFlow: NEFilterFlowFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NEFilterSocketFlowClass) Alloc() NEFilterSocketFlow {
	rv := objc.Send[NEFilterSocketFlow](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEFilterSocketFlowClass) New() NEFilterSocketFlow {
	rv := objc.Send[NEFilterSocketFlow](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEFilterSocketFlow) Init() NEFilterSocketFlow {
	rv := objc.Send[NEFilterSocketFlow](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEFilterSocketFlow) Autorelease() NEFilterSocketFlow {
	rv := objc.Send[NEFilterSocketFlow](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterSocketFlow creates a new NEFilterSocketFlow instance.
func NewNEFilterSocketFlow() NEFilterSocketFlow {
	return getNEFilterSocketFlowClass().New()
}


// An object containing details about the socket’s local endpoint.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefiltersocketflow/localendpoint
func (n_ NEFilterSocketFlow) LocalEndpoint() NWEndpoint {
	rv := objc.Send[NWEndpoint](n_.ID, objc.Sel("localEndpoint"))
	return rv
}


// SetLocalEndpoint sets the value of the localEndpoint property.
// An object containing details about the socket’s local endpoint.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefiltersocketflow/localendpoint
func (n_ NEFilterSocketFlow) SetLocalEndpoint(value INWEndpoint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setLocalEndpoint:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefiltersocketflow/localflowendpoint-89z3l
func (n_ NEFilterSocketFlow) LocalFlowEndpoint() NWEndpoint {
	rv := objc.Send[NWEndpoint](n_.ID, objc.Sel("localFlowEndpoint"))
	return rv
}


// SetLocalFlowEndpoint sets the value of the localFlowEndpoint property.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefiltersocketflow/localflowendpoint-89z3l
func (n_ NEFilterSocketFlow) SetLocalFlowEndpoint(value INWEndpoint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setLocalFlowEndpoint:"), value)
}

// An object containing details about the socket’s remote endpoint.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefiltersocketflow/remoteendpoint
func (n_ NEFilterSocketFlow) RemoteEndpoint() NWEndpoint {
	rv := objc.Send[NWEndpoint](n_.ID, objc.Sel("remoteEndpoint"))
	return rv
}


// SetRemoteEndpoint sets the value of the remoteEndpoint property.
// An object containing details about the socket’s remote endpoint.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefiltersocketflow/remoteendpoint
func (n_ NEFilterSocketFlow) SetRemoteEndpoint(value INWEndpoint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setRemoteEndpoint:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefiltersocketflow/remoteflowendpoint-6bnas
func (n_ NEFilterSocketFlow) RemoteFlowEndpoint() NWEndpoint {
	rv := objc.Send[NWEndpoint](n_.ID, objc.Sel("remoteFlowEndpoint"))
	return rv
}


// SetRemoteFlowEndpoint sets the value of the remoteFlowEndpoint property.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefiltersocketflow/remoteflowendpoint-6bnas
func (n_ NEFilterSocketFlow) SetRemoteFlowEndpoint(value INWEndpoint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setRemoteFlowEndpoint:"), value)
}

// The flow’s remote hostname, if applicable.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefiltersocketflow/remotehostname
func (n_ NEFilterSocketFlow) RemoteHostname() string {
	rv := objc.Send[string](n_.ID, objc.Sel("remoteHostname"))
	return rv
}


// SetRemoteHostname sets the value of the remoteHostname property.
// The flow’s remote hostname, if applicable.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefiltersocketflow/remotehostname
func (n_ NEFilterSocketFlow) SetRemoteHostname(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setRemoteHostname:"), objc.String(value))
}

// The protocol family of the socket.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefiltersocketflow/socketfamily
func (n_ NEFilterSocketFlow) SocketFamily() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("socketFamily"))
	return rv
}


// SetSocketFamily sets the value of the socketFamily property.
// The protocol family of the socket.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefiltersocketflow/socketfamily
func (n_ NEFilterSocketFlow) SetSocketFamily(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSocketFamily:"), value)
}

// The protocol of the socket.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefiltersocketflow/socketprotocol
func (n_ NEFilterSocketFlow) SocketProtocol() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("socketProtocol"))
	return rv
}


// SetSocketProtocol sets the value of the socketProtocol property.
// The protocol of the socket.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefiltersocketflow/socketprotocol
func (n_ NEFilterSocketFlow) SetSocketProtocol(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSocketProtocol:"), value)
}

// The type of the socket.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefiltersocketflow/sockettype
func (n_ NEFilterSocketFlow) SocketType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("socketType"))
	return rv
}


// SetSocketType sets the value of the socketType property.
// The type of the socket.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefiltersocketflow/sockettype
func (n_ NEFilterSocketFlow) SetSocketType(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSocketType:"), value)
}



