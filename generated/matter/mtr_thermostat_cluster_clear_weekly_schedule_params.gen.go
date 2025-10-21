// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRThermostatClusterClearWeeklyScheduleParams] class.
var (
	MTRThermostatClusterClearWeeklyScheduleParamsClass     _MTRThermostatClusterClearWeeklyScheduleParamsClass
	MTRThermostatClusterClearWeeklyScheduleParamsClassOnce sync.Once
)

func getMTRThermostatClusterClearWeeklyScheduleParamsClass() _MTRThermostatClusterClearWeeklyScheduleParamsClass {
	MTRThermostatClusterClearWeeklyScheduleParamsClassOnce.Do(func() {
		MTRThermostatClusterClearWeeklyScheduleParamsClass = _MTRThermostatClusterClearWeeklyScheduleParamsClass{objc.GetClass("MTRThermostatClusterClearWeeklyScheduleParams")}
	})
	return MTRThermostatClusterClearWeeklyScheduleParamsClass
}

type _MTRThermostatClusterClearWeeklyScheduleParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRThermostatClusterClearWeeklyScheduleParams] class.
type IMTRThermostatClusterClearWeeklyScheduleParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterClearWeeklyScheduleParams
type MTRThermostatClusterClearWeeklyScheduleParams struct {
	objectivec.Object
}

// MTRThermostatClusterClearWeeklyScheduleParamsFrom constructs a [MTRThermostatClusterClearWeeklyScheduleParams] from an unsafe.Pointer.
func MTRThermostatClusterClearWeeklyScheduleParamsFrom(ptr unsafe.Pointer) MTRThermostatClusterClearWeeklyScheduleParams {
	return MTRThermostatClusterClearWeeklyScheduleParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterClearWeeklyScheduleParamsClass) Alloc() MTRThermostatClusterClearWeeklyScheduleParams {
	rv := objc.Send[MTRThermostatClusterClearWeeklyScheduleParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThermostatClusterClearWeeklyScheduleParamsClass) New() MTRThermostatClusterClearWeeklyScheduleParams {
	rv := objc.Send[MTRThermostatClusterClearWeeklyScheduleParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThermostatClusterClearWeeklyScheduleParams) Init() MTRThermostatClusterClearWeeklyScheduleParams {
	rv := objc.Send[MTRThermostatClusterClearWeeklyScheduleParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThermostatClusterClearWeeklyScheduleParams) Autorelease() MTRThermostatClusterClearWeeklyScheduleParams {
	rv := objc.Send[MTRThermostatClusterClearWeeklyScheduleParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThermostatClusterClearWeeklyScheduleParams creates a new MTRThermostatClusterClearWeeklyScheduleParams instance.
func NewMTRThermostatClusterClearWeeklyScheduleParams() MTRThermostatClusterClearWeeklyScheduleParams {
	return getMTRThermostatClusterClearWeeklyScheduleParamsClass().New()
}




