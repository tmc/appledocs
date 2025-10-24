// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MXAppRunTimeMetric */


/* debug [class_header]: Header for MXAppRunTimeMetric */
// The class instance for the [MXAppRunTimeMetric] class.
var (
	MXAppRunTimeMetricClass     _MXAppRunTimeMetricClass
	MXAppRunTimeMetricClassOnce sync.Once
)

func getMXAppRunTimeMetricClass() _MXAppRunTimeMetricClass {
	MXAppRunTimeMetricClassOnce.Do(func() {
		MXAppRunTimeMetricClass = _MXAppRunTimeMetricClass{objc.GetClass("MXAppRunTimeMetric")}
	})
	return MXAppRunTimeMetricClass
}

type _MXAppRunTimeMetricClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MXAppRunTimeMetric */
// An interface definition for the [MXAppRunTimeMetric] class.
type IMXAppRunTimeMetric interface {
	IMXMetric
	
/* debug [class_interface_properties]: Properties for MXAppRunTimeMetric */
	// properties:
	CumulativeBackgroundAudioTime() unsafe.Pointer
	CumulativeBackgroundLocationTime() unsafe.Pointer
	CumulativeBackgroundTime() unsafe.Pointer
	CumulativeForegroundTime() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MXAppRunTimeMetric */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MXAppRunTimeMetric */
// Alloc allocates a new instance without initialization.
func (mc _MXAppRunTimeMetricClass) Alloc() MXAppRunTimeMetric {
	rv := objc.Send[MXAppRunTimeMetric](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MXAppRunTimeMetricClass) New() MXAppRunTimeMetric {
	rv := objc.Send[MXAppRunTimeMetric](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXAppRunTimeMetric) Init() MXAppRunTimeMetric {
	rv := objc.Send[MXAppRunTimeMetric](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXAppRunTimeMetric) Autorelease() MXAppRunTimeMetric {
	rv := objc.Send[MXAppRunTimeMetric](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXAppRunTimeMetric creates a new MXAppRunTimeMetric instance.
func NewMXAppRunTimeMetric() MXAppRunTimeMetric {
	return getMXAppRunTimeMetricClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MXAppRunTimeMetric */
// An object representing metrics about the amount of time the app is active.


// An object representing metrics about the amount of time the app is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAppRunTimeMetric
type MXAppRunTimeMetric struct {
	MXMetric
}

// MXAppRunTimeMetricFrom constructs a [MXAppRunTimeMetric] from an unsafe.Pointer.
//
// An object representing metrics about the amount of time the app is active.
func MXAppRunTimeMetricFrom(ptr unsafe.Pointer) MXAppRunTimeMetric {
	return MXAppRunTimeMetric{
		MXMetric: MXMetricFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MXAppRunTimeMetric *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MXAppRunTimeMetric */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MXAppRunTimeMetric */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MXAppRunTimeMetric */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MXAppRunTimeMetric */

// The total time the app is in the background and playing audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAppRunTimeMetric/cumulativeBackgroundAudioTime
func (m_ MXAppRunTimeMetric) CumulativeBackgroundAudioTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cumulativeBackgroundAudioTime"))
	return rv
}/* debug [instance_properties/getter]: cumulativeBackgroundAudioTime */


// The total time the app is in the background and using location services.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAppRunTimeMetric/cumulativeBackgroundLocationTime
func (m_ MXAppRunTimeMetric) CumulativeBackgroundLocationTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cumulativeBackgroundLocationTime"))
	return rv
}/* debug [instance_properties/getter]: cumulativeBackgroundLocationTime */


// The total time the app is active in the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAppRunTimeMetric/cumulativeBackgroundTime
func (m_ MXAppRunTimeMetric) CumulativeBackgroundTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cumulativeBackgroundTime"))
	return rv
}/* debug [instance_properties/getter]: cumulativeBackgroundTime */


// The total time the app is in the foreground.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAppRunTimeMetric/cumulativeForegroundTime
func (m_ MXAppRunTimeMetric) CumulativeForegroundTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cumulativeForegroundTime"))
	return rv
}/* debug [instance_properties/getter]: cumulativeForegroundTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MXAppRunTimeMetric */



