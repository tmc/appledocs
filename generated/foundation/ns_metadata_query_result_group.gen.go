// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MetadataQueryResultGroup] class.
var (
	MetadataQueryResultGroupClass     _MetadataQueryResultGroupClass
	MetadataQueryResultGroupClassOnce sync.Once
)

func getMetadataQueryResultGroupClass() _MetadataQueryResultGroupClass {
	MetadataQueryResultGroupClassOnce.Do(func() {
		MetadataQueryResultGroupClass = _MetadataQueryResultGroupClass{objc.GetClass("NSMetadataQueryResultGroup")}
	})
	return MetadataQueryResultGroupClass
}

type _MetadataQueryResultGroupClass struct {
	class objc.Class
}

// An interface definition for the [MetadataQueryResultGroup] class.
type IMetadataQueryResultGroup interface {
	objectivec.IObject
	// properties:
	GroupedResults() IMetadataQueryResultGroup
	SetGroupedResults(value IMetadataQueryResultGroup)
	OperationQueue() IOperationQueue
	SetOperationQueue(value IOperationQueue)
	ResultCount() int /* primitive/slice/pointer. */
	SetResultCount(value int /* primitive/slice/pointer. */)
	Results() unsafe.Pointer
	SetResults(value unsafe.Pointer)
	ValueLists() IMetadataQueryAttributeValueTuple
	SetValueLists(value IMetadataQueryAttributeValueTuple)
	Attribute() IString
	SetAttribute(value IString)
	Subgroups() IMetadataQueryResultGroup
	SetSubgroups(value IMetadataQueryResultGroup)
	Value() unsafe.Pointer
	SetValue(value unsafe.Pointer)
	// methods:
}

// The class represents a collection of grouped attribute results returned by an object.


// The class represents a collection of grouped attribute results returned by an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQueryResultGroup
type MetadataQueryResultGroup struct {
	objectivec.Object
}

// MetadataQueryResultGroupFrom constructs a [MetadataQueryResultGroup] from an unsafe.Pointer.
//
// The class represents a collection of grouped attribute results returned by an object.
func MetadataQueryResultGroupFrom(ptr unsafe.Pointer) MetadataQueryResultGroup {
	return MetadataQueryResultGroup{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MetadataQueryResultGroupClass) Alloc() MetadataQueryResultGroup {
	rv := objc.Send[MetadataQueryResultGroup](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MetadataQueryResultGroupClass) New() MetadataQueryResultGroup {
	rv := objc.Send[MetadataQueryResultGroup](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetadataQueryResultGroup) Init() MetadataQueryResultGroup {
	rv := objc.Send[MetadataQueryResultGroup](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetadataQueryResultGroup) Autorelease() MetadataQueryResultGroup {
	rv := objc.Send[MetadataQueryResultGroup](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetadataQueryResultGroup creates a new MetadataQueryResultGroup instance.
func NewMetadataQueryResultGroup() MetadataQueryResultGroup {
	return getMetadataQueryResultGroupClass().New()
}



// An array containing hierarchical groups of query results. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/groupedresults
func (m_ MetadataQueryResultGroup) GroupedResults() IMetadataQueryResultGroup {
	rv := objc.Send[MetadataQueryResultGroup](m_.ID, objc.Sel("groupedResults"))
	return rv
}


// An array containing hierarchical groups of query results. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/groupedresults
func (m_ MetadataQueryResultGroup) SetGroupedResults(value IMetadataQueryResultGroup) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupedResults:"), value)
}


// The queue on which query result notifications are posted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/operationqueue
func (m_ MetadataQueryResultGroup) OperationQueue() IOperationQueue {
	rv := objc.Send[OperationQueue](m_.ID, objc.Sel("operationQueue"))
	return rv
}


// The queue on which query result notifications are posted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/operationqueue
func (m_ MetadataQueryResultGroup) SetOperationQueue(value IOperationQueue) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationQueue:"), value)
}


// The number of results returned by the query. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/resultcount
func (m_ MetadataQueryResultGroup) ResultCount() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](m_.ID, objc.Sel("resultCount"))
	return rv
}


// The number of results returned by the query. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/resultcount
func (m_ MetadataQueryResultGroup) SetResultCount(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setResultCount:"), value)
}


// An array containing the query’s results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/results
func (m_ MetadataQueryResultGroup) Results() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("results"))
	return rv
}


// An array containing the query’s results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/results
func (m_ MetadataQueryResultGroup) SetResults(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setResults:"), value)
}


// A dictionary containing the value lists generated by the query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/valuelists
func (m_ MetadataQueryResultGroup) ValueLists() IMetadataQueryAttributeValueTuple {
	rv := objc.Send[MetadataQueryAttributeValueTuple](m_.ID, objc.Sel("valueLists"))
	return rv
}


// A dictionary containing the value lists generated by the query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/valuelists
func (m_ MetadataQueryResultGroup) SetValueLists(value IMetadataQueryAttributeValueTuple) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValueLists:"), value)
}


// The result group’s attribute name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataqueryresultgroup/attribute
func (m_ MetadataQueryResultGroup) Attribute() IString {
	rv := objc.Send[String](m_.ID, objc.Sel("attribute"))
	return rv
}


// The result group’s attribute name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataqueryresultgroup/attribute
func (m_ MetadataQueryResultGroup) SetAttribute(value IString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttribute:"), value)
}


// An array containing the result group’s subgroups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataqueryresultgroup/subgroups
func (m_ MetadataQueryResultGroup) Subgroups() IMetadataQueryResultGroup {
	rv := objc.Send[MetadataQueryResultGroup](m_.ID, objc.Sel("subgroups"))
	return rv
}


// An array containing the result group’s subgroups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataqueryresultgroup/subgroups
func (m_ MetadataQueryResultGroup) SetSubgroups(value IMetadataQueryResultGroup) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSubgroups:"), value)
}


// The result group’s value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataqueryresultgroup/value
func (m_ MetadataQueryResultGroup) Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("value"))
	return rv
}


// The result group’s value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataqueryresultgroup/value
func (m_ MetadataQueryResultGroup) SetValue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}



