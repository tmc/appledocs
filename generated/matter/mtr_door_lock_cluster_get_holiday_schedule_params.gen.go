// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRDoorLockClusterGetHolidayScheduleParams] class.
var (
	MTRDoorLockClusterGetHolidayScheduleParamsClass     _MTRDoorLockClusterGetHolidayScheduleParamsClass
	MTRDoorLockClusterGetHolidayScheduleParamsClassOnce sync.Once
)

func getMTRDoorLockClusterGetHolidayScheduleParamsClass() _MTRDoorLockClusterGetHolidayScheduleParamsClass {
	MTRDoorLockClusterGetHolidayScheduleParamsClassOnce.Do(func() {
		MTRDoorLockClusterGetHolidayScheduleParamsClass = _MTRDoorLockClusterGetHolidayScheduleParamsClass{objc.GetClass("MTRDoorLockClusterGetHolidayScheduleParams")}
	})
	return MTRDoorLockClusterGetHolidayScheduleParamsClass
}

type _MTRDoorLockClusterGetHolidayScheduleParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterGetHolidayScheduleParams] class.
type IMTRDoorLockClusterGetHolidayScheduleParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetHolidayScheduleParams
type MTRDoorLockClusterGetHolidayScheduleParams struct {
	objectivec.Object
}

// MTRDoorLockClusterGetHolidayScheduleParamsFrom constructs a [MTRDoorLockClusterGetHolidayScheduleParams] from an unsafe.Pointer.
func MTRDoorLockClusterGetHolidayScheduleParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterGetHolidayScheduleParams {
	return MTRDoorLockClusterGetHolidayScheduleParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterGetHolidayScheduleParamsClass) Alloc() MTRDoorLockClusterGetHolidayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterGetHolidayScheduleParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterGetHolidayScheduleParamsClass) New() MTRDoorLockClusterGetHolidayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterGetHolidayScheduleParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterGetHolidayScheduleParams) Init() MTRDoorLockClusterGetHolidayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterGetHolidayScheduleParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterGetHolidayScheduleParams) Autorelease() MTRDoorLockClusterGetHolidayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterGetHolidayScheduleParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterGetHolidayScheduleParams creates a new MTRDoorLockClusterGetHolidayScheduleParams instance.
func NewMTRDoorLockClusterGetHolidayScheduleParams() MTRDoorLockClusterGetHolidayScheduleParams {
	return getMTRDoorLockClusterGetHolidayScheduleParamsClass().New()
}




