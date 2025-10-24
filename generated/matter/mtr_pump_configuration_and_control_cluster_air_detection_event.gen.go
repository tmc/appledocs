// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRPumpConfigurationAndControlClusterAirDetectionEvent] class.
var (
	MTRPumpConfigurationAndControlClusterAirDetectionEventClass     _MTRPumpConfigurationAndControlClusterAirDetectionEventClass
	MTRPumpConfigurationAndControlClusterAirDetectionEventClassOnce sync.Once
)

func getMTRPumpConfigurationAndControlClusterAirDetectionEventClass() _MTRPumpConfigurationAndControlClusterAirDetectionEventClass {
	MTRPumpConfigurationAndControlClusterAirDetectionEventClassOnce.Do(func() {
		MTRPumpConfigurationAndControlClusterAirDetectionEventClass = _MTRPumpConfigurationAndControlClusterAirDetectionEventClass{objc.GetClass("MTRPumpConfigurationAndControlClusterAirDetectionEvent")}
	})
	return MTRPumpConfigurationAndControlClusterAirDetectionEventClass
}

type _MTRPumpConfigurationAndControlClusterAirDetectionEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRPumpConfigurationAndControlClusterAirDetectionEvent] class.
type IMTRPumpConfigurationAndControlClusterAirDetectionEvent interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPumpConfigurationAndControlClusterAirDetectionEvent
type MTRPumpConfigurationAndControlClusterAirDetectionEvent struct {
	objectivec.Object
}

// MTRPumpConfigurationAndControlClusterAirDetectionEventFrom constructs a [MTRPumpConfigurationAndControlClusterAirDetectionEvent] from an unsafe.Pointer.
func MTRPumpConfigurationAndControlClusterAirDetectionEventFrom(ptr unsafe.Pointer) MTRPumpConfigurationAndControlClusterAirDetectionEvent {
	return MTRPumpConfigurationAndControlClusterAirDetectionEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRPumpConfigurationAndControlClusterAirDetectionEventClass) Alloc() MTRPumpConfigurationAndControlClusterAirDetectionEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterAirDetectionEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRPumpConfigurationAndControlClusterAirDetectionEventClass) New() MTRPumpConfigurationAndControlClusterAirDetectionEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterAirDetectionEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPumpConfigurationAndControlClusterAirDetectionEvent) Init() MTRPumpConfigurationAndControlClusterAirDetectionEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterAirDetectionEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPumpConfigurationAndControlClusterAirDetectionEvent) Autorelease() MTRPumpConfigurationAndControlClusterAirDetectionEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterAirDetectionEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPumpConfigurationAndControlClusterAirDetectionEvent creates a new MTRPumpConfigurationAndControlClusterAirDetectionEvent instance.
func NewMTRPumpConfigurationAndControlClusterAirDetectionEvent() MTRPumpConfigurationAndControlClusterAirDetectionEvent {
	return getMTRPumpConfigurationAndControlClusterAirDetectionEventClass().New()
}




