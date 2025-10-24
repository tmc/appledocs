// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ODQuery */


/* debug [class_header]: Header for ODQuery */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ODQuery */
// An interface definition for the [ODQuery] class.
type IODQuery interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ODQuery */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	OperationQueue() foundation.OperationQueue
	SetOperationQueue(value foundation.OperationQueue)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ODQuery */
	// methods:
	RemoveFromRunLoopForMode(inRunLoop foundation.RunLoop, inMode objc.IObject /* cross-framework: NSString */)
	ResultsAllowingPartialError(inAllowPartialResults bool, outError unsafe.Pointer) foundation.Array
	ScheduleInRunLoopForMode(inRunLoop foundation.RunLoop, inMode objc.IObject /* cross-framework: NSString */)
	Synchronize()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ODQuery */
// Alloc allocates a new instance without initialization.
func (oc _ODQueryClass) Alloc() ODQuery {
	rv := objc.Send[ODQuery](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ODQuery */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ODQuery */

// Creates a query object with provided parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQuery/init(node:forRecordTypes:attribute:matchType:queryValues:returnAttributes:maximumResults:)
func NewODQueryWithNodeForRecordTypesAttributeMatchTypeQueryValuesReturnAttributesMaximumResultsError(inNode IODNode, inRecordTypeOrList objc.IObject, inAttribute ODAttributeType /* typedef */, inMatchType ODMatchType /* typedef */, inQueryValueOrList objc.IObject, inReturnAttributeOrList objc.IObject, inMaximumResults int, outError unsafe.Pointer) ODQuery {
	instance := getODQueryClass().Alloc()
	rv := objc.Send[ODQuery](instance.ID, objc.Sel("initWithNode:forRecordTypes:attribute:matchType:queryValues:returnAttributes:maximumResults:error:"), inNode, inRecordTypeOrList, inAttribute, inMatchType, inQueryValueOrList, inReturnAttributeOrList, inMaximumResults, outError)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewODQueryWithNodeForRecordTypesAttributeMatchTypeQueryValuesReturnAttributesMaximumResultsError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ODQuery */

// Returns an autoreleased query object created with provided parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQuery/queryWithNode:forRecordTypes:attribute:matchType:queryValues:returnAttributes:maximumResults:error:
func (oc _ODQueryClass) QueryWithNodeForRecordTypesAttributeMatchTypeQueryValuesReturnAttributesMaximumResultsError(inNode IODNode, inRecordTypeOrList objc.IObject, inAttribute ODAttributeType /* typedef */, inMatchType ODMatchType /* typedef */, inQueryValueOrList objc.IObject, inReturnAttributeOrList objc.IObject, inMaximumResults int, outError unsafe.Pointer) ODQuery {
	rv := objc.Send[ODQuery](objc.ID(oc.class), objc.Sel("queryWithNode:forRecordTypes:attribute:matchType:queryValues:returnAttributes:maximumResults:error:"), inNode, inRecordTypeOrList, inAttribute, inMatchType, inQueryValueOrList, inReturnAttributeOrList, inMaximumResults, outError)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=QueryWithNodeForRecordTypesAttributeMatchTypeQueryValuesReturnAttributesMaximumResultsError) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ODQuery */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ODQuery */

// Removes the query from a specified run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQuery/remove(from:forMode:)
func (o_ ODQuery) RemoveFromRunLoopForMode(inRunLoop foundation.RunLoop, inMode objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("removeFromRunLoop:forMode:"), inRunLoop, inMode)
}/* debug [instance_methods/method]: RemoveFromRunLoopForMode */


// Returns results from a query synchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQuery/resultsAllowingPartial(_:)
func (o_ ODQuery) ResultsAllowingPartialError(inAllowPartialResults bool, outError unsafe.Pointer) foundation.Array {
	rv := objc.Send[foundation.Array](o_.ID, objc.Sel("resultsAllowingPartial:error:"), inAllowPartialResults, outError)
	return rv
}/* debug [instance_methods/method]: ResultsAllowingPartialError */


// Retrieves results from a query asynchronously by scheduling the query in a run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQuery/schedule(in:forMode:)
func (o_ ODQuery) ScheduleInRunLoopForMode(inRunLoop foundation.RunLoop, inMode objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("scheduleInRunLoop:forMode:"), inRunLoop, inMode)
}/* debug [instance_methods/method]: ScheduleInRunLoopForMode */


// Restarts a query, disposing of any results it has obtained.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQuery/synchronize()
func (o_ ODQuery) Synchronize() {
	objc.Send[objc.ID](o_.ID, objc.Sel("synchronize"))
}/* debug [instance_methods/method]: Synchronize */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ODQuery */

// The query’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQuery/delegate
func (o_ ODQuery) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The query’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQuery/delegate
func (o_ ODQuery) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The queue on which asynchronous results are delivered to the delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQuery/operationQueue
func (o_ ODQuery) OperationQueue() foundation.OperationQueue {
	rv := objc.Send[foundation.OperationQueue](o_.ID, objc.Sel("operationQueue"))
	return rv
}/* debug [instance_properties/getter]: operationQueue */


// The queue on which asynchronous results are delivered to the delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQuery/operationQueue
func (o_ ODQuery) SetOperationQueue(value foundation.OperationQueue) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setOperationQueue:"), value)
}/* debug [instance_properties/setter]: operationQueue */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ODQuery */


