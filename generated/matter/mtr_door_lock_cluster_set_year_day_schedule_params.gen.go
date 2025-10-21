// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRDoorLockClusterSetYearDayScheduleParams] class.
var (
	MTRDoorLockClusterSetYearDayScheduleParamsClass     _MTRDoorLockClusterSetYearDayScheduleParamsClass
	MTRDoorLockClusterSetYearDayScheduleParamsClassOnce sync.Once
)

func getMTRDoorLockClusterSetYearDayScheduleParamsClass() _MTRDoorLockClusterSetYearDayScheduleParamsClass {
	MTRDoorLockClusterSetYearDayScheduleParamsClassOnce.Do(func() {
		MTRDoorLockClusterSetYearDayScheduleParamsClass = _MTRDoorLockClusterSetYearDayScheduleParamsClass{objc.GetClass("MTRDoorLockClusterSetYearDayScheduleParams")}
	})
	return MTRDoorLockClusterSetYearDayScheduleParamsClass
}

type _MTRDoorLockClusterSetYearDayScheduleParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterSetYearDayScheduleParams] class.
type IMTRDoorLockClusterSetYearDayScheduleParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetYearDayScheduleParams
type MTRDoorLockClusterSetYearDayScheduleParams struct {
	objectivec.Object
}

// MTRDoorLockClusterSetYearDayScheduleParamsFrom constructs a [MTRDoorLockClusterSetYearDayScheduleParams] from an unsafe.Pointer.
func MTRDoorLockClusterSetYearDayScheduleParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterSetYearDayScheduleParams {
	return MTRDoorLockClusterSetYearDayScheduleParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterSetYearDayScheduleParamsClass) Alloc() MTRDoorLockClusterSetYearDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterSetYearDayScheduleParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterSetYearDayScheduleParamsClass) New() MTRDoorLockClusterSetYearDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterSetYearDayScheduleParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterSetYearDayScheduleParams) Init() MTRDoorLockClusterSetYearDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterSetYearDayScheduleParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterSetYearDayScheduleParams) Autorelease() MTRDoorLockClusterSetYearDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterSetYearDayScheduleParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterSetYearDayScheduleParams creates a new MTRDoorLockClusterSetYearDayScheduleParams instance.
func NewMTRDoorLockClusterSetYearDayScheduleParams() MTRDoorLockClusterSetYearDayScheduleParams {
	return getMTRDoorLockClusterSetYearDayScheduleParamsClass().New()
}




