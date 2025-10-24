// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	AverageMemory() objc.IObject /* cross-framework: UnitInformationStorage */
	SetAverageMemory(value objc.IObject /* cross-framework: UnitInformationStorage */)
	CumulativeCPUTime() objc.IObject /* cross-framework: UnitDuration */
	SetCumulativeCPUTime(value objc.IObject /* cross-framework: UnitDuration */)
	CumulativeHitchTimeRatio() objc.IObject /* cross-framework: Unit */
	SetCumulativeHitchTimeRatio(value objc.IObject /* cross-framework: Unit */)
	CumulativeLogicalWrites() objc.IObject /* cross-framework: UnitInformationStorage */
	SetCumulativeLogicalWrites(value objc.IObject /* cross-framework: UnitInformationStorage */)
	HistogrammedSignpostDuration() objc.IObject /* cross-framework: UnitDuration */
	SetHistogrammedSignpostDuration(value objc.IObject /* cross-framework: UnitDuration */)
	SignpostIntervalData() IMXSignpostIntervalData
	SetSignpostIntervalData(value IMXSignpostIntervalData)
	// methods:
}

// A data object representing the captured data for a custom metric.


// A data object representing the captured data for a custom metric.
//
// [Full Topic]
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



// The average memory used during the logged intervals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxsignpostintervaldata/averagememory
func (m_ MXSignpostIntervalData) AverageMemory() objc.IObject /* cross-framework: UnitInformationStorage */ {
	rv := objc.Send[foundation.UnitInformationStorage](m_.ID, objc.Sel("averageMemory"))
	return rv
}


// The average memory used during the logged intervals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxsignpostintervaldata/averagememory
func (m_ MXSignpostIntervalData) SetAverageMemory(value objc.IObject /* cross-framework: UnitInformationStorage */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAverageMemory:"), value)
}


// The total amount of CPU time used during the logged intervals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxsignpostintervaldata/cumulativecputime
func (m_ MXSignpostIntervalData) CumulativeCPUTime() objc.IObject /* cross-framework: UnitDuration */ {
	rv := objc.Send[foundation.UnitDuration](m_.ID, objc.Sel("cumulativeCPUTime"))
	return rv
}


// The total amount of CPU time used during the logged intervals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxsignpostintervaldata/cumulativecputime
func (m_ MXSignpostIntervalData) SetCumulativeCPUTime(value objc.IObject /* cross-framework: UnitDuration */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCumulativeCPUTime:"), value)
}


// The ratio of the total time spent hitching to the total time spent animating during the logged intervals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxsignpostintervaldata/cumulativehitchtimeratio
func (m_ MXSignpostIntervalData) CumulativeHitchTimeRatio() objc.IObject /* cross-framework: Unit */ {
	rv := objc.Send[foundation.Unit](m_.ID, objc.Sel("cumulativeHitchTimeRatio"))
	return rv
}


// The ratio of the total time spent hitching to the total time spent animating during the logged intervals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxsignpostintervaldata/cumulativehitchtimeratio
func (m_ MXSignpostIntervalData) SetCumulativeHitchTimeRatio(value objc.IObject /* cross-framework: Unit */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCumulativeHitchTimeRatio:"), value)
}


// The total amount of data written to disk or other long term storage during the logged intervals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxsignpostintervaldata/cumulativelogicalwrites
func (m_ MXSignpostIntervalData) CumulativeLogicalWrites() objc.IObject /* cross-framework: UnitInformationStorage */ {
	rv := objc.Send[foundation.UnitInformationStorage](m_.ID, objc.Sel("cumulativeLogicalWrites"))
	return rv
}


// The total amount of data written to disk or other long term storage during the logged intervals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxsignpostintervaldata/cumulativelogicalwrites
func (m_ MXSignpostIntervalData) SetCumulativeLogicalWrites(value objc.IObject /* cross-framework: UnitInformationStorage */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCumulativeLogicalWrites:"), value)
}


// A histogram of the different time intervals of a custom metric event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxsignpostintervaldata/histogrammedsignpostduration
func (m_ MXSignpostIntervalData) HistogrammedSignpostDuration() objc.IObject /* cross-framework: UnitDuration */ {
	rv := objc.Send[foundation.UnitDuration](m_.ID, objc.Sel("histogrammedSignpostDuration"))
	return rv
}


// A histogram of the different time intervals of a custom metric event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxsignpostintervaldata/histogrammedsignpostduration
func (m_ MXSignpostIntervalData) SetHistogrammedSignpostDuration(value objc.IObject /* cross-framework: UnitDuration */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHistogrammedSignpostDuration:"), value)
}


// The data captured for a custom metric.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxsignpostmetric/signpostintervaldata
func (m_ MXSignpostIntervalData) SignpostIntervalData() IMXSignpostIntervalData {
	rv := objc.Send[MXSignpostIntervalData](m_.ID, objc.Sel("signpostIntervalData"))
	return rv
}


// The data captured for a custom metric.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxsignpostmetric/signpostintervaldata
func (m_ MXSignpostIntervalData) SetSignpostIntervalData(value IMXSignpostIntervalData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSignpostIntervalData:"), value)
}



