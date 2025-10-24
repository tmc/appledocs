// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKCircleRenderer */


/* debug [class_header]: Header for MKCircleRenderer */
// The class instance for the [MKCircleRenderer] class.
var (
	MKCircleRendererClass     _MKCircleRendererClass
	MKCircleRendererClassOnce sync.Once
)

func getMKCircleRendererClass() _MKCircleRendererClass {
	MKCircleRendererClassOnce.Do(func() {
		MKCircleRendererClass = _MKCircleRendererClass{objc.GetClass("MKCircleRenderer")}
	})
	return MKCircleRendererClass
}

type _MKCircleRendererClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKCircleRenderer */
// An interface definition for the [MKCircleRenderer] class.
type IMKCircleRenderer interface {
	IMKOverlayPathRenderer
	
/* debug [class_interface_properties]: Properties for MKCircleRenderer */
	// properties:
	Circle() IMKCircle
	StrokeEnd() float64
	SetStrokeEnd(value float64)
	StrokeStart() float64
	SetStrokeStart(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKCircleRenderer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKCircleRenderer */
// Alloc allocates a new instance without initialization.
func (mc _MKCircleRendererClass) Alloc() MKCircleRenderer {
	rv := objc.Send[MKCircleRenderer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKCircleRendererClass) New() MKCircleRenderer {
	rv := objc.Send[MKCircleRenderer](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKCircleRenderer) Init() MKCircleRenderer {
	rv := objc.Send[MKCircleRenderer](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKCircleRenderer) Autorelease() MKCircleRenderer {
	rv := objc.Send[MKCircleRenderer](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKCircleRenderer creates a new MKCircleRenderer instance.
func NewMKCircleRenderer() MKCircleRenderer {
	return getMKCircleRendererClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKCircleRenderer */
// The visual representation of a circular overlay.
//
// This renderer fills and strokes the circular region that the overlay object represents. You can change the color and other drawing attributes of the circle by modifying the properties it inherits from the main class. You typically use this class as-is and don’t subclass it. You create an instance of this class in your map view delegate’s method.


// The visual representation of a circular overlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCircleRenderer
type MKCircleRenderer struct {
	MKOverlayPathRenderer
}

// MKCircleRendererFrom constructs a [MKCircleRenderer] from an unsafe.Pointer.
//
// The visual representation of a circular overlay.
func MKCircleRendererFrom(ptr unsafe.Pointer) MKCircleRenderer {
	return MKCircleRenderer{
		MKOverlayPathRenderer: MKOverlayPathRendererFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKCircleRenderer */

// Creates a new overlay view using the specified circle overlay object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCircleRenderer/init(circle:)
func NewMKCircleRendererWithCircle(circle IMKCircle) MKCircleRenderer {
	instance := getMKCircleRendererClass().Alloc()
	rv := objc.Send[MKCircleRenderer](instance.ID, objc.Sel("initWithCircle:"), circle)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKCircleRendererWithCircle */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKCircleRenderer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKCircleRenderer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKCircleRenderer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKCircleRenderer */

// The circle overlay object that contains the information for drawing the overlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCircleRenderer/circle
func (m_ MKCircleRenderer) Circle() IMKCircle {
	rv := objc.Send[MKCircle](m_.ID, objc.Sel("circle"))
	return rv
}/* debug [instance_properties/getter]: circle */


// The unit distance along the circle where the stroke ends.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCircleRenderer/strokeEnd
func (m_ MKCircleRenderer) StrokeEnd() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("strokeEnd"))
	return rv
}/* debug [instance_properties/getter]: strokeEnd */


// The unit distance along the circle where the stroke ends.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCircleRenderer/strokeEnd
func (m_ MKCircleRenderer) SetStrokeEnd(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStrokeEnd:"), value)
}/* debug [instance_properties/setter]: strokeEnd */


// The unit distance along the circle where the stroke starts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCircleRenderer/strokeStart
func (m_ MKCircleRenderer) StrokeStart() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("strokeStart"))
	return rv
}/* debug [instance_properties/getter]: strokeStart */


// The unit distance along the circle where the stroke starts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCircleRenderer/strokeStart
func (m_ MKCircleRenderer) SetStrokeStart(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStrokeStart:"), value)
}/* debug [instance_properties/setter]: strokeStart */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKCircleRenderer */


