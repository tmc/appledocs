// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterSetWeekDayScheduleParams] class.
var (
	MTRDoorLockClusterSetWeekDayScheduleParamsClass     _MTRDoorLockClusterSetWeekDayScheduleParamsClass
	MTRDoorLockClusterSetWeekDayScheduleParamsClassOnce sync.Once
)

func getMTRDoorLockClusterSetWeekDayScheduleParamsClass() _MTRDoorLockClusterSetWeekDayScheduleParamsClass {
	MTRDoorLockClusterSetWeekDayScheduleParamsClassOnce.Do(func() {
		MTRDoorLockClusterSetWeekDayScheduleParamsClass = _MTRDoorLockClusterSetWeekDayScheduleParamsClass{objc.GetClass("MTRDoorLockClusterSetWeekDayScheduleParams")}
	})
	return MTRDoorLockClusterSetWeekDayScheduleParamsClass
}

type _MTRDoorLockClusterSetWeekDayScheduleParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterSetWeekDayScheduleParams] class.
type IMTRDoorLockClusterSetWeekDayScheduleParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetWeekDayScheduleParams
type MTRDoorLockClusterSetWeekDayScheduleParams struct {
	objectivec.Object
}

// MTRDoorLockClusterSetWeekDayScheduleParamsFrom constructs a [MTRDoorLockClusterSetWeekDayScheduleParams] from an unsafe.Pointer.
func MTRDoorLockClusterSetWeekDayScheduleParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterSetWeekDayScheduleParams {
	return MTRDoorLockClusterSetWeekDayScheduleParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterSetWeekDayScheduleParamsClass) Alloc() MTRDoorLockClusterSetWeekDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterSetWeekDayScheduleParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterSetWeekDayScheduleParamsClass) New() MTRDoorLockClusterSetWeekDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterSetWeekDayScheduleParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) Init() MTRDoorLockClusterSetWeekDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterSetWeekDayScheduleParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) Autorelease() MTRDoorLockClusterSetWeekDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterSetWeekDayScheduleParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterSetWeekDayScheduleParams creates a new MTRDoorLockClusterSetWeekDayScheduleParams instance.
func NewMTRDoorLockClusterSetWeekDayScheduleParams() MTRDoorLockClusterSetWeekDayScheduleParams {
	return getMTRDoorLockClusterSetWeekDayScheduleParamsClass().New()
}




