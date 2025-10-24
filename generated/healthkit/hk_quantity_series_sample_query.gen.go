// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKQuantitySeriesSampleQuery */


/* debug [class_header]: Header for HKQuantitySeriesSampleQuery */
// The class instance for the [HKQuantitySeriesSampleQuery] class.
var (
	HKQuantitySeriesSampleQueryClass     _HKQuantitySeriesSampleQueryClass
	HKQuantitySeriesSampleQueryClassOnce sync.Once
)

func getHKQuantitySeriesSampleQueryClass() _HKQuantitySeriesSampleQueryClass {
	HKQuantitySeriesSampleQueryClassOnce.Do(func() {
		HKQuantitySeriesSampleQueryClass = _HKQuantitySeriesSampleQueryClass{objc.GetClass("HKQuantitySeriesSampleQuery")}
	})
	return HKQuantitySeriesSampleQueryClass
}

type _HKQuantitySeriesSampleQueryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKQuantitySeriesSampleQuery */
// An interface definition for the [HKQuantitySeriesSampleQuery] class.
type IHKQuantitySeriesSampleQuery interface {
	IHKQuery
	
/* debug [class_interface_properties]: Properties for HKQuantitySeriesSampleQuery */
	// properties:
	IncludeSample() bool
	SetIncludeSample(value bool)
	OrderByQuantitySampleStartDate() bool
	SetOrderByQuantitySampleStartDate(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKQuantitySeriesSampleQuery */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKQuantitySeriesSampleQuery */
// Alloc allocates a new instance without initialization.
func (hc _HKQuantitySeriesSampleQueryClass) Alloc() HKQuantitySeriesSampleQuery {
	rv := objc.Send[HKQuantitySeriesSampleQuery](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKQuantitySeriesSampleQueryClass) New() HKQuantitySeriesSampleQuery {
	rv := objc.Send[HKQuantitySeriesSampleQuery](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKQuantitySeriesSampleQuery) Init() HKQuantitySeriesSampleQuery {
	rv := objc.Send[HKQuantitySeriesSampleQuery](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKQuantitySeriesSampleQuery) Autorelease() HKQuantitySeriesSampleQuery {
	rv := objc.Send[HKQuantitySeriesSampleQuery](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKQuantitySeriesSampleQuery creates a new HKQuantitySeriesSampleQuery instance.
func NewHKQuantitySeriesSampleQuery() HKQuantitySeriesSampleQuery {
	return getHKQuantitySeriesSampleQueryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKQuantitySeriesSampleQuery */
// A query that accesses the series data associated with a quantity sample.
//
// Use a series query to access the individual objects added to a sample using an .


// A query that accesses the series data associated with a quantity sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySeriesSampleQuery
type HKQuantitySeriesSampleQuery struct {
	HKQuery
}

// HKQuantitySeriesSampleQueryFrom constructs a [HKQuantitySeriesSampleQuery] from an unsafe.Pointer.
//
// A query that accesses the series data associated with a quantity sample.
func HKQuantitySeriesSampleQueryFrom(ptr unsafe.Pointer) HKQuantitySeriesSampleQuery {
	return HKQuantitySeriesSampleQuery{
		HKQuery: HKQueryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKQuantitySeriesSampleQuery */

// Creates a new query for a series of the specified quantity type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySeriesSampleQuery/init(quantityType:predicate:quantityHandler:)
func NewHKQuantitySeriesSampleQueryWithQuantityTypePredicateQuantityHandler(quantityType IHKQuantityType, predicate foundation.Predicate, quantityHandler unsafe.Pointer) HKQuantitySeriesSampleQuery {
	instance := getHKQuantitySeriesSampleQueryClass().Alloc()
	rv := objc.Send[HKQuantitySeriesSampleQuery](instance.ID, objc.Sel("initWithQuantityType:predicate:quantityHandler:"), quantityType, predicate, quantityHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKQuantitySeriesSampleQueryWithQuantityTypePredicateQuantityHandler */


// Creates a new series query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySeriesSampleQuery/init(sample:quantityHandler:)
func NewHKQuantitySeriesSampleQueryWithSampleQuantityHandler(quantitySample IHKQuantitySample, quantityHandler unsafe.Pointer) HKQuantitySeriesSampleQuery {
	instance := getHKQuantitySeriesSampleQueryClass().Alloc()
	rv := objc.Send[HKQuantitySeriesSampleQuery](instance.ID, objc.Sel("initWithSample:quantityHandler:"), quantitySample, quantityHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKQuantitySeriesSampleQueryWithSampleQuantityHandler */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKQuantitySeriesSampleQuery */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKQuantitySeriesSampleQuery */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKQuantitySeriesSampleQuery */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKQuantitySeriesSampleQuery */

// A Boolean value that determines whether the query should return the series sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySeriesSampleQuery/includeSample
func (h_ HKQuantitySeriesSampleQuery) IncludeSample() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("includeSample"))
	return rv
}/* debug [instance_properties/getter]: includeSample */


// A Boolean value that determines whether the query should return the series sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySeriesSampleQuery/includeSample
func (h_ HKQuantitySeriesSampleQuery) SetIncludeSample(value bool) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setIncludeSample:"), value)
}/* debug [instance_properties/setter]: includeSample */


// A Boolean value that determines whether the query groups the results based on the quantity sample’s start date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySeriesSampleQuery/orderByQuantitySampleStartDate
func (h_ HKQuantitySeriesSampleQuery) OrderByQuantitySampleStartDate() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("orderByQuantitySampleStartDate"))
	return rv
}/* debug [instance_properties/getter]: orderByQuantitySampleStartDate */


// A Boolean value that determines whether the query groups the results based on the quantity sample’s start date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySeriesSampleQuery/orderByQuantitySampleStartDate
func (h_ HKQuantitySeriesSampleQuery) SetOrderByQuantitySampleStartDate(value bool) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setOrderByQuantitySampleStartDate:"), value)
}/* debug [instance_properties/setter]: orderByQuantitySampleStartDate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKQuantitySeriesSampleQuery */


