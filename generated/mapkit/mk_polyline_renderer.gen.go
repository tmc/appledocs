// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKPolylineRenderer */


/* debug [class_header]: Header for MKPolylineRenderer */
// The class instance for the [MKPolylineRenderer] class.
var (
	MKPolylineRendererClass     _MKPolylineRendererClass
	MKPolylineRendererClassOnce sync.Once
)

func getMKPolylineRendererClass() _MKPolylineRendererClass {
	MKPolylineRendererClassOnce.Do(func() {
		MKPolylineRendererClass = _MKPolylineRendererClass{objc.GetClass("MKPolylineRenderer")}
	})
	return MKPolylineRendererClass
}

type _MKPolylineRendererClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKPolylineRenderer */
// An interface definition for the [MKPolylineRenderer] class.
type IMKPolylineRenderer interface {
	IMKOverlayPathRenderer
	
/* debug [class_interface_properties]: Properties for MKPolylineRenderer */
	// properties:
	Polyline() IMKPolyline
	StrokeEnd() float64
	SetStrokeEnd(value float64)
	StrokeStart() float64
	SetStrokeStart(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKPolylineRenderer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKPolylineRenderer */
// Alloc allocates a new instance without initialization.
func (mc _MKPolylineRendererClass) Alloc() MKPolylineRenderer {
	rv := objc.Send[MKPolylineRenderer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKPolylineRendererClass) New() MKPolylineRenderer {
	rv := objc.Send[MKPolylineRenderer](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKPolylineRenderer) Init() MKPolylineRenderer {
	rv := objc.Send[MKPolylineRenderer](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKPolylineRenderer) Autorelease() MKPolylineRenderer {
	rv := objc.Send[MKPolylineRenderer](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKPolylineRenderer creates a new MKPolylineRenderer instance.
func NewMKPolylineRenderer() MKPolylineRenderer {
	return getMKPolylineRendererClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKPolylineRenderer */
// A visual representation of any polyline overlay object.
//
// This renderer strokes the line only; it doesn’t fill it. You can change the color and other drawing attributes of the polyline by modifying the properties it inherits from the main class. You typically use this class as-is and don’t subclass it.


// A visual representation of any polyline overlay object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolylineRenderer
type MKPolylineRenderer struct {
	MKOverlayPathRenderer
}

// MKPolylineRendererFrom constructs a [MKPolylineRenderer] from an unsafe.Pointer.
//
// A visual representation of any polyline overlay object.
func MKPolylineRendererFrom(ptr unsafe.Pointer) MKPolylineRenderer {
	return MKPolylineRenderer{
		MKOverlayPathRenderer: MKOverlayPathRendererFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKPolylineRenderer */

// Creates a new overlay view using the specified polyline overlay object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolylineRenderer/init(polyline:)
func NewMKPolylineRendererWithPolyline(polyline IMKPolyline) MKPolylineRenderer {
	instance := getMKPolylineRendererClass().Alloc()
	rv := objc.Send[MKPolylineRenderer](instance.ID, objc.Sel("initWithPolyline:"), polyline)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKPolylineRendererWithPolyline */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKPolylineRenderer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKPolylineRenderer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKPolylineRenderer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKPolylineRenderer */

// The polyline overlay object that contains the information for drawing the overlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolylineRenderer/polyline
func (m_ MKPolylineRenderer) Polyline() IMKPolyline {
	rv := objc.Send[MKPolyline](m_.ID, objc.Sel("polyline"))
	return rv
}/* debug [instance_properties/getter]: polyline */


// The unit distance along the line where the stroke ends.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolylineRenderer/strokeEnd
func (m_ MKPolylineRenderer) StrokeEnd() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("strokeEnd"))
	return rv
}/* debug [instance_properties/getter]: strokeEnd */


// The unit distance along the line where the stroke ends.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolylineRenderer/strokeEnd
func (m_ MKPolylineRenderer) SetStrokeEnd(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStrokeEnd:"), value)
}/* debug [instance_properties/setter]: strokeEnd */


// The unit distance along the line where the stroke starts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolylineRenderer/strokeStart
func (m_ MKPolylineRenderer) StrokeStart() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("strokeStart"))
	return rv
}/* debug [instance_properties/getter]: strokeStart */


// The unit distance along the line where the stroke starts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolylineRenderer/strokeStart
func (m_ MKPolylineRenderer) SetStrokeStart(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStrokeStart:"), value)
}/* debug [instance_properties/setter]: strokeStart */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKPolylineRenderer */


