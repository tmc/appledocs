// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MXAverage] class.
type IMXAverage interface {
	objectivec.IObject
	AverageMeasurement() unsafe.Pointer
	SampleCount() int
	StandardDeviation() float64
	MXErrorDomain() string
}

// A unit of measure for an average.
//
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

// Alloc allocates a new instance without initialization.
func (mc _MXAverageClass) Alloc() MXAverage {
	rv := objc.Send[MXAverage](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The value of the average.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAverage/averageMeasurement
func (m_ MXAverage) AverageMeasurement() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("averageMeasurement"))
	return rv
}

// The number of samples used to calculate the average.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAverage/sampleCount
func (m_ MXAverage) SampleCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("sampleCount"))
	return rv
}

// The standard deviation of the distribution of values used to calculate the average.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAverage/standardDeviation
func (m_ MXAverage) StandardDeviation() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("standardDeviation"))
	return rv
}

// Error domain for error values from app metrics.
//
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxerrordomain
func (m_ MXAverage) MXErrorDomain() string {
	rv := objc.Send[string](m_.ID, objc.Sel("MXErrorDomain"))
	return rv
}



