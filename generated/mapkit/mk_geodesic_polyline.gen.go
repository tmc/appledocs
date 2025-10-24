// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKGeodesicPolyline */


/* debug [class_header]: Header for MKGeodesicPolyline */
// The class instance for the [MKGeodesicPolyline] class.
var (
	MKGeodesicPolylineClass     _MKGeodesicPolylineClass
	MKGeodesicPolylineClassOnce sync.Once
)

func getMKGeodesicPolylineClass() _MKGeodesicPolylineClass {
	MKGeodesicPolylineClassOnce.Do(func() {
		MKGeodesicPolylineClass = _MKGeodesicPolylineClass{objc.GetClass("MKGeodesicPolyline")}
	})
	return MKGeodesicPolylineClass
}

type _MKGeodesicPolylineClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKGeodesicPolyline */
// An interface definition for the [MKGeodesicPolyline] class.
type IMKGeodesicPolyline interface {
	IMKPolyline
	
/* debug [class_interface_properties]: Properties for MKGeodesicPolyline */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKGeodesicPolyline */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKGeodesicPolyline */
// Alloc allocates a new instance without initialization.
func (mc _MKGeodesicPolylineClass) Alloc() MKGeodesicPolyline {
	rv := objc.Send[MKGeodesicPolyline](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKGeodesicPolylineClass) New() MKGeodesicPolyline {
	rv := objc.Send[MKGeodesicPolyline](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKGeodesicPolyline) Init() MKGeodesicPolyline {
	rv := objc.Send[MKGeodesicPolyline](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKGeodesicPolyline) Autorelease() MKGeodesicPolyline {
	rv := objc.Send[MKGeodesicPolyline](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKGeodesicPolyline creates a new MKGeodesicPolyline instance.
func NewMKGeodesicPolyline() MKGeodesicPolyline {
	return getMKGeodesicPolylineClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKGeodesicPolyline */
// An open polygon overlay consisting of line segments that follow the contours of the Earth to create the shortest path between the specified points.
//
// A geodesic polyline contains a set of points that connect end-to-end in the order that you provide them. The first and last points don’t automatically connect to each other. When displaying on a two-dimensional map view, the line segment between any two points may appear curved.


// An open polygon overlay consisting of line segments that follow the contours of the Earth to create the shortest path between the specified points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKGeodesicPolyline
type MKGeodesicPolyline struct {
	MKPolyline
}

// MKGeodesicPolylineFrom constructs a [MKGeodesicPolyline] from an unsafe.Pointer.
//
// An open polygon overlay consisting of line segments that follow the contours of the Earth to create the shortest path between the specified points.
func MKGeodesicPolylineFrom(ptr unsafe.Pointer) MKGeodesicPolyline {
	return MKGeodesicPolyline{
		MKPolyline: MKPolylineFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKGeodesicPolyline */

// Creates and returns a geodesic polyline using the specified coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKGeodesicPolyline/init(coordinates:count:)
func NewMKGeodesicPolylineWithCoordinatesCount(coords LocationCoordinate2D /* not a class type */, count uint) MKGeodesicPolyline {
	rv := objc.Send[MKGeodesicPolyline](objc.ID(getMKGeodesicPolylineClass().class), objc.Sel("polylineWithCoordinates:count:"), coords, count)
	return rv
}/* debug [class_init_methods/constructor]: NewMKGeodesicPolylineWithCoordinatesCount */


// Creates and returns a geodesic polyline using the specified map points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKGeodesicPolyline/init(points:count:)
func NewMKGeodesicPolylineWithPointsCount(points objc.IObject /* cross-framework: MKMapPoint */, count uint) MKGeodesicPolyline {
	rv := objc.Send[MKGeodesicPolyline](objc.ID(getMKGeodesicPolylineClass().class), objc.Sel("polylineWithPoints:count:"), points, count)
	return rv
}/* debug [class_init_methods/constructor]: NewMKGeodesicPolylineWithPointsCount */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKGeodesicPolyline */

// Creates and returns a geodesic polyline using the specified coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKGeodesicPolyline/init(coordinates:count:)
func (mc _MKGeodesicPolylineClass) PolylineWithCoordinatesCount(coords LocationCoordinate2D /* not a class type */, count uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("polylineWithCoordinates:count:"), coords, count)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PolylineWithCoordinatesCount) */


// Creates and returns a geodesic polyline using the specified map points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKGeodesicPolyline/init(points:count:)
func (mc _MKGeodesicPolylineClass) PolylineWithPointsCount(points objc.IObject /* cross-framework: MKMapPoint */, count uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("polylineWithPoints:count:"), points, count)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PolylineWithPointsCount) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKGeodesicPolyline */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKGeodesicPolyline */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKGeodesicPolyline */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKGeodesicPolyline */


