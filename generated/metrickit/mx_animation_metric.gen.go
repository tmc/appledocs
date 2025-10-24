// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MXAnimationMetric */


/* debug [class_header]: Header for MXAnimationMetric */
// The class instance for the [MXAnimationMetric] class.
var (
	MXAnimationMetricClass     _MXAnimationMetricClass
	MXAnimationMetricClassOnce sync.Once
)

func getMXAnimationMetricClass() _MXAnimationMetricClass {
	MXAnimationMetricClassOnce.Do(func() {
		MXAnimationMetricClass = _MXAnimationMetricClass{objc.GetClass("MXAnimationMetric")}
	})
	return MXAnimationMetricClass
}

type _MXAnimationMetricClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MXAnimationMetric */
// An interface definition for the [MXAnimationMetric] class.
type IMXAnimationMetric interface {
	IMXMetric
	
/* debug [class_interface_properties]: Properties for MXAnimationMetric */
	// properties:
	HitchTimeRatio() unsafe.Pointer
	ScrollHitchTimeRatio() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MXAnimationMetric */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MXAnimationMetric */
// Alloc allocates a new instance without initialization.
func (mc _MXAnimationMetricClass) Alloc() MXAnimationMetric {
	rv := objc.Send[MXAnimationMetric](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MXAnimationMetricClass) New() MXAnimationMetric {
	rv := objc.Send[MXAnimationMetric](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXAnimationMetric) Init() MXAnimationMetric {
	rv := objc.Send[MXAnimationMetric](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXAnimationMetric) Autorelease() MXAnimationMetric {
	rv := objc.Send[MXAnimationMetric](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXAnimationMetric creates a new MXAnimationMetric instance.
func NewMXAnimationMetric() MXAnimationMetric {
	return getMXAnimationMetricClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MXAnimationMetric */
// An object representing metrics about the responsiveness of animation in the app.


// An object representing metrics about the responsiveness of animation in the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAnimationMetric
type MXAnimationMetric struct {
	MXMetric
}

// MXAnimationMetricFrom constructs a [MXAnimationMetric] from an unsafe.Pointer.
//
// An object representing metrics about the responsiveness of animation in the app.
func MXAnimationMetricFrom(ptr unsafe.Pointer) MXAnimationMetric {
	return MXAnimationMetric{
		MXMetric: MXMetricFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MXAnimationMetric *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MXAnimationMetric */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MXAnimationMetric */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MXAnimationMetric */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MXAnimationMetric */

// The ratio of time spent hitching during tracked animations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAnimationMetric/hitchTimeRatio
func (m_ MXAnimationMetric) HitchTimeRatio() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("hitchTimeRatio"))
	return rv
}/* debug [instance_properties/getter]: hitchTimeRatio */


// The ratio of the time spent hitching while scrolling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAnimationMetric/scrollHitchTimeRatio
func (m_ MXAnimationMetric) ScrollHitchTimeRatio() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("scrollHitchTimeRatio"))
	return rv
}/* debug [instance_properties/getter]: scrollHitchTimeRatio */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MXAnimationMetric */



