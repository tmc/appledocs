// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MXCellularConditionMetric] class.
var (
	MXCellularConditionMetricClass     _MXCellularConditionMetricClass
	MXCellularConditionMetricClassOnce sync.Once
)

func getMXCellularConditionMetricClass() _MXCellularConditionMetricClass {
	MXCellularConditionMetricClassOnce.Do(func() {
		MXCellularConditionMetricClass = _MXCellularConditionMetricClass{objc.GetClass("MXCellularConditionMetric")}
	})
	return MXCellularConditionMetricClass
}

type _MXCellularConditionMetricClass struct {
	class objc.Class
}

// An interface definition for the [MXCellularConditionMetric] class.
type IMXCellularConditionMetric interface {
	IMXMetric
}

// An object representing metrics about the condition of the cellular network.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCellularConditionMetric
type MXCellularConditionMetric struct {
	MXMetric
}

// MXCellularConditionMetricFrom constructs a [MXCellularConditionMetric] from an unsafe.Pointer.
//
// An object representing metrics about the condition of the cellular network.
func MXCellularConditionMetricFrom(ptr unsafe.Pointer) MXCellularConditionMetric {
	return MXCellularConditionMetric{
		MXMetric: MXMetricFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MXCellularConditionMetricClass) Alloc() MXCellularConditionMetric {
	rv := objc.Send[MXCellularConditionMetric](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MXCellularConditionMetricClass) New() MXCellularConditionMetric {
	rv := objc.Send[MXCellularConditionMetric](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXCellularConditionMetric) Init() MXCellularConditionMetric {
	rv := objc.Send[MXCellularConditionMetric](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXCellularConditionMetric) Autorelease() MXCellularConditionMetric {
	rv := objc.Send[MXCellularConditionMetric](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXCellularConditionMetric creates a new MXCellularConditionMetric instance.
func NewMXCellularConditionMetric() MXCellularConditionMetric {
	return getMXCellularConditionMetricClass().New()
}


// An object representing the distribution of the different levels of connectivity to the cellular network.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCellularConditionMetric/histogrammedCellularConditionTime
func (m_ MXCellularConditionMetric) HistogrammedCellularConditionTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("histogrammedCellularConditionTime"))
	return rv
}



