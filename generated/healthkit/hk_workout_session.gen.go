// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKWorkoutSession */


/* debug [class_header]: Header for HKWorkoutSession */
// The class instance for the [HKWorkoutSession] class.
var (
	HKWorkoutSessionClass     _HKWorkoutSessionClass
	HKWorkoutSessionClassOnce sync.Once
)

func getHKWorkoutSessionClass() _HKWorkoutSessionClass {
	HKWorkoutSessionClassOnce.Do(func() {
		HKWorkoutSessionClass = _HKWorkoutSessionClass{objc.GetClass("HKWorkoutSession")}
	})
	return HKWorkoutSessionClass
}

type _HKWorkoutSessionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKWorkoutSession */
// An interface definition for the [HKWorkoutSession] class.
type IHKWorkoutSession interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKWorkoutSession */
	// properties:
	ActivityType() HKWorkoutActivityType
	CurrentActivity() IHKWorkoutActivity
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	EndDate() objc.IObject /* cross-framework: NSDate */
	LocationType() HKWorkoutSessionLocationType
	StartDate() objc.IObject /* cross-framework: NSDate */
	State() HKWorkoutSessionState
	Type() HKWorkoutSessionType
	WorkoutConfiguration() IHKWorkoutConfiguration
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKWorkoutSession */
	// methods:
	BeginNewActivityWithConfigurationDateMetadata(workoutConfiguration IHKWorkoutConfiguration, date objc.IObject /* cross-framework: NSDate */, metadata foundation.IDictionary)
	End()
	EndCurrentActivityOnDate(date objc.IObject /* cross-framework: NSDate */)
	Pause()
	Prepare()
	Resume()
	StartActivityWithDate(date objc.IObject /* cross-framework: NSDate */)
	StopActivityWithDate(date objc.IObject /* cross-framework: NSDate */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKWorkoutSession */
// Alloc allocates a new instance without initialization.
func (hc _HKWorkoutSessionClass) Alloc() HKWorkoutSession {
	rv := objc.Send[HKWorkoutSession](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKWorkoutSessionClass) New() HKWorkoutSession {
	rv := objc.Send[HKWorkoutSession](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKWorkoutSession) Init() HKWorkoutSession {
	rv := objc.Send[HKWorkoutSession](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKWorkoutSession) Autorelease() HKWorkoutSession {
	rv := objc.Send[HKWorkoutSession](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKWorkoutSession creates a new HKWorkoutSession instance.
func NewHKWorkoutSession() HKWorkoutSession {
	return getHKWorkoutSessionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKWorkoutSession */
// A session that tracks a person’s workout.
//
// The session fine-tunes Apple Watch’s sensors for the specified activity. All workout sessions generate high-frequency heart rate samples; however, an outdoor cycling activity generates accurate location data, while an indoor cycling activity doesn’t. Collecting heart rate data on iPhone or iPad requires pairing with an external heart rate sensor because these devices don’t have one. iPhone and iPad can collect various workout metrics, but the system may generate different samples than those specifically requested by an app. You can modify the default types of data collected during a workout. After someone saves a workout, you can access and display summary statistics or chart metrics over time. iPhone typically locks during workouts. For privacy reasons, health data usually isn’t accessible while the device is locked. However, the system can prompt someone to provide your app access to workout data even when their device is locked. You can then display Live Activities on the Lock Screen, providing health metrics without requiring the person to unlock their phone. Siri support extends to the Lock Screen, allowing people to start, pause, resume, or cancel workouts hands-free. You can integrate Siri intents into your apps to enable this functionality. Apple Watch runs one workout session at a time. If a second workout starts while your workout is running, your object receives an error, and your session ends.


// A session that tracks a person’s workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSession
type HKWorkoutSession struct {
	objectivec.Object
}

// HKWorkoutSessionFrom constructs a [HKWorkoutSession] from an unsafe.Pointer.
//
// A session that tracks a person’s workout.
func HKWorkoutSessionFrom(ptr unsafe.Pointer) HKWorkoutSession {
	return HKWorkoutSession{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKWorkoutSession */

// Returns a newly instantiated workout session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSession/init(activityType:locationType:)
func NewHKWorkoutSessionWithActivityTypeLocationType(activityType HKWorkoutActivityType, locationType HKWorkoutSessionLocationType) HKWorkoutSession {
	instance := getHKWorkoutSessionClass().Alloc()
	rv := objc.Send[HKWorkoutSession](instance.ID, objc.Sel("initWithActivityType:locationType:"), activityType, locationType)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKWorkoutSessionWithActivityTypeLocationType */


// Returns a newly instantiated workout session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSession/init(configuration:)
func NewHKWorkoutSessionWithConfigurationError(workoutConfiguration IHKWorkoutConfiguration, error_ objectivec.IObject) HKWorkoutSession {
	instance := getHKWorkoutSessionClass().Alloc()
	rv := objc.Send[HKWorkoutSession](instance.ID, objc.Sel("initWithConfiguration:error:"), workoutConfiguration, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKWorkoutSessionWithConfigurationError */


// Returns a newly instantiated workout session with an associated workout builder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSession/init(healthStore:configuration:)
func NewHKWorkoutSessionWithHealthStoreConfigurationError(healthStore IHKHealthStore, workoutConfiguration IHKWorkoutConfiguration, error_ objectivec.IObject) HKWorkoutSession {
	instance := getHKWorkoutSessionClass().Alloc()
	rv := objc.Send[HKWorkoutSession](instance.ID, objc.Sel("initWithHealthStore:configuration:error:"), healthStore, workoutConfiguration, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKWorkoutSessionWithHealthStoreConfigurationError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKWorkoutSession */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKWorkoutSession */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKWorkoutSession */

// Begins a new workout activity in the workout session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSession/beginNewActivity(configuration:date:metadata:)
func (h_ HKWorkoutSession) BeginNewActivityWithConfigurationDateMetadata(workoutConfiguration IHKWorkoutConfiguration, date objc.IObject /* cross-framework: NSDate */, metadata foundation.IDictionary) {
	objc.Send[objc.ID](h_.ID, objc.Sel("beginNewActivityWithConfiguration:date:metadata:"), workoutConfiguration, date, metadata)
}/* debug [instance_methods/method]: BeginNewActivityWithConfigurationDateMetadata */


// Ends the workout session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSession/end()
func (h_ HKWorkoutSession) End() {
	objc.Send[objc.ID](h_.ID, objc.Sel("end"))
}/* debug [instance_methods/method]: End */


// Ends the current workout activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSession/endCurrentActivity(on:)
func (h_ HKWorkoutSession) EndCurrentActivityOnDate(date objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("endCurrentActivityOnDate:"), date)
}/* debug [instance_methods/method]: EndCurrentActivityOnDate */


// Pauses the workout session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSession/pause()
func (h_ HKWorkoutSession) Pause() {
	objc.Send[objc.ID](h_.ID, objc.Sel("pause"))
}/* debug [instance_methods/method]: Pause */


// Prepares the workout session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSession/prepare()
func (h_ HKWorkoutSession) Prepare() {
	objc.Send[objc.ID](h_.ID, objc.Sel("prepare"))
}/* debug [instance_methods/method]: Prepare */


// Resumes the workout session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSession/resume()
func (h_ HKWorkoutSession) Resume() {
	objc.Send[objc.ID](h_.ID, objc.Sel("resume"))
}/* debug [instance_methods/method]: Resume */


// Starts the workout session activity, and sets the start date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSession/startActivity(with:)
func (h_ HKWorkoutSession) StartActivityWithDate(date objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("startActivityWithDate:"), date)
}/* debug [instance_methods/method]: StartActivityWithDate */


// Stops the workout session activity, and sets the end date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSession/stopActivity(with:)
func (h_ HKWorkoutSession) StopActivityWithDate(date objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("stopActivityWithDate:"), date)
}/* debug [instance_methods/method]: StopActivityWithDate */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKWorkoutSession */

// The workout activity performed during this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSession/activityType
func (h_ HKWorkoutSession) ActivityType() HKWorkoutActivityType {
	rv := objc.Send[HKWorkoutActivityType](h_.ID, objc.Sel("activityType"))
	return rv
}/* debug [instance_properties/getter]: activityType */


// The current workout activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSession/currentActivity
func (h_ HKWorkoutSession) CurrentActivity() IHKWorkoutActivity {
	rv := objc.Send[HKWorkoutActivity](h_.ID, objc.Sel("currentActivity"))
	return rv
}/* debug [instance_properties/getter]: currentActivity */


// The workout session’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSession/delegate
func (h_ HKWorkoutSession) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The workout session’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSession/delegate
func (h_ HKWorkoutSession) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The ending time and date for this workout session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSession/endDate
func (h_ HKWorkoutSession) EndDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](h_.ID, objc.Sel("endDate"))
	return rv
}/* debug [instance_properties/getter]: endDate */


// A value that indicates whether the workout session occurred indoors or outdoors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSession/locationType
func (h_ HKWorkoutSession) LocationType() HKWorkoutSessionLocationType {
	rv := objc.Send[HKWorkoutSessionLocationType](h_.ID, objc.Sel("locationType"))
	return rv
}/* debug [instance_properties/getter]: locationType */


// The starting time and date for this workout session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSession/startDate
func (h_ HKWorkoutSession) StartDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](h_.ID, objc.Sel("startDate"))
	return rv
}/* debug [instance_properties/getter]: startDate */


// The workout session’s current state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSession/state
func (h_ HKWorkoutSession) State() HKWorkoutSessionState {
	rv := objc.Send[HKWorkoutSessionState](h_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */


// A value that indicates whether the session is a primary session or a mirrored session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSession/type
func (h_ HKWorkoutSession) Type() HKWorkoutSessionType {
	rv := objc.Send[HKWorkoutSessionType](h_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// The configuration object that describes this workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSession/workoutConfiguration
func (h_ HKWorkoutSession) WorkoutConfiguration() IHKWorkoutConfiguration {
	rv := objc.Send[HKWorkoutConfiguration](h_.ID, objc.Sel("workoutConfiguration"))
	return rv
}/* debug [instance_properties/getter]: workoutConfiguration */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKWorkoutSession */


