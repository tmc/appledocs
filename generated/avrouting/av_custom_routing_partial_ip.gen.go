// Code generated from Apple documentation for AVRouting. DO NOT EDIT.

package avrouting

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCustomRoutingPartialIP */


/* debug [class_header]: Header for AVCustomRoutingPartialIP */
// The class instance for the [CustomRoutingPartialIP] class.
var (
	CustomRoutingPartialIPClass     _CustomRoutingPartialIPClass
	CustomRoutingPartialIPClassOnce sync.Once
)

func getCustomRoutingPartialIPClass() _CustomRoutingPartialIPClass {
	CustomRoutingPartialIPClassOnce.Do(func() {
		CustomRoutingPartialIPClass = _CustomRoutingPartialIPClass{objc.GetClass("AVCustomRoutingPartialIP")}
	})
	return CustomRoutingPartialIPClass
}

type _CustomRoutingPartialIPClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CustomRoutingPartialIP */
// An interface definition for the [CustomRoutingPartialIP] class.
type ICustomRoutingPartialIP interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CustomRoutingPartialIP */
	// properties:
	KnownRouteIPs() IAVCustomRoutingPartialIP
	SetKnownRouteIPs(value IAVCustomRoutingPartialIP)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CustomRoutingPartialIP */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CustomRoutingPartialIP */
// Alloc allocates a new instance without initialization.
func (cc _CustomRoutingPartialIPClass) Alloc() CustomRoutingPartialIP {
	rv := objc.Send[CustomRoutingPartialIP](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CustomRoutingPartialIPClass) New() CustomRoutingPartialIP {
	rv := objc.Send[CustomRoutingPartialIP](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CustomRoutingPartialIP) Init() CustomRoutingPartialIP {
	rv := objc.Send[CustomRoutingPartialIP](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CustomRoutingPartialIP) Autorelease() CustomRoutingPartialIP {
	rv := objc.Send[CustomRoutingPartialIP](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCustomRoutingPartialIP creates a new CustomRoutingPartialIP instance.
func NewCustomRoutingPartialIP() CustomRoutingPartialIP {
	return getCustomRoutingPartialIPClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CustomRoutingPartialIP */
// An object that represents a full or partial IP address.
//
// Use this type to define the IP address and subnet mask of known routes on a local network. Create an instance of this class and add it to a custom routing controller’s array like shown below:


// An object that represents a full or partial IP address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomRoutingPartialIP
type CustomRoutingPartialIP struct {
	objectivec.Object
}

// CustomRoutingPartialIPFrom constructs a [CustomRoutingPartialIP] from an unsafe.Pointer.
//
// An object that represents a full or partial IP address.
func CustomRoutingPartialIPFrom(ptr unsafe.Pointer) CustomRoutingPartialIP {
	return CustomRoutingPartialIP{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CustomRoutingPartialIP */

// Creates an IP fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomRoutingPartialIP/init(address:mask:)
func NewCustomRoutingPartialIPWithAddressMask(address objc.IObject /* cross-framework: NSData */, mask objc.IObject /* cross-framework: NSData */) CustomRoutingPartialIP {
	instance := getCustomRoutingPartialIPClass().Alloc()
	rv := objc.Send[CustomRoutingPartialIP](instance.ID, objc.Sel("initWithAddress:mask:"), address, mask)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCustomRoutingPartialIPWithAddressMask */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CustomRoutingPartialIP */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CustomRoutingPartialIP */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CustomRoutingPartialIP */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CustomRoutingPartialIP */

// An array of route addresses known to be on the local network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avrouting/avcustomroutingcontroller/knownrouteips
func (c_ CustomRoutingPartialIP) KnownRouteIPs() IAVCustomRoutingPartialIP {
	rv := objc.Send[CustomRoutingPartialIP](c_.ID, objc.Sel("knownRouteIPs"))
	return rv
}/* debug [instance_properties/getter]: knownRouteIPs */


// An array of route addresses known to be on the local network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avrouting/avcustomroutingcontroller/knownrouteips
func (c_ CustomRoutingPartialIP) SetKnownRouteIPs(value IAVCustomRoutingPartialIP) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKnownRouteIPs:"), value)
}/* debug [instance_properties/setter]: knownRouteIPs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCustomRoutingPartialIP */


