// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ODQuery] class.
var (
	ODQueryClass     _ODQueryClass
	ODQueryClassOnce sync.Once
)

func getODQueryClass() _ODQueryClass {
	ODQueryClassOnce.Do(func() {
		ODQueryClass = _ODQueryClass{objc.GetClass("ODQuery")}
	})
	return ODQueryClass
}

type _ODQueryClass struct {
	class objc.Class
}

// An interface definition for the [ODQuery] class.
type IODQuery interface {
	objectivec.IObject
	// properties:
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	OperationQueue() objc.IObject /* cross-framework: OperationQueue */
	SetOperationQueue(value objc.IObject /* cross-framework: OperationQueue */)
	// methods:
	RemoveFromRunLoopForMode(inRunLoop objc.IObject /* cross-framework: RunLoop */, inMode objc.IObject /* cross-framework: NSString */)
	ResultsAllowingPartialError(inAllowPartialResults bool, outError unsafe.Pointer) objc.IObject /* cross-framework: Array */
	ScheduleInRunLoopForMode(inRunLoop objc.IObject /* cross-framework: RunLoop */, inMode objc.IObject /* cross-framework: NSString */)
	Synchronize()
}

// An object serves as a Cocoa wrapper for an Open Directory query.


// An object serves as a Cocoa wrapper for an Open Directory query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQuery
type ODQuery struct {
	objectivec.Object
}

// ODQueryFrom constructs a [ODQuery] from an unsafe.Pointer.
//
// An object serves as a Cocoa wrapper for an Open Directory query.
func ODQueryFrom(ptr unsafe.Pointer) ODQuery {
	return ODQuery{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _ODQueryClass) Alloc() ODQuery {
	rv := objc.Send[ODQuery](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _ODQueryClass) New() ODQuery {
	rv := objc.Send[ODQuery](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ ODQuery) Init() ODQuery {
	rv := objc.Send[ODQuery](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ ODQuery) Autorelease() ODQuery {
	rv := objc.Send[ODQuery](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewODQuery creates a new ODQuery instance.
func NewODQuery() ODQuery {
	return getODQueryClass().New()
}



// Creates a query object with provided parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQuery/init(node:forRecordTypes:attribute:matchType:queryValues:returnAttributes:maximumResults:)
func NewODQueryWithNodeForRecordTypesAttributeMatchTypeQueryValuesReturnAttributesMaximumResultsError(inNode IODNode, inRecordTypeOrList objectivec.IObject, inAttribute ODAttributeType /* typedef */, inMatchType ODMatchType /* typedef */, inQueryValueOrList objectivec.IObject, inReturnAttributeOrList objectivec.IObject, inMaximumResults int, outError unsafe.Pointer) ODQuery {
	instance := getODQueryClass().Alloc()
	rv := objc.Send[ODQuery](instance.ID, objc.Sel("initWithNode:forRecordTypes:attribute:matchType:queryValues:returnAttributes:maximumResults:error:"), inNode, inRecordTypeOrList, inAttribute, inMatchType, inQueryValueOrList, inReturnAttributeOrList, inMaximumResults, outError)
	rv.Autorelease()
	return rv
}



// Returns an autoreleased query object created with provided parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQuery/queryWithNode:forRecordTypes:attribute:matchType:queryValues:returnAttributes:maximumResults:error:
func (oc _ODQueryClass) QueryWithNodeForRecordTypesAttributeMatchTypeQueryValuesReturnAttributesMaximumResultsError(inNode IODNode, inRecordTypeOrList objectivec.IObject, inAttribute ODAttributeType /* typedef */, inMatchType ODMatchType /* typedef */, inQueryValueOrList objectivec.IObject, inReturnAttributeOrList objectivec.IObject, inMaximumResults int, outError unsafe.Pointer) ODQuery {
	rv := objc.Send[ODQuery](objc.ID(oc.class), objc.Sel("queryWithNode:forRecordTypes:attribute:matchType:queryValues:returnAttributes:maximumResults:error:"), inNode, inRecordTypeOrList, inAttribute, inMatchType, inQueryValueOrList, inReturnAttributeOrList, inMaximumResults, outError)
	return rv
}


// Removes the query from a specified run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQuery/remove(from:forMode:)
func (o_ ODQuery) RemoveFromRunLoopForMode(inRunLoop objc.IObject /* cross-framework: RunLoop */, inMode objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("removeFromRunLoop:forMode:"), inRunLoop, inMode)
}


// Returns results from a query synchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQuery/resultsAllowingPartial(_:)
func (o_ ODQuery) ResultsAllowingPartialError(inAllowPartialResults bool, outError unsafe.Pointer) objc.IObject /* cross-framework: Array */ {
	rv := objc.Send[foundation.Array](o_.ID, objc.Sel("resultsAllowingPartial:error:"), inAllowPartialResults, outError)
	return rv
}


// Retrieves results from a query asynchronously by scheduling the query in a run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQuery/schedule(in:forMode:)
func (o_ ODQuery) ScheduleInRunLoopForMode(inRunLoop objc.IObject /* cross-framework: RunLoop */, inMode objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("scheduleInRunLoop:forMode:"), inRunLoop, inMode)
}


// Restarts a query, disposing of any results it has obtained.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQuery/synchronize()
func (o_ ODQuery) Synchronize() {
	objc.Send[objc.ID](o_.ID, objc.Sel("synchronize"))
}


// The query’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQuery/delegate
func (o_ ODQuery) Delegate() objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("delegate"))
	return rv
}


// The query’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQuery/delegate
func (o_ ODQuery) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setDelegate:"), value)
}


// The queue on which asynchronous results are delivered to the delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQuery/operationQueue
func (o_ ODQuery) OperationQueue() objc.IObject /* cross-framework: OperationQueue */ {
	rv := objc.Send[foundation.OperationQueue](o_.ID, objc.Sel("operationQueue"))
	return rv
}


// The queue on which asynchronous results are delivered to the delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQuery/operationQueue
func (o_ ODQuery) SetOperationQueue(value objc.IObject /* cross-framework: OperationQueue */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setOperationQueue:"), value)
}


