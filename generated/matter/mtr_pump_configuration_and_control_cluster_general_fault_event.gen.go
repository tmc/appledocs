// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRPumpConfigurationAndControlClusterGeneralFaultEvent] class.
var (
	MTRPumpConfigurationAndControlClusterGeneralFaultEventClass     _MTRPumpConfigurationAndControlClusterGeneralFaultEventClass
	MTRPumpConfigurationAndControlClusterGeneralFaultEventClassOnce sync.Once
)

func getMTRPumpConfigurationAndControlClusterGeneralFaultEventClass() _MTRPumpConfigurationAndControlClusterGeneralFaultEventClass {
	MTRPumpConfigurationAndControlClusterGeneralFaultEventClassOnce.Do(func() {
		MTRPumpConfigurationAndControlClusterGeneralFaultEventClass = _MTRPumpConfigurationAndControlClusterGeneralFaultEventClass{objc.GetClass("MTRPumpConfigurationAndControlClusterGeneralFaultEvent")}
	})
	return MTRPumpConfigurationAndControlClusterGeneralFaultEventClass
}

type _MTRPumpConfigurationAndControlClusterGeneralFaultEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRPumpConfigurationAndControlClusterGeneralFaultEvent] class.
type IMTRPumpConfigurationAndControlClusterGeneralFaultEvent interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPumpConfigurationAndControlClusterGeneralFaultEvent
type MTRPumpConfigurationAndControlClusterGeneralFaultEvent struct {
	objectivec.Object
}

// MTRPumpConfigurationAndControlClusterGeneralFaultEventFrom constructs a [MTRPumpConfigurationAndControlClusterGeneralFaultEvent] from an unsafe.Pointer.
func MTRPumpConfigurationAndControlClusterGeneralFaultEventFrom(ptr unsafe.Pointer) MTRPumpConfigurationAndControlClusterGeneralFaultEvent {
	return MTRPumpConfigurationAndControlClusterGeneralFaultEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRPumpConfigurationAndControlClusterGeneralFaultEventClass) Alloc() MTRPumpConfigurationAndControlClusterGeneralFaultEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterGeneralFaultEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRPumpConfigurationAndControlClusterGeneralFaultEventClass) New() MTRPumpConfigurationAndControlClusterGeneralFaultEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterGeneralFaultEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPumpConfigurationAndControlClusterGeneralFaultEvent) Init() MTRPumpConfigurationAndControlClusterGeneralFaultEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterGeneralFaultEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPumpConfigurationAndControlClusterGeneralFaultEvent) Autorelease() MTRPumpConfigurationAndControlClusterGeneralFaultEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterGeneralFaultEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPumpConfigurationAndControlClusterGeneralFaultEvent creates a new MTRPumpConfigurationAndControlClusterGeneralFaultEvent instance.
func NewMTRPumpConfigurationAndControlClusterGeneralFaultEvent() MTRPumpConfigurationAndControlClusterGeneralFaultEvent {
	return getMTRPumpConfigurationAndControlClusterGeneralFaultEventClass().New()
}




