// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	AuthorizationDeniedGlobally() bool
	AuthorizationRequestInProgress() bool
	ConditionLimitExceeded() bool
	ConditionUnsupported() bool
	Date() foundation.NSDate
	Identifier() string
	InsufficientlyInUse() bool
	PersistenceUnavailable() bool
	ServiceSessionRequired() bool
	State() MonitoringState
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



// A Boolean value that indicates whether the app has system-wide authorization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitoringEvent/authorizationDeniedGlobally
func (m_ MonitoringEvent) AuthorizationDeniedGlobally() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("authorizationDeniedGlobally"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitoringEvent/authorizationRequestInProgress
func (m_ MonitoringEvent) AuthorizationRequestInProgress() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("authorizationRequestInProgress"))
	return rv
}


// A Boolean value that indicates whether the app receives location updates based on other monitoring conditions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitoringEvent/conditionLimitExceeded
func (m_ MonitoringEvent) ConditionLimitExceeded() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("conditionLimitExceeded"))
	return rv
}


// A Boolean value that indicates whether the app receives location updates based on the supported condition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitoringEvent/conditionUnsupported
func (m_ MonitoringEvent) ConditionUnsupported() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("conditionUnsupported"))
	return rv
}


// The date the event occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitoringEvent/date
func (m_ MonitoringEvent) Date() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("date"))
	return rv
}


// A string that represents the identifier of a monitored condition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitoringEvent/identifier
func (m_ MonitoringEvent) Identifier() string {
	rv := objc.Send[string](m_.ID, objc.Sel("identifier"))
	return rv
}


// A Boolean value that indicates whether the app receives location updates because it’s insufficiently in use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitoringEvent/insufficientlyInUse
func (m_ MonitoringEvent) InsufficientlyInUse() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("insufficientlyInUse"))
	return rv
}


// A Boolean value that indicates whether it receives location updates based on successful persistence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitoringEvent/persistenceUnavailable
func (m_ MonitoringEvent) PersistenceUnavailable() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("persistenceUnavailable"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitoringEvent/serviceSessionRequired
func (m_ MonitoringEvent) ServiceSessionRequired() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("serviceSessionRequired"))
	return rv
}


// The state of the condition at the time of the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitoringEvent/state
func (m_ MonitoringEvent) State() MonitoringState {
	rv := objc.Send[MonitoringState](m_.ID, objc.Sel("state"))
	return rv
}



