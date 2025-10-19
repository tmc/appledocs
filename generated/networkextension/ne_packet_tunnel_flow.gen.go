// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEPacketTunnelFlow] class.
var nEPacketTunnelFlowClass = _NEPacketTunnelFlowClass{objc.GetClass("NEPacketTunnelFlow")}

type _NEPacketTunnelFlowClass struct {
	class objc.Class
}

// An interface definition for the [NEPacketTunnelFlow] class.
type INEPacketTunnelFlow interface {
	objectivec.IObject
}

// An object you use to read and write packets to and from the tunnel’s virtual interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelFlow

type NEPacketTunnelFlow struct {
	objectivec.Object
}

// NEPacketTunnelFlowFrom constructs a [NEPacketTunnelFlow] from an unsafe.Pointer.
//
// An object you use to read and write packets to and from the tunnel’s virtual interface.
func NEPacketTunnelFlowFrom(ptr unsafe.Pointer) NEPacketTunnelFlow {
	return NEPacketTunnelFlow{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (nc _NEPacketTunnelFlowClass) Alloc() NEPacketTunnelFlow {
	rv := objc.Send[NEPacketTunnelFlow](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (nc _NEPacketTunnelFlowClass) New() NEPacketTunnelFlow {
	rv := objc.Send[NEPacketTunnelFlow](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEPacketTunnelFlow) Init() NEPacketTunnelFlow {
	rv := objc.Send[NEPacketTunnelFlow](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEPacketTunnelFlow) Autorelease() NEPacketTunnelFlow {
	rv := objc.Send[NEPacketTunnelFlow](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEPacketTunnelFlow creates a new NEPacketTunnelFlow instance.
func NewNEPacketTunnelFlow() NEPacketTunnelFlow {
	return nEPacketTunnelFlowClass.New()
}




