// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MXDiskIOMetric] class.
var (
	MXDiskIOMetricClass     _MXDiskIOMetricClass
	MXDiskIOMetricClassOnce sync.Once
)

func getMXDiskIOMetricClass() _MXDiskIOMetricClass {
	MXDiskIOMetricClassOnce.Do(func() {
		MXDiskIOMetricClass = _MXDiskIOMetricClass{objc.GetClass("MXDiskIOMetric")}
	})
	return MXDiskIOMetricClass
}

type _MXDiskIOMetricClass struct {
	class objc.Class
}

// An interface definition for the [MXDiskIOMetric] class.
type IMXDiskIOMetric interface {
	IMXMetric
}

// An object representing metrics about disk usage.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiskIOMetric
type MXDiskIOMetric struct {
	MXMetric
}

// MXDiskIOMetricFrom constructs a [MXDiskIOMetric] from an unsafe.Pointer.
//
// An object representing metrics about disk usage.
func MXDiskIOMetricFrom(ptr unsafe.Pointer) MXDiskIOMetric {
	return MXDiskIOMetric{
		MXMetric: MXMetricFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MXDiskIOMetricClass) Alloc() MXDiskIOMetric {
	rv := objc.Send[MXDiskIOMetric](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MXDiskIOMetricClass) New() MXDiskIOMetric {
	rv := objc.Send[MXDiskIOMetric](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXDiskIOMetric) Init() MXDiskIOMetric {
	rv := objc.Send[MXDiskIOMetric](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXDiskIOMetric) Autorelease() MXDiskIOMetric {
	rv := objc.Send[MXDiskIOMetric](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXDiskIOMetric creates a new MXDiskIOMetric instance.
func NewMXDiskIOMetric() MXDiskIOMetric {
	return getMXDiskIOMetricClass().New()
}


// The total amount of data written to disk or other long term storage.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiskIOMetric/cumulativeLogicalWrites
func (m_ MXDiskIOMetric) CumulativeLogicalWrites() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cumulativeLogicalWrites"))
	return rv
}



