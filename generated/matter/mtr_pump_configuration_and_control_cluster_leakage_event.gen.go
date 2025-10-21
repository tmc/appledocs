// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRPumpConfigurationAndControlClusterLeakageEvent] class.
var (
	MTRPumpConfigurationAndControlClusterLeakageEventClass     _MTRPumpConfigurationAndControlClusterLeakageEventClass
	MTRPumpConfigurationAndControlClusterLeakageEventClassOnce sync.Once
)

func getMTRPumpConfigurationAndControlClusterLeakageEventClass() _MTRPumpConfigurationAndControlClusterLeakageEventClass {
	MTRPumpConfigurationAndControlClusterLeakageEventClassOnce.Do(func() {
		MTRPumpConfigurationAndControlClusterLeakageEventClass = _MTRPumpConfigurationAndControlClusterLeakageEventClass{objc.GetClass("MTRPumpConfigurationAndControlClusterLeakageEvent")}
	})
	return MTRPumpConfigurationAndControlClusterLeakageEventClass
}

type _MTRPumpConfigurationAndControlClusterLeakageEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRPumpConfigurationAndControlClusterLeakageEvent] class.
type IMTRPumpConfigurationAndControlClusterLeakageEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPumpConfigurationAndControlClusterLeakageEvent
type MTRPumpConfigurationAndControlClusterLeakageEvent struct {
	objectivec.Object
}

// MTRPumpConfigurationAndControlClusterLeakageEventFrom constructs a [MTRPumpConfigurationAndControlClusterLeakageEvent] from an unsafe.Pointer.
func MTRPumpConfigurationAndControlClusterLeakageEventFrom(ptr unsafe.Pointer) MTRPumpConfigurationAndControlClusterLeakageEvent {
	return MTRPumpConfigurationAndControlClusterLeakageEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRPumpConfigurationAndControlClusterLeakageEventClass) Alloc() MTRPumpConfigurationAndControlClusterLeakageEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterLeakageEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRPumpConfigurationAndControlClusterLeakageEventClass) New() MTRPumpConfigurationAndControlClusterLeakageEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterLeakageEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPumpConfigurationAndControlClusterLeakageEvent) Init() MTRPumpConfigurationAndControlClusterLeakageEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterLeakageEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPumpConfigurationAndControlClusterLeakageEvent) Autorelease() MTRPumpConfigurationAndControlClusterLeakageEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterLeakageEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPumpConfigurationAndControlClusterLeakageEvent creates a new MTRPumpConfigurationAndControlClusterLeakageEvent instance.
func NewMTRPumpConfigurationAndControlClusterLeakageEvent() MTRPumpConfigurationAndControlClusterLeakageEvent {
	return getMTRPumpConfigurationAndControlClusterLeakageEventClass().New()
}




