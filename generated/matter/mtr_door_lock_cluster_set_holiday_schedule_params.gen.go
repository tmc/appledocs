// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterSetHolidayScheduleParams] class.
var (
	MTRDoorLockClusterSetHolidayScheduleParamsClass     _MTRDoorLockClusterSetHolidayScheduleParamsClass
	MTRDoorLockClusterSetHolidayScheduleParamsClassOnce sync.Once
)

func getMTRDoorLockClusterSetHolidayScheduleParamsClass() _MTRDoorLockClusterSetHolidayScheduleParamsClass {
	MTRDoorLockClusterSetHolidayScheduleParamsClassOnce.Do(func() {
		MTRDoorLockClusterSetHolidayScheduleParamsClass = _MTRDoorLockClusterSetHolidayScheduleParamsClass{objc.GetClass("MTRDoorLockClusterSetHolidayScheduleParams")}
	})
	return MTRDoorLockClusterSetHolidayScheduleParamsClass
}

type _MTRDoorLockClusterSetHolidayScheduleParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterSetHolidayScheduleParams] class.
type IMTRDoorLockClusterSetHolidayScheduleParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetHolidayScheduleParams
type MTRDoorLockClusterSetHolidayScheduleParams struct {
	objectivec.Object
}

// MTRDoorLockClusterSetHolidayScheduleParamsFrom constructs a [MTRDoorLockClusterSetHolidayScheduleParams] from an unsafe.Pointer.
func MTRDoorLockClusterSetHolidayScheduleParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterSetHolidayScheduleParams {
	return MTRDoorLockClusterSetHolidayScheduleParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterSetHolidayScheduleParamsClass) Alloc() MTRDoorLockClusterSetHolidayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterSetHolidayScheduleParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterSetHolidayScheduleParamsClass) New() MTRDoorLockClusterSetHolidayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterSetHolidayScheduleParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) Init() MTRDoorLockClusterSetHolidayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterSetHolidayScheduleParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) Autorelease() MTRDoorLockClusterSetHolidayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterSetHolidayScheduleParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterSetHolidayScheduleParams creates a new MTRDoorLockClusterSetHolidayScheduleParams instance.
func NewMTRDoorLockClusterSetHolidayScheduleParams() MTRDoorLockClusterSetHolidayScheduleParams {
	return getMTRDoorLockClusterSetHolidayScheduleParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetholidayscheduleparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetholidayscheduleparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetholidayscheduleparams/localstarttime
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) LocalStartTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("localStartTime"))
	return rv
}


// SetLocalStartTime sets the value of the localStartTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetholidayscheduleparams/localstarttime
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) SetLocalStartTime(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocalStartTime:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetholidayscheduleparams/holidayindex
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) HolidayIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("holidayIndex"))
	return rv
}


// SetHolidayIndex sets the value of the holidayIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetholidayscheduleparams/holidayindex
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) SetHolidayIndex(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHolidayIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetholidayscheduleparams/localendtime
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) LocalEndTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("localEndTime"))
	return rv
}


// SetLocalEndTime sets the value of the localEndTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetholidayscheduleparams/localendtime
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) SetLocalEndTime(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocalEndTime:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetholidayscheduleparams/operatingmode
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) OperatingMode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("operatingMode"))
	return rv
}


// SetOperatingMode sets the value of the operatingMode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetholidayscheduleparams/operatingmode
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) SetOperatingMode(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperatingMode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetholidayscheduleparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetholidayscheduleparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



