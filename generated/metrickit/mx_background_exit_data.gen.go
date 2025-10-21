// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
}

// An object representing counts for the different types of background app exits.
//
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


// The number of times the app exited abnormally from the background.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXBackgroundExitData/cumulativeAbnormalExitCount
func (m_ MXBackgroundExitData) CumulativeAbnormalExitCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("cumulativeAbnormalExitCount"))
	return rv
}

// The number of times the system watchdog terminated the app from the background.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXBackgroundExitData/cumulativeAppWatchdogExitCount
func (m_ MXBackgroundExitData) CumulativeAppWatchdogExitCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("cumulativeAppWatchdogExitCount"))
	return rv
}

// The number of times the system terminated the app from the background for exceeding the allocated time for a background task.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXBackgroundExitData/cumulativeBackgroundTaskAssertionTimeoutExitCount
func (m_ MXBackgroundExitData) CumulativeBackgroundTaskAssertionTimeoutExitCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("cumulativeBackgroundTaskAssertionTimeoutExitCount"))
	return rv
}

// The number of times the system terminated the app from the background for attempting an invalid memory access.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXBackgroundExitData/cumulativeBadAccessExitCount
func (m_ MXBackgroundExitData) CumulativeBadAccessExitCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("cumulativeBadAccessExitCount"))
	return rv
}

// The number of times the system terminated the app from the background for using too much CPU time.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXBackgroundExitData/cumulativeCPUResourceLimitExitCount
func (m_ MXBackgroundExitData) CumulativeCPUResourceLimitExitCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("cumulativeCPUResourceLimitExitCount"))
	return rv
}

// The number of times the system terminated the app from the background for attempting to execute an illegal or undefined instruction.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXBackgroundExitData/cumulativeIllegalInstructionExitCount
func (m_ MXBackgroundExitData) CumulativeIllegalInstructionExitCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("cumulativeIllegalInstructionExitCount"))
	return rv
}

// The number of times the system terminated the app from the background to free up memory.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXBackgroundExitData/cumulativeMemoryPressureExitCount
func (m_ MXBackgroundExitData) CumulativeMemoryPressureExitCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("cumulativeMemoryPressureExitCount"))
	return rv
}

// The number of times the system terminated the app from the background for using too much memory.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXBackgroundExitData/cumulativeMemoryResourceLimitExitCount
func (m_ MXBackgroundExitData) CumulativeMemoryResourceLimitExitCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("cumulativeMemoryResourceLimitExitCount"))
	return rv
}

// The number of times the app exited normally from the background.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXBackgroundExitData/cumulativeNormalAppExitCount
func (m_ MXBackgroundExitData) CumulativeNormalAppExitCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("cumulativeNormalAppExitCount"))
	return rv
}

// The number of times the system terminated the app from the background while being suspended and having file locks.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXBackgroundExitData/cumulativeSuspendedWithLockedFileExitCount
func (m_ MXBackgroundExitData) CumulativeSuspendedWithLockedFileExitCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("cumulativeSuspendedWithLockedFileExitCount"))
	return rv
}



