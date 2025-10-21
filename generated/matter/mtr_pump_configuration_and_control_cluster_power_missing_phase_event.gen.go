// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent] class.
var (
	MTRPumpConfigurationAndControlClusterPowerMissingPhaseEventClass     _MTRPumpConfigurationAndControlClusterPowerMissingPhaseEventClass
	MTRPumpConfigurationAndControlClusterPowerMissingPhaseEventClassOnce sync.Once
)

func getMTRPumpConfigurationAndControlClusterPowerMissingPhaseEventClass() _MTRPumpConfigurationAndControlClusterPowerMissingPhaseEventClass {
	MTRPumpConfigurationAndControlClusterPowerMissingPhaseEventClassOnce.Do(func() {
		MTRPumpConfigurationAndControlClusterPowerMissingPhaseEventClass = _MTRPumpConfigurationAndControlClusterPowerMissingPhaseEventClass{objc.GetClass("MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent")}
	})
	return MTRPumpConfigurationAndControlClusterPowerMissingPhaseEventClass
}

type _MTRPumpConfigurationAndControlClusterPowerMissingPhaseEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent] class.
type IMTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent
type MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent struct {
	objectivec.Object
}

// MTRPumpConfigurationAndControlClusterPowerMissingPhaseEventFrom constructs a [MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent] from an unsafe.Pointer.
func MTRPumpConfigurationAndControlClusterPowerMissingPhaseEventFrom(ptr unsafe.Pointer) MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent {
	return MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRPumpConfigurationAndControlClusterPowerMissingPhaseEventClass) Alloc() MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRPumpConfigurationAndControlClusterPowerMissingPhaseEventClass) New() MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent) Init() MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent) Autorelease() MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent creates a new MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent instance.
func NewMTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent() MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent {
	return getMTRPumpConfigurationAndControlClusterPowerMissingPhaseEventClass().New()
}




