// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MXAppExitMetric] class.
var (
	MXAppExitMetricClass     _MXAppExitMetricClass
	MXAppExitMetricClassOnce sync.Once
)

func getMXAppExitMetricClass() _MXAppExitMetricClass {
	MXAppExitMetricClassOnce.Do(func() {
		MXAppExitMetricClass = _MXAppExitMetricClass{objc.GetClass("MXAppExitMetric")}
	})
	return MXAppExitMetricClass
}

type _MXAppExitMetricClass struct {
	class objc.Class
}

// An interface definition for the [MXAppExitMetric] class.
type IMXAppExitMetric interface {
	IMXMetric
	// properties:
	BackgroundExitData() IMXBackgroundExitData
	ForegroundExitData() IMXForegroundExitData
	// methods:
}

// An object representing metrics about the types of foreground and background app exits.


// An object representing metrics about the types of foreground and background app exits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAppExitMetric
type MXAppExitMetric struct {
	MXMetric
}

// MXAppExitMetricFrom constructs a [MXAppExitMetric] from an unsafe.Pointer.
//
// An object representing metrics about the types of foreground and background app exits.
func MXAppExitMetricFrom(ptr unsafe.Pointer) MXAppExitMetric {
	return MXAppExitMetric{
		MXMetric: MXMetricFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MXAppExitMetricClass) Alloc() MXAppExitMetric {
	rv := objc.Send[MXAppExitMetric](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MXAppExitMetricClass) New() MXAppExitMetric {
	rv := objc.Send[MXAppExitMetric](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXAppExitMetric) Init() MXAppExitMetric {
	rv := objc.Send[MXAppExitMetric](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXAppExitMetric) Autorelease() MXAppExitMetric {
	rv := objc.Send[MXAppExitMetric](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXAppExitMetric creates a new MXAppExitMetric instance.
func NewMXAppExitMetric() MXAppExitMetric {
	return getMXAppExitMetricClass().New()
}



// The metrics for the background app exits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAppExitMetric/backgroundExitData
func (m_ MXAppExitMetric) BackgroundExitData() IMXBackgroundExitData {
	rv := objc.Send[MXBackgroundExitData](m_.ID, objc.Sel("backgroundExitData"))
	return rv
}


// The metrics for the foreground app exits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAppExitMetric/foregroundExitData
func (m_ MXAppExitMetric) ForegroundExitData() IMXForegroundExitData {
	rv := objc.Send[MXForegroundExitData](m_.ID, objc.Sel("foregroundExitData"))
	return rv
}



