// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKMultiPolyline */


/* debug [class_header]: Header for MKMultiPolyline */
// The class instance for the [MKMultiPolyline] class.
var (
	MKMultiPolylineClass     _MKMultiPolylineClass
	MKMultiPolylineClassOnce sync.Once
)

func getMKMultiPolylineClass() _MKMultiPolylineClass {
	MKMultiPolylineClassOnce.Do(func() {
		MKMultiPolylineClass = _MKMultiPolylineClass{objc.GetClass("MKMultiPolyline")}
	})
	return MKMultiPolylineClass
}

type _MKMultiPolylineClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKMultiPolyline */
// An interface definition for the [MKMultiPolyline] class.
type IMKMultiPolyline interface {
	IMKShape
	
/* debug [class_interface_properties]: Properties for MKMultiPolyline */
	// properties:
	Polylines() []MKPolyline
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKMultiPolyline */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKMultiPolyline */
// Alloc allocates a new instance without initialization.
func (mc _MKMultiPolylineClass) Alloc() MKMultiPolyline {
	rv := objc.Send[MKMultiPolyline](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKMultiPolylineClass) New() MKMultiPolyline {
	rv := objc.Send[MKMultiPolyline](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKMultiPolyline) Init() MKMultiPolyline {
	rv := objc.Send[MKMultiPolyline](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKMultiPolyline) Autorelease() MKMultiPolyline {
	rv := objc.Send[MKMultiPolyline](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKMultiPolyline creates a new MKMultiPolyline instance.
func NewMKMultiPolyline() MKMultiPolyline {
	return getMKMultiPolylineClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKMultiPolyline */
// A collection of multipolyline shapes, each consisting of one or more connected line segments.
//
// Use a object when you have multiple distinct polyline shapes that you intend to render using the same style.


// A collection of multipolyline shapes, each consisting of one or more connected line segments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMultiPolyline
type MKMultiPolyline struct {
	MKShape
}

// MKMultiPolylineFrom constructs a [MKMultiPolyline] from an unsafe.Pointer.
//
// A collection of multipolyline shapes, each consisting of one or more connected line segments.
func MKMultiPolylineFrom(ptr unsafe.Pointer) MKMultiPolyline {
	return MKMultiPolyline{
		MKShape: MKShapeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKMultiPolyline */

// Creates a multipolyline object using the provided polylines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMultiPolyline/init(_:)
func NewMKMultiPolylineWithPolylines(polylines []MKPolyline) MKMultiPolyline {
	instance := getMKMultiPolylineClass().Alloc()
	rv := objc.Send[MKMultiPolyline](instance.ID, objc.Sel("initWithPolylines:"), polylines)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKMultiPolylineWithPolylines */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKMultiPolyline */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKMultiPolyline */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKMultiPolyline */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKMultiPolyline */

// An array containing the polyline objects that make up the multipolyline object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMultiPolyline/polylines
func (m_ MKMultiPolyline) Polylines() []MKPolyline {
	rv := objc.Send[[]MKPolyline](m_.ID, objc.Sel("polylines"))
	return rv
}/* debug [instance_properties/getter]: polylines */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKMultiPolyline */


