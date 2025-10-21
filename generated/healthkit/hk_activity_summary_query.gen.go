// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [HKActivitySummaryQuery] class.
type IHKActivitySummaryQuery interface {
	IHKQuery
}

// A query for reading activity summary objects from the HealthKit store.
//
// Activity summary query objects are mostly immutable. You can assign the query’s property after instantiating the object, but before executing the query. All other properties must be set when you instantiate the object, and they can’t change.
//
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

// Alloc allocates a new instance without initialization.
func (hc _HKActivitySummaryQueryClass) Alloc() HKActivitySummaryQuery {
	rv := objc.Send[HKActivitySummaryQuery](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Initializes a new active summary query.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummaryQuery/init(predicate:resultsHandler:)
func NewHKActivitySummaryQueryWithPredicateResultsHandler(predicate foundation.IPredicate, handler unsafe.Pointer) HKActivitySummaryQuery {
	instance := getHKActivitySummaryQueryClass().Alloc()
	rv := objc.Send[HKActivitySummaryQuery](instance.ID, objc.Sel("initWithPredicate:resultsHandler:"), predicate, handler)
	rv.Autorelease()
	return rv
}


// The handler for monitoring updates to activity summaries saved in the HealthKit store.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummaryquery/updatehandler
func (h_ HKActivitySummaryQuery) UpdateHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("updateHandler"))
	return rv
}


// SetUpdateHandler sets the value of the updateHandler property.
// The handler for monitoring updates to activity summaries saved in the HealthKit store.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitysummaryquery/updatehandler
func (h_ HKActivitySummaryQuery) SetUpdateHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setUpdateHandler:"), value)
}


