// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INEndWorkoutIntent] class.
var (
	INEndWorkoutIntentClass     _INEndWorkoutIntentClass
	INEndWorkoutIntentClassOnce sync.Once
)

func getINEndWorkoutIntentClass() _INEndWorkoutIntentClass {
	INEndWorkoutIntentClassOnce.Do(func() {
		INEndWorkoutIntentClass = _INEndWorkoutIntentClass{objc.GetClass("INEndWorkoutIntent")}
	})
	return INEndWorkoutIntentClass
}

type _INEndWorkoutIntentClass struct {
	class objc.Class
}

// An interface definition for the [INEndWorkoutIntent] class.
type IINEndWorkoutIntent interface {
	IINIntent
	WorkoutName() INSpeakableString
	SetWorkoutName(value INSpeakableString)
}

// A request to end the current workout that also validates workout parameters and saves the results.
//
// SiriKit creates an object when the user asks to finish an in-progress workout. Finishing a workout stops it and records any progress made toward the workout’s goals. Use this intent object to validate the workout parameters. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object that indicates it’s possible to end the workout. For the successful handling of the intent, SiriKit launches your app and passes it an object your app must then use to end the workout.

// A request to end the current workout that also validates workout parameters and saves the results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INEndWorkoutIntent
type INEndWorkoutIntent struct {
	INIntent
}

// INEndWorkoutIntentFrom constructs a [INEndWorkoutIntent] from an unsafe.Pointer.
//
// A request to end the current workout that also validates workout parameters and saves the results.
func INEndWorkoutIntentFrom(ptr unsafe.Pointer) INEndWorkoutIntent {
	return INEndWorkoutIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INEndWorkoutIntentClass) Alloc() INEndWorkoutIntent {
	rv := objc.Send[INEndWorkoutIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INEndWorkoutIntentClass) New() INEndWorkoutIntent {
	rv := objc.Send[INEndWorkoutIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INEndWorkoutIntent) Init() INEndWorkoutIntent {
	rv := objc.Send[INEndWorkoutIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INEndWorkoutIntent) Autorelease() INEndWorkoutIntent {
	rv := objc.Send[INEndWorkoutIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINEndWorkoutIntent creates a new INEndWorkoutIntent instance.
func NewINEndWorkoutIntent() INEndWorkoutIntent {
	return getINEndWorkoutIntentClass().New()
}

// Initializes an intent object with the specified workout name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INEndWorkoutIntent/init(workoutName:)
func NewINEndWorkoutIntentWithWorkoutName(workoutName INSpeakableString) INEndWorkoutIntent {
	instance := getINEndWorkoutIntentClass().Alloc()
	rv := objc.Send[INEndWorkoutIntent](instance.ID, objc.Sel("initWithWorkoutName:"), workoutName)
	rv.Autorelease()
	return rv
}

// The name of the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inendworkoutintent/workoutname
func (i_ INEndWorkoutIntent) WorkoutName() INSpeakableString {
	rv := objc.Send[INSpeakableString](i_.ID, objc.Sel("workoutName"))
	return rv
}

// The name of the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inendworkoutintent/workoutname
func (i_ INEndWorkoutIntent) SetWorkoutName(value INSpeakableString) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setWorkoutName:"), value)
}
