// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKSampleQuery */


/* debug [class_header]: Header for HKSampleQuery */
// The class instance for the [HKSampleQuery] class.
var (
	HKSampleQueryClass     _HKSampleQueryClass
	HKSampleQueryClassOnce sync.Once
)

func getHKSampleQueryClass() _HKSampleQueryClass {
	HKSampleQueryClassOnce.Do(func() {
		HKSampleQueryClass = _HKSampleQueryClass{objc.GetClass("HKSampleQuery")}
	})
	return HKSampleQueryClass
}

type _HKSampleQueryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKSampleQuery */
// An interface definition for the [HKSampleQuery] class.
type IHKSampleQuery interface {
	IHKQuery
	
/* debug [class_interface_properties]: Properties for HKSampleQuery */
	// properties:
	Limit() uint
	SortDescriptors() []cloudkit.SortDescriptor
	HKObjectQueryNoLimit() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKSampleQuery */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKSampleQuery */
// Alloc allocates a new instance without initialization.
func (hc _HKSampleQueryClass) Alloc() HKSampleQuery {
	rv := objc.Send[HKSampleQuery](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKSampleQueryClass) New() HKSampleQuery {
	rv := objc.Send[HKSampleQuery](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKSampleQuery) Init() HKSampleQuery {
	rv := objc.Send[HKSampleQuery](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKSampleQuery) Autorelease() HKSampleQuery {
	rv := objc.Send[HKSampleQuery](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKSampleQuery creates a new HKSampleQuery instance.
func NewHKSampleQuery() HKSampleQuery {
	return getHKSampleQueryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKSampleQuery */
// A general query that returns a snapshot of all the matching samples currently saved in the HealthKit store.
//
// You can use sample queries to search for any concrete subclasses of the class, including , , , and objects. The sample query returns sample objects that match the provided type and predicate. You can provide a sort order for the returned samples, or limit the number of samples returned. Other query classes can be used to perform more specialized searches and calculations. For more information, see . Sample queries are immutable: The query’s properties are set when the query is first created, and they can’t change.


// A general query that returns a snapshot of all the matching samples currently saved in the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSampleQuery
type HKSampleQuery struct {
	HKQuery
}

// HKSampleQueryFrom constructs a [HKSampleQuery] from an unsafe.Pointer.
//
// A general query that returns a snapshot of all the matching samples currently saved in the HealthKit store.
func HKSampleQueryFrom(ptr unsafe.Pointer) HKSampleQuery {
	return HKSampleQuery{
		HKQuery: HKQueryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKSampleQuery */

// Creates a query for samples that match any of the descriptors you provided.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSampleQuery/init(queryDescriptors:limit:resultsHandler:)
func NewHKSampleQueryWithQueryDescriptorsLimitResultsHandler(queryDescriptors []HKQueryDescriptor, limit int, resultsHandler unsafe.Pointer) HKSampleQuery {
	instance := getHKSampleQueryClass().Alloc()
	rv := objc.Send[HKSampleQuery](instance.ID, objc.Sel("initWithQueryDescriptors:limit:resultsHandler:"), queryDescriptors, limit, resultsHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKSampleQueryWithQueryDescriptorsLimitResultsHandler */


// Creates a query for samples that match any of the query descriptors you provided, sorted by the sort descriptors you provided.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSampleQuery/init(queryDescriptors:limit:sortDescriptors:resultsHandler:)
func NewHKSampleQueryWithQueryDescriptorsLimitSortDescriptorsResultsHandler(queryDescriptors []HKQueryDescriptor, limit int, sortDescriptors []cloudkit.SortDescriptor, resultsHandler unsafe.Pointer) HKSampleQuery {
	instance := getHKSampleQueryClass().Alloc()
	rv := objc.Send[HKSampleQuery](instance.ID, objc.Sel("initWithQueryDescriptors:limit:sortDescriptors:resultsHandler:"), queryDescriptors, limit, sortDescriptors, resultsHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKSampleQueryWithQueryDescriptorsLimitSortDescriptorsResultsHandler */


// Instantiates and returns a sample query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSampleQuery/init(sampleType:predicate:limit:sortDescriptors:resultsHandler:)
func NewHKSampleQueryWithSampleTypePredicateLimitSortDescriptorsResultsHandler(sampleType IHKSampleType, predicate foundation.Predicate, limit uint, sortDescriptors []cloudkit.SortDescriptor, resultsHandler unsafe.Pointer) HKSampleQuery {
	instance := getHKSampleQueryClass().Alloc()
	rv := objc.Send[HKSampleQuery](instance.ID, objc.Sel("initWithSampleType:predicate:limit:sortDescriptors:resultsHandler:"), sampleType, predicate, limit, sortDescriptors, resultsHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKSampleQueryWithSampleTypePredicateLimitSortDescriptorsResultsHandler */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKSampleQuery */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKSampleQuery */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKSampleQuery */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKSampleQuery */

// The maximum number of samples that this query returns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSampleQuery/limit
func (h_ HKSampleQuery) Limit() uint {
	rv := objc.Send[uint](h_.ID, objc.Sel("limit"))
	return rv
}/* debug [instance_properties/getter]: limit */


// The sort descriptors that specify the order of the results returned by this query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSampleQuery/sortDescriptors
func (h_ HKSampleQuery) SortDescriptors() []cloudkit.SortDescriptor {
	rv := objc.Send[[]cloudkit.SortDescriptor](h_.ID, objc.Sel("sortDescriptors"))
	return rv
}/* debug [instance_properties/getter]: sortDescriptors */


// A value indicating that the query returns all the matching samples in the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkobjectquerynolimit
func (h_ HKSampleQuery) HKObjectQueryNoLimit() int {
	rv := objc.Send[int](h_.ID, objc.Sel("HKObjectQueryNoLimit"))
	return rv
}/* debug [instance_properties/getter]: HKObjectQueryNoLimit */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKSampleQuery */


