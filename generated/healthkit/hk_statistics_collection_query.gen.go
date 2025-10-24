// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKStatisticsCollectionQuery */


/* debug [class_header]: Header for HKStatisticsCollectionQuery */
// The class instance for the [HKStatisticsCollectionQuery] class.
var (
	HKStatisticsCollectionQueryClass     _HKStatisticsCollectionQueryClass
	HKStatisticsCollectionQueryClassOnce sync.Once
)

func getHKStatisticsCollectionQueryClass() _HKStatisticsCollectionQueryClass {
	HKStatisticsCollectionQueryClassOnce.Do(func() {
		HKStatisticsCollectionQueryClass = _HKStatisticsCollectionQueryClass{objc.GetClass("HKStatisticsCollectionQuery")}
	})
	return HKStatisticsCollectionQueryClass
}

type _HKStatisticsCollectionQueryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKStatisticsCollectionQuery */
// An interface definition for the [HKStatisticsCollectionQuery] class.
type IHKStatisticsCollectionQuery interface {
	IHKQuery
	
/* debug [class_interface_properties]: Properties for HKStatisticsCollectionQuery */
	// properties:
	AnchorDate() objc.IObject /* cross-framework: NSDate */
	InitialResultsHandler() unsafe.Pointer
	SetInitialResultsHandler(value unsafe.Pointer)
	IntervalComponents() foundation.DateComponents
	Options() HKStatisticsOptions
	StatisticsUpdateHandler() unsafe.Pointer
	SetStatisticsUpdateHandler(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKStatisticsCollectionQuery */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKStatisticsCollectionQuery */
// Alloc allocates a new instance without initialization.
func (hc _HKStatisticsCollectionQueryClass) Alloc() HKStatisticsCollectionQuery {
	rv := objc.Send[HKStatisticsCollectionQuery](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKStatisticsCollectionQueryClass) New() HKStatisticsCollectionQuery {
	rv := objc.Send[HKStatisticsCollectionQuery](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKStatisticsCollectionQuery) Init() HKStatisticsCollectionQuery {
	rv := objc.Send[HKStatisticsCollectionQuery](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKStatisticsCollectionQuery) Autorelease() HKStatisticsCollectionQuery {
	rv := objc.Send[HKStatisticsCollectionQuery](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKStatisticsCollectionQuery creates a new HKStatisticsCollectionQuery instance.
func NewHKStatisticsCollectionQuery() HKStatisticsCollectionQuery {
	return getHKStatisticsCollectionQueryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKStatisticsCollectionQuery */
// A query that performs multiple statistics queries over a series of fixed-length time intervals.
//
// Statistics collection queries are often used to produce data for graphs and charts. For example, you might create a statistics collection query that calculates the total number of steps for each day or the average heart rate for each hour. Like observer queries, collection queries can also act as long-running queries, receiving updates when the HealthKit store’s content changes. Statistics collection queries are mostly immutable. You can assign the query’s and properties after instantiating the object. You must set all other properties when you instantiate the object, and they can’t change. For more information about statistics queries, see .


// A query that performs multiple statistics queries over a series of fixed-length time intervals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatisticsCollectionQuery
type HKStatisticsCollectionQuery struct {
	HKQuery
}

// HKStatisticsCollectionQueryFrom constructs a [HKStatisticsCollectionQuery] from an unsafe.Pointer.
//
// A query that performs multiple statistics queries over a series of fixed-length time intervals.
func HKStatisticsCollectionQueryFrom(ptr unsafe.Pointer) HKStatisticsCollectionQuery {
	return HKStatisticsCollectionQuery{
		HKQuery: HKQueryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKStatisticsCollectionQuery */

// Initializes a statistics collection query to perform the specified calculations over a set of time intervals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatisticsCollectionQuery/init(quantityType:quantitySamplePredicate:options:anchorDate:intervalComponents:)
func NewHKStatisticsCollectionQueryWithQuantityTypeQuantitySamplePredicateOptionsAnchorDateIntervalComponents(quantityType IHKQuantityType, quantitySamplePredicate foundation.Predicate, options HKStatisticsOptions, anchorDate objc.IObject /* cross-framework: NSDate */, intervalComponents foundation.DateComponents) HKStatisticsCollectionQuery {
	instance := getHKStatisticsCollectionQueryClass().Alloc()
	rv := objc.Send[HKStatisticsCollectionQuery](instance.ID, objc.Sel("initWithQuantityType:quantitySamplePredicate:options:anchorDate:intervalComponents:"), quantityType, quantitySamplePredicate, options, anchorDate, intervalComponents)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKStatisticsCollectionQueryWithQuantityTypeQuantitySamplePredicateOptionsAnchorDateIntervalComponents */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKStatisticsCollectionQuery */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKStatisticsCollectionQuery */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKStatisticsCollectionQuery */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKStatisticsCollectionQuery */

// The anchor date for the collection’s time intervals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatisticsCollectionQuery/anchorDate
func (h_ HKStatisticsCollectionQuery) AnchorDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](h_.ID, objc.Sel("anchorDate"))
	return rv
}/* debug [instance_properties/getter]: anchorDate */


// The results handler for the query’s initial results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatisticsCollectionQuery/initialResultsHandler
func (h_ HKStatisticsCollectionQuery) InitialResultsHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("initialResultsHandler"))
	return rv
}/* debug [instance_properties/getter]: initialResultsHandler */


// The results handler for the query’s initial results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatisticsCollectionQuery/initialResultsHandler
func (h_ HKStatisticsCollectionQuery) SetInitialResultsHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setInitialResultsHandler:"), value)
}/* debug [instance_properties/setter]: initialResultsHandler */


// The date components that define the time interval for each statistics object in the collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatisticsCollectionQuery/intervalComponents
func (h_ HKStatisticsCollectionQuery) IntervalComponents() foundation.DateComponents {
	rv := objc.Send[foundation.DateComponents](h_.ID, objc.Sel("intervalComponents"))
	return rv
}/* debug [instance_properties/getter]: intervalComponents */


// A list of options that define the type of statistical calculations performed and the way in which data from multiple sources are merged.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatisticsCollectionQuery/options
func (h_ HKStatisticsCollectionQuery) Options() HKStatisticsOptions {
	rv := objc.Send[HKStatisticsOptions](h_.ID, objc.Sel("options"))
	return rv
}/* debug [instance_properties/getter]: options */


// The results handler for monitoring updates to the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatisticsCollectionQuery/statisticsUpdateHandler
func (h_ HKStatisticsCollectionQuery) StatisticsUpdateHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("statisticsUpdateHandler"))
	return rv
}/* debug [instance_properties/getter]: statisticsUpdateHandler */


// The results handler for monitoring updates to the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatisticsCollectionQuery/statisticsUpdateHandler
func (h_ HKStatisticsCollectionQuery) SetStatisticsUpdateHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setStatisticsUpdateHandler:"), value)
}/* debug [instance_properties/setter]: statisticsUpdateHandler */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKStatisticsCollectionQuery */


