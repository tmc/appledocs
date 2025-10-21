// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRPumpConfigurationAndControlClusterSystemPressureHighEvent] class.
var (
	MTRPumpConfigurationAndControlClusterSystemPressureHighEventClass     _MTRPumpConfigurationAndControlClusterSystemPressureHighEventClass
	MTRPumpConfigurationAndControlClusterSystemPressureHighEventClassOnce sync.Once
)

func getMTRPumpConfigurationAndControlClusterSystemPressureHighEventClass() _MTRPumpConfigurationAndControlClusterSystemPressureHighEventClass {
	MTRPumpConfigurationAndControlClusterSystemPressureHighEventClassOnce.Do(func() {
		MTRPumpConfigurationAndControlClusterSystemPressureHighEventClass = _MTRPumpConfigurationAndControlClusterSystemPressureHighEventClass{objc.GetClass("MTRPumpConfigurationAndControlClusterSystemPressureHighEvent")}
	})
	return MTRPumpConfigurationAndControlClusterSystemPressureHighEventClass
}

type _MTRPumpConfigurationAndControlClusterSystemPressureHighEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRPumpConfigurationAndControlClusterSystemPressureHighEvent] class.
type IMTRPumpConfigurationAndControlClusterSystemPressureHighEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPumpConfigurationAndControlClusterSystemPressureHighEvent
type MTRPumpConfigurationAndControlClusterSystemPressureHighEvent struct {
	objectivec.Object
}

// MTRPumpConfigurationAndControlClusterSystemPressureHighEventFrom constructs a [MTRPumpConfigurationAndControlClusterSystemPressureHighEvent] from an unsafe.Pointer.
func MTRPumpConfigurationAndControlClusterSystemPressureHighEventFrom(ptr unsafe.Pointer) MTRPumpConfigurationAndControlClusterSystemPressureHighEvent {
	return MTRPumpConfigurationAndControlClusterSystemPressureHighEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRPumpConfigurationAndControlClusterSystemPressureHighEventClass) Alloc() MTRPumpConfigurationAndControlClusterSystemPressureHighEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSystemPressureHighEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRPumpConfigurationAndControlClusterSystemPressureHighEventClass) New() MTRPumpConfigurationAndControlClusterSystemPressureHighEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSystemPressureHighEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPumpConfigurationAndControlClusterSystemPressureHighEvent) Init() MTRPumpConfigurationAndControlClusterSystemPressureHighEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSystemPressureHighEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPumpConfigurationAndControlClusterSystemPressureHighEvent) Autorelease() MTRPumpConfigurationAndControlClusterSystemPressureHighEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSystemPressureHighEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPumpConfigurationAndControlClusterSystemPressureHighEvent creates a new MTRPumpConfigurationAndControlClusterSystemPressureHighEvent instance.
func NewMTRPumpConfigurationAndControlClusterSystemPressureHighEvent() MTRPumpConfigurationAndControlClusterSystemPressureHighEvent {
	return getMTRPumpConfigurationAndControlClusterSystemPressureHighEventClass().New()
}




