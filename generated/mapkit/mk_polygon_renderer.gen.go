// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKPolygonRenderer */


/* debug [class_header]: Header for MKPolygonRenderer */
// The class instance for the [MKPolygonRenderer] class.
var (
	MKPolygonRendererClass     _MKPolygonRendererClass
	MKPolygonRendererClassOnce sync.Once
)

func getMKPolygonRendererClass() _MKPolygonRendererClass {
	MKPolygonRendererClassOnce.Do(func() {
		MKPolygonRendererClass = _MKPolygonRendererClass{objc.GetClass("MKPolygonRenderer")}
	})
	return MKPolygonRendererClass
}

type _MKPolygonRendererClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKPolygonRenderer */
// An interface definition for the [MKPolygonRenderer] class.
type IMKPolygonRenderer interface {
	IMKOverlayPathRenderer
	
/* debug [class_interface_properties]: Properties for MKPolygonRenderer */
	// properties:
	Polygon() IMKPolygon
	StrokeEnd() float64
	SetStrokeEnd(value float64)
	StrokeStart() float64
	SetStrokeStart(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKPolygonRenderer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKPolygonRenderer */
// Alloc allocates a new instance without initialization.
func (mc _MKPolygonRendererClass) Alloc() MKPolygonRenderer {
	rv := objc.Send[MKPolygonRenderer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKPolygonRendererClass) New() MKPolygonRenderer {
	rv := objc.Send[MKPolygonRenderer](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKPolygonRenderer) Init() MKPolygonRenderer {
	rv := objc.Send[MKPolygonRenderer](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKPolygonRenderer) Autorelease() MKPolygonRenderer {
	rv := objc.Send[MKPolygonRenderer](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKPolygonRenderer creates a new MKPolygonRenderer instance.
func NewMKPolygonRenderer() MKPolygonRenderer {
	return getMKPolygonRendererClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKPolygonRenderer */
// The visual representation of a single polygon overlay.
//
// This renderer creates the polygon overlay by first filling the shape and then representing its outline with strokes. You can change the color and other drawing attributes of the polygon by modifying the properties inherited from the parent class.


// The visual representation of a single polygon overlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolygonRenderer
type MKPolygonRenderer struct {
	MKOverlayPathRenderer
}

// MKPolygonRendererFrom constructs a [MKPolygonRenderer] from an unsafe.Pointer.
//
// The visual representation of a single polygon overlay.
func MKPolygonRendererFrom(ptr unsafe.Pointer) MKPolygonRenderer {
	return MKPolygonRenderer{
		MKOverlayPathRenderer: MKOverlayPathRendererFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKPolygonRenderer */

// Creates a new renderer that handles drawing for the specified polygon overlay object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolygonRenderer/init(polygon:)
func NewMKPolygonRendererWithPolygon(polygon IMKPolygon) MKPolygonRenderer {
	instance := getMKPolygonRendererClass().Alloc()
	rv := objc.Send[MKPolygonRenderer](instance.ID, objc.Sel("initWithPolygon:"), polygon)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKPolygonRendererWithPolygon */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKPolygonRenderer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKPolygonRenderer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKPolygonRenderer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKPolygonRenderer */

// The polygon object that contains the information used to draw the overlay’s contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolygonRenderer/polygon
func (m_ MKPolygonRenderer) Polygon() IMKPolygon {
	rv := objc.Send[MKPolygon](m_.ID, objc.Sel("polygon"))
	return rv
}/* debug [instance_properties/getter]: polygon */


// The unit distance along the polygon where the stroke ends.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolygonRenderer/strokeEnd
func (m_ MKPolygonRenderer) StrokeEnd() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("strokeEnd"))
	return rv
}/* debug [instance_properties/getter]: strokeEnd */


// The unit distance along the polygon where the stroke ends.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolygonRenderer/strokeEnd
func (m_ MKPolygonRenderer) SetStrokeEnd(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStrokeEnd:"), value)
}/* debug [instance_properties/setter]: strokeEnd */


// The unit distance along the polygon where the stroke starts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolygonRenderer/strokeStart
func (m_ MKPolygonRenderer) StrokeStart() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("strokeStart"))
	return rv
}/* debug [instance_properties/getter]: strokeStart */


// The unit distance along the polygon where the stroke starts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolygonRenderer/strokeStart
func (m_ MKPolygonRenderer) SetStrokeStart(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStrokeStart:"), value)
}/* debug [instance_properties/setter]: strokeStart */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKPolygonRenderer */


