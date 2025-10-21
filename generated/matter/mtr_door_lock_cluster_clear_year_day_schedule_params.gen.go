// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRDoorLockClusterClearYearDayScheduleParams] class.
var (
	MTRDoorLockClusterClearYearDayScheduleParamsClass     _MTRDoorLockClusterClearYearDayScheduleParamsClass
	MTRDoorLockClusterClearYearDayScheduleParamsClassOnce sync.Once
)

func getMTRDoorLockClusterClearYearDayScheduleParamsClass() _MTRDoorLockClusterClearYearDayScheduleParamsClass {
	MTRDoorLockClusterClearYearDayScheduleParamsClassOnce.Do(func() {
		MTRDoorLockClusterClearYearDayScheduleParamsClass = _MTRDoorLockClusterClearYearDayScheduleParamsClass{objc.GetClass("MTRDoorLockClusterClearYearDayScheduleParams")}
	})
	return MTRDoorLockClusterClearYearDayScheduleParamsClass
}

type _MTRDoorLockClusterClearYearDayScheduleParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterClearYearDayScheduleParams] class.
type IMTRDoorLockClusterClearYearDayScheduleParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearYearDayScheduleParams
type MTRDoorLockClusterClearYearDayScheduleParams struct {
	objectivec.Object
}

// MTRDoorLockClusterClearYearDayScheduleParamsFrom constructs a [MTRDoorLockClusterClearYearDayScheduleParams] from an unsafe.Pointer.
func MTRDoorLockClusterClearYearDayScheduleParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterClearYearDayScheduleParams {
	return MTRDoorLockClusterClearYearDayScheduleParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterClearYearDayScheduleParamsClass) Alloc() MTRDoorLockClusterClearYearDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterClearYearDayScheduleParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterClearYearDayScheduleParamsClass) New() MTRDoorLockClusterClearYearDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterClearYearDayScheduleParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterClearYearDayScheduleParams) Init() MTRDoorLockClusterClearYearDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterClearYearDayScheduleParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterClearYearDayScheduleParams) Autorelease() MTRDoorLockClusterClearYearDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterClearYearDayScheduleParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterClearYearDayScheduleParams creates a new MTRDoorLockClusterClearYearDayScheduleParams instance.
func NewMTRDoorLockClusterClearYearDayScheduleParams() MTRDoorLockClusterClearYearDayScheduleParams {
	return getMTRDoorLockClusterClearYearDayScheduleParamsClass().New()
}




