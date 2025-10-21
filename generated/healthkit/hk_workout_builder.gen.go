// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HKWorkoutBuilder] class.
var (
	HKWorkoutBuilderClass     _HKWorkoutBuilderClass
	HKWorkoutBuilderClassOnce sync.Once
)

func getHKWorkoutBuilderClass() _HKWorkoutBuilderClass {
	HKWorkoutBuilderClassOnce.Do(func() {
		HKWorkoutBuilderClass = _HKWorkoutBuilderClass{objc.GetClass("HKWorkoutBuilder")}
	})
	return HKWorkoutBuilderClass
}

type _HKWorkoutBuilderClass struct {
	class objc.Class
}

// An interface definition for the [HKWorkoutBuilder] class.
type IHKWorkoutBuilder interface {
	objectivec.IObject
	AddWorkoutActivityCompletion(workoutActivity unsafe.Pointer, completion unsafe.Pointer)
	BeginCollectionWithStartDateCompletion(startDate unsafe.Pointer, completion unsafe.Pointer)
	FinishWorkoutWithCompletion(completion unsafe.Pointer)
	UpdateActivityWithUUIDEndDateCompletion(UUID unsafe.Pointer, endDate unsafe.Pointer, completion unsafe.Pointer)
}

// A builder object that incrementally constructs a workout.
//
// Incrementally collect samples and events associated with a workout. When the workout ends, call to create an sample and save it to the HealthKit store. For watchOS, use an and an instead.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutBuilder
type HKWorkoutBuilder struct {
	objectivec.Object
}

// HKWorkoutBuilderFrom constructs a [HKWorkoutBuilder] from an unsafe.Pointer.
//
// A builder object that incrementally constructs a workout.
func HKWorkoutBuilderFrom(ptr unsafe.Pointer) HKWorkoutBuilder {
	return HKWorkoutBuilder{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKWorkoutBuilderClass) Alloc() HKWorkoutBuilder {
	rv := objc.Send[HKWorkoutBuilder](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKWorkoutBuilderClass) New() HKWorkoutBuilder {
	rv := objc.Send[HKWorkoutBuilder](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKWorkoutBuilder) Init() HKWorkoutBuilder {
	rv := objc.Send[HKWorkoutBuilder](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKWorkoutBuilder) Autorelease() HKWorkoutBuilder {
	rv := objc.Send[HKWorkoutBuilder](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKWorkoutBuilder creates a new HKWorkoutBuilder instance.
func NewHKWorkoutBuilder() HKWorkoutBuilder {
	return getHKWorkoutBuilderClass().New()
}


// Adds a workout activity to the workout builder.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutBuilder/addWorkoutActivity(_:completion:)
func (h_ HKWorkoutBuilder) AddWorkoutActivityCompletion(workoutActivity unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("addWorkoutActivity:completion:"), workoutActivity, completion)
}

// Sets the workout’s start date and begins building the workout.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutBuilder/beginCollection(withStart:completion:)
func (h_ HKWorkoutBuilder) BeginCollectionWithStartDateCompletion(startDate unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("beginCollectionWithStartDate:completion:"), startDate, completion)
}

// Creates the workout, using the samples and events added to the builder, and saves it to the HealthKit store.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutBuilder/finishWorkout(completion:)
func (h_ HKWorkoutBuilder) FinishWorkoutWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("finishWorkoutWithCompletion:"), completion)
}

// Sets the end date for a workout activity that you’ve already added to the workout builder.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutBuilder/updateActivity(uuid:end:completion:)
func (h_ HKWorkoutBuilder) UpdateActivityWithUUIDEndDateCompletion(UUID unsafe.Pointer, endDate unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("updateActivityWithUUID:endDate:completion:"), UUID, endDate, completion)
}

// The device associated with the workout.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutBuilder/device
func (h_ HKWorkoutBuilder) Device() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("device"))
	return rv
}



