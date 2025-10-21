// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
}

//
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




