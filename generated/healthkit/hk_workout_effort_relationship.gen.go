// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HKWorkoutEffortRelationship] class.
var (
	HKWorkoutEffortRelationshipClass     _HKWorkoutEffortRelationshipClass
	HKWorkoutEffortRelationshipClassOnce sync.Once
)

func getHKWorkoutEffortRelationshipClass() _HKWorkoutEffortRelationshipClass {
	HKWorkoutEffortRelationshipClassOnce.Do(func() {
		HKWorkoutEffortRelationshipClass = _HKWorkoutEffortRelationshipClass{objc.GetClass("HKWorkoutEffortRelationship")}
	})
	return HKWorkoutEffortRelationshipClass
}

type _HKWorkoutEffortRelationshipClass struct {
	class objc.Class
}

// An interface definition for the [HKWorkoutEffortRelationship] class.
type IHKWorkoutEffortRelationship interface {
	objectivec.IObject
	Activity() HKWorkoutActivity
	SetActivity(value IHKWorkoutActivity)
	Samples() HKSample
	SetSamples(value IHKSample)
	Workout() HKWorkout
	SetWorkout(value IHKWorkout)
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutEffortRelationship

type HKWorkoutEffortRelationship struct {
	objectivec.Object
}

// HKWorkoutEffortRelationshipFrom constructs a [HKWorkoutEffortRelationship] from an unsafe.Pointer.
func HKWorkoutEffortRelationshipFrom(ptr unsafe.Pointer) HKWorkoutEffortRelationship {
	return HKWorkoutEffortRelationship{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKWorkoutEffortRelationshipClass) Alloc() HKWorkoutEffortRelationship {
	rv := objc.Send[HKWorkoutEffortRelationship](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKWorkoutEffortRelationshipClass) New() HKWorkoutEffortRelationship {
	rv := objc.Send[HKWorkoutEffortRelationship](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKWorkoutEffortRelationship) Init() HKWorkoutEffortRelationship {
	rv := objc.Send[HKWorkoutEffortRelationship](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKWorkoutEffortRelationship) Autorelease() HKWorkoutEffortRelationship {
	rv := objc.Send[HKWorkoutEffortRelationship](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKWorkoutEffortRelationship creates a new HKWorkoutEffortRelationship instance.
func NewHKWorkoutEffortRelationship() HKWorkoutEffortRelationship {
	return getHKWorkoutEffortRelationshipClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkouteffortrelationship/activity

func (h_ HKWorkoutEffortRelationship) Activity() HKWorkoutActivity {
	rv := objc.Send[HKWorkoutActivity](h_.ID, objc.Sel("activity"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkouteffortrelationship/activity

func (h_ HKWorkoutEffortRelationship) SetActivity(value IHKWorkoutActivity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setActivity:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkouteffortrelationship/samples

func (h_ HKWorkoutEffortRelationship) Samples() HKSample {
	rv := objc.Send[HKSample](h_.ID, objc.Sel("samples"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkouteffortrelationship/samples

func (h_ HKWorkoutEffortRelationship) SetSamples(value IHKSample) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSamples:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkouteffortrelationship/workout

func (h_ HKWorkoutEffortRelationship) Workout() HKWorkout {
	rv := objc.Send[HKWorkout](h_.ID, objc.Sel("workout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkouteffortrelationship/workout

func (h_ HKWorkoutEffortRelationship) SetWorkout(value IHKWorkout) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setWorkout:"), value)
}



