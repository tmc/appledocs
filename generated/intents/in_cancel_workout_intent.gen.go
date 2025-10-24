// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INCancelWorkoutIntent] class.
var (
	INCancelWorkoutIntentClass     _INCancelWorkoutIntentClass
	INCancelWorkoutIntentClassOnce sync.Once
)

func getINCancelWorkoutIntentClass() _INCancelWorkoutIntentClass {
	INCancelWorkoutIntentClassOnce.Do(func() {
		INCancelWorkoutIntentClass = _INCancelWorkoutIntentClass{objc.GetClass("INCancelWorkoutIntent")}
	})
	return INCancelWorkoutIntentClass
}

type _INCancelWorkoutIntentClass struct {
	class objc.Class
}

// An interface definition for the [INCancelWorkoutIntent] class.
type IINCancelWorkoutIntent interface {
	IINIntent
	WorkoutName() INSpeakableString
	SetWorkoutName(value INSpeakableString)
}

// A request to cancel an active workout.
//
// SiriKit creates an object when the user asks to cancel a currently running workout. Canceling a workout stops the workout and doesn’t record any progress toward the workout goals. Use this intent object to get the workout parameters. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object to indicate it’s possible to cancel the workout. For the successful handling of the intent, SiriKit launches your app and passes it an object that your app must then use to cancel the workout.

// A request to cancel an active workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INCancelWorkoutIntent
type INCancelWorkoutIntent struct {
	INIntent
}

// INCancelWorkoutIntentFrom constructs a [INCancelWorkoutIntent] from an unsafe.Pointer.
//
// A request to cancel an active workout.
func INCancelWorkoutIntentFrom(ptr unsafe.Pointer) INCancelWorkoutIntent {
	return INCancelWorkoutIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INCancelWorkoutIntentClass) Alloc() INCancelWorkoutIntent {
	rv := objc.Send[INCancelWorkoutIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INCancelWorkoutIntentClass) New() INCancelWorkoutIntent {
	rv := objc.Send[INCancelWorkoutIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INCancelWorkoutIntent) Init() INCancelWorkoutIntent {
	rv := objc.Send[INCancelWorkoutIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INCancelWorkoutIntent) Autorelease() INCancelWorkoutIntent {
	rv := objc.Send[INCancelWorkoutIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINCancelWorkoutIntent creates a new INCancelWorkoutIntent instance.
func NewINCancelWorkoutIntent() INCancelWorkoutIntent {
	return getINCancelWorkoutIntentClass().New()
}

// The name of the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/incancelworkoutintent/workoutname
func (i_ INCancelWorkoutIntent) WorkoutName() INSpeakableString {
	rv := objc.Send[INSpeakableString](i_.ID, objc.Sel("workoutName"))
	return rv
}

// The name of the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/incancelworkoutintent/workoutname
func (i_ INCancelWorkoutIntent) SetWorkoutName(value INSpeakableString) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setWorkoutName:"), value)
}
