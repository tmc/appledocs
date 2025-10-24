// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKAnchoredObjectQuery */


/* debug [class_header]: Header for HKAnchoredObjectQuery */
// The class instance for the [HKAnchoredObjectQuery] class.
var (
	HKAnchoredObjectQueryClass     _HKAnchoredObjectQueryClass
	HKAnchoredObjectQueryClassOnce sync.Once
)

func getHKAnchoredObjectQueryClass() _HKAnchoredObjectQueryClass {
	HKAnchoredObjectQueryClassOnce.Do(func() {
		HKAnchoredObjectQueryClass = _HKAnchoredObjectQueryClass{objc.GetClass("HKAnchoredObjectQuery")}
	})
	return HKAnchoredObjectQueryClass
}

type _HKAnchoredObjectQueryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKAnchoredObjectQuery */
// An interface definition for the [HKAnchoredObjectQuery] class.
type IHKAnchoredObjectQuery interface {
	IHKQuery
	
/* debug [class_interface_properties]: Properties for HKAnchoredObjectQuery */
	// properties:
	UpdateHandler() unsafe.Pointer
	SetUpdateHandler(value unsafe.Pointer)
	HKObjectQueryNoLimit() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKAnchoredObjectQuery */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKAnchoredObjectQuery */
// Alloc allocates a new instance without initialization.
func (hc _HKAnchoredObjectQueryClass) Alloc() HKAnchoredObjectQuery {
	rv := objc.Send[HKAnchoredObjectQuery](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKAnchoredObjectQueryClass) New() HKAnchoredObjectQuery {
	rv := objc.Send[HKAnchoredObjectQuery](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKAnchoredObjectQuery) Init() HKAnchoredObjectQuery {
	rv := objc.Send[HKAnchoredObjectQuery](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKAnchoredObjectQuery) Autorelease() HKAnchoredObjectQuery {
	rv := objc.Send[HKAnchoredObjectQuery](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKAnchoredObjectQuery creates a new HKAnchoredObjectQuery instance.
func NewHKAnchoredObjectQuery() HKAnchoredObjectQuery {
	return getHKAnchoredObjectQueryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKAnchoredObjectQuery */
// A query that returns changes to the HealthKit store, including a snapshot of new changes and continuous monitoring as a long-running query.
//
// Anchored object queries provide an easy way to search for new data in the HealthKit store. An returns an anchor value that corresponds to the last sample or deleted object received by that query. Subsequent queries can use this anchor to restrict their results to only newer saved or deleted objects. Anchored object queries are mostly immutable. You can assign the query’s property after instantiating the object, but you must set all other properties when you instantiate the object. You can’t change them.


// A query that returns changes to the HealthKit store, including a snapshot of new changes and continuous monitoring as a long-running query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAnchoredObjectQuery
type HKAnchoredObjectQuery struct {
	HKQuery
}

// HKAnchoredObjectQueryFrom constructs a [HKAnchoredObjectQuery] from an unsafe.Pointer.
//
// A query that returns changes to the HealthKit store, including a snapshot of new changes and continuous monitoring as a long-running query.
func HKAnchoredObjectQueryFrom(ptr unsafe.Pointer) HKAnchoredObjectQuery {
	return HKAnchoredObjectQuery{
		HKQuery: HKQueryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKAnchoredObjectQuery */

// Creates an anchored object query that matches any of the query descriptors you provided.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAnchoredObjectQuery/init(queryDescriptors:anchor:limit:resultsHandler:)
func NewHKAnchoredObjectQueryWithQueryDescriptorsAnchorLimitResultsHandler(queryDescriptors []HKQueryDescriptor, anchor IHKQueryAnchor, limit int, handler unsafe.Pointer) HKAnchoredObjectQuery {
	instance := getHKAnchoredObjectQueryClass().Alloc()
	rv := objc.Send[HKAnchoredObjectQuery](instance.ID, objc.Sel("initWithQueryDescriptors:anchor:limit:resultsHandler:"), queryDescriptors, anchor, limit, handler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKAnchoredObjectQueryWithQueryDescriptorsAnchorLimitResultsHandler */


// Initializes a new anchored object query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAnchoredObjectQuery/init(type:predicate:anchor:limit:completionHandler:)
func NewHKAnchoredObjectQueryWithTypePredicateAnchorLimitCompletionHandler(type_ IHKSampleType, predicate foundation.Predicate, anchor uint, limit uint, handler unsafe.Pointer) HKAnchoredObjectQuery {
	instance := getHKAnchoredObjectQueryClass().Alloc()
	rv := objc.Send[HKAnchoredObjectQuery](instance.ID, objc.Sel("initWithType:predicate:anchor:limit:completionHandler:"), type_, predicate, anchor, limit, handler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKAnchoredObjectQueryWithTypePredicateAnchorLimitCompletionHandler */


// Initializes a new anchored object query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAnchoredObjectQuery/init(type:predicate:anchor:limit:resultsHandler:)
func NewHKAnchoredObjectQueryWithTypePredicateAnchorLimitResultsHandler(type_ IHKSampleType, predicate foundation.Predicate, anchor IHKQueryAnchor, limit uint, handler unsafe.Pointer) HKAnchoredObjectQuery {
	instance := getHKAnchoredObjectQueryClass().Alloc()
	rv := objc.Send[HKAnchoredObjectQuery](instance.ID, objc.Sel("initWithType:predicate:anchor:limit:resultsHandler:"), type_, predicate, anchor, limit, handler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKAnchoredObjectQueryWithTypePredicateAnchorLimitResultsHandler */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKAnchoredObjectQuery */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKAnchoredObjectQuery */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKAnchoredObjectQuery */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKAnchoredObjectQuery */

// Handler for monitoring updates to the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAnchoredObjectQuery/updateHandler
func (h_ HKAnchoredObjectQuery) UpdateHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("updateHandler"))
	return rv
}/* debug [instance_properties/getter]: updateHandler */


// Handler for monitoring updates to the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAnchoredObjectQuery/updateHandler
func (h_ HKAnchoredObjectQuery) SetUpdateHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setUpdateHandler:"), value)
}/* debug [instance_properties/setter]: updateHandler */


// A value indicating that the query returns all the matching samples in the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkobjectquerynolimit
func (h_ HKAnchoredObjectQuery) HKObjectQueryNoLimit() int {
	rv := objc.Send[int](h_.ID, objc.Sel("HKObjectQueryNoLimit"))
	return rv
}/* debug [instance_properties/getter]: HKObjectQueryNoLimit */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKAnchoredObjectQuery */


