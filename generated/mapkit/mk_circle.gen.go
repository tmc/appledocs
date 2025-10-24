// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKCircle */


/* debug [class_header]: Header for MKCircle */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKCircle */
// An interface definition for the [MKCircle] class.
type IMKCircle interface {
	IMKShape
	
/* debug [class_interface_properties]: Properties for MKCircle */
	// properties:
	BoundingMapRect() objc.IObject /* cross-framework: MKMapRect */
	Coordinate() LocationCoordinate2D /* not a class type */
	Radius() LocationDistance /* not a class type */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKCircle */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKCircle */
// Alloc allocates a new instance without initialization.
func (mc _MKCircleClass) Alloc() MKCircle {
	rv := objc.Send[MKCircle](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKCircle */
// A circular overlay with a configurable radius that you center on a geographic coordinate.
//
// This class defines the portion of the map that the overlay covers. To draw the region, return an object from the method of your map view delegate.


// A circular overlay with a configurable radius that you center on a geographic coordinate.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKCircle */

// Creates and returns a circle object using the specified coordinate and radius.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCircle/init(center:radius:)
func NewMKCircleWithCenterCoordinateRadius(coord LocationCoordinate2D /* not a class type */, radius LocationDistance /* not a class type */) MKCircle {
	rv := objc.Send[MKCircle](objc.ID(getMKCircleClass().class), objc.Sel("circleWithCenterCoordinate:radius:"), coord, radius)
	return rv
}/* debug [class_init_methods/constructor]: NewMKCircleWithCenterCoordinateRadius */


// Creates and returns a circle object that derives the circular area from the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCircle/init(mapRect:)
func NewMKCircleWithMapRect(mapRect objc.IObject /* cross-framework: MKMapRect */) MKCircle {
	rv := objc.Send[MKCircle](objc.ID(getMKCircleClass().class), objc.Sel("circleWithMapRect:"), mapRect)
	return rv
}/* debug [class_init_methods/constructor]: NewMKCircleWithMapRect */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKCircle */

// Creates and returns a circle object using the specified coordinate and radius.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCircle/init(center:radius:)
func (mc _MKCircleClass) CircleWithCenterCoordinateRadius(coord LocationCoordinate2D /* not a class type */, radius LocationDistance /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("circleWithCenterCoordinate:radius:"), coord, radius)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CircleWithCenterCoordinateRadius) */


// Creates and returns a circle object that derives the circular area from the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCircle/init(mapRect:)
func (mc _MKCircleClass) CircleWithMapRect(mapRect objc.IObject /* cross-framework: MKMapRect */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("circleWithMapRect:"), mapRect)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CircleWithMapRect) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKCircle */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKCircle */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKCircle */

// The bounding rectangle of the circular area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCircle/boundingMapRect
func (m_ MKCircle) BoundingMapRect() objc.IObject /* cross-framework: MKMapRect */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("boundingMapRect"))
	return rv
}/* debug [instance_properties/getter]: boundingMapRect */


// The center point of the circular area, specified as a latitude and longitude.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCircle/coordinate
func (m_ MKCircle) Coordinate() LocationCoordinate2D /* not a class type */ {
	rv := objc.Send[LocationCoordinate2D](m_.ID, objc.Sel("coordinate"))
	return rv
}/* debug [instance_properties/getter]: coordinate */


// The radius of the circular area, in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCircle/radius
func (m_ MKCircle) Radius() LocationDistance /* not a class type */ {
	rv := objc.Send[LocationDistance](m_.ID, objc.Sel("radius"))
	return rv
}/* debug [instance_properties/getter]: radius */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKCircle */


