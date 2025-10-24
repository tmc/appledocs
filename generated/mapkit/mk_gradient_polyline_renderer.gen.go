// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MKGradientPolylineRenderer */


/* debug [class_header]: Header for MKGradientPolylineRenderer */
// The class instance for the [MKGradientPolylineRenderer] class.
var (
	MKGradientPolylineRendererClass     _MKGradientPolylineRendererClass
	MKGradientPolylineRendererClassOnce sync.Once
)

func getMKGradientPolylineRendererClass() _MKGradientPolylineRendererClass {
	MKGradientPolylineRendererClassOnce.Do(func() {
		MKGradientPolylineRendererClass = _MKGradientPolylineRendererClass{objc.GetClass("MKGradientPolylineRenderer")}
	})
	return MKGradientPolylineRendererClass
}

type _MKGradientPolylineRendererClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKGradientPolylineRenderer */
// An interface definition for the [MKGradientPolylineRenderer] class.
type IMKGradientPolylineRenderer interface {
	IMKPolylineRenderer
	
/* debug [class_interface_properties]: Properties for MKGradientPolylineRenderer */
	// properties:
	Colors() []appkit.Color
	Locations() []foundation.Number
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKGradientPolylineRenderer */
	// methods:
	SetColorsAtLocations(colors []appkit.Color, locations []foundation.Number)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKGradientPolylineRenderer */
// Alloc allocates a new instance without initialization.
func (mc _MKGradientPolylineRendererClass) Alloc() MKGradientPolylineRenderer {
	rv := objc.Send[MKGradientPolylineRenderer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKGradientPolylineRendererClass) New() MKGradientPolylineRenderer {
	rv := objc.Send[MKGradientPolylineRenderer](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKGradientPolylineRenderer) Init() MKGradientPolylineRenderer {
	rv := objc.Send[MKGradientPolylineRenderer](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKGradientPolylineRenderer) Autorelease() MKGradientPolylineRenderer {
	rv := objc.Send[MKGradientPolylineRenderer](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKGradientPolylineRenderer creates a new MKGradientPolylineRenderer instance.
func NewMKGradientPolylineRenderer() MKGradientPolylineRenderer {
	return getMKGradientPolylineRendererClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKGradientPolylineRenderer */
// A visual representation of any polyline overlay object with a gradient.
//
// This renderer only applies a stroke to the line; it doesn’t fill it. Set the gradients with and pair colors to locations that MapKit represents as unit distance values along the distance of the polyline. Don’t subclass . Use the class as-is. The gradient displays itself along the direction of the line.


// A visual representation of any polyline overlay object with a gradient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKGradientPolylineRenderer
type MKGradientPolylineRenderer struct {
	MKPolylineRenderer
}

// MKGradientPolylineRendererFrom constructs a [MKGradientPolylineRenderer] from an unsafe.Pointer.
//
// A visual representation of any polyline overlay object with a gradient.
func MKGradientPolylineRendererFrom(ptr unsafe.Pointer) MKGradientPolylineRenderer {
	return MKGradientPolylineRenderer{
		MKPolylineRenderer: MKPolylineRendererFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKGradientPolylineRenderer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKGradientPolylineRenderer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKGradientPolylineRenderer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKGradientPolylineRenderer */

// Sets the colors and corresponding unit distance values to create gradients.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKGradientPolylineRenderer/setColors:atLocations:
func (m_ MKGradientPolylineRenderer) SetColorsAtLocations(colors []appkit.Color, locations []foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColors:atLocations:"), colors, locations)
}/* debug [instance_methods/method]: SetColorsAtLocations */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKGradientPolylineRenderer */

// An array that represents the gradient’s color transition points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKGradientPolylineRenderer/colors
func (m_ MKGradientPolylineRenderer) Colors() []appkit.Color {
	rv := objc.Send[[]appkit.Color](m_.ID, objc.Sel("colors"))
	return rv
}/* debug [instance_properties/getter]: colors */


// An array of location indices corresponding to their respective colors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKGradientPolylineRenderer/locations-50knt
func (m_ MKGradientPolylineRenderer) Locations() []foundation.Number {
	rv := objc.Send[[]foundation.Number](m_.ID, objc.Sel("locations"))
	return rv
}/* debug [instance_properties/getter]: locations */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKGradientPolylineRenderer */



