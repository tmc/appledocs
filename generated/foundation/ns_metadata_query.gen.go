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
