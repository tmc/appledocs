// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKWorkoutActivity */


/* debug [class_header]: Header for HKWorkoutActivity */
// The class instance for the [HKWorkoutActivity] class.
var (
	HKWorkoutActivityClass     _HKWorkoutActivityClass
	HKWorkoutActivityClassOnce sync.Once
)

func getHKWorkoutActivityClass() _HKWorkoutActivityClass {
	HKWorkoutActivityClassOnce.Do(func() {
		HKWorkoutActivityClass = _HKWorkoutActivityClass{objc.GetClass("HKWorkoutActivity")}
	})
	return HKWorkoutActivityClass
}

type _HKWorkoutActivityClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKWorkoutActivity */
// An interface definition for the [HKWorkoutActivity] class.
type IHKWorkoutActivity interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKWorkoutActivity */
	// properties:
	AllStatistics() foundation.IDictionary
	Duration() float64
	EndDate() objc.IObject /* cross-framework: NSDate */
	Metadata() foundation.IDictionary
	StartDate() objc.IObject /* cross-framework: NSDate */
	UUID() foundation.UUID
	WorkoutConfiguration() IHKWorkoutConfiguration
	WorkoutEvents() []HKWorkoutEvent
	HKPredicateKeyPathWorkoutActivity() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathWorkoutActivityAverageQuantity() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathWorkoutActivityDuration() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathWorkoutActivityEndDate() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathWorkoutActivityMaximumQuantity() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathWorkoutActivityMinimumQuantity() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathWorkoutActivityStartDate() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathWorkoutActivitySumQuantity() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathWorkoutActivityType() objc.IObject /* cross-framework: NSString */
	HKWorkoutTypeIdentifier() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKWorkoutActivity */
	// methods:
	StatisticsForType(quantityType IHKQuantityType) IHKStatistics
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKWorkoutActivity */
// Alloc allocates a new instance without initialization.
func (hc _HKWorkoutActivityClass) Alloc() HKWorkoutActivity {
	rv := objc.Send[HKWorkoutActivity](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKWorkoutActivityClass) New() HKWorkoutActivity {
	rv := objc.Send[HKWorkoutActivity](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKWorkoutActivity) Init() HKWorkoutActivity {
	rv := objc.Send[HKWorkoutActivity](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKWorkoutActivity) Autorelease() HKWorkoutActivity {
	rv := objc.Send[HKWorkoutActivity](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKWorkoutActivity creates a new HKWorkoutActivity instance.
func NewHKWorkoutActivity() HKWorkoutActivity {
	return getHKWorkoutActivityClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKWorkoutActivity */
// An object that describes an activity within a longer workout.
//
// Workout activity objects partition a workout into a set of separate activities. For example, you can use workout activities to record the swim, bike, and running portions of a multisport event, like a triathlon, or to represent the active and rest periods during interval training. All instance have at least one, associated . If you don’t explicitly set workout activities, HealthKit assigns a workout activity that matches the object’s activity type. For more information, see .


// An object that describes an activity within a longer workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivity
type HKWorkoutActivity struct {
	objectivec.Object
}

// HKWorkoutActivityFrom constructs a [HKWorkoutActivity] from an unsafe.Pointer.
//
// An object that describes an activity within a longer workout.
func HKWorkoutActivityFrom(ptr unsafe.Pointer) HKWorkoutActivity {
	return HKWorkoutActivity{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKWorkoutActivity */

// Creates a workout activity using the provided configuration, start date, end date, and metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivity/init(workoutConfiguration:start:end:metadata:)
func NewHKWorkoutActivityWithWorkoutConfigurationStartDateEndDateMetadata(workoutConfiguration IHKWorkoutConfiguration, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, metadata foundation.IDictionary) HKWorkoutActivity {
	instance := getHKWorkoutActivityClass().Alloc()
	rv := objc.Send[HKWorkoutActivity](instance.ID, objc.Sel("initWithWorkoutConfiguration:startDate:endDate:metadata:"), workoutConfiguration, startDate, endDate, metadata)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKWorkoutActivityWithWorkoutConfigurationStartDateEndDateMetadata */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKWorkoutActivity */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKWorkoutActivity */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKWorkoutActivity */

// Returns the activity’s statistics for the provided quantity type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivity/statistics(for:)
func (h_ HKWorkoutActivity) StatisticsForType(quantityType IHKQuantityType) IHKStatistics {
	rv := objc.Send[HKStatistics](h_.ID, objc.Sel("statisticsForType:"), quantityType)
	return rv
}/* debug [instance_methods/method]: StatisticsForType */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKWorkoutActivity */

// A dictionary that contains all the statistics for the activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivity/allStatistics
func (h_ HKWorkoutActivity) AllStatistics() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](h_.ID, objc.Sel("allStatistics"))
	return rv
}/* debug [instance_properties/getter]: allStatistics */


// The activity’s duration, measured in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivity/duration
func (h_ HKWorkoutActivity) Duration() float64 {
	rv := objc.Send[float64](h_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// The activity’s end date and time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivity/endDate
func (h_ HKWorkoutActivity) EndDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](h_.ID, objc.Sel("endDate"))
	return rv
}/* debug [instance_properties/getter]: endDate */


// Metadata that describes the activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivity/metadata
func (h_ HKWorkoutActivity) Metadata() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](h_.ID, objc.Sel("metadata"))
	return rv
}/* debug [instance_properties/getter]: metadata */


// The activitiy’s start date and time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivity/startDate
func (h_ HKWorkoutActivity) StartDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](h_.ID, objc.Sel("startDate"))
	return rv
}/* debug [instance_properties/getter]: startDate */


// The activity’s universally unique identifier (UUID).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivity/uuid
func (h_ HKWorkoutActivity) UUID() foundation.UUID {
	rv := objc.Send[foundation.UUID](h_.ID, objc.Sel("UUID"))
	return rv
}/* debug [instance_properties/getter]: UUID */


// The configuration information for this part of the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivity/workoutConfiguration
func (h_ HKWorkoutActivity) WorkoutConfiguration() IHKWorkoutConfiguration {
	rv := objc.Send[HKWorkoutConfiguration](h_.ID, objc.Sel("workoutConfiguration"))
	return rv
}/* debug [instance_properties/getter]: workoutConfiguration */


// An array of events associated with the containing workout and occurring during the activity’s duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivity/workoutEvents
func (h_ HKWorkoutActivity) WorkoutEvents() []HKWorkoutEvent {
	rv := objc.Send[[]HKWorkoutEvent](h_.ID, objc.Sel("workoutEvents"))
	return rv
}/* debug [instance_properties/getter]: workoutEvents */


// The key path for accessing a specific workout activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutactivity
func (h_ HKWorkoutActivity) HKPredicateKeyPathWorkoutActivity() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutActivity"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathWorkoutActivity */


// The key path for accessing activities with a matching average quantity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutactivityaveragequantity
func (h_ HKWorkoutActivity) HKPredicateKeyPathWorkoutActivityAverageQuantity() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutActivityAverageQuantity"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathWorkoutActivityAverageQuantity */


// The key path for accessing activities with a matching duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutactivityduration
func (h_ HKWorkoutActivity) HKPredicateKeyPathWorkoutActivityDuration() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutActivityDuration"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathWorkoutActivityDuration */


// The key path for accessing activities with a matching end date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutactivityenddate
func (h_ HKWorkoutActivity) HKPredicateKeyPathWorkoutActivityEndDate() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutActivityEndDate"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathWorkoutActivityEndDate */


// The key path for accessing activities with a matching maximum quantity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutactivitymaximumquantity
func (h_ HKWorkoutActivity) HKPredicateKeyPathWorkoutActivityMaximumQuantity() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutActivityMaximumQuantity"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathWorkoutActivityMaximumQuantity */


// The key path for accessing activities with a matching minimum quantity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutactivityminimumquantity
func (h_ HKWorkoutActivity) HKPredicateKeyPathWorkoutActivityMinimumQuantity() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutActivityMinimumQuantity"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathWorkoutActivityMinimumQuantity */


// The key path for accessing activities with a matching start date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutactivitystartdate
func (h_ HKWorkoutActivity) HKPredicateKeyPathWorkoutActivityStartDate() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutActivityStartDate"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathWorkoutActivityStartDate */


// The key path for accessing activities with a matching sum.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutactivitysumquantity
func (h_ HKWorkoutActivity) HKPredicateKeyPathWorkoutActivitySumQuantity() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutActivitySumQuantity"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathWorkoutActivitySumQuantity */


// The key path for accessing activities that match a workout activity type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutactivitytype
func (h_ HKWorkoutActivity) HKPredicateKeyPathWorkoutActivityType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutActivityType"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathWorkoutActivityType */


// The workout type identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkouttypeidentifier
func (h_ HKWorkoutActivity) HKWorkoutTypeIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKWorkoutTypeIdentifier"))
	return rv
}/* debug [instance_properties/getter]: HKWorkoutTypeIdentifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKWorkoutActivity */


