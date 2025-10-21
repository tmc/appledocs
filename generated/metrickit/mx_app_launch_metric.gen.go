// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MXAppLaunchMetric] class.
var (
	MXAppLaunchMetricClass     _MXAppLaunchMetricClass
	MXAppLaunchMetricClassOnce sync.Once
)

func getMXAppLaunchMetricClass() _MXAppLaunchMetricClass {
	MXAppLaunchMetricClassOnce.Do(func() {
		MXAppLaunchMetricClass = _MXAppLaunchMetricClass{objc.GetClass("MXAppLaunchMetric")}
	})
	return MXAppLaunchMetricClass
}

type _MXAppLaunchMetricClass struct {
	class objc.Class
}

// An interface definition for the [MXAppLaunchMetric] class.
type IMXAppLaunchMetric interface {
	IMXMetric
}

// An object representing metrics about app launch time.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAppLaunchMetric
type MXAppLaunchMetric struct {
	MXMetric
}

// MXAppLaunchMetricFrom constructs a [MXAppLaunchMetric] from an unsafe.Pointer.
//
// An object representing metrics about app launch time.
func MXAppLaunchMetricFrom(ptr unsafe.Pointer) MXAppLaunchMetric {
	return MXAppLaunchMetric{
		MXMetric: MXMetricFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MXAppLaunchMetricClass) Alloc() MXAppLaunchMetric {
	rv := objc.Send[MXAppLaunchMetric](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MXAppLaunchMetricClass) New() MXAppLaunchMetric {
	rv := objc.Send[MXAppLaunchMetric](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXAppLaunchMetric) Init() MXAppLaunchMetric {
	rv := objc.Send[MXAppLaunchMetric](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXAppLaunchMetric) Autorelease() MXAppLaunchMetric {
	rv := objc.Send[MXAppLaunchMetric](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXAppLaunchMetric creates a new MXAppLaunchMetric instance.
func NewMXAppLaunchMetric() MXAppLaunchMetric {
	return getMXAppLaunchMetricClass().New()
}


// A histogram of the different amounts of time taken to resume the app from the background.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAppLaunchMetric/histogrammedApplicationResumeTime
func (m_ MXAppLaunchMetric) HistogrammedApplicationResumeTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("histogrammedApplicationResumeTime"))
	return rv
}

// A histogram of the different amounts of time taken to launch the app, including the extended launch tasks.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAppLaunchMetric/histogrammedExtendedLaunch
func (m_ MXAppLaunchMetric) HistogrammedExtendedLaunch() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("histogrammedExtendedLaunch"))
	return rv
}

// A histogram of the different amounts of time associated with prewarmed app launches.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAppLaunchMetric/histogrammedOptimizedTimeToFirstDraw
func (m_ MXAppLaunchMetric) HistogrammedOptimizedTimeToFirstDraw() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("histogrammedOptimizedTimeToFirstDraw"))
	return rv
}

// A histogram of the different amounts of time taken to launch the app.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAppLaunchMetric/histogrammedTimeToFirstDraw
func (m_ MXAppLaunchMetric) HistogrammedTimeToFirstDraw() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("histogrammedTimeToFirstDraw"))
	return rv
}



