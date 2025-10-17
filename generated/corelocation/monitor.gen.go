// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Monitor] class.
var monitorClass = _MonitorClass{objc.GetClass("CLMonitor")}

type _MonitorClass struct {
	class objc.Class
}

// An object that monitors the conditions you add to it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitor-6ynwz

type Monitor struct {
	objectivec.Object
}

// MonitorFrom constructs a [Monitor] from an unsafe.Pointer.
//
// An object that monitors the conditions you add to it.
func MonitorFrom(ptr unsafe.Pointer) Monitor {
	return Monitor{objectivec.Object{objc.ID(ptr)}}
}

// Creates a location monitor with the configuration and event handler you provide. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitor-6ynwz/requestMonitorWithConfiguration:completion:
func (mc _MonitorClass) RequestMonitorWithConfigurationCompletion(config unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("requestMonitorWithConfiguration:completion:"), config, completionHandler)
}
// Adds a condition to monitor with the identifier you provide. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitor-6ynwz/addConditionForMonitoring:identifier:
func (m_ Monitor) AddConditionForMonitoringIdentifier(condition unsafe.Pointer, identifier string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addConditionForMonitoring:identifier:"), condition, identifier)
}
// Adds a condition to monitor with the state and identifier you provide. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitor-6ynwz/addConditionForMonitoring:identifier:assumedState:
func (m_ Monitor) AddConditionForMonitoringIdentifierAssumedState(condition unsafe.Pointer, identifier string, state unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addConditionForMonitoring:identifier:assumedState:"), condition, identifier, state)
}
// Gets the monitoring record containing the condition and most recent monitoring event for the identifier you supply, if applicable. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitor-6ynwz/monitoringRecordForIdentifier:
func (m_ Monitor) MonitoringRecordForIdentifier(identifier string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("monitoringRecordForIdentifier:"), identifier)
	return rv
}
// Removes the monitoring record with the identifier from monitoring. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitor-6ynwz/removeConditionFromMonitoringWithIdentifier:
func (m_ Monitor) RemoveConditionFromMonitoringWithIdentifier(identifier string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeConditionFromMonitoringWithIdentifier:"), identifier)
}


