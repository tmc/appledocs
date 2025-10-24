// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterGetWeekDayScheduleParams] class.
var (
	MTRDoorLockClusterGetWeekDayScheduleParamsClass     _MTRDoorLockClusterGetWeekDayScheduleParamsClass
	MTRDoorLockClusterGetWeekDayScheduleParamsClassOnce sync.Once
)

func getMTRDoorLockClusterGetWeekDayScheduleParamsClass() _MTRDoorLockClusterGetWeekDayScheduleParamsClass {
	MTRDoorLockClusterGetWeekDayScheduleParamsClassOnce.Do(func() {
		MTRDoorLockClusterGetWeekDayScheduleParamsClass = _MTRDoorLockClusterGetWeekDayScheduleParamsClass{objc.GetClass("MTRDoorLockClusterGetWeekDayScheduleParams")}
	})
	return MTRDoorLockClusterGetWeekDayScheduleParamsClass
}

type _MTRDoorLockClusterGetWeekDayScheduleParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterGetWeekDayScheduleParams] class.
type IMTRDoorLockClusterGetWeekDayScheduleParams interface {
	objectivec.IObject
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	UserIndex() objc.IObject /* cross-framework: NSNumber */
	SetUserIndex(value objc.IObject /* cross-framework: NSNumber */)
	WeekDayIndex() objc.IObject /* cross-framework: NSNumber */
	SetWeekDayIndex(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetWeekDayScheduleParams
type MTRDoorLockClusterGetWeekDayScheduleParams struct {
	objectivec.Object
}

// MTRDoorLockClusterGetWeekDayScheduleParamsFrom constructs a [MTRDoorLockClusterGetWeekDayScheduleParams] from an unsafe.Pointer.
func MTRDoorLockClusterGetWeekDayScheduleParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterGetWeekDayScheduleParams {
	return MTRDoorLockClusterGetWeekDayScheduleParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterGetWeekDayScheduleParamsClass) Alloc() MTRDoorLockClusterGetWeekDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterGetWeekDayScheduleParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterGetWeekDayScheduleParamsClass) New() MTRDoorLockClusterGetWeekDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterGetWeekDayScheduleParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterGetWeekDayScheduleParams) Init() MTRDoorLockClusterGetWeekDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterGetWeekDayScheduleParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterGetWeekDayScheduleParams) Autorelease() MTRDoorLockClusterGetWeekDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterGetWeekDayScheduleParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterGetWeekDayScheduleParams creates a new MTRDoorLockClusterGetWeekDayScheduleParams instance.
func NewMTRDoorLockClusterGetWeekDayScheduleParams() MTRDoorLockClusterGetWeekDayScheduleParams {
	return getMTRDoorLockClusterGetWeekDayScheduleParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterGetWeekDayScheduleParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterGetWeekDayScheduleParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterGetWeekDayScheduleParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterGetWeekDayScheduleParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleparams/userindex
func (m_ MTRDoorLockClusterGetWeekDayScheduleParams) UserIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleparams/userindex
func (m_ MTRDoorLockClusterGetWeekDayScheduleParams) SetUserIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleparams/weekdayindex
func (m_ MTRDoorLockClusterGetWeekDayScheduleParams) WeekDayIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("weekDayIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetweekdayscheduleparams/weekdayindex
func (m_ MTRDoorLockClusterGetWeekDayScheduleParams) SetWeekDayIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWeekDayIndex:"), value)
}



