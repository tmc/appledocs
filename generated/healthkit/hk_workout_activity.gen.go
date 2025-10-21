// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HKWorkoutActivity] class.
var (
	HKWorkoutActivityClass     _HKWorkoutActivityClass
	HKWorkoutActivityClassOnce sync.Once
)

func getHKWorkoutActivityClass() _HKWorkoutActivityClass {
	HKWorkoutActivityClassOnce.Do(func() {
		HKWorkoutActivityClass = _HKWorkoutActivityClass{objc.GetClass("HKWorkoutActivity")}
	})
	return HKWorkoutActivityClass
}

type _HKWorkoutActivityClass struct {
	class objc.Class
}

// An interface definition for the [HKWorkoutActivity] class.
type IHKWorkoutActivity interface {
	objectivec.IObject
}

// An object that describes an activity within a longer workout.
//
// Workout activity objects partition a workout into a set of separate activities. For example, you can use workout activities to record the swim, bike, and running portions of a multisport event, like a triathlon, or to represent the active and rest periods during interval training. All instance have at least one, associated . If you don’t explicitly set workout activities, HealthKit assigns a workout activity that matches the object’s activity type. For more information, see .
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivity
type HKWorkoutActivity struct {
	objectivec.Object
}

// HKWorkoutActivityFrom constructs a [HKWorkoutActivity] from an unsafe.Pointer.
//
// An object that describes an activity within a longer workout.
func HKWorkoutActivityFrom(ptr unsafe.Pointer) HKWorkoutActivity {
	return HKWorkoutActivity{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKWorkoutActivityClass) Alloc() HKWorkoutActivity {
	rv := objc.Send[HKWorkoutActivity](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKWorkoutActivityClass) New() HKWorkoutActivity {
	rv := objc.Send[HKWorkoutActivity](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKWorkoutActivity) Init() HKWorkoutActivity {
	rv := objc.Send[HKWorkoutActivity](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKWorkoutActivity) Autorelease() HKWorkoutActivity {
	rv := objc.Send[HKWorkoutActivity](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKWorkoutActivity creates a new HKWorkoutActivity instance.
func NewHKWorkoutActivity() HKWorkoutActivity {
	return getHKWorkoutActivityClass().New()
}


// The activity’s duration, measured in seconds.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivity/duration
func (h_ HKWorkoutActivity) Duration() TimeInterval {
	rv := objc.Send[TimeInterval](h_.ID, objc.Sel("duration"))
	return rv
}



