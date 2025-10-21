// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRThermostatClusterThermostatScheduleTransition] class.
var (
	MTRThermostatClusterThermostatScheduleTransitionClass     _MTRThermostatClusterThermostatScheduleTransitionClass
	MTRThermostatClusterThermostatScheduleTransitionClassOnce sync.Once
)

func getMTRThermostatClusterThermostatScheduleTransitionClass() _MTRThermostatClusterThermostatScheduleTransitionClass {
	MTRThermostatClusterThermostatScheduleTransitionClassOnce.Do(func() {
		MTRThermostatClusterThermostatScheduleTransitionClass = _MTRThermostatClusterThermostatScheduleTransitionClass{objc.GetClass("MTRThermostatClusterThermostatScheduleTransition")}
	})
	return MTRThermostatClusterThermostatScheduleTransitionClass
}

type _MTRThermostatClusterThermostatScheduleTransitionClass struct {
	class objc.Class
}

// An interface definition for the [MTRThermostatClusterThermostatScheduleTransition] class.
type IMTRThermostatClusterThermostatScheduleTransition interface {
	IMTRThermostatClusterWeeklyScheduleTransitionStruct
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterThermostatScheduleTransition
type MTRThermostatClusterThermostatScheduleTransition struct {
	MTRThermostatClusterWeeklyScheduleTransitionStruct
}

// MTRThermostatClusterThermostatScheduleTransitionFrom constructs a [MTRThermostatClusterThermostatScheduleTransition] from an unsafe.Pointer.
func MTRThermostatClusterThermostatScheduleTransitionFrom(ptr unsafe.Pointer) MTRThermostatClusterThermostatScheduleTransition {
	return MTRThermostatClusterThermostatScheduleTransition{
		MTRThermostatClusterWeeklyScheduleTransitionStruct: MTRThermostatClusterWeeklyScheduleTransitionStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterThermostatScheduleTransitionClass) Alloc() MTRThermostatClusterThermostatScheduleTransition {
	rv := objc.Send[MTRThermostatClusterThermostatScheduleTransition](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThermostatClusterThermostatScheduleTransitionClass) New() MTRThermostatClusterThermostatScheduleTransition {
	rv := objc.Send[MTRThermostatClusterThermostatScheduleTransition](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThermostatClusterThermostatScheduleTransition) Init() MTRThermostatClusterThermostatScheduleTransition {
	rv := objc.Send[MTRThermostatClusterThermostatScheduleTransition](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThermostatClusterThermostatScheduleTransition) Autorelease() MTRThermostatClusterThermostatScheduleTransition {
	rv := objc.Send[MTRThermostatClusterThermostatScheduleTransition](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThermostatClusterThermostatScheduleTransition creates a new MTRThermostatClusterThermostatScheduleTransition instance.
func NewMTRThermostatClusterThermostatScheduleTransition() MTRThermostatClusterThermostatScheduleTransition {
	return getMTRThermostatClusterThermostatScheduleTransitionClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterthermostatscheduletransition/coolsetpoint
func (m_ MTRThermostatClusterThermostatScheduleTransition) CoolSetpoint() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("coolSetpoint"))
	return rv
}


// SetCoolSetpoint sets the value of the coolSetpoint property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterthermostatscheduletransition/coolsetpoint
func (m_ MTRThermostatClusterThermostatScheduleTransition) SetCoolSetpoint(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCoolSetpoint:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterthermostatscheduletransition/heatsetpoint
func (m_ MTRThermostatClusterThermostatScheduleTransition) HeatSetpoint() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("heatSetpoint"))
	return rv
}


// SetHeatSetpoint sets the value of the heatSetpoint property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterthermostatscheduletransition/heatsetpoint
func (m_ MTRThermostatClusterThermostatScheduleTransition) SetHeatSetpoint(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHeatSetpoint:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterthermostatscheduletransition/transitiontime
func (m_ MTRThermostatClusterThermostatScheduleTransition) TransitionTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("transitionTime"))
	return rv
}


// SetTransitionTime sets the value of the transitionTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterthermostatscheduletransition/transitiontime
func (m_ MTRThermostatClusterThermostatScheduleTransition) SetTransitionTime(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}



