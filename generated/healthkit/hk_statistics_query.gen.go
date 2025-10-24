// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKStatisticsQuery */


/* debug [class_header]: Header for HKStatisticsQuery */
// The class instance for the [HKStatisticsQuery] class.
var (
	HKStatisticsQueryClass     _HKStatisticsQueryClass
	HKStatisticsQueryClassOnce sync.Once
)

func getHKStatisticsQueryClass() _HKStatisticsQueryClass {
	HKStatisticsQueryClassOnce.Do(func() {
		HKStatisticsQueryClass = _HKStatisticsQueryClass{objc.GetClass("HKStatisticsQuery")}
	})
	return HKStatisticsQueryClass
}

type _HKStatisticsQueryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKStatisticsQuery */
// An interface definition for the [HKStatisticsQuery] class.
type IHKStatisticsQuery interface {
	IHKQuery
	
/* debug [class_interface_properties]: Properties for HKStatisticsQuery */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKStatisticsQuery */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKStatisticsQuery */
// Alloc allocates a new instance without initialization.
func (hc _HKStatisticsQueryClass) Alloc() HKStatisticsQuery {
	rv := objc.Send[HKStatisticsQuery](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKStatisticsQueryClass) New() HKStatisticsQuery {
	rv := objc.Send[HKStatisticsQuery](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKStatisticsQuery) Init() HKStatisticsQuery {
	rv := objc.Send[HKStatisticsQuery](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKStatisticsQuery) Autorelease() HKStatisticsQuery {
	rv := objc.Send[HKStatisticsQuery](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKStatisticsQuery creates a new HKStatisticsQuery instance.
func NewHKStatisticsQuery() HKStatisticsQuery {
	return getHKStatisticsQueryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKStatisticsQuery */
// A query that performs statistical calculations over a set of matching quantity samples, and returns the results.
//
// Statistics queries calculate common statistics over the set of matching samples. You can use statistical queries to calculate the minimum, maximum, or average value of a set of discrete quantities, or use them to calculate the sum for cumulative quantities. For the complete list of possible calculations, see . For more information about the available quantity types and to learn whether they are discrete or cumulative values, see . You can use statistics queries with quantity samples only. If you want to calculate statistics over workouts or correlation samples, you must perform the appropriate query and process the data yourself. Statistics queries are immutable. Their properties are set when they are first created, and they can’t change.


// A query that performs statistical calculations over a set of matching quantity samples, and returns the results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatisticsQuery
type HKStatisticsQuery struct {
	HKQuery
}

// HKStatisticsQueryFrom constructs a [HKStatisticsQuery] from an unsafe.Pointer.
//
// A query that performs statistical calculations over a set of matching quantity samples, and returns the results.
func HKStatisticsQueryFrom(ptr unsafe.Pointer) HKStatisticsQuery {
	return HKStatisticsQuery{
		HKQuery: HKQueryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKStatisticsQuery */

// Initializes a statistics query instance that performs the specified calculations over the matching samples in the HeathKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatisticsQuery/init(quantityType:quantitySamplePredicate:options:completionHandler:)
func NewHKStatisticsQueryWithQuantityTypeQuantitySamplePredicateOptionsCompletionHandler(quantityType IHKQuantityType, quantitySamplePredicate foundation.Predicate, options HKStatisticsOptions, handler unsafe.Pointer) HKStatisticsQuery {
	instance := getHKStatisticsQueryClass().Alloc()
	rv := objc.Send[HKStatisticsQuery](instance.ID, objc.Sel("initWithQuantityType:quantitySamplePredicate:options:completionHandler:"), quantityType, quantitySamplePredicate, options, handler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKStatisticsQueryWithQuantityTypeQuantitySamplePredicateOptionsCompletionHandler */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKStatisticsQuery */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKStatisticsQuery */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKStatisticsQuery */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKStatisticsQuery */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKStatisticsQuery */


