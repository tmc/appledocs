// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	AllStatistics() IHKStatistics
	SetAllStatistics(value IHKStatistics)
	Device() IHKDevice
	SetDevice(value IHKDevice)
	EndDate() foundation.objc.IObject /* cross-framework: Date */
	SetEndDate(value foundation.objc.IObject /* cross-framework: Date */)
	Metadata() string /* primitive/slice/pointer. */
	SetMetadata(value string /* primitive/slice/pointer. */)
	StartDate() foundation.objc.IObject /* cross-framework: Date */
	SetStartDate(value foundation.objc.IObject /* cross-framework: Date */)
	WorkoutActivities() IHKWorkoutActivity
	SetWorkoutActivities(value IHKWorkoutActivity)
	WorkoutConfiguration() IHKWorkoutConfiguration
	SetWorkoutConfiguration(value IHKWorkoutConfiguration)
	WorkoutEvents() IHKWorkoutEvent
	SetWorkoutEvents(value IHKWorkoutEvent)
	HKWorkoutTypeIdentifier() string /* primitive/slice/pointer. */
	// methods:
}

// A builder object that incrementally constructs a workout.
//
// Incrementally collect samples and events associated with a workout. When the workout ends, call to create an sample and save it to the HealthKit store. For watchOS, use an and an instead.


// A builder object that incrementally constructs a workout.
//
// [Full Topic]
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



// A dictionary that contains all the statistics for the workout builder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutbuilder/allstatistics
func (h_ HKWorkoutBuilder) AllStatistics() IHKStatistics {
	rv := objc.Send[HKStatistics](h_.ID, objc.Sel("allStatistics"))
	return rv
}


// A dictionary that contains all the statistics for the workout builder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutbuilder/allstatistics
func (h_ HKWorkoutBuilder) SetAllStatistics(value IHKStatistics) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAllStatistics:"), value)
}


// The device associated with the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutbuilder/device
func (h_ HKWorkoutBuilder) Device() IHKDevice {
	rv := objc.Send[HKDevice](h_.ID, objc.Sel("device"))
	return rv
}


// The device associated with the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutbuilder/device
func (h_ HKWorkoutBuilder) SetDevice(value IHKDevice) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setDevice:"), value)
}


// The workout’s end date and time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutbuilder/enddate
func (h_ HKWorkoutBuilder) EndDate() foundation.objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](h_.ID, objc.Sel("endDate"))
	return rv
}


// The workout’s end date and time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutbuilder/enddate
func (h_ HKWorkoutBuilder) SetEndDate(value foundation.objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setEndDate:"), value)
}


// The metadata the builder saves with the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutbuilder/metadata
func (h_ HKWorkoutBuilder) Metadata() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](h_.ID, objc.Sel("metadata"))
	return rv
}


// The metadata the builder saves with the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutbuilder/metadata
func (h_ HKWorkoutBuilder) SetMetadata(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setMetadata:"), objc.String(value))
}


// The workout’s start date and time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutbuilder/startdate
func (h_ HKWorkoutBuilder) StartDate() foundation.objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](h_.ID, objc.Sel("startDate"))
	return rv
}


// The workout’s start date and time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutbuilder/startdate
func (h_ HKWorkoutBuilder) SetStartDate(value foundation.objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setStartDate:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutbuilder/workoutactivities
func (h_ HKWorkoutBuilder) WorkoutActivities() IHKWorkoutActivity {
	rv := objc.Send[HKWorkoutActivity](h_.ID, objc.Sel("workoutActivities"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutbuilder/workoutactivities
func (h_ HKWorkoutBuilder) SetWorkoutActivities(value IHKWorkoutActivity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setWorkoutActivities:"), value)
}


// The configuration information for the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutbuilder/workoutconfiguration
func (h_ HKWorkoutBuilder) WorkoutConfiguration() IHKWorkoutConfiguration {
	rv := objc.Send[HKWorkoutConfiguration](h_.ID, objc.Sel("workoutConfiguration"))
	return rv
}


// The configuration information for the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutbuilder/workoutconfiguration
func (h_ HKWorkoutBuilder) SetWorkoutConfiguration(value IHKWorkoutConfiguration) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setWorkoutConfiguration:"), value)
}


// The list of events added to the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutbuilder/workoutevents
func (h_ HKWorkoutBuilder) WorkoutEvents() IHKWorkoutEvent {
	rv := objc.Send[HKWorkoutEvent](h_.ID, objc.Sel("workoutEvents"))
	return rv
}


// The list of events added to the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutbuilder/workoutevents
func (h_ HKWorkoutBuilder) SetWorkoutEvents(value IHKWorkoutEvent) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setWorkoutEvents:"), value)
}


// The workout type identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkouttypeidentifier
func (h_ HKWorkoutBuilder) HKWorkoutTypeIdentifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](h_.ID, objc.Sel("HKWorkoutTypeIdentifier"))
	return rv
}



