// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [HKWorkout] class.
type IHKWorkout interface {
	IHKSample
	AllStatistics() unsafe.Pointer
	Duration() foundation.TimeInterval
	TotalDistance() HKQuantity
	TotalEnergyBurned() HKQuantity
	TotalFlightsClimbed() HKQuantity
	WorkoutActivities() []HKWorkoutActivity
	HKPredicateKeyPathWorkoutAverageQuantity() string
	HKPredicateKeyPathWorkoutDuration() string
	HKPredicateKeyPathWorkoutMaximumQuantity() string
	HKPredicateKeyPathWorkoutMinimumQuantity() string
	HKPredicateKeyPathWorkoutSumQuantity() string
	HKPredicateKeyPathWorkoutTotalDistance() string
	HKPredicateKeyPathWorkoutTotalEnergyBurned() string
	HKPredicateKeyPathWorkoutType() string
	TotalSwimmingStrokeCount() HKQuantity
	SetTotalSwimmingStrokeCount(value IHKQuantity)
	WorkoutActivityType() HKWorkoutActivityType
	SetWorkoutActivityType(value HKWorkoutActivityType)
	WorkoutEvents() HKWorkoutEvent
	SetWorkoutEvents(value IHKWorkoutEvent)
	WorkoutPlan() unsafe.Pointer
	SetWorkoutPlan(value unsafe.Pointer)
	HKWorkoutSortIdentifierDuration() string
	HKWorkoutSortIdentifierTotalDistance() string
	HKWorkoutSortIdentifierTotalEnergyBurned() string
	HKWorkoutTypeIdentifier() string
}

// A workout sample that stores information about a single physical activity.
//
// The class is a concrete subclass of the class; however, they behave somewhat differently than other sample types. You don’t need a specific type identifier to create the instance. All workouts use the same type identifier. You must provide an value for each workout. This value defines the type of activity performed during the workout. After saving the workout to the HealthKit store, you must associate additional samples with the workout (for example, active energy burned or distance samples). These samples provide fine-grained details. Use the method to associate them with the workout. The workout records a summary of information about a single physical activity (for example, the duration, total distance, and total energy burned). It also acts as a container for other objects. You can associate any number of samples with a workout, adding details over the course of the workout. For example, you may want to break a single run into a number of shorter intervals, and then add samples to track the user’s heart rate, energy burned, distance traveled, and steps taken for each interval. For more information, see . HealthKit supports a wide range of activity types. For a complete list, see . Workouts are mostly immutable. You set their properties when you instantiate the workout, and they can’t change. However, you can continue to add samples to the workouts.
//
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

// Alloc allocates a new instance without initialization.
func (hc _HKWorkoutClass) Alloc() HKWorkout {
	rv := objc.Send[HKWorkout](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Instantiates a workout using a variety of data, including the number of strokes while swimming.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/init(activityType:start:end:workoutEvents:totalEnergyBurned:totalDistance:totalSwimmingStrokeCount:device:metadata:)
func NewHKWorkoutWithActivityTypeStartDateEndDateWorkoutEventsTotalEnergyBurnedTotalDistanceTotalSwimmingStrokeCountDeviceMetadata(workoutActivityType HKWorkoutActivityType, startDate foundation.IDate, endDate foundation.IDate, workoutEvents []HKWorkoutEvent, totalEnergyBurned IHKQuantity, totalDistance IHKQuantity, totalSwimmingStrokeCount IHKQuantity, device IHKDevice, metadata unsafe.Pointer) HKWorkout {
	rv := objc.Send[HKWorkout](objc.ID(getHKWorkoutClass().class), objc.Sel("workoutWithActivityType:startDate:endDate:workoutEvents:totalEnergyBurned:totalDistance:totalSwimmingStrokeCount:device:metadata:"), workoutActivityType, startDate, endDate, workoutEvents, totalEnergyBurned, totalDistance, totalSwimmingStrokeCount, device, metadata)
	return rv
}


// Instantiates a workout using a variety of data, including the number of strokes while swimming.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/init(activityType:start:end:workoutEvents:totalEnergyBurned:totalDistance:totalSwimmingStrokeCount:device:metadata:)
func (hc _HKWorkoutClass) WorkoutWithActivityTypeStartDateEndDateWorkoutEventsTotalEnergyBurnedTotalDistanceTotalSwimmingStrokeCountDeviceMetadata(workoutActivityType HKWorkoutActivityType, startDate foundation.IDate, endDate foundation.IDate, workoutEvents []HKWorkoutEvent, totalEnergyBurned IHKQuantity, totalDistance IHKQuantity, totalSwimmingStrokeCount IHKQuantity, device IHKDevice, metadata unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("workoutWithActivityType:startDate:endDate:workoutEvents:totalEnergyBurned:totalDistance:totalSwimmingStrokeCount:device:metadata:"), workoutActivityType, startDate, endDate, workoutEvents, totalEnergyBurned, totalDistance, totalSwimmingStrokeCount, device, metadata)
	return rv
}

// A dictionary that contains all the statistics for the workout.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/allStatistics
func (h_ HKWorkout) AllStatistics() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("allStatistics"))
	return rv
}

// The workout’s duration.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/duration
func (h_ HKWorkout) Duration() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](h_.ID, objc.Sel("duration"))
	return rv
}

// The total distance traveled during the workout.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/totalDistance
func (h_ HKWorkout) TotalDistance() HKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("totalDistance"))
	return rv
}

