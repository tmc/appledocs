// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [HKWorkoutActivity] class.
type IHKWorkoutActivity interface {
	objectivec.IObject
	// properties:
	HKPredicateKeyPathWorkoutActivity() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathWorkoutActivityAverageQuantity() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathWorkoutActivityDuration() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathWorkoutActivityEndDate() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathWorkoutActivityMaximumQuantity() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathWorkoutActivityMinimumQuantity() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathWorkoutActivityStartDate() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathWorkoutActivitySumQuantity() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathWorkoutActivityType() objc.IObject /* cross-framework: NSString */
	AllStatistics() IHKStatistics
	SetAllStatistics(value IHKStatistics)
	Duration() float64
	SetDuration(value float64)
	EndDate() objc.IObject /* cross-framework: Date */
	SetEndDate(value objc.IObject /* cross-framework: Date */)
	Metadata() objc.IObject /* cross-framework: NSString */
	SetMetadata(value objc.IObject /* cross-framework: NSString */)
	StartDate() objc.IObject /* cross-framework: Date */
	SetStartDate(value objc.IObject /* cross-framework: Date */)
	Uuid() objc.IObject /* cross-framework: UUID */
	SetUuid(value objc.IObject /* cross-framework: UUID */)
	WorkoutConfiguration() IHKWorkoutConfiguration
	SetWorkoutConfiguration(value IHKWorkoutConfiguration)
	WorkoutEvents() IHKWorkoutEvent
	SetWorkoutEvents(value IHKWorkoutEvent)
	HKWorkoutTypeIdentifier() objc.IObject /* cross-framework: NSString */
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (hc _HKWorkoutActivityClass) Alloc() HKWorkoutActivity {
	rv := objc.Send[HKWorkoutActivity](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The key path for accessing a specific workout activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutactivity
func (h_ HKWorkoutActivity) HKPredicateKeyPathWorkoutActivity() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutActivity"))
	return rv
}


// The key path for accessing activities with a matching average quantity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutactivityaveragequantity
func (h_ HKWorkoutActivity) HKPredicateKeyPathWorkoutActivityAverageQuantity() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutActivityAverageQuantity"))
	return rv
}


// The key path for accessing activities with a matching duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutactivityduration
func (h_ HKWorkoutActivity) HKPredicateKeyPathWorkoutActivityDuration() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutActivityDuration"))
	return rv
}


// The key path for accessing activities with a matching end date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutactivityenddate
func (h_ HKWorkoutActivity) HKPredicateKeyPathWorkoutActivityEndDate() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutActivityEndDate"))
	return rv
}


// The key path for accessing activities with a matching maximum quantity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutactivitymaximumquantity
func (h_ HKWorkoutActivity) HKPredicateKeyPathWorkoutActivityMaximumQuantity() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutActivityMaximumQuantity"))
	return rv
}


// The key path for accessing activities with a matching minimum quantity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutactivityminimumquantity
func (h_ HKWorkoutActivity) HKPredicateKeyPathWorkoutActivityMinimumQuantity() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutActivityMinimumQuantity"))
	return rv
}


// The key path for accessing activities with a matching start date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutactivitystartdate
func (h_ HKWorkoutActivity) HKPredicateKeyPathWorkoutActivityStartDate() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutActivityStartDate"))
	return rv
}


// The key path for accessing activities with a matching sum.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutactivitysumquantity
func (h_ HKWorkoutActivity) HKPredicateKeyPathWorkoutActivitySumQuantity() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutActivitySumQuantity"))
	return rv
}


// The key path for accessing activities that match a workout activity type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutactivitytype
func (h_ HKWorkoutActivity) HKPredicateKeyPathWorkoutActivityType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutActivityType"))
	return rv
}


// A dictionary that contains all the statistics for the activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutactivity/allstatistics
func (h_ HKWorkoutActivity) AllStatistics() IHKStatistics {
	rv := objc.Send[HKStatistics](h_.ID, objc.Sel("allStatistics"))
	return rv
}


// A dictionary that contains all the statistics for the activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutactivity/allstatistics
func (h_ HKWorkoutActivity) SetAllStatistics(value IHKStatistics) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAllStatistics:"), value)
}


// The activity’s duration, measured in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutactivity/duration
func (h_ HKWorkoutActivity) Duration() float64 {
	rv := objc.Send[float64](h_.ID, objc.Sel("duration"))
	return rv
}


// The activity’s duration, measured in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutactivity/duration
func (h_ HKWorkoutActivity) SetDuration(value float64) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setDuration:"), value)
}


// The activity’s end date and time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutactivity/enddate
func (h_ HKWorkoutActivity) EndDate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](h_.ID, objc.Sel("endDate"))
	return rv
}


// The activity’s end date and time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutactivity/enddate
func (h_ HKWorkoutActivity) SetEndDate(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setEndDate:"), value)
}


// Metadata that describes the activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutactivity/metadata
func (h_ HKWorkoutActivity) Metadata() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("metadata"))
	return rv
}


// Metadata that describes the activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutactivity/metadata
func (h_ HKWorkoutActivity) SetMetadata(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setMetadata:"), value)
}


// The activitiy’s start date and time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutactivity/startdate
func (h_ HKWorkoutActivity) StartDate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](h_.ID, objc.Sel("startDate"))
	return rv
}


// The activitiy’s start date and time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutactivity/startdate
func (h_ HKWorkoutActivity) SetStartDate(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setStartDate:"), value)
}


// The activity’s universally unique identifier (UUID).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutactivity/uuid
func (h_ HKWorkoutActivity) Uuid() objc.IObject /* cross-framework: UUID */ {
	rv := objc.Send[foundation.UUID](h_.ID, objc.Sel("uuid"))
	return rv
}


// The activity’s universally unique identifier (UUID).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutactivity/uuid
func (h_ HKWorkoutActivity) SetUuid(value objc.IObject /* cross-framework: UUID */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setUuid:"), value)
}


// The configuration information for this part of the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutactivity/workoutconfiguration
func (h_ HKWorkoutActivity) WorkoutConfiguration() IHKWorkoutConfiguration {
	rv := objc.Send[HKWorkoutConfiguration](h_.ID, objc.Sel("workoutConfiguration"))
	return rv
}


// The configuration information for this part of the workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutactivity/workoutconfiguration
func (h_ HKWorkoutActivity) SetWorkoutConfiguration(value IHKWorkoutConfiguration) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setWorkoutConfiguration:"), value)
}


// An array of events associated with the containing workout and occurring during the activity’s duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutactivity/workoutevents
func (h_ HKWorkoutActivity) WorkoutEvents() IHKWorkoutEvent {
	rv := objc.Send[HKWorkoutEvent](h_.ID, objc.Sel("workoutEvents"))
	return rv
}


// An array of events associated with the containing workout and occurring during the activity’s duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutactivity/workoutevents
func (h_ HKWorkoutActivity) SetWorkoutEvents(value IHKWorkoutEvent) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setWorkoutEvents:"), value)
}


// The workout type identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkouttypeidentifier
func (h_ HKWorkoutActivity) HKWorkoutTypeIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKWorkoutTypeIdentifier"))
	return rv
}



