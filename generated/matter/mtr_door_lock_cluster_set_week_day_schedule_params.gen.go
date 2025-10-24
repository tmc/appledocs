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
	// properties:
	DaysMask() objc.IObject /* cross-framework: NSNumber */
	SetDaysMask(value objc.IObject /* cross-framework: NSNumber */)
	EndHour() objc.IObject /* cross-framework: NSNumber */
	SetEndHour(value objc.IObject /* cross-framework: NSNumber */)
	EndMinute() objc.IObject /* cross-framework: NSNumber */
	SetEndMinute(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	StartHour() objc.IObject /* cross-framework: NSNumber */
	SetStartHour(value objc.IObject /* cross-framework: NSNumber */)
	StartMinute() objc.IObject /* cross-framework: NSNumber */
	SetStartMinute(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	UserIndex() objc.IObject /* cross-framework: NSNumber */
	SetUserIndex(value objc.IObject /* cross-framework: NSNumber */)
	WeekDayIndex() objc.IObject /* cross-framework: NSNumber */
	SetWeekDayIndex(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetweekdayscheduleparams/daysmask
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) DaysMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("daysMask"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetweekdayscheduleparams/daysmask
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) SetDaysMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDaysMask:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetweekdayscheduleparams/endhour
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) EndHour() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endHour"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetweekdayscheduleparams/endhour
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) SetEndHour(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndHour:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetweekdayscheduleparams/endminute
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) EndMinute() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endMinute"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetweekdayscheduleparams/endminute
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) SetEndMinute(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndMinute:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetweekdayscheduleparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetweekdayscheduleparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetweekdayscheduleparams/starthour
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) StartHour() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("startHour"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetweekdayscheduleparams/starthour
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) SetStartHour(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartHour:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetweekdayscheduleparams/startminute
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) StartMinute() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("startMinute"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetweekdayscheduleparams/startminute
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) SetStartMinute(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartMinute:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetweekdayscheduleparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetweekdayscheduleparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetweekdayscheduleparams/userindex
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) UserIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetweekdayscheduleparams/userindex
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) SetUserIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetweekdayscheduleparams/weekdayindex
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) WeekDayIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("weekDayIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetweekdayscheduleparams/weekdayindex
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) SetWeekDayIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWeekDayIndex:"), value)
}



