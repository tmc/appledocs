// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INPauseWorkoutIntent] class.
var (
	INPauseWorkoutIntentClass     _INPauseWorkoutIntentClass
	INPauseWorkoutIntentClassOnce sync.Once
)

func getINPauseWorkoutIntentClass() _INPauseWorkoutIntentClass {
	INPauseWorkoutIntentClassOnce.Do(func() {
		INPauseWorkoutIntentClass = _INPauseWorkoutIntentClass{objc.GetClass("INPauseWorkoutIntent")}
	})
	return INPauseWorkoutIntentClass
}

type _INPauseWorkoutIntentClass struct {
	class objc.Class
}

// An interface definition for the [INPauseWorkoutIntent] class.
type IINPauseWorkoutIntent interface {
	IINIntent
}

// A request to pause the current workout that also stops the gathering of workout data.
//
// SiriKit creates an object when the user asks to pause a currently running workout. Pausing a workout stops the gathering of workout data without clearing the current workout progress information. Use this intent object to validate workout parameters. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object that indicates it’s possible to pause the workout. SiriKit launches your app and passes it an object that your app must then use to pause the workout.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INPauseWorkoutIntent
type INPauseWorkoutIntent struct {
	INIntent
}

// INPauseWorkoutIntentFrom constructs a [INPauseWorkoutIntent] from an unsafe.Pointer.
//
// A request to pause the current workout that also stops the gathering of workout data.
func INPauseWorkoutIntentFrom(ptr unsafe.Pointer) INPauseWorkoutIntent {
	return INPauseWorkoutIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INPauseWorkoutIntentClass) Alloc() INPauseWorkoutIntent {
	rv := objc.Send[INPauseWorkoutIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INPauseWorkoutIntentClass) New() INPauseWorkoutIntent {
	rv := objc.Send[INPauseWorkoutIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INPauseWorkoutIntent) Init() INPauseWorkoutIntent {
	rv := objc.Send[INPauseWorkoutIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INPauseWorkoutIntent) Autorelease() INPauseWorkoutIntent {
	rv := objc.Send[INPauseWorkoutIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINPauseWorkoutIntent creates a new INPauseWorkoutIntent instance.
func NewINPauseWorkoutIntent() INPauseWorkoutIntent {
	return getINPauseWorkoutIntentClass().New()
}


// The name of the workout to pause.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INPauseWorkoutIntent/workoutName
func (i_ INPauseWorkoutIntent) WorkoutName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("workoutName"))
	return rv
}



