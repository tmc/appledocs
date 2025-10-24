// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MXAppResponsivenessMetric */


/* debug [class_header]: Header for MXAppResponsivenessMetric */
// The class instance for the [MXAppResponsivenessMetric] class.
var (
	MXAppResponsivenessMetricClass     _MXAppResponsivenessMetricClass
	MXAppResponsivenessMetricClassOnce sync.Once
)

func getMXAppResponsivenessMetricClass() _MXAppResponsivenessMetricClass {
	MXAppResponsivenessMetricClassOnce.Do(func() {
		MXAppResponsivenessMetricClass = _MXAppResponsivenessMetricClass{objc.GetClass("MXAppResponsivenessMetric")}
	})
	return MXAppResponsivenessMetricClass
}

type _MXAppResponsivenessMetricClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MXAppResponsivenessMetric */
// An interface definition for the [MXAppResponsivenessMetric] class.
type IMXAppResponsivenessMetric interface {
	IMXMetric
	
/* debug [class_interface_properties]: Properties for MXAppResponsivenessMetric */
	// properties:
	HistogrammedApplicationHangTime() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MXAppResponsivenessMetric */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MXAppResponsivenessMetric */
// Alloc allocates a new instance without initialization.
func (mc _MXAppResponsivenessMetricClass) Alloc() MXAppResponsivenessMetric {
	rv := objc.Send[MXAppResponsivenessMetric](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MXAppResponsivenessMetricClass) New() MXAppResponsivenessMetric {
	rv := objc.Send[MXAppResponsivenessMetric](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXAppResponsivenessMetric) Init() MXAppResponsivenessMetric {
	rv := objc.Send[MXAppResponsivenessMetric](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXAppResponsivenessMetric) Autorelease() MXAppResponsivenessMetric {
	rv := objc.Send[MXAppResponsivenessMetric](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXAppResponsivenessMetric creates a new MXAppResponsivenessMetric instance.
func NewMXAppResponsivenessMetric() MXAppResponsivenessMetric {
	return getMXAppResponsivenessMetricClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MXAppResponsivenessMetric */
// An object representing metrics about the responsiveness of the app to user interaction.


// An object representing metrics about the responsiveness of the app to user interaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAppResponsivenessMetric
type MXAppResponsivenessMetric struct {
	MXMetric
}

// MXAppResponsivenessMetricFrom constructs a [MXAppResponsivenessMetric] from an unsafe.Pointer.
//
// An object representing metrics about the responsiveness of the app to user interaction.
func MXAppResponsivenessMetricFrom(ptr unsafe.Pointer) MXAppResponsivenessMetric {
	return MXAppResponsivenessMetric{
		MXMetric: MXMetricFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MXAppResponsivenessMetric *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MXAppResponsivenessMetric */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MXAppResponsivenessMetric */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MXAppResponsivenessMetric */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MXAppResponsivenessMetric */

// A histogram of the different durations of time in which the app is too busy to handle user interaction responsively.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAppResponsivenessMetric/histogrammedApplicationHangTime
func (m_ MXAppResponsivenessMetric) HistogrammedApplicationHangTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("histogrammedApplicationHangTime"))
	return rv
}/* debug [instance_properties/getter]: histogrammedApplicationHangTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MXAppResponsivenessMetric */



