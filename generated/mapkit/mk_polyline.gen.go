// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKPolyline */


/* debug [class_header]: Header for MKPolyline */
// The class instance for the [MKPolyline] class.
var (
	MKPolylineClass     _MKPolylineClass
	MKPolylineClassOnce sync.Once
)

func getMKPolylineClass() _MKPolylineClass {
	MKPolylineClassOnce.Do(func() {
		MKPolylineClass = _MKPolylineClass{objc.GetClass("MKPolyline")}
	})
	return MKPolylineClass
}

type _MKPolylineClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKPolyline */
// An interface definition for the [MKPolyline] class.
type IMKPolyline interface {
	IMKMultiPoint
	
/* debug [class_interface_properties]: Properties for MKPolyline */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKPolyline */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKPolyline */
// Alloc allocates a new instance without initialization.
func (mc _MKPolylineClass) Alloc() MKPolyline {
	rv := objc.Send[MKPolyline](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKPolylineClass) New() MKPolyline {
	rv := objc.Send[MKPolyline](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKPolyline) Init() MKPolyline {
	rv := objc.Send[MKPolyline](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKPolyline) Autorelease() MKPolyline {
	rv := objc.Send[MKPolyline](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKPolyline creates a new MKPolyline instance.
func NewMKPolyline() MKPolyline {
	return getMKPolylineClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKPolyline */
// An open polygon overlay consisting of one or more connected line segments.
//
// The points connect end-to-end in the order that you provide them. The first and last points don’t automatically connect to each other.


// An open polygon overlay consisting of one or more connected line segments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolyline
type MKPolyline struct {
	MKMultiPoint
}

// MKPolylineFrom constructs a [MKPolyline] from an unsafe.Pointer.
//
// An open polygon overlay consisting of one or more connected line segments.
func MKPolylineFrom(ptr unsafe.Pointer) MKPolyline {
	return MKPolyline{
		MKMultiPoint: MKMultiPointFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKPolyline */

// Creates a polyline object from the specified set of coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolyline/init(coordinates:count:)
func NewMKPolylineWithCoordinatesCount(coords LocationCoordinate2D /* not a class type */, count uint) MKPolyline {
	rv := objc.Send[MKPolyline](objc.ID(getMKPolylineClass().class), objc.Sel("polylineWithCoordinates:count:"), coords, count)
	return rv
}/* debug [class_init_methods/constructor]: NewMKPolylineWithCoordinatesCount */


// Creates a polyline object from the specified set of map points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolyline/init(points:count:)
func NewMKPolylineWithPointsCount(points objc.IObject /* cross-framework: MKMapPoint */, count uint) MKPolyline {
	rv := objc.Send[MKPolyline](objc.ID(getMKPolylineClass().class), objc.Sel("polylineWithPoints:count:"), points, count)
	return rv
}/* debug [class_init_methods/constructor]: NewMKPolylineWithPointsCount */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKPolyline */

// Creates a polyline object from the specified set of coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolyline/init(coordinates:count:)
func (mc _MKPolylineClass) PolylineWithCoordinatesCount(coords LocationCoordinate2D /* not a class type */, count uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("polylineWithCoordinates:count:"), coords, count)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PolylineWithCoordinatesCount) */


// Creates a polyline object from the specified set of map points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolyline/init(points:count:)
func (mc _MKPolylineClass) PolylineWithPointsCount(points objc.IObject /* cross-framework: MKMapPoint */, count uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("polylineWithPoints:count:"), points, count)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PolylineWithPointsCount) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKPolyline */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKPolyline */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKPolyline */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKPolyline */


