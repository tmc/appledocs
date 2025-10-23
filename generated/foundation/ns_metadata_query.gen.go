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
	DisableUpdates()
	EnableUpdates()
	EnumerateResultsUsingBlock(block unsafe.Pointer)
	EnumerateResultsWithOptionsUsingBlock(opts NSEnumerationOptions, block unsafe.Pointer)
	IndexOfResult(result objectivec.IObject) uint
	ResultAtIndex(idx uint) objc.ID
	StartQuery() bool
	StopQuery()
	ValueOfAttributeForResultAtIndex(attrName string, idx uint) objc.ID
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	GroupedResults() []MetadataQueryResultGroup
	GroupingAttributes() []string
	SetGroupingAttributes(value []string)
	Gathering() bool
	Started() bool
	Stopped() bool
	NotificationBatchingInterval() TimeInterval
	SetNotificationBatchingInterval(value TimeInterval)
	OperationQueue() IOperationQueue
	SetOperationQueue(value IOperationQueue)
	Predicate() IPredicate
	SetPredicate(value IPredicate)
	ResultCount() uint
	Results() objc.ID
	SearchItems() objc.ID
	SetSearchItems(value objc.ID)
	SearchScopes() objc.ID
	SetSearchScopes(value objc.ID)
	SortDescriptors() []SortDescriptor
	SetSortDescriptors(value []SortDescriptor)
	ValueListAttributes() []string
	SetValueListAttributes(value []string)
	ValueLists() IDictionary
	IsGathering() bool
	SetIsGathering(value bool)
	IsStarted() bool
	SetIsStarted(value bool)
	IsStopped() bool
	SetIsStopped(value bool)
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



// Disables updates to the query results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/disableUpdates()
func (m_ MetadataQuery) DisableUpdates() {
	objc.Send[objc.ID](m_.ID, objc.Sel("disableUpdates"))
}


// Enables updates to the query results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/enableUpdates()
func (m_ MetadataQuery) EnableUpdates() {
	objc.Send[objc.ID](m_.ID, objc.Sel("enableUpdates"))
}


// Enumerates the current set of results using the given block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/enumerateResults(_:)
func (m_ MetadataQuery) EnumerateResultsUsingBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("enumerateResultsUsingBlock:"), block)
}


// Enumerates the current set of results using the given options and block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/enumerateResults(options:using:)
func (m_ MetadataQuery) EnumerateResultsWithOptionsUsingBlock(opts NSEnumerationOptions, block unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("enumerateResultsWithOptions:usingBlock:"), opts, block)
}


// Returns the index of a query result object in the receiver’s results array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/index(ofResult:)
func (m_ MetadataQuery) IndexOfResult(result objectivec.IObject) uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("indexOfResult:"), result)
	return rv
}


// Returns the query result at a specific index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/result(at:)
func (m_ MetadataQuery) ResultAtIndex(idx uint) objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("resultAtIndex:"), idx)
	return rv
}


// Attempts to start the query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/start()
func (m_ MetadataQuery) StartQuery() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("startQuery"))
	return rv
}


// Stops the receiver’s current query from gathering any further results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/stop()
func (m_ MetadataQuery) StopQuery() {
	objc.Send[objc.ID](m_.ID, objc.Sel("stopQuery"))
}


// Returns the value for the attribute name at the index in the results specified by .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/value(ofAttribute:forResultAt:)
func (m_ MetadataQuery) ValueOfAttributeForResultAtIndex(attrName string, idx uint) objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("valueOfAttribute:forResultAtIndex:"), objc.String(attrName), idx)
	return rv
}


// The query’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/delegate
func (m_ MetadataQuery) Delegate() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("delegate"))
	return rv
}


// The query’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/delegate
func (m_ MetadataQuery) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelegate:"), value)
}


// An array containing hierarchical groups of query results. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/groupedResults
func (m_ MetadataQuery) GroupedResults() []MetadataQueryResultGroup {
	rv := objc.Send[[]MetadataQueryResultGroup](m_.ID, objc.Sel("groupedResults"))
	return rv
}


// An array of grouping attributes. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/groupingAttributes
func (m_ MetadataQuery) GroupingAttributes() []string {
	rv := objc.Send[[]string](m_.ID, objc.Sel("groupingAttributes"))
	return rv
}


// An array of grouping attributes. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/groupingAttributes
func (m_ MetadataQuery) SetGroupingAttributes(value []string) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupingAttributes:"), nsArray)
}


// A Boolean value that indicates whether the receiver is in the initial gathering phase of the query. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/isGathering
func (m_ MetadataQuery) Gathering() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("gathering"))
	return rv
}


// A Boolean value that indicates whether the query has started. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/isStarted
func (m_ MetadataQuery) Started() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("started"))
	return rv
}


// A Boolean value that indicates whether the query has stopped.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/isStopped
func (m_ MetadataQuery) Stopped() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("stopped"))
	return rv
}


// The interval at which notification of updated results occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/notificationBatchingInterval
func (m_ MetadataQuery) NotificationBatchingInterval() TimeInterval {
	rv := objc.Send[TimeInterval](m_.ID, objc.Sel("notificationBatchingInterval"))
	return rv
}


// The interval at which notification of updated results occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/notificationBatchingInterval
func (m_ MetadataQuery) SetNotificationBatchingInterval(value TimeInterval) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNotificationBatchingInterval:"), value)
}


// The queue on which query result notifications are posted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/operationQueue
func (m_ MetadataQuery) OperationQueue() IOperationQueue {
	rv := objc.Send[NSOperationQueue](m_.ID, objc.Sel("operationQueue"))
	return rv
}


// The queue on which query result notifications are posted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/operationQueue
func (m_ MetadataQuery) SetOperationQueue(value IOperationQueue) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationQueue:"), value)
}


// The predicate used to filter query results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/predicate
func (m_ MetadataQuery) Predicate() IPredicate {
	rv := objc.Send[NSPredicate](m_.ID, objc.Sel("predicate"))
	return rv
}


// The predicate used to filter query results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/predicate
func (m_ MetadataQuery) SetPredicate(value IPredicate) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPredicate:"), value)
}


// The number of results returned by the query. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/resultCount
func (m_ MetadataQuery) ResultCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("resultCount"))
	return rv
}


// An array containing the query’s results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/results
func (m_ MetadataQuery) Results() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("results"))
	return rv
}


// An array of objects that define the query’s scope.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/searchItems
func (m_ MetadataQuery) SearchItems() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("searchItems"))
	return rv
}


// An array of objects that define the query’s scope.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/searchItems
func (m_ MetadataQuery) SetSearchItems(value objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSearchItems:"), value)
}


// An array containing the search scopes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/searchScopes
func (m_ MetadataQuery) SearchScopes() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("searchScopes"))
	return rv
}


// An array containing the search scopes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/searchScopes
func (m_ MetadataQuery) SetSearchScopes(value objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSearchScopes:"), value)
}


// An array of sort descriptor objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/sortDescriptors
func (m_ MetadataQuery) SortDescriptors() []SortDescriptor {
	rv := objc.Send[[]SortDescriptor](m_.ID, objc.Sel("sortDescriptors"))
	return rv
}


// An array of sort descriptor objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/sortDescriptors
func (m_ MetadataQuery) SetSortDescriptors(value []SortDescriptor) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setSortDescriptors:"), nsArray)
}


// An array of attributes whose values are gathered by the query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/valueListAttributes
func (m_ MetadataQuery) ValueListAttributes() []string {
	rv := objc.Send[[]string](m_.ID, objc.Sel("valueListAttributes"))
	return rv
}


// An array of attributes whose values are gathered by the query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/valueListAttributes
func (m_ MetadataQuery) SetValueListAttributes(value []string) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setValueListAttributes:"), nsArray)
}


// A dictionary containing the value lists generated by the query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery/valueLists
func (m_ MetadataQuery) ValueLists() IDictionary {
	rv := objc.Send[IDictionary](m_.ID, objc.Sel("valueLists"))
	return rv
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



