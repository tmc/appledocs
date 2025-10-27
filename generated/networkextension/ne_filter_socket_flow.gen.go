// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	

	// properties:
	LocalEndpoint() INWEndpoint
	LocalFlowEndpoint() objectivec.IObject
	RemoteEndpoint() INWEndpoint
	RemoteFlowEndpoint() objectivec.IObject
	RemoteHostname() foundation.foundation.INSString
	SocketFamily() int
	SocketProtocol() int
	SocketType() int


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (nc _NEFilterSocketFlowClass) Alloc() NEFilterSocketFlow {
	rv := objc.Send[NEFilterSocketFlow](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// A flow of network data that the filter examines.


// A flow of network data that the filter examines.
//
// [Full Topic]
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

























// An object containing details about the socket’s local endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterSocketFlow/localEndpoint
func (n_ NEFilterSocketFlow) LocalEndpoint() INWEndpoint {
	rv := objc.Send[NWEndpoint](n_.ID, objc.Sel("localEndpoint"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterSocketFlow/localFlowEndpoint-4nt54
func (n_ NEFilterSocketFlow) LocalFlowEndpoint() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("localFlowEndpoint"))
	return rv
}


// An object containing details about the socket’s remote endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterSocketFlow/remoteEndpoint
func (n_ NEFilterSocketFlow) RemoteEndpoint() INWEndpoint {
	rv := objc.Send[NWEndpoint](n_.ID, objc.Sel("remoteEndpoint"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterSocketFlow/remoteFlowEndpoint-52dxr
func (n_ NEFilterSocketFlow) RemoteFlowEndpoint() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("remoteFlowEndpoint"))
	return rv
}


// The flow’s remote hostname, if applicable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterSocketFlow/remoteHostname
func (n_ NEFilterSocketFlow) RemoteHostname() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("remoteHostname"))
	return rv
}


// The protocol family of the socket.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterSocketFlow/socketFamily
func (n_ NEFilterSocketFlow) SocketFamily() int {
	rv := objc.Send[int](n_.ID, objc.Sel("socketFamily"))
	return rv
}


// The protocol of the socket.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterSocketFlow/socketProtocol
func (n_ NEFilterSocketFlow) SocketProtocol() int {
	rv := objc.Send[int](n_.ID, objc.Sel("socketProtocol"))
	return rv
}


// The type of the socket.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterSocketFlow/socketType
func (n_ NEFilterSocketFlow) SocketType() int {
	rv := objc.Send[int](n_.ID, objc.Sel("socketType"))
	return rv
}








