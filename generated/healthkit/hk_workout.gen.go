// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
func NewHKWorkoutWithActivityTypeStartDateEndDateWorkoutEventsTotalEnergyBurnedTotalDistanceTotalSwimmingStrokeCountDeviceMetadata(workoutActivityType unsafe.Pointer, startDate unsafe.Pointer, endDate unsafe.Pointer, workoutEvents unsafe.Pointer, totalEnergyBurned unsafe.Pointer, totalDistance unsafe.Pointer, totalSwimmingStrokeCount unsafe.Pointer, device unsafe.Pointer, metadata unsafe.Pointer) HKWorkout {
	rv := objc.Send[HKWorkout](objc.ID(getHKWorkoutClass().class), objc.Sel("workoutWithActivityType:startDate:endDate:workoutEvents:totalEnergyBurned:totalDistance:totalSwimmingStrokeCount:device:metadata:"), workoutActivityType, startDate, endDate, workoutEvents, totalEnergyBurned, totalDistance, totalSwimmingStrokeCount, device, metadata)
	return rv
}


// Instantiates a workout using a variety of data, including the number of strokes while swimming.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/init(activityType:start:end:workoutEvents:totalEnergyBurned:totalDistance:totalSwimmingStrokeCount:device:metadata:)
func (hc _HKWorkoutClass) WorkoutWithActivityTypeStartDateEndDateWorkoutEventsTotalEnergyBurnedTotalDistanceTotalSwimmingStrokeCountDeviceMetadata(workoutActivityType unsafe.Pointer, startDate unsafe.Pointer, endDate unsafe.Pointer, workoutEvents unsafe.Pointer, totalEnergyBurned unsafe.Pointer, totalDistance unsafe.Pointer, totalSwimmingStrokeCount unsafe.Pointer, device unsafe.Pointer, metadata unsafe.Pointer) unsafe.Pointer {
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
func (h_ HKWorkout) Duration() TimeInterval {
	rv := objc.Send[TimeInterval](h_.ID, objc.Sel("duration"))
	return rv
}

// The total distance traveled during the workout.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/totalDistance
func (h_ HKWorkout) TotalDistance() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("totalDistance"))
	return rv
}

// The total active energy burned during the workout.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/totalEnergyBurned
func (h_ HKWorkout) TotalEnergyBurned() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("totalEnergyBurned"))
	return rv
}

// The total number of flights of stairs climbed during the workout.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/totalFlightsClimbed
func (h_ HKWorkout) TotalFlightsClimbed() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("totalFlightsClimbed"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkout/workoutActivities
func (h_ HKWorkout) WorkoutActivities() []HKWorkoutActivity {
	rv := objc.Send[[]HKWorkoutActivity](h_.ID, objc.Sel("workoutActivities"))
	return rv
}


