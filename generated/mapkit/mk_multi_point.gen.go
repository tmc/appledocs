// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
)

/* debug [class.gen.go]: Generating class MKMultiPoint */


/* debug [class_header]: Header for MKMultiPoint */
// The class instance for the [MKMultiPoint] class.
var (
	MKMultiPointClass     _MKMultiPointClass
	MKMultiPointClassOnce sync.Once
)

func getMKMultiPointClass() _MKMultiPointClass {
	MKMultiPointClassOnce.Do(func() {
		MKMultiPointClass = _MKMultiPointClass{objc.GetClass("MKMultiPoint")}
	})
	return MKMultiPointClass
}

type _MKMultiPointClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKMultiPoint */
// An interface definition for the [MKMultiPoint] class.
type IMKMultiPoint interface {
	IMKShape
	
/* debug [class_interface_properties]: Properties for MKMultiPoint */
	// properties:
	PointCount() uint
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKMultiPoint */
	// methods:
	GetCoordinatesRange(coords LocationCoordinate2D /* not a class type */, range_ corefoundation.Range)
	LocationAtPointIndex(index uint) float64
	LocationsAtPointIndexes(indexes foundation.IndexSet) []foundation.Number
	Points() objc.IObject /* cross-framework: MKMapPoint */
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKMultiPoint */
// Alloc allocates a new instance without initialization.
func (mc _MKMultiPointClass) Alloc() MKMultiPoint {
	rv := objc.Send[MKMultiPoint](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKMultiPointClass) New() MKMultiPoint {
	rv := objc.Send[MKMultiPoint](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKMultiPoint) Init() MKMultiPoint {
	rv := objc.Send[MKMultiPoint](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKMultiPoint) Autorelease() MKMultiPoint {
	rv := objc.Send[MKMultiPoint](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKMultiPoint creates a new MKMultiPoint instance.
func NewMKMultiPoint() MKMultiPoint {
	return getMKMultiPointClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKMultiPoint */
// An abstract class that defines the common behavior that open and closed polygon overlays share.
//
// Don’t create instances of this class directly. Instead, create instances of the or classes. However, you can use the methods and property of this class to access information about the specific points associated with the line or polygon.


// An abstract class that defines the common behavior that open and closed polygon overlays share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMultiPoint
type MKMultiPoint struct {
	MKShape
}

// MKMultiPointFrom constructs a [MKMultiPoint] from an unsafe.Pointer.
//
// An abstract class that defines the common behavior that open and closed polygon overlays share.
func MKMultiPointFrom(ptr unsafe.Pointer) MKMultiPoint {
	return MKMultiPoint{
		MKShape: MKShapeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKMultiPoint *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKMultiPoint */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKMultiPoint */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKMultiPoint */

// Retrieves one or more points associated with the shape and converts them to coordinate values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMultiPoint/getCoordinates(_:range:)
func (m_ MKMultiPoint) GetCoordinatesRange(coords LocationCoordinate2D /* not a class type */, range_ corefoundation.Range) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getCoordinates:range:"), coords, range_)
}/* debug [instance_methods/method]: GetCoordinatesRange */


// Translates a point index into a unit distance along the shape.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMultiPoint/location(atPointIndex:)
func (m_ MKMultiPoint) LocationAtPointIndex(index uint) float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("locationAtPointIndex:"), index)
	return rv
}/* debug [instance_methods/method]: LocationAtPointIndex */


// Returns a set of unit distance values that correspond to the point indexes along the shape.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMultiPoint/locationsAtPointIndexes:
func (m_ MKMultiPoint) LocationsAtPointIndexes(indexes foundation.IndexSet) []foundation.Number {
	rv := objc.Send[[]foundation.Number](m_.ID, objc.Sel("locationsAtPointIndexes:"), indexes)
	return rv
}/* debug [instance_methods/method]: LocationsAtPointIndexes */


// Returns an array of map points associated with the shape.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMultiPoint/points()
func (m_ MKMultiPoint) Points() objc.IObject /* cross-framework: MKMapPoint */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("points"))
	return rv
}/* debug [instance_methods/method]: Points */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKMultiPoint */

// The number of points associated with the shape.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMultiPoint/pointCount
func (m_ MKMultiPoint) PointCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("pointCount"))
	return rv
}/* debug [instance_properties/getter]: pointCount */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKMultiPoint */



