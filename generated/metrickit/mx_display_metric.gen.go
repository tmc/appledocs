// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MXDisplayMetric */


/* debug [class_header]: Header for MXDisplayMetric */
// The class instance for the [MXDisplayMetric] class.
var (
	MXDisplayMetricClass     _MXDisplayMetricClass
	MXDisplayMetricClassOnce sync.Once
)

func getMXDisplayMetricClass() _MXDisplayMetricClass {
	MXDisplayMetricClassOnce.Do(func() {
		MXDisplayMetricClass = _MXDisplayMetricClass{objc.GetClass("MXDisplayMetric")}
	})
	return MXDisplayMetricClass
}

type _MXDisplayMetricClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MXDisplayMetric */
// An interface definition for the [MXDisplayMetric] class.
type IMXDisplayMetric interface {
	IMXMetric
	
/* debug [class_interface_properties]: Properties for MXDisplayMetric */
	// properties:
	AveragePixelLuminance() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MXDisplayMetric */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MXDisplayMetric */
// Alloc allocates a new instance without initialization.
func (mc _MXDisplayMetricClass) Alloc() MXDisplayMetric {
	rv := objc.Send[MXDisplayMetric](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MXDisplayMetricClass) New() MXDisplayMetric {
	rv := objc.Send[MXDisplayMetric](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXDisplayMetric) Init() MXDisplayMetric {
	rv := objc.Send[MXDisplayMetric](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXDisplayMetric) Autorelease() MXDisplayMetric {
	rv := objc.Send[MXDisplayMetric](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXDisplayMetric creates a new MXDisplayMetric instance.
func NewMXDisplayMetric() MXDisplayMetric {
	return getMXDisplayMetricClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MXDisplayMetric */
// An object representing metrics about the power used to display the app on the screen.


// An object representing metrics about the power used to display the app on the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDisplayMetric
type MXDisplayMetric struct {
	MXMetric
}

// MXDisplayMetricFrom constructs a [MXDisplayMetric] from an unsafe.Pointer.
//
// An object representing metrics about the power used to display the app on the screen.
func MXDisplayMetricFrom(ptr unsafe.Pointer) MXDisplayMetric {
	return MXDisplayMetric{
		MXMetric: MXMetricFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MXDisplayMetric *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MXDisplayMetric */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MXDisplayMetric */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MXDisplayMetric */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MXDisplayMetric */

// The average amount of luminosity of the pixels on an OLED display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDisplayMetric/averagePixelLuminance
func (m_ MXDisplayMetric) AveragePixelLuminance() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("averagePixelLuminance"))
	return rv
}/* debug [instance_properties/getter]: averagePixelLuminance */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MXDisplayMetric */



