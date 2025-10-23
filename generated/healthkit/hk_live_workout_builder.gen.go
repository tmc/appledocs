// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	CurrentWorkoutActivity() IHKWorkoutActivity
	SetCurrentWorkoutActivity(value IHKWorkoutActivity)
	DataSource() IHKLiveWorkoutDataSource
	SetDataSource(value IHKLiveWorkoutDataSource)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	ElapsedTime() unsafe.Pointer
	SetElapsedTime(value unsafe.Pointer)
	ShouldCollectWorkoutEvents() bool /* primitive/slice/pointer. */
	SetShouldCollectWorkoutEvents(value bool /* primitive/slice/pointer. */)
	WorkoutSession() IHKWorkoutSession
	SetWorkoutSession(value IHKWorkoutSession)
	// methods:
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
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkliveworkoutbuilder/currentworkoutactivity
func (h_ HKLiveWorkoutBuilder) CurrentWorkoutActivity() IHKWorkoutActivity {
	rv := objc.Send[HKWorkoutActivity](h_.ID, objc.Sel("currentWorkoutActivity"))
	return rv
}


// The current workout activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkliveworkoutbuilder/currentworkoutactivity
func (h_ HKLiveWorkoutBuilder) SetCurrentWorkoutActivity(value IHKWorkoutActivity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setCurrentWorkoutActivity:"), value)
}


// A data source that provides live data from a workout session automatically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkliveworkoutbuilder/datasource
func (h_ HKLiveWorkoutBuilder) DataSource() IHKLiveWorkoutDataSource {
	rv := objc.Send[HKLiveWorkoutDataSource](h_.ID, objc.Sel("dataSource"))
	return rv
}


// A data source that provides live data from a workout session automatically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkliveworkoutbuilder/datasource
func (h_ HKLiveWorkoutBuilder) SetDataSource(value IHKLiveWorkoutDataSource) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setDataSource:"), value)
}


// The live builder’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkliveworkoutbuilder/delegate
func (h_ HKLiveWorkoutBuilder) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("delegate"))
	return rv
}


// The live builder’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkliveworkoutbuilder/delegate
func (h_ HKLiveWorkoutBuilder) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setDelegate:"), value)
}


// The elapsed time for the workout based on the builder’s current contents, including pauses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkliveworkoutbuilder/elapsedtime
func (h_ HKLiveWorkoutBuilder) ElapsedTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("elapsedTime"))
	return rv
}


// The elapsed time for the workout based on the builder’s current contents, including pauses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkliveworkoutbuilder/elapsedtime
func (h_ HKLiveWorkoutBuilder) SetElapsedTime(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setElapsedTime:"), value)
}


// A Boolean value that determines whether the workout builder automatically adds events generated by the workout session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkliveworkoutbuilder/shouldcollectworkoutevents
func (h_ HKLiveWorkoutBuilder) ShouldCollectWorkoutEvents() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](h_.ID, objc.Sel("shouldCollectWorkoutEvents"))
	return rv
}


// A Boolean value that determines whether the workout builder automatically adds events generated by the workout session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkliveworkoutbuilder/shouldcollectworkoutevents
func (h_ HKLiveWorkoutBuilder) SetShouldCollectWorkoutEvents(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setShouldCollectWorkoutEvents:"), value)
}


// The workout session created by the data source and associated with this builder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkliveworkoutbuilder/workoutsession
func (h_ HKLiveWorkoutBuilder) WorkoutSession() IHKWorkoutSession {
	rv := objc.Send[HKWorkoutSession](h_.ID, objc.Sel("workoutSession"))
	return rv
}


// The workout session created by the data source and associated with this builder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkliveworkoutbuilder/workoutsession
func (h_ HKLiveWorkoutBuilder) SetWorkoutSession(value IHKWorkoutSession) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setWorkoutSession:"), value)
}



