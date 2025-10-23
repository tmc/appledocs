// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CircularRegion] class.
var (
	CircularRegionClass     _CircularRegionClass
	CircularRegionClassOnce sync.Once
)

func getCircularRegionClass() _CircularRegionClass {
	CircularRegionClassOnce.Do(func() {
		CircularRegionClass = _CircularRegionClass{objc.GetClass("CLCircularRegion")}
	})
	return CircularRegionClass
}

type _CircularRegionClass struct {
	class objc.Class
}

// An interface definition for the [CircularRegion] class.
type ICircularRegion interface {
	IRegion
	// properties:
	Radius() LocationDistance /* not a class type */
	Center() LocationCoordinate2D /* not a class type */
	SetCenter(value LocationCoordinate2D /* not a class type */)
	// methods:
}

// A circular geographic region that a center point and radius deine.
//
// The class defines the location and boundaries for a circular geographic region. You can use instances of this class to define geofences for a specific location. The crossing of a geofence’s boundary causes the location manager to notify its delegate.


// A circular geographic region that a center point and radius deine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLCircularRegion
type CircularRegion struct {
	Region
}

// CircularRegionFrom constructs a [CircularRegion] from an unsafe.Pointer.
//
// A circular geographic region that a center point and radius deine.
func CircularRegionFrom(ptr unsafe.Pointer) CircularRegion {
	return CircularRegion{
		Region: RegionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CircularRegionClass) Alloc() CircularRegion {
	rv := objc.Send[CircularRegion](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CircularRegionClass) New() CircularRegion {
	rv := objc.Send[CircularRegion](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CircularRegion) Init() CircularRegion {
	rv := objc.Send[CircularRegion](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CircularRegion) Autorelease() CircularRegion {
	rv := objc.Send[CircularRegion](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCircularRegion creates a new CircularRegion instance.
func NewCircularRegion() CircularRegion {
	return getCircularRegionClass().New()
}



// The radius (measured in meters) that defines the geographic area’s outer boundary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLCircularRegion/radius
func (c_ CircularRegion) Radius() LocationDistance /* not a class type */ {
	rv := objc.Send[LocationDistance](c_.ID, objc.Sel("radius"))
	return rv
}


// The center point of the geographic area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clcircularregion/center
func (c_ CircularRegion) Center() LocationCoordinate2D /* not a class type */ {
	rv := objc.Send[LocationCoordinate2D](c_.ID, objc.Sel("center"))
	return rv
}


// The center point of the geographic area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clcircularregion/center
func (c_ CircularRegion) SetCenter(value LocationCoordinate2D /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCenter:"), value)
}



