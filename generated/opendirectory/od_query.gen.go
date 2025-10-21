// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
	RemoveFromRunLoopForMode(inRunLoop unsafe.Pointer, inMode string)
	ResultsAllowingPartialError(inAllowPartialResults bool, outError unsafe.Pointer) unsafe.Pointer
	ScheduleInRunLoopForMode(inRunLoop unsafe.Pointer, inMode string)
	Synchronize()
}

// An object serves as a Cocoa wrapper for an Open Directory query.
//
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
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQuery/init(node:forRecordTypes:attribute:matchType:queryValues:returnAttributes:maximumResults:)
func NewODQueryWithNodeForRecordTypesAttributeMatchTypeQueryValuesReturnAttributesMaximumResultsError(inNode unsafe.Pointer, inRecordTypeOrList objc.ID, inAttribute unsafe.Pointer, inMatchType unsafe.Pointer, inQueryValueOrList objc.ID, inReturnAttributeOrList objc.ID, inMaximumResults int, outError unsafe.Pointer) ODQuery {
	instance := getODQueryClass().Alloc()
	rv := objc.Send[ODQuery](instance.ID, objc.Sel("initWithNode:forRecordTypes:attribute:matchType:queryValues:returnAttributes:maximumResults:error:"), inNode, inRecordTypeOrList, inAttribute, inMatchType, inQueryValueOrList, inReturnAttributeOrList, inMaximumResults, outError)
	rv.Autorelease()
	return rv
}


// Returns an autoreleased query object created with provided parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQuery/queryWithNode:forRecordTypes:attribute:matchType:queryValues:returnAttributes:maximumResults:error:
func (oc _ODQueryClass) QueryWithNodeForRecordTypesAttributeMatchTypeQueryValuesReturnAttributesMaximumResultsError(inNode unsafe.Pointer, inRecordTypeOrList objc.ID, inAttribute unsafe.Pointer, inMatchType unsafe.Pointer, inQueryValueOrList objc.ID, inReturnAttributeOrList objc.ID, inMaximumResults int, outError unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("queryWithNode:forRecordTypes:attribute:matchType:queryValues:returnAttributes:maximumResults:error:"), inNode, inRecordTypeOrList, inAttribute, inMatchType, inQueryValueOrList, inReturnAttributeOrList, inMaximumResults, outError)
	return rv
}

// Removes the query from a specified run loop.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQuery/remove(from:forMode:)
func (o_ ODQuery) RemoveFromRunLoopForMode(inRunLoop unsafe.Pointer, inMode string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("removeFromRunLoop:forMode:"), inRunLoop, objc.String(inMode))
}

// Returns results from a query synchronously.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQuery/resultsAllowingPartial(_:)
func (o_ ODQuery) ResultsAllowingPartialError(inAllowPartialResults bool, outError unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("resultsAllowingPartial:error:"), inAllowPartialResults, outError)
	return rv
}

// Retrieves results from a query asynchronously by scheduling the query in a run loop.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQuery/schedule(in:forMode:)
func (o_ ODQuery) ScheduleInRunLoopForMode(inRunLoop unsafe.Pointer, inMode string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("scheduleInRunLoop:forMode:"), inRunLoop, objc.String(inMode))
}

// Restarts a query, disposing of any results it has obtained.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQuery/synchronize()
func (o_ ODQuery) Synchronize() {
	objc.Send[objc.ID](o_.ID, objc.Sel("synchronize"))
}

// The query’s delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQuery/delegate
func (o_ ODQuery) Delegate() objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The query’s delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQuery/delegate
func (o_ ODQuery) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setDelegate:"), value)
}

// The queue on which asynchronous results are delivered to the delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQuery/operationQueue
func (o_ ODQuery) OperationQueue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("operationQueue"))
	return rv
}


// SetOperationQueue sets the value of the operationQueue property.
// The queue on which asynchronous results are delivered to the delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQuery/operationQueue
func (o_ ODQuery) SetOperationQueue(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setOperationQueue:"), value)
}


