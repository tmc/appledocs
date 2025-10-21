// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
}

//
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




