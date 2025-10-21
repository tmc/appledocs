// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterClearWeekDayScheduleParams] class.
var (
	MTRDoorLockClusterClearWeekDayScheduleParamsClass     _MTRDoorLockClusterClearWeekDayScheduleParamsClass
	MTRDoorLockClusterClearWeekDayScheduleParamsClassOnce sync.Once
)

func getMTRDoorLockClusterClearWeekDayScheduleParamsClass() _MTRDoorLockClusterClearWeekDayScheduleParamsClass {
	MTRDoorLockClusterClearWeekDayScheduleParamsClassOnce.Do(func() {
		MTRDoorLockClusterClearWeekDayScheduleParamsClass = _MTRDoorLockClusterClearWeekDayScheduleParamsClass{objc.GetClass("MTRDoorLockClusterClearWeekDayScheduleParams")}
	})
	return MTRDoorLockClusterClearWeekDayScheduleParamsClass
}

type _MTRDoorLockClusterClearWeekDayScheduleParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterClearWeekDayScheduleParams] class.
type IMTRDoorLockClusterClearWeekDayScheduleParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearWeekDayScheduleParams
type MTRDoorLockClusterClearWeekDayScheduleParams struct {
	objectivec.Object
}

// MTRDoorLockClusterClearWeekDayScheduleParamsFrom constructs a [MTRDoorLockClusterClearWeekDayScheduleParams] from an unsafe.Pointer.
func MTRDoorLockClusterClearWeekDayScheduleParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterClearWeekDayScheduleParams {
	return MTRDoorLockClusterClearWeekDayScheduleParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterClearWeekDayScheduleParamsClass) Alloc() MTRDoorLockClusterClearWeekDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterClearWeekDayScheduleParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterClearWeekDayScheduleParamsClass) New() MTRDoorLockClusterClearWeekDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterClearWeekDayScheduleParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterClearWeekDayScheduleParams) Init() MTRDoorLockClusterClearWeekDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterClearWeekDayScheduleParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterClearWeekDayScheduleParams) Autorelease() MTRDoorLockClusterClearWeekDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterClearWeekDayScheduleParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterClearWeekDayScheduleParams creates a new MTRDoorLockClusterClearWeekDayScheduleParams instance.
func NewMTRDoorLockClusterClearWeekDayScheduleParams() MTRDoorLockClusterClearWeekDayScheduleParams {
	return getMTRDoorLockClusterClearWeekDayScheduleParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearweekdayscheduleparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterClearWeekDayScheduleParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearweekdayscheduleparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterClearWeekDayScheduleParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearweekdayscheduleparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterClearWeekDayScheduleParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearweekdayscheduleparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterClearWeekDayScheduleParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearweekdayscheduleparams/userindex
func (m_ MTRDoorLockClusterClearWeekDayScheduleParams) UserIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("userIndex"))
	return rv
}


// SetUserIndex sets the value of the userIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearweekdayscheduleparams/userindex
func (m_ MTRDoorLockClusterClearWeekDayScheduleParams) SetUserIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearweekdayscheduleparams/weekdayindex
func (m_ MTRDoorLockClusterClearWeekDayScheduleParams) WeekDayIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("weekDayIndex"))
	return rv
}


// SetWeekDayIndex sets the value of the weekDayIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearweekdayscheduleparams/weekdayindex
func (m_ MTRDoorLockClusterClearWeekDayScheduleParams) SetWeekDayIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWeekDayIndex:"), value)
}



