// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKWorkoutRouteQuery] class.
var (
	HKWorkoutRouteQueryClass     _HKWorkoutRouteQueryClass
	HKWorkoutRouteQueryClassOnce sync.Once
)

func getHKWorkoutRouteQueryClass() _HKWorkoutRouteQueryClass {
	HKWorkoutRouteQueryClassOnce.Do(func() {
		HKWorkoutRouteQueryClass = _HKWorkoutRouteQueryClass{objc.GetClass("HKWorkoutRouteQuery")}
	})
	return HKWorkoutRouteQueryClass
}

type _HKWorkoutRouteQueryClass struct {
	class objc.Class
}

// An interface definition for the [HKWorkoutRouteQuery] class.
type IHKWorkoutRouteQuery interface {
	IHKQuery
}

// A query to access the location data stored in a workout route.
//
// Use a workout route query to access the location data associated with an . Because a route sample can include a large number of objects, the query asynchronously returns the locations in batches. For detailed instructions, see .


// A query to access the location data stored in a workout route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutRouteQuery

type HKWorkoutRouteQuery struct {
	HKQuery
}

// HKWorkoutRouteQueryFrom constructs a [HKWorkoutRouteQuery] from an unsafe.Pointer.
//
// A query to access the location data stored in a workout route.
func HKWorkoutRouteQueryFrom(ptr unsafe.Pointer) HKWorkoutRouteQuery {
	return HKWorkoutRouteQuery{
		HKQuery: HKQueryFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKWorkoutRouteQueryClass) Alloc() HKWorkoutRouteQuery {
	rv := objc.Send[HKWorkoutRouteQuery](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKWorkoutRouteQueryClass) New() HKWorkoutRouteQuery {
	rv := objc.Send[HKWorkoutRouteQuery](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKWorkoutRouteQuery) Init() HKWorkoutRouteQuery {
	rv := objc.Send[HKWorkoutRouteQuery](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKWorkoutRouteQuery) Autorelease() HKWorkoutRouteQuery {
	rv := objc.Send[HKWorkoutRouteQuery](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKWorkoutRouteQuery creates a new HKWorkoutRouteQuery instance.
func NewHKWorkoutRouteQuery() HKWorkoutRouteQuery {
	return getHKWorkoutRouteQueryClass().New()
}




