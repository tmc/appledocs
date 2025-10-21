// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRThermostatClusterSetWeeklyScheduleParams] class.
var (
	MTRThermostatClusterSetWeeklyScheduleParamsClass     _MTRThermostatClusterSetWeeklyScheduleParamsClass
	MTRThermostatClusterSetWeeklyScheduleParamsClassOnce sync.Once
)

func getMTRThermostatClusterSetWeeklyScheduleParamsClass() _MTRThermostatClusterSetWeeklyScheduleParamsClass {
	MTRThermostatClusterSetWeeklyScheduleParamsClassOnce.Do(func() {
		MTRThermostatClusterSetWeeklyScheduleParamsClass = _MTRThermostatClusterSetWeeklyScheduleParamsClass{objc.GetClass("MTRThermostatClusterSetWeeklyScheduleParams")}
	})
	return MTRThermostatClusterSetWeeklyScheduleParamsClass
}

type _MTRThermostatClusterSetWeeklyScheduleParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRThermostatClusterSetWeeklyScheduleParams] class.
type IMTRThermostatClusterSetWeeklyScheduleParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetWeeklyScheduleParams
type MTRThermostatClusterSetWeeklyScheduleParams struct {
	objectivec.Object
}

// MTRThermostatClusterSetWeeklyScheduleParamsFrom constructs a [MTRThermostatClusterSetWeeklyScheduleParams] from an unsafe.Pointer.
func MTRThermostatClusterSetWeeklyScheduleParamsFrom(ptr unsafe.Pointer) MTRThermostatClusterSetWeeklyScheduleParams {
	return MTRThermostatClusterSetWeeklyScheduleParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterSetWeeklyScheduleParamsClass) Alloc() MTRThermostatClusterSetWeeklyScheduleParams {
	rv := objc.Send[MTRThermostatClusterSetWeeklyScheduleParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThermostatClusterSetWeeklyScheduleParamsClass) New() MTRThermostatClusterSetWeeklyScheduleParams {
	rv := objc.Send[MTRThermostatClusterSetWeeklyScheduleParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThermostatClusterSetWeeklyScheduleParams) Init() MTRThermostatClusterSetWeeklyScheduleParams {
	rv := objc.Send[MTRThermostatClusterSetWeeklyScheduleParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThermostatClusterSetWeeklyScheduleParams) Autorelease() MTRThermostatClusterSetWeeklyScheduleParams {
	rv := objc.Send[MTRThermostatClusterSetWeeklyScheduleParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThermostatClusterSetWeeklyScheduleParams creates a new MTRThermostatClusterSetWeeklyScheduleParams instance.
func NewMTRThermostatClusterSetWeeklyScheduleParams() MTRThermostatClusterSetWeeklyScheduleParams {
	return getMTRThermostatClusterSetWeeklyScheduleParamsClass().New()
}




