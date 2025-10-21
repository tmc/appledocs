// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [HKActivitySummary] class.
var (
	HKActivitySummaryClass     _HKActivitySummaryClass
	HKActivitySummaryClassOnce sync.Once
)

func getHKActivitySummaryClass() _HKActivitySummaryClass {
	HKActivitySummaryClassOnce.Do(func() {
		HKActivitySummaryClass = _HKActivitySummaryClass{objc.GetClass("HKActivitySummary")}
	})
	return HKActivitySummaryClass
}

type _HKActivitySummaryClass struct {
	class objc.Class
}

// An interface definition for the [HKActivitySummary] class.
type IHKActivitySummary interface {
	objectivec.IObject
	DateComponentsForCalendar(calendar unsafe.Pointer) unsafe.Pointer
}

// An object that contains the move, exercise, and stand data for a given day.
//
// You can read objects from the HealthKit store using an object. Unlike the subclasses, instances are mutable, but changes made to the object’s properties have no affect on the values in the HealthKit store. You can instantiate your own objects (if needed), but you can’t save objects to the store. You can display an active summary in iOS using the class or in watchOS using the class.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary
type HKActivitySummary struct {
	objectivec.Object
}

// HKActivitySummaryFrom constructs a [HKActivitySummary] from an unsafe.Pointer.
//
// An object that contains the move, exercise, and stand data for a given day.
func HKActivitySummaryFrom(ptr unsafe.Pointer) HKActivitySummary {
	return HKActivitySummary{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKActivitySummaryClass) Alloc() HKActivitySummary {
	rv := objc.Send[HKActivitySummary](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKActivitySummaryClass) New() HKActivitySummary {
	rv := objc.Send[HKActivitySummary](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKActivitySummary) Init() HKActivitySummary {
	rv := objc.Send[HKActivitySummary](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKActivitySummary) Autorelease() HKActivitySummary {
	rv := objc.Send[HKActivitySummary](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKActivitySummary creates a new HKActivitySummary instance.
func NewHKActivitySummary() HKActivitySummary {
	return getHKActivitySummaryClass().New()
}


// Date components that uniquely identify the day represented by the summary object.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/dateComponents(for:)
func (h_ HKActivitySummary) DateComponentsForCalendar(calendar unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("dateComponentsForCalendar:"), calendar)
	return rv
}

// The amount of active energy the user burned during the specified day.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/activeEnergyBurned
func (h_ HKActivitySummary) ActiveEnergyBurned() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("activeEnergyBurned"))
	return rv
}


// SetActiveEnergyBurned sets the value of the activeEnergyBurned property.
// The amount of active energy the user burned during the specified day.

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/activeEnergyBurned
func (h_ HKActivitySummary) SetActiveEnergyBurned(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setActiveEnergyBurned:"), value)
}
// The user’s daily goal for active energy burned.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/activeEnergyBurnedGoal
func (h_ HKActivitySummary) ActiveEnergyBurnedGoal() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("activeEnergyBurnedGoal"))
	return rv
}


// SetActiveEnergyBurnedGoal sets the value of the activeEnergyBurnedGoal property.
// The user’s daily goal for active energy burned.

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/activeEnergyBurnedGoal
func (h_ HKActivitySummary) SetActiveEnergyBurnedGoal(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setActiveEnergyBurnedGoal:"), value)
}
// The move mode that they system used for this activity summary.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/activityMoveMode
func (h_ HKActivitySummary) ActivityMoveMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("activityMoveMode"))
	return rv
}


// SetActivityMoveMode sets the value of the activityMoveMode property.
// The move mode that they system used for this activity summary.

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/activityMoveMode
func (h_ HKActivitySummary) SetActivityMoveMode(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setActivityMoveMode:"), value)
}
// The amount of time that the user has spent exercising during the specified day.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/appleExerciseTime
func (h_ HKActivitySummary) AppleExerciseTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("appleExerciseTime"))
	return rv
}


// SetAppleExerciseTime sets the value of the appleExerciseTime property.
// The amount of time that the user has spent exercising during the specified day.

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/appleExerciseTime
func (h_ HKActivitySummary) SetAppleExerciseTime(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAppleExerciseTime:"), value)
}
// The user’s daily exercise goal.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/appleExerciseTimeGoal
func (h_ HKActivitySummary) AppleExerciseTimeGoal() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("appleExerciseTimeGoal"))
	return rv
}


// SetAppleExerciseTimeGoal sets the value of the appleExerciseTimeGoal property.
// The user’s daily exercise goal.

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/appleExerciseTimeGoal
func (h_ HKActivitySummary) SetAppleExerciseTimeGoal(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAppleExerciseTimeGoal:"), value)
}
// The amount of time the user spent performing activities that involve full-body movements during the specified day.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/appleMoveTime
func (h_ HKActivitySummary) AppleMoveTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("appleMoveTime"))
	return rv
}


// SetAppleMoveTime sets the value of the appleMoveTime property.
// The amount of time the user spent performing activities that involve full-body movements during the specified day.

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/appleMoveTime
func (h_ HKActivitySummary) SetAppleMoveTime(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAppleMoveTime:"), value)
}
// The user’s daily goal for move time.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/appleMoveTimeGoal
func (h_ HKActivitySummary) AppleMoveTimeGoal() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("appleMoveTimeGoal"))
	return rv
}


// SetAppleMoveTimeGoal sets the value of the appleMoveTimeGoal property.
// The user’s daily goal for move time.

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/appleMoveTimeGoal
func (h_ HKActivitySummary) SetAppleMoveTimeGoal(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAppleMoveTimeGoal:"), value)
}
// The number hours in the specified day during which the user has stood and moved for at least a minute per hour.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/appleStandHours
func (h_ HKActivitySummary) AppleStandHours() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("appleStandHours"))
	return rv
}


// SetAppleStandHours sets the value of the appleStandHours property.
// The number hours in the specified day during which the user has stood and moved for at least a minute per hour.

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/appleStandHours
func (h_ HKActivitySummary) SetAppleStandHours(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAppleStandHours:"), value)
}
// The user’s daily goal for stand hours.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/appleStandHoursGoal
func (h_ HKActivitySummary) AppleStandHoursGoal() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("appleStandHoursGoal"))
	return rv
}


// SetAppleStandHoursGoal sets the value of the appleStandHoursGoal property.
// The user’s daily goal for stand hours.

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/appleStandHoursGoal
func (h_ HKActivitySummary) SetAppleStandHoursGoal(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAppleStandHoursGoal:"), value)
}
// The user’s daily goal for exercise time.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/exerciseTimeGoal
func (h_ HKActivitySummary) ExerciseTimeGoal() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("exerciseTimeGoal"))
	return rv
}


// SetExerciseTimeGoal sets the value of the exerciseTimeGoal property.
// The user’s daily goal for exercise time.

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/exerciseTimeGoal
func (h_ HKActivitySummary) SetExerciseTimeGoal(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setExerciseTimeGoal:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/isPaused
func (h_ HKActivitySummary) Paused() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("paused"))
	return rv
}


// SetPaused sets the value of the paused property.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/isPaused
func (h_ HKActivitySummary) SetPaused(value bool) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setPaused:"), value)
}
// The user’s daily goal for stand hours.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/standHoursGoal
func (h_ HKActivitySummary) StandHoursGoal() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("standHoursGoal"))
	return rv
}


// SetStandHoursGoal sets the value of the standHoursGoal property.
// The user’s daily goal for stand hours.

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/standHoursGoal
func (h_ HKActivitySummary) SetStandHoursGoal(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setStandHoursGoal:"), value)
}


