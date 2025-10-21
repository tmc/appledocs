// Code generated from Apple documentation for NearbyInteraction. DO NOT EDIT.

package nearbyinteraction

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NINearbyObject] class.
var (
	NINearbyObjectClass     _NINearbyObjectClass
	NINearbyObjectClassOnce sync.Once
)

func getNINearbyObjectClass() _NINearbyObjectClass {
	NINearbyObjectClassOnce.Do(func() {
		NINearbyObjectClass = _NINearbyObjectClass{objc.GetClass("NINearbyObject")}
	})
	return NINearbyObjectClass
}

type _NINearbyObjectClass struct {
	class objc.Class
}

// An interface definition for the [NINearbyObject] class.
type ININearbyObject interface {
	objectivec.IObject
}

// Location information for a peer device in an interaction session.
//
// A nearby object refers to a peer Apple device or third-party accessory. When the framework is ready to provide your app with information about a nearby object’s relative position, it calls your delegate’s implementation. If a session can’t provide peer direction or distance, it sets the values to . In Objective-C, the session uses the and values to indicate missing direction or distance. For more information, see .
//
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyObject
type NINearbyObject struct {
	objectivec.Object
}

// NINearbyObjectFrom constructs a [NINearbyObject] from an unsafe.Pointer.
//
// Location information for a peer device in an interaction session.
func NINearbyObjectFrom(ptr unsafe.Pointer) NINearbyObject {
	return NINearbyObject{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NINearbyObjectClass) Alloc() NINearbyObject {
	rv := objc.Send[NINearbyObject](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NINearbyObjectClass) New() NINearbyObject {
	rv := objc.Send[NINearbyObject](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NINearbyObject) Init() NINearbyObject {
	rv := objc.Send[NINearbyObject](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NINearbyObject) Autorelease() NINearbyObject {
	rv := objc.Send[NINearbyObject](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNINearbyObject creates a new NINearbyObject instance.
func NewNINearbyObject() NINearbyObject {
	return getNINearbyObjectClass().New()
}


// A unique identifier for a peer device in the session.
//
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyObject/discoveryToken
func (n_ NINearbyObject) DiscoveryToken() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("discoveryToken"))
	return rv
}

// The estimation of a nearby object’s vertical position as it relates to the user’s device.
//
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyObject/verticalDirectionEstimate-swift.property
func (n_ NINearbyObject) VerticalDirectionEstimate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("verticalDirectionEstimate"))
	return rv
}

// A vector that points from the user’s device in the direction of the peer device.
//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/ninearbyobject/direction-4qh5w
func (n_ NINearbyObject) Direction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("direction"))
	return rv
}


// SetDirection sets the value of the direction property.
// A vector that points from the user’s device in the direction of the peer device.

//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/ninearbyobject/direction-4qh5w
func (n_ NINearbyObject) SetDirection(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDirection:"), value)
}

// The distance from the user’s device to the peer device in meters.
//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/ninearbyobject/distance-676dm
func (n_ NINearbyObject) Distance() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("distance"))
	return rv
}


// SetDistance sets the value of the distance property.
// The distance from the user’s device to the peer device in meters.

//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/ninearbyobject/distance-676dm
func (n_ NINearbyObject) SetDistance(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDistance:"), value)
}

// An angle in radians that indicates the azimuthal direction to the nearby object.
//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/ninearbyobject/horizontalangle-hsg
func (n_ NINearbyObject) HorizontalAngle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("horizontalAngle"))
	return rv
}


// SetHorizontalAngle sets the value of the horizontalAngle property.
// An angle in radians that indicates the azimuthal direction to the nearby object.

//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/ninearbyobject/horizontalangle-hsg
func (n_ NINearbyObject) SetHorizontalAngle(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHorizontalAngle:"), value)
}



