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
	ResultAtIndex(idx uint) objc.ID
	Attribute() string
	ResultCount() uint
	Results() objc.ID
	Subgroups() []MetadataQueryResultGroup
	Value() objc.ID
	GroupedResults() NSMetadataQueryResultGroup
	SetGroupedResults(value IMetadataQueryResultGroup)
	OperationQueue() NSOperationQueue
	SetOperationQueue(value IOperationQueue)
	ValueLists() NSMetadataQueryAttributeValueTuple
	SetValueLists(value IMetadataQueryAttributeValueTuple)
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




// Returns the query result at a specific index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQueryResultGroup/result(at:)

func (m_ MetadataQueryResultGroup) ResultAtIndex(idx uint) objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("resultAtIndex:"), idx)
	return rv
}


// The result group’s attribute name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQueryResultGroup/attribute

func (m_ MetadataQueryResultGroup) Attribute() string {
	rv := objc.Send[string](m_.ID, objc.Sel("attribute"))
	return rv
}


// The number of results returned by the result group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQueryResultGroup/resultCount

func (m_ MetadataQueryResultGroup) ResultCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("resultCount"))
	return rv
}


// An array containing the result group’s result objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQueryResultGroup/results

func (m_ MetadataQueryResultGroup) Results() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("results"))
	return rv
}


// An array containing the result group’s subgroups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQueryResultGroup/subgroups

func (m_ MetadataQueryResultGroup) Subgroups() []MetadataQueryResultGroup {
	rv := objc.Send[[]MetadataQueryResultGroup](m_.ID, objc.Sel("subgroups"))
	return rv
}


// The result group’s value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQueryResultGroup/value

func (m_ MetadataQueryResultGroup) Value() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("value"))
	return rv
}


// An array containing hierarchical groups of query results. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/groupedresults

func (m_ MetadataQueryResultGroup) GroupedResults() NSMetadataQueryResultGroup {
	rv := objc.Send[NSMetadataQueryResultGroup](m_.ID, objc.Sel("groupedResults"))
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

func (m_ MetadataQueryResultGroup) OperationQueue() NSOperationQueue {
	rv := objc.Send[NSOperationQueue](m_.ID, objc.Sel("operationQueue"))
	return rv
}


// The queue on which query result notifications are posted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/operationqueue

func (m_ MetadataQueryResultGroup) SetOperationQueue(value IOperationQueue) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationQueue:"), value)
}


// A dictionary containing the value lists generated by the query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/valuelists

func (m_ MetadataQueryResultGroup) ValueLists() NSMetadataQueryAttributeValueTuple {
	rv := objc.Send[NSMetadataQueryAttributeValueTuple](m_.ID, objc.Sel("valueLists"))
	return rv
}


// A dictionary containing the value lists generated by the query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/valuelists

func (m_ MetadataQueryResultGroup) SetValueLists(value IMetadataQueryAttributeValueTuple) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValueLists:"), value)
}



