// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MXCellularConditionMetric */


/* debug [class_header]: Header for MXCellularConditionMetric */
// The class instance for the [MXCellularConditionMetric] class.
var (
	MXCellularConditionMetricClass     _MXCellularConditionMetricClass
	MXCellularConditionMetricClassOnce sync.Once
)

func getMXCellularConditionMetricClass() _MXCellularConditionMetricClass {
	MXCellularConditionMetricClassOnce.Do(func() {
		MXCellularConditionMetricClass = _MXCellularConditionMetricClass{objc.GetClass("MXCellularConditionMetric")}
	})
	return MXCellularConditionMetricClass
}

type _MXCellularConditionMetricClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MXCellularConditionMetric */
// An interface definition for the [MXCellularConditionMetric] class.
type IMXCellularConditionMetric interface {
	IMXMetric
	
/* debug [class_interface_properties]: Properties for MXCellularConditionMetric */
	// properties:
	HistogrammedCellularConditionTime() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MXCellularConditionMetric */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MXCellularConditionMetric */
// Alloc allocates a new instance without initialization.
func (mc _MXCellularConditionMetricClass) Alloc() MXCellularConditionMetric {
	rv := objc.Send[MXCellularConditionMetric](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MXCellularConditionMetricClass) New() MXCellularConditionMetric {
	rv := objc.Send[MXCellularConditionMetric](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXCellularConditionMetric) Init() MXCellularConditionMetric {
	rv := objc.Send[MXCellularConditionMetric](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXCellularConditionMetric) Autorelease() MXCellularConditionMetric {
	rv := objc.Send[MXCellularConditionMetric](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXCellularConditionMetric creates a new MXCellularConditionMetric instance.
func NewMXCellularConditionMetric() MXCellularConditionMetric {
	return getMXCellularConditionMetricClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MXCellularConditionMetric */
// An object representing metrics about the condition of the cellular network.


// An object representing metrics about the condition of the cellular network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCellularConditionMetric
type MXCellularConditionMetric struct {
	MXMetric
}

// MXCellularConditionMetricFrom constructs a [MXCellularConditionMetric] from an unsafe.Pointer.
//
// An object representing metrics about the condition of the cellular network.
func MXCellularConditionMetricFrom(ptr unsafe.Pointer) MXCellularConditionMetric {
	return MXCellularConditionMetric{
		MXMetric: MXMetricFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MXCellularConditionMetric *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MXCellularConditionMetric */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MXCellularConditionMetric */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MXCellularConditionMetric */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MXCellularConditionMetric */

// An object representing the distribution of the different levels of connectivity to the cellular network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCellularConditionMetric/histogrammedCellularConditionTime
func (m_ MXCellularConditionMetric) HistogrammedCellularConditionTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("histogrammedCellularConditionTime"))
	return rv
}/* debug [instance_properties/getter]: histogrammedCellularConditionTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MXCellularConditionMetric */



