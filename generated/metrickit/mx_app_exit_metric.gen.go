// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MXAppExitMetric */


/* debug [class_header]: Header for MXAppExitMetric */
// The class instance for the [MXAppExitMetric] class.
var (
	MXAppExitMetricClass     _MXAppExitMetricClass
	MXAppExitMetricClassOnce sync.Once
)

func getMXAppExitMetricClass() _MXAppExitMetricClass {
	MXAppExitMetricClassOnce.Do(func() {
		MXAppExitMetricClass = _MXAppExitMetricClass{objc.GetClass("MXAppExitMetric")}
	})
	return MXAppExitMetricClass
}

type _MXAppExitMetricClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MXAppExitMetric */
// An interface definition for the [MXAppExitMetric] class.
type IMXAppExitMetric interface {
	IMXMetric
	
/* debug [class_interface_properties]: Properties for MXAppExitMetric */
	// properties:
	BackgroundExitData() IMXBackgroundExitData
	ForegroundExitData() IMXForegroundExitData
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MXAppExitMetric */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MXAppExitMetric */
// Alloc allocates a new instance without initialization.
func (mc _MXAppExitMetricClass) Alloc() MXAppExitMetric {
	rv := objc.Send[MXAppExitMetric](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MXAppExitMetricClass) New() MXAppExitMetric {
	rv := objc.Send[MXAppExitMetric](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXAppExitMetric) Init() MXAppExitMetric {
	rv := objc.Send[MXAppExitMetric](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXAppExitMetric) Autorelease() MXAppExitMetric {
	rv := objc.Send[MXAppExitMetric](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXAppExitMetric creates a new MXAppExitMetric instance.
func NewMXAppExitMetric() MXAppExitMetric {
	return getMXAppExitMetricClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MXAppExitMetric */
// An object representing metrics about the types of foreground and background app exits.


// An object representing metrics about the types of foreground and background app exits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAppExitMetric
type MXAppExitMetric struct {
	MXMetric
}

// MXAppExitMetricFrom constructs a [MXAppExitMetric] from an unsafe.Pointer.
//
// An object representing metrics about the types of foreground and background app exits.
func MXAppExitMetricFrom(ptr unsafe.Pointer) MXAppExitMetric {
	return MXAppExitMetric{
		MXMetric: MXMetricFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MXAppExitMetric *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MXAppExitMetric */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MXAppExitMetric */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MXAppExitMetric */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MXAppExitMetric */

// The metrics for the background app exits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAppExitMetric/backgroundExitData
func (m_ MXAppExitMetric) BackgroundExitData() IMXBackgroundExitData {
	rv := objc.Send[MXBackgroundExitData](m_.ID, objc.Sel("backgroundExitData"))
	return rv
}/* debug [instance_properties/getter]: backgroundExitData */


// The metrics for the foreground app exits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAppExitMetric/foregroundExitData
func (m_ MXAppExitMetric) ForegroundExitData() IMXForegroundExitData {
	rv := objc.Send[MXForegroundExitData](m_.ID, objc.Sel("foregroundExitData"))
	return rv
}/* debug [instance_properties/getter]: foregroundExitData */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MXAppExitMetric */



