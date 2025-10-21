// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKWorkoutRoute] class.
var (
	HKWorkoutRouteClass     _HKWorkoutRouteClass
	HKWorkoutRouteClassOnce sync.Once
)

func getHKWorkoutRouteClass() _HKWorkoutRouteClass {
	HKWorkoutRouteClassOnce.Do(func() {
		HKWorkoutRouteClass = _HKWorkoutRouteClass{objc.GetClass("HKWorkoutRoute")}
	})
	return HKWorkoutRouteClass
}

type _HKWorkoutRouteClass struct {
	class objc.Class
}

// An interface definition for the [HKWorkoutRoute] class.
type IHKWorkoutRoute interface {
	IHKSeriesSample
}

// A sample that contains a workout’s route data.
//
// When creating a workout route, you do not instantiate the objects directly. Instead, create a object, and provide it with location data throughout the workout. After the workout ends, call the route builder’s method to create the route. For detailed instructions, see . The route’s location data is stored as an array of objects. Because the route may contain a large number of location objects, use a object to asynchronously read the location data from the HealthKit store in batches. For more information, see .
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutRoute
type HKWorkoutRoute struct {
	HKSeriesSample
}

// HKWorkoutRouteFrom constructs a [HKWorkoutRoute] from an unsafe.Pointer.
//
// A sample that contains a workout’s route data.
func HKWorkoutRouteFrom(ptr unsafe.Pointer) HKWorkoutRoute {
	return HKWorkoutRoute{
		HKSeriesSample: HKSeriesSampleFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKWorkoutRouteClass) Alloc() HKWorkoutRoute {
	rv := objc.Send[HKWorkoutRoute](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKWorkoutRouteClass) New() HKWorkoutRoute {
	rv := objc.Send[HKWorkoutRoute](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKWorkoutRoute) Init() HKWorkoutRoute {
	rv := objc.Send[HKWorkoutRoute](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKWorkoutRoute) Autorelease() HKWorkoutRoute {
	rv := objc.Send[HKWorkoutRoute](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKWorkoutRoute creates a new HKWorkoutRoute instance.
func NewHKWorkoutRoute() HKWorkoutRoute {
	return getHKWorkoutRouteClass().New()
}


// A series sample containing location data that defines the route the user took during a workout.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutroutetypeidentifier
func (h_ HKWorkoutRoute) HKWorkoutRouteTypeIdentifier() string {
	rv := objc.Send[string](h_.ID, objc.Sel("HKWorkoutRouteTypeIdentifier"))
	return rv
}



