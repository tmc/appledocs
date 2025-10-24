// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MonitoringEvent] class.
var (
	MonitoringEventClass     _MonitoringEventClass
	MonitoringEventClassOnce sync.Once
)

func getMonitoringEventClass() _MonitoringEventClass {
	MonitoringEventClassOnce.Do(func() {
		MonitoringEventClass = _MonitoringEventClass{objc.GetClass("CLMonitoringEvent")}
	})
	return MonitoringEventClass
}

type _MonitoringEventClass struct {
	class objc.Class
}

// An interface definition for the [MonitoringEvent] class.
type IMonitoringEvent interface {
	objectivec.IObject
	// properties:
	// methods:
}

// The object that the framework passes to the monitor’s callback handler upon receiving an event.
//
// Instances of contain detailed information about an event in the monitoring of a by a .

// The object that the framework passes to the monitor’s callback handler upon receiving an event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitoringEvent
type MonitoringEvent struct {
	objectivec.Object
}

// MonitoringEventFrom constructs a [MonitoringEvent] from an unsafe.Pointer.
//
// The object that the framework passes to the monitor’s callback handler upon receiving an event.
func MonitoringEventFrom(ptr unsafe.Pointer) MonitoringEvent {
	return MonitoringEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MonitoringEventClass) Alloc() MonitoringEvent {
	rv := objc.Send[MonitoringEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MonitoringEventClass) New() MonitoringEvent {
	rv := objc.Send[MonitoringEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MonitoringEvent) Init() MonitoringEvent {
	rv := objc.Send[MonitoringEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MonitoringEvent) Autorelease() MonitoringEvent {
	rv := objc.Send[MonitoringEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMonitoringEvent creates a new MonitoringEvent instance.
func NewMonitoringEvent() MonitoringEvent {
	return getMonitoringEventClass().New()
}
