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
	// properties:
	DaysMask() objc.IObject /* cross-framework: NSNumber */
	SetDaysMask(value objc.IObject /* cross-framework: NSNumber */)
	EndHour() objc.IObject /* cross-framework: NSNumber */
	SetEndHour(value objc.IObject /* cross-framework: NSNumber */)
	EndMinute() objc.IObject /* cross-framework: NSNumber */
	SetEndMinute(value objc.IObject /* cross-framework: NSNumber */)
	StartHour() objc.IObject /* cross-framework: NSNumber */
	SetStartHour(value objc.IObject /* cross-framework: NSNumber */)
	StartMinute() objc.IObject /* cross-framework: NSNumber */
	SetStartMinute(value objc.IObject /* cross-framework: NSNumber */)
	Status() objc.IObject /* cross-framework: NSNumber */
	SetStatus(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	UserIndex() objc.IObject /* cross-framework: NSNumber */
	SetUserIndex(value objc.IObject /* cross-framework: NSNumber */)
	WeekDayIndex() objc.IObject /* cross-framework: NSNumber */
	SetWeekDayIndex(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleresponseparams/daysmask
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) DaysMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("daysMask"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleresponseparams/daysmask
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) SetDaysMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDaysMask:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleresponseparams/endhour
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) EndHour() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endHour"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleresponseparams/endhour
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) SetEndHour(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndHour:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleresponseparams/endminute
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) EndMinute() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endMinute"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleresponseparams/endminute
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) SetEndMinute(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndMinute:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleresponseparams/starthour
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) StartHour() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("startHour"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleresponseparams/starthour
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) SetStartHour(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartHour:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleresponseparams/startminute
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) StartMinute() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("startMinute"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleresponseparams/startminute
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) SetStartMinute(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartMinute:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleresponseparams/status
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleresponseparams/status
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleresponseparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleresponseparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleresponseparams/userindex
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) UserIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleresponseparams/userindex
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) SetUserIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleresponseparams/weekdayindex
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) WeekDayIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("weekDayIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleresponseparams/weekdayindex
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) SetWeekDayIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWeekDayIndex:"), value)
}



