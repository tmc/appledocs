// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	WorkoutGoalUnitType() unsafe.Pointer
	WorkoutName() INSpeakableString
	GoalValue() float64
	SetGoalValue(value float64)
	IsOpenEnded() bool
	SetIsOpenEnded(value bool)
	WorkoutLocationType() unsafe.Pointer
	SetWorkoutLocationType(value unsafe.Pointer)
}

// A request to start a workout for the user.
//
// SiriKit creates an object when the user asks to start a workout using your app. A start workout intent identifies the user-selected workout type and goals. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object that indicates it’s possible to start the workout. For the successful handling of the intent, SiriKit launches your app and passes it an object your app must then use to start the workout.


// A request to start a workout for the user.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INStartWorkoutIntent/initWithWorkoutName:goalValue:workoutGoalUnitType:workoutLocationType:isOpenEnded:
func NewINStartWorkoutIntentWithWorkoutNameGoalValueWorkoutGoalUnitTypeWorkoutLocationTypeIsOpenEnded(workoutName INSpeakableString, goalValue foundation.INumber, workoutGoalUnitType unsafe.Pointer, workoutLocationType unsafe.Pointer, isOpenEnded foundation.INumber) INStartWorkoutIntent {
	instance := getINStartWorkoutIntentClass().Alloc()
	rv := objc.Send[INStartWorkoutIntent](instance.ID, objc.Sel("initWithWorkoutName:goalValue:workoutGoalUnitType:workoutLocationType:isOpenEnded:"), workoutName, goalValue, workoutGoalUnitType, workoutLocationType, isOpenEnded)
	rv.Autorelease()
	return rv
}



// The units associated with the workout goal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INStartWorkoutIntent/workoutGoalUnitType
func (i_ INStartWorkoutIntent) WorkoutGoalUnitType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("workoutGoalUnitType"))
	return rv
}


// The name of the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INStartWorkoutIntent/workoutName
func (i_ INStartWorkoutIntent) WorkoutName() INSpeakableString {
	rv := objc.Send[INSpeakableString](i_.ID, objc.Sel("workoutName"))
	return rv
}


// The user-supplied numerical goal of the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartworkoutintent/goalvalue-5oazy
func (i_ INStartWorkoutIntent) GoalValue() float64 {
	rv := objc.Send[float64](i_.ID, objc.Sel("goalValue"))
	return rv
}


// The user-supplied numerical goal of the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartworkoutintent/goalvalue-5oazy
func (i_ INStartWorkoutIntent) SetGoalValue(value float64) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setGoalValue:"), value)
}


// A Boolean value that indicates whether the workout is open ended.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartworkoutintent/isopenended-8hecn
func (i_ INStartWorkoutIntent) IsOpenEnded() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isOpenEnded"))
	return rv
}


// A Boolean value that indicates whether the workout is open ended.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartworkoutintent/isopenended-8hecn
func (i_ INStartWorkoutIntent) SetIsOpenEnded(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsOpenEnded:"), value)
}


// The location of the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartworkoutintent/workoutlocationtype
func (i_ INStartWorkoutIntent) WorkoutLocationType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("workoutLocationType"))
	return rv
}


// The location of the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartworkoutintent/workoutlocationtype
func (i_ INStartWorkoutIntent) SetWorkoutLocationType(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setWorkoutLocationType:"), value)
}


