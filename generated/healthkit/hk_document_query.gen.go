// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKDocumentQuery] class.
var (
	HKDocumentQueryClass     _HKDocumentQueryClass
	HKDocumentQueryClassOnce sync.Once
)

func getHKDocumentQueryClass() _HKDocumentQueryClass {
	HKDocumentQueryClassOnce.Do(func() {
		HKDocumentQueryClass = _HKDocumentQueryClass{objc.GetClass("HKDocumentQuery")}
	})
	return HKDocumentQueryClass
}

type _HKDocumentQueryClass struct {
	class objc.Class
}

// An interface definition for the [HKDocumentQuery] class.
type IHKDocumentQuery interface {
	IHKQuery
	// properties:
	IncludeDocumentData() bool /* primitive/slice/pointer. */
	SetIncludeDocumentData(value bool /* primitive/slice/pointer. */)
	Limit() int /* primitive/slice/pointer. */
	SetLimit(value int /* primitive/slice/pointer. */)
	SortDescriptors() SortDescriptor /* not a class type */
	SetSortDescriptors(value SortDescriptor /* not a class type */)
	HKObjectQueryNoLimit() int /* primitive/slice/pointer. */
	// methods:
}

// A query that returns a snapshot of all matching documents currently saved in the HealthKit store.
//
// Use an object to search for documents in the HealthKit store. You can provide a predicate to filter the search results, a sort order for the returned samples, or even a limit to the number of samples returned. Document queries are immutable: The query’s properties are set when the query is first created. They cannot change.


// A query that returns a snapshot of all matching documents currently saved in the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDocumentQuery
type HKDocumentQuery struct {
	HKQuery
}

// HKDocumentQueryFrom constructs a [HKDocumentQuery] from an unsafe.Pointer.
//
// A query that returns a snapshot of all matching documents currently saved in the HealthKit store.
func HKDocumentQueryFrom(ptr unsafe.Pointer) HKDocumentQuery {
	return HKDocumentQuery{
		HKQuery: HKQueryFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKDocumentQueryClass) Alloc() HKDocumentQuery {
	rv := objc.Send[HKDocumentQuery](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKDocumentQueryClass) New() HKDocumentQuery {
	rv := objc.Send[HKDocumentQuery](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKDocumentQuery) Init() HKDocumentQuery {
	rv := objc.Send[HKDocumentQuery](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKDocumentQuery) Autorelease() HKDocumentQuery {
	rv := objc.Send[HKDocumentQuery](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKDocumentQuery creates a new HKDocumentQuery instance.
func NewHKDocumentQuery() HKDocumentQuery {
	return getHKDocumentQueryClass().New()
}



// A Boolean value that indicates whether the sample includes the full document’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdocumentquery/includedocumentdata
func (h_ HKDocumentQuery) IncludeDocumentData() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](h_.ID, objc.Sel("includeDocumentData"))
	return rv
}


// A Boolean value that indicates whether the sample includes the full document’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdocumentquery/includedocumentdata
func (h_ HKDocumentQuery) SetIncludeDocumentData(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setIncludeDocumentData:"), value)
}


// The maximum number of documents the receiver will return upon completion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdocumentquery/limit
func (h_ HKDocumentQuery) Limit() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](h_.ID, objc.Sel("limit"))
	return rv
}


// The maximum number of documents the receiver will return upon completion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdocumentquery/limit
func (h_ HKDocumentQuery) SetLimit(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setLimit:"), value)
}


// An array of sort descriptors that specify the order of the results returned by this query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdocumentquery/sortdescriptors
func (h_ HKDocumentQuery) SortDescriptors() SortDescriptor /* not a class type */ {
	rv := objc.Send[SortDescriptor](h_.ID, objc.Sel("sortDescriptors"))
	return rv
}


// An array of sort descriptors that specify the order of the results returned by this query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdocumentquery/sortdescriptors
func (h_ HKDocumentQuery) SetSortDescriptors(value SortDescriptor /* not a class type */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSortDescriptors:"), value)
}


// A value indicating that the query returns all the matching samples in the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkobjectquerynolimit
func (h_ HKDocumentQuery) HKObjectQueryNoLimit() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](h_.ID, objc.Sel("HKObjectQueryNoLimit"))
	return rv
}