// The total active energy burned during the workout.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/totalEnergyBurned
func (h_ HKWorkout) TotalEnergyBurned() HKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("totalEnergyBurned"))
	return rv
}

// The total number of flights of stairs climbed during the workout.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/totalFlightsClimbed
func (h_ HKWorkout) TotalFlightsClimbed() HKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("totalFlightsClimbed"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/workoutActivities
func (h_ HKWorkout) WorkoutActivities() []HKWorkoutActivity {
	rv := objc.Send[[]HKWorkoutActivity](h_.ID, objc.Sel("workoutActivities"))
	return rv
}

// The key path for accessing workouts with a matching average quantity.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutaveragequantity
func (h_ HKWorkout) HKPredicateKeyPathWorkoutAverageQuantity() string {
	rv := objc.Send[string](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutAverageQuantity"))
	return rv
}

// The key path for accessing the workout’s duration.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutduration
func (h_ HKWorkout) HKPredicateKeyPathWorkoutDuration() string {
	rv := objc.Send[string](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutDuration"))
	return rv
}

// The key path for accessing workouts with a matching maximum quantity.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutmaximumquantity
func (h_ HKWorkout) HKPredicateKeyPathWorkoutMaximumQuantity() string {
	rv := objc.Send[string](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutMaximumQuantity"))
	return rv
}

// The key path for accessing workouts with a matching minimum quantity.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutminimumquantity
func (h_ HKWorkout) HKPredicateKeyPathWorkoutMinimumQuantity() string {
	rv := objc.Send[string](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutMinimumQuantity"))
	return rv
}

// The key path for accessing workouts with a matching sum.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutsumquantity
func (h_ HKWorkout) HKPredicateKeyPathWorkoutSumQuantity() string {
	rv := objc.Send[string](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutSumQuantity"))
	return rv
}

// The key path for accessing the workout’s total distance.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkouttotaldistance
func (h_ HKWorkout) HKPredicateKeyPathWorkoutTotalDistance() string {
	rv := objc.Send[string](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutTotalDistance"))
	return rv
}

// The key path for accessing the workout’s total energy burned.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkouttotalenergyburned
func (h_ HKWorkout) HKPredicateKeyPathWorkoutTotalEnergyBurned() string {
	rv := objc.Send[string](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutTotalEnergyBurned"))
	return rv
}

// The key path for accessing the workout’s type.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkouttype
func (h_ HKWorkout) HKPredicateKeyPathWorkoutType() string {
	rv := objc.Send[string](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutType"))
	return rv
}

// The total stroke count for the workout.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkout/totalswimmingstrokecount
func (h_ HKWorkout) TotalSwimmingStrokeCount() HKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("totalSwimmingStrokeCount"))
	return rv
}


// SetTotalSwimmingStrokeCount sets the value of the totalSwimmingStrokeCount property.
// The total stroke count for the workout.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkout/totalswimmingstrokecount
func (h_ HKWorkout) SetTotalSwimmingStrokeCount(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setTotalSwimmingStrokeCount:"), value)
}

// The type of activity performed during the workout.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkout/workoutactivitytype
func (h_ HKWorkout) WorkoutActivityType() HKWorkoutActivityType {
	rv := objc.Send[HKWorkoutActivityType](h_.ID, objc.Sel("workoutActivityType"))
	return rv
}


// SetWorkoutActivityType sets the value of the workoutActivityType property.
// The type of activity performed during the workout.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkout/workoutactivitytype
func (h_ HKWorkout) SetWorkoutActivityType(value HKWorkoutActivityType) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setWorkoutActivityType:"), value)
}

// An array of workout event objects.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkout/workoutevents
func (h_ HKWorkout) WorkoutEvents() HKWorkoutEvent {
	rv := objc.Send[HKWorkoutEvent](h_.ID, objc.Sel("workoutEvents"))
	return rv
}


// SetWorkoutEvents sets the value of the workoutEvents property.
// An array of workout event objects.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkout/workoutevents
func (h_ HKWorkout) SetWorkoutEvents(value IHKWorkoutEvent) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setWorkoutEvents:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkout/workoutplan
func (h_ HKWorkout) WorkoutPlan() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("workoutPlan"))
	return rv
}


// SetWorkoutPlan sets the value of the workoutPlan property.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkout/workoutplan
func (h_ HKWorkout) SetWorkoutPlan(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setWorkoutPlan:"), value)
}

// A constant for sorting workouts based on their duration.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutsortidentifierduration
func (h_ HKWorkout) HKWorkoutSortIdentifierDuration() string {
	rv := objc.Send[string](h_.ID, objc.Sel("HKWorkoutSortIdentifierDuration"))
	return rv
}

// A constant for sorting workouts based on their total distance.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutsortidentifiertotaldistance
func (h_ HKWorkout) HKWorkoutSortIdentifierTotalDistance() string {
	rv := objc.Send[string](h_.ID, objc.Sel("HKWorkoutSortIdentifierTotalDistance"))
	return rv
}

// A constant for sorting workouts based on the total energy burned.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutsortidentifiertotalenergyburned
func (h_ HKWorkout) HKWorkoutSortIdentifierTotalEnergyBurned() string {
	rv := objc.Send[string](h_.ID, objc.Sel("HKWorkoutSortIdentifierTotalEnergyBurned"))
	return rv
}

// The workout type identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkouttypeidentifier
func (h_ HKWorkout) HKWorkoutTypeIdentifier() string {
	rv := objc.Send[string](h_.ID, objc.Sel("HKWorkoutTypeIdentifier"))
	return rv
}


