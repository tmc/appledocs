// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRPumpConfigurationAndControlClusterTurbineOperationEvent] class.
var (
	MTRPumpConfigurationAndControlClusterTurbineOperationEventClass     _MTRPumpConfigurationAndControlClusterTurbineOperationEventClass
	MTRPumpConfigurationAndControlClusterTurbineOperationEventClassOnce sync.Once
)

func getMTRPumpConfigurationAndControlClusterTurbineOperationEventClass() _MTRPumpConfigurationAndControlClusterTurbineOperationEventClass {
	MTRPumpConfigurationAndControlClusterTurbineOperationEventClassOnce.Do(func() {
		MTRPumpConfigurationAndControlClusterTurbineOperationEventClass = _MTRPumpConfigurationAndControlClusterTurbineOperationEventClass{objc.GetClass("MTRPumpConfigurationAndControlClusterTurbineOperationEvent")}
	})
	return MTRPumpConfigurationAndControlClusterTurbineOperationEventClass
}

type _MTRPumpConfigurationAndControlClusterTurbineOperationEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRPumpConfigurationAndControlClusterTurbineOperationEvent] class.
type IMTRPumpConfigurationAndControlClusterTurbineOperationEvent interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPumpConfigurationAndControlClusterTurbineOperationEvent
type MTRPumpConfigurationAndControlClusterTurbineOperationEvent struct {
	objectivec.Object
}

// MTRPumpConfigurationAndControlClusterTurbineOperationEventFrom constructs a [MTRPumpConfigurationAndControlClusterTurbineOperationEvent] from an unsafe.Pointer.
func MTRPumpConfigurationAndControlClusterTurbineOperationEventFrom(ptr unsafe.Pointer) MTRPumpConfigurationAndControlClusterTurbineOperationEvent {
	return MTRPumpConfigurationAndControlClusterTurbineOperationEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRPumpConfigurationAndControlClusterTurbineOperationEventClass) Alloc() MTRPumpConfigurationAndControlClusterTurbineOperationEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterTurbineOperationEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRPumpConfigurationAndControlClusterTurbineOperationEventClass) New() MTRPumpConfigurationAndControlClusterTurbineOperationEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterTurbineOperationEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPumpConfigurationAndControlClusterTurbineOperationEvent) Init() MTRPumpConfigurationAndControlClusterTurbineOperationEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterTurbineOperationEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPumpConfigurationAndControlClusterTurbineOperationEvent) Autorelease() MTRPumpConfigurationAndControlClusterTurbineOperationEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterTurbineOperationEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPumpConfigurationAndControlClusterTurbineOperationEvent creates a new MTRPumpConfigurationAndControlClusterTurbineOperationEvent instance.
func NewMTRPumpConfigurationAndControlClusterTurbineOperationEvent() MTRPumpConfigurationAndControlClusterTurbineOperationEvent {
	return getMTRPumpConfigurationAndControlClusterTurbineOperationEventClass().New()
}




