// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [HKCorrelationQuery] class.
type IHKCorrelationQuery interface {
	IHKQuery
}

// A query that performs complex searches based on the correlation’s contents, and returns a snapshot of all matching samples.
//
// Correlation samples act as a container, grouping multiple quantity or category samples. While you can use objects to search for correlations, correlation queries allow more complex filtering based on the contained samples. Specifically, correlation queries let you provide a separate predicate for each of the sample types stored in the correlation. A correlation is returned only if the correlation’s predicate and all of the sample predicates match.
//
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

// Alloc allocates a new instance without initialization.
func (hc _HKCorrelationQueryClass) Alloc() HKCorrelationQuery {
	rv := objc.Send[HKCorrelationQuery](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Instantiates and returns a correlation query.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCorrelationQuery/init(type:predicate:samplePredicates:completion:)
func NewHKCorrelationQueryWithTypePredicateSamplePredicatesCompletion(correlationType unsafe.Pointer, predicate unsafe.Pointer, samplePredicates unsafe.Pointer, completion unsafe.Pointer) HKCorrelationQuery {
	instance := getHKCorrelationQueryClass().Alloc()
	rv := objc.Send[HKCorrelationQuery](instance.ID, objc.Sel("initWithType:predicate:samplePredicates:completion:"), correlationType, predicate, samplePredicates, completion)
	rv.Autorelease()
	return rv
}


// The type of correlation to search for.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCorrelationQuery/correlationType
func (h_ HKCorrelationQuery) CorrelationType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("correlationType"))
	return rv
}

// A dictionary whose keys are instances and whose values are instances.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCorrelationQuery/samplePredicates
func (h_ HKCorrelationQuery) SamplePredicates() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("samplePredicates"))
	return rv
}


