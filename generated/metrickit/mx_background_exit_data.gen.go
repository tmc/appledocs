// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MXBackgroundExitData] class.
var (
	MXBackgroundExitDataClass     _MXBackgroundExitDataClass
	MXBackgroundExitDataClassOnce sync.Once
)

func getMXBackgroundExitDataClass() _MXBackgroundExitDataClass {
	MXBackgroundExitDataClassOnce.Do(func() {
		MXBackgroundExitDataClass = _MXBackgroundExitDataClass{objc.GetClass("MXBackgroundExitData")}
	})
	return MXBackgroundExitDataClass
}

type _MXBackgroundExitDataClass struct {
	class objc.Class
}

// An interface definition for the [MXBackgroundExitData] class.
type IMXBackgroundExitData interface {
	objectivec.IObject
	// properties:
	BackgroundExitData() IMXBackgroundExitData
	SetBackgroundExitData(value IMXBackgroundExitData)
	CumulativeAbnormalExitCount() int
	SetCumulativeAbnormalExitCount(value int)
	CumulativeAppWatchdogExitCount() int
	SetCumulativeAppWatchdogExitCount(value int)
	CumulativeBackgroundTaskAssertionTimeoutExitCount() int
	SetCumulativeBackgroundTaskAssertionTimeoutExitCount(value int)
	CumulativeBadAccessExitCount() int
	SetCumulativeBadAccessExitCount(value int)
	CumulativeCPUResourceLimitExitCount() int
	SetCumulativeCPUResourceLimitExitCount(value int)
	CumulativeIllegalInstructionExitCount() int
	SetCumulativeIllegalInstructionExitCount(value int)
	CumulativeMemoryPressureExitCount() int
	SetCumulativeMemoryPressureExitCount(value int)
	CumulativeMemoryResourceLimitExitCount() int
	SetCumulativeMemoryResourceLimitExitCount(value int)
	CumulativeNormalAppExitCount() int
	SetCumulativeNormalAppExitCount(value int)
	CumulativeSuspendedWithLockedFileExitCount() int
	SetCumulativeSuspendedWithLockedFileExitCount(value int)
	// methods:
}

// An object representing counts for the different types of background app exits.


// An object representing counts for the different types of background app exits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXBackgroundExitData
type MXBackgroundExitData struct {
	objectivec.Object
}

// MXBackgroundExitDataFrom constructs a [MXBackgroundExitData] from an unsafe.Pointer.
//
// An object representing counts for the different types of background app exits.
func MXBackgroundExitDataFrom(ptr unsafe.Pointer) MXBackgroundExitData {
	return MXBackgroundExitData{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MXBackgroundExitDataClass) Alloc() MXBackgroundExitData {
	rv := objc.Send[MXBackgroundExitData](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MXBackgroundExitDataClass) New() MXBackgroundExitData {
	rv := objc.Send[MXBackgroundExitData](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXBackgroundExitData) Init() MXBackgroundExitData {
	rv := objc.Send[MXBackgroundExitData](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXBackgroundExitData) Autorelease() MXBackgroundExitData {
	rv := objc.Send[MXBackgroundExitData](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXBackgroundExitData creates a new MXBackgroundExitData instance.
func NewMXBackgroundExitData() MXBackgroundExitData {
	return getMXBackgroundExitDataClass().New()
}



// The metrics for the background app exits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxappexitmetric/backgroundexitdata
func (m_ MXBackgroundExitData) BackgroundExitData() IMXBackgroundExitData {
	rv := objc.Send[MXBackgroundExitData](m_.ID, objc.Sel("backgroundExitData"))
	return rv
}


// The metrics for the background app exits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxappexitmetric/backgroundexitdata
func (m_ MXBackgroundExitData) SetBackgroundExitData(value IMXBackgroundExitData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBackgroundExitData:"), value)
}


// The number of times the app exited abnormally from the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxbackgroundexitdata/cumulativeabnormalexitcount
func (m_ MXBackgroundExitData) CumulativeAbnormalExitCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("cumulativeAbnormalExitCount"))
	return rv
}


// The number of times the app exited abnormally from the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxbackgroundexitdata/cumulativeabnormalexitcount
func (m_ MXBackgroundExitData) SetCumulativeAbnormalExitCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCumulativeAbnormalExitCount:"), value)
}


// The number of times the system watchdog terminated the app from the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxbackgroundexitdata/cumulativeappwatchdogexitcount
func (m_ MXBackgroundExitData) CumulativeAppWatchdogExitCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("cumulativeAppWatchdogExitCount"))
	return rv
}


// The number of times the system watchdog terminated the app from the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxbackgroundexitdata/cumulativeappwatchdogexitcount
func (m_ MXBackgroundExitData) SetCumulativeAppWatchdogExitCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCumulativeAppWatchdogExitCount:"), value)
}


