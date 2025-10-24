// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent] class.
var (
	MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEventClass     _MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEventClass
	MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEventClassOnce sync.Once
)

func getMTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEventClass() _MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEventClass {
	MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEventClassOnce.Do(func() {
		MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEventClass = _MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEventClass{objc.GetClass("MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent")}
	})
	return MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEventClass
}

type _MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent] class.
type IMTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent
type MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent struct {
	objectivec.Object
}

// MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEventFrom constructs a [MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent] from an unsafe.Pointer.
func MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEventFrom(ptr unsafe.Pointer) MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent {
	return MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEventClass) Alloc() MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEventClass) New() MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent) Init() MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent) Autorelease() MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent creates a new MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent instance.
func NewMTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent() MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent {
	return getMTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEventClass().New()
}




