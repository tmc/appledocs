// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterGetWeekDayScheduleResponseParams] class.
var (
	MTRDoorLockClusterGetWeekDayScheduleResponseParamsClass     _MTRDoorLockClusterGetWeekDayScheduleResponseParamsClass
	MTRDoorLockClusterGetWeekDayScheduleResponseParamsClassOnce sync.Once
)

func getMTRDoorLockClusterGetWeekDayScheduleResponseParamsClass() _MTRDoorLockClusterGetWeekDayScheduleResponseParamsClass {
	MTRDoorLockClusterGetWeekDayScheduleResponseParamsClassOnce.Do(func() {
		MTRDoorLockClusterGetWeekDayScheduleResponseParamsClass = _MTRDoorLockClusterGetWeekDayScheduleResponseParamsClass{objc.GetClass("MTRDoorLockClusterGetWeekDayScheduleResponseParams")}
	})
	return MTRDoorLockClusterGetWeekDayScheduleResponseParamsClass
}

type _MTRDoorLockClusterGetWeekDayScheduleResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterGetWeekDayScheduleResponseParams] class.
type IMTRDoorLockClusterGetWeekDayScheduleResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetWeekDayScheduleResponseParams
type MTRDoorLockClusterGetWeekDayScheduleResponseParams struct {
	objectivec.Object
}

// MTRDoorLockClusterGetWeekDayScheduleResponseParamsFrom constructs a [MTRDoorLockClusterGetWeekDayScheduleResponseParams] from an unsafe.Pointer.
func MTRDoorLockClusterGetWeekDayScheduleResponseParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterGetWeekDayScheduleResponseParams {
	return MTRDoorLockClusterGetWeekDayScheduleResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterGetWeekDayScheduleResponseParamsClass) Alloc() MTRDoorLockClusterGetWeekDayScheduleResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetWeekDayScheduleResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterGetWeekDayScheduleResponseParamsClass) New() MTRDoorLockClusterGetWeekDayScheduleResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetWeekDayScheduleResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) Init() MTRDoorLockClusterGetWeekDayScheduleResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetWeekDayScheduleResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) Autorelease() MTRDoorLockClusterGetWeekDayScheduleResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetWeekDayScheduleResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterGetWeekDayScheduleResponseParams creates a new MTRDoorLockClusterGetWeekDayScheduleResponseParams instance.
func NewMTRDoorLockClusterGetWeekDayScheduleResponseParams() MTRDoorLockClusterGetWeekDayScheduleResponseParams {
	return getMTRDoorLockClusterGetWeekDayScheduleResponseParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleresponseparams/daysmask
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) DaysMask() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("daysMask"))
	return rv
}


// SetDaysMask sets the value of the daysMask property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleresponseparams/daysmask
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) SetDaysMask(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDaysMask:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleresponseparams/endhour
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) EndHour() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("endHour"))
	return rv
}


// SetEndHour sets the value of the endHour property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleresponseparams/endhour
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) SetEndHour(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndHour:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleresponseparams/endminute
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) EndMinute() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("endMinute"))
	return rv
}


// SetEndMinute sets the value of the endMinute property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleresponseparams/endminute
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) SetEndMinute(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndMinute:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleresponseparams/starthour
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) StartHour() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("startHour"))
	return rv
}


// SetStartHour sets the value of the startHour property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleresponseparams/starthour
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) SetStartHour(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartHour:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleresponseparams/startminute
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) StartMinute() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("startMinute"))
	return rv
}


// SetStartMinute sets the value of the startMinute property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleresponseparams/startminute
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) SetStartMinute(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartMinute:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleresponseparams/status
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) Status() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("status"))
	return rv
}


// SetStatus sets the value of the status property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleresponseparams/status
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) SetStatus(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleresponseparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleresponseparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleresponseparams/userindex
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) UserIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("userIndex"))
	return rv
}


// SetUserIndex sets the value of the userIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleresponseparams/userindex
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) SetUserIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleresponseparams/weekdayindex
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) WeekDayIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("weekDayIndex"))
	return rv
}


// SetWeekDayIndex sets the value of the weekDayIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleresponseparams/weekdayindex
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) SetWeekDayIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWeekDayIndex:"), value)
}



