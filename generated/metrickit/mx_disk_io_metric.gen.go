// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MXDiskIOMetric */


/* debug [class_header]: Header for MXDiskIOMetric */
// The class instance for the [MXDiskIOMetric] class.
var (
	MXDiskIOMetricClass     _MXDiskIOMetricClass
	MXDiskIOMetricClassOnce sync.Once
)

func getMXDiskIOMetricClass() _MXDiskIOMetricClass {
	MXDiskIOMetricClassOnce.Do(func() {
		MXDiskIOMetricClass = _MXDiskIOMetricClass{objc.GetClass("MXDiskIOMetric")}
	})
	return MXDiskIOMetricClass
}

type _MXDiskIOMetricClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MXDiskIOMetric */
// An interface definition for the [MXDiskIOMetric] class.
type IMXDiskIOMetric interface {
	IMXMetric
	
/* debug [class_interface_properties]: Properties for MXDiskIOMetric */
	// properties:
	CumulativeLogicalWrites() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MXDiskIOMetric */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MXDiskIOMetric */
// Alloc allocates a new instance without initialization.
func (mc _MXDiskIOMetricClass) Alloc() MXDiskIOMetric {
	rv := objc.Send[MXDiskIOMetric](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MXDiskIOMetricClass) New() MXDiskIOMetric {
	rv := objc.Send[MXDiskIOMetric](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXDiskIOMetric) Init() MXDiskIOMetric {
	rv := objc.Send[MXDiskIOMetric](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXDiskIOMetric) Autorelease() MXDiskIOMetric {
	rv := objc.Send[MXDiskIOMetric](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXDiskIOMetric creates a new MXDiskIOMetric instance.
func NewMXDiskIOMetric() MXDiskIOMetric {
	return getMXDiskIOMetricClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MXDiskIOMetric */
// An object representing metrics about disk usage.


// An object representing metrics about disk usage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiskIOMetric
type MXDiskIOMetric struct {
	MXMetric
}

// MXDiskIOMetricFrom constructs a [MXDiskIOMetric] from an unsafe.Pointer.
//
// An object representing metrics about disk usage.
func MXDiskIOMetricFrom(ptr unsafe.Pointer) MXDiskIOMetric {
	return MXDiskIOMetric{
		MXMetric: MXMetricFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MXDiskIOMetric *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MXDiskIOMetric */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MXDiskIOMetric */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MXDiskIOMetric */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MXDiskIOMetric */

// The total amount of data written to disk or other long term storage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiskIOMetric/cumulativeLogicalWrites
func (m_ MXDiskIOMetric) CumulativeLogicalWrites() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cumulativeLogicalWrites"))
	return rv
}/* debug [instance_properties/getter]: cumulativeLogicalWrites */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MXDiskIOMetric */



