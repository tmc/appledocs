// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKActivitySummary */


/* debug [class_header]: Header for HKActivitySummary */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKActivitySummary */
// An interface definition for the [HKActivitySummary] class.
type IHKActivitySummary interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKActivitySummary */
	// properties:
	ActiveEnergyBurned() IHKQuantity
	SetActiveEnergyBurned(value IHKQuantity)
	ActiveEnergyBurnedGoal() IHKQuantity
	SetActiveEnergyBurnedGoal(value IHKQuantity)
	ActivityMoveMode() HKActivityMoveMode
	SetActivityMoveMode(value HKActivityMoveMode)
	AppleExerciseTime() IHKQuantity
	SetAppleExerciseTime(value IHKQuantity)
	AppleExerciseTimeGoal() IHKQuantity
	SetAppleExerciseTimeGoal(value IHKQuantity)
	AppleMoveTime() IHKQuantity
	SetAppleMoveTime(value IHKQuantity)
	AppleMoveTimeGoal() IHKQuantity
	SetAppleMoveTimeGoal(value IHKQuantity)
	AppleStandHours() IHKQuantity
	SetAppleStandHours(value IHKQuantity)
	AppleStandHoursGoal() IHKQuantity
	SetAppleStandHoursGoal(value IHKQuantity)
	ExerciseTimeGoal() IHKQuantity
	SetExerciseTimeGoal(value IHKQuantity)
	Paused() bool
	SetPaused(value bool)
	StandHoursGoal() IHKQuantity
	SetStandHoursGoal(value IHKQuantity)
	IsPaused() bool
	SetIsPaused(value bool)
	HKPredicateKeyPathDateComponents() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKActivitySummary */
	// methods:
	DateComponentsForCalendar(calendar foundation.Calendar) foundation.DateComponents
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKActivitySummary */
// Alloc allocates a new instance without initialization.
func (hc _HKActivitySummaryClass) Alloc() HKActivitySummary {
	rv := objc.Send[HKActivitySummary](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKActivitySummary */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKActivitySummary *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKActivitySummary */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKActivitySummary */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKActivitySummary */

// Date components that uniquely identify the day represented by the summary object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/dateComponents(for:)
func (h_ HKActivitySummary) DateComponentsForCalendar(calendar foundation.Calendar) foundation.DateComponents {
	rv := objc.Send[foundation.DateComponents](h_.ID, objc.Sel("dateComponentsForCalendar:"), calendar)
	return rv
}/* debug [instance_methods/method]: DateComponentsForCalendar */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKActivitySummary */

// The amount of active energy the user burned during the specified day.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/activeEnergyBurned
func (h_ HKActivitySummary) ActiveEnergyBurned() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("activeEnergyBurned"))
	return rv
}/* debug [instance_properties/getter]: activeEnergyBurned */


// The amount of active energy the user burned during the specified day.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/activeEnergyBurned
func (h_ HKActivitySummary) SetActiveEnergyBurned(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setActiveEnergyBurned:"), value)
}/* debug [instance_properties/setter]: activeEnergyBurned */


// The user’s daily goal for active energy burned.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/activeEnergyBurnedGoal
func (h_ HKActivitySummary) ActiveEnergyBurnedGoal() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("activeEnergyBurnedGoal"))
	return rv
}/* debug [instance_properties/getter]: activeEnergyBurnedGoal */


// The user’s daily goal for active energy burned.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/activeEnergyBurnedGoal
func (h_ HKActivitySummary) SetActiveEnergyBurnedGoal(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setActiveEnergyBurnedGoal:"), value)
}/* debug [instance_properties/setter]: activeEnergyBurnedGoal */


// The move mode that they system used for this activity summary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/activityMoveMode
func (h_ HKActivitySummary) ActivityMoveMode() HKActivityMoveMode {
	rv := objc.Send[HKActivityMoveMode](h_.ID, objc.Sel("activityMoveMode"))
	return rv
}/* debug [instance_properties/getter]: activityMoveMode */


// The move mode that they system used for this activity summary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/activityMoveMode
func (h_ HKActivitySummary) SetActivityMoveMode(value HKActivityMoveMode) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setActivityMoveMode:"), value)
}/* debug [instance_properties/setter]: activityMoveMode */


// The amount of time that the user has spent exercising during the specified day.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/appleExerciseTime
func (h_ HKActivitySummary) AppleExerciseTime() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("appleExerciseTime"))
	return rv
}/* debug [instance_properties/getter]: appleExerciseTime */


// The amount of time that the user has spent exercising during the specified day.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/appleExerciseTime
func (h_ HKActivitySummary) SetAppleExerciseTime(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAppleExerciseTime:"), value)
}/* debug [instance_properties/setter]: appleExerciseTime */


// The user’s daily exercise goal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/appleExerciseTimeGoal
func (h_ HKActivitySummary) AppleExerciseTimeGoal() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("appleExerciseTimeGoal"))
	return rv
}/* debug [instance_properties/getter]: appleExerciseTimeGoal */


// The user’s daily exercise goal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/appleExerciseTimeGoal
func (h_ HKActivitySummary) SetAppleExerciseTimeGoal(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAppleExerciseTimeGoal:"), value)
}/* debug [instance_properties/setter]: appleExerciseTimeGoal */


// The amount of time the user spent performing activities that involve full-body movements during the specified day.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/appleMoveTime
func (h_ HKActivitySummary) AppleMoveTime() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("appleMoveTime"))
	return rv
}/* debug [instance_properties/getter]: appleMoveTime */


// The amount of time the user spent performing activities that involve full-body movements during the specified day.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/appleMoveTime
func (h_ HKActivitySummary) SetAppleMoveTime(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAppleMoveTime:"), value)
}/* debug [instance_properties/setter]: appleMoveTime */


// The user’s daily goal for move time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/appleMoveTimeGoal
func (h_ HKActivitySummary) AppleMoveTimeGoal() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("appleMoveTimeGoal"))
	return rv
}/* debug [instance_properties/getter]: appleMoveTimeGoal */


// The user’s daily goal for move time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/appleMoveTimeGoal
func (h_ HKActivitySummary) SetAppleMoveTimeGoal(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAppleMoveTimeGoal:"), value)
}/* debug [instance_properties/setter]: appleMoveTimeGoal */


// The number hours in the specified day during which the user has stood and moved for at least a minute per hour.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/appleStandHours
func (h_ HKActivitySummary) AppleStandHours() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("appleStandHours"))
	return rv
}/* debug [instance_properties/getter]: appleStandHours */


// The number hours in the specified day during which the user has stood and moved for at least a minute per hour.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/appleStandHours
func (h_ HKActivitySummary) SetAppleStandHours(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAppleStandHours:"), value)
}/* debug [instance_properties/setter]: appleStandHours */


// The user’s daily goal for stand hours.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/appleStandHoursGoal
func (h_ HKActivitySummary) AppleStandHoursGoal() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("appleStandHoursGoal"))
	return rv
}/* debug [instance_properties/getter]: appleStandHoursGoal */


// The user’s daily goal for stand hours.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/appleStandHoursGoal
func (h_ HKActivitySummary) SetAppleStandHoursGoal(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAppleStandHoursGoal:"), value)
}/* debug [instance_properties/setter]: appleStandHoursGoal */


// The user’s daily goal for exercise time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/exerciseTimeGoal
func (h_ HKActivitySummary) ExerciseTimeGoal() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("exerciseTimeGoal"))
	return rv
}/* debug [instance_properties/getter]: exerciseTimeGoal */


// The user’s daily goal for exercise time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/exerciseTimeGoal
func (h_ HKActivitySummary) SetExerciseTimeGoal(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setExerciseTimeGoal:"), value)
}/* debug [instance_properties/setter]: exerciseTimeGoal */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/isPaused
func (h_ HKActivitySummary) Paused() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("paused"))
	return rv
}/* debug [instance_properties/getter]: paused */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/isPaused
func (h_ HKActivitySummary) SetPaused(value bool) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setPaused:"), value)
}/* debug [instance_properties/setter]: paused */


// The user’s daily goal for stand hours.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/standHoursGoal
func (h_ HKActivitySummary) StandHoursGoal() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("standHoursGoal"))
	return rv
}/* debug [instance_properties/getter]: standHoursGoal */


// The user’s daily goal for stand hours.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummary/standHoursGoal
func (h_ HKActivitySummary) SetStandHoursGoal(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setStandHoursGoal:"), value)
}/* debug [instance_properties/setter]: standHoursGoal */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/ispaused
func (h_ HKActivitySummary) IsPaused() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("isPaused"))
	return rv
}/* debug [instance_properties/getter]: isPaused */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/ispaused
func (h_ HKActivitySummary) SetIsPaused(value bool) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setIsPaused:"), value)
}/* debug [instance_properties/setter]: isPaused */


// The key path for accessing an activity summary’s date components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathdatecomponents
func (h_ HKActivitySummary) HKPredicateKeyPathDateComponents() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathDateComponents"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathDateComponents */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKActivitySummary */



