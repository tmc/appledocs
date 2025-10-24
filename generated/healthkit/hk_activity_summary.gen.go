// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	AppleMoveTime() IHKQuantity
	SetAppleMoveTime(value IHKQuantity)
	ActiveEnergyBurned() IHKQuantity
	SetActiveEnergyBurned(value IHKQuantity)
	ActiveEnergyBurnedGoal() IHKQuantity
	SetActiveEnergyBurnedGoal(value IHKQuantity)
	ActivityMoveMode() unsafe.Pointer
	SetActivityMoveMode(value unsafe.Pointer)
	AppleExerciseTime() IHKQuantity
	SetAppleExerciseTime(value IHKQuantity)
	AppleExerciseTimeGoal() IHKQuantity
	SetAppleExerciseTimeGoal(value IHKQuantity)
	AppleMoveTimeGoal() IHKQuantity
	SetAppleMoveTimeGoal(value IHKQuantity)
	AppleStandHours() IHKQuantity
	SetAppleStandHours(value IHKQuantity)
	AppleStandHoursGoal() IHKQuantity
	SetAppleStandHoursGoal(value IHKQuantity)
	ExerciseTimeGoal() IHKQuantity
	SetExerciseTimeGoal(value IHKQuantity)
	IsPaused() bool
	SetIsPaused(value bool)
	StandHoursGoal() IHKQuantity
	SetStandHoursGoal(value IHKQuantity)
	HKPredicateKeyPathDateComponents() objc.IObject /* cross-framework: NSString */
	// methods:
}

// An object that contains the move, exercise, and stand data for a given day.
//
// You can read objects from the HealthKit store using an object. Unlike the subclasses, instances are mutable, but changes made to the object’s properties have no affect on the values in the HealthKit store. You can instantiate your own objects (if needed), but you can’t save objects to the store. You can display an active summary in iOS using the class or in watchOS using the class.


// An object that contains the move, exercise, and stand data for a given day.
//
// [Full Topic]
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



// The amount of time the user spent performing activities that involve full-body movements during the specified day.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/appleMoveTime
func (h_ HKActivitySummary) AppleMoveTime() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("appleMoveTime"))
	return rv
}


// The amount of time the user spent performing activities that involve full-body movements during the specified day.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/appleMoveTime
func (h_ HKActivitySummary) SetAppleMoveTime(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAppleMoveTime:"), value)
}


// The amount of active energy the user burned during the specified day.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/activeenergyburned
func (h_ HKActivitySummary) ActiveEnergyBurned() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("activeEnergyBurned"))
	return rv
}


// The amount of active energy the user burned during the specified day.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/activeenergyburned
func (h_ HKActivitySummary) SetActiveEnergyBurned(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setActiveEnergyBurned:"), value)
}


// The user’s daily goal for active energy burned.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/activeenergyburnedgoal
func (h_ HKActivitySummary) ActiveEnergyBurnedGoal() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("activeEnergyBurnedGoal"))
	return rv
}


// The user’s daily goal for active energy burned.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/activeenergyburnedgoal
func (h_ HKActivitySummary) SetActiveEnergyBurnedGoal(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setActiveEnergyBurnedGoal:"), value)
}


// The move mode that they system used for this activity summary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/activitymovemode
func (h_ HKActivitySummary) ActivityMoveMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("activityMoveMode"))
	return rv
}


// The move mode that they system used for this activity summary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/activitymovemode
func (h_ HKActivitySummary) SetActivityMoveMode(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setActivityMoveMode:"), value)
}


// The amount of time that the user has spent exercising during the specified day.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/appleexercisetime
func (h_ HKActivitySummary) AppleExerciseTime() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("appleExerciseTime"))
	return rv
}


// The amount of time that the user has spent exercising during the specified day.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/appleexercisetime
func (h_ HKActivitySummary) SetAppleExerciseTime(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAppleExerciseTime:"), value)
}


// The user’s daily exercise goal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/appleexercisetimegoal
func (h_ HKActivitySummary) AppleExerciseTimeGoal() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("appleExerciseTimeGoal"))
	return rv
}


// The user’s daily exercise goal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/appleexercisetimegoal
func (h_ HKActivitySummary) SetAppleExerciseTimeGoal(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAppleExerciseTimeGoal:"), value)
}


// The user’s daily goal for move time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/applemovetimegoal
func (h_ HKActivitySummary) AppleMoveTimeGoal() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("appleMoveTimeGoal"))
	return rv
}


// The user’s daily goal for move time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/applemovetimegoal
func (h_ HKActivitySummary) SetAppleMoveTimeGoal(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAppleMoveTimeGoal:"), value)
}


// The number hours in the specified day during which the user has stood and moved for at least a minute per hour.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/applestandhours
func (h_ HKActivitySummary) AppleStandHours() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("appleStandHours"))
	return rv
}


// The number hours in the specified day during which the user has stood and moved for at least a minute per hour.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/applestandhours
func (h_ HKActivitySummary) SetAppleStandHours(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAppleStandHours:"), value)
}


// The user’s daily goal for stand hours.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/applestandhoursgoal
func (h_ HKActivitySummary) AppleStandHoursGoal() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("appleStandHoursGoal"))
	return rv
}


// The user’s daily goal for stand hours.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/applestandhoursgoal
func (h_ HKActivitySummary) SetAppleStandHoursGoal(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAppleStandHoursGoal:"), value)
}


// The user’s daily goal for exercise time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/exercisetimegoal
func (h_ HKActivitySummary) ExerciseTimeGoal() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("exerciseTimeGoal"))
	return rv
}


// The user’s daily goal for exercise time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/exercisetimegoal
func (h_ HKActivitySummary) SetExerciseTimeGoal(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setExerciseTimeGoal:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/ispaused
func (h_ HKActivitySummary) IsPaused() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("isPaused"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/ispaused
func (h_ HKActivitySummary) SetIsPaused(value bool) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setIsPaused:"), value)
}


// The user’s daily goal for stand hours.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/standhoursgoal
func (h_ HKActivitySummary) StandHoursGoal() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("standHoursGoal"))
	return rv
}


// The user’s daily goal for stand hours.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/standhoursgoal
func (h_ HKActivitySummary) SetStandHoursGoal(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setStandHoursGoal:"), value)
}


// The key path for accessing an activity summary’s date components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathdatecomponents
func (h_ HKActivitySummary) HKPredicateKeyPathDateComponents() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathDateComponents"))
	return rv
}



