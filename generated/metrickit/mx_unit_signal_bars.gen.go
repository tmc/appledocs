// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MXUnitSignalBars */


/* debug [class_header]: Header for MXUnitSignalBars */
// The class instance for the [MXUnitSignalBars] class.
var (
	MXUnitSignalBarsClass     _MXUnitSignalBarsClass
	MXUnitSignalBarsClassOnce sync.Once
)

func getMXUnitSignalBarsClass() _MXUnitSignalBarsClass {
	MXUnitSignalBarsClassOnce.Do(func() {
		MXUnitSignalBarsClass = _MXUnitSignalBarsClass{objc.GetClass("MXUnitSignalBars")}
	})
	return MXUnitSignalBarsClass
}

type _MXUnitSignalBarsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MXUnitSignalBars */
// An interface definition for the [MXUnitSignalBars] class.
type IMXUnitSignalBars interface {
	foundation.IDimension
	
/* debug [class_interface_properties]: Properties for MXUnitSignalBars */
	// properties:
	HistogrammedCellularConditionTime() IMXUnitSignalBars
	SetHistogrammedCellularConditionTime(value IMXUnitSignalBars)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MXUnitSignalBars */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MXUnitSignalBars */
// Alloc allocates a new instance without initialization.
func (mc _MXUnitSignalBarsClass) Alloc() MXUnitSignalBars {
	rv := objc.Send[MXUnitSignalBars](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MXUnitSignalBarsClass) New() MXUnitSignalBars {
	rv := objc.Send[MXUnitSignalBars](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXUnitSignalBars) Init() MXUnitSignalBars {
	rv := objc.Send[MXUnitSignalBars](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXUnitSignalBars) Autorelease() MXUnitSignalBars {
	rv := objc.Send[MXUnitSignalBars](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXUnitSignalBars creates a new MXUnitSignalBars instance.
func NewMXUnitSignalBars() MXUnitSignalBars {
	return getMXUnitSignalBarsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MXUnitSignalBars */
// A unit of measure for the number of bars of cellular network connectivity.
//
// Cellular connectivity measures the relative strength of the device’s signal reception in decibels, which usually falls in a range of 0 to -110. defines the base unit as bars corresponding to the bars in the cellular connection status icon at the top of a device screen.


// A unit of measure for the number of bars of cellular network connectivity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXUnitSignalBars
type MXUnitSignalBars struct {
	foundation.Dimension
}

// MXUnitSignalBarsFrom constructs a [MXUnitSignalBars] from an unsafe.Pointer.
//
// A unit of measure for the number of bars of cellular network connectivity.
func MXUnitSignalBarsFrom(ptr unsafe.Pointer) MXUnitSignalBars {
	return MXUnitSignalBars{
		Dimension: foundation.DimensionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MXUnitSignalBars *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MXUnitSignalBars */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MXUnitSignalBars */

// The number of bars of connectivity to the cellular network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXUnitSignalBars/bars
func (mc _MXUnitSignalBarsClass) Bars() MXUnitSignalBars {
	rv := objc.Send[MXUnitSignalBars](objc.ID(mc.class), objc.Sel("bars"))
	return rv
}/* debug [class_properties_class/property]: bars */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MXUnitSignalBars */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MXUnitSignalBars */

// The number of bars of connectivity to the cellular network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXUnitSignalBars/bars
func (m_ MXUnitSignalBars) Bars() IMXUnitSignalBars {
	rv := objc.Send[MXUnitSignalBars](m_.ID, objc.Sel("bars"))
	return rv
}/* debug [instance_properties/getter]: bars */


// An object representing the distribution of the different levels of connectivity to the cellular network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxcellularconditionmetric/histogrammedcellularconditiontime
func (m_ MXUnitSignalBars) HistogrammedCellularConditionTime() IMXUnitSignalBars {
	rv := objc.Send[MXUnitSignalBars](m_.ID, objc.Sel("histogrammedCellularConditionTime"))
	return rv
}/* debug [instance_properties/getter]: histogrammedCellularConditionTime */


// An object representing the distribution of the different levels of connectivity to the cellular network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxcellularconditionmetric/histogrammedcellularconditiontime
func (m_ MXUnitSignalBars) SetHistogrammedCellularConditionTime(value IMXUnitSignalBars) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHistogrammedCellularConditionTime:"), value)
}/* debug [instance_properties/setter]: histogrammedCellularConditionTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MXUnitSignalBars */





