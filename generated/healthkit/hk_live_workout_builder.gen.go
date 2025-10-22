// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [HKLiveWorkoutBuilder] class.
var (
	HKLiveWorkoutBuilderClass     _HKLiveWorkoutBuilderClass
	HKLiveWorkoutBuilderClassOnce sync.Once
)

func getHKLiveWorkoutBuilderClass() _HKLiveWorkoutBuilderClass {
	HKLiveWorkoutBuilderClassOnce.Do(func() {
		HKLiveWorkoutBuilderClass = _HKLiveWorkoutBuilderClass{objc.GetClass("HKLiveWorkoutBuilder")}
	})
	return HKLiveWorkoutBuilderClass
}

type _HKLiveWorkoutBuilderClass struct {
	class objc.Class
}

// An interface definition for the [HKLiveWorkoutBuilder] class.
type IHKLiveWorkoutBuilder interface {
	IHKWorkoutBuilder
	CurrentWorkoutActivity() HKWorkoutActivity
	DataSource() HKLiveWorkoutDataSource
	SetDataSource(value IHKLiveWorkoutDataSource)
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	ElapsedTime() foundation.TimeInterval
	ShouldCollectWorkoutEvents() bool
	SetShouldCollectWorkoutEvents(value bool)
	WorkoutSession() HKWorkoutSession
}

// A builder object that constructs a workout incrementally based on live data from an active workout session.
//
// Use a live workout builder to create an sample during an active . For complete instructions on running workout sessions on Apple Watch, see .


// A builder object that constructs a workout incrementally based on live data from an active workout session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKLiveWorkoutBuilder

type HKLiveWorkoutBuilder struct {
	HKWorkoutBuilder
}

// HKLiveWorkoutBuilderFrom constructs a [HKLiveWorkoutBuilder] from an unsafe.Pointer.
//
// A builder object that constructs a workout incrementally based on live data from an active workout session.
func HKLiveWorkoutBuilderFrom(ptr unsafe.Pointer) HKLiveWorkoutBuilder {
	return HKLiveWorkoutBuilder{
		HKWorkoutBuilder: HKWorkoutBuilderFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKLiveWorkoutBuilderClass) Alloc() HKLiveWorkoutBuilder {
	rv := objc.Send[HKLiveWorkoutBuilder](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKLiveWorkoutBuilderClass) New() HKLiveWorkoutBuilder {
	rv := objc.Send[HKLiveWorkoutBuilder](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKLiveWorkoutBuilder) Init() HKLiveWorkoutBuilder {
	rv := objc.Send[HKLiveWorkoutBuilder](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKLiveWorkoutBuilder) Autorelease() HKLiveWorkoutBuilder {
	rv := objc.Send[HKLiveWorkoutBuilder](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKLiveWorkoutBuilder creates a new HKLiveWorkoutBuilder instance.
func NewHKLiveWorkoutBuilder() HKLiveWorkoutBuilder {
	return getHKLiveWorkoutBuilderClass().New()
}



// The current workout activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKLiveWorkoutBuilder/currentWorkoutActivity

func (h_ HKLiveWorkoutBuilder) CurrentWorkoutActivity() HKWorkoutActivity {
	rv := objc.Send[HKWorkoutActivity](h_.ID, objc.Sel("currentWorkoutActivity"))
	return rv
}


// A data source that provides live data from a workout session automatically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKLiveWorkoutBuilder/dataSource

func (h_ HKLiveWorkoutBuilder) DataSource() HKLiveWorkoutDataSource {
	rv := objc.Send[HKLiveWorkoutDataSource](h_.ID, objc.Sel("dataSource"))
	return rv
}


// A data source that provides live data from a workout session automatically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKLiveWorkoutBuilder/dataSource

func (h_ HKLiveWorkoutBuilder) SetDataSource(value IHKLiveWorkoutDataSource) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setDataSource:"), value)
}


// The live builder’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKLiveWorkoutBuilder/delegate

func (h_ HKLiveWorkoutBuilder) Delegate() objc.ID {
	rv := objc.Send[objc.ID](h_.ID, objc.Sel("delegate"))
	return rv
}


// The live builder’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKLiveWorkoutBuilder/delegate

func (h_ HKLiveWorkoutBuilder) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setDelegate:"), value)
}


// The elapsed time for the workout based on the builder’s current contents, including pauses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKLiveWorkoutBuilder/elapsedTime

func (h_ HKLiveWorkoutBuilder) ElapsedTime() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](h_.ID, objc.Sel("elapsedTime"))
	return rv
}


// A Boolean value that determines whether the workout builder automatically adds events generated by the workout session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKLiveWorkoutBuilder/shouldCollectWorkoutEvents

func (h_ HKLiveWorkoutBuilder) ShouldCollectWorkoutEvents() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("shouldCollectWorkoutEvents"))
	return rv
}


// A Boolean value that determines whether the workout builder automatically adds events generated by the workout session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKLiveWorkoutBuilder/shouldCollectWorkoutEvents

func (h_ HKLiveWorkoutBuilder) SetShouldCollectWorkoutEvents(value bool) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setShouldCollectWorkoutEvents:"), value)
}


// The workout session created by the data source and associated with this builder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKLiveWorkoutBuilder/workoutSession

func (h_ HKLiveWorkoutBuilder) WorkoutSession() HKWorkoutSession {
	rv := objc.Send[HKWorkoutSession](h_.ID, objc.Sel("workoutSession"))
	return rv
}



