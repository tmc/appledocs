// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterGetHolidayScheduleResponseParams] class.
var (
	MTRDoorLockClusterGetHolidayScheduleResponseParamsClass     _MTRDoorLockClusterGetHolidayScheduleResponseParamsClass
	MTRDoorLockClusterGetHolidayScheduleResponseParamsClassOnce sync.Once
)

func getMTRDoorLockClusterGetHolidayScheduleResponseParamsClass() _MTRDoorLockClusterGetHolidayScheduleResponseParamsClass {
	MTRDoorLockClusterGetHolidayScheduleResponseParamsClassOnce.Do(func() {
		MTRDoorLockClusterGetHolidayScheduleResponseParamsClass = _MTRDoorLockClusterGetHolidayScheduleResponseParamsClass{objc.GetClass("MTRDoorLockClusterGetHolidayScheduleResponseParams")}
	})
	return MTRDoorLockClusterGetHolidayScheduleResponseParamsClass
}

type _MTRDoorLockClusterGetHolidayScheduleResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterGetHolidayScheduleResponseParams] class.
type IMTRDoorLockClusterGetHolidayScheduleResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetHolidayScheduleResponseParams
type MTRDoorLockClusterGetHolidayScheduleResponseParams struct {
	objectivec.Object
}

// MTRDoorLockClusterGetHolidayScheduleResponseParamsFrom constructs a [MTRDoorLockClusterGetHolidayScheduleResponseParams] from an unsafe.Pointer.
func MTRDoorLockClusterGetHolidayScheduleResponseParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterGetHolidayScheduleResponseParams {
	return MTRDoorLockClusterGetHolidayScheduleResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterGetHolidayScheduleResponseParamsClass) Alloc() MTRDoorLockClusterGetHolidayScheduleResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetHolidayScheduleResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterGetHolidayScheduleResponseParamsClass) New() MTRDoorLockClusterGetHolidayScheduleResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetHolidayScheduleResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) Init() MTRDoorLockClusterGetHolidayScheduleResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetHolidayScheduleResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) Autorelease() MTRDoorLockClusterGetHolidayScheduleResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetHolidayScheduleResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterGetHolidayScheduleResponseParams creates a new MTRDoorLockClusterGetHolidayScheduleResponseParams instance.
func NewMTRDoorLockClusterGetHolidayScheduleResponseParams() MTRDoorLockClusterGetHolidayScheduleResponseParams {
	return getMTRDoorLockClusterGetHolidayScheduleResponseParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetholidayscheduleresponseparams/holidayindex
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) HolidayIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("holidayIndex"))
	return rv
}


// SetHolidayIndex sets the value of the holidayIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetholidayscheduleresponseparams/holidayindex
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) SetHolidayIndex(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHolidayIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetholidayscheduleresponseparams/localendtime
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) LocalEndTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("localEndTime"))
	return rv
}


// SetLocalEndTime sets the value of the localEndTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetholidayscheduleresponseparams/localendtime
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) SetLocalEndTime(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocalEndTime:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetholidayscheduleresponseparams/localstarttime
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) LocalStartTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("localStartTime"))
	return rv
}


// SetLocalStartTime sets the value of the localStartTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetholidayscheduleresponseparams/localstarttime
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) SetLocalStartTime(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocalStartTime:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetholidayscheduleresponseparams/operatingmode
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) OperatingMode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("operatingMode"))
	return rv
}


// SetOperatingMode sets the value of the operatingMode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetholidayscheduleresponseparams/operatingmode
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) SetOperatingMode(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperatingMode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetholidayscheduleresponseparams/status
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) Status() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("status"))
	return rv
}


// SetStatus sets the value of the status property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetholidayscheduleresponseparams/status
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) SetStatus(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetholidayscheduleresponseparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetholidayscheduleresponseparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



