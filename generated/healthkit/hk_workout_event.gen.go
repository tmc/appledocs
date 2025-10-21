// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
}

// An object representing an important event during a workout.
//
// You can use workout events to toggle a workout between an active and an inactive state, or to mark points of interest during a workout. Workouts start in an active state. A pause event switches it to an inactive state; a resume event switches it back to an active state. Adding a pause event when the workout is already inactive, or a resume event when the workout is already active, does not affect the workout’s state. These events are ignored. The lap, segment, and marker events are used to identify periods of interest during a workout. Use lap events to partition a workout into segments of equal distance. Segment events mark important periods during the workout, while markers identify important points in time.
//
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


// The type of workout event.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutEvent/type
func (h_ HKWorkoutEvent) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("type"))
	return rv
}



