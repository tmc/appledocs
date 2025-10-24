// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MXLocationActivityMetric */


/* debug [class_header]: Header for MXLocationActivityMetric */
// The class instance for the [MXLocationActivityMetric] class.
var (
	MXLocationActivityMetricClass     _MXLocationActivityMetricClass
	MXLocationActivityMetricClassOnce sync.Once
)

func getMXLocationActivityMetricClass() _MXLocationActivityMetricClass {
	MXLocationActivityMetricClassOnce.Do(func() {
		MXLocationActivityMetricClass = _MXLocationActivityMetricClass{objc.GetClass("MXLocationActivityMetric")}
	})
	return MXLocationActivityMetricClass
}

type _MXLocationActivityMetricClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MXLocationActivityMetric */
// An interface definition for the [MXLocationActivityMetric] class.
type IMXLocationActivityMetric interface {
	IMXMetric
	
/* debug [class_interface_properties]: Properties for MXLocationActivityMetric */
	// properties:
	CumulativeBestAccuracyForNavigationTime() unsafe.Pointer
	CumulativeBestAccuracyTime() unsafe.Pointer
	CumulativeHundredMetersAccuracyTime() unsafe.Pointer
	CumulativeKilometerAccuracyTime() unsafe.Pointer
	CumulativeNearestTenMetersAccuracyTime() unsafe.Pointer
	CumulativeThreeKilometersAccuracyTime() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MXLocationActivityMetric */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MXLocationActivityMetric */
// Alloc allocates a new instance without initialization.
func (mc _MXLocationActivityMetricClass) Alloc() MXLocationActivityMetric {
	rv := objc.Send[MXLocationActivityMetric](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MXLocationActivityMetricClass) New() MXLocationActivityMetric {
	rv := objc.Send[MXLocationActivityMetric](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXLocationActivityMetric) Init() MXLocationActivityMetric {
	rv := objc.Send[MXLocationActivityMetric](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXLocationActivityMetric) Autorelease() MXLocationActivityMetric {
	rv := objc.Send[MXLocationActivityMetric](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXLocationActivityMetric creates a new MXLocationActivityMetric instance.
func NewMXLocationActivityMetric() MXLocationActivityMetric {
	return getMXLocationActivityMetricClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MXLocationActivityMetric */
// An object representing metrics about the use of location-tracking features of a device.


// An object representing metrics about the use of location-tracking features of a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXLocationActivityMetric
type MXLocationActivityMetric struct {
	MXMetric
}

// MXLocationActivityMetricFrom constructs a [MXLocationActivityMetric] from an unsafe.Pointer.
//
// An object representing metrics about the use of location-tracking features of a device.
func MXLocationActivityMetricFrom(ptr unsafe.Pointer) MXLocationActivityMetric {
	return MXLocationActivityMetric{
		MXMetric: MXMetricFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MXLocationActivityMetric *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MXLocationActivityMetric */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MXLocationActivityMetric */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MXLocationActivityMetric */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MXLocationActivityMetric */

// The total time spent tracking the current location at the best accuracy for navigation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXLocationActivityMetric/cumulativeBestAccuracyForNavigationTime
func (m_ MXLocationActivityMetric) CumulativeBestAccuracyForNavigationTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cumulativeBestAccuracyForNavigationTime"))
	return rv
}/* debug [instance_properties/getter]: cumulativeBestAccuracyForNavigationTime */


// The total time spent tracking the current location at the best accuracy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXLocationActivityMetric/cumulativeBestAccuracyTime
func (m_ MXLocationActivityMetric) CumulativeBestAccuracyTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cumulativeBestAccuracyTime"))
	return rv
}/* debug [instance_properties/getter]: cumulativeBestAccuracyTime */


// The total time spent tracking the current location to an accuracy of 100 meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXLocationActivityMetric/cumulativeHundredMetersAccuracyTime
func (m_ MXLocationActivityMetric) CumulativeHundredMetersAccuracyTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cumulativeHundredMetersAccuracyTime"))
	return rv
}/* debug [instance_properties/getter]: cumulativeHundredMetersAccuracyTime */


// The total time spent tracking the current location to an accuracy of 1 kilometer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXLocationActivityMetric/cumulativeKilometerAccuracyTime
func (m_ MXLocationActivityMetric) CumulativeKilometerAccuracyTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cumulativeKilometerAccuracyTime"))
	return rv
}/* debug [instance_properties/getter]: cumulativeKilometerAccuracyTime */


// The total time spent tracking the current location to an accuracy of 10 meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXLocationActivityMetric/cumulativeNearestTenMetersAccuracyTime
func (m_ MXLocationActivityMetric) CumulativeNearestTenMetersAccuracyTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cumulativeNearestTenMetersAccuracyTime"))
	return rv
}/* debug [instance_properties/getter]: cumulativeNearestTenMetersAccuracyTime */


// The total time spent tracking the current location to an accuracy of 3 kilometers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXLocationActivityMetric/cumulativeThreeKilometersAccuracyTime
func (m_ MXLocationActivityMetric) CumulativeThreeKilometersAccuracyTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cumulativeThreeKilometersAccuracyTime"))
	return rv
}/* debug [instance_properties/getter]: cumulativeThreeKilometersAccuracyTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MXLocationActivityMetric */



