// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent] class.
var (
	MTRPumpConfigurationAndControlClusterSupplyVoltageLowEventClass     _MTRPumpConfigurationAndControlClusterSupplyVoltageLowEventClass
	MTRPumpConfigurationAndControlClusterSupplyVoltageLowEventClassOnce sync.Once
)

func getMTRPumpConfigurationAndControlClusterSupplyVoltageLowEventClass() _MTRPumpConfigurationAndControlClusterSupplyVoltageLowEventClass {
	MTRPumpConfigurationAndControlClusterSupplyVoltageLowEventClassOnce.Do(func() {
		MTRPumpConfigurationAndControlClusterSupplyVoltageLowEventClass = _MTRPumpConfigurationAndControlClusterSupplyVoltageLowEventClass{objc.GetClass("MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent")}
	})
	return MTRPumpConfigurationAndControlClusterSupplyVoltageLowEventClass
}

type _MTRPumpConfigurationAndControlClusterSupplyVoltageLowEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent] class.
type IMTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent
type MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent struct {
	objectivec.Object
}

// MTRPumpConfigurationAndControlClusterSupplyVoltageLowEventFrom constructs a [MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent] from an unsafe.Pointer.
func MTRPumpConfigurationAndControlClusterSupplyVoltageLowEventFrom(ptr unsafe.Pointer) MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent {
	return MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRPumpConfigurationAndControlClusterSupplyVoltageLowEventClass) Alloc() MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRPumpConfigurationAndControlClusterSupplyVoltageLowEventClass) New() MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent) Init() MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent) Autorelease() MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent creates a new MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent instance.
func NewMTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent() MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent {
	return getMTRPumpConfigurationAndControlClusterSupplyVoltageLowEventClass().New()
}




