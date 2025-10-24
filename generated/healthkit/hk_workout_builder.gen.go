// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKWorkoutBuilder */


/* debug [class_header]: Header for HKWorkoutBuilder */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKWorkoutBuilder */
// An interface definition for the [HKWorkoutBuilder] class.
type IHKWorkoutBuilder interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKWorkoutBuilder */
	// properties:
	AllStatistics() foundation.IDictionary
	Device() IHKDevice
	EndDate() objc.IObject /* cross-framework: NSDate */
	Metadata() foundation.IDictionary
	StartDate() objc.IObject /* cross-framework: NSDate */
	WorkoutActivities() []HKWorkoutActivity
	WorkoutConfiguration() IHKWorkoutConfiguration
	WorkoutEvents() []HKWorkoutEvent
	HKWorkoutTypeIdentifier() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKWorkoutBuilder */
	// methods:
	AddSamplesCompletion(samples []HKSample, completion unsafe.Pointer)
	AddMetadataCompletion(metadata foundation.IDictionary, completion unsafe.Pointer)
	AddWorkoutActivityCompletion(workoutActivity IHKWorkoutActivity, completion unsafe.Pointer)
	AddWorkoutEventsCompletion(workoutEvents []HKWorkoutEvent, completion unsafe.Pointer)
	BeginCollectionWithStartDateCompletion(startDate objc.IObject /* cross-framework: NSDate */, completion unsafe.Pointer)
	DiscardWorkout()
	ElapsedTimeAtDate(date objc.IObject /* cross-framework: NSDate */) float64
	EndCollectionWithEndDateCompletion(endDate objc.IObject /* cross-framework: NSDate */, completion unsafe.Pointer)
	FinishWorkoutWithCompletion(completion unsafe.Pointer)
	SeriesBuilderForType(seriesType IHKSeriesType) IHKSeriesBuilder
	StatisticsForType(quantityType IHKQuantityType) IHKStatistics
	UpdateActivityWithUUIDAddMedatataCompletion(UUID foundation.UUID, metadata foundation.IDictionary, completion unsafe.Pointer)
	UpdateActivityWithUUIDEndDateCompletion(UUID foundation.UUID, endDate objc.IObject /* cross-framework: NSDate */, completion unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKWorkoutBuilder */
// Alloc allocates a new instance without initialization.
func (hc _HKWorkoutBuilderClass) Alloc() HKWorkoutBuilder {
	rv := objc.Send[HKWorkoutBuilder](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKWorkoutBuilder */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKWorkoutBuilder */

// Returns a new workout builder object that is not connected to a workout session or other data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutBuilder/init(healthStore:configuration:device:)
func NewHKWorkoutBuilderWithHealthStoreConfigurationDevice(healthStore IHKHealthStore, configuration IHKWorkoutConfiguration, device IHKDevice) HKWorkoutBuilder {
	instance := getHKWorkoutBuilderClass().Alloc()
	rv := objc.Send[HKWorkoutBuilder](instance.ID, objc.Sel("initWithHealthStore:configuration:device:"), healthStore, configuration, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKWorkoutBuilderWithHealthStoreConfigurationDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKWorkoutBuilder */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKWorkoutBuilder */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKWorkoutBuilder */

// Adds a sample to be associated with the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutBuilder/add(_:completion:)
func (h_ HKWorkoutBuilder) AddSamplesCompletion(samples []HKSample, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("addSamples:completion:"), samples, completion)
}/* debug [instance_methods/method]: AddSamplesCompletion */


// Adds metadata to be saved with the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutBuilder/addMetadata(_:completion:)
func (h_ HKWorkoutBuilder) AddMetadataCompletion(metadata foundation.IDictionary, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("addMetadata:completion:"), metadata, completion)
}/* debug [instance_methods/method]: AddMetadataCompletion */


// Adds a workout activity to the workout builder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutBuilder/addWorkoutActivity(_:completion:)
func (h_ HKWorkoutBuilder) AddWorkoutActivityCompletion(workoutActivity IHKWorkoutActivity, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("addWorkoutActivity:completion:"), workoutActivity, completion)
}/* debug [instance_methods/method]: AddWorkoutActivityCompletion */


// Adds a workout event to the builder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutBuilder/addWorkoutEvents(_:completion:)
func (h_ HKWorkoutBuilder) AddWorkoutEventsCompletion(workoutEvents []HKWorkoutEvent, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("addWorkoutEvents:completion:"), workoutEvents, completion)
}/* debug [instance_methods/method]: AddWorkoutEventsCompletion */


// Sets the workout’s start date and begins building the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutBuilder/beginCollection(withStart:completion:)
func (h_ HKWorkoutBuilder) BeginCollectionWithStartDateCompletion(startDate objc.IObject /* cross-framework: NSDate */, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("beginCollectionWithStartDate:completion:"), startDate, completion)
}/* debug [instance_methods/method]: BeginCollectionWithStartDateCompletion */


// Stops the collection of data and discards the current results without saving the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutBuilder/discardWorkout()
func (h_ HKWorkoutBuilder) DiscardWorkout() {
	objc.Send[objc.ID](h_.ID, objc.Sel("discardWorkout"))
}/* debug [instance_methods/method]: DiscardWorkout */


// Calculates the duration of the workout at the specified time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutBuilder/elapsedTime(at:)
func (h_ HKWorkoutBuilder) ElapsedTimeAtDate(date objc.IObject /* cross-framework: NSDate */) float64 {
	rv := objc.Send[float64](h_.ID, objc.Sel("elapsedTimeAtDate:"), date)
	return rv
}/* debug [instance_methods/method]: ElapsedTimeAtDate */


// Stops the collection of data, sets the workout’s end date, and deactivates the workout builder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutBuilder/endCollection(withEnd:completion:)
func (h_ HKWorkoutBuilder) EndCollectionWithEndDateCompletion(endDate objc.IObject /* cross-framework: NSDate */, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("endCollectionWithEndDate:completion:"), endDate, completion)
}/* debug [instance_methods/method]: EndCollectionWithEndDateCompletion */


// Creates the workout, using the samples and events added to the builder, and saves it to the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutBuilder/finishWorkout(completion:)
func (h_ HKWorkoutBuilder) FinishWorkoutWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("finishWorkoutWithCompletion:"), completion)
}/* debug [instance_methods/method]: FinishWorkoutWithCompletion */


// Returns the series builder for the specified type, creating a new builder, if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutBuilder/seriesBuilder(for:)
func (h_ HKWorkoutBuilder) SeriesBuilderForType(seriesType IHKSeriesType) IHKSeriesBuilder {
	rv := objc.Send[HKSeriesBuilder](h_.ID, objc.Sel("seriesBuilderForType:"), seriesType)
	return rv
}/* debug [instance_methods/method]: SeriesBuilderForType */


// Returns the statistics calculated for matching samples added to the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutBuilder/statistics(for:)
func (h_ HKWorkoutBuilder) StatisticsForType(quantityType IHKQuantityType) IHKStatistics {
	rv := objc.Send[HKStatistics](h_.ID, objc.Sel("statisticsForType:"), quantityType)
	return rv
}/* debug [instance_methods/method]: StatisticsForType */


// Adds metadata to a workout activity that you’ve already added to the workout builder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutBuilder/updateActivity(uuid:adding:completion:)
func (h_ HKWorkoutBuilder) UpdateActivityWithUUIDAddMedatataCompletion(UUID foundation.UUID, metadata foundation.IDictionary, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("updateActivityWithUUID:addMedatata:completion:"), UUID, metadata, completion)
}/* debug [instance_methods/method]: UpdateActivityWithUUIDAddMedatataCompletion */


// Sets the end date for a workout activity that you’ve already added to the workout builder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutBuilder/updateActivity(uuid:end:completion:)
func (h_ HKWorkoutBuilder) UpdateActivityWithUUIDEndDateCompletion(UUID foundation.UUID, endDate objc.IObject /* cross-framework: NSDate */, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("updateActivityWithUUID:endDate:completion:"), UUID, endDate, completion)
}/* debug [instance_methods/method]: UpdateActivityWithUUIDEndDateCompletion */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKWorkoutBuilder */

// A dictionary that contains all the statistics for the workout builder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutBuilder/allStatistics
func (h_ HKWorkoutBuilder) AllStatistics() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](h_.ID, objc.Sel("allStatistics"))
	return rv
}/* debug [instance_properties/getter]: allStatistics */


// The device associated with the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutBuilder/device
func (h_ HKWorkoutBuilder) Device() IHKDevice {
	rv := objc.Send[HKDevice](h_.ID, objc.Sel("device"))
	return rv
}/* debug [instance_properties/getter]: device */


// The workout’s end date and time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutBuilder/endDate
func (h_ HKWorkoutBuilder) EndDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](h_.ID, objc.Sel("endDate"))
	return rv
}/* debug [instance_properties/getter]: endDate */


// The metadata the builder saves with the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutBuilder/metadata
func (h_ HKWorkoutBuilder) Metadata() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](h_.ID, objc.Sel("metadata"))
	return rv
}/* debug [instance_properties/getter]: metadata */


// The workout’s start date and time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutBuilder/startDate
func (h_ HKWorkoutBuilder) StartDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](h_.ID, objc.Sel("startDate"))
	return rv
}/* debug [instance_properties/getter]: startDate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutBuilder/workoutActivities
func (h_ HKWorkoutBuilder) WorkoutActivities() []HKWorkoutActivity {
	rv := objc.Send[[]HKWorkoutActivity](h_.ID, objc.Sel("workoutActivities"))
	return rv
}/* debug [instance_properties/getter]: workoutActivities */


// The configuration information for the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutBuilder/workoutConfiguration
func (h_ HKWorkoutBuilder) WorkoutConfiguration() IHKWorkoutConfiguration {
	rv := objc.Send[HKWorkoutConfiguration](h_.ID, objc.Sel("workoutConfiguration"))
	return rv
}/* debug [instance_properties/getter]: workoutConfiguration */


// The list of events added to the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutBuilder/workoutEvents
func (h_ HKWorkoutBuilder) WorkoutEvents() []HKWorkoutEvent {
	rv := objc.Send[[]HKWorkoutEvent](h_.ID, objc.Sel("workoutEvents"))
	return rv
}/* debug [instance_properties/getter]: workoutEvents */


// The workout type identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkouttypeidentifier
func (h_ HKWorkoutBuilder) HKWorkoutTypeIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKWorkoutTypeIdentifier"))
	return rv
}/* debug [instance_properties/getter]: HKWorkoutTypeIdentifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKWorkoutBuilder */


