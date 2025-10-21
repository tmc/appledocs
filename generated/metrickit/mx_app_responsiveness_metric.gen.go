// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MXAppResponsivenessMetric] class.
var (
	MXAppResponsivenessMetricClass     _MXAppResponsivenessMetricClass
	MXAppResponsivenessMetricClassOnce sync.Once
)

func getMXAppResponsivenessMetricClass() _MXAppResponsivenessMetricClass {
	MXAppResponsivenessMetricClassOnce.Do(func() {
		MXAppResponsivenessMetricClass = _MXAppResponsivenessMetricClass{objc.GetClass("MXAppResponsivenessMetric")}
	})
	return MXAppResponsivenessMetricClass
}

type _MXAppResponsivenessMetricClass struct {
	class objc.Class
}

// An interface definition for the [MXAppResponsivenessMetric] class.
type IMXAppResponsivenessMetric interface {
	IMXMetric
}

// An object representing metrics about the responsiveness of the app to user interaction.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAppResponsivenessMetric
type MXAppResponsivenessMetric struct {
	MXMetric
}

// MXAppResponsivenessMetricFrom constructs a [MXAppResponsivenessMetric] from an unsafe.Pointer.
//
// An object representing metrics about the responsiveness of the app to user interaction.
func MXAppResponsivenessMetricFrom(ptr unsafe.Pointer) MXAppResponsivenessMetric {
	return MXAppResponsivenessMetric{
		MXMetric: MXMetricFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MXAppResponsivenessMetricClass) Alloc() MXAppResponsivenessMetric {
	rv := objc.Send[MXAppResponsivenessMetric](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MXAppResponsivenessMetricClass) New() MXAppResponsivenessMetric {
	rv := objc.Send[MXAppResponsivenessMetric](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXAppResponsivenessMetric) Init() MXAppResponsivenessMetric {
	rv := objc.Send[MXAppResponsivenessMetric](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXAppResponsivenessMetric) Autorelease() MXAppResponsivenessMetric {
	rv := objc.Send[MXAppResponsivenessMetric](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXAppResponsivenessMetric creates a new MXAppResponsivenessMetric instance.
func NewMXAppResponsivenessMetric() MXAppResponsivenessMetric {
	return getMXAppResponsivenessMetricClass().New()
}


// A histogram of the different durations of time in which the app is too busy to handle user interaction responsively.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAppResponsivenessMetric/histogrammedApplicationHangTime
func (m_ MXAppResponsivenessMetric) HistogrammedApplicationHangTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("histogrammedApplicationHangTime"))
	return rv
}



