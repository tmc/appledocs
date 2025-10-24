// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MXAppLaunchMetric */


/* debug [class_header]: Header for MXAppLaunchMetric */
// The class instance for the [MXAppLaunchMetric] class.
var (
	MXAppLaunchMetricClass     _MXAppLaunchMetricClass
	MXAppLaunchMetricClassOnce sync.Once
)

func getMXAppLaunchMetricClass() _MXAppLaunchMetricClass {
	MXAppLaunchMetricClassOnce.Do(func() {
		MXAppLaunchMetricClass = _MXAppLaunchMetricClass{objc.GetClass("MXAppLaunchMetric")}
	})
	return MXAppLaunchMetricClass
}

type _MXAppLaunchMetricClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MXAppLaunchMetric */
// An interface definition for the [MXAppLaunchMetric] class.
type IMXAppLaunchMetric interface {
	IMXMetric
	
/* debug [class_interface_properties]: Properties for MXAppLaunchMetric */
	// properties:
	HistogrammedApplicationResumeTime() unsafe.Pointer
	HistogrammedExtendedLaunch() unsafe.Pointer
	HistogrammedOptimizedTimeToFirstDraw() unsafe.Pointer
	HistogrammedTimeToFirstDraw() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MXAppLaunchMetric */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MXAppLaunchMetric */
// Alloc allocates a new instance without initialization.
func (mc _MXAppLaunchMetricClass) Alloc() MXAppLaunchMetric {
	rv := objc.Send[MXAppLaunchMetric](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MXAppLaunchMetricClass) New() MXAppLaunchMetric {
	rv := objc.Send[MXAppLaunchMetric](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXAppLaunchMetric) Init() MXAppLaunchMetric {
	rv := objc.Send[MXAppLaunchMetric](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXAppLaunchMetric) Autorelease() MXAppLaunchMetric {
	rv := objc.Send[MXAppLaunchMetric](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXAppLaunchMetric creates a new MXAppLaunchMetric instance.
func NewMXAppLaunchMetric() MXAppLaunchMetric {
	return getMXAppLaunchMetricClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MXAppLaunchMetric */
// An object representing metrics about app launch time.


// An object representing metrics about app launch time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAppLaunchMetric
type MXAppLaunchMetric struct {
	MXMetric
}

// MXAppLaunchMetricFrom constructs a [MXAppLaunchMetric] from an unsafe.Pointer.
//
// An object representing metrics about app launch time.
func MXAppLaunchMetricFrom(ptr unsafe.Pointer) MXAppLaunchMetric {
	return MXAppLaunchMetric{
		MXMetric: MXMetricFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MXAppLaunchMetric *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MXAppLaunchMetric */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MXAppLaunchMetric */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MXAppLaunchMetric */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MXAppLaunchMetric */

// A histogram of the different amounts of time taken to resume the app from the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAppLaunchMetric/histogrammedApplicationResumeTime
func (m_ MXAppLaunchMetric) HistogrammedApplicationResumeTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("histogrammedApplicationResumeTime"))
	return rv
}/* debug [instance_properties/getter]: histogrammedApplicationResumeTime */


// A histogram of the different amounts of time taken to launch the app, including the extended launch tasks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAppLaunchMetric/histogrammedExtendedLaunch
func (m_ MXAppLaunchMetric) HistogrammedExtendedLaunch() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("histogrammedExtendedLaunch"))
	return rv
}/* debug [instance_properties/getter]: histogrammedExtendedLaunch */


// A histogram of the different amounts of time associated with prewarmed app launches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAppLaunchMetric/histogrammedOptimizedTimeToFirstDraw
func (m_ MXAppLaunchMetric) HistogrammedOptimizedTimeToFirstDraw() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("histogrammedOptimizedTimeToFirstDraw"))
	return rv
}/* debug [instance_properties/getter]: histogrammedOptimizedTimeToFirstDraw */


// A histogram of the different amounts of time taken to launch the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAppLaunchMetric/histogrammedTimeToFirstDraw
func (m_ MXAppLaunchMetric) HistogrammedTimeToFirstDraw() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("histogrammedTimeToFirstDraw"))
	return rv
}/* debug [instance_properties/getter]: histogrammedTimeToFirstDraw */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MXAppLaunchMetric */



