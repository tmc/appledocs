// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterGetHolidayScheduleParams] class.
var (
	MTRDoorLockClusterGetHolidayScheduleParamsClass     _MTRDoorLockClusterGetHolidayScheduleParamsClass
	MTRDoorLockClusterGetHolidayScheduleParamsClassOnce sync.Once
)

func getMTRDoorLockClusterGetHolidayScheduleParamsClass() _MTRDoorLockClusterGetHolidayScheduleParamsClass {
	MTRDoorLockClusterGetHolidayScheduleParamsClassOnce.Do(func() {
		MTRDoorLockClusterGetHolidayScheduleParamsClass = _MTRDoorLockClusterGetHolidayScheduleParamsClass{objc.GetClass("MTRDoorLockClusterGetHolidayScheduleParams")}
	})
	return MTRDoorLockClusterGetHolidayScheduleParamsClass
}

type _MTRDoorLockClusterGetHolidayScheduleParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterGetHolidayScheduleParams] class.
type IMTRDoorLockClusterGetHolidayScheduleParams interface {
	objectivec.IObject
	HolidayIndex() foundation.Number
	SetHolidayIndex(value foundation.INumber)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetHolidayScheduleParams
type MTRDoorLockClusterGetHolidayScheduleParams struct {
	objectivec.Object
}

// MTRDoorLockClusterGetHolidayScheduleParamsFrom constructs a [MTRDoorLockClusterGetHolidayScheduleParams] from an unsafe.Pointer.
func MTRDoorLockClusterGetHolidayScheduleParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterGetHolidayScheduleParams {
	return MTRDoorLockClusterGetHolidayScheduleParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterGetHolidayScheduleParamsClass) Alloc() MTRDoorLockClusterGetHolidayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterGetHolidayScheduleParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterGetHolidayScheduleParamsClass) New() MTRDoorLockClusterGetHolidayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterGetHolidayScheduleParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterGetHolidayScheduleParams) Init() MTRDoorLockClusterGetHolidayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterGetHolidayScheduleParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterGetHolidayScheduleParams) Autorelease() MTRDoorLockClusterGetHolidayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterGetHolidayScheduleParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterGetHolidayScheduleParams creates a new MTRDoorLockClusterGetHolidayScheduleParams instance.
func NewMTRDoorLockClusterGetHolidayScheduleParams() MTRDoorLockClusterGetHolidayScheduleParams {
	return getMTRDoorLockClusterGetHolidayScheduleParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetholidayscheduleparams/holidayindex
func (m_ MTRDoorLockClusterGetHolidayScheduleParams) HolidayIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("holidayIndex"))
	return rv
}


// SetHolidayIndex sets the value of the holidayIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetholidayscheduleparams/holidayindex
func (m_ MTRDoorLockClusterGetHolidayScheduleParams) SetHolidayIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHolidayIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetholidayscheduleparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterGetHolidayScheduleParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetholidayscheduleparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterGetHolidayScheduleParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetholidayscheduleparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterGetHolidayScheduleParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetholidayscheduleparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterGetHolidayScheduleParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



