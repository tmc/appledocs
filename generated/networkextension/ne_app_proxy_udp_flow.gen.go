// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [NEAppProxyUDPFlow] class.
var (
	NEAppProxyUDPFlowClass     _NEAppProxyUDPFlowClass
	NEAppProxyUDPFlowClassOnce sync.Once
)

func getNEAppProxyUDPFlowClass() _NEAppProxyUDPFlowClass {
	NEAppProxyUDPFlowClassOnce.Do(func() {
		NEAppProxyUDPFlowClass = _NEAppProxyUDPFlowClass{objc.GetClass("NEAppProxyUDPFlow")}
	})
	return NEAppProxyUDPFlowClass
}

type _NEAppProxyUDPFlowClass struct {
	class objc.Class
}

// An interface definition for the [NEAppProxyUDPFlow] class.
type INEAppProxyUDPFlow interface {
	INEAppProxyFlow
	// properties:
	LocalEndpoint() INWEndpoint
	LocalFlowEndpoint() INWEndpoint
	SetLocalFlowEndpoint(value INWEndpoint)
	// methods:
	WriteDatagramsSentByFlowEndpointsCompletionHandler(datagrams []objc.IObject /* cross-framework: Data */, remoteEndpoints unsafe.Pointer, completionHandler unsafe.Pointer)
}

// An object for reading and writing data to and from a UDP conversation being proxied by the provider.
//
// App Proxy Providers receive UDP connections to be proxied in the form of objects.


// An object for reading and writing data to and from a UDP conversation being proxied by the provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyUDPFlow
type NEAppProxyUDPFlow struct {
	NEAppProxyFlow
}

// NEAppProxyUDPFlowFrom constructs a [NEAppProxyUDPFlow] from an unsafe.Pointer.
//
// An object for reading and writing data to and from a UDP conversation being proxied by the provider.
func NEAppProxyUDPFlowFrom(ptr unsafe.Pointer) NEAppProxyUDPFlow {
	return NEAppProxyUDPFlow{
		NEAppProxyFlow: NEAppProxyFlowFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NEAppProxyUDPFlowClass) Alloc() NEAppProxyUDPFlow {
	rv := objc.Send[NEAppProxyUDPFlow](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEAppProxyUDPFlowClass) New() NEAppProxyUDPFlow {
	rv := objc.Send[NEAppProxyUDPFlow](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEAppProxyUDPFlow) Init() NEAppProxyUDPFlow {
	rv := objc.Send[NEAppProxyUDPFlow](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEAppProxyUDPFlow) Autorelease() NEAppProxyUDPFlow {
	rv := objc.Send[NEAppProxyUDPFlow](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEAppProxyUDPFlow creates a new NEAppProxyUDPFlow instance.
func NewNEAppProxyUDPFlow() NEAppProxyUDPFlow {
	return getNEAppProxyUDPFlowClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyUDPFlow/writeDatagrams:sentByFlowEndpoints:completionHandler:
func (n_ NEAppProxyUDPFlow) WriteDatagramsSentByFlowEndpointsCompletionHandler(datagrams []objc.IObject /* cross-framework: Data */, remoteEndpoints unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("writeDatagrams:sentByFlowEndpoints:completionHandler:"), datagrams, remoteEndpoints, completionHandler)
}


// An object containing information about the local endpoint of the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyUDPFlow/localEndpoint
func (n_ NEAppProxyUDPFlow) LocalEndpoint() INWEndpoint {
	rv := objc.Send[NWEndpoint](n_.ID, objc.Sel("localEndpoint"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neappproxyudpflow/localflowendpoint-7ukb6
func (n_ NEAppProxyUDPFlow) LocalFlowEndpoint() INWEndpoint {
	rv := objc.Send[NWEndpoint](n_.ID, objc.Sel("localFlowEndpoint"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neappproxyudpflow/localflowendpoint-7ukb6
func (n_ NEAppProxyUDPFlow) SetLocalFlowEndpoint(value INWEndpoint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setLocalFlowEndpoint:"), value)
}



