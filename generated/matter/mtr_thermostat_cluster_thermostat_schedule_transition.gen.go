// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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




