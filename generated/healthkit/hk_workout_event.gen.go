// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKWorkoutEvent */


/* debug [class_header]: Header for HKWorkoutEvent */
// The class instance for the [HKWorkoutEvent] class.
var (
	HKWorkoutEventClass     _HKWorkoutEventClass
	HKWorkoutEventClassOnce sync.Once
)

func getHKWorkoutEventClass() _HKWorkoutEventClass {
	HKWorkoutEventClassOnce.Do(func() {
		HKWorkoutEventClass = _HKWorkoutEventClass{objc.GetClass("HKWorkoutEvent")}
	})
	return HKWorkoutEventClass
}

type _HKWorkoutEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKWorkoutEvent */
// An interface definition for the [HKWorkoutEvent] class.
type IHKWorkoutEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKWorkoutEvent */
	// properties:
	Date() objc.IObject /* cross-framework: NSDate */
	DateInterval() foundation.DateInterval
	Metadata() foundation.IDictionary
	Type() HKWorkoutEventType
	HKWorkoutTypeIdentifier() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKWorkoutEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKWorkoutEvent */
// Alloc allocates a new instance without initialization.
func (hc _HKWorkoutEventClass) Alloc() HKWorkoutEvent {
	rv := objc.Send[HKWorkoutEvent](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKWorkoutEventClass) New() HKWorkoutEvent {
	rv := objc.Send[HKWorkoutEvent](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKWorkoutEvent) Init() HKWorkoutEvent {
	rv := objc.Send[HKWorkoutEvent](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKWorkoutEvent) Autorelease() HKWorkoutEvent {
	rv := objc.Send[HKWorkoutEvent](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKWorkoutEvent creates a new HKWorkoutEvent instance.
func NewHKWorkoutEvent() HKWorkoutEvent {
	return getHKWorkoutEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKWorkoutEvent */
// An object representing an important event during a workout.
//
// You can use workout events to toggle a workout between an active and an inactive state, or to mark points of interest during a workout. Workouts start in an active state. A pause event switches it to an inactive state; a resume event switches it back to an active state. Adding a pause event when the workout is already inactive, or a resume event when the workout is already active, does not affect the workout’s state. These events are ignored. The lap, segment, and marker events are used to identify periods of interest during a workout. Use lap events to partition a workout into segments of equal distance. Segment events mark important periods during the workout, while markers identify important points in time.


// An object representing an important event during a workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutEvent
type HKWorkoutEvent struct {
	objectivec.Object
}

// HKWorkoutEventFrom constructs a [HKWorkoutEvent] from an unsafe.Pointer.
//
// An object representing an important event during a workout.
func HKWorkoutEventFrom(ptr unsafe.Pointer) HKWorkoutEvent {
	return HKWorkoutEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKWorkoutEvent */

// Instantiates and returns a new workout event with the specified type and date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutEvent/init(type:date:)
func NewHKWorkoutEventWithTypeDate(type_ HKWorkoutEventType, date objc.IObject /* cross-framework: NSDate */) HKWorkoutEvent {
	rv := objc.Send[HKWorkoutEvent](objc.ID(getHKWorkoutEventClass().class), objc.Sel("workoutEventWithType:date:"), type_, date)
	return rv
}/* debug [class_init_methods/constructor]: NewHKWorkoutEventWithTypeDate */


// Instantiates and returns a new workout event with the specified type, date interval, and metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutEvent/init(type:dateInterval:metadata:)
func NewHKWorkoutEventWithTypeDateIntervalMetadata(type_ HKWorkoutEventType, dateInterval foundation.DateInterval, metadata foundation.IDictionary) HKWorkoutEvent {
	rv := objc.Send[HKWorkoutEvent](objc.ID(getHKWorkoutEventClass().class), objc.Sel("workoutEventWithType:dateInterval:metadata:"), type_, dateInterval, metadata)
	return rv
}/* debug [class_init_methods/constructor]: NewHKWorkoutEventWithTypeDateIntervalMetadata */


// Instantiates and returns a new workout event with the specified type, date, and metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutEvent/init(type:date:metadata:)
func NewHKWorkoutEventWithTypeDateMetadata(type_ HKWorkoutEventType, date objc.IObject /* cross-framework: NSDate */, metadata foundation.IDictionary) HKWorkoutEvent {
	rv := objc.Send[HKWorkoutEvent](objc.ID(getHKWorkoutEventClass().class), objc.Sel("workoutEventWithType:date:metadata:"), type_, date, metadata)
	return rv
}/* debug [class_init_methods/constructor]: NewHKWorkoutEventWithTypeDateMetadata */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKWorkoutEvent */

// Instantiates and returns a new workout event with the specified type and date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutEvent/init(type:date:)
func (hc _HKWorkoutEventClass) WorkoutEventWithTypeDate(type_ HKWorkoutEventType, date objc.IObject /* cross-framework: NSDate */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("workoutEventWithType:date:"), type_, date)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WorkoutEventWithTypeDate) */


// Instantiates and returns a new workout event with the specified type, date, and metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutEvent/init(type:date:metadata:)
func (hc _HKWorkoutEventClass) WorkoutEventWithTypeDateMetadata(type_ HKWorkoutEventType, date objc.IObject /* cross-framework: NSDate */, metadata foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("workoutEventWithType:date:metadata:"), type_, date, metadata)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WorkoutEventWithTypeDateMetadata) */


// Instantiates and returns a new workout event with the specified type, date interval, and metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutEvent/init(type:dateInterval:metadata:)
func (hc _HKWorkoutEventClass) WorkoutEventWithTypeDateIntervalMetadata(type_ HKWorkoutEventType, dateInterval foundation.DateInterval, metadata foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("workoutEventWithType:dateInterval:metadata:"), type_, dateInterval, metadata)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WorkoutEventWithTypeDateIntervalMetadata) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKWorkoutEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKWorkoutEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKWorkoutEvent */

// The time when the transition occurred.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutEvent/date
func (h_ HKWorkoutEvent) Date() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](h_.ID, objc.Sel("date"))
	return rv
}/* debug [instance_properties/getter]: date */


// The time and duration of the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutEvent/dateInterval
func (h_ HKWorkoutEvent) DateInterval() foundation.DateInterval {
	rv := objc.Send[foundation.DateInterval](h_.ID, objc.Sel("dateInterval"))
	return rv
}/* debug [instance_properties/getter]: dateInterval */


// The metadata associated with the workout event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutEvent/metadata
func (h_ HKWorkoutEvent) Metadata() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](h_.ID, objc.Sel("metadata"))
	return rv
}/* debug [instance_properties/getter]: metadata */


// The type of workout event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutEvent/type
func (h_ HKWorkoutEvent) Type() HKWorkoutEventType {
	rv := objc.Send[HKWorkoutEventType](h_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// The workout type identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkouttypeidentifier
func (h_ HKWorkoutEvent) HKWorkoutTypeIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKWorkoutTypeIdentifier"))
	return rv
}/* debug [instance_properties/getter]: HKWorkoutTypeIdentifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKWorkoutEvent */


