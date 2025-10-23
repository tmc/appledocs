// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKWorkoutType] class.
var (
	HKWorkoutTypeClass     _HKWorkoutTypeClass
	HKWorkoutTypeClassOnce sync.Once
)

func getHKWorkoutTypeClass() _HKWorkoutTypeClass {
	HKWorkoutTypeClassOnce.Do(func() {
		HKWorkoutTypeClass = _HKWorkoutTypeClass{objc.GetClass("HKWorkoutType")}
	})
	return HKWorkoutTypeClass
}

type _HKWorkoutTypeClass struct {
	class objc.Class
}

// An interface definition for the [HKWorkoutType] class.
type IHKWorkoutType interface {
	IHKSampleType
}

// A type that identifies samples that store information about a workout.
//
// The class is a concrete subclass of the class. To create a workout type instances, use the convenience method. All workouts use the same workout type instance.


// A type that identifies samples that store information about a workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutType
type HKWorkoutType struct {
	HKSampleType
}

// HKWorkoutTypeFrom constructs a [HKWorkoutType] from an unsafe.Pointer.
//
// A type that identifies samples that store information about a workout.
func HKWorkoutTypeFrom(ptr unsafe.Pointer) HKWorkoutType {
	return HKWorkoutType{
		HKSampleType: HKSampleTypeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKWorkoutTypeClass) Alloc() HKWorkoutType {
	rv := objc.Send[HKWorkoutType](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKWorkoutTypeClass) New() HKWorkoutType {
	rv := objc.Send[HKWorkoutType](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKWorkoutType) Init() HKWorkoutType {
	rv := objc.Send[HKWorkoutType](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKWorkoutType) Autorelease() HKWorkoutType {
	rv := objc.Send[HKWorkoutType](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKWorkoutType creates a new HKWorkoutType instance.
func NewHKWorkoutType() HKWorkoutType {
	return getHKWorkoutTypeClass().New()
}





