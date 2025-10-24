// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRPumpConfigurationAndControlClusterSensorFailureEvent] class.
var (
	MTRPumpConfigurationAndControlClusterSensorFailureEventClass     _MTRPumpConfigurationAndControlClusterSensorFailureEventClass
	MTRPumpConfigurationAndControlClusterSensorFailureEventClassOnce sync.Once
)

func getMTRPumpConfigurationAndControlClusterSensorFailureEventClass() _MTRPumpConfigurationAndControlClusterSensorFailureEventClass {
	MTRPumpConfigurationAndControlClusterSensorFailureEventClassOnce.Do(func() {
		MTRPumpConfigurationAndControlClusterSensorFailureEventClass = _MTRPumpConfigurationAndControlClusterSensorFailureEventClass{objc.GetClass("MTRPumpConfigurationAndControlClusterSensorFailureEvent")}
	})
	return MTRPumpConfigurationAndControlClusterSensorFailureEventClass
}

type _MTRPumpConfigurationAndControlClusterSensorFailureEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRPumpConfigurationAndControlClusterSensorFailureEvent] class.
type IMTRPumpConfigurationAndControlClusterSensorFailureEvent interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPumpConfigurationAndControlClusterSensorFailureEvent
type MTRPumpConfigurationAndControlClusterSensorFailureEvent struct {
	objectivec.Object
}

// MTRPumpConfigurationAndControlClusterSensorFailureEventFrom constructs a [MTRPumpConfigurationAndControlClusterSensorFailureEvent] from an unsafe.Pointer.
func MTRPumpConfigurationAndControlClusterSensorFailureEventFrom(ptr unsafe.Pointer) MTRPumpConfigurationAndControlClusterSensorFailureEvent {
	return MTRPumpConfigurationAndControlClusterSensorFailureEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRPumpConfigurationAndControlClusterSensorFailureEventClass) Alloc() MTRPumpConfigurationAndControlClusterSensorFailureEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSensorFailureEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRPumpConfigurationAndControlClusterSensorFailureEventClass) New() MTRPumpConfigurationAndControlClusterSensorFailureEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSensorFailureEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPumpConfigurationAndControlClusterSensorFailureEvent) Init() MTRPumpConfigurationAndControlClusterSensorFailureEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSensorFailureEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPumpConfigurationAndControlClusterSensorFailureEvent) Autorelease() MTRPumpConfigurationAndControlClusterSensorFailureEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSensorFailureEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPumpConfigurationAndControlClusterSensorFailureEvent creates a new MTRPumpConfigurationAndControlClusterSensorFailureEvent instance.
func NewMTRPumpConfigurationAndControlClusterSensorFailureEvent() MTRPumpConfigurationAndControlClusterSensorFailureEvent {
	return getMTRPumpConfigurationAndControlClusterSensorFailureEventClass().New()
}




