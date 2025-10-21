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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetyeardayscheduleresponseparams/localendtime
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) LocalEndTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("localEndTime"))
	return rv
}


// SetLocalEndTime sets the value of the localEndTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetyeardayscheduleresponseparams/localendtime
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) SetLocalEndTime(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocalEndTime:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetyeardayscheduleresponseparams/localstarttime
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) LocalStartTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("localStartTime"))
	return rv
}


// SetLocalStartTime sets the value of the localStartTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetyeardayscheduleresponseparams/localstarttime
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) SetLocalStartTime(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocalStartTime:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetyeardayscheduleresponseparams/status
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) Status() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("status"))
	return rv
}


// SetStatus sets the value of the status property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetyeardayscheduleresponseparams/status
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) SetStatus(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetyeardayscheduleresponseparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetyeardayscheduleresponseparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetyeardayscheduleresponseparams/userindex
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) UserIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("userIndex"))
	return rv
}


// SetUserIndex sets the value of the userIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetyeardayscheduleresponseparams/userindex
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) SetUserIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetyeardayscheduleresponseparams/yeardayindex
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) YearDayIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("yearDayIndex"))
	return rv
}


// SetYearDayIndex sets the value of the yearDayIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetyeardayscheduleresponseparams/yeardayindex
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) SetYearDayIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setYearDayIndex:"), value)
}



