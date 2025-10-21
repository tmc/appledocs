// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HKWorkoutSession] class.
var (
	HKWorkoutSessionClass     _HKWorkoutSessionClass
	HKWorkoutSessionClassOnce sync.Once
)

func getHKWorkoutSessionClass() _HKWorkoutSessionClass {
	HKWorkoutSessionClassOnce.Do(func() {
		HKWorkoutSessionClass = _HKWorkoutSessionClass{objc.GetClass("HKWorkoutSession")}
	})
	return HKWorkoutSessionClass
}

type _HKWorkoutSessionClass struct {
	class objc.Class
}

// An interface definition for the [HKWorkoutSession] class.
type IHKWorkoutSession interface {
	objectivec.IObject
	End()
	EndCurrentActivityOnDate(date unsafe.Pointer)
	Pause()
}

// A session that tracks a person’s workout.
//
// The session fine-tunes Apple Watch’s sensors for the specified activity. All workout sessions generate high-frequency heart rate samples; however, an outdoor cycling activity generates accurate location data, while an indoor cycling activity doesn’t. Collecting heart rate data on iPhone or iPad requires pairing with an external heart rate sensor because these devices don’t have one. iPhone and iPad can collect various workout metrics, but the system may generate different samples than those specifically requested by an app. You can modify the default types of data collected during a workout. After someone saves a workout, you can access and display summary statistics or chart metrics over time. iPhone typically locks during workouts. For privacy reasons, health data usually isn’t accessible while the device is locked. However, the system can prompt someone to provide your app access to workout data even when their device is locked. You can then display Live Activities on the Lock Screen, providing health metrics without requiring the person to unlock their phone. Siri support extends to the Lock Screen, allowing people to start, pause, resume, or cancel workouts hands-free. You can integrate Siri intents into your apps to enable this functionality. Apple Watch runs one workout session at a time. If a second workout starts while your workout is running, your object receives an error, and your session ends.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSession
type HKWorkoutSession struct {
	objectivec.Object
}

// HKWorkoutSessionFrom constructs a [HKWorkoutSession] from an unsafe.Pointer.
//
// A session that tracks a person’s workout.
func HKWorkoutSessionFrom(ptr unsafe.Pointer) HKWorkoutSession {
	return HKWorkoutSession{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKWorkoutSessionClass) Alloc() HKWorkoutSession {
	rv := objc.Send[HKWorkoutSession](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKWorkoutSessionClass) New() HKWorkoutSession {
	rv := objc.Send[HKWorkoutSession](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKWorkoutSession) Init() HKWorkoutSession {
	rv := objc.Send[HKWorkoutSession](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKWorkoutSession) Autorelease() HKWorkoutSession {
	rv := objc.Send[HKWorkoutSession](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKWorkoutSession creates a new HKWorkoutSession instance.
func NewHKWorkoutSession() HKWorkoutSession {
	return getHKWorkoutSessionClass().New()
}


// Ends the workout session.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSession/end()
func (h_ HKWorkoutSession) End() {
	objc.Send[objc.ID](h_.ID, objc.Sel("end"))
}

// Ends the current workout activity.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSession/endCurrentActivity(on:)
func (h_ HKWorkoutSession) EndCurrentActivityOnDate(date unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("endCurrentActivityOnDate:"), date)
}

// Pauses the workout session.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSession/pause()
func (h_ HKWorkoutSession) Pause() {
	objc.Send[objc.ID](h_.ID, objc.Sel("pause"))
}

// A value that indicates whether the workout session occurred indoors or outdoors.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSession/locationType
func (h_ HKWorkoutSession) LocationType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("locationType"))
	return rv
}

// The workout session’s current state.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSession/state
func (h_ HKWorkoutSession) State() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("state"))
	return rv
}



