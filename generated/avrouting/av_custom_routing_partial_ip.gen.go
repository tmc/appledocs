// Code generated from Apple documentation for AVRouting. DO NOT EDIT.

package avrouting

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [CustomRoutingPartialIP] class.
type ICustomRoutingPartialIP interface {
	objectivec.IObject
}

// An object that represents a full or partial IP address.
//
// Use this type to define the IP address and subnet mask of known routes on a local network. Create an instance of this class and add it to a custom routing controller’s array like shown below:
//
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

// Alloc allocates a new instance without initialization.
func (cc _CustomRoutingPartialIPClass) Alloc() CustomRoutingPartialIP {
	rv := objc.Send[CustomRoutingPartialIP](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// A mask that represents how many octets of the IP address to respect.
//
// [Full Topic]: https://developer.apple.com/documentation/avrouting/avcustomroutingpartialip/mask
func (c_ CustomRoutingPartialIP) Mask() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("mask"))
	return rv
}


// SetMask sets the value of the mask property.
// A mask that represents how many octets of the IP address to respect.

//
// [Full Topic]: https://developer.apple.com/documentation/avrouting/avcustomroutingpartialip/mask
func (c_ CustomRoutingPartialIP) SetMask(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMask:"), value)
}

// An array of route addresses known to be on the local network.
//
// [Full Topic]: https://developer.apple.com/documentation/avrouting/avcustomroutingcontroller/knownrouteips
func (c_ CustomRoutingPartialIP) KnownRouteIPs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("knownRouteIPs"))
	return rv
}


// SetKnownRouteIPs sets the value of the knownRouteIPs property.
// An array of route addresses known to be on the local network.

//
// [Full Topic]: https://developer.apple.com/documentation/avrouting/avcustomroutingcontroller/knownrouteips
func (c_ CustomRoutingPartialIP) SetKnownRouteIPs(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKnownRouteIPs:"), value)
}

// A full or partial IP address for a device known to be on the network.
//
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomRoutingPartialIP/address
func (c_ CustomRoutingPartialIP) Address() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("address"))
	return rv
}



