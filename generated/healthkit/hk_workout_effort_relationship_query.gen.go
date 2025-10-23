// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKWorkoutEffortRelationshipQuery] class.
var (
	HKWorkoutEffortRelationshipQueryClass     _HKWorkoutEffortRelationshipQueryClass
	HKWorkoutEffortRelationshipQueryClassOnce sync.Once
)

func getHKWorkoutEffortRelationshipQueryClass() _HKWorkoutEffortRelationshipQueryClass {
	HKWorkoutEffortRelationshipQueryClassOnce.Do(func() {
		HKWorkoutEffortRelationshipQueryClass = _HKWorkoutEffortRelationshipQueryClass{objc.GetClass("HKWorkoutEffortRelationshipQuery")}
	})
	return HKWorkoutEffortRelationshipQueryClass
}

type _HKWorkoutEffortRelationshipQueryClass struct {
	class objc.Class
}

// An interface definition for the [HKWorkoutEffortRelationshipQuery] class.
type IHKWorkoutEffortRelationshipQuery interface {
	IHKQuery
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutEffortRelationshipQuery
type HKWorkoutEffortRelationshipQuery struct {
	HKQuery
}

// HKWorkoutEffortRelationshipQueryFrom constructs a [HKWorkoutEffortRelationshipQuery] from an unsafe.Pointer.
func HKWorkoutEffortRelationshipQueryFrom(ptr unsafe.Pointer) HKWorkoutEffortRelationshipQuery {
	return HKWorkoutEffortRelationshipQuery{
		HKQuery: HKQueryFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKWorkoutEffortRelationshipQueryClass) Alloc() HKWorkoutEffortRelationshipQuery {
	rv := objc.Send[HKWorkoutEffortRelationshipQuery](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKWorkoutEffortRelationshipQueryClass) New() HKWorkoutEffortRelationshipQuery {
	rv := objc.Send[HKWorkoutEffortRelationshipQuery](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKWorkoutEffortRelationshipQuery) Init() HKWorkoutEffortRelationshipQuery {
	rv := objc.Send[HKWorkoutEffortRelationshipQuery](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKWorkoutEffortRelationshipQuery) Autorelease() HKWorkoutEffortRelationshipQuery {
	rv := objc.Send[HKWorkoutEffortRelationshipQuery](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKWorkoutEffortRelationshipQuery creates a new HKWorkoutEffortRelationshipQuery instance.
func NewHKWorkoutEffortRelationshipQuery() HKWorkoutEffortRelationshipQuery {
	return getHKWorkoutEffortRelationshipQueryClass().New()
}




