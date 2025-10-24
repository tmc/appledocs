// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MXUnitAveragePixelLuminance */


/* debug [class_header]: Header for MXUnitAveragePixelLuminance */
// The class instance for the [MXUnitAveragePixelLuminance] class.
var (
	MXUnitAveragePixelLuminanceClass     _MXUnitAveragePixelLuminanceClass
	MXUnitAveragePixelLuminanceClassOnce sync.Once
)

func getMXUnitAveragePixelLuminanceClass() _MXUnitAveragePixelLuminanceClass {
	MXUnitAveragePixelLuminanceClassOnce.Do(func() {
		MXUnitAveragePixelLuminanceClass = _MXUnitAveragePixelLuminanceClass{objc.GetClass("MXUnitAveragePixelLuminance")}
	})
	return MXUnitAveragePixelLuminanceClass
}

type _MXUnitAveragePixelLuminanceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MXUnitAveragePixelLuminance */
// An interface definition for the [MXUnitAveragePixelLuminance] class.
type IMXUnitAveragePixelLuminance interface {
	foundation.IDimension
	
/* debug [class_interface_properties]: Properties for MXUnitAveragePixelLuminance */
	// properties:
	AveragePixelLuminance() IMXUnitAveragePixelLuminance
	SetAveragePixelLuminance(value IMXUnitAveragePixelLuminance)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MXUnitAveragePixelLuminance */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MXUnitAveragePixelLuminance */
// Alloc allocates a new instance without initialization.
func (mc _MXUnitAveragePixelLuminanceClass) Alloc() MXUnitAveragePixelLuminance {
	rv := objc.Send[MXUnitAveragePixelLuminance](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MXUnitAveragePixelLuminanceClass) New() MXUnitAveragePixelLuminance {
	rv := objc.Send[MXUnitAveragePixelLuminance](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXUnitAveragePixelLuminance) Init() MXUnitAveragePixelLuminance {
	rv := objc.Send[MXUnitAveragePixelLuminance](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXUnitAveragePixelLuminance) Autorelease() MXUnitAveragePixelLuminance {
	rv := objc.Send[MXUnitAveragePixelLuminance](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXUnitAveragePixelLuminance creates a new MXUnitAveragePixelLuminance instance.
func NewMXUnitAveragePixelLuminance() MXUnitAveragePixelLuminance {
	return getMXUnitAveragePixelLuminanceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MXUnitAveragePixelLuminance */
// A unit of measure of pixel luminosity on an OLED display.
//
// Luminosity represents the brightness of each red, green, and blue component pixel. Unlike LCD displays, each pixel requires power to display a color, and white draws the most power per pixel. defines the base unit as the average luminance of all the pixels on the screen for some period of time. Reducing the average luminance of the display reduces the amount of power consumed by the app.


// A unit of measure of pixel luminosity on an OLED display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXUnitAveragePixelLuminance
type MXUnitAveragePixelLuminance struct {
	foundation.Dimension
}

// MXUnitAveragePixelLuminanceFrom constructs a [MXUnitAveragePixelLuminance] from an unsafe.Pointer.
//
// A unit of measure of pixel luminosity on an OLED display.
func MXUnitAveragePixelLuminanceFrom(ptr unsafe.Pointer) MXUnitAveragePixelLuminance {
	return MXUnitAveragePixelLuminance{
		Dimension: foundation.DimensionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MXUnitAveragePixelLuminance *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MXUnitAveragePixelLuminance */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MXUnitAveragePixelLuminance */

// The average number of powered pixels on a OLED display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXUnitAveragePixelLuminance/apl
func (mc _MXUnitAveragePixelLuminanceClass) Apl() MXUnitAveragePixelLuminance {
	rv := objc.Send[MXUnitAveragePixelLuminance](objc.ID(mc.class), objc.Sel("apl"))
	return rv
}/* debug [class_properties_class/property]: apl */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MXUnitAveragePixelLuminance */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MXUnitAveragePixelLuminance */

// The average number of powered pixels on a OLED display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXUnitAveragePixelLuminance/apl
func (m_ MXUnitAveragePixelLuminance) Apl() IMXUnitAveragePixelLuminance {
	rv := objc.Send[MXUnitAveragePixelLuminance](m_.ID, objc.Sel("apl"))
	return rv
}/* debug [instance_properties/getter]: apl */


// The average amount of luminosity of the pixels on an OLED display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxdisplaymetric/averagepixelluminance
func (m_ MXUnitAveragePixelLuminance) AveragePixelLuminance() IMXUnitAveragePixelLuminance {
	rv := objc.Send[MXUnitAveragePixelLuminance](m_.ID, objc.Sel("averagePixelLuminance"))
	return rv
}/* debug [instance_properties/getter]: averagePixelLuminance */


// The average amount of luminosity of the pixels on an OLED display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxdisplaymetric/averagepixelluminance
func (m_ MXUnitAveragePixelLuminance) SetAveragePixelLuminance(value IMXUnitAveragePixelLuminance) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAveragePixelLuminance:"), value)
}/* debug [instance_properties/setter]: averagePixelLuminance */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MXUnitAveragePixelLuminance */



