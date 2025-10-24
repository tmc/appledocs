// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent] class.
var (
	MTRPumpConfigurationAndControlClusterMotorTemperatureHighEventClass     _MTRPumpConfigurationAndControlClusterMotorTemperatureHighEventClass
	MTRPumpConfigurationAndControlClusterMotorTemperatureHighEventClassOnce sync.Once
)

func getMTRPumpConfigurationAndControlClusterMotorTemperatureHighEventClass() _MTRPumpConfigurationAndControlClusterMotorTemperatureHighEventClass {
	MTRPumpConfigurationAndControlClusterMotorTemperatureHighEventClassOnce.Do(func() {
		MTRPumpConfigurationAndControlClusterMotorTemperatureHighEventClass = _MTRPumpConfigurationAndControlClusterMotorTemperatureHighEventClass{objc.GetClass("MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent")}
	})
	return MTRPumpConfigurationAndControlClusterMotorTemperatureHighEventClass
}

type _MTRPumpConfigurationAndControlClusterMotorTemperatureHighEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent] class.
type IMTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent
type MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent struct {
	objectivec.Object
}

// MTRPumpConfigurationAndControlClusterMotorTemperatureHighEventFrom constructs a [MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent] from an unsafe.Pointer.
func MTRPumpConfigurationAndControlClusterMotorTemperatureHighEventFrom(ptr unsafe.Pointer) MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent {
	return MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRPumpConfigurationAndControlClusterMotorTemperatureHighEventClass) Alloc() MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRPumpConfigurationAndControlClusterMotorTemperatureHighEventClass) New() MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent) Init() MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent) Autorelease() MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent creates a new MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent instance.
func NewMTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent() MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent {
	return getMTRPumpConfigurationAndControlClusterMotorTemperatureHighEventClass().New()
}




