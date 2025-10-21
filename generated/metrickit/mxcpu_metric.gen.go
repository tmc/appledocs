// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MXCPUMetric] class.
var (
	MXCPUMetricClass     _MXCPUMetricClass
	MXCPUMetricClassOnce sync.Once
)

func getMXCPUMetricClass() _MXCPUMetricClass {
	MXCPUMetricClassOnce.Do(func() {
		MXCPUMetricClass = _MXCPUMetricClass{objc.GetClass("MXCPUMetric")}
	})
	return MXCPUMetricClass
}

type _MXCPUMetricClass struct {
	class objc.Class
}

// An interface definition for the [MXCPUMetric] class.
type IMXCPUMetric interface {
	IMXMetric
}

// An object representing metrics about the use of the CPU.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCPUMetric
type MXCPUMetric struct {
	MXMetric
}

// MXCPUMetricFrom constructs a [MXCPUMetric] from an unsafe.Pointer.
//
// An object representing metrics about the use of the CPU.
func MXCPUMetricFrom(ptr unsafe.Pointer) MXCPUMetric {
	return MXCPUMetric{
		MXMetric: MXMetricFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MXCPUMetricClass) Alloc() MXCPUMetric {
	rv := objc.Send[MXCPUMetric](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MXCPUMetricClass) New() MXCPUMetric {
	rv := objc.Send[MXCPUMetric](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXCPUMetric) Init() MXCPUMetric {
	rv := objc.Send[MXCPUMetric](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXCPUMetric) Autorelease() MXCPUMetric {
	rv := objc.Send[MXCPUMetric](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXCPUMetric creates a new MXCPUMetric instance.
func NewMXCPUMetric() MXCPUMetric {
	return getMXCPUMetricClass().New()
}


// The total number of CPU instructions the app executed during the reporting period.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCPUMetric/cumulativeCPUInstructions
func (m_ MXCPUMetric) CumulativeCPUInstructions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cumulativeCPUInstructions"))
	return rv
}

// The total amount of CPU the app used.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCPUMetric/cumulativeCPUTime
func (m_ MXCPUMetric) CumulativeCPUTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cumulativeCPUTime"))
	return rv
}



