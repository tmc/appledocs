// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	RemoteHostname() objc.IObject /* cross-framework: NSString */
	NEAppProxyErrorDomain() objc.IObject /* cross-framework: NSString */
	Interface() unsafe.Pointer
	SetInterface(value unsafe.Pointer)
	IsBound() bool
	SetIsBound(value bool)
	MetaData() objc.IObject /* cross-framework: NEFlowMetaData */
	SetMetaData(value objc.IObject /* cross-framework: NEFlowMetaData */)
	NetworkInterface() unsafe.Pointer
	SetNetworkInterface(value unsafe.Pointer)
	// methods:
}

// An abstract base class shared by NEAppProxyTCPFlow and NEAppProxyUDPFlow.
//
// App Proxy Providers receive network connections to be proxied in the form of objects, which are passed to the App Proxy Provider via the method. objects are initially in an unopened state. Before they can be used to transmit network data, they must be opened using the method. When you are finished with a flow, you should call and , and then release the object.


// An abstract base class shared by NEAppProxyTCPFlow and NEAppProxyUDPFlow.
//
// [Full Topic]
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



// The remote host name for flows created from a hostname.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlow/remoteHostname
func (n_ NEAppProxyFlow) RemoteHostname() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("remoteHostname"))
	return rv
}


// The domain used for app proxy errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neappproxyerrordomain
func (n_ NEAppProxyFlow) NEAppProxyErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("NEAppProxyErrorDomain"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neappproxyflow/interface
func (n_ NEAppProxyFlow) Interface() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("interface"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neappproxyflow/interface
func (n_ NEAppProxyFlow) SetInterface(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setInterface:"), value)
}


// A Boolean value that indicates whether the flow has a binding to a specific interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neappproxyflow/isbound
func (n_ NEAppProxyFlow) IsBound() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isBound"))
	return rv
}


// A Boolean value that indicates whether the flow has a binding to a specific interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neappproxyflow/isbound
func (n_ NEAppProxyFlow) SetIsBound(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsBound:"), value)
}


// A metadata object containing information about the source app of the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neappproxyflow/metadata
func (n_ NEAppProxyFlow) MetaData() objc.IObject /* cross-framework: NEFlowMetaData */ {
	rv := objc.Send[NEFlowMetaData](n_.ID, objc.Sel("metaData"))
	return rv
}


// A metadata object containing information about the source app of the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neappproxyflow/metadata
func (n_ NEAppProxyFlow) SetMetaData(value objc.IObject /* cross-framework: NEFlowMetaData */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMetaData:"), value)
}


// The network interface, if any, used by this flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neappproxyflow/networkinterface
func (n_ NEAppProxyFlow) NetworkInterface() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("networkInterface"))
	return rv
}


// The network interface, if any, used by this flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neappproxyflow/networkinterface
func (n_ NEAppProxyFlow) SetNetworkInterface(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNetworkInterface:"), value)
}



