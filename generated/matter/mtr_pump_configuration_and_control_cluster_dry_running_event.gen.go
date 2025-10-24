// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRPumpConfigurationAndControlClusterDryRunningEvent] class.
var (
	MTRPumpConfigurationAndControlClusterDryRunningEventClass     _MTRPumpConfigurationAndControlClusterDryRunningEventClass
	MTRPumpConfigurationAndControlClusterDryRunningEventClassOnce sync.Once
)

func getMTRPumpConfigurationAndControlClusterDryRunningEventClass() _MTRPumpConfigurationAndControlClusterDryRunningEventClass {
	MTRPumpConfigurationAndControlClusterDryRunningEventClassOnce.Do(func() {
		MTRPumpConfigurationAndControlClusterDryRunningEventClass = _MTRPumpConfigurationAndControlClusterDryRunningEventClass{objc.GetClass("MTRPumpConfigurationAndControlClusterDryRunningEvent")}
	})
	return MTRPumpConfigurationAndControlClusterDryRunningEventClass
}

type _MTRPumpConfigurationAndControlClusterDryRunningEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRPumpConfigurationAndControlClusterDryRunningEvent] class.
type IMTRPumpConfigurationAndControlClusterDryRunningEvent interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPumpConfigurationAndControlClusterDryRunningEvent
type MTRPumpConfigurationAndControlClusterDryRunningEvent struct {
	objectivec.Object
}

// MTRPumpConfigurationAndControlClusterDryRunningEventFrom constructs a [MTRPumpConfigurationAndControlClusterDryRunningEvent] from an unsafe.Pointer.
func MTRPumpConfigurationAndControlClusterDryRunningEventFrom(ptr unsafe.Pointer) MTRPumpConfigurationAndControlClusterDryRunningEvent {
	return MTRPumpConfigurationAndControlClusterDryRunningEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRPumpConfigurationAndControlClusterDryRunningEventClass) Alloc() MTRPumpConfigurationAndControlClusterDryRunningEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterDryRunningEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRPumpConfigurationAndControlClusterDryRunningEventClass) New() MTRPumpConfigurationAndControlClusterDryRunningEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterDryRunningEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPumpConfigurationAndControlClusterDryRunningEvent) Init() MTRPumpConfigurationAndControlClusterDryRunningEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterDryRunningEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPumpConfigurationAndControlClusterDryRunningEvent) Autorelease() MTRPumpConfigurationAndControlClusterDryRunningEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterDryRunningEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPumpConfigurationAndControlClusterDryRunningEvent creates a new MTRPumpConfigurationAndControlClusterDryRunningEvent instance.
func NewMTRPumpConfigurationAndControlClusterDryRunningEvent() MTRPumpConfigurationAndControlClusterDryRunningEvent {
	return getMTRPumpConfigurationAndControlClusterDryRunningEventClass().New()
}




