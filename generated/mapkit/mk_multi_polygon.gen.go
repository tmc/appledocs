// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKMultiPolygon */


/* debug [class_header]: Header for MKMultiPolygon */
// The class instance for the [MKMultiPolygon] class.
var (
	MKMultiPolygonClass     _MKMultiPolygonClass
	MKMultiPolygonClassOnce sync.Once
)

func getMKMultiPolygonClass() _MKMultiPolygonClass {
	MKMultiPolygonClassOnce.Do(func() {
		MKMultiPolygonClass = _MKMultiPolygonClass{objc.GetClass("MKMultiPolygon")}
	})
	return MKMultiPolygonClass
}

type _MKMultiPolygonClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKMultiPolygon */
// An interface definition for the [MKMultiPolygon] class.
type IMKMultiPolygon interface {
	IMKShape
	
/* debug [class_interface_properties]: Properties for MKMultiPolygon */
	// properties:
	Polygons() []MKPolygon
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKMultiPolygon */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKMultiPolygon */
// Alloc allocates a new instance without initialization.
func (mc _MKMultiPolygonClass) Alloc() MKMultiPolygon {
	rv := objc.Send[MKMultiPolygon](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKMultiPolygonClass) New() MKMultiPolygon {
	rv := objc.Send[MKMultiPolygon](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKMultiPolygon) Init() MKMultiPolygon {
	rv := objc.Send[MKMultiPolygon](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKMultiPolygon) Autorelease() MKMultiPolygon {
	rv := objc.Send[MKMultiPolygon](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKMultiPolygon creates a new MKMultiPolygon instance.
func NewMKMultiPolygon() MKMultiPolygon {
	return getMKMultiPolygonClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKMultiPolygon */
// A collection of multiple closed polygon overlays.
//
// Use a when you have multiple distinct polygon shapes that you intend to render using the same style.


// A collection of multiple closed polygon overlays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMultiPolygon
type MKMultiPolygon struct {
	MKShape
}

// MKMultiPolygonFrom constructs a [MKMultiPolygon] from an unsafe.Pointer.
//
// A collection of multiple closed polygon overlays.
func MKMultiPolygonFrom(ptr unsafe.Pointer) MKMultiPolygon {
	return MKMultiPolygon{
		MKShape: MKShapeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKMultiPolygon */

// Creates a multipolygon object using the provided polygons.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMultiPolygon/init(_:)
func NewMKMultiPolygonWithPolygons(polygons []MKPolygon) MKMultiPolygon {
	instance := getMKMultiPolygonClass().Alloc()
	rv := objc.Send[MKMultiPolygon](instance.ID, objc.Sel("initWithPolygons:"), polygons)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKMultiPolygonWithPolygons */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKMultiPolygon */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKMultiPolygon */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKMultiPolygon */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKMultiPolygon */

// An array containing the polygons that make up the multipolygon object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMultiPolygon/polygons
func (m_ MKMultiPolygon) Polygons() []MKPolygon {
	rv := objc.Send[[]MKPolygon](m_.ID, objc.Sel("polygons"))
	return rv
}/* debug [instance_properties/getter]: polygons */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKMultiPolygon */


