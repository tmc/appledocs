// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MXSignpostIntervalData] class.
var (
	MXSignpostIntervalDataClass     _MXSignpostIntervalDataClass
	MXSignpostIntervalDataClassOnce sync.Once
)

func getMXSignpostIntervalDataClass() _MXSignpostIntervalDataClass {
	MXSignpostIntervalDataClassOnce.Do(func() {
		MXSignpostIntervalDataClass = _MXSignpostIntervalDataClass{objc.GetClass("MXSignpostIntervalData")}
	})
	return MXSignpostIntervalDataClass
}

type _MXSignpostIntervalDataClass struct {
	class objc.Class
}

// An interface definition for the [MXSignpostIntervalData] class.
type IMXSignpostIntervalData interface {
	objectivec.IObject
}

// A data object representing the captured data for a custom metric.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXSignpostIntervalData
type MXSignpostIntervalData struct {
	objectivec.Object
}

// MXSignpostIntervalDataFrom constructs a [MXSignpostIntervalData] from an unsafe.Pointer.
//
// A data object representing the captured data for a custom metric.
func MXSignpostIntervalDataFrom(ptr unsafe.Pointer) MXSignpostIntervalData {
	return MXSignpostIntervalData{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MXSignpostIntervalDataClass) Alloc() MXSignpostIntervalData {
	rv := objc.Send[MXSignpostIntervalData](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MXSignpostIntervalDataClass) New() MXSignpostIntervalData {
	rv := objc.Send[MXSignpostIntervalData](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXSignpostIntervalData) Init() MXSignpostIntervalData {
	rv := objc.Send[MXSignpostIntervalData](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXSignpostIntervalData) Autorelease() MXSignpostIntervalData {
	rv := objc.Send[MXSignpostIntervalData](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXSignpostIntervalData creates a new MXSignpostIntervalData instance.
func NewMXSignpostIntervalData() MXSignpostIntervalData {
	return getMXSignpostIntervalDataClass().New()
}


// The data captured for a custom metric.
//
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxsignpostmetric/signpostintervaldata
func (m_ MXSignpostIntervalData) SignpostIntervalData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("signpostIntervalData"))
	return rv
}


// SetSignpostIntervalData sets the value of the signpostIntervalData property.
// The data captured for a custom metric.

//
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxsignpostmetric/signpostintervaldata
func (m_ MXSignpostIntervalData) SetSignpostIntervalData(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSignpostIntervalData:"), value)
}

// The average memory used during the logged intervals.
//
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxsignpostintervaldata/averagememory
func (m_ MXSignpostIntervalData) AverageMemory() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("averageMemory"))
	return rv
}


// SetAverageMemory sets the value of the averageMemory property.
// The average memory used during the logged intervals.

//
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxsignpostintervaldata/averagememory
func (m_ MXSignpostIntervalData) SetAverageMemory(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAverageMemory:"), value)
}

// The ratio of the total time spent hitching to the total time spent animating during the logged intervals.
//
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxsignpostintervaldata/cumulativehitchtimeratio
func (m_ MXSignpostIntervalData) CumulativeHitchTimeRatio() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cumulativeHitchTimeRatio"))
	return rv
}


// SetCumulativeHitchTimeRatio sets the value of the cumulativeHitchTimeRatio property.
// The ratio of the total time spent hitching to the total time spent animating during the logged intervals.

//
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxsignpostintervaldata/cumulativehitchtimeratio
func (m_ MXSignpostIntervalData) SetCumulativeHitchTimeRatio(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCumulativeHitchTimeRatio:"), value)
}

// The total amount of data written to disk or other long term storage during the logged intervals.
//
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxsignpostintervaldata/cumulativelogicalwrites
func (m_ MXSignpostIntervalData) CumulativeLogicalWrites() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cumulativeLogicalWrites"))
	return rv
}


// SetCumulativeLogicalWrites sets the value of the cumulativeLogicalWrites property.
// The total amount of data written to disk or other long term storage during the logged intervals.

//
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxsignpostintervaldata/cumulativelogicalwrites
func (m_ MXSignpostIntervalData) SetCumulativeLogicalWrites(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCumulativeLogicalWrites:"), value)
}

// A histogram of the different time intervals of a custom metric event.
//
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxsignpostintervaldata/histogrammedsignpostduration
func (m_ MXSignpostIntervalData) HistogrammedSignpostDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("histogrammedSignpostDuration"))
	return rv
}


// SetHistogrammedSignpostDuration sets the value of the histogrammedSignpostDuration property.
// A histogram of the different time intervals of a custom metric event.

//
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxsignpostintervaldata/histogrammedsignpostduration
func (m_ MXSignpostIntervalData) SetHistogrammedSignpostDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHistogrammedSignpostDuration:"), value)
}

// The total amount of CPU time used during the logged intervals.
//
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxsignpostintervaldata/cumulativecputime
func (m_ MXSignpostIntervalData) CumulativeCPUTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cumulativeCPUTime"))
	return rv
}


// SetCumulativeCPUTime sets the value of the cumulativeCPUTime property.
// The total amount of CPU time used during the logged intervals.

//
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxsignpostintervaldata/cumulativecputime
func (m_ MXSignpostIntervalData) SetCumulativeCPUTime(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCumulativeCPUTime:"), value)
}



