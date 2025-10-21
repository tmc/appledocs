// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MXGPUMetric] class.
var (
	MXGPUMetricClass     _MXGPUMetricClass
	MXGPUMetricClassOnce sync.Once
)

func getMXGPUMetricClass() _MXGPUMetricClass {
	MXGPUMetricClassOnce.Do(func() {
		MXGPUMetricClass = _MXGPUMetricClass{objc.GetClass("MXGPUMetric")}
	})
	return MXGPUMetricClass
}

type _MXGPUMetricClass struct {
	class objc.Class
}

// An interface definition for the [MXGPUMetric] class.
type IMXGPUMetric interface {
	IMXMetric
}

// An object representing metrics about the use of the GPU.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXGPUMetric
type MXGPUMetric struct {
	MXMetric
}

// MXGPUMetricFrom constructs a [MXGPUMetric] from an unsafe.Pointer.
//
// An object representing metrics about the use of the GPU.
func MXGPUMetricFrom(ptr unsafe.Pointer) MXGPUMetric {
	return MXGPUMetric{
		MXMetric: MXMetricFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MXGPUMetricClass) Alloc() MXGPUMetric {
	rv := objc.Send[MXGPUMetric](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MXGPUMetricClass) New() MXGPUMetric {
	rv := objc.Send[MXGPUMetric](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXGPUMetric) Init() MXGPUMetric {
	rv := objc.Send[MXGPUMetric](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXGPUMetric) Autorelease() MXGPUMetric {
	rv := objc.Send[MXGPUMetric](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXGPUMetric creates a new MXGPUMetric instance.
func NewMXGPUMetric() MXGPUMetric {
	return getMXGPUMetricClass().New()
}


// The total amount of GPU time used by the app.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXGPUMetric/cumulativeGPUTime
func (m_ MXGPUMetric) CumulativeGPUTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cumulativeGPUTime"))
	return rv
}



