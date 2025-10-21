// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRThermostatClusterWeeklyScheduleTransitionStruct] class.
var (
	MTRThermostatClusterWeeklyScheduleTransitionStructClass     _MTRThermostatClusterWeeklyScheduleTransitionStructClass
	MTRThermostatClusterWeeklyScheduleTransitionStructClassOnce sync.Once
)

func getMTRThermostatClusterWeeklyScheduleTransitionStructClass() _MTRThermostatClusterWeeklyScheduleTransitionStructClass {
	MTRThermostatClusterWeeklyScheduleTransitionStructClassOnce.Do(func() {
		MTRThermostatClusterWeeklyScheduleTransitionStructClass = _MTRThermostatClusterWeeklyScheduleTransitionStructClass{objc.GetClass("MTRThermostatClusterWeeklyScheduleTransitionStruct")}
	})
	return MTRThermostatClusterWeeklyScheduleTransitionStructClass
}

type _MTRThermostatClusterWeeklyScheduleTransitionStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRThermostatClusterWeeklyScheduleTransitionStruct] class.
type IMTRThermostatClusterWeeklyScheduleTransitionStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterWeeklyScheduleTransitionStruct
type MTRThermostatClusterWeeklyScheduleTransitionStruct struct {
	objectivec.Object
}

// MTRThermostatClusterWeeklyScheduleTransitionStructFrom constructs a [MTRThermostatClusterWeeklyScheduleTransitionStruct] from an unsafe.Pointer.
func MTRThermostatClusterWeeklyScheduleTransitionStructFrom(ptr unsafe.Pointer) MTRThermostatClusterWeeklyScheduleTransitionStruct {
	return MTRThermostatClusterWeeklyScheduleTransitionStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterWeeklyScheduleTransitionStructClass) Alloc() MTRThermostatClusterWeeklyScheduleTransitionStruct {
	rv := objc.Send[MTRThermostatClusterWeeklyScheduleTransitionStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThermostatClusterWeeklyScheduleTransitionStructClass) New() MTRThermostatClusterWeeklyScheduleTransitionStruct {
	rv := objc.Send[MTRThermostatClusterWeeklyScheduleTransitionStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThermostatClusterWeeklyScheduleTransitionStruct) Init() MTRThermostatClusterWeeklyScheduleTransitionStruct {
	rv := objc.Send[MTRThermostatClusterWeeklyScheduleTransitionStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThermostatClusterWeeklyScheduleTransitionStruct) Autorelease() MTRThermostatClusterWeeklyScheduleTransitionStruct {
	rv := objc.Send[MTRThermostatClusterWeeklyScheduleTransitionStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThermostatClusterWeeklyScheduleTransitionStruct creates a new MTRThermostatClusterWeeklyScheduleTransitionStruct instance.
func NewMTRThermostatClusterWeeklyScheduleTransitionStruct() MTRThermostatClusterWeeklyScheduleTransitionStruct {
	return getMTRThermostatClusterWeeklyScheduleTransitionStructClass().New()
}




