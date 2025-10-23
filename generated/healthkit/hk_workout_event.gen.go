// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [HKWorkoutEvent] class.
type IHKWorkoutEvent interface {
	objectivec.IObject
	// properties:
	Date() foundation.objc.IObject /* cross-framework: Date */
	SetDate(value foundation.objc.IObject /* cross-framework: Date */)
	DateInterval() foundation.objc.IObject /* cross-framework: DateInterval */
	SetDateInterval(value foundation.objc.IObject /* cross-framework: DateInterval */)
	Metadata() string /* primitive/slice/pointer. */
	SetMetadata(value string /* primitive/slice/pointer. */)
	Type() unsafe.Pointer
	SetType(value unsafe.Pointer)
	HKWorkoutTypeIdentifier() string /* primitive/slice/pointer. */
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (hc _HKWorkoutEventClass) Alloc() HKWorkoutEvent {
	rv := objc.Send[HKWorkoutEvent](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The time when the transition occurred.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutevent/date
func (h_ HKWorkoutEvent) Date() foundation.objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](h_.ID, objc.Sel("date"))
	return rv
}


// The time when the transition occurred.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutevent/date
func (h_ HKWorkoutEvent) SetDate(value foundation.objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setDate:"), value)
}


// The time and duration of the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutevent/dateinterval
func (h_ HKWorkoutEvent) DateInterval() foundation.objc.IObject /* cross-framework: DateInterval */ {
	rv := objc.Send[foundation.DateInterval](h_.ID, objc.Sel("dateInterval"))
	return rv
}


// The time and duration of the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutevent/dateinterval
func (h_ HKWorkoutEvent) SetDateInterval(value foundation.objc.IObject /* cross-framework: DateInterval */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setDateInterval:"), value)
}


// The metadata associated with the workout event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutevent/metadata
func (h_ HKWorkoutEvent) Metadata() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](h_.ID, objc.Sel("metadata"))
	return rv
}


// The metadata associated with the workout event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutevent/metadata
func (h_ HKWorkoutEvent) SetMetadata(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setMetadata:"), objc.String(value))
}


// The type of workout event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutevent/type
func (h_ HKWorkoutEvent) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("type"))
	return rv
}


// The type of workout event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutevent/type
func (h_ HKWorkoutEvent) SetType(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setType:"), value)
}


// The workout type identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkouttypeidentifier
func (h_ HKWorkoutEvent) HKWorkoutTypeIdentifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](h_.ID, objc.Sel("HKWorkoutTypeIdentifier"))
	return rv
}



