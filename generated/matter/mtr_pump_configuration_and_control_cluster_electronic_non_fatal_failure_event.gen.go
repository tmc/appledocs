// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent] class.
var (
	MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEventClass     _MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEventClass
	MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEventClassOnce sync.Once
)

func getMTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEventClass() _MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEventClass {
	MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEventClassOnce.Do(func() {
		MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEventClass = _MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEventClass{objc.GetClass("MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent")}
	})
	return MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEventClass
}

type _MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent] class.
type IMTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent
type MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent struct {
	objectivec.Object
}

// MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEventFrom constructs a [MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent] from an unsafe.Pointer.
func MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEventFrom(ptr unsafe.Pointer) MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent {
	return MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEventClass) Alloc() MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEventClass) New() MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent) Init() MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent) Autorelease() MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent creates a new MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent instance.
func NewMTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent() MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent {
	return getMTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEventClass().New()
}




