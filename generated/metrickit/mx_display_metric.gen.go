// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [MXDisplayMetric] class.
type IMXDisplayMetric interface {
	IMXMetric
	AveragePixelLuminance() unsafe.Pointer
}

// An object representing metrics about the power used to display the app on the screen.
//
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

// Alloc allocates a new instance without initialization.
func (mc _MXDisplayMetricClass) Alloc() MXDisplayMetric {
	rv := objc.Send[MXDisplayMetric](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The average amount of luminosity of the pixels on an OLED display.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDisplayMetric/averagePixelLuminance
func (m_ MXDisplayMetric) AveragePixelLuminance() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("averagePixelLuminance"))
	return rv
}



