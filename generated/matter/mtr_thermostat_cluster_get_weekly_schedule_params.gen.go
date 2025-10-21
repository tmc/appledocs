// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRThermostatClusterGetWeeklyScheduleParams] class.
var (
	MTRThermostatClusterGetWeeklyScheduleParamsClass     _MTRThermostatClusterGetWeeklyScheduleParamsClass
	MTRThermostatClusterGetWeeklyScheduleParamsClassOnce sync.Once
)

func getMTRThermostatClusterGetWeeklyScheduleParamsClass() _MTRThermostatClusterGetWeeklyScheduleParamsClass {
	MTRThermostatClusterGetWeeklyScheduleParamsClassOnce.Do(func() {
		MTRThermostatClusterGetWeeklyScheduleParamsClass = _MTRThermostatClusterGetWeeklyScheduleParamsClass{objc.GetClass("MTRThermostatClusterGetWeeklyScheduleParams")}
	})
	return MTRThermostatClusterGetWeeklyScheduleParamsClass
}

type _MTRThermostatClusterGetWeeklyScheduleParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRThermostatClusterGetWeeklyScheduleParams] class.
type IMTRThermostatClusterGetWeeklyScheduleParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterGetWeeklyScheduleParams
type MTRThermostatClusterGetWeeklyScheduleParams struct {
	objectivec.Object
}

// MTRThermostatClusterGetWeeklyScheduleParamsFrom constructs a [MTRThermostatClusterGetWeeklyScheduleParams] from an unsafe.Pointer.
func MTRThermostatClusterGetWeeklyScheduleParamsFrom(ptr unsafe.Pointer) MTRThermostatClusterGetWeeklyScheduleParams {
	return MTRThermostatClusterGetWeeklyScheduleParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterGetWeeklyScheduleParamsClass) Alloc() MTRThermostatClusterGetWeeklyScheduleParams {
	rv := objc.Send[MTRThermostatClusterGetWeeklyScheduleParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThermostatClusterGetWeeklyScheduleParamsClass) New() MTRThermostatClusterGetWeeklyScheduleParams {
	rv := objc.Send[MTRThermostatClusterGetWeeklyScheduleParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThermostatClusterGetWeeklyScheduleParams) Init() MTRThermostatClusterGetWeeklyScheduleParams {
	rv := objc.Send[MTRThermostatClusterGetWeeklyScheduleParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThermostatClusterGetWeeklyScheduleParams) Autorelease() MTRThermostatClusterGetWeeklyScheduleParams {
	rv := objc.Send[MTRThermostatClusterGetWeeklyScheduleParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThermostatClusterGetWeeklyScheduleParams creates a new MTRThermostatClusterGetWeeklyScheduleParams instance.
func NewMTRThermostatClusterGetWeeklyScheduleParams() MTRThermostatClusterGetWeeklyScheduleParams {
	return getMTRThermostatClusterGetWeeklyScheduleParamsClass().New()
}




