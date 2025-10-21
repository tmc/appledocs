// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [MXAnimationMetric] class.
type IMXAnimationMetric interface {
	IMXMetric
}

// An object representing metrics about the responsiveness of animation in the app.
//
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

// Alloc allocates a new instance without initialization.
func (mc _MXAnimationMetricClass) Alloc() MXAnimationMetric {
	rv := objc.Send[MXAnimationMetric](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The ratio of time spent hitching during tracked animations.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAnimationMetric/hitchTimeRatio
func (m_ MXAnimationMetric) HitchTimeRatio() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("hitchTimeRatio"))
	return rv
}

// The ratio of the time spent hitching while scrolling.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAnimationMetric/scrollHitchTimeRatio
func (m_ MXAnimationMetric) ScrollHitchTimeRatio() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("scrollHitchTimeRatio"))
	return rv
}



