// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MetadataQuery] class.
var (
	MetadataQueryClass     _MetadataQueryClass
	MetadataQueryClassOnce sync.Once
)

func getMetadataQueryClass() _MetadataQueryClass {
	MetadataQueryClassOnce.Do(func() {
		MetadataQueryClass = _MetadataQueryClass{objc.GetClass("NSMetadataQuery")}
	})
	return MetadataQueryClass
}

type _MetadataQueryClass struct {
	class objc.Class
}

// An interface definition for the [MetadataQuery] class.
type IMetadataQuery interface {
	objectivec.IObject
	EnumerateResultsUsingBlock(block unsafe.Pointer)
	ResultAtIndex(idx uint) objc.ID
}

// A query that you perform against Spotlight metadata.
//
// The class encapsulates the functionality provided by the opaque type for querying the Spotlight metadata. objects provide metadata query results in several ways: As individual attribute values for requested attributes. As value lists that contain the distinct values for given attributes in the query results. As a result array proxy, containing all the query results. This is suitable for use with Cocoa bindings. As a hierarchical collection of results, grouping together items with the same values for specified grouping attributes. This is also suitable for use with Cocoa bindings. Queries have two phases: the initial gathering phase that collects all currently matching results and a second live-update phase. By default, the receiver has no limitation on its search scope. Use the property to customize. By default, notification of updated results occurs at 1.0 seconds. Use the property to customize. You must set a predicate with the property before starting a query.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery
type MetadataQuery struct {
	objectivec.Object
}

// MetadataQueryFrom constructs a [MetadataQuery] from an unsafe.Pointer.
//
// A query that you perform against Spotlight metadata.
func MetadataQueryFrom(ptr unsafe.Pointer) MetadataQuery {
	return MetadataQuery{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MetadataQueryClass) Alloc() MetadataQuery {
	rv := objc.Send[MetadataQuery](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MetadataQueryClass) New() MetadataQuery {
	rv := objc.Send[MetadataQuery](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetadataQuery) Init() MetadataQuery {
	rv := objc.Send[MetadataQuery](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetadataQuery) Autorelease() MetadataQuery {
	rv := objc.Send[MetadataQuery](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetadataQuery creates a new MetadataQuery instance.
func NewMetadataQuery() MetadataQuery {
	return getMetadataQueryClass().New()
}


// Enumerates the current set of results using the given block.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/enumerateResults(_:)
func (m_ MetadataQuery) EnumerateResultsUsingBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("enumerateResultsUsingBlock:"), block)
}

// Returns the query result at a specific index.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/result(at:)
func (m_ MetadataQuery) ResultAtIndex(idx uint) objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("resultAtIndex:"), idx)
	return rv
}

// An array of grouping attributes. (read-only)
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/groupingattributes
func (m_ MetadataQuery) GroupingAttributes() string {
	rv := objc.Send[string](m_.ID, objc.Sel("groupingAttributes"))
	return rv
}


// SetGroupingAttributes sets the value of the groupingAttributes property.
// An array of grouping attributes. (read-only)

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/groupingattributes
func (m_ MetadataQuery) SetGroupingAttributes(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupingAttributes:"), objc.String(value))
}

// A Boolean value that indicates whether the query has started. (read-only)
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/isstarted
func (m_ MetadataQuery) IsStarted() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isStarted"))
	return rv
}


// SetIsStarted sets the value of the isStarted property.
// A Boolean value that indicates whether the query has started. (read-only)

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/isstarted
func (m_ MetadataQuery) SetIsStarted(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsStarted:"), value)
}

// An array of sort descriptor objects.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/sortdescriptors
func (m_ MetadataQuery) SortDescriptors() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("sortDescriptors"))
	return rv
}


// SetSortDescriptors sets the value of the sortDescriptors property.
// An array of sort descriptor objects.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/sortdescriptors
func (m_ MetadataQuery) SetSortDescriptors(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSortDescriptors:"), value)
}

// The query’s delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/delegate
func (m_ MetadataQuery) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The query’s delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/delegate
func (m_ MetadataQuery) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelegate:"), value)
}

// A dictionary containing the value lists generated by the query.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/valuelists
func (m_ MetadataQuery) ValueLists() string {
	rv := objc.Send[string](m_.ID, objc.Sel("valueLists"))
	return rv
}


