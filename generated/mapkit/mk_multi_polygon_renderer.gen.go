// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKMultiPolygonRenderer */


/* debug [class_header]: Header for MKMultiPolygonRenderer */
// The class instance for the [MKMultiPolygonRenderer] class.
var (
	MKMultiPolygonRendererClass     _MKMultiPolygonRendererClass
	MKMultiPolygonRendererClassOnce sync.Once
)

func getMKMultiPolygonRendererClass() _MKMultiPolygonRendererClass {
	MKMultiPolygonRendererClassOnce.Do(func() {
		MKMultiPolygonRendererClass = _MKMultiPolygonRendererClass{objc.GetClass("MKMultiPolygonRenderer")}
	})
	return MKMultiPolygonRendererClass
}

type _MKMultiPolygonRendererClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKMultiPolygonRenderer */
// An interface definition for the [MKMultiPolygonRenderer] class.
type IMKMultiPolygonRenderer interface {
	IMKOverlayPathRenderer
	
/* debug [class_interface_properties]: Properties for MKMultiPolygonRenderer */
	// properties:
	MultiPolygon() IMKMultiPolygon
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKMultiPolygonRenderer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKMultiPolygonRenderer */
// Alloc allocates a new instance without initialization.
func (mc _MKMultiPolygonRendererClass) Alloc() MKMultiPolygonRenderer {
	rv := objc.Send[MKMultiPolygonRenderer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKMultiPolygonRendererClass) New() MKMultiPolygonRenderer {
	rv := objc.Send[MKMultiPolygonRenderer](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKMultiPolygonRenderer) Init() MKMultiPolygonRenderer {
	rv := objc.Send[MKMultiPolygonRenderer](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKMultiPolygonRenderer) Autorelease() MKMultiPolygonRenderer {
	rv := objc.Send[MKMultiPolygonRenderer](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKMultiPolygonRenderer creates a new MKMultiPolygonRenderer instance.
func NewMKMultiPolygonRenderer() MKMultiPolygonRenderer {
	return getMKMultiPolygonRendererClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKMultiPolygonRenderer */
// The visual representation of multiple polygon overlays.
//
// Use this renderer to provide the style for multiple polygons created using .


// The visual representation of multiple polygon overlays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMultiPolygonRenderer
type MKMultiPolygonRenderer struct {
	MKOverlayPathRenderer
}

// MKMultiPolygonRendererFrom constructs a [MKMultiPolygonRenderer] from an unsafe.Pointer.
//
// The visual representation of multiple polygon overlays.
func MKMultiPolygonRendererFrom(ptr unsafe.Pointer) MKMultiPolygonRenderer {
	return MKMultiPolygonRenderer{
		MKOverlayPathRenderer: MKOverlayPathRendererFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKMultiPolygonRenderer */

// Creates and returns a renderer that handles drawing for the specified multipolygon overlay object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMultiPolygonRenderer/init(multiPolygon:)
func NewMKMultiPolygonRendererWithMultiPolygon(multiPolygon IMKMultiPolygon) MKMultiPolygonRenderer {
	instance := getMKMultiPolygonRendererClass().Alloc()
	rv := objc.Send[MKMultiPolygonRenderer](instance.ID, objc.Sel("initWithMultiPolygon:"), multiPolygon)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKMultiPolygonRendererWithMultiPolygon */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKMultiPolygonRenderer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKMultiPolygonRenderer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKMultiPolygonRenderer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKMultiPolygonRenderer */

// The multipolygon object that the renderer uses to draw the overlay’s contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMultiPolygonRenderer/multiPolygon
func (m_ MKMultiPolygonRenderer) MultiPolygon() IMKMultiPolygon {
	rv := objc.Send[MKMultiPolygon](m_.ID, objc.Sel("multiPolygon"))
	return rv
}/* debug [instance_properties/getter]: multiPolygon */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKMultiPolygonRenderer */


