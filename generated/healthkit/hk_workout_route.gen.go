// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class HKWorkoutRoute */


/* debug [class_header]: Header for HKWorkoutRoute */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKWorkoutRoute */
// An interface definition for the [HKWorkoutRoute] class.
type IHKWorkoutRoute interface {
	IHKSeriesSample
	
/* debug [class_interface_properties]: Properties for HKWorkoutRoute */
	// properties:
	HKWorkoutRouteTypeIdentifier() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKWorkoutRoute */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKWorkoutRoute */
// Alloc allocates a new instance without initialization.
func (hc _HKWorkoutRouteClass) Alloc() HKWorkoutRoute {
	rv := objc.Send[HKWorkoutRoute](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKWorkoutRoute */
// A sample that contains a workout’s route data.
//
// When creating a workout route, you do not instantiate the objects directly. Instead, create a object, and provide it with location data throughout the workout. After the workout ends, call the route builder’s method to create the route. For detailed instructions, see . The route’s location data is stored as an array of objects. Because the route may contain a large number of location objects, use a object to asynchronously read the location data from the HealthKit store in batches. For more information, see .


// A sample that contains a workout’s route data.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKWorkoutRoute *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKWorkoutRoute */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKWorkoutRoute */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKWorkoutRoute */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKWorkoutRoute */

// A series sample containing location data that defines the route the user took during a workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutroutetypeidentifier
func (h_ HKWorkoutRoute) HKWorkoutRouteTypeIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKWorkoutRouteTypeIdentifier"))
	return rv
}/* debug [instance_properties/getter]: HKWorkoutRouteTypeIdentifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKWorkoutRoute */



