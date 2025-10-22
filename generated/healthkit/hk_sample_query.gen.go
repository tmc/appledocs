// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [HKSampleQuery] class.
type IHKSampleQuery interface {
	IHKQuery
	HKObjectQueryNoLimit() int
	Limit() int
	SetLimit(value int)
	SortDescriptors() foundation.SortDescriptor
	SetSortDescriptors(value foundation.ISortDescriptor)
}

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

// Alloc allocates a new instance without initialization.
func (hc _HKSampleQueryClass) Alloc() HKSampleQuery {
	rv := objc.Send[HKSampleQuery](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates a query for samples that match any of the query descriptors you provided, sorted by the sort descriptors you provided.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSampleQuery/init(queryDescriptors:limit:sortDescriptors:resultsHandler:)

func NewHKSampleQueryWithQueryDescriptorsLimitSortDescriptorsResultsHandler(queryDescriptors []HKQueryDescriptor, limit int, sortDescriptors []foundation.ISortDescriptor, resultsHandler unsafe.Pointer) HKSampleQuery {
	instance := getHKSampleQueryClass().Alloc()
	rv := objc.Send[HKSampleQuery](instance.ID, objc.Sel("initWithQueryDescriptors:limit:sortDescriptors:resultsHandler:"), queryDescriptors, limit, sortDescriptors, resultsHandler)
	rv.Autorelease()
	return rv
}



// A value indicating that the query returns all the matching samples in the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkobjectquerynolimit

func (h_ HKSampleQuery) HKObjectQueryNoLimit() int {
	rv := objc.Send[int](h_.ID, objc.Sel("HKObjectQueryNoLimit"))
	return rv
}


// The maximum number of samples that this query returns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hksamplequery/limit

func (h_ HKSampleQuery) Limit() int {
	rv := objc.Send[int](h_.ID, objc.Sel("limit"))
	return rv
}


// The maximum number of samples that this query returns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hksamplequery/limit

func (h_ HKSampleQuery) SetLimit(value int) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setLimit:"), value)
}


// The sort descriptors that specify the order of the results returned by this query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hksamplequery/sortdescriptors

func (h_ HKSampleQuery) SortDescriptors() foundation.SortDescriptor {
	rv := objc.Send[foundation.SortDescriptor](h_.ID, objc.Sel("sortDescriptors"))
	return rv
}


// The sort descriptors that specify the order of the results returned by this query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hksamplequery/sortdescriptors

func (h_ HKSampleQuery) SetSortDescriptors(value foundation.ISortDescriptor) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSortDescriptors:"), value)
}


