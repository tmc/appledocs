// Code generated from Apple documentation for SafetyKit. DO NOT EDIT.

package safetykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corelocation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SACrashDetectionEvent */


/* debug [class_header]: Header for SACrashDetectionEvent */
// The class instance for the [SACrashDetectionEvent] class.
var (
	SACrashDetectionEventClass     _SACrashDetectionEventClass
	SACrashDetectionEventClassOnce sync.Once
)

func getSACrashDetectionEventClass() _SACrashDetectionEventClass {
	SACrashDetectionEventClassOnce.Do(func() {
		SACrashDetectionEventClass = _SACrashDetectionEventClass{objc.GetClass("SACrashDetectionEvent")}
	})
	return SACrashDetectionEventClass
}

type _SACrashDetectionEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SACrashDetectionEvent */
// An interface definition for the [SACrashDetectionEvent] class.
type ISACrashDetectionEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SACrashDetectionEvent */
	// properties:
	Date() objc.IObject /* cross-framework: NSDate */
	Location() corelocation.Location
	Response() SACrashDetectionEventResponse
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SACrashDetectionEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SACrashDetectionEvent */
// Alloc allocates a new instance without initialization.
func (sc _SACrashDetectionEventClass) Alloc() SACrashDetectionEvent {
	rv := objc.Send[SACrashDetectionEvent](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SACrashDetectionEventClass) New() SACrashDetectionEvent {
	rv := objc.Send[SACrashDetectionEvent](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SACrashDetectionEvent) Init() SACrashDetectionEvent {
	rv := objc.Send[SACrashDetectionEvent](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SACrashDetectionEvent) Autorelease() SACrashDetectionEvent {
	rv := objc.Send[SACrashDetectionEvent](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSACrashDetectionEvent creates a new SACrashDetectionEvent instance.
func NewSACrashDetectionEvent() SACrashDetectionEvent {
	return getSACrashDetectionEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SACrashDetectionEvent */
// Describes the information about a vehicular crash.
//
// When a vehicular crash occurs, SafetyKit calls your delegate’s method with an object. Inspect this object to determine information about the crash, including the date and time, location, and if the system attempted to contact emergency services.


// Describes the information about a vehicular crash.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SACrashDetectionEvent
type SACrashDetectionEvent struct {
	objectivec.Object
}

// SACrashDetectionEventFrom constructs a [SACrashDetectionEvent] from an unsafe.Pointer.
//
// Describes the information about a vehicular crash.
func SACrashDetectionEventFrom(ptr unsafe.Pointer) SACrashDetectionEvent {
	return SACrashDetectionEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SACrashDetectionEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SACrashDetectionEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SACrashDetectionEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SACrashDetectionEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SACrashDetectionEvent */

// The date and time the crash occurred.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SACrashDetectionEvent/date
func (s_ SACrashDetectionEvent) Date() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](s_.ID, objc.Sel("date"))
	return rv
}/* debug [instance_properties/getter]: date */


// The longitude and latitude where the crash detection occurred.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SACrashDetectionEvent/location
func (s_ SACrashDetectionEvent) Location() corelocation.Location {
	rv := objc.Send[corelocation.Location](s_.ID, objc.Sel("location"))
	return rv
}/* debug [instance_properties/getter]: location */


// An indication of whether the system attempted to call an Emergency SOS provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SACrashDetectionEvent/response-swift.property
func (s_ SACrashDetectionEvent) Response() SACrashDetectionEventResponse {
	rv := objc.Send[SACrashDetectionEventResponse](s_.ID, objc.Sel("response"))
	return rv
}/* debug [instance_properties/getter]: response */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SACrashDetectionEvent */



