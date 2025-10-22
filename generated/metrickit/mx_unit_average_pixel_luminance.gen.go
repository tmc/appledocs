// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [MXUnitAveragePixelLuminance] class.
type IMXUnitAveragePixelLuminance interface {
	foundation.IDimension
	AveragePixelLuminance() MXUnitAveragePixelLuminance
	SetAveragePixelLuminance(value IMXUnitAveragePixelLuminance)
}

// A unit of measure of pixel luminosity on an OLED display.
//
// Luminosity represents the brightness of each red, green, and blue component pixel. Unlike LCD displays, each pixel requires power to display a color, and white draws the most power per pixel. defines the base unit as the average luminance of all the pixels on the screen for some period of time. Reducing the average luminance of the display reduces the amount of power consumed by the app.
//
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

// Alloc allocates a new instance without initialization.
func (mc _MXUnitAveragePixelLuminanceClass) Alloc() MXUnitAveragePixelLuminance {
	rv := objc.Send[MXUnitAveragePixelLuminance](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The average amount of luminosity of the pixels on an OLED display.
//
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxdisplaymetric/averagepixelluminance
func (m_ MXUnitAveragePixelLuminance) AveragePixelLuminance() MXUnitAveragePixelLuminance {
	rv := objc.Send[MXUnitAveragePixelLuminance](m_.ID, objc.Sel("averagePixelLuminance"))
	return rv
}


// SetAveragePixelLuminance sets the value of the averagePixelLuminance property.
// The average amount of luminosity of the pixels on an OLED display.

//
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxdisplaymetric/averagepixelluminance
func (m_ MXUnitAveragePixelLuminance) SetAveragePixelLuminance(value IMXUnitAveragePixelLuminance) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAveragePixelLuminance:"), value)
}



