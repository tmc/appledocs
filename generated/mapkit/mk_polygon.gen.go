// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKPolygon */


/* debug [class_header]: Header for MKPolygon */
// The class instance for the [MKPolygon] class.
var (
	MKPolygonClass     _MKPolygonClass
	MKPolygonClassOnce sync.Once
)

func getMKPolygonClass() _MKPolygonClass {
	MKPolygonClassOnce.Do(func() {
		MKPolygonClass = _MKPolygonClass{objc.GetClass("MKPolygon")}
	})
	return MKPolygonClass
}

type _MKPolygonClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKPolygon */
// An interface definition for the [MKPolygon] class.
type IMKPolygon interface {
	IMKMultiPoint
	
/* debug [class_interface_properties]: Properties for MKPolygon */
	// properties:
	InteriorPolygons() []MKPolygon
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKPolygon */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKPolygon */
// Alloc allocates a new instance without initialization.
func (mc _MKPolygonClass) Alloc() MKPolygon {
	rv := objc.Send[MKPolygon](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKPolygonClass) New() MKPolygon {
	rv := objc.Send[MKPolygon](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKPolygon) Init() MKPolygon {
	rv := objc.Send[MKPolygon](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKPolygon) Autorelease() MKPolygon {
	rv := objc.Send[MKPolygon](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKPolygon creates a new MKPolygon instance.
func NewMKPolygon() MKPolygon {
	return getMKPolygonClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKPolygon */
// A closed polygon overlay.
//
// The points you add to this overlay connect end-to-end in the order you provide them. The first and last points connect to each other to create a closed shape. When creating a polygon, you can mask out portions of the polygon by specifying one or more interior polygons. For the polygons you specify, this class uses the even-odd fill rule to determine the final occupied area. When applied to overlapping polygons, this rule can cause the framework to mask specific regions out and thereby remove them from the total occupied area. For more information about how fill rules apply to paths, see in .


// A closed polygon overlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolygon
type MKPolygon struct {
	MKMultiPoint
}

// MKPolygonFrom constructs a [MKPolygon] from an unsafe.Pointer.
//
// A closed polygon overlay.
func MKPolygonFrom(ptr unsafe.Pointer) MKPolygon {
	return MKPolygon{
		MKMultiPoint: MKMultiPointFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKPolygon */

// Creates and returns a polygon object from the specified set of coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolygon/init(coordinates:count:)
func NewMKPolygonWithCoordinatesCount(coords LocationCoordinate2D /* not a class type */, count uint) MKPolygon {
	rv := objc.Send[MKPolygon](objc.ID(getMKPolygonClass().class), objc.Sel("polygonWithCoordinates:count:"), coords, count)
	return rv
}/* debug [class_init_methods/constructor]: NewMKPolygonWithCoordinatesCount */


// Creates and returns a polygon object from the specified set of coordinates and interior polygons.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolygon/init(coordinates:count:interiorPolygons:)
func NewMKPolygonWithCoordinatesCountInteriorPolygons(coords LocationCoordinate2D /* not a class type */, count uint, interiorPolygons []MKPolygon) MKPolygon {
	rv := objc.Send[MKPolygon](objc.ID(getMKPolygonClass().class), objc.Sel("polygonWithCoordinates:count:interiorPolygons:"), coords, count, interiorPolygons)
	return rv
}/* debug [class_init_methods/constructor]: NewMKPolygonWithCoordinatesCountInteriorPolygons */


// Creates and returns a polygon object from the specified set of map points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolygon/init(points:count:)
func NewMKPolygonWithPointsCount(points objc.IObject /* cross-framework: MKMapPoint */, count uint) MKPolygon {
	rv := objc.Send[MKPolygon](objc.ID(getMKPolygonClass().class), objc.Sel("polygonWithPoints:count:"), points, count)
	return rv
}/* debug [class_init_methods/constructor]: NewMKPolygonWithPointsCount */


// Creates and returns a polygon object from the specified set of map points and interior polygons.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolygon/init(points:count:interiorPolygons:)
func NewMKPolygonWithPointsCountInteriorPolygons(points objc.IObject /* cross-framework: MKMapPoint */, count uint, interiorPolygons []MKPolygon) MKPolygon {
	rv := objc.Send[MKPolygon](objc.ID(getMKPolygonClass().class), objc.Sel("polygonWithPoints:count:interiorPolygons:"), points, count, interiorPolygons)
	return rv
}/* debug [class_init_methods/constructor]: NewMKPolygonWithPointsCountInteriorPolygons */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKPolygon */

// Creates and returns a polygon object from the specified set of coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolygon/init(coordinates:count:)
func (mc _MKPolygonClass) PolygonWithCoordinatesCount(coords LocationCoordinate2D /* not a class type */, count uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("polygonWithCoordinates:count:"), coords, count)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PolygonWithCoordinatesCount) */


// Creates and returns a polygon object from the specified set of coordinates and interior polygons.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolygon/init(coordinates:count:interiorPolygons:)
func (mc _MKPolygonClass) PolygonWithCoordinatesCountInteriorPolygons(coords LocationCoordinate2D /* not a class type */, count uint, interiorPolygons []MKPolygon) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("polygonWithCoordinates:count:interiorPolygons:"), coords, count, interiorPolygons)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PolygonWithCoordinatesCountInteriorPolygons) */


// Creates and returns a polygon object from the specified set of map points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolygon/init(points:count:)
func (mc _MKPolygonClass) PolygonWithPointsCount(points objc.IObject /* cross-framework: MKMapPoint */, count uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("polygonWithPoints:count:"), points, count)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PolygonWithPointsCount) */


// Creates and returns a polygon object from the specified set of map points and interior polygons.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolygon/init(points:count:interiorPolygons:)
func (mc _MKPolygonClass) PolygonWithPointsCountInteriorPolygons(points objc.IObject /* cross-framework: MKMapPoint */, count uint, interiorPolygons []MKPolygon) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("polygonWithPoints:count:interiorPolygons:"), points, count, interiorPolygons)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PolygonWithPointsCountInteriorPolygons) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKPolygon */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKPolygon */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKPolygon */

// The array of polygons that nest inside the enclosing polygon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolygon/interiorPolygons
func (m_ MKPolygon) InteriorPolygons() []MKPolygon {
	rv := objc.Send[[]MKPolygon](m_.ID, objc.Sel("interiorPolygons"))
	return rv
}/* debug [instance_properties/getter]: interiorPolygons */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKPolygon */


