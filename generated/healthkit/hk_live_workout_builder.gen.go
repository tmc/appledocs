// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class HKLiveWorkoutBuilder */


/* debug [class_header]: Header for HKLiveWorkoutBuilder */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKLiveWorkoutBuilder */
// An interface definition for the [HKLiveWorkoutBuilder] class.
type IHKLiveWorkoutBuilder interface {
	IHKWorkoutBuilder
	
/* debug [class_interface_properties]: Properties for HKLiveWorkoutBuilder */
	// properties:
	CurrentWorkoutActivity() IHKWorkoutActivity
	DataSource() IHKLiveWorkoutDataSource
	SetDataSource(value IHKLiveWorkoutDataSource)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	ElapsedTime() float64
	ShouldCollectWorkoutEvents() bool
	SetShouldCollectWorkoutEvents(value bool)
	WorkoutSession() IHKWorkoutSession
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKLiveWorkoutBuilder */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKLiveWorkoutBuilder */
// Alloc allocates a new instance without initialization.
func (hc _HKLiveWorkoutBuilderClass) Alloc() HKLiveWorkoutBuilder {
	rv := objc.Send[HKLiveWorkoutBuilder](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKLiveWorkoutBuilder */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKLiveWorkoutBuilder *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKLiveWorkoutBuilder */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKLiveWorkoutBuilder */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKLiveWorkoutBuilder */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKLiveWorkoutBuilder */

// The current workout activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKLiveWorkoutBuilder/currentWorkoutActivity
func (h_ HKLiveWorkoutBuilder) CurrentWorkoutActivity() IHKWorkoutActivity {
	rv := objc.Send[HKWorkoutActivity](h_.ID, objc.Sel("currentWorkoutActivity"))
	return rv
}/* debug [instance_properties/getter]: currentWorkoutActivity */


// A data source that provides live data from a workout session automatically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKLiveWorkoutBuilder/dataSource
func (h_ HKLiveWorkoutBuilder) DataSource() IHKLiveWorkoutDataSource {
	rv := objc.Send[HKLiveWorkoutDataSource](h_.ID, objc.Sel("dataSource"))
	return rv
}/* debug [instance_properties/getter]: dataSource */


// A data source that provides live data from a workout session automatically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKLiveWorkoutBuilder/dataSource
func (h_ HKLiveWorkoutBuilder) SetDataSource(value IHKLiveWorkoutDataSource) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setDataSource:"), value)
}/* debug [instance_properties/setter]: dataSource */


// The live builder’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKLiveWorkoutBuilder/delegate
func (h_ HKLiveWorkoutBuilder) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The live builder’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKLiveWorkoutBuilder/delegate
func (h_ HKLiveWorkoutBuilder) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The elapsed time for the workout based on the builder’s current contents, including pauses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKLiveWorkoutBuilder/elapsedTime
func (h_ HKLiveWorkoutBuilder) ElapsedTime() float64 {
	rv := objc.Send[float64](h_.ID, objc.Sel("elapsedTime"))
	return rv
}/* debug [instance_properties/getter]: elapsedTime */


// A Boolean value that determines whether the workout builder automatically adds events generated by the workout session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKLiveWorkoutBuilder/shouldCollectWorkoutEvents
func (h_ HKLiveWorkoutBuilder) ShouldCollectWorkoutEvents() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("shouldCollectWorkoutEvents"))
	return rv
}/* debug [instance_properties/getter]: shouldCollectWorkoutEvents */


// A Boolean value that determines whether the workout builder automatically adds events generated by the workout session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKLiveWorkoutBuilder/shouldCollectWorkoutEvents
func (h_ HKLiveWorkoutBuilder) SetShouldCollectWorkoutEvents(value bool) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setShouldCollectWorkoutEvents:"), value)
}/* debug [instance_properties/setter]: shouldCollectWorkoutEvents */


// The workout session created by the data source and associated with this builder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKLiveWorkoutBuilder/workoutSession
func (h_ HKLiveWorkoutBuilder) WorkoutSession() IHKWorkoutSession {
	rv := objc.Send[HKWorkoutSession](h_.ID, objc.Sel("workoutSession"))
	return rv
}/* debug [instance_properties/getter]: workoutSession */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKLiveWorkoutBuilder */



