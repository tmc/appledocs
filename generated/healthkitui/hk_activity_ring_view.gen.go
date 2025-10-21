// Code generated from Apple documentation for HealthKitUI. DO NOT EDIT.

package healthkitui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [HKActivityRingView] class.
var (
	HKActivityRingViewClass     _HKActivityRingViewClass
	HKActivityRingViewClassOnce sync.Once
)

func getHKActivityRingViewClass() _HKActivityRingViewClass {
	HKActivityRingViewClassOnce.Do(func() {
		HKActivityRingViewClass = _HKActivityRingViewClass{objc.GetClass("HKActivityRingView")}
	})
	return HKActivityRingViewClass
}

type _HKActivityRingViewClass struct {
	class objc.Class
}

// An interface definition for the [HKActivityRingView] class.
type IHKActivityRingView interface {
	appkit.IView
	SetActivitySummaryAnimated(activitySummary unsafe.Pointer, animated bool)
}

// A view that uses the Move, Exercise, and Stand activity rings to display data from a HealthKit activity summary object.
//
// Use to display data from an object. For example, the following image shows how the rings can display a summary view of a person’s activity. To display activity summary data from the HealthKit store, use an object. You can also instantiate and display your own objects, as needed. The activity ring view always appears as a black rectangle with colored concentric rings. The rings are centered in the view and are sized to fit the available space. The activity ring view displays different rings depending on the properties defined in the ring view’s property. When the view’s has set to and values for and the ring only displays the red Move ring. Otherwise, it displays the Move, Exercise, and Stand activity as red, green, and blue concentric rings. Summary data from the HealthKit store only displays the Move ring when the person hasn’t paired an Apple Watch. The rings can display as either empty or with a dot at the top of the ring to display a lack of data. Empty rings indicate that the activity summary is missing, and a dot at the top indicates that the activity summary’s values are set to zero. If the ring has a -valued quantity properties, the rings appear empty. Use this to indicate that there is no summary data available for the specified day. For example, dates in the future. If the summary has zero-valued quantities set for its value properties, the ring displays a dot at the top of the ring. Use this to indicate that the person hasn’t burned any active calories, exercised, or earned any stand hours for the specified day. To display data for a ring, the object must have a non- quantity for both the corresponding value property and the goal property. Move only ring properties: The activity ring view colors a percentage of each ring based on these properties, as shown here: The following code snippet shows how to manually display only the Move ring:
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKitUI/HKActivityRingView
type HKActivityRingView struct {
	appkit.View
}

// HKActivityRingViewFrom constructs a [HKActivityRingView] from an unsafe.Pointer.
//
// A view that uses the Move, Exercise, and Stand activity rings to display data from a HealthKit activity summary object.
func HKActivityRingViewFrom(ptr unsafe.Pointer) HKActivityRingView {
	return HKActivityRingView{
		View: appkit.ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKActivityRingViewClass) Alloc() HKActivityRingView {
	rv := objc.Send[HKActivityRingView](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKActivityRingViewClass) New() HKActivityRingView {
	rv := objc.Send[HKActivityRingView](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKActivityRingView) Init() HKActivityRingView {
	rv := objc.Send[HKActivityRingView](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKActivityRingView) Autorelease() HKActivityRingView {
	rv := objc.Send[HKActivityRingView](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKActivityRingView creates a new HKActivityRingView instance.
func NewHKActivityRingView() HKActivityRingView {
	return getHKActivityRingViewClass().New()
}


// Sets the activity summary displayed by the activity ring view.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKitUI/HKActivityRingView/setActivitySummary(_:animated:)
func (h_ HKActivityRingView) SetActivitySummaryAnimated(activitySummary unsafe.Pointer, animated bool) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setActivitySummary:animated:"), activitySummary, animated)
}

// The number hours in the specified day during which the user has stood and moved for at least a minute per hour.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/applestandhours
func (h_ HKActivityRingView) AppleStandHours() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("appleStandHours"))
	return rv
}


// SetAppleStandHours sets the value of the appleStandHours property.
// The number hours in the specified day during which the user has stood and moved for at least a minute per hour.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/applestandhours
func (h_ HKActivityRingView) SetAppleStandHours(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAppleStandHours:"), value)
}

// The user’s daily exercise goal.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/appleexercisetimegoal
func (h_ HKActivityRingView) AppleExerciseTimeGoal() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("appleExerciseTimeGoal"))
	return rv
}


// SetAppleExerciseTimeGoal sets the value of the appleExerciseTimeGoal property.
// The user’s daily exercise goal.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/appleexercisetimegoal
func (h_ HKActivityRingView) SetAppleExerciseTimeGoal(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAppleExerciseTimeGoal:"), value)
}

// The user’s daily goal for move time.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/applemovetimegoal
func (h_ HKActivityRingView) AppleMoveTimeGoal() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("appleMoveTimeGoal"))
	return rv
}


// SetAppleMoveTimeGoal sets the value of the appleMoveTimeGoal property.
// The user’s daily goal for move time.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/applemovetimegoal
func (h_ HKActivityRingView) SetAppleMoveTimeGoal(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAppleMoveTimeGoal:"), value)
}

// The user’s daily goal for active energy burned.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/activeenergyburnedgoal
func (h_ HKActivityRingView) ActiveEnergyBurnedGoal() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("activeEnergyBurnedGoal"))
	return rv
}


// SetActiveEnergyBurnedGoal sets the value of the activeEnergyBurnedGoal property.
// The user’s daily goal for active energy burned.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/activeenergyburnedgoal
func (h_ HKActivityRingView) SetActiveEnergyBurnedGoal(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setActiveEnergyBurnedGoal:"), value)
}

// The amount of time that the user has spent exercising during the specified day.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/appleexercisetime
func (h_ HKActivityRingView) AppleExerciseTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("appleExerciseTime"))
	return rv
}


// SetAppleExerciseTime sets the value of the appleExerciseTime property.
// The amount of time that the user has spent exercising during the specified day.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/appleexercisetime
func (h_ HKActivityRingView) SetAppleExerciseTime(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAppleExerciseTime:"), value)
}

// The user’s daily goal for stand hours.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/applestandhoursgoal
func (h_ HKActivityRingView) AppleStandHoursGoal() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("appleStandHoursGoal"))
	return rv
}


// SetAppleStandHoursGoal sets the value of the appleStandHoursGoal property.
// The user’s daily goal for stand hours.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/applestandhoursgoal
func (h_ HKActivityRingView) SetAppleStandHoursGoal(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAppleStandHoursGoal:"), value)
}

// A property that contains the movement mode value.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitymovemodeobject/activitymovemode
func (h_ HKActivityRingView) ActivityMoveMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("activityMoveMode"))
	return rv
}


// SetActivityMoveMode sets the value of the activityMoveMode property.
// A property that contains the movement mode value.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitymovemodeobject/activitymovemode
func (h_ HKActivityRingView) SetActivityMoveMode(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setActivityMoveMode:"), value)
}

// The amount of time the user spent performing activities that involve full-body movements during the specified day.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/applemovetime
func (h_ HKActivityRingView) AppleMoveTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("appleMoveTime"))
	return rv
}


// SetAppleMoveTime sets the value of the appleMoveTime property.
// The amount of time the user spent performing activities that involve full-body movements during the specified day.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/applemovetime
func (h_ HKActivityRingView) SetAppleMoveTime(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAppleMoveTime:"), value)
}

// The amount of active energy the user burned during the specified day.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/activeenergyburned
func (h_ HKActivityRingView) ActiveEnergyBurned() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("activeEnergyBurned"))
	return rv
}


// SetActiveEnergyBurned sets the value of the activeEnergyBurned property.
// The amount of active energy the user burned during the specified day.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummary/activeenergyburned
func (h_ HKActivityRingView) SetActiveEnergyBurned(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setActiveEnergyBurned:"), value)
}

// The active summary displayed by the activity ring view.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKitUI/HKActivityRingView/activitySummary
func (h_ HKActivityRingView) ActivitySummary() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("activitySummary"))
	return rv
}


// SetActivitySummary sets the value of the activitySummary property.
// The active summary displayed by the activity ring view.

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKitUI/HKActivityRingView/activitySummary
func (h_ HKActivityRingView) SetActivitySummary(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setActivitySummary:"), value)
}




