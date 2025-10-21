// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [NEAppProxyTCPFlow] class.
var (
	NEAppProxyTCPFlowClass     _NEAppProxyTCPFlowClass
	NEAppProxyTCPFlowClassOnce sync.Once
)

func getNEAppProxyTCPFlowClass() _NEAppProxyTCPFlowClass {
	NEAppProxyTCPFlowClassOnce.Do(func() {
		NEAppProxyTCPFlowClass = _NEAppProxyTCPFlowClass{objc.GetClass("NEAppProxyTCPFlow")}
	})
	return NEAppProxyTCPFlowClass
}

type _NEAppProxyTCPFlowClass struct {
	class objc.Class
}

// An interface definition for the [NEAppProxyTCPFlow] class.
type INEAppProxyTCPFlow interface {
	INEAppProxyFlow
	ReadDataWithCompletionHandler(completionHandler unsafe.Pointer)
	WriteDataWithCompletionHandler(data foundation.IData, completionHandler unsafe.Pointer)
}

// An object for reading and writing data to and from a TCP connection being proxied by the provider.
//
// App Proxy Providers receive TCP connections to be proxied in the form of objects.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyTCPFlow
type NEAppProxyTCPFlow struct {
	NEAppProxyFlow
}

// NEAppProxyTCPFlowFrom constructs a [NEAppProxyTCPFlow] from an unsafe.Pointer.
//
// An object for reading and writing data to and from a TCP connection being proxied by the provider.
func NEAppProxyTCPFlowFrom(ptr unsafe.Pointer) NEAppProxyTCPFlow {
	return NEAppProxyTCPFlow{
		NEAppProxyFlow: NEAppProxyFlowFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NEAppProxyTCPFlowClass) Alloc() NEAppProxyTCPFlow {
	rv := objc.Send[NEAppProxyTCPFlow](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEAppProxyTCPFlowClass) New() NEAppProxyTCPFlow {
	rv := objc.Send[NEAppProxyTCPFlow](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEAppProxyTCPFlow) Init() NEAppProxyTCPFlow {
	rv := objc.Send[NEAppProxyTCPFlow](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEAppProxyTCPFlow) Autorelease() NEAppProxyTCPFlow {
	rv := objc.Send[NEAppProxyTCPFlow](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEAppProxyTCPFlow creates a new NEAppProxyTCPFlow instance.
func NewNEAppProxyTCPFlow() NEAppProxyTCPFlow {
	return getNEAppProxyTCPFlowClass().New()
}


// Read data from the flow.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyTCPFlow/readData(completionHandler:)
func (n_ NEAppProxyTCPFlow) ReadDataWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("readDataWithCompletionHandler:"), completionHandler)
}

// Write data to the flow.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyTCPFlow/write(_:withCompletionHandler:)
func (n_ NEAppProxyTCPFlow) WriteDataWithCompletionHandler(data foundation.IData, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("writeData:withCompletionHandler:"), data, completionHandler)
}

// An object containing information about the intended remote endpoint of the flow.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyTCPFlow/remoteEndpoint
func (n_ NEAppProxyTCPFlow) RemoteEndpoint() NWEndpoint {
	rv := objc.Send[NWEndpoint](n_.ID, objc.Sel("remoteEndpoint"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyTCPFlow/remoteFlowEndpoint-9lvob
func (n_ NEAppProxyTCPFlow) RemoteFlowEndpoint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("remoteFlowEndpoint"))
	return rv
}



