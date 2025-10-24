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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearweekdayscheduleparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterClearWeekDayScheduleParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearweekdayscheduleparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterClearWeekDayScheduleParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearweekdayscheduleparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterClearWeekDayScheduleParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearweekdayscheduleparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterClearWeekDayScheduleParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearweekdayscheduleparams/userindex
func (m_ MTRDoorLockClusterClearWeekDayScheduleParams) UserIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearweekdayscheduleparams/userindex
func (m_ MTRDoorLockClusterClearWeekDayScheduleParams) SetUserIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearweekdayscheduleparams/weekdayindex
func (m_ MTRDoorLockClusterClearWeekDayScheduleParams) WeekDayIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("weekDayIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearweekdayscheduleparams/weekdayindex
func (m_ MTRDoorLockClusterClearWeekDayScheduleParams) SetWeekDayIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWeekDayIndex:"), value)
}



