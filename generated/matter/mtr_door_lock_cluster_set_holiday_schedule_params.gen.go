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
	// properties:
	HolidayIndex() objc.IObject /* cross-framework: NSNumber */
	SetHolidayIndex(value objc.IObject /* cross-framework: NSNumber */)
	LocalEndTime() objc.IObject /* cross-framework: NSNumber */
	SetLocalEndTime(value objc.IObject /* cross-framework: NSNumber */)
	LocalStartTime() objc.IObject /* cross-framework: NSNumber */
	SetLocalStartTime(value objc.IObject /* cross-framework: NSNumber */)
	OperatingMode() objc.IObject /* cross-framework: NSNumber */
	SetOperatingMode(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetholidayscheduleparams/holidayindex
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) HolidayIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("holidayIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetholidayscheduleparams/holidayindex
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) SetHolidayIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHolidayIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetholidayscheduleparams/localendtime
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) LocalEndTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("localEndTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetholidayscheduleparams/localendtime
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) SetLocalEndTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocalEndTime:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetholidayscheduleparams/localstarttime
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) LocalStartTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("localStartTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetholidayscheduleparams/localstarttime
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) SetLocalStartTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocalStartTime:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetholidayscheduleparams/operatingmode
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) OperatingMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("operatingMode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetholidayscheduleparams/operatingmode
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) SetOperatingMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperatingMode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetholidayscheduleparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetholidayscheduleparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetholidayscheduleparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetholidayscheduleparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



