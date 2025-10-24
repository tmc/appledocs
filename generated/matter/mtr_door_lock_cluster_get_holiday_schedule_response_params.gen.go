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
	// properties:
	HolidayIndex() objc.IObject /* cross-framework: NSNumber */
	SetHolidayIndex(value objc.IObject /* cross-framework: NSNumber */)
	LocalEndTime() objc.IObject /* cross-framework: NSNumber */
	SetLocalEndTime(value objc.IObject /* cross-framework: NSNumber */)
	LocalStartTime() objc.IObject /* cross-framework: NSNumber */
	SetLocalStartTime(value objc.IObject /* cross-framework: NSNumber */)
	OperatingMode() objc.IObject /* cross-framework: NSNumber */
	SetOperatingMode(value objc.IObject /* cross-framework: NSNumber */)
	Status() objc.IObject /* cross-framework: NSNumber */
	SetStatus(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetholidayscheduleresponseparams/holidayindex
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) HolidayIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("holidayIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetholidayscheduleresponseparams/holidayindex
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) SetHolidayIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHolidayIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetholidayscheduleresponseparams/localendtime
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) LocalEndTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("localEndTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetholidayscheduleresponseparams/localendtime
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) SetLocalEndTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocalEndTime:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetholidayscheduleresponseparams/localstarttime
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) LocalStartTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("localStartTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetholidayscheduleresponseparams/localstarttime
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) SetLocalStartTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocalStartTime:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetholidayscheduleresponseparams/operatingmode
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) OperatingMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("operatingMode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetholidayscheduleresponseparams/operatingmode
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) SetOperatingMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperatingMode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetholidayscheduleresponseparams/status
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetholidayscheduleresponseparams/status
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetholidayscheduleresponseparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetholidayscheduleresponseparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



