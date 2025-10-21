// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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




