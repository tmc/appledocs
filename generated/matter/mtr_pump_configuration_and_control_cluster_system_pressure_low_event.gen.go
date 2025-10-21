// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRPumpConfigurationAndControlClusterSystemPressureLowEvent] class.
var (
	MTRPumpConfigurationAndControlClusterSystemPressureLowEventClass     _MTRPumpConfigurationAndControlClusterSystemPressureLowEventClass
	MTRPumpConfigurationAndControlClusterSystemPressureLowEventClassOnce sync.Once
)

func getMTRPumpConfigurationAndControlClusterSystemPressureLowEventClass() _MTRPumpConfigurationAndControlClusterSystemPressureLowEventClass {
	MTRPumpConfigurationAndControlClusterSystemPressureLowEventClassOnce.Do(func() {
		MTRPumpConfigurationAndControlClusterSystemPressureLowEventClass = _MTRPumpConfigurationAndControlClusterSystemPressureLowEventClass{objc.GetClass("MTRPumpConfigurationAndControlClusterSystemPressureLowEvent")}
	})
	return MTRPumpConfigurationAndControlClusterSystemPressureLowEventClass
}

type _MTRPumpConfigurationAndControlClusterSystemPressureLowEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRPumpConfigurationAndControlClusterSystemPressureLowEvent] class.
type IMTRPumpConfigurationAndControlClusterSystemPressureLowEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPumpConfigurationAndControlClusterSystemPressureLowEvent
type MTRPumpConfigurationAndControlClusterSystemPressureLowEvent struct {
	objectivec.Object
}

// MTRPumpConfigurationAndControlClusterSystemPressureLowEventFrom constructs a [MTRPumpConfigurationAndControlClusterSystemPressureLowEvent] from an unsafe.Pointer.
func MTRPumpConfigurationAndControlClusterSystemPressureLowEventFrom(ptr unsafe.Pointer) MTRPumpConfigurationAndControlClusterSystemPressureLowEvent {
	return MTRPumpConfigurationAndControlClusterSystemPressureLowEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRPumpConfigurationAndControlClusterSystemPressureLowEventClass) Alloc() MTRPumpConfigurationAndControlClusterSystemPressureLowEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSystemPressureLowEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRPumpConfigurationAndControlClusterSystemPressureLowEventClass) New() MTRPumpConfigurationAndControlClusterSystemPressureLowEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSystemPressureLowEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPumpConfigurationAndControlClusterSystemPressureLowEvent) Init() MTRPumpConfigurationAndControlClusterSystemPressureLowEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSystemPressureLowEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPumpConfigurationAndControlClusterSystemPressureLowEvent) Autorelease() MTRPumpConfigurationAndControlClusterSystemPressureLowEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSystemPressureLowEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPumpConfigurationAndControlClusterSystemPressureLowEvent creates a new MTRPumpConfigurationAndControlClusterSystemPressureLowEvent instance.
func NewMTRPumpConfigurationAndControlClusterSystemPressureLowEvent() MTRPumpConfigurationAndControlClusterSystemPressureLowEvent {
	return getMTRPumpConfigurationAndControlClusterSystemPressureLowEventClass().New()
}




