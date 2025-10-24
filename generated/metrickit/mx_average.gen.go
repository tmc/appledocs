// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MXAverage */


/* debug [class_header]: Header for MXAverage */
// The class instance for the [MXAverage] class.
var (
	MXAverageClass     _MXAverageClass
	MXAverageClassOnce sync.Once
)

func getMXAverageClass() _MXAverageClass {
	MXAverageClassOnce.Do(func() {
		MXAverageClass = _MXAverageClass{objc.GetClass("MXAverage")}
	})
	return MXAverageClass
}

type _MXAverageClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MXAverage */
// An interface definition for the [MXAverage] class.
type IMXAverage interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MXAverage */
	// properties:
	AverageMeasurement() unsafe.Pointer
	SampleCount() int
	StandardDeviation() float64
	MXErrorDomain() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MXAverage */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MXAverage */
// Alloc allocates a new instance without initialization.
func (mc _MXAverageClass) Alloc() MXAverage {
	rv := objc.Send[MXAverage](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MXAverageClass) New() MXAverage {
	rv := objc.Send[MXAverage](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXAverage) Init() MXAverage {
	rv := objc.Send[MXAverage](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXAverage) Autorelease() MXAverage {
	rv := objc.Send[MXAverage](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXAverage creates a new MXAverage instance.
func NewMXAverage() MXAverage {
	return getMXAverageClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MXAverage */
// A unit of measure for an average.


// A unit of measure for an average.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAverage
type MXAverage struct {
	objectivec.Object
}

// MXAverageFrom constructs a [MXAverage] from an unsafe.Pointer.
//
// A unit of measure for an average.
func MXAverageFrom(ptr unsafe.Pointer) MXAverage {
	return MXAverage{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MXAverage *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MXAverage */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MXAverage */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MXAverage */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MXAverage */

// The value of the average.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAverage/averageMeasurement
func (m_ MXAverage) AverageMeasurement() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("averageMeasurement"))
	return rv
}/* debug [instance_properties/getter]: averageMeasurement */


// The number of samples used to calculate the average.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAverage/sampleCount
func (m_ MXAverage) SampleCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("sampleCount"))
	return rv
}/* debug [instance_properties/getter]: sampleCount */


// The standard deviation of the distribution of values used to calculate the average.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAverage/standardDeviation
func (m_ MXAverage) StandardDeviation() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("standardDeviation"))
	return rv
}/* debug [instance_properties/getter]: standardDeviation */


// Error domain for error values from app metrics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxerrordomain
func (m_ MXAverage) MXErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MXErrorDomain"))
	return rv
}/* debug [instance_properties/getter]: MXErrorDomain */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MXAverage */



