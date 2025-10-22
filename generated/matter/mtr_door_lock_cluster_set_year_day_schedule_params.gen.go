// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterSetYearDayScheduleParams] class.
var (
	MTRDoorLockClusterSetYearDayScheduleParamsClass     _MTRDoorLockClusterSetYearDayScheduleParamsClass
	MTRDoorLockClusterSetYearDayScheduleParamsClassOnce sync.Once
)

func getMTRDoorLockClusterSetYearDayScheduleParamsClass() _MTRDoorLockClusterSetYearDayScheduleParamsClass {
	MTRDoorLockClusterSetYearDayScheduleParamsClassOnce.Do(func() {
		MTRDoorLockClusterSetYearDayScheduleParamsClass = _MTRDoorLockClusterSetYearDayScheduleParamsClass{objc.GetClass("MTRDoorLockClusterSetYearDayScheduleParams")}
	})
	return MTRDoorLockClusterSetYearDayScheduleParamsClass
}

type _MTRDoorLockClusterSetYearDayScheduleParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterSetYearDayScheduleParams] class.
type IMTRDoorLockClusterSetYearDayScheduleParams interface {
	objectivec.IObject
	LocalEndTime() foundation.Number
	SetLocalEndTime(value foundation.INumber)
	LocalStartTime() foundation.Number
	SetLocalStartTime(value foundation.INumber)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
	UserIndex() foundation.Number
	SetUserIndex(value foundation.INumber)
	YearDayIndex() foundation.Number
	SetYearDayIndex(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetYearDayScheduleParams
type MTRDoorLockClusterSetYearDayScheduleParams struct {
	objectivec.Object
}

// MTRDoorLockClusterSetYearDayScheduleParamsFrom constructs a [MTRDoorLockClusterSetYearDayScheduleParams] from an unsafe.Pointer.
func MTRDoorLockClusterSetYearDayScheduleParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterSetYearDayScheduleParams {
	return MTRDoorLockClusterSetYearDayScheduleParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterSetYearDayScheduleParamsClass) Alloc() MTRDoorLockClusterSetYearDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterSetYearDayScheduleParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterSetYearDayScheduleParamsClass) New() MTRDoorLockClusterSetYearDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterSetYearDayScheduleParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterSetYearDayScheduleParams) Init() MTRDoorLockClusterSetYearDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterSetYearDayScheduleParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterSetYearDayScheduleParams) Autorelease() MTRDoorLockClusterSetYearDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterSetYearDayScheduleParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterSetYearDayScheduleParams creates a new MTRDoorLockClusterSetYearDayScheduleParams instance.
func NewMTRDoorLockClusterSetYearDayScheduleParams() MTRDoorLockClusterSetYearDayScheduleParams {
	return getMTRDoorLockClusterSetYearDayScheduleParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetyeardayscheduleparams/localendtime
func (m_ MTRDoorLockClusterSetYearDayScheduleParams) LocalEndTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("localEndTime"))
	return rv
}


// SetLocalEndTime sets the value of the localEndTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetyeardayscheduleparams/localendtime
func (m_ MTRDoorLockClusterSetYearDayScheduleParams) SetLocalEndTime(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocalEndTime:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetyeardayscheduleparams/localstarttime
func (m_ MTRDoorLockClusterSetYearDayScheduleParams) LocalStartTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("localStartTime"))
	return rv
}


// SetLocalStartTime sets the value of the localStartTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetyeardayscheduleparams/localstarttime
func (m_ MTRDoorLockClusterSetYearDayScheduleParams) SetLocalStartTime(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocalStartTime:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetyeardayscheduleparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterSetYearDayScheduleParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetyeardayscheduleparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterSetYearDayScheduleParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetyeardayscheduleparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterSetYearDayScheduleParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetyeardayscheduleparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterSetYearDayScheduleParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetyeardayscheduleparams/userindex
func (m_ MTRDoorLockClusterSetYearDayScheduleParams) UserIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("userIndex"))
	return rv
}


// SetUserIndex sets the value of the userIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetyeardayscheduleparams/userindex
func (m_ MTRDoorLockClusterSetYearDayScheduleParams) SetUserIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetyeardayscheduleparams/yeardayindex
func (m_ MTRDoorLockClusterSetYearDayScheduleParams) YearDayIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("yearDayIndex"))
	return rv
}


// SetYearDayIndex sets the value of the yearDayIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetyeardayscheduleparams/yeardayindex
func (m_ MTRDoorLockClusterSetYearDayScheduleParams) SetYearDayIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setYearDayIndex:"), value)
}



