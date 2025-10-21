// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
}

// An object that contains configuration information about a workout session.
//
// Like many HealthKit classes, the class is not extendable and should not be subclassed.
//
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
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutConfiguration/activityType
func (h_ HKWorkoutConfiguration) ActivityType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("activityType"))
	return rv
}


// SetActivityType sets the value of the activityType property.
// The workout session’s activity type.

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutConfiguration/activityType
func (h_ HKWorkoutConfiguration) SetActivityType(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setActivityType:"), value)
}

// The length of the lap for a workout session.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutConfiguration/lapLength
func (h_ HKWorkoutConfiguration) LapLength() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("lapLength"))
	return rv
}


// SetLapLength sets the value of the lapLength property.
// The length of the lap for a workout session.

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutConfiguration/lapLength
func (h_ HKWorkoutConfiguration) SetLapLength(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setLapLength:"), value)
}

// The workout session’s location.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutConfiguration/locationType
func (h_ HKWorkoutConfiguration) LocationType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("locationType"))
	return rv
}


// SetLocationType sets the value of the locationType property.
// The workout session’s location.

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutConfiguration/locationType
func (h_ HKWorkoutConfiguration) SetLocationType(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setLocationType:"), value)
}

// The workout session’s swimming location.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutConfiguration/swimmingLocationType
func (h_ HKWorkoutConfiguration) SwimmingLocationType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("swimmingLocationType"))
	return rv
}


// SetSwimmingLocationType sets the value of the swimmingLocationType property.
// The workout session’s swimming location.

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutConfiguration/swimmingLocationType
func (h_ HKWorkoutConfiguration) SetSwimmingLocationType(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSwimmingLocationType:"), value)
}