// The number of times the system terminated the app from the background for exceeding the allocated time for a background task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxbackgroundexitdata/cumulativebackgroundtaskassertiontimeoutexitcount
func (m_ MXBackgroundExitData) CumulativeBackgroundTaskAssertionTimeoutExitCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("cumulativeBackgroundTaskAssertionTimeoutExitCount"))
	return rv
}


// The number of times the system terminated the app from the background for exceeding the allocated time for a background task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxbackgroundexitdata/cumulativebackgroundtaskassertiontimeoutexitcount
func (m_ MXBackgroundExitData) SetCumulativeBackgroundTaskAssertionTimeoutExitCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCumulativeBackgroundTaskAssertionTimeoutExitCount:"), value)
}


// The number of times the system terminated the app from the background for attempting an invalid memory access.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxbackgroundexitdata/cumulativebadaccessexitcount
func (m_ MXBackgroundExitData) CumulativeBadAccessExitCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("cumulativeBadAccessExitCount"))
	return rv
}


// The number of times the system terminated the app from the background for attempting an invalid memory access.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxbackgroundexitdata/cumulativebadaccessexitcount
func (m_ MXBackgroundExitData) SetCumulativeBadAccessExitCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCumulativeBadAccessExitCount:"), value)
}


// The number of times the system terminated the app from the background for using too much CPU time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxbackgroundexitdata/cumulativecpuresourcelimitexitcount
func (m_ MXBackgroundExitData) CumulativeCPUResourceLimitExitCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("cumulativeCPUResourceLimitExitCount"))
	return rv
}


// The number of times the system terminated the app from the background for using too much CPU time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxbackgroundexitdata/cumulativecpuresourcelimitexitcount
func (m_ MXBackgroundExitData) SetCumulativeCPUResourceLimitExitCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCumulativeCPUResourceLimitExitCount:"), value)
}


// The number of times the system terminated the app from the background for attempting to execute an illegal or undefined instruction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxbackgroundexitdata/cumulativeillegalinstructionexitcount
func (m_ MXBackgroundExitData) CumulativeIllegalInstructionExitCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("cumulativeIllegalInstructionExitCount"))
	return rv
}


// The number of times the system terminated the app from the background for attempting to execute an illegal or undefined instruction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxbackgroundexitdata/cumulativeillegalinstructionexitcount
func (m_ MXBackgroundExitData) SetCumulativeIllegalInstructionExitCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCumulativeIllegalInstructionExitCount:"), value)
}


// The number of times the system terminated the app from the background to free up memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxbackgroundexitdata/cumulativememorypressureexitcount
func (m_ MXBackgroundExitData) CumulativeMemoryPressureExitCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("cumulativeMemoryPressureExitCount"))
	return rv
}


// The number of times the system terminated the app from the background to free up memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxbackgroundexitdata/cumulativememorypressureexitcount
func (m_ MXBackgroundExitData) SetCumulativeMemoryPressureExitCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCumulativeMemoryPressureExitCount:"), value)
}


// The number of times the system terminated the app from the background for using too much memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxbackgroundexitdata/cumulativememoryresourcelimitexitcount
func (m_ MXBackgroundExitData) CumulativeMemoryResourceLimitExitCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("cumulativeMemoryResourceLimitExitCount"))
	return rv
}


// The number of times the system terminated the app from the background for using too much memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxbackgroundexitdata/cumulativememoryresourcelimitexitcount
func (m_ MXBackgroundExitData) SetCumulativeMemoryResourceLimitExitCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCumulativeMemoryResourceLimitExitCount:"), value)
}


// The number of times the app exited normally from the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxbackgroundexitdata/cumulativenormalappexitcount
func (m_ MXBackgroundExitData) CumulativeNormalAppExitCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("cumulativeNormalAppExitCount"))
	return rv
}


// The number of times the app exited normally from the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxbackgroundexitdata/cumulativenormalappexitcount
func (m_ MXBackgroundExitData) SetCumulativeNormalAppExitCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCumulativeNormalAppExitCount:"), value)
}


// The number of times the system terminated the app from the background while being suspended and having file locks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxbackgroundexitdata/cumulativesuspendedwithlockedfileexitcount
func (m_ MXBackgroundExitData) CumulativeSuspendedWithLockedFileExitCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("cumulativeSuspendedWithLockedFileExitCount"))
	return rv
}


// The number of times the system terminated the app from the background while being suspended and having file locks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxbackgroundexitdata/cumulativesuspendedwithlockedfileexitcount
func (m_ MXBackgroundExitData) SetCumulativeSuspendedWithLockedFileExitCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCumulativeSuspendedWithLockedFileExitCount:"), value)
}



