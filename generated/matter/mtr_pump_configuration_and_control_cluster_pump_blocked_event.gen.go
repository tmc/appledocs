// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRPumpConfigurationAndControlClusterPumpBlockedEvent] class.
var (
	MTRPumpConfigurationAndControlClusterPumpBlockedEventClass     _MTRPumpConfigurationAndControlClusterPumpBlockedEventClass
	MTRPumpConfigurationAndControlClusterPumpBlockedEventClassOnce sync.Once
)

func getMTRPumpConfigurationAndControlClusterPumpBlockedEventClass() _MTRPumpConfigurationAndControlClusterPumpBlockedEventClass {
	MTRPumpConfigurationAndControlClusterPumpBlockedEventClassOnce.Do(func() {
		MTRPumpConfigurationAndControlClusterPumpBlockedEventClass = _MTRPumpConfigurationAndControlClusterPumpBlockedEventClass{objc.GetClass("MTRPumpConfigurationAndControlClusterPumpBlockedEvent")}
	})
	return MTRPumpConfigurationAndControlClusterPumpBlockedEventClass
}

type _MTRPumpConfigurationAndControlClusterPumpBlockedEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRPumpConfigurationAndControlClusterPumpBlockedEvent] class.
type IMTRPumpConfigurationAndControlClusterPumpBlockedEvent interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPumpConfigurationAndControlClusterPumpBlockedEvent
type MTRPumpConfigurationAndControlClusterPumpBlockedEvent struct {
	objectivec.Object
}

// MTRPumpConfigurationAndControlClusterPumpBlockedEventFrom constructs a [MTRPumpConfigurationAndControlClusterPumpBlockedEvent] from an unsafe.Pointer.
func MTRPumpConfigurationAndControlClusterPumpBlockedEventFrom(ptr unsafe.Pointer) MTRPumpConfigurationAndControlClusterPumpBlockedEvent {
	return MTRPumpConfigurationAndControlClusterPumpBlockedEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRPumpConfigurationAndControlClusterPumpBlockedEventClass) Alloc() MTRPumpConfigurationAndControlClusterPumpBlockedEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterPumpBlockedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRPumpConfigurationAndControlClusterPumpBlockedEventClass) New() MTRPumpConfigurationAndControlClusterPumpBlockedEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterPumpBlockedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPumpConfigurationAndControlClusterPumpBlockedEvent) Init() MTRPumpConfigurationAndControlClusterPumpBlockedEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterPumpBlockedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPumpConfigurationAndControlClusterPumpBlockedEvent) Autorelease() MTRPumpConfigurationAndControlClusterPumpBlockedEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterPumpBlockedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPumpConfigurationAndControlClusterPumpBlockedEvent creates a new MTRPumpConfigurationAndControlClusterPumpBlockedEvent instance.
func NewMTRPumpConfigurationAndControlClusterPumpBlockedEvent() MTRPumpConfigurationAndControlClusterPumpBlockedEvent {
	return getMTRPumpConfigurationAndControlClusterPumpBlockedEventClass().New()
}




