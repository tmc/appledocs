// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKWorkoutRouteQuery */


/* debug [class_header]: Header for HKWorkoutRouteQuery */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKWorkoutRouteQuery */
// An interface definition for the [HKWorkoutRouteQuery] class.
type IHKWorkoutRouteQuery interface {
	IHKQuery
	
/* debug [class_interface_properties]: Properties for HKWorkoutRouteQuery */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKWorkoutRouteQuery */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKWorkoutRouteQuery */
// Alloc allocates a new instance without initialization.
func (hc _HKWorkoutRouteQueryClass) Alloc() HKWorkoutRouteQuery {
	rv := objc.Send[HKWorkoutRouteQuery](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKWorkoutRouteQuery */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKWorkoutRouteQuery */

// Creates a new query to access the location data associated with a workout route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutRouteQuery/init(route:dataHandler:)
func NewHKWorkoutRouteQueryWithRouteDataHandler(workoutRoute IHKWorkoutRoute, dataHandler unsafe.Pointer) HKWorkoutRouteQuery {
	instance := getHKWorkoutRouteQueryClass().Alloc()
	rv := objc.Send[HKWorkoutRouteQuery](instance.ID, objc.Sel("initWithRoute:dataHandler:"), workoutRoute, dataHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKWorkoutRouteQueryWithRouteDataHandler */


// Creates a new query to access the location data associated with a workout route during the specified date interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutRouteQuery/init(route:dateInterval:dataHandler:)
func NewHKWorkoutRouteQueryWithRouteDateIntervalDataHandler(workoutRoute IHKWorkoutRoute, dateInterval foundation.DateInterval, dataHandler unsafe.Pointer) HKWorkoutRouteQuery {
	instance := getHKWorkoutRouteQueryClass().Alloc()
	rv := objc.Send[HKWorkoutRouteQuery](instance.ID, objc.Sel("initWithRoute:dateInterval:dataHandler:"), workoutRoute, dateInterval, dataHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKWorkoutRouteQueryWithRouteDateIntervalDataHandler */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKWorkoutRouteQuery */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKWorkoutRouteQuery */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKWorkoutRouteQuery */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKWorkoutRouteQuery */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKWorkoutRouteQuery */


