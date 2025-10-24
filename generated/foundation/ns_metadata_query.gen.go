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
	// properties:
	SearchScopes() IArray
	SetSearchScopes(value IArray)
	Delegate() objc.IObject /* cross-framework: MetadataQueryDelegate */
	SetDelegate(value objc.IObject /* cross-framework: MetadataQueryDelegate */)
	GroupedResults() MetadataQueryResultGroup /* not a class type */
	SetGroupedResults(value MetadataQueryResultGroup /* not a class type */)
	GroupingAttributes() IString
	SetGroupingAttributes(value IString)
	IsGathering() bool
	SetIsGathering(value bool)
	IsStarted() bool
	SetIsStarted(value bool)
	IsStopped() bool
	SetIsStopped(value bool)
	NotificationBatchingInterval() float64
	SetNotificationBatchingInterval(value float64)
	OperationQueue() IOperationQueue
	SetOperationQueue(value IOperationQueue)
	Predicate() IPredicate
	SetPredicate(value IPredicate)
	ResultCount() int
	SetResultCount(value int)
	SortDescriptors() ISortDescriptor
	SetSortDescriptors(value ISortDescriptor)
	ValueListAttributes() IString
	SetValueListAttributes(value IString)
	ValueLists() MetadataQueryAttributeValueTuple /* not a class type */
	SetValueLists(value MetadataQueryAttributeValueTuple /* not a class type */)
	// methods:
}

// A query that you perform against Spotlight metadata.
//
// The class encapsulates the functionality provided by the opaque type for querying the Spotlight metadata. objects provide metadata query results in several ways: As individual attribute values for requested attributes. As value lists that contain the distinct values for given attributes in the query results. As a result array proxy, containing all the query results. This is suitable for use with Cocoa bindings. As a hierarchical collection of results, grouping together items with the same values for specified grouping attributes. This is also suitable for use with Cocoa bindings. Queries have two phases: the initial gathering phase that collects all currently matching results and a second live-update phase. By default, the receiver has no limitation on its search scope. Use the property to customize. By default, notification of updated results occurs at 1.0 seconds. Use the property to customize. You must set a predicate with the property before starting a query.


// A query that you perform against Spotlight metadata.
//
// [Full Topic]
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



// An array containing the search scopes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/searchScopes
func (m_ MetadataQuery) SearchScopes() IArray {
	rv := objc.Send[Array](m_.ID, objc.Sel("searchScopes"))
	return rv
}


// An array containing the search scopes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/searchScopes
func (m_ MetadataQuery) SetSearchScopes(value IArray) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSearchScopes:"), value)
}


// The query’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/delegate
func (m_ MetadataQuery) Delegate() objc.IObject /* cross-framework: MetadataQueryDelegate */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("delegate"))
	return rv
}


// The query’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/delegate
func (m_ MetadataQuery) SetDelegate(value objc.IObject /* cross-framework: MetadataQueryDelegate */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelegate:"), value)
}


// An array containing hierarchical groups of query results. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/groupedresults
func (m_ MetadataQuery) GroupedResults() MetadataQueryResultGroup /* not a class type */ {
	rv := objc.Send[MetadataQueryResultGroup](m_.ID, objc.Sel("groupedResults"))
	return rv
}


// An array containing hierarchical groups of query results. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/groupedresults
func (m_ MetadataQuery) SetGroupedResults(value MetadataQueryResultGroup /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupedResults:"), value)
}


// An array of grouping attributes. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/groupingattributes
func (m_ MetadataQuery) GroupingAttributes() IString {
	rv := objc.Send[String](m_.ID, objc.Sel("groupingAttributes"))
	return rv
}


// An array of grouping attributes. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/groupingattributes
func (m_ MetadataQuery) SetGroupingAttributes(value IString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupingAttributes:"), value)
}


// A Boolean value that indicates whether the receiver is in the initial gathering phase of the query. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/isgathering
func (m_ MetadataQuery) IsGathering() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isGathering"))
	return rv
}


// A Boolean value that indicates whether the receiver is in the initial gathering phase of the query. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/isgathering
func (m_ MetadataQuery) SetIsGathering(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsGathering:"), value)
}


// A Boolean value that indicates whether the query has started. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/isstarted
func (m_ MetadataQuery) IsStarted() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isStarted"))
	return rv
}


// A Boolean value that indicates whether the query has started. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/isstarted
func (m_ MetadataQuery) SetIsStarted(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsStarted:"), value)
}


// A Boolean value that indicates whether the query has stopped.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/isstopped
func (m_ MetadataQuery) IsStopped() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isStopped"))
	return rv
}


// A Boolean value that indicates whether the query has stopped.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/isstopped
func (m_ MetadataQuery) SetIsStopped(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsStopped:"), value)
}


// The interval at which notification of updated results occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/notificationbatchinginterval
func (m_ MetadataQuery) NotificationBatchingInterval() float64 {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("notificationBatchingInterval"))
	return rv
}


// The interval at which notification of updated results occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/notificationbatchinginterval
func (m_ MetadataQuery) SetNotificationBatchingInterval(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNotificationBatchingInterval:"), value)
}


// The queue on which query result notifications are posted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/operationqueue
func (m_ MetadataQuery) OperationQueue() IOperationQueue {
	rv := objc.Send[OperationQueue](m_.ID, objc.Sel("operationQueue"))
	return rv
}


// The queue on which query result notifications are posted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/operationqueue
func (m_ MetadataQuery) SetOperationQueue(value IOperationQueue) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationQueue:"), value)
}


// The predicate used to filter query results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/predicate
func (m_ MetadataQuery) Predicate() IPredicate {
	rv := objc.Send[Predicate](m_.ID, objc.Sel("predicate"))
	return rv
}


// The predicate used to filter query results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/predicate
func (m_ MetadataQuery) SetPredicate(value IPredicate) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPredicate:"), value)
}


// The number of results returned by the query. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/resultcount
func (m_ MetadataQuery) ResultCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("resultCount"))
	return rv
}


// The number of results returned by the query. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/resultcount
func (m_ MetadataQuery) SetResultCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setResultCount:"), value)
}


// An array of sort descriptor objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/sortdescriptors
func (m_ MetadataQuery) SortDescriptors() ISortDescriptor {
	rv := objc.Send[SortDescriptor](m_.ID, objc.Sel("sortDescriptors"))
	return rv
}


// An array of sort descriptor objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/sortdescriptors
func (m_ MetadataQuery) SetSortDescriptors(value ISortDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSortDescriptors:"), value)
}


// An array of attributes whose values are gathered by the query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/valuelistattributes
func (m_ MetadataQuery) ValueListAttributes() IString {
	rv := objc.Send[String](m_.ID, objc.Sel("valueListAttributes"))
	return rv
}


// An array of attributes whose values are gathered by the query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/valuelistattributes
func (m_ MetadataQuery) SetValueListAttributes(value IString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValueListAttributes:"), value)
}


// A dictionary containing the value lists generated by the query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/valuelists
func (m_ MetadataQuery) ValueLists() MetadataQueryAttributeValueTuple /* not a class type */ {
	rv := objc.Send[MetadataQueryAttributeValueTuple](m_.ID, objc.Sel("valueLists"))
	return rv
}


// A dictionary containing the value lists generated by the query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/valuelists
func (m_ MetadataQuery) SetValueLists(value MetadataQueryAttributeValueTuple /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValueLists:"), value)
}



