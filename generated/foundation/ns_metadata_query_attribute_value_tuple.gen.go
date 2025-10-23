// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MetadataQueryAttributeValueTuple] class.
var (
	MetadataQueryAttributeValueTupleClass     _MetadataQueryAttributeValueTupleClass
	MetadataQueryAttributeValueTupleClassOnce sync.Once
)

func getMetadataQueryAttributeValueTupleClass() _MetadataQueryAttributeValueTupleClass {
	MetadataQueryAttributeValueTupleClassOnce.Do(func() {
		MetadataQueryAttributeValueTupleClass = _MetadataQueryAttributeValueTupleClass{objc.GetClass("NSMetadataQueryAttributeValueTuple")}
	})
	return MetadataQueryAttributeValueTupleClass
}

type _MetadataQueryAttributeValueTupleClass struct {
	class objc.Class
}

// An interface definition for the [MetadataQueryAttributeValueTuple] class.
type IMetadataQueryAttributeValueTuple interface {
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
	Count() int /* primitive/slice/pointer. */
	SetCount(value int /* primitive/slice/pointer. */)
	Value() unsafe.Pointer
	SetValue(value unsafe.Pointer)
	// methods:
}

// The class represents attribute-value tuples, which are objects that contain the attribute name and value of a metadata attribute.
//
// Attribute-value tuples are returned by objects as the results in the value lists. Each attribute/value tuple contains the attribute name, the value, and the number of instances of that value that exist for the attribute name.


// The class represents attribute-value tuples, which are objects that contain the attribute name and value of a metadata attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQueryAttributeValueTuple
type MetadataQueryAttributeValueTuple struct {
	objectivec.Object
}

// MetadataQueryAttributeValueTupleFrom constructs a [MetadataQueryAttributeValueTuple] from an unsafe.Pointer.
//
// The class represents attribute-value tuples, which are objects that contain the attribute name and value of a metadata attribute.
func MetadataQueryAttributeValueTupleFrom(ptr unsafe.Pointer) MetadataQueryAttributeValueTuple {
	return MetadataQueryAttributeValueTuple{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MetadataQueryAttributeValueTupleClass) Alloc() MetadataQueryAttributeValueTuple {
	rv := objc.Send[MetadataQueryAttributeValueTuple](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MetadataQueryAttributeValueTupleClass) New() MetadataQueryAttributeValueTuple {
	rv := objc.Send[MetadataQueryAttributeValueTuple](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetadataQueryAttributeValueTuple) Init() MetadataQueryAttributeValueTuple {
	rv := objc.Send[MetadataQueryAttributeValueTuple](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetadataQueryAttributeValueTuple) Autorelease() MetadataQueryAttributeValueTuple {
	rv := objc.Send[MetadataQueryAttributeValueTuple](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetadataQueryAttributeValueTuple creates a new MetadataQueryAttributeValueTuple instance.
func NewMetadataQueryAttributeValueTuple() MetadataQueryAttributeValueTuple {
	return getMetadataQueryAttributeValueTupleClass().New()
}



// An array containing hierarchical groups of query results. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/groupedresults
func (m_ MetadataQueryAttributeValueTuple) GroupedResults() IMetadataQueryResultGroup {
	rv := objc.Send[MetadataQueryResultGroup](m_.ID, objc.Sel("groupedResults"))
	return rv
}


// An array containing hierarchical groups of query results. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/groupedresults
func (m_ MetadataQueryAttributeValueTuple) SetGroupedResults(value IMetadataQueryResultGroup) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupedResults:"), value)
}


// The queue on which query result notifications are posted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/operationqueue
func (m_ MetadataQueryAttributeValueTuple) OperationQueue() IOperationQueue {
	rv := objc.Send[OperationQueue](m_.ID, objc.Sel("operationQueue"))
	return rv
}


// The queue on which query result notifications are posted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/operationqueue
func (m_ MetadataQueryAttributeValueTuple) SetOperationQueue(value IOperationQueue) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationQueue:"), value)
}


// The number of results returned by the query. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/resultcount
func (m_ MetadataQueryAttributeValueTuple) ResultCount() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](m_.ID, objc.Sel("resultCount"))
	return rv
}


// The number of results returned by the query. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/resultcount
func (m_ MetadataQueryAttributeValueTuple) SetResultCount(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setResultCount:"), value)
}


// An array containing the query’s results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/results
func (m_ MetadataQueryAttributeValueTuple) Results() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("results"))
	return rv
}


// An array containing the query’s results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/results
func (m_ MetadataQueryAttributeValueTuple) SetResults(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setResults:"), value)
}


// A dictionary containing the value lists generated by the query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/valuelists
func (m_ MetadataQueryAttributeValueTuple) ValueLists() IMetadataQueryAttributeValueTuple {
	rv := objc.Send[MetadataQueryAttributeValueTuple](m_.ID, objc.Sel("valueLists"))
	return rv
}


// A dictionary containing the value lists generated by the query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/valuelists
func (m_ MetadataQueryAttributeValueTuple) SetValueLists(value IMetadataQueryAttributeValueTuple) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValueLists:"), value)
}


// The attribute name for the tuple’s value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataqueryattributevaluetuple/attribute
func (m_ MetadataQueryAttributeValueTuple) Attribute() IString {
	rv := objc.Send[String](m_.ID, objc.Sel("attribute"))
	return rv
}


// The attribute name for the tuple’s value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataqueryattributevaluetuple/attribute
func (m_ MetadataQueryAttributeValueTuple) SetAttribute(value IString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttribute:"), value)
}


// The number of instances of the value for the tuple’s attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataqueryattributevaluetuple/count
func (m_ MetadataQueryAttributeValueTuple) Count() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](m_.ID, objc.Sel("count"))
	return rv
}


// The number of instances of the value for the tuple’s attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataqueryattributevaluetuple/count
func (m_ MetadataQueryAttributeValueTuple) SetCount(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCount:"), value)
}


// The value of the tuple’s attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataqueryattributevaluetuple/value
func (m_ MetadataQueryAttributeValueTuple) Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("value"))
	return rv
}


// The value of the tuple’s attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataqueryattributevaluetuple/value
func (m_ MetadataQueryAttributeValueTuple) SetValue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}



