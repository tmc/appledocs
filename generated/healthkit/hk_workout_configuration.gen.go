// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKWorkoutConfiguration */


/* debug [class_header]: Header for HKWorkoutConfiguration */
// The class instance for the [HKWorkoutConfiguration] class.
var (
	HKWorkoutConfigurationClass     _HKWorkoutConfigurationClass
	HKWorkoutConfigurationClassOnce sync.Once
)

func getHKWorkoutConfigurationClass() _HKWorkoutConfigurationClass {
	HKWorkoutConfigurationClassOnce.Do(func() {
		HKWorkoutConfigurationClass = _HKWorkoutConfigurationClass{objc.GetClass("HKWorkoutConfiguration")}
	})
	return HKWorkoutConfigurationClass
}

type _HKWorkoutConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKWorkoutConfiguration */
// An interface definition for the [HKWorkoutConfiguration] class.
type IHKWorkoutConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKWorkoutConfiguration */
	// properties:
	ActivityType() HKWorkoutActivityType
	SetActivityType(value HKWorkoutActivityType)
	LapLength() IHKQuantity
	SetLapLength(value IHKQuantity)
	LocationType() HKWorkoutSessionLocationType
	SetLocationType(value HKWorkoutSessionLocationType)
	SwimmingLocationType() HKWorkoutSwimmingLocationType
	SetSwimmingLocationType(value HKWorkoutSwimmingLocationType)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKWorkoutConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKWorkoutConfiguration */
// Alloc allocates a new instance without initialization.
func (hc _HKWorkoutConfigurationClass) Alloc() HKWorkoutConfiguration {
	rv := objc.Send[HKWorkoutConfiguration](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKWorkoutConfigurationClass) New() HKWorkoutConfiguration {
	rv := objc.Send[HKWorkoutConfiguration](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKWorkoutConfiguration) Init() HKWorkoutConfiguration {
	rv := objc.Send[HKWorkoutConfiguration](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKWorkoutConfiguration) Autorelease() HKWorkoutConfiguration {
	rv := objc.Send[HKWorkoutConfiguration](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKWorkoutConfiguration creates a new HKWorkoutConfiguration instance.
func NewHKWorkoutConfiguration() HKWorkoutConfiguration {
	return getHKWorkoutConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKWorkoutConfiguration */
// An object that contains configuration information about a workout session.
//
// Like many HealthKit classes, the class is not extendable and should not be subclassed.


// An object that contains configuration information about a workout session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutConfiguration
type HKWorkoutConfiguration struct {
	objectivec.Object
}

// HKWorkoutConfigurationFrom constructs a [HKWorkoutConfiguration] from an unsafe.Pointer.
//
// An object that contains configuration information about a workout session.
func HKWorkoutConfigurationFrom(ptr unsafe.Pointer) HKWorkoutConfiguration {
	return HKWorkoutConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKWorkoutConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKWorkoutConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKWorkoutConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKWorkoutConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKWorkoutConfiguration */

// The workout session’s activity type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutConfiguration/activityType
func (h_ HKWorkoutConfiguration) ActivityType() HKWorkoutActivityType {
	rv := objc.Send[HKWorkoutActivityType](h_.ID, objc.Sel("activityType"))
	return rv
}/* debug [instance_properties/getter]: activityType */


// The workout session’s activity type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutConfiguration/activityType
func (h_ HKWorkoutConfiguration) SetActivityType(value HKWorkoutActivityType) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setActivityType:"), value)
}/* debug [instance_properties/setter]: activityType */


// The length of the lap for a workout session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutConfiguration/lapLength
func (h_ HKWorkoutConfiguration) LapLength() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("lapLength"))
	return rv
}/* debug [instance_properties/getter]: lapLength */


// The length of the lap for a workout session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutConfiguration/lapLength
func (h_ HKWorkoutConfiguration) SetLapLength(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setLapLength:"), value)
}/* debug [instance_properties/setter]: lapLength */


// The workout session’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutConfiguration/locationType
func (h_ HKWorkoutConfiguration) LocationType() HKWorkoutSessionLocationType {
	rv := objc.Send[HKWorkoutSessionLocationType](h_.ID, objc.Sel("locationType"))
	return rv
}/* debug [instance_properties/getter]: locationType */


// The workout session’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutConfiguration/locationType
func (h_ HKWorkoutConfiguration) SetLocationType(value HKWorkoutSessionLocationType) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setLocationType:"), value)
}/* debug [instance_properties/setter]: locationType */


// The workout session’s swimming location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutConfiguration/swimmingLocationType
func (h_ HKWorkoutConfiguration) SwimmingLocationType() HKWorkoutSwimmingLocationType {
	rv := objc.Send[HKWorkoutSwimmingLocationType](h_.ID, objc.Sel("swimmingLocationType"))
	return rv
}/* debug [instance_properties/getter]: swimmingLocationType */


// The workout session’s swimming location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutConfiguration/swimmingLocationType
func (h_ HKWorkoutConfiguration) SetSwimmingLocationType(value HKWorkoutSwimmingLocationType) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSwimmingLocationType:"), value)
}/* debug [instance_properties/setter]: swimmingLocationType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKWorkoutConfiguration */



