// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEPacket */


/* debug [class_header]: Header for NEPacket */
// The class instance for the [NEPacket] class.
var (
	NEPacketClass     _NEPacketClass
	NEPacketClassOnce sync.Once
)

func getNEPacketClass() _NEPacketClass {
	NEPacketClassOnce.Do(func() {
		NEPacketClass = _NEPacketClass{objc.GetClass("NEPacket")}
	})
	return NEPacketClass
}

type _NEPacketClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEPacket */
// An interface definition for the [NEPacket] class.
type INEPacket interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NEPacket */
	// properties:
	Data() objc.IObject /* cross-framework: NSData */
	Direction() NETrafficDirection
	Metadata() INEFlowMetaData
	ProtocolFamily() objectivec.IObject
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEPacket */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEPacket */
// Alloc allocates a new instance without initialization.
func (nc _NEPacketClass) Alloc() NEPacket {
	rv := objc.Send[NEPacket](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEPacketClass) New() NEPacket {
	rv := objc.Send[NEPacket](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEPacket) Init() NEPacket {
	rv := objc.Send[NEPacket](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEPacket) Autorelease() NEPacket {
	rv := objc.Send[NEPacket](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEPacket creates a new NEPacket instance.
func NewNEPacket() NEPacket {
	return getNEPacketClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEPacket */
// A network packet and its associated properties.


// A network packet and its associated properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacket
type NEPacket struct {
	objectivec.Object
}

// NEPacketFrom constructs a [NEPacket] from an unsafe.Pointer.
//
// A network packet and its associated properties.
func NEPacketFrom(ptr unsafe.Pointer) NEPacket {
	return NEPacket{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEPacket */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacket/init(data:protocolFamily:)
func NewNEPacketWithDataProtocolFamily(data objc.IObject /* cross-framework: NSData */, protocolFamily objectivec.IObject) NEPacket {
	instance := getNEPacketClass().Alloc()
	rv := objc.Send[NEPacket](instance.ID, objc.Sel("initWithData:protocolFamily:"), data, protocolFamily)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNEPacketWithDataProtocolFamily */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEPacket */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEPacket */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEPacket */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEPacket */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacket/data
func (n_ NEPacket) Data() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](n_.ID, objc.Sel("data"))
	return rv
}/* debug [instance_properties/getter]: data */


// The direction of the packet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacket/direction
func (n_ NEPacket) Direction() NETrafficDirection {
	rv := objc.Send[NETrafficDirection](n_.ID, objc.Sel("direction"))
	return rv
}/* debug [instance_properties/getter]: direction */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacket/metadata
func (n_ NEPacket) Metadata() INEFlowMetaData {
	rv := objc.Send[NEFlowMetaData](n_.ID, objc.Sel("metadata"))
	return rv
}/* debug [instance_properties/getter]: metadata */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacket/protocolFamily
func (n_ NEPacket) ProtocolFamily() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("protocolFamily"))
	return rv
}/* debug [instance_properties/getter]: protocolFamily */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEPacket */


