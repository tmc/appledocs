// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterGetYearDayScheduleParams] class.
var (
	MTRDoorLockClusterGetYearDayScheduleParamsClass     _MTRDoorLockClusterGetYearDayScheduleParamsClass
	MTRDoorLockClusterGetYearDayScheduleParamsClassOnce sync.Once
)

func getMTRDoorLockClusterGetYearDayScheduleParamsClass() _MTRDoorLockClusterGetYearDayScheduleParamsClass {
	MTRDoorLockClusterGetYearDayScheduleParamsClassOnce.Do(func() {
		MTRDoorLockClusterGetYearDayScheduleParamsClass = _MTRDoorLockClusterGetYearDayScheduleParamsClass{objc.GetClass("MTRDoorLockClusterGetYearDayScheduleParams")}
	})
	return MTRDoorLockClusterGetYearDayScheduleParamsClass
}

type _MTRDoorLockClusterGetYearDayScheduleParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterGetYearDayScheduleParams] class.
type IMTRDoorLockClusterGetYearDayScheduleParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetYearDayScheduleParams
type MTRDoorLockClusterGetYearDayScheduleParams struct {
	objectivec.Object
}

// MTRDoorLockClusterGetYearDayScheduleParamsFrom constructs a [MTRDoorLockClusterGetYearDayScheduleParams] from an unsafe.Pointer.
func MTRDoorLockClusterGetYearDayScheduleParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterGetYearDayScheduleParams {
	return MTRDoorLockClusterGetYearDayScheduleParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterGetYearDayScheduleParamsClass) Alloc() MTRDoorLockClusterGetYearDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterGetYearDayScheduleParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterGetYearDayScheduleParamsClass) New() MTRDoorLockClusterGetYearDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterGetYearDayScheduleParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterGetYearDayScheduleParams) Init() MTRDoorLockClusterGetYearDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterGetYearDayScheduleParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterGetYearDayScheduleParams) Autorelease() MTRDoorLockClusterGetYearDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterGetYearDayScheduleParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterGetYearDayScheduleParams creates a new MTRDoorLockClusterGetYearDayScheduleParams instance.
func NewMTRDoorLockClusterGetYearDayScheduleParams() MTRDoorLockClusterGetYearDayScheduleParams {
	return getMTRDoorLockClusterGetYearDayScheduleParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetyeardayscheduleparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterGetYearDayScheduleParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetyeardayscheduleparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterGetYearDayScheduleParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetyeardayscheduleparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterGetYearDayScheduleParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetyeardayscheduleparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterGetYearDayScheduleParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetyeardayscheduleparams/userindex
func (m_ MTRDoorLockClusterGetYearDayScheduleParams) UserIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("userIndex"))
	return rv
}


// SetUserIndex sets the value of the userIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetyeardayscheduleparams/userindex
func (m_ MTRDoorLockClusterGetYearDayScheduleParams) SetUserIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetyeardayscheduleparams/yeardayindex
func (m_ MTRDoorLockClusterGetYearDayScheduleParams) YearDayIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("yearDayIndex"))
	return rv
}


// SetYearDayIndex sets the value of the yearDayIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetyeardayscheduleparams/yeardayindex
func (m_ MTRDoorLockClusterGetYearDayScheduleParams) SetYearDayIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setYearDayIndex:"), value)
}



