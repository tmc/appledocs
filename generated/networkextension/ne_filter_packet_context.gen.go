// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEFilterPacketContext */


/* debug [class_header]: Header for NEFilterPacketContext */
// The class instance for the [NEFilterPacketContext] class.
var (
	NEFilterPacketContextClass     _NEFilterPacketContextClass
	NEFilterPacketContextClassOnce sync.Once
)

func getNEFilterPacketContextClass() _NEFilterPacketContextClass {
	NEFilterPacketContextClassOnce.Do(func() {
		NEFilterPacketContextClass = _NEFilterPacketContextClass{objc.GetClass("NEFilterPacketContext")}
	})
	return NEFilterPacketContextClass
}

type _NEFilterPacketContextClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEFilterPacketContext */
// An interface definition for the [NEFilterPacketContext] class.
type INEFilterPacketContext interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NEFilterPacketContext */
	// properties:
	PacketHandler() objectivec.IObject
	SetPacketHandler(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEFilterPacketContext */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEFilterPacketContext */
// Alloc allocates a new instance without initialization.
func (nc _NEFilterPacketContextClass) Alloc() NEFilterPacketContext {
	rv := objc.Send[NEFilterPacketContext](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEFilterPacketContextClass) New() NEFilterPacketContext {
	rv := objc.Send[NEFilterPacketContext](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEFilterPacketContext) Init() NEFilterPacketContext {
	rv := objc.Send[NEFilterPacketContext](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEFilterPacketContext) Autorelease() NEFilterPacketContext {
	rv := objc.Send[NEFilterPacketContext](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterPacketContext creates a new NEFilterPacketContext instance.
func NewNEFilterPacketContext() NEFilterPacketContext {
	return getNEFilterPacketContextClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEFilterPacketContext */
// The context object provided to the filter packet handler.


// The context object provided to the filter packet handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterPacketContext
type NEFilterPacketContext struct {
	objectivec.Object
}

// NEFilterPacketContextFrom constructs a [NEFilterPacketContext] from an unsafe.Pointer.
//
// The context object provided to the filter packet handler.
func NEFilterPacketContextFrom(ptr unsafe.Pointer) NEFilterPacketContext {
	return NEFilterPacketContext{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEFilterPacketContext *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEFilterPacketContext */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEFilterPacketContext */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEFilterPacketContext */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEFilterPacketContext */

// A Swift closure or an ObjectiveC block that handles each packet received by the filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterpacketprovider/packethandler
func (n_ NEFilterPacketContext) PacketHandler() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("packetHandler"))
	return rv
}/* debug [instance_properties/getter]: packetHandler */


// A Swift closure or an ObjectiveC block that handles each packet received by the filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterpacketprovider/packethandler
func (n_ NEFilterPacketContext) SetPacketHandler(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPacketHandler:"), value)
}/* debug [instance_properties/setter]: packetHandler */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEFilterPacketContext */



