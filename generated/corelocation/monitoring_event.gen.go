// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MonitoringEvent] class.
var monitoringEventClass = _MonitoringEventClass{objc.GetClass("CLMonitoringEvent")}

type _MonitoringEventClass struct {
	class objc.Class
}

// The object that the framework passes to the monitor’s callback handler upon receiving an event. [Full Topic]
//
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



