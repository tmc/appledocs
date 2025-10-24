// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEPacketTunnelFlow */


/* debug [class_header]: Header for NEPacketTunnelFlow */
// The class instance for the [NEPacketTunnelFlow] class.
var (
	NEPacketTunnelFlowClass     _NEPacketTunnelFlowClass
	NEPacketTunnelFlowClassOnce sync.Once
)

func getNEPacketTunnelFlowClass() _NEPacketTunnelFlowClass {
	NEPacketTunnelFlowClassOnce.Do(func() {
		NEPacketTunnelFlowClass = _NEPacketTunnelFlowClass{objc.GetClass("NEPacketTunnelFlow")}
	})
	return NEPacketTunnelFlowClass
}

type _NEPacketTunnelFlowClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEPacketTunnelFlow */
// An interface definition for the [NEPacketTunnelFlow] class.
type INEPacketTunnelFlow interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NEPacketTunnelFlow */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEPacketTunnelFlow */
	// methods:
	ReadPacketObjectsWithCompletionHandler(completionHandler unsafe.Pointer)
	ReadPacketsWithCompletionHandler(completionHandler unsafe.Pointer)
	WritePacketObjects(packets []NEPacket) bool
	WritePacketsWithProtocols(packets []foundation.Data, protocols []foundation.Number) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEPacketTunnelFlow */
// Alloc allocates a new instance without initialization.
func (nc _NEPacketTunnelFlowClass) Alloc() NEPacketTunnelFlow {
	rv := objc.Send[NEPacketTunnelFlow](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
	return getNEPacketTunnelFlowClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEPacketTunnelFlow */
// An object you use to read and write packets to and from the tunnel’s virtual interface.
//
// Use the class to implement a custom-IP tunneling protocol for your packet tunnel. For example, use the APIs in this class to read packets from the virtual interface, so you can then encapsulate these packets and send them to a packet-tunnel server. Likewise, read packets from your packet-tunnel server and use these APIs to write the packets back to the tunnel’s virtual interface.


// An object you use to read and write packets to and from the tunnel’s virtual interface.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEPacketTunnelFlow *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEPacketTunnelFlow */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEPacketTunnelFlow */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEPacketTunnelFlow */

// Read multiple IP packets from the TUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelFlow/readPacketObjects(completionHandler:)
func (n_ NEPacketTunnelFlow) ReadPacketObjectsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("readPacketObjectsWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: ReadPacketObjectsWithCompletionHandler */


// Reads IP packets from the TUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelFlow/readPackets(completionHandler:)
func (n_ NEPacketTunnelFlow) ReadPacketsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("readPacketsWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: ReadPacketsWithCompletionHandler */


// Write multiple IP packets to the TUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelFlow/writePacketObjects(_:)
func (n_ NEPacketTunnelFlow) WritePacketObjects(packets []NEPacket) bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("writePacketObjects:"), packets)
	return rv
}/* debug [instance_methods/method]: WritePacketObjects */


// Writes IP packets to the TUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelFlow/writePackets(_:withProtocols:)
func (n_ NEPacketTunnelFlow) WritePacketsWithProtocols(packets []foundation.Data, protocols []foundation.Number) bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("writePackets:withProtocols:"), packets, protocols)
	return rv
}/* debug [instance_methods/method]: WritePacketsWithProtocols */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEPacketTunnelFlow */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEPacketTunnelFlow */



