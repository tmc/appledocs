// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent] class.
var (
	MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEventClass     _MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEventClass
	MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEventClassOnce sync.Once
)

func getMTRPumpConfigurationAndControlClusterElectronicTemperatureHighEventClass() _MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEventClass {
	MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEventClassOnce.Do(func() {
		MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEventClass = _MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEventClass{objc.GetClass("MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent")}
	})
	return MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEventClass
}

type _MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent] class.
type IMTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent
type MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent struct {
	objectivec.Object
}

// MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEventFrom constructs a [MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent] from an unsafe.Pointer.
func MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEventFrom(ptr unsafe.Pointer) MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent {
	return MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEventClass) Alloc() MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEventClass) New() MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent) Init() MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent) Autorelease() MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent creates a new MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent instance.
func NewMTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent() MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent {
	return getMTRPumpConfigurationAndControlClusterElectronicTemperatureHighEventClass().New()
}




