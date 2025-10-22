// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MXAppRunTimeMetric] class.
var (
	MXAppRunTimeMetricClass     _MXAppRunTimeMetricClass
	MXAppRunTimeMetricClassOnce sync.Once
)

func getMXAppRunTimeMetricClass() _MXAppRunTimeMetricClass {
	MXAppRunTimeMetricClassOnce.Do(func() {
		MXAppRunTimeMetricClass = _MXAppRunTimeMetricClass{objc.GetClass("MXAppRunTimeMetric")}
	})
	return MXAppRunTimeMetricClass
}

type _MXAppRunTimeMetricClass struct {
	class objc.Class
}

// An interface definition for the [MXAppRunTimeMetric] class.
type IMXAppRunTimeMetric interface {
	IMXMetric
	CumulativeBackgroundAudioTime() unsafe.Pointer
	CumulativeBackgroundLocationTime() unsafe.Pointer
	CumulativeBackgroundTime() unsafe.Pointer
	CumulativeForegroundTime() unsafe.Pointer
}

// An object representing metrics about the amount of time the app is active.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAppRunTimeMetric
type MXAppRunTimeMetric struct {
	MXMetric
}

// MXAppRunTimeMetricFrom constructs a [MXAppRunTimeMetric] from an unsafe.Pointer.
//
// An object representing metrics about the amount of time the app is active.
func MXAppRunTimeMetricFrom(ptr unsafe.Pointer) MXAppRunTimeMetric {
	return MXAppRunTimeMetric{
		MXMetric: MXMetricFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MXAppRunTimeMetricClass) Alloc() MXAppRunTimeMetric {
	rv := objc.Send[MXAppRunTimeMetric](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MXAppRunTimeMetricClass) New() MXAppRunTimeMetric {
	rv := objc.Send[MXAppRunTimeMetric](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXAppRunTimeMetric) Init() MXAppRunTimeMetric {
	rv := objc.Send[MXAppRunTimeMetric](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXAppRunTimeMetric) Autorelease() MXAppRunTimeMetric {
	rv := objc.Send[MXAppRunTimeMetric](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXAppRunTimeMetric creates a new MXAppRunTimeMetric instance.
func NewMXAppRunTimeMetric() MXAppRunTimeMetric {
	return getMXAppRunTimeMetricClass().New()
}


// The total time the app is in the background and playing audio.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAppRunTimeMetric/cumulativeBackgroundAudioTime
func (m_ MXAppRunTimeMetric) CumulativeBackgroundAudioTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cumulativeBackgroundAudioTime"))
	return rv
}

// The total time the app is in the background and using location services.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAppRunTimeMetric/cumulativeBackgroundLocationTime
func (m_ MXAppRunTimeMetric) CumulativeBackgroundLocationTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cumulativeBackgroundLocationTime"))
	return rv
}

// The total time the app is active in the background.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAppRunTimeMetric/cumulativeBackgroundTime
func (m_ MXAppRunTimeMetric) CumulativeBackgroundTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cumulativeBackgroundTime"))
	return rv
}

// The total time the app is in the foreground.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAppRunTimeMetric/cumulativeForegroundTime
func (m_ MXAppRunTimeMetric) CumulativeForegroundTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cumulativeForegroundTime"))
	return rv
}



