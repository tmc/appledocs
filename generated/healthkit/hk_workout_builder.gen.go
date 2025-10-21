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

// The list of events added to the workout.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutbuilder/workoutevents
func (h_ HKWorkoutBuilder) WorkoutEvents() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("workoutEvents"))
	return rv
}


// SetWorkoutEvents sets the value of the workoutEvents property.
// The list of events added to the workout.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutbuilder/workoutevents
func (h_ HKWorkoutBuilder) SetWorkoutEvents(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setWorkoutEvents:"), value)
}

// A dictionary that contains all the statistics for the workout builder.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutbuilder/allstatistics
func (h_ HKWorkoutBuilder) AllStatistics() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("allStatistics"))
	return rv
}


// SetAllStatistics sets the value of the allStatistics property.
// A dictionary that contains all the statistics for the workout builder.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutbuilder/allstatistics
func (h_ HKWorkoutBuilder) SetAllStatistics(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAllStatistics:"), value)
}

// The workout’s end date and time.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutbuilder/enddate
func (h_ HKWorkoutBuilder) EndDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("endDate"))
	return rv
}


// SetEndDate sets the value of the endDate property.
// The workout’s end date and time.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutbuilder/enddate
func (h_ HKWorkoutBuilder) SetEndDate(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setEndDate:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutbuilder/workoutactivities
func (h_ HKWorkoutBuilder) WorkoutActivities() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("workoutActivities"))
	return rv
}


// SetWorkoutActivities sets the value of the workoutActivities property.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutbuilder/workoutactivities
func (h_ HKWorkoutBuilder) SetWorkoutActivities(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setWorkoutActivities:"), value)
}

// The configuration information for the workout.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutbuilder/workoutconfiguration
func (h_ HKWorkoutBuilder) WorkoutConfiguration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("workoutConfiguration"))
	return rv
}


// SetWorkoutConfiguration sets the value of the workoutConfiguration property.
// The configuration information for the workout.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutbuilder/workoutconfiguration
func (h_ HKWorkoutBuilder) SetWorkoutConfiguration(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setWorkoutConfiguration:"), value)
}

// The workout type identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkouttypeidentifier
func (h_ HKWorkoutBuilder) HKWorkoutTypeIdentifier() string {
	rv := objc.Send[string](h_.ID, objc.Sel("HKWorkoutTypeIdentifier"))
	return rv
}

// The workout’s start date and time.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutbuilder/startdate
func (h_ HKWorkoutBuilder) StartDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("startDate"))
	return rv
}


// SetStartDate sets the value of the startDate property.
// The workout’s start date and time.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutbuilder/startdate
func (h_ HKWorkoutBuilder) SetStartDate(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setStartDate:"), value)
}

// The metadata the builder saves with the workout.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutbuilder/metadata
func (h_ HKWorkoutBuilder) Metadata() string {
	rv := objc.Send[string](h_.ID, objc.Sel("metadata"))
	return rv
}


// SetMetadata sets the value of the metadata property.
// The metadata the builder saves with the workout.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutbuilder/metadata
func (h_ HKWorkoutBuilder) SetMetadata(value string) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setMetadata:"), objc.String(value))
}

// The device associated with the workout.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutBuilder/device
func (h_ HKWorkoutBuilder) Device() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("device"))
	return rv
}



