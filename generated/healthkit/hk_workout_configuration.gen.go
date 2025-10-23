// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [HKWorkoutConfiguration] class.
type IHKWorkoutConfiguration interface {
	objectivec.IObject
	// properties:
	ActivityType() HKWorkoutActivityType
	SetActivityType(value HKWorkoutActivityType)
	LapLength() IHKQuantity
	SetLapLength(value IHKQuantity)
	LocationType() unsafe.Pointer
	SetLocationType(value unsafe.Pointer)
	SwimmingLocationType() unsafe.Pointer
	SetSwimmingLocationType(value unsafe.Pointer)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (hc _HKWorkoutConfigurationClass) Alloc() HKWorkoutConfiguration {
	rv := objc.Send[HKWorkoutConfiguration](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The workout session’s activity type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutconfiguration/activitytype
func (h_ HKWorkoutConfiguration) ActivityType() HKWorkoutActivityType {
	rv := objc.Send[HKWorkoutActivityType](h_.ID, objc.Sel("activityType"))
	return rv
}


// The workout session’s activity type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutconfiguration/activitytype
func (h_ HKWorkoutConfiguration) SetActivityType(value HKWorkoutActivityType) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setActivityType:"), value)
}


// The length of the lap for a workout session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutconfiguration/laplength
func (h_ HKWorkoutConfiguration) LapLength() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("lapLength"))
	return rv
}


// The length of the lap for a workout session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutconfiguration/laplength
func (h_ HKWorkoutConfiguration) SetLapLength(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setLapLength:"), value)
}


// The workout session’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutconfiguration/locationtype
func (h_ HKWorkoutConfiguration) LocationType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("locationType"))
	return rv
}


// The workout session’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutconfiguration/locationtype
func (h_ HKWorkoutConfiguration) SetLocationType(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setLocationType:"), value)
}


// The workout session’s swimming location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutconfiguration/swimminglocationtype
func (h_ HKWorkoutConfiguration) SwimmingLocationType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("swimmingLocationType"))
	return rv
}


// The workout session’s swimming location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutconfiguration/swimminglocationtype
func (h_ HKWorkoutConfiguration) SetSwimmingLocationType(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSwimmingLocationType:"), value)
}



