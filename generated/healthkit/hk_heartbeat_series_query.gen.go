// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKHeartbeatSeriesQuery */


/* debug [class_header]: Header for HKHeartbeatSeriesQuery */
// The class instance for the [HKHeartbeatSeriesQuery] class.
var (
	HKHeartbeatSeriesQueryClass     _HKHeartbeatSeriesQueryClass
	HKHeartbeatSeriesQueryClassOnce sync.Once
)

func getHKHeartbeatSeriesQueryClass() _HKHeartbeatSeriesQueryClass {
	HKHeartbeatSeriesQueryClassOnce.Do(func() {
		HKHeartbeatSeriesQueryClass = _HKHeartbeatSeriesQueryClass{objc.GetClass("HKHeartbeatSeriesQuery")}
	})
	return HKHeartbeatSeriesQueryClass
}

type _HKHeartbeatSeriesQueryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKHeartbeatSeriesQuery */
// An interface definition for the [HKHeartbeatSeriesQuery] class.
type IHKHeartbeatSeriesQuery interface {
	IHKQuery
	
/* debug [class_interface_properties]: Properties for HKHeartbeatSeriesQuery */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKHeartbeatSeriesQuery */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKHeartbeatSeriesQuery */
// Alloc allocates a new instance without initialization.
func (hc _HKHeartbeatSeriesQueryClass) Alloc() HKHeartbeatSeriesQuery {
	rv := objc.Send[HKHeartbeatSeriesQuery](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKHeartbeatSeriesQueryClass) New() HKHeartbeatSeriesQuery {
	rv := objc.Send[HKHeartbeatSeriesQuery](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKHeartbeatSeriesQuery) Init() HKHeartbeatSeriesQuery {
	rv := objc.Send[HKHeartbeatSeriesQuery](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKHeartbeatSeriesQuery) Autorelease() HKHeartbeatSeriesQuery {
	rv := objc.Send[HKHeartbeatSeriesQuery](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKHeartbeatSeriesQuery creates a new HKHeartbeatSeriesQuery instance.
func NewHKHeartbeatSeriesQuery() HKHeartbeatSeriesQuery {
	return getHKHeartbeatSeriesQueryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKHeartbeatSeriesQuery */
// A query that returns the heartbeat data contained in a heartbeat series sample.


// A query that returns the heartbeat data contained in a heartbeat series sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHeartbeatSeriesQuery
type HKHeartbeatSeriesQuery struct {
	HKQuery
}

// HKHeartbeatSeriesQueryFrom constructs a [HKHeartbeatSeriesQuery] from an unsafe.Pointer.
//
// A query that returns the heartbeat data contained in a heartbeat series sample.
func HKHeartbeatSeriesQueryFrom(ptr unsafe.Pointer) HKHeartbeatSeriesQuery {
	return HKHeartbeatSeriesQuery{
		HKQuery: HKQueryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKHeartbeatSeriesQuery */

// Creates a new heartbeat series query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHeartbeatSeriesQuery/init(heartbeatSeries:dataHandler:)
func NewHKHeartbeatSeriesQueryWithHeartbeatSeriesDataHandler(heartbeatSeries IHKHeartbeatSeriesSample, dataHandler unsafe.Pointer) HKHeartbeatSeriesQuery {
	instance := getHKHeartbeatSeriesQueryClass().Alloc()
	rv := objc.Send[HKHeartbeatSeriesQuery](instance.ID, objc.Sel("initWithHeartbeatSeries:dataHandler:"), heartbeatSeries, dataHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKHeartbeatSeriesQueryWithHeartbeatSeriesDataHandler */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKHeartbeatSeriesQuery */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKHeartbeatSeriesQuery */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKHeartbeatSeriesQuery */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKHeartbeatSeriesQuery */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKHeartbeatSeriesQuery */


