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
	// properties:
	DiscoveryToken() objc.IObject /* cross-framework: NIDiscoveryToken */
	SetDiscoveryToken(value objc.IObject /* cross-framework: NIDiscoveryToken */)
	HorizontalAngle() float32
	SetHorizontalAngle(value float32)
	// methods:
}

// Location information for a peer device in an interaction session.
//
// A nearby object refers to a peer Apple device or third-party accessory. When the framework is ready to provide your app with information about a nearby object’s relative position, it calls your delegate’s implementation. If a session can’t provide peer direction or distance, it sets the values to . In Objective-C, the session uses the and values to indicate missing direction or distance. For more information, see .


// Location information for a peer device in an interaction session.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/ninearbyobject/discoverytoken
func (n_ NINearbyObject) DiscoveryToken() objc.IObject /* cross-framework: NIDiscoveryToken */ {
	rv := objc.Send[NIDiscoveryToken](n_.ID, objc.Sel("discoveryToken"))
	return rv
}


// A unique identifier for a peer device in the session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/ninearbyobject/discoverytoken
func (n_ NINearbyObject) SetDiscoveryToken(value objc.IObject /* cross-framework: NIDiscoveryToken */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDiscoveryToken:"), value)
}


// An angle in radians that indicates the azimuthal direction to the nearby object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/ninearbyobject/horizontalangle-hsg
func (n_ NINearbyObject) HorizontalAngle() float32 {
	rv := objc.Send[float32](n_.ID, objc.Sel("horizontalAngle"))
	return rv
}


// An angle in radians that indicates the azimuthal direction to the nearby object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/ninearbyobject/horizontalangle-hsg
func (n_ NINearbyObject) SetHorizontalAngle(value float32) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHorizontalAngle:"), value)
}


