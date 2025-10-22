// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Monitor] class.
var (
	MonitorClass     _MonitorClass
	MonitorClassOnce sync.Once
)

func getMonitorClass() _MonitorClass {
	MonitorClassOnce.Do(func() {
		MonitorClass = _MonitorClass{objc.GetClass("CLMonitor")}
	})
	return MonitorClass
}

type _MonitorClass struct {
	class objc.Class
}

// An interface definition for the [Monitor] class.
type IMonitor interface {
	objectivec.IObject
	AddConditionForMonitoringIdentifier(condition ICLCondition, identifier string)
	AddConditionForMonitoringIdentifierAssumedState(condition ICLCondition, identifier string, state MonitoringState)
	MonitoringRecordForIdentifier(identifier string) MonitoringRecord
	RemoveConditionFromMonitoringWithIdentifier(identifier string)
	MonitoredIdentifiers() []string
	Name() string
}

// An object that monitors the conditions you add to it.
//
// Use to monitor for and observe events such as the entry to a specific geographic area or proximity to a beacon with characteristics that you specify. This service is unavailable in a compatible iPad or iPhone app running in visionOS.
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

// Alloc allocates a new instance without initialization.
func (mc _MonitorClass) Alloc() Monitor {
	rv := objc.Send[Monitor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MonitorClass) New() Monitor {
	rv := objc.Send[Monitor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ Monitor) Init() Monitor {
	rv := objc.Send[Monitor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ Monitor) Autorelease() Monitor {
	rv := objc.Send[Monitor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMonitor creates a new Monitor instance.
func NewMonitor() Monitor {
	return getMonitorClass().New()
}


// Creates a location monitor with the configuration and event handler you provide.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitor-6ynwz/requestMonitorWithConfiguration:completion:
func (mc _MonitorClass) RequestMonitorWithConfigurationCompletion(config ICLMonitorConfiguration, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("requestMonitorWithConfiguration:completion:"), config, completionHandler)
}

// Adds a condition to monitor with the identifier you provide.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitor-6ynwz/addConditionForMonitoring:identifier:
func (m_ Monitor) AddConditionForMonitoringIdentifier(condition ICLCondition, identifier string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addConditionForMonitoring:identifier:"), condition, objc.String(identifier))
}

// Adds a condition to monitor with the state and identifier you provide.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitor-6ynwz/addConditionForMonitoring:identifier:assumedState:
func (m_ Monitor) AddConditionForMonitoringIdentifierAssumedState(condition ICLCondition, identifier string, state MonitoringState) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addConditionForMonitoring:identifier:assumedState:"), condition, objc.String(identifier), state)
}

// Gets the monitoring record containing the condition and most recent monitoring event for the identifier you supply, if applicable.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitor-6ynwz/monitoringRecordForIdentifier:
func (m_ Monitor) MonitoringRecordForIdentifier(identifier string) MonitoringRecord {
	rv := objc.Send[MonitoringRecord](m_.ID, objc.Sel("monitoringRecordForIdentifier:"), objc.String(identifier))
	return rv
}

// Removes the monitoring record with the identifier from monitoring.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitor-6ynwz/removeConditionFromMonitoringWithIdentifier:
func (m_ Monitor) RemoveConditionFromMonitoringWithIdentifier(identifier string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeConditionFromMonitoringWithIdentifier:"), objc.String(identifier))
}

// An array that contains all the identifiers for each condition that the monitor is monitoring.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitor-6ynwz/monitoredIdentifiers
func (m_ Monitor) MonitoredIdentifiers() []string {
	rv := objc.Send[[]string](m_.ID, objc.Sel("monitoredIdentifiers"))
	return rv
}

// The name associated with the location monitor instance.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitor-6ynwz/name
func (m_ Monitor) Name() string {
	rv := objc.Send[string](m_.ID, objc.Sel("name"))
	return rv
}



