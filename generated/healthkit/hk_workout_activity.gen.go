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
}

// An object that describes an activity within a longer workout.
//
// Workout activity objects partition a workout into a set of separate activities. For example, you can use workout activities to record the swim, bike, and running portions of a multisport event, like a triathlon, or to represent the active and rest periods during interval training. All instance have at least one, associated . If you don’t explicitly set workout activities, HealthKit assigns a workout activity that matches the object’s activity type. For more information, see .
//
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


// The activity’s duration, measured in seconds.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivity/duration
func (h_ HKWorkoutActivity) Duration() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](h_.ID, objc.Sel("duration"))
	return rv
}

// The key path for accessing a specific workout activity.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutactivity
func (h_ HKWorkoutActivity) HKPredicateKeyPathWorkoutActivity() appkit.string {
	rv := objc.Send[appkit.string](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutActivity"))
	return rv
}

// The key path for accessing activities with a matching average quantity.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutactivityaveragequantity
func (h_ HKWorkoutActivity) HKPredicateKeyPathWorkoutActivityAverageQuantity() appkit.string {
	rv := objc.Send[appkit.string](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutActivityAverageQuantity"))
	return rv
}

// The key path for accessing activities with a matching duration.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutactivityduration
func (h_ HKWorkoutActivity) HKPredicateKeyPathWorkoutActivityDuration() appkit.string {
	rv := objc.Send[appkit.string](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutActivityDuration"))
	return rv
}

// The key path for accessing activities with a matching end date.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutactivityenddate
func (h_ HKWorkoutActivity) HKPredicateKeyPathWorkoutActivityEndDate() appkit.string {
	rv := objc.Send[appkit.string](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutActivityEndDate"))
	return rv
}

// The key path for accessing activities with a matching maximum quantity.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutactivitymaximumquantity
func (h_ HKWorkoutActivity) HKPredicateKeyPathWorkoutActivityMaximumQuantity() appkit.string {
	rv := objc.Send[appkit.string](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutActivityMaximumQuantity"))
	return rv
}

// The key path for accessing activities with a matching minimum quantity.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutactivityminimumquantity
func (h_ HKWorkoutActivity) HKPredicateKeyPathWorkoutActivityMinimumQuantity() appkit.string {
	rv := objc.Send[appkit.string](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutActivityMinimumQuantity"))
	return rv
}

// The key path for accessing activities with a matching start date.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutactivitystartdate
func (h_ HKWorkoutActivity) HKPredicateKeyPathWorkoutActivityStartDate() appkit.string {
	rv := objc.Send[appkit.string](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutActivityStartDate"))
	return rv
}

// The key path for accessing activities with a matching sum.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutactivitysumquantity
func (h_ HKWorkoutActivity) HKPredicateKeyPathWorkoutActivitySumQuantity() appkit.string {
	rv := objc.Send[appkit.string](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutActivitySumQuantity"))
	return rv
}

// The key path for accessing activities that match a workout activity type.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathworkoutactivitytype
func (h_ HKWorkoutActivity) HKPredicateKeyPathWorkoutActivityType() appkit.string {
	rv := objc.Send[appkit.string](h_.ID, objc.Sel("HKPredicateKeyPathWorkoutActivityType"))
	return rv
}

// A dictionary that contains all the statistics for the activity.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutactivity/allstatistics
func (h_ HKWorkoutActivity) AllStatistics() HKStatistics {
	rv := objc.Send[HKStatistics](h_.ID, objc.Sel("allStatistics"))
	return rv
}


// SetAllStatistics sets the value of the allStatistics property.
// A dictionary that contains all the statistics for the activity.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutactivity/allstatistics
func (h_ HKWorkoutActivity) SetAllStatistics(value IHKStatistics) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAllStatistics:"), value)
}

// The activity’s end date and time.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutactivity/enddate
func (h_ HKWorkoutActivity) EndDate() foundation.Date {
	rv := objc.Send[foundation.Date](h_.ID, objc.Sel("endDate"))
	return rv
}


// SetEndDate sets the value of the endDate property.
// The activity’s end date and time.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutactivity/enddate
func (h_ HKWorkoutActivity) SetEndDate(value foundation.IDate) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setEndDate:"), value)
}

// Metadata that describes the activity.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutactivity/metadata
func (h_ HKWorkoutActivity) Metadata() appkit.string {
	rv := objc.Send[appkit.string](h_.ID, objc.Sel("metadata"))
	return rv
}


// SetMetadata sets the value of the metadata property.
// Metadata that describes the activity.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutactivity/metadata
func (h_ HKWorkoutActivity) SetMetadata(value appkit.string) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setMetadata:"), value)
}

// The activitiy’s start date and time.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutactivity/startdate
func (h_ HKWorkoutActivity) StartDate() foundation.Date {
	rv := objc.Send[foundation.Date](h_.ID, objc.Sel("startDate"))
	return rv
}


// SetStartDate sets the value of the startDate property.
// The activitiy’s start date and time.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutactivity/startdate
func (h_ HKWorkoutActivity) SetStartDate(value foundation.IDate) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setStartDate:"), value)
}

// The activity’s universally unique identifier (UUID).
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutactivity/uuid
func (h_ HKWorkoutActivity) Uuid() foundation.UUID {
	rv := objc.Send[foundation.UUID](h_.ID, objc.Sel("uuid"))
	return rv
}


// SetUuid sets the value of the uuid property.
// The activity’s universally unique identifier (UUID).

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutactivity/uuid
func (h_ HKWorkoutActivity) SetUuid(value foundation.IUUID) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setUuid:"), value)
}

// The configuration information for this part of the workout.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutactivity/workoutconfiguration
func (h_ HKWorkoutActivity) WorkoutConfiguration() HKWorkoutConfiguration {
	rv := objc.Send[HKWorkoutConfiguration](h_.ID, objc.Sel("workoutConfiguration"))
	return rv
}


// SetWorkoutConfiguration sets the value of the workoutConfiguration property.
// The configuration information for this part of the workout.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutactivity/workoutconfiguration
func (h_ HKWorkoutActivity) SetWorkoutConfiguration(value IHKWorkoutConfiguration) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setWorkoutConfiguration:"), value)
}

// An array of events associated with the containing workout and occurring during the activity’s duration.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutactivity/workoutevents
func (h_ HKWorkoutActivity) WorkoutEvents() HKWorkoutEvent {
	rv := objc.Send[HKWorkoutEvent](h_.ID, objc.Sel("workoutEvents"))
	return rv
}


// SetWorkoutEvents sets the value of the workoutEvents property.
// An array of events associated with the containing workout and occurring during the activity’s duration.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutactivity/workoutevents
func (h_ HKWorkoutActivity) SetWorkoutEvents(value IHKWorkoutEvent) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setWorkoutEvents:"), value)
}

// The workout type identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkouttypeidentifier
func (h_ HKWorkoutActivity) HKWorkoutTypeIdentifier() appkit.string {
	rv := objc.Send[appkit.string](h_.ID, objc.Sel("HKWorkoutTypeIdentifier"))
	return rv
}



