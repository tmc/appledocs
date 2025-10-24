// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MXGPUMetric */


/* debug [class_header]: Header for MXGPUMetric */
// The class instance for the [MXGPUMetric] class.
var (
	MXGPUMetricClass     _MXGPUMetricClass
	MXGPUMetricClassOnce sync.Once
)

func getMXGPUMetricClass() _MXGPUMetricClass {
	MXGPUMetricClassOnce.Do(func() {
		MXGPUMetricClass = _MXGPUMetricClass{objc.GetClass("MXGPUMetric")}
	})
	return MXGPUMetricClass
}

type _MXGPUMetricClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MXGPUMetric */
// An interface definition for the [MXGPUMetric] class.
type IMXGPUMetric interface {
	IMXMetric
	
/* debug [class_interface_properties]: Properties for MXGPUMetric */
	// properties:
	CumulativeGPUTime() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MXGPUMetric */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MXGPUMetric */
// Alloc allocates a new instance without initialization.
func (mc _MXGPUMetricClass) Alloc() MXGPUMetric {
	rv := objc.Send[MXGPUMetric](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MXGPUMetricClass) New() MXGPUMetric {
	rv := objc.Send[MXGPUMetric](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXGPUMetric) Init() MXGPUMetric {
	rv := objc.Send[MXGPUMetric](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXGPUMetric) Autorelease() MXGPUMetric {
	rv := objc.Send[MXGPUMetric](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXGPUMetric creates a new MXGPUMetric instance.
func NewMXGPUMetric() MXGPUMetric {
	return getMXGPUMetricClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MXGPUMetric */
// An object representing metrics about the use of the GPU.


// An object representing metrics about the use of the GPU.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXGPUMetric
type MXGPUMetric struct {
	MXMetric
}

// MXGPUMetricFrom constructs a [MXGPUMetric] from an unsafe.Pointer.
//
// An object representing metrics about the use of the GPU.
func MXGPUMetricFrom(ptr unsafe.Pointer) MXGPUMetric {
	return MXGPUMetric{
		MXMetric: MXMetricFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MXGPUMetric *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MXGPUMetric */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MXGPUMetric */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MXGPUMetric */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MXGPUMetric */

// The total amount of GPU time used by the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXGPUMetric/cumulativeGPUTime
func (m_ MXGPUMetric) CumulativeGPUTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cumulativeGPUTime"))
	return rv
}/* debug [instance_properties/getter]: cumulativeGPUTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MXGPUMetric */



