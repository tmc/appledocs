// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterGetWeekDayScheduleResponseParams] class.
var (
	MTRDoorLockClusterGetWeekDayScheduleResponseParamsClass     _MTRDoorLockClusterGetWeekDayScheduleResponseParamsClass
	MTRDoorLockClusterGetWeekDayScheduleResponseParamsClassOnce sync.Once
)

func getMTRDoorLockClusterGetWeekDayScheduleResponseParamsClass() _MTRDoorLockClusterGetWeekDayScheduleResponseParamsClass {
	MTRDoorLockClusterGetWeekDayScheduleResponseParamsClassOnce.Do(func() {
		MTRDoorLockClusterGetWeekDayScheduleResponseParamsClass = _MTRDoorLockClusterGetWeekDayScheduleResponseParamsClass{objc.GetClass("MTRDoorLockClusterGetWeekDayScheduleResponseParams")}
	})
	return MTRDoorLockClusterGetWeekDayScheduleResponseParamsClass
}

type _MTRDoorLockClusterGetWeekDayScheduleResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterGetWeekDayScheduleResponseParams] class.
type IMTRDoorLockClusterGetWeekDayScheduleResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetWeekDayScheduleResponseParams
type MTRDoorLockClusterGetWeekDayScheduleResponseParams struct {
	objectivec.Object
}

// MTRDoorLockClusterGetWeekDayScheduleResponseParamsFrom constructs a [MTRDoorLockClusterGetWeekDayScheduleResponseParams] from an unsafe.Pointer.
func MTRDoorLockClusterGetWeekDayScheduleResponseParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterGetWeekDayScheduleResponseParams {
	return MTRDoorLockClusterGetWeekDayScheduleResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterGetWeekDayScheduleResponseParamsClass) Alloc() MTRDoorLockClusterGetWeekDayScheduleResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetWeekDayScheduleResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterGetWeekDayScheduleResponseParamsClass) New() MTRDoorLockClusterGetWeekDayScheduleResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetWeekDayScheduleResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) Init() MTRDoorLockClusterGetWeekDayScheduleResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetWeekDayScheduleResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) Autorelease() MTRDoorLockClusterGetWeekDayScheduleResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetWeekDayScheduleResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterGetWeekDayScheduleResponseParams creates a new MTRDoorLockClusterGetWeekDayScheduleResponseParams instance.
func NewMTRDoorLockClusterGetWeekDayScheduleResponseParams() MTRDoorLockClusterGetWeekDayScheduleResponseParams {
	return getMTRDoorLockClusterGetWeekDayScheduleResponseParamsClass().New()
}




