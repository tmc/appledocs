// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRThermostatClusterGetWeeklyScheduleResponseParams] class.
var (
	MTRThermostatClusterGetWeeklyScheduleResponseParamsClass     _MTRThermostatClusterGetWeeklyScheduleResponseParamsClass
	MTRThermostatClusterGetWeeklyScheduleResponseParamsClassOnce sync.Once
)

func getMTRThermostatClusterGetWeeklyScheduleResponseParamsClass() _MTRThermostatClusterGetWeeklyScheduleResponseParamsClass {
	MTRThermostatClusterGetWeeklyScheduleResponseParamsClassOnce.Do(func() {
		MTRThermostatClusterGetWeeklyScheduleResponseParamsClass = _MTRThermostatClusterGetWeeklyScheduleResponseParamsClass{objc.GetClass("MTRThermostatClusterGetWeeklyScheduleResponseParams")}
	})
	return MTRThermostatClusterGetWeeklyScheduleResponseParamsClass
}

type _MTRThermostatClusterGetWeeklyScheduleResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRThermostatClusterGetWeeklyScheduleResponseParams] class.
type IMTRThermostatClusterGetWeeklyScheduleResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterGetWeeklyScheduleResponseParams
type MTRThermostatClusterGetWeeklyScheduleResponseParams struct {
	objectivec.Object
}

// MTRThermostatClusterGetWeeklyScheduleResponseParamsFrom constructs a [MTRThermostatClusterGetWeeklyScheduleResponseParams] from an unsafe.Pointer.
func MTRThermostatClusterGetWeeklyScheduleResponseParamsFrom(ptr unsafe.Pointer) MTRThermostatClusterGetWeeklyScheduleResponseParams {
	return MTRThermostatClusterGetWeeklyScheduleResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterGetWeeklyScheduleResponseParamsClass) Alloc() MTRThermostatClusterGetWeeklyScheduleResponseParams {
	rv := objc.Send[MTRThermostatClusterGetWeeklyScheduleResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThermostatClusterGetWeeklyScheduleResponseParamsClass) New() MTRThermostatClusterGetWeeklyScheduleResponseParams {
	rv := objc.Send[MTRThermostatClusterGetWeeklyScheduleResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThermostatClusterGetWeeklyScheduleResponseParams) Init() MTRThermostatClusterGetWeeklyScheduleResponseParams {
	rv := objc.Send[MTRThermostatClusterGetWeeklyScheduleResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThermostatClusterGetWeeklyScheduleResponseParams) Autorelease() MTRThermostatClusterGetWeeklyScheduleResponseParams {
	rv := objc.Send[MTRThermostatClusterGetWeeklyScheduleResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThermostatClusterGetWeeklyScheduleResponseParams creates a new MTRThermostatClusterGetWeeklyScheduleResponseParams instance.
func NewMTRThermostatClusterGetWeeklyScheduleResponseParams() MTRThermostatClusterGetWeeklyScheduleResponseParams {
	return getMTRThermostatClusterGetWeeklyScheduleResponseParamsClass().New()
}




