// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterGetYearDayScheduleResponseParams] class.
var (
	MTRDoorLockClusterGetYearDayScheduleResponseParamsClass     _MTRDoorLockClusterGetYearDayScheduleResponseParamsClass
	MTRDoorLockClusterGetYearDayScheduleResponseParamsClassOnce sync.Once
)

func getMTRDoorLockClusterGetYearDayScheduleResponseParamsClass() _MTRDoorLockClusterGetYearDayScheduleResponseParamsClass {
	MTRDoorLockClusterGetYearDayScheduleResponseParamsClassOnce.Do(func() {
		MTRDoorLockClusterGetYearDayScheduleResponseParamsClass = _MTRDoorLockClusterGetYearDayScheduleResponseParamsClass{objc.GetClass("MTRDoorLockClusterGetYearDayScheduleResponseParams")}
	})
	return MTRDoorLockClusterGetYearDayScheduleResponseParamsClass
}

type _MTRDoorLockClusterGetYearDayScheduleResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterGetYearDayScheduleResponseParams] class.
type IMTRDoorLockClusterGetYearDayScheduleResponseParams interface {
	objectivec.IObject
	// properties:
	LocalEndTime() objc.IObject /* cross-framework: NSNumber */
	SetLocalEndTime(value objc.IObject /* cross-framework: NSNumber */)
	LocalStartTime() objc.IObject /* cross-framework: NSNumber */
	SetLocalStartTime(value objc.IObject /* cross-framework: NSNumber */)
	Status() objc.IObject /* cross-framework: NSNumber */
	SetStatus(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	UserIndex() objc.IObject /* cross-framework: NSNumber */
	SetUserIndex(value objc.IObject /* cross-framework: NSNumber */)
	YearDayIndex() objc.IObject /* cross-framework: NSNumber */
	SetYearDayIndex(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetYearDayScheduleResponseParams
type MTRDoorLockClusterGetYearDayScheduleResponseParams struct {
	objectivec.Object
}

// MTRDoorLockClusterGetYearDayScheduleResponseParamsFrom constructs a [MTRDoorLockClusterGetYearDayScheduleResponseParams] from an unsafe.Pointer.
func MTRDoorLockClusterGetYearDayScheduleResponseParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterGetYearDayScheduleResponseParams {
	return MTRDoorLockClusterGetYearDayScheduleResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterGetYearDayScheduleResponseParamsClass) Alloc() MTRDoorLockClusterGetYearDayScheduleResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetYearDayScheduleResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterGetYearDayScheduleResponseParamsClass) New() MTRDoorLockClusterGetYearDayScheduleResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetYearDayScheduleResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) Init() MTRDoorLockClusterGetYearDayScheduleResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetYearDayScheduleResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) Autorelease() MTRDoorLockClusterGetYearDayScheduleResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetYearDayScheduleResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterGetYearDayScheduleResponseParams creates a new MTRDoorLockClusterGetYearDayScheduleResponseParams instance.
func NewMTRDoorLockClusterGetYearDayScheduleResponseParams() MTRDoorLockClusterGetYearDayScheduleResponseParams {
	return getMTRDoorLockClusterGetYearDayScheduleResponseParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetyeardayscheduleresponseparams/localendtime
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) LocalEndTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("localEndTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetyeardayscheduleresponseparams/localendtime
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) SetLocalEndTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocalEndTime:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetyeardayscheduleresponseparams/localstarttime
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) LocalStartTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("localStartTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetyeardayscheduleresponseparams/localstarttime
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) SetLocalStartTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocalStartTime:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetyeardayscheduleresponseparams/status
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetyeardayscheduleresponseparams/status
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetyeardayscheduleresponseparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetyeardayscheduleresponseparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetyeardayscheduleresponseparams/userindex
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) UserIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetyeardayscheduleresponseparams/userindex
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) SetUserIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetyeardayscheduleresponseparams/yeardayindex
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) YearDayIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("yearDayIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetyeardayscheduleresponseparams/yeardayindex
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) SetYearDayIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setYearDayIndex:"), value)
}



