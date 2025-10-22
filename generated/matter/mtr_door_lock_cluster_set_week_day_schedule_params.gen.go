// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterSetWeekDayScheduleParams] class.
var (
	MTRDoorLockClusterSetWeekDayScheduleParamsClass     _MTRDoorLockClusterSetWeekDayScheduleParamsClass
	MTRDoorLockClusterSetWeekDayScheduleParamsClassOnce sync.Once
)

func getMTRDoorLockClusterSetWeekDayScheduleParamsClass() _MTRDoorLockClusterSetWeekDayScheduleParamsClass {
	MTRDoorLockClusterSetWeekDayScheduleParamsClassOnce.Do(func() {
		MTRDoorLockClusterSetWeekDayScheduleParamsClass = _MTRDoorLockClusterSetWeekDayScheduleParamsClass{objc.GetClass("MTRDoorLockClusterSetWeekDayScheduleParams")}
	})
	return MTRDoorLockClusterSetWeekDayScheduleParamsClass
}

type _MTRDoorLockClusterSetWeekDayScheduleParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterSetWeekDayScheduleParams] class.
type IMTRDoorLockClusterSetWeekDayScheduleParams interface {
	objectivec.IObject
	DaysMask() foundation.Number
	SetDaysMask(value foundation.INumber)
	EndHour() foundation.Number
	SetEndHour(value foundation.INumber)
	EndMinute() foundation.Number
	SetEndMinute(value foundation.INumber)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	StartHour() foundation.Number
	SetStartHour(value foundation.INumber)
	StartMinute() foundation.Number
	SetStartMinute(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
	UserIndex() foundation.Number
	SetUserIndex(value foundation.INumber)
	WeekDayIndex() foundation.Number
	SetWeekDayIndex(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetWeekDayScheduleParams
type MTRDoorLockClusterSetWeekDayScheduleParams struct {
	objectivec.Object
}

// MTRDoorLockClusterSetWeekDayScheduleParamsFrom constructs a [MTRDoorLockClusterSetWeekDayScheduleParams] from an unsafe.Pointer.
func MTRDoorLockClusterSetWeekDayScheduleParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterSetWeekDayScheduleParams {
	return MTRDoorLockClusterSetWeekDayScheduleParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterSetWeekDayScheduleParamsClass) Alloc() MTRDoorLockClusterSetWeekDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterSetWeekDayScheduleParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterSetWeekDayScheduleParamsClass) New() MTRDoorLockClusterSetWeekDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterSetWeekDayScheduleParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) Init() MTRDoorLockClusterSetWeekDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterSetWeekDayScheduleParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) Autorelease() MTRDoorLockClusterSetWeekDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterSetWeekDayScheduleParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterSetWeekDayScheduleParams creates a new MTRDoorLockClusterSetWeekDayScheduleParams instance.
func NewMTRDoorLockClusterSetWeekDayScheduleParams() MTRDoorLockClusterSetWeekDayScheduleParams {
	return getMTRDoorLockClusterSetWeekDayScheduleParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetweekdayscheduleparams/daysmask
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) DaysMask() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("daysMask"))
	return rv
}


// SetDaysMask sets the value of the daysMask property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetweekdayscheduleparams/daysmask
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) SetDaysMask(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDaysMask:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetweekdayscheduleparams/endhour
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) EndHour() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("endHour"))
	return rv
}


// SetEndHour sets the value of the endHour property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetweekdayscheduleparams/endhour
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) SetEndHour(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndHour:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetweekdayscheduleparams/endminute
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) EndMinute() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("endMinute"))
	return rv
}


// SetEndMinute sets the value of the endMinute property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetweekdayscheduleparams/endminute
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) SetEndMinute(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndMinute:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetweekdayscheduleparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetweekdayscheduleparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetweekdayscheduleparams/starthour
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) StartHour() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("startHour"))
	return rv
}


// SetStartHour sets the value of the startHour property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetweekdayscheduleparams/starthour
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) SetStartHour(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartHour:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetweekdayscheduleparams/startminute
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) StartMinute() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("startMinute"))
	return rv
}


// SetStartMinute sets the value of the startMinute property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetweekdayscheduleparams/startminute
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) SetStartMinute(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartMinute:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetweekdayscheduleparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetweekdayscheduleparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetweekdayscheduleparams/userindex
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) UserIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("userIndex"))
	return rv
}


// SetUserIndex sets the value of the userIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetweekdayscheduleparams/userindex
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) SetUserIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetweekdayscheduleparams/weekdayindex
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) WeekDayIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("weekDayIndex"))
	return rv
}


// SetWeekDayIndex sets the value of the weekDayIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetweekdayscheduleparams/weekdayindex
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) SetWeekDayIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWeekDayIndex:"), value)
}



