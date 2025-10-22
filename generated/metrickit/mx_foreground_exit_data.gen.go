// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MXForegroundExitData] class.
var (
	MXForegroundExitDataClass     _MXForegroundExitDataClass
	MXForegroundExitDataClassOnce sync.Once
)

func getMXForegroundExitDataClass() _MXForegroundExitDataClass {
	MXForegroundExitDataClassOnce.Do(func() {
		MXForegroundExitDataClass = _MXForegroundExitDataClass{objc.GetClass("MXForegroundExitData")}
	})
	return MXForegroundExitDataClass
}

type _MXForegroundExitDataClass struct {
	class objc.Class
}

// An interface definition for the [MXForegroundExitData] class.
type IMXForegroundExitData interface {
	objectivec.IObject
	CumulativeBadAccessExitCount() uint
	CumulativeIllegalInstructionExitCount() uint
	CumulativeMemoryResourceLimitExitCount() uint
	CumulativeNormalAppExitCount() uint
	ForegroundExitData() MXForegroundExitData
	SetForegroundExitData(value IMXForegroundExitData)
	CumulativeAbnormalExitCount() int
	SetCumulativeAbnormalExitCount(value int)
	CumulativeAppWatchdogExitCount() int
	SetCumulativeAppWatchdogExitCount(value int)
}

// An object representing counts for the different types of foreground app exits.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXForegroundExitData
type MXForegroundExitData struct {
	objectivec.Object
}

// MXForegroundExitDataFrom constructs a [MXForegroundExitData] from an unsafe.Pointer.
//
// An object representing counts for the different types of foreground app exits.
func MXForegroundExitDataFrom(ptr unsafe.Pointer) MXForegroundExitData {
	return MXForegroundExitData{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MXForegroundExitDataClass) Alloc() MXForegroundExitData {
	rv := objc.Send[MXForegroundExitData](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MXForegroundExitDataClass) New() MXForegroundExitData {
	rv := objc.Send[MXForegroundExitData](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXForegroundExitData) Init() MXForegroundExitData {
	rv := objc.Send[MXForegroundExitData](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXForegroundExitData) Autorelease() MXForegroundExitData {
	rv := objc.Send[MXForegroundExitData](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXForegroundExitData creates a new MXForegroundExitData instance.
func NewMXForegroundExitData() MXForegroundExitData {
	return getMXForegroundExitDataClass().New()
}


// The number of times the system terminated the app from the foreground for attempting an invalid memory access.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXForegroundExitData/cumulativeBadAccessExitCount
func (m_ MXForegroundExitData) CumulativeBadAccessExitCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("cumulativeBadAccessExitCount"))
	return rv
}

// The number of times the system terminated the app from the foreground for attempting to execute an illegal or undefined instruction.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXForegroundExitData/cumulativeIllegalInstructionExitCount
func (m_ MXForegroundExitData) CumulativeIllegalInstructionExitCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("cumulativeIllegalInstructionExitCount"))
	return rv
}

// The number of times the system terminated the app from the foreground for using too much memory.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXForegroundExitData/cumulativeMemoryResourceLimitExitCount
func (m_ MXForegroundExitData) CumulativeMemoryResourceLimitExitCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("cumulativeMemoryResourceLimitExitCount"))
	return rv
}

// The number of times the app exited normally from the foreground.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXForegroundExitData/cumulativeNormalAppExitCount
func (m_ MXForegroundExitData) CumulativeNormalAppExitCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("cumulativeNormalAppExitCount"))
	return rv
}

// The metrics for the foreground app exits.
//
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxappexitmetric/foregroundexitdata
func (m_ MXForegroundExitData) ForegroundExitData() MXForegroundExitData {
	rv := objc.Send[MXForegroundExitData](m_.ID, objc.Sel("foregroundExitData"))
	return rv
}


// SetForegroundExitData sets the value of the foregroundExitData property.
// The metrics for the foreground app exits.

//
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxappexitmetric/foregroundexitdata
func (m_ MXForegroundExitData) SetForegroundExitData(value IMXForegroundExitData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setForegroundExitData:"), value)
}

// The number of times the app exited abnormally from the foreground.
//
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxforegroundexitdata/cumulativeabnormalexitcount
func (m_ MXForegroundExitData) CumulativeAbnormalExitCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("cumulativeAbnormalExitCount"))
	return rv
}


// SetCumulativeAbnormalExitCount sets the value of the cumulativeAbnormalExitCount property.
// The number of times the app exited abnormally from the foreground.

//
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxforegroundexitdata/cumulativeabnormalexitcount
func (m_ MXForegroundExitData) SetCumulativeAbnormalExitCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCumulativeAbnormalExitCount:"), value)
}

// The number of times the system watchdog terminated the app from the foreground.
//
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxforegroundexitdata/cumulativeappwatchdogexitcount
func (m_ MXForegroundExitData) CumulativeAppWatchdogExitCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("cumulativeAppWatchdogExitCount"))
	return rv
}


// SetCumulativeAppWatchdogExitCount sets the value of the cumulativeAppWatchdogExitCount property.
// The number of times the system watchdog terminated the app from the foreground.

//
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxforegroundexitdata/cumulativeappwatchdogexitcount
func (m_ MXForegroundExitData) SetCumulativeAppWatchdogExitCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCumulativeAppWatchdogExitCount:"), value)
}



