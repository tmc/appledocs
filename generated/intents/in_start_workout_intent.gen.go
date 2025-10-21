// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INStartWorkoutIntent] class.
var (
	INStartWorkoutIntentClass     _INStartWorkoutIntentClass
	INStartWorkoutIntentClassOnce sync.Once
)

func getINStartWorkoutIntentClass() _INStartWorkoutIntentClass {
	INStartWorkoutIntentClassOnce.Do(func() {
		INStartWorkoutIntentClass = _INStartWorkoutIntentClass{objc.GetClass("INStartWorkoutIntent")}
	})
	return INStartWorkoutIntentClass
}

type _INStartWorkoutIntentClass struct {
	class objc.Class
}

// An interface definition for the [INStartWorkoutIntent] class.
type IINStartWorkoutIntent interface {
	IINIntent
}

// A request to start a workout for the user.
//
// SiriKit creates an object when the user asks to start a workout using your app. A start workout intent identifies the user-selected workout type and goals. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object that indicates it’s possible to start the workout. For the successful handling of the intent, SiriKit launches your app and passes it an object your app must then use to start the workout.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INStartWorkoutIntent
type INStartWorkoutIntent struct {
	INIntent
}

// INStartWorkoutIntentFrom constructs a [INStartWorkoutIntent] from an unsafe.Pointer.
//
// A request to start a workout for the user.
func INStartWorkoutIntentFrom(ptr unsafe.Pointer) INStartWorkoutIntent {
	return INStartWorkoutIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INStartWorkoutIntentClass) Alloc() INStartWorkoutIntent {
	rv := objc.Send[INStartWorkoutIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INStartWorkoutIntentClass) New() INStartWorkoutIntent {
	rv := objc.Send[INStartWorkoutIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INStartWorkoutIntent) Init() INStartWorkoutIntent {
	rv := objc.Send[INStartWorkoutIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INStartWorkoutIntent) Autorelease() INStartWorkoutIntent {
	rv := objc.Send[INStartWorkoutIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINStartWorkoutIntent creates a new INStartWorkoutIntent instance.
func NewINStartWorkoutIntent() INStartWorkoutIntent {
	return getINStartWorkoutIntentClass().New()
}




// Initializes an intent object with the specified workout information.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INStartWorkoutIntent/initWithWorkoutName:goalValue:workoutGoalUnitType:workoutLocationType:isOpenEnded:
func NewINStartWorkoutIntentWithWorkoutNameGoalValueWorkoutGoalUnitTypeWorkoutLocationTypeIsOpenEnded(workoutName unsafe.Pointer, goalValue unsafe.Pointer, workoutGoalUnitType unsafe.Pointer, workoutLocationType unsafe.Pointer, isOpenEnded unsafe.Pointer) INStartWorkoutIntent {
	instance := getINStartWorkoutIntentClass().Alloc()
	rv := objc.Send[INStartWorkoutIntent](instance.ID, objc.Sel("initWithWorkoutName:goalValue:workoutGoalUnitType:workoutLocationType:isOpenEnded:"), workoutName, goalValue, workoutGoalUnitType, workoutLocationType, isOpenEnded)
	rv.Autorelease()
	return rv
}


// The units associated with the workout goal.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INStartWorkoutIntent/workoutGoalUnitType
func (i_ INStartWorkoutIntent) WorkoutGoalUnitType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("workoutGoalUnitType"))
	return rv
}

// The name of the workout.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INStartWorkoutIntent/workoutName
func (i_ INStartWorkoutIntent) WorkoutName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("workoutName"))
	return rv
}


