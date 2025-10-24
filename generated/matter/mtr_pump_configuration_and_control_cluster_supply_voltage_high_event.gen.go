// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent] class.
var (
	MTRPumpConfigurationAndControlClusterSupplyVoltageHighEventClass     _MTRPumpConfigurationAndControlClusterSupplyVoltageHighEventClass
	MTRPumpConfigurationAndControlClusterSupplyVoltageHighEventClassOnce sync.Once
)

func getMTRPumpConfigurationAndControlClusterSupplyVoltageHighEventClass() _MTRPumpConfigurationAndControlClusterSupplyVoltageHighEventClass {
	MTRPumpConfigurationAndControlClusterSupplyVoltageHighEventClassOnce.Do(func() {
		MTRPumpConfigurationAndControlClusterSupplyVoltageHighEventClass = _MTRPumpConfigurationAndControlClusterSupplyVoltageHighEventClass{objc.GetClass("MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent")}
	})
	return MTRPumpConfigurationAndControlClusterSupplyVoltageHighEventClass
}

type _MTRPumpConfigurationAndControlClusterSupplyVoltageHighEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent] class.
type IMTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent
type MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent struct {
	objectivec.Object
}

// MTRPumpConfigurationAndControlClusterSupplyVoltageHighEventFrom constructs a [MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent] from an unsafe.Pointer.
func MTRPumpConfigurationAndControlClusterSupplyVoltageHighEventFrom(ptr unsafe.Pointer) MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent {
	return MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRPumpConfigurationAndControlClusterSupplyVoltageHighEventClass) Alloc() MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRPumpConfigurationAndControlClusterSupplyVoltageHighEventClass) New() MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent) Init() MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent) Autorelease() MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent creates a new MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent instance.
func NewMTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent() MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent {
	return getMTRPumpConfigurationAndControlClusterSupplyVoltageHighEventClass().New()
}




