// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NEAppProxyUDPFlow] class.
var nEAppProxyUDPFlowClass = _NEAppProxyUDPFlowClass{objc.GetClass("NEAppProxyUDPFlow")}

type _NEAppProxyUDPFlowClass struct {
	class objc.Class
}

// An interface definition for the [NEAppProxyUDPFlow] class.
type INEAppProxyUDPFlow interface {
	INEAppProxyFlow
}

// An object for reading and writing data to and from a UDP conversation being proxied by the provider. [Full Topic]
//
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

// New creates and returns a new instance with a +1 retain count.
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
	return nEAppProxyUDPFlowClass.New()
}




