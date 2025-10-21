// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INResumeWorkoutIntent] class.
var (
	INResumeWorkoutIntentClass     _INResumeWorkoutIntentClass
	INResumeWorkoutIntentClassOnce sync.Once
)

func getINResumeWorkoutIntentClass() _INResumeWorkoutIntentClass {
	INResumeWorkoutIntentClassOnce.Do(func() {
		INResumeWorkoutIntentClass = _INResumeWorkoutIntentClass{objc.GetClass("INResumeWorkoutIntent")}
	})
	return INResumeWorkoutIntentClass
}

type _INResumeWorkoutIntentClass struct {
	class objc.Class
}

// An interface definition for the [INResumeWorkoutIntent] class.
type IINResumeWorkoutIntent interface {
	IINIntent
}

// A request to resume a paused workout.
//
// SiriKit creates an object when the user asks to resume a currently paused workout. Resuming a workout resumes gathering workout data and counting that data toward the user’s current workout goal. Use this intent object to validate the workout parameters. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object that indicates it’s possible to resume the workout. For the successful handling of the intent, SiriKit launches your app and passes it an object your app must then use to resume the workout.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INResumeWorkoutIntent
type INResumeWorkoutIntent struct {
	INIntent
}

// INResumeWorkoutIntentFrom constructs a [INResumeWorkoutIntent] from an unsafe.Pointer.
//
// A request to resume a paused workout.
func INResumeWorkoutIntentFrom(ptr unsafe.Pointer) INResumeWorkoutIntent {
	return INResumeWorkoutIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INResumeWorkoutIntentClass) Alloc() INResumeWorkoutIntent {
	rv := objc.Send[INResumeWorkoutIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INResumeWorkoutIntentClass) New() INResumeWorkoutIntent {
	rv := objc.Send[INResumeWorkoutIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INResumeWorkoutIntent) Init() INResumeWorkoutIntent {
	rv := objc.Send[INResumeWorkoutIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INResumeWorkoutIntent) Autorelease() INResumeWorkoutIntent {
	rv := objc.Send[INResumeWorkoutIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINResumeWorkoutIntent creates a new INResumeWorkoutIntent instance.
func NewINResumeWorkoutIntent() INResumeWorkoutIntent {
	return getINResumeWorkoutIntentClass().New()
}


// The name of the workout.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inresumeworkoutintent/workoutname
func (i_ INResumeWorkoutIntent) WorkoutName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("workoutName"))
	return rv
}


// SetWorkoutName sets the value of the workoutName property.
// The name of the workout.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inresumeworkoutintent/workoutname
func (i_ INResumeWorkoutIntent) SetWorkoutName(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setWorkoutName:"), value)
}



