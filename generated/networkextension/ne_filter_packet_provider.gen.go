// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEFilterPacketProvider */


/* debug [class_header]: Header for NEFilterPacketProvider */
// The class instance for the [NEFilterPacketProvider] class.
var (
	NEFilterPacketProviderClass     _NEFilterPacketProviderClass
	NEFilterPacketProviderClassOnce sync.Once
)

func getNEFilterPacketProviderClass() _NEFilterPacketProviderClass {
	NEFilterPacketProviderClassOnce.Do(func() {
		NEFilterPacketProviderClass = _NEFilterPacketProviderClass{objc.GetClass("NEFilterPacketProvider")}
	})
	return NEFilterPacketProviderClass
}

type _NEFilterPacketProviderClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEFilterPacketProvider */
// An interface definition for the [NEFilterPacketProvider] class.
type INEFilterPacketProvider interface {
	INEFilterProvider
	
/* debug [class_interface_properties]: Properties for NEFilterPacketProvider */
	// properties:
	PacketHandler() objectivec.IObject
	SetPacketHandler(value objectivec.IObject)
	Handler() objectivec.IObject
	SetHandler(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEFilterPacketProvider */
	// methods:
	AllowPacket(packet INEPacket)
	DelayCurrentPacket(context INEFilterPacketContext) INEPacket
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEFilterPacketProvider */
// Alloc allocates a new instance without initialization.
func (nc _NEFilterPacketProviderClass) Alloc() NEFilterPacketProvider {
	rv := objc.Send[NEFilterPacketProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEFilterPacketProviderClass) New() NEFilterPacketProvider {
	rv := objc.Send[NEFilterPacketProvider](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEFilterPacketProvider) Init() NEFilterPacketProvider {
	rv := objc.Send[NEFilterPacketProvider](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEFilterPacketProvider) Autorelease() NEFilterPacketProvider {
	rv := objc.Send[NEFilterPacketProvider](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterPacketProvider creates a new NEFilterPacketProvider instance.
func NewNEFilterPacketProvider() NEFilterPacketProvider {
	return getNEFilterPacketProviderClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEFilterPacketProvider */
// A filter provider that evaluates network packets and decides whether to block, allow, or delay the packets.


// A filter provider that evaluates network packets and decides whether to block, allow, or delay the packets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterPacketProvider
type NEFilterPacketProvider struct {
	NEFilterProvider
}

// NEFilterPacketProviderFrom constructs a [NEFilterPacketProvider] from an unsafe.Pointer.
//
// A filter provider that evaluates network packets and decides whether to block, allow, or delay the packets.
func NEFilterPacketProviderFrom(ptr unsafe.Pointer) NEFilterPacketProvider {
	return NEFilterPacketProvider{
		NEFilterProvider: NEFilterProviderFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEFilterPacketProvider *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEFilterPacketProvider */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEFilterPacketProvider */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEFilterPacketProvider */

// Allow delivery of a previously-delayed packet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterPacketProvider/allow(_:)
func (n_ NEFilterPacketProvider) AllowPacket(packet INEPacket) {
	objc.Send[objc.ID](n_.ID, objc.Sel("allowPacket:"), packet)
}/* debug [instance_methods/method]: AllowPacket */


// Delay a packet currently processed by a packet handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterPacketProvider/delayCurrentPacket(_:)
func (n_ NEFilterPacketProvider) DelayCurrentPacket(context INEFilterPacketContext) INEPacket {
	rv := objc.Send[NEPacket](n_.ID, objc.Sel("delayCurrentPacket:"), context)
	return rv
}/* debug [instance_methods/method]: DelayCurrentPacket */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEFilterPacketProvider */

// A Swift closure or an ObjectiveC block that handles each packet received by the filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterPacketProvider/packetHandler
func (n_ NEFilterPacketProvider) PacketHandler() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("packetHandler"))
	return rv
}/* debug [instance_properties/getter]: packetHandler */


// A Swift closure or an ObjectiveC block that handles each packet received by the filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterPacketProvider/packetHandler
func (n_ NEFilterPacketProvider) SetPacketHandler(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPacketHandler:"), value)
}/* debug [instance_properties/setter]: packetHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterpacketprovider/handler
func (n_ NEFilterPacketProvider) Handler() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("handler"))
	return rv
}/* debug [instance_properties/getter]: handler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterpacketprovider/handler
func (n_ NEFilterPacketProvider) SetHandler(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHandler:"), value)
}/* debug [instance_properties/setter]: handler */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEFilterPacketProvider */



