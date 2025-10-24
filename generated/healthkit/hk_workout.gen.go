// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKWorkout */


/* debug [class_header]: Header for HKWorkout */
// The class instance for the [HKWorkout] class.
var (
	HKWorkoutClass     _HKWorkoutClass
	HKWorkoutClassOnce sync.Once
)

func getHKWorkoutClass() _HKWorkoutClass {
	HKWorkoutClassOnce.Do(func() {
		HKWorkoutClass = _HKWorkoutClass{objc.GetClass("HKWorkout")}
	})
	return HKWorkoutClass
}

type _HKWorkoutClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKWorkout */
// An interface definition for the [HKWorkout] class.
type IHKWorkout interface {
	IHKSample
	
/* debug [class_interface_properties]: Properties for HKWorkout */
	// properties:
	AllStatistics() foundation.IDictionary
	Duration() float64
	TotalDistance() IHKQuantity
	TotalEnergyBurned() IHKQuantity
	TotalFlightsClimbed() IHKQuantity
	TotalSwimmingStrokeCount() IHKQuantity
	WorkoutActivities() []HKWorkoutActivity
	WorkoutActivityType() HKWorkoutActivityType
	WorkoutEvents() []HKWorkoutEvent
	HKPredicateKeyPathWorkoutAverageQuantity() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathWorkoutDuration() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathWorkoutMaximumQuantity() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathWorkoutMinimumQuantity() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathWorkoutSumQuantity() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathWorkoutTotalDistance() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathWorkoutTotalEnergyBurned() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathWorkoutType() objc.IObject /* cross-framework: NSString */
	WorkoutPlan() objectivec.IObject
	SetWorkoutPlan(value objectivec.IObject)
	HKWorkoutSortIdentifierDuration() objc.IObject /* cross-framework: NSString */
	HKWorkoutSortIdentifierTotalDistance() objc.IObject /* cross-framework: NSString */
	HKWorkoutSortIdentifierTotalEnergyBurned() objc.IObject /* cross-framework: NSString */
	HKWorkoutTypeIdentifier() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKWorkout */
	// methods:
	StatisticsForType(quantityType IHKQuantityType) IHKStatistics
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKWorkout */
// Alloc allocates a new instance without initialization.
func (hc _HKWorkoutClass) Alloc() HKWorkout {
	rv := objc.Send[HKWorkout](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKWorkoutClass) New() HKWorkout {
	rv := objc.Send[HKWorkout](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKWorkout) Init() HKWorkout {
	rv := objc.Send[HKWorkout](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKWorkout) Autorelease() HKWorkout {
	rv := objc.Send[HKWorkout](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKWorkout creates a new HKWorkout instance.
func NewHKWorkout() HKWorkout {
	return getHKWorkoutClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKWorkout */
// A workout sample that stores information about a single physical activity.
//
// The class is a concrete subclass of the class; however, they behave somewhat differently than other sample types. You don’t need a specific type identifier to create the instance. All workouts use the same type identifier. You must provide an value for each workout. This value defines the type of activity performed during the workout. After saving the workout to the HealthKit store, you must associate additional samples with the workout (for example, active energy burned or distance samples). These samples provide fine-grained details. Use the method to associate them with the workout. The workout records a summary of information about a single physical activity (for example, the duration, total distance, and total energy burned). It also acts as a container for other objects. You can associate any number of samples with a workout, adding details over the course of the workout. For example, you may want to break a single run into a number of shorter intervals, and then add samples to track the user’s heart rate, energy burned, distance traveled, and steps taken for each interval. For more information, see . HealthKit supports a wide range of activity types. For a complete list, see . Workouts are mostly immutable. You set their properties when you instantiate the workout, and they can’t change. However, you can continue to add samples to the workouts.


// A workout sample that stores information about a single physical activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout
type HKWorkout struct {
	HKSample
}

// HKWorkoutFrom constructs a [HKWorkout] from an unsafe.Pointer.
//
// A workout sample that stores information about a single physical activity.
func HKWorkoutFrom(ptr unsafe.Pointer) HKWorkout {
	return HKWorkout{
		HKSample: HKSampleFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKWorkout */

// Instantiates a new workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/init(activityType:start:end:)
func NewHKWorkoutWithActivityTypeStartDateEndDate(workoutActivityType HKWorkoutActivityType, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */) HKWorkout {
	rv := objc.Send[HKWorkout](objc.ID(getHKWorkoutClass().class), objc.Sel("workoutWithActivityType:startDate:endDate:"), workoutActivityType, startDate, endDate)
	return rv
}/* debug [class_init_methods/constructor]: NewHKWorkoutWithActivityTypeStartDateEndDate */


// Instantiates a new workout activity that includes the device that produced the sample data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/init(activityType:start:end:duration:totalEnergyBurned:totalDistance:device:metadata:)
func NewHKWorkoutWithActivityTypeStartDateEndDateDurationTotalEnergyBurnedTotalDistanceDeviceMetadata(workoutActivityType HKWorkoutActivityType, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, duration float64, totalEnergyBurned IHKQuantity, totalDistance IHKQuantity, device IHKDevice, metadata foundation.IDictionary) HKWorkout {
	rv := objc.Send[HKWorkout](objc.ID(getHKWorkoutClass().class), objc.Sel("workoutWithActivityType:startDate:endDate:duration:totalEnergyBurned:totalDistance:device:metadata:"), workoutActivityType, startDate, endDate, duration, totalEnergyBurned, totalDistance, device, metadata)
	return rv
}/* debug [class_init_methods/constructor]: NewHKWorkoutWithActivityTypeStartDateEndDateDurationTotalEnergyBurnedTotalDistanceDeviceMetadata */


// Instantiates a new workout that includes the energy burned, distance, and metadata for the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/init(activityType:start:end:duration:totalEnergyBurned:totalDistance:metadata:)
func NewHKWorkoutWithActivityTypeStartDateEndDateDurationTotalEnergyBurnedTotalDistanceMetadata(workoutActivityType HKWorkoutActivityType, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, duration float64, totalEnergyBurned IHKQuantity, totalDistance IHKQuantity, metadata foundation.IDictionary) HKWorkout {
	rv := objc.Send[HKWorkout](objc.ID(getHKWorkoutClass().class), objc.Sel("workoutWithActivityType:startDate:endDate:duration:totalEnergyBurned:totalDistance:metadata:"), workoutActivityType, startDate, endDate, duration, totalEnergyBurned, totalDistance, metadata)
	return rv
}/* debug [class_init_methods/constructor]: NewHKWorkoutWithActivityTypeStartDateEndDateDurationTotalEnergyBurnedTotalDistanceMetadata */


// Instantiates a workout that includes both workout events and the device that produced the sample data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/init(activityType:start:end:workoutEvents:totalEnergyBurned:totalDistance:device:metadata:)
func NewHKWorkoutWithActivityTypeStartDateEndDateWorkoutEventsTotalEnergyBurnedTotalDistanceDeviceMetadata(workoutActivityType HKWorkoutActivityType, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, workoutEvents []HKWorkoutEvent, totalEnergyBurned IHKQuantity, totalDistance IHKQuantity, device IHKDevice, metadata foundation.IDictionary) HKWorkout {
	rv := objc.Send[HKWorkout](objc.ID(getHKWorkoutClass().class), objc.Sel("workoutWithActivityType:startDate:endDate:workoutEvents:totalEnergyBurned:totalDistance:device:metadata:"), workoutActivityType, startDate, endDate, workoutEvents, totalEnergyBurned, totalDistance, device, metadata)
	return rv
}/* debug [class_init_methods/constructor]: NewHKWorkoutWithActivityTypeStartDateEndDateWorkoutEventsTotalEnergyBurnedTotalDistanceDeviceMetadata */


// Instantiates a new workout whose duration is calculated based on the start and end dates and the provided workout events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/init(activityType:start:end:workoutEvents:totalEnergyBurned:totalDistance:metadata:)
func NewHKWorkoutWithActivityTypeStartDateEndDateWorkoutEventsTotalEnergyBurnedTotalDistanceMetadata(workoutActivityType HKWorkoutActivityType, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, workoutEvents []HKWorkoutEvent, totalEnergyBurned IHKQuantity, totalDistance IHKQuantity, metadata foundation.IDictionary) HKWorkout {
	rv := objc.Send[HKWorkout](objc.ID(getHKWorkoutClass().class), objc.Sel("workoutWithActivityType:startDate:endDate:workoutEvents:totalEnergyBurned:totalDistance:metadata:"), workoutActivityType, startDate, endDate, workoutEvents, totalEnergyBurned, totalDistance, metadata)
	return rv
}/* debug [class_init_methods/constructor]: NewHKWorkoutWithActivityTypeStartDateEndDateWorkoutEventsTotalEnergyBurnedTotalDistanceMetadata */


// Instantiates a workout using a variety of data, including the number of flights of stairs climbed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/init(activityType:start:end:workoutEvents:totalEnergyBurned:totalDistance:totalFlightsClimbed:device:metadata:)
func NewHKWorkoutWithActivityTypeStartDateEndDateWorkoutEventsTotalEnergyBurnedTotalDistanceTotalFlightsClimbedDeviceMetadata(workoutActivityType HKWorkoutActivityType, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, workoutEvents []HKWorkoutEvent, totalEnergyBurned IHKQuantity, totalDistance IHKQuantity, totalFlightsClimbed IHKQuantity, device IHKDevice, metadata foundation.IDictionary) HKWorkout {
	rv := objc.Send[HKWorkout](objc.ID(getHKWorkoutClass().class), objc.Sel("workoutWithActivityType:startDate:endDate:workoutEvents:totalEnergyBurned:totalDistance:totalFlightsClimbed:device:metadata:"), workoutActivityType, startDate, endDate, workoutEvents, totalEnergyBurned, totalDistance, totalFlightsClimbed, device, metadata)
	return rv
}/* debug [class_init_methods/constructor]: NewHKWorkoutWithActivityTypeStartDateEndDateWorkoutEventsTotalEnergyBurnedTotalDistanceTotalFlightsClimbedDeviceMetadata */


// Instantiates a workout using a variety of data, including the number of strokes while swimming.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/init(activityType:start:end:workoutEvents:totalEnergyBurned:totalDistance:totalSwimmingStrokeCount:device:metadata:)
func NewHKWorkoutWithActivityTypeStartDateEndDateWorkoutEventsTotalEnergyBurnedTotalDistanceTotalSwimmingStrokeCountDeviceMetadata(workoutActivityType HKWorkoutActivityType, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, workoutEvents []HKWorkoutEvent, totalEnergyBurned IHKQuantity, totalDistance IHKQuantity, totalSwimmingStrokeCount IHKQuantity, device IHKDevice, metadata foundation.IDictionary) HKWorkout {
	rv := objc.Send[HKWorkout](objc.ID(getHKWorkoutClass().class), objc.Sel("workoutWithActivityType:startDate:endDate:workoutEvents:totalEnergyBurned:totalDistance:totalSwimmingStrokeCount:device:metadata:"), workoutActivityType, startDate, endDate, workoutEvents, totalEnergyBurned, totalDistance, totalSwimmingStrokeCount, device, metadata)
	return rv
}/* debug [class_init_methods/constructor]: NewHKWorkoutWithActivityTypeStartDateEndDateWorkoutEventsTotalEnergyBurnedTotalDistanceTotalSwimmingStrokeCountDeviceMetadata */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKWorkout */

// Instantiates a new workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/init(activityType:start:end:)
func (hc _HKWorkoutClass) WorkoutWithActivityTypeStartDateEndDate(workoutActivityType HKWorkoutActivityType, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("workoutWithActivityType:startDate:endDate:"), workoutActivityType, startDate, endDate)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WorkoutWithActivityTypeStartDateEndDate) */


// Instantiates a new workout activity that includes the device that produced the sample data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/init(activityType:start:end:duration:totalEnergyBurned:totalDistance:device:metadata:)
func (hc _HKWorkoutClass) WorkoutWithActivityTypeStartDateEndDateDurationTotalEnergyBurnedTotalDistanceDeviceMetadata(workoutActivityType HKWorkoutActivityType, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, duration float64, totalEnergyBurned IHKQuantity, totalDistance IHKQuantity, device IHKDevice, metadata foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("workoutWithActivityType:startDate:endDate:duration:totalEnergyBurned:totalDistance:device:metadata:"), workoutActivityType, startDate, endDate, duration, totalEnergyBurned, totalDistance, device, metadata)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WorkoutWithActivityTypeStartDateEndDateDurationTotalEnergyBurnedTotalDistanceDeviceMetadata) */


// Instantiates a new workout that includes the energy burned, distance, and metadata for the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/init(activityType:start:end:duration:totalEnergyBurned:totalDistance:metadata:)
func (hc _HKWorkoutClass) WorkoutWithActivityTypeStartDateEndDateDurationTotalEnergyBurnedTotalDistanceMetadata(workoutActivityType HKWorkoutActivityType, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, duration float64, totalEnergyBurned IHKQuantity, totalDistance IHKQuantity, metadata foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("workoutWithActivityType:startDate:endDate:duration:totalEnergyBurned:totalDistance:metadata:"), workoutActivityType, startDate, endDate, duration, totalEnergyBurned, totalDistance, metadata)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WorkoutWithActivityTypeStartDateEndDateDurationTotalEnergyBurnedTotalDistanceMetadata) */


// Instantiates a workout that includes both workout events and the device that produced the sample data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/init(activityType:start:end:workoutEvents:totalEnergyBurned:totalDistance:device:metadata:)
func (hc _HKWorkoutClass) WorkoutWithActivityTypeStartDateEndDateWorkoutEventsTotalEnergyBurnedTotalDistanceDeviceMetadata(workoutActivityType HKWorkoutActivityType, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, workoutEvents []HKWorkoutEvent, totalEnergyBurned IHKQuantity, totalDistance IHKQuantity, device IHKDevice, metadata foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("workoutWithActivityType:startDate:endDate:workoutEvents:totalEnergyBurned:totalDistance:device:metadata:"), workoutActivityType, startDate, endDate, workoutEvents, totalEnergyBurned, totalDistance, device, metadata)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WorkoutWithActivityTypeStartDateEndDateWorkoutEventsTotalEnergyBurnedTotalDistanceDeviceMetadata) */


// Instantiates a new workout whose duration is calculated based on the start and end dates and the provided workout events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/init(activityType:start:end:workoutEvents:totalEnergyBurned:totalDistance:metadata:)
func (hc _HKWorkoutClass) WorkoutWithActivityTypeStartDateEndDateWorkoutEventsTotalEnergyBurnedTotalDistanceMetadata(workoutActivityType HKWorkoutActivityType, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, workoutEvents []HKWorkoutEvent, totalEnergyBurned IHKQuantity, totalDistance IHKQuantity, metadata foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("workoutWithActivityType:startDate:endDate:workoutEvents:totalEnergyBurned:totalDistance:metadata:"), workoutActivityType, startDate, endDate, workoutEvents, totalEnergyBurned, totalDistance, metadata)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WorkoutWithActivityTypeStartDateEndDateWorkoutEventsTotalEnergyBurnedTotalDistanceMetadata) */


// Instantiates a workout using a variety of data, including the number of flights of stairs climbed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/init(activityType:start:end:workoutEvents:totalEnergyBurned:totalDistance:totalFlightsClimbed:device:metadata:)
func (hc _HKWorkoutClass) WorkoutWithActivityTypeStartDateEndDateWorkoutEventsTotalEnergyBurnedTotalDistanceTotalFlightsClimbedDeviceMetadata(workoutActivityType HKWorkoutActivityType, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, workoutEvents []HKWorkoutEvent, totalEnergyBurned IHKQuantity, totalDistance IHKQuantity, totalFlightsClimbed IHKQuantity, device IHKDevice, metadata foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("workoutWithActivityType:startDate:endDate:workoutEvents:totalEnergyBurned:totalDistance:totalFlightsClimbed:device:metadata:"), workoutActivityType, startDate, endDate, workoutEvents, totalEnergyBurned, totalDistance, totalFlightsClimbed, device, metadata)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WorkoutWithActivityTypeStartDateEndDateWorkoutEventsTotalEnergyBurnedTotalDistanceTotalFlightsClimbedDeviceMetadata) */


// Instantiates a workout using a variety of data, including the number of strokes while swimming.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/init(activityType:start:end:workoutEvents:totalEnergyBurned:totalDistance:totalSwimmingStrokeCount:device:metadata:)
func (hc _HKWorkoutClass) WorkoutWithActivityTypeStartDateEndDateWorkoutEventsTotalEnergyBurnedTotalDistanceTotalSwimmingStrokeCountDeviceMetadata(workoutActivityType HKWorkoutActivityType, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, workoutEvents []HKWorkoutEvent, totalEnergyBurned IHKQuantity, totalDistance IHKQuantity, totalSwimmingStrokeCount IHKQuantity, device IHKDevice, metadata foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("workoutWithActivityType:startDate:endDate:workoutEvents:totalEnergyBurned:totalDistance:totalSwimmingStrokeCount:device:metadata:"), workoutActivityType, startDate, endDate, workoutEvents, totalEnergyBurned, totalDistance, totalSwimmingStrokeCount, device, metadata)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WorkoutWithActivityTypeStartDateEndDateWorkoutEventsTotalEnergyBurnedTotalDistanceTotalSwimmingStrokeCountDeviceMetadata) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKWorkout */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKWorkout */

// Returns the workout’s statistics for the provided quantity type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/statistics(for:)
func (h_ HKWorkout) StatisticsForType(quantityType IHKQuantityType) IHKStatistics {
	rv := objc.Send[HKStatistics](h_.ID, objc.Sel("statisticsForType:"), quantityType)
	return rv
}/* debug [instance_methods/method]: StatisticsForType */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKWorkout */

// A dictionary that contains all the statistics for the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/allStatistics
func (h_ HKWorkout) AllStatistics() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](h_.ID, objc.Sel("allStatistics"))
	return rv
}/* debug [instance_properties/getter]: allStatistics */


// The workout’s duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/duration
func (h_ HKWorkout) Duration() float64 {
	rv := objc.Send[float64](h_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// The total distance traveled during the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/totalDistance
func (h_ HKWorkout) TotalDistance() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("totalDistance"))
	return rv
}/* debug [instance_properties/getter]: totalDistance */


// The total active energy burned during the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/totalEnergyBurned
func (h_ HKWorkout) TotalEnergyBurned() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("totalEnergyBurned"))
	return rv
}/* debug [instance_properties/getter]: totalEnergyBurned */


// The total number of flights of stairs climbed during the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/totalFlightsClimbed
func (h_ HKWorkout) TotalFlightsClimbed() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("totalFlightsClimbed"))
	return rv
}/* debug [instance_properties/getter]: totalFlightsClimbed */


// The total stroke count for the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/totalSwimmingStrokeCount
func (h_ HKWorkout) TotalSwimmingStrokeCount() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("totalSwimmingStrokeCount"))
	return rv
}/* debug [instance_properties/getter]: totalSwimmingStrokeCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/workoutActivities
func (h_ HKWorkout) WorkoutActivities() []HKWorkoutActivity {
	rv := objc.Send[[]HKWorkoutActivity](h_.ID, objc.Sel("workoutActivities"))
	return rv
}/* debug [instance_properties/getter]: workoutActivities */


// The type of activity performed during the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/workoutActivityType
func (h_ HKWorkout) WorkoutActivityType() HKWorkoutActivityType {
	rv := objc.Send[HKWorkoutActivityType](h_.ID, objc.Sel("workoutActivityType"))
	return rv
}/* debug [instance_properties/getter]: workoutActivityType */


// An array of workout event objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/workoutEvents
func (h_ HKWorkout) WorkoutEvents() []HKWorkoutEvent {
	rv := objc.Send[[]HKWorkoutEvent](h_.ID, objc.Sel("workoutEvents"))
	return rv
}/* debug [instance_properties/getter]: workoutEvents */


// The key path for accessing workouts with a matching average quantity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutaveragequantity
func (h_ HKWorkout) HKPredicateKeyPathWorkoutAverageQuantity() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutAverageQuantity"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathWorkoutAverageQuantity */


// The key path for accessing the workout’s duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutduration
func (h_ HKWorkout) HKPredicateKeyPathWorkoutDuration() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutDuration"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathWorkoutDuration */


// The key path for accessing workouts with a matching maximum quantity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutmaximumquantity
func (h_ HKWorkout) HKPredicateKeyPathWorkoutMaximumQuantity() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutMaximumQuantity"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathWorkoutMaximumQuantity */


// The key path for accessing workouts with a matching minimum quantity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutminimumquantity
func (h_ HKWorkout) HKPredicateKeyPathWorkoutMinimumQuantity() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutMinimumQuantity"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathWorkoutMinimumQuantity */


// The key path for accessing workouts with a matching sum.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutsumquantity
func (h_ HKWorkout) HKPredicateKeyPathWorkoutSumQuantity() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutSumQuantity"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathWorkoutSumQuantity */


// The key path for accessing the workout’s total distance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkouttotaldistance
func (h_ HKWorkout) HKPredicateKeyPathWorkoutTotalDistance() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutTotalDistance"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathWorkoutTotalDistance */


// The key path for accessing the workout’s total energy burned.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkouttotalenergyburned
func (h_ HKWorkout) HKPredicateKeyPathWorkoutTotalEnergyBurned() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutTotalEnergyBurned"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathWorkoutTotalEnergyBurned */


// The key path for accessing the workout’s type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkouttype
func (h_ HKWorkout) HKPredicateKeyPathWorkoutType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutType"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathWorkoutType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkout/workoutplan
func (h_ HKWorkout) WorkoutPlan() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](h_.ID, objc.Sel("workoutPlan"))
	return rv
}/* debug [instance_properties/getter]: workoutPlan */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkout/workoutplan
func (h_ HKWorkout) SetWorkoutPlan(value objectivec.IObject) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setWorkoutPlan:"), value)
}/* debug [instance_properties/setter]: workoutPlan */


// A constant for sorting workouts based on their duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutsortidentifierduration
func (h_ HKWorkout) HKWorkoutSortIdentifierDuration() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKWorkoutSortIdentifierDuration"))
	return rv
}/* debug [instance_properties/getter]: HKWorkoutSortIdentifierDuration */


// A constant for sorting workouts based on their total distance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutsortidentifiertotaldistance
func (h_ HKWorkout) HKWorkoutSortIdentifierTotalDistance() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKWorkoutSortIdentifierTotalDistance"))
	return rv
}/* debug [instance_properties/getter]: HKWorkoutSortIdentifierTotalDistance */


// A constant for sorting workouts based on the total energy burned.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutsortidentifiertotalenergyburned
func (h_ HKWorkout) HKWorkoutSortIdentifierTotalEnergyBurned() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKWorkoutSortIdentifierTotalEnergyBurned"))
	return rv
}/* debug [instance_properties/getter]: HKWorkoutSortIdentifierTotalEnergyBurned */


// The workout type identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkouttypeidentifier
func (h_ HKWorkout) HKWorkoutTypeIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKWorkoutTypeIdentifier"))
	return rv
}/* debug [instance_properties/getter]: HKWorkoutTypeIdentifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKWorkout */


