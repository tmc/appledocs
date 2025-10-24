// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKActivitySummaryQuery */


/* debug [class_header]: Header for HKActivitySummaryQuery */
// The class instance for the [HKActivitySummaryQuery] class.
var (
	HKActivitySummaryQueryClass     _HKActivitySummaryQueryClass
	HKActivitySummaryQueryClassOnce sync.Once
)

func getHKActivitySummaryQueryClass() _HKActivitySummaryQueryClass {
	HKActivitySummaryQueryClassOnce.Do(func() {
		HKActivitySummaryQueryClass = _HKActivitySummaryQueryClass{objc.GetClass("HKActivitySummaryQuery")}
	})
	return HKActivitySummaryQueryClass
}

type _HKActivitySummaryQueryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKActivitySummaryQuery */
// An interface definition for the [HKActivitySummaryQuery] class.
type IHKActivitySummaryQuery interface {
	IHKQuery
	
/* debug [class_interface_properties]: Properties for HKActivitySummaryQuery */
	// properties:
	UpdateHandler() unsafe.Pointer
	SetUpdateHandler(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKActivitySummaryQuery */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKActivitySummaryQuery */
// Alloc allocates a new instance without initialization.
func (hc _HKActivitySummaryQueryClass) Alloc() HKActivitySummaryQuery {
	rv := objc.Send[HKActivitySummaryQuery](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKActivitySummaryQueryClass) New() HKActivitySummaryQuery {
	rv := objc.Send[HKActivitySummaryQuery](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKActivitySummaryQuery) Init() HKActivitySummaryQuery {
	rv := objc.Send[HKActivitySummaryQuery](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKActivitySummaryQuery) Autorelease() HKActivitySummaryQuery {
	rv := objc.Send[HKActivitySummaryQuery](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKActivitySummaryQuery creates a new HKActivitySummaryQuery instance.
func NewHKActivitySummaryQuery() HKActivitySummaryQuery {
	return getHKActivitySummaryQueryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKActivitySummaryQuery */
// A query for reading activity summary objects from the HealthKit store.
//
// Activity summary query objects are mostly immutable. You can assign the query’s property after instantiating the object, but before executing the query. All other properties must be set when you instantiate the object, and they can’t change.


// A query for reading activity summary objects from the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummaryQuery
type HKActivitySummaryQuery struct {
	HKQuery
}

// HKActivitySummaryQueryFrom constructs a [HKActivitySummaryQuery] from an unsafe.Pointer.
//
// A query for reading activity summary objects from the HealthKit store.
func HKActivitySummaryQueryFrom(ptr unsafe.Pointer) HKActivitySummaryQuery {
	return HKActivitySummaryQuery{
		HKQuery: HKQueryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKActivitySummaryQuery */

// Initializes a new active summary query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummaryQuery/init(predicate:resultsHandler:)
func NewHKActivitySummaryQueryWithPredicateResultsHandler(predicate foundation.Predicate, handler unsafe.Pointer) HKActivitySummaryQuery {
	instance := getHKActivitySummaryQueryClass().Alloc()
	rv := objc.Send[HKActivitySummaryQuery](instance.ID, objc.Sel("initWithPredicate:resultsHandler:"), predicate, handler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKActivitySummaryQueryWithPredicateResultsHandler */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKActivitySummaryQuery */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKActivitySummaryQuery */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKActivitySummaryQuery */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKActivitySummaryQuery */

// The handler for monitoring updates to activity summaries saved in the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummaryQuery/updateHandler
func (h_ HKActivitySummaryQuery) UpdateHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("updateHandler"))
	return rv
}/* debug [instance_properties/getter]: updateHandler */


// The handler for monitoring updates to activity summaries saved in the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummaryQuery/updateHandler
func (h_ HKActivitySummaryQuery) SetUpdateHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setUpdateHandler:"), value)
}/* debug [instance_properties/setter]: updateHandler */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKActivitySummaryQuery */


