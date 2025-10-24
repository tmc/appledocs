// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKCorrelationQuery */


/* debug [class_header]: Header for HKCorrelationQuery */
// The class instance for the [HKCorrelationQuery] class.
var (
	HKCorrelationQueryClass     _HKCorrelationQueryClass
	HKCorrelationQueryClassOnce sync.Once
)

func getHKCorrelationQueryClass() _HKCorrelationQueryClass {
	HKCorrelationQueryClassOnce.Do(func() {
		HKCorrelationQueryClass = _HKCorrelationQueryClass{objc.GetClass("HKCorrelationQuery")}
	})
	return HKCorrelationQueryClass
}

type _HKCorrelationQueryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKCorrelationQuery */
// An interface definition for the [HKCorrelationQuery] class.
type IHKCorrelationQuery interface {
	IHKQuery
	
/* debug [class_interface_properties]: Properties for HKCorrelationQuery */
	// properties:
	CorrelationType() IHKCorrelationType
	SamplePredicates() foundation.IDictionary
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKCorrelationQuery */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKCorrelationQuery */
// Alloc allocates a new instance without initialization.
func (hc _HKCorrelationQueryClass) Alloc() HKCorrelationQuery {
	rv := objc.Send[HKCorrelationQuery](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKCorrelationQueryClass) New() HKCorrelationQuery {
	rv := objc.Send[HKCorrelationQuery](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKCorrelationQuery) Init() HKCorrelationQuery {
	rv := objc.Send[HKCorrelationQuery](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKCorrelationQuery) Autorelease() HKCorrelationQuery {
	rv := objc.Send[HKCorrelationQuery](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKCorrelationQuery creates a new HKCorrelationQuery instance.
func NewHKCorrelationQuery() HKCorrelationQuery {
	return getHKCorrelationQueryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKCorrelationQuery */
// A query that performs complex searches based on the correlation’s contents, and returns a snapshot of all matching samples.
//
// Correlation samples act as a container, grouping multiple quantity or category samples. While you can use objects to search for correlations, correlation queries allow more complex filtering based on the contained samples. Specifically, correlation queries let you provide a separate predicate for each of the sample types stored in the correlation. A correlation is returned only if the correlation’s predicate and all of the sample predicates match.


// A query that performs complex searches based on the correlation’s contents, and returns a snapshot of all matching samples.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCorrelationQuery
type HKCorrelationQuery struct {
	HKQuery
}

// HKCorrelationQueryFrom constructs a [HKCorrelationQuery] from an unsafe.Pointer.
//
// A query that performs complex searches based on the correlation’s contents, and returns a snapshot of all matching samples.
func HKCorrelationQueryFrom(ptr unsafe.Pointer) HKCorrelationQuery {
	return HKCorrelationQuery{
		HKQuery: HKQueryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKCorrelationQuery */

// Instantiates and returns a correlation query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCorrelationQuery/init(type:predicate:samplePredicates:completion:)
func NewHKCorrelationQueryWithTypePredicateSamplePredicatesCompletion(correlationType IHKCorrelationType, predicate foundation.Predicate, samplePredicates foundation.IDictionary, completion unsafe.Pointer) HKCorrelationQuery {
	instance := getHKCorrelationQueryClass().Alloc()
	rv := objc.Send[HKCorrelationQuery](instance.ID, objc.Sel("initWithType:predicate:samplePredicates:completion:"), correlationType, predicate, samplePredicates, completion)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKCorrelationQueryWithTypePredicateSamplePredicatesCompletion */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKCorrelationQuery */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKCorrelationQuery */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKCorrelationQuery */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKCorrelationQuery */

// The type of correlation to search for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCorrelationQuery/correlationType
func (h_ HKCorrelationQuery) CorrelationType() IHKCorrelationType {
	rv := objc.Send[HKCorrelationType](h_.ID, objc.Sel("correlationType"))
	return rv
}/* debug [instance_properties/getter]: correlationType */


// A dictionary whose keys are instances and whose values are instances.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCorrelationQuery/samplePredicates
func (h_ HKCorrelationQuery) SamplePredicates() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](h_.ID, objc.Sel("samplePredicates"))
	return rv
}/* debug [instance_properties/getter]: samplePredicates */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKCorrelationQuery */


