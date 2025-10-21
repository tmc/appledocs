// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEAppProxyFlow] class.
var (
	NEAppProxyFlowClass     _NEAppProxyFlowClass
	NEAppProxyFlowClassOnce sync.Once
)

func getNEAppProxyFlowClass() _NEAppProxyFlowClass {
	NEAppProxyFlowClassOnce.Do(func() {
		NEAppProxyFlowClass = _NEAppProxyFlowClass{objc.GetClass("NEAppProxyFlow")}
	})
	return NEAppProxyFlowClass
}

type _NEAppProxyFlowClass struct {
	class objc.Class
}

// An interface definition for the [NEAppProxyFlow] class.
type INEAppProxyFlow interface {
	objectivec.IObject
	CloseReadWithError(error_ unsafe.Pointer)
	CloseWriteWithError(error_ unsafe.Pointer)
	OpenWithLocalEndpointCompletionHandler(localEndpoint unsafe.Pointer, completionHandler unsafe.Pointer)
	OpenWithLocalFlowEndpointCompletionHandler(localEndpoint unsafe.Pointer, completionHandler unsafe.Pointer)
	SetMetadata(parameters unsafe.Pointer)
}

// An abstract base class shared by NEAppProxyTCPFlow and NEAppProxyUDPFlow.
//
// App Proxy Providers receive network connections to be proxied in the form of objects, which are passed to the App Proxy Provider via the method. objects are initially in an unopened state. Before they can be used to transmit network data, they must be opened using the method. When you are finished with a flow, you should call and , and then release the object.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlow
type NEAppProxyFlow struct {
	objectivec.Object
}

// NEAppProxyFlowFrom constructs a [NEAppProxyFlow] from an unsafe.Pointer.
//
// An abstract base class shared by NEAppProxyTCPFlow and NEAppProxyUDPFlow.
func NEAppProxyFlowFrom(ptr unsafe.Pointer) NEAppProxyFlow {
	return NEAppProxyFlow{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEAppProxyFlowClass) Alloc() NEAppProxyFlow {
	rv := objc.Send[NEAppProxyFlow](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEAppProxyFlowClass) New() NEAppProxyFlow {
	rv := objc.Send[NEAppProxyFlow](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEAppProxyFlow) Init() NEAppProxyFlow {
	rv := objc.Send[NEAppProxyFlow](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEAppProxyFlow) Autorelease() NEAppProxyFlow {
	rv := objc.Send[NEAppProxyFlow](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEAppProxyFlow creates a new NEAppProxyFlow instance.
func NewNEAppProxyFlow() NEAppProxyFlow {
	return getNEAppProxyFlowClass().New()
}


// Close the flow for further read operations.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlow/closeReadWithError(_:)
func (n_ NEAppProxyFlow) CloseReadWithError(error_ unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("closeReadWithError:"), error_)
}

// Close the flow for further write operations.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlow/closeWriteWithError(_:)
func (n_ NEAppProxyFlow) CloseWriteWithError(error_ unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("closeWriteWithError:"), error_)
}

// Opens the flow, indicating to the system that the caller is ready to start receiving and sending data.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlow/open(withLocalEndpoint:completionHandler:)
func (n_ NEAppProxyFlow) OpenWithLocalEndpointCompletionHandler(localEndpoint unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("openWithLocalEndpoint:completionHandler:"), localEndpoint, completionHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlow/openWithLocalFlowEndpoint:completionHandler:
func (n_ NEAppProxyFlow) OpenWithLocalFlowEndpointCompletionHandler(localEndpoint unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("openWithLocalFlowEndpoint:completionHandler:"), localEndpoint, completionHandler)
}

// Sets the flow’s metadata for use by proxy providers.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlow/setMetadata(_:)
func (n_ NEAppProxyFlow) SetMetadata(parameters unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMetadata:"), parameters)
}

// A Boolean value that indicates whether the flow has a binding to a specific interface.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlow/isBound
func (n_ NEAppProxyFlow) IsBound() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isBound"))
	return rv
}

// A metadata object containing information about the source app of the flow.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlow/metaData
func (n_ NEAppProxyFlow) MetaData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("metaData"))
	return rv
}

// The network interface, if any, used by this flow.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlow/networkInterface
func (n_ NEAppProxyFlow) NetworkInterface() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("networkInterface"))
	return rv
}


// SetNetworkInterface sets the value of the networkInterface property.
// The network interface, if any, used by this flow.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlow/networkInterface
func (n_ NEAppProxyFlow) SetNetworkInterface(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNetworkInterface:"), value)
}

// The remote host name for flows created from a hostname.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlow/remoteHostname
func (n_ NEAppProxyFlow) RemoteHostname() string {
	rv := objc.Send[string](n_.ID, objc.Sel("remoteHostname"))
	return rv
}



