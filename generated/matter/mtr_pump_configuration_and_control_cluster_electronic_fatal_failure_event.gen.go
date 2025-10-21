// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent] class.
var (
	MTRPumpConfigurationAndControlClusterElectronicFatalFailureEventClass     _MTRPumpConfigurationAndControlClusterElectronicFatalFailureEventClass
	MTRPumpConfigurationAndControlClusterElectronicFatalFailureEventClassOnce sync.Once
)

func getMTRPumpConfigurationAndControlClusterElectronicFatalFailureEventClass() _MTRPumpConfigurationAndControlClusterElectronicFatalFailureEventClass {
	MTRPumpConfigurationAndControlClusterElectronicFatalFailureEventClassOnce.Do(func() {
		MTRPumpConfigurationAndControlClusterElectronicFatalFailureEventClass = _MTRPumpConfigurationAndControlClusterElectronicFatalFailureEventClass{objc.GetClass("MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent")}
	})
	return MTRPumpConfigurationAndControlClusterElectronicFatalFailureEventClass
}

type _MTRPumpConfigurationAndControlClusterElectronicFatalFailureEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent] class.
type IMTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent
type MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent struct {
	objectivec.Object
}

// MTRPumpConfigurationAndControlClusterElectronicFatalFailureEventFrom constructs a [MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent] from an unsafe.Pointer.
func MTRPumpConfigurationAndControlClusterElectronicFatalFailureEventFrom(ptr unsafe.Pointer) MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent {
	return MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRPumpConfigurationAndControlClusterElectronicFatalFailureEventClass) Alloc() MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRPumpConfigurationAndControlClusterElectronicFatalFailureEventClass) New() MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent) Init() MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent) Autorelease() MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent creates a new MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent instance.
func NewMTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent() MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent {
	return getMTRPumpConfigurationAndControlClusterElectronicFatalFailureEventClass().New()
}




