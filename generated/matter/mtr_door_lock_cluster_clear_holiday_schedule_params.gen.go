// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterClearHolidayScheduleParams] class.
var (
	MTRDoorLockClusterClearHolidayScheduleParamsClass     _MTRDoorLockClusterClearHolidayScheduleParamsClass
	MTRDoorLockClusterClearHolidayScheduleParamsClassOnce sync.Once
)

func getMTRDoorLockClusterClearHolidayScheduleParamsClass() _MTRDoorLockClusterClearHolidayScheduleParamsClass {
	MTRDoorLockClusterClearHolidayScheduleParamsClassOnce.Do(func() {
		MTRDoorLockClusterClearHolidayScheduleParamsClass = _MTRDoorLockClusterClearHolidayScheduleParamsClass{objc.GetClass("MTRDoorLockClusterClearHolidayScheduleParams")}
	})
	return MTRDoorLockClusterClearHolidayScheduleParamsClass
}

type _MTRDoorLockClusterClearHolidayScheduleParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterClearHolidayScheduleParams] class.
type IMTRDoorLockClusterClearHolidayScheduleParams interface {
	objectivec.IObject
	HolidayIndex() foundation.Number
	SetHolidayIndex(value foundation.INumber)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearHolidayScheduleParams
type MTRDoorLockClusterClearHolidayScheduleParams struct {
	objectivec.Object
}

// MTRDoorLockClusterClearHolidayScheduleParamsFrom constructs a [MTRDoorLockClusterClearHolidayScheduleParams] from an unsafe.Pointer.
func MTRDoorLockClusterClearHolidayScheduleParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterClearHolidayScheduleParams {
	return MTRDoorLockClusterClearHolidayScheduleParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterClearHolidayScheduleParamsClass) Alloc() MTRDoorLockClusterClearHolidayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterClearHolidayScheduleParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterClearHolidayScheduleParamsClass) New() MTRDoorLockClusterClearHolidayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterClearHolidayScheduleParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterClearHolidayScheduleParams) Init() MTRDoorLockClusterClearHolidayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterClearHolidayScheduleParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterClearHolidayScheduleParams) Autorelease() MTRDoorLockClusterClearHolidayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterClearHolidayScheduleParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterClearHolidayScheduleParams creates a new MTRDoorLockClusterClearHolidayScheduleParams instance.
func NewMTRDoorLockClusterClearHolidayScheduleParams() MTRDoorLockClusterClearHolidayScheduleParams {
	return getMTRDoorLockClusterClearHolidayScheduleParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearholidayscheduleparams/holidayindex
func (m_ MTRDoorLockClusterClearHolidayScheduleParams) HolidayIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("holidayIndex"))
	return rv
}


// SetHolidayIndex sets the value of the holidayIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearholidayscheduleparams/holidayindex
func (m_ MTRDoorLockClusterClearHolidayScheduleParams) SetHolidayIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHolidayIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearholidayscheduleparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterClearHolidayScheduleParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearholidayscheduleparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterClearHolidayScheduleParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearholidayscheduleparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterClearHolidayScheduleParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearholidayscheduleparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterClearHolidayScheduleParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