// SetValueLists sets the value of the valueLists property.
// A dictionary containing the value lists generated by the query.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/valuelists
func (m_ MetadataQuery) SetValueLists(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValueLists:"), objc.String(value))
}

// The number of results returned by the query. (read-only)
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/resultcount
func (m_ MetadataQuery) ResultCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("resultCount"))
	return rv
}


// SetResultCount sets the value of the resultCount property.
// The number of results returned by the query. (read-only)

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/resultcount
func (m_ MetadataQuery) SetResultCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setResultCount:"), value)
}

// An array containing hierarchical groups of query results. (read-only)
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/groupedresults
func (m_ MetadataQuery) GroupedResults() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("groupedResults"))
	return rv
}


// SetGroupedResults sets the value of the groupedResults property.
// An array containing hierarchical groups of query results. (read-only)

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/groupedresults
func (m_ MetadataQuery) SetGroupedResults(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupedResults:"), value)
}

// A Boolean value that indicates whether the receiver is in the initial gathering phase of the query. (read-only)
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/isgathering
func (m_ MetadataQuery) IsGathering() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isGathering"))
	return rv
}


// SetIsGathering sets the value of the isGathering property.
// A Boolean value that indicates whether the receiver is in the initial gathering phase of the query. (read-only)

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/isgathering
func (m_ MetadataQuery) SetIsGathering(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsGathering:"), value)
}

// An array of attributes whose values are gathered by the query.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/valuelistattributes
func (m_ MetadataQuery) ValueListAttributes() string {
	rv := objc.Send[string](m_.ID, objc.Sel("valueListAttributes"))
	return rv
}


// SetValueListAttributes sets the value of the valueListAttributes property.
// An array of attributes whose values are gathered by the query.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/valuelistattributes
func (m_ MetadataQuery) SetValueListAttributes(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValueListAttributes:"), objc.String(value))
}

// An array of objects that define the query’s scope.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/searchitems
func (m_ MetadataQuery) SearchItems() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("searchItems"))
	return rv
}


// SetSearchItems sets the value of the searchItems property.
// An array of objects that define the query’s scope.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/searchitems
func (m_ MetadataQuery) SetSearchItems(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSearchItems:"), value)
}

// A Boolean value that indicates whether the query has stopped.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/isstopped
func (m_ MetadataQuery) IsStopped() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isStopped"))
	return rv
}


// SetIsStopped sets the value of the isStopped property.
// A Boolean value that indicates whether the query has stopped.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/isstopped
func (m_ MetadataQuery) SetIsStopped(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsStopped:"), value)
}

// The interval at which notification of updated results occurs.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/notificationBatchingInterval
func (m_ MetadataQuery) NotificationBatchingInterval() TimeInterval {
	rv := objc.Send[TimeInterval](m_.ID, objc.Sel("notificationBatchingInterval"))
	return rv
}


// SetNotificationBatchingInterval sets the value of the notificationBatchingInterval property.
// The interval at which notification of updated results occurs.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/notificationBatchingInterval
func (m_ MetadataQuery) SetNotificationBatchingInterval(value TimeInterval) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNotificationBatchingInterval:"), value)
}

// The queue on which query result notifications are posted.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/operationQueue
func (m_ MetadataQuery) OperationQueue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("operationQueue"))
	return rv
}


// SetOperationQueue sets the value of the operationQueue property.
// The queue on which query result notifications are posted.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/operationQueue
func (m_ MetadataQuery) SetOperationQueue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationQueue:"), value)
}

// The predicate used to filter query results.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/predicate
func (m_ MetadataQuery) Predicate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("predicate"))
	return rv
}


// SetPredicate sets the value of the predicate property.
// The predicate used to filter query results.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/predicate
func (m_ MetadataQuery) SetPredicate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPredicate:"), value)
}

// An array containing the query’s results.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/results
func (m_ MetadataQuery) Results() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("results"))
	return rv
}

// An array containing the search scopes.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/searchScopes
func (m_ MetadataQuery) SearchScopes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("searchScopes"))
	return rv
}


// SetSearchScopes sets the value of the searchScopes property.
// An array containing the search scopes.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/searchScopes
func (m_ MetadataQuery) SetSearchScopes(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSearchScopes:"), value)
}



