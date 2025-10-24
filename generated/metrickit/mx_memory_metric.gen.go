// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MXMemoryMetric */


/* debug [class_header]: Header for MXMemoryMetric */
// The class instance for the [MXMemoryMetric] class.
var (
	MXMemoryMetricClass     _MXMemoryMetricClass
	MXMemoryMetricClassOnce sync.Once
)

func getMXMemoryMetricClass() _MXMemoryMetricClass {
	MXMemoryMetricClassOnce.Do(func() {
		MXMemoryMetricClass = _MXMemoryMetricClass{objc.GetClass("MXMemoryMetric")}
	})
	return MXMemoryMetricClass
}

type _MXMemoryMetricClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MXMemoryMetric */
// An interface definition for the [MXMemoryMetric] class.
type IMXMemoryMetric interface {
	IMXMetric
	
/* debug [class_interface_properties]: Properties for MXMemoryMetric */
	// properties:
	AverageSuspendedMemory() unsafe.Pointer
	PeakMemoryUsage() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MXMemoryMetric */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MXMemoryMetric */
// Alloc allocates a new instance without initialization.
func (mc _MXMemoryMetricClass) Alloc() MXMemoryMetric {
	rv := objc.Send[MXMemoryMetric](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MXMemoryMetricClass) New() MXMemoryMetric {
	rv := objc.Send[MXMemoryMetric](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXMemoryMetric) Init() MXMemoryMetric {
	rv := objc.Send[MXMemoryMetric](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXMemoryMetric) Autorelease() MXMemoryMetric {
	rv := objc.Send[MXMemoryMetric](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXMemoryMetric creates a new MXMemoryMetric instance.
func NewMXMemoryMetric() MXMemoryMetric {
	return getMXMemoryMetricClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MXMemoryMetric */
// An object representing metrics about the app’s memory use.


// An object representing metrics about the app’s memory use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMemoryMetric
type MXMemoryMetric struct {
	MXMetric
}

// MXMemoryMetricFrom constructs a [MXMemoryMetric] from an unsafe.Pointer.
//
// An object representing metrics about the app’s memory use.
func MXMemoryMetricFrom(ptr unsafe.Pointer) MXMemoryMetric {
	return MXMemoryMetric{
		MXMetric: MXMetricFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MXMemoryMetric *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MXMemoryMetric */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MXMemoryMetric */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MXMemoryMetric */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MXMemoryMetric */

// The average amount of memory in use by the app when it’s suspended.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMemoryMetric/averageSuspendedMemory
func (m_ MXMemoryMetric) AverageSuspendedMemory() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("averageSuspendedMemory"))
	return rv
}/* debug [instance_properties/getter]: averageSuspendedMemory */


// The largest amount of memory used by the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMemoryMetric/peakMemoryUsage
func (m_ MXMemoryMetric) PeakMemoryUsage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("peakMemoryUsage"))
	return rv
}/* debug [instance_properties/getter]: peakMemoryUsage */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MXMemoryMetric */



