// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKMultiPolylineRenderer */


/* debug [class_header]: Header for MKMultiPolylineRenderer */
// The class instance for the [MKMultiPolylineRenderer] class.
var (
	MKMultiPolylineRendererClass     _MKMultiPolylineRendererClass
	MKMultiPolylineRendererClassOnce sync.Once
)

func getMKMultiPolylineRendererClass() _MKMultiPolylineRendererClass {
	MKMultiPolylineRendererClassOnce.Do(func() {
		MKMultiPolylineRendererClass = _MKMultiPolylineRendererClass{objc.GetClass("MKMultiPolylineRenderer")}
	})
	return MKMultiPolylineRendererClass
}

type _MKMultiPolylineRendererClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKMultiPolylineRenderer */
// An interface definition for the [MKMultiPolylineRenderer] class.
type IMKMultiPolylineRenderer interface {
	IMKOverlayPathRenderer
	
/* debug [class_interface_properties]: Properties for MKMultiPolylineRenderer */
	// properties:
	MultiPolyline() IMKMultiPolyline
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKMultiPolylineRenderer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKMultiPolylineRenderer */
// Alloc allocates a new instance without initialization.
func (mc _MKMultiPolylineRendererClass) Alloc() MKMultiPolylineRenderer {
	rv := objc.Send[MKMultiPolylineRenderer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKMultiPolylineRendererClass) New() MKMultiPolylineRenderer {
	rv := objc.Send[MKMultiPolylineRenderer](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKMultiPolylineRenderer) Init() MKMultiPolylineRenderer {
	rv := objc.Send[MKMultiPolylineRenderer](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKMultiPolylineRenderer) Autorelease() MKMultiPolylineRenderer {
	rv := objc.Send[MKMultiPolylineRenderer](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKMultiPolylineRenderer creates a new MKMultiPolylineRenderer instance.
func NewMKMultiPolylineRenderer() MKMultiPolylineRenderer {
	return getMKMultiPolylineRendererClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKMultiPolylineRenderer */
// A visual representation of multiple polyline overlay objects.
//
// Use the multipolyline renderer to provide the styling of multiple polylines that you create using .


// A visual representation of multiple polyline overlay objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMultiPolylineRenderer
type MKMultiPolylineRenderer struct {
	MKOverlayPathRenderer
}

// MKMultiPolylineRendererFrom constructs a [MKMultiPolylineRenderer] from an unsafe.Pointer.
//
// A visual representation of multiple polyline overlay objects.
func MKMultiPolylineRendererFrom(ptr unsafe.Pointer) MKMultiPolylineRenderer {
	return MKMultiPolylineRenderer{
		MKOverlayPathRenderer: MKOverlayPathRendererFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKMultiPolylineRenderer */

// Creates an object that renders a visual representation of multiple polyline objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMultiPolylineRenderer/init(multiPolyline:)
func NewMKMultiPolylineRendererWithMultiPolyline(multiPolyline IMKMultiPolyline) MKMultiPolylineRenderer {
	instance := getMKMultiPolylineRendererClass().Alloc()
	rv := objc.Send[MKMultiPolylineRenderer](instance.ID, objc.Sel("initWithMultiPolyline:"), multiPolyline)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKMultiPolylineRendererWithMultiPolyline */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKMultiPolylineRenderer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKMultiPolylineRenderer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKMultiPolylineRenderer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKMultiPolylineRenderer */

// An object that represents multiple polyline shapes, each consisting of one or more connected line segments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMultiPolylineRenderer/multiPolyline
func (m_ MKMultiPolylineRenderer) MultiPolyline() IMKMultiPolyline {
	rv := objc.Send[MKMultiPolyline](m_.ID, objc.Sel("multiPolyline"))
	return rv
}/* debug [instance_properties/getter]: multiPolyline */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKMultiPolylineRenderer */


