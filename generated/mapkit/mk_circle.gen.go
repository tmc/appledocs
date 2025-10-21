// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MKCircle] class.
var (
	MKCircleClass     _MKCircleClass
	MKCircleClassOnce sync.Once
)

func getMKCircleClass() _MKCircleClass {
	MKCircleClassOnce.Do(func() {
		MKCircleClass = _MKCircleClass{objc.GetClass("MKCircle")}
	})
	return MKCircleClass
}

type _MKCircleClass struct {
	class objc.Class
}

// An interface definition for the [MKCircle] class.
type IMKCircle interface {
	IMKShape
}

// A circular overlay with a configurable radius that you center on a geographic coordinate.
//
// This class defines the portion of the map that the overlay covers. To draw the region, return an object from the method of your map view delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCircle
type MKCircle struct {
	MKShape
}

// MKCircleFrom constructs a [MKCircle] from an unsafe.Pointer.
//
// A circular overlay with a configurable radius that you center on a geographic coordinate.
func MKCircleFrom(ptr unsafe.Pointer) MKCircle {
	return MKCircle{
		MKShape: MKShapeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MKCircleClass) Alloc() MKCircle {
	rv := objc.Send[MKCircle](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKCircleClass) New() MKCircle {
	rv := objc.Send[MKCircle](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKCircle) Init() MKCircle {
	rv := objc.Send[MKCircle](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKCircle) Autorelease() MKCircle {
	rv := objc.Send[MKCircle](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKCircle creates a new MKCircle instance.
func NewMKCircle() MKCircle {
	return getMKCircleClass().New()
}




// Creates and returns a circle object using the specified coordinate and radius.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCircle/init(center:radius:)
func NewMKCircleWithCenterCoordinateRadius(coord unsafe.Pointer, radius unsafe.Pointer) MKCircle {
	rv := objc.Send[MKCircle](objc.ID(getMKCircleClass().class), objc.Sel("circleWithCenterCoordinate:radius:"), coord, radius)
	return rv
}



// Creates and returns a circle object that derives the circular area from the specified rectangle.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCircle/init(mapRect:)
func NewMKCircleWithMapRect(mapRect unsafe.Pointer) MKCircle {
	rv := objc.Send[MKCircle](objc.ID(getMKCircleClass().class), objc.Sel("circleWithMapRect:"), mapRect)
	return rv
}


// Creates and returns a circle object using the specified coordinate and radius.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCircle/init(center:radius:)
func (mc _MKCircleClass) CircleWithCenterCoordinateRadius(coord unsafe.Pointer, radius unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("circleWithCenterCoordinate:radius:"), coord, radius)
	return rv
}

// Creates and returns a circle object that derives the circular area from the specified rectangle.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCircle/init(mapRect:)
func (mc _MKCircleClass) CircleWithMapRect(mapRect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("circleWithMapRect:"), mapRect)
	return rv
}

// The bounding rectangle of the circular area.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCircle/boundingMapRect
func (m_ MKCircle) BoundingMapRect() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("boundingMapRect"))
	return rv
}

// The center point of the circular area, specified as a latitude and longitude.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCircle/coordinate
func (m_ MKCircle) Coordinate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("coordinate"))
	return rv
}

// The radius of the circular area, in meters.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCircle/radius
func (m_ MKCircle) Radius() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("radius"))
	return rv
}


