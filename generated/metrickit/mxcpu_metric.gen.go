// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MXCPUMetric */


/* debug [class_header]: Header for MXCPUMetric */
// The class instance for the [MXCPUMetric] class.
var (
	MXCPUMetricClass     _MXCPUMetricClass
	MXCPUMetricClassOnce sync.Once
)

func getMXCPUMetricClass() _MXCPUMetricClass {
	MXCPUMetricClassOnce.Do(func() {
		MXCPUMetricClass = _MXCPUMetricClass{objc.GetClass("MXCPUMetric")}
	})
	return MXCPUMetricClass
}

type _MXCPUMetricClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MXCPUMetric */
// An interface definition for the [MXCPUMetric] class.
type IMXCPUMetric interface {
	IMXMetric
	
/* debug [class_interface_properties]: Properties for MXCPUMetric */
	// properties:
	CumulativeCPUInstructions() unsafe.Pointer
	CumulativeCPUTime() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MXCPUMetric */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MXCPUMetric */
// Alloc allocates a new instance without initialization.
func (mc _MXCPUMetricClass) Alloc() MXCPUMetric {
	rv := objc.Send[MXCPUMetric](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MXCPUMetricClass) New() MXCPUMetric {
	rv := objc.Send[MXCPUMetric](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXCPUMetric) Init() MXCPUMetric {
	rv := objc.Send[MXCPUMetric](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXCPUMetric) Autorelease() MXCPUMetric {
	rv := objc.Send[MXCPUMetric](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXCPUMetric creates a new MXCPUMetric instance.
func NewMXCPUMetric() MXCPUMetric {
	return getMXCPUMetricClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MXCPUMetric */
// An object representing metrics about the use of the CPU.


// An object representing metrics about the use of the CPU.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCPUMetric
type MXCPUMetric struct {
	MXMetric
}

// MXCPUMetricFrom constructs a [MXCPUMetric] from an unsafe.Pointer.
//
// An object representing metrics about the use of the CPU.
func MXCPUMetricFrom(ptr unsafe.Pointer) MXCPUMetric {
	return MXCPUMetric{
		MXMetric: MXMetricFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MXCPUMetric *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MXCPUMetric */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MXCPUMetric */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MXCPUMetric */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MXCPUMetric */

// The total number of CPU instructions the app executed during the reporting period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCPUMetric/cumulativeCPUInstructions
func (m_ MXCPUMetric) CumulativeCPUInstructions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cumulativeCPUInstructions"))
	return rv
}/* debug [instance_properties/getter]: cumulativeCPUInstructions */


// The total amount of CPU the app used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCPUMetric/cumulativeCPUTime
func (m_ MXCPUMetric) CumulativeCPUTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cumulativeCPUTime"))
	return rv
}/* debug [instance_properties/getter]: cumulativeCPUTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MXCPUMetric */



