// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ExpressionDescription] class.
var (
	ExpressionDescriptionClass     _ExpressionDescriptionClass
	ExpressionDescriptionClassOnce sync.Once
)

func getExpressionDescriptionClass() _ExpressionDescriptionClass {
	ExpressionDescriptionClassOnce.Do(func() {
		ExpressionDescriptionClass = _ExpressionDescriptionClass{objc.GetClass("NSExpressionDescription")}
	})
	return ExpressionDescriptionClass
}

type _ExpressionDescriptionClass struct {
	class objc.Class
}

// An interface definition for the [ExpressionDescription] class.
type IExpressionDescription interface {
	IPropertyDescription
	// properties:
	Expression() objc.IObject /* cross-framework: Expression */
	SetExpression(value objc.IObject /* cross-framework: Expression */)
	ExpressionResultType() AttributeType
	SetExpressionResultType(value AttributeType)
	Properties() IPropertyDescription
	SetProperties(value IPropertyDescription)
	ResultType() AttributeType
	SetResultType(value AttributeType)
	AffectedStores() IPersistentStore
	SetAffectedStores(value IPersistentStore)
	FetchBatchSize() int /* primitive/slice/pointer. */
	SetFetchBatchSize(value int /* primitive/slice/pointer. */)
	FetchLimit() int /* primitive/slice/pointer. */
	SetFetchLimit(value int /* primitive/slice/pointer. */)
	FetchOffset() int /* primitive/slice/pointer. */
	SetFetchOffset(value int /* primitive/slice/pointer. */)
	Predicate() objc.IObject /* cross-framework: Predicate */
	SetPredicate(value objc.IObject /* cross-framework: Predicate */)
	PropertiesToFetch() unsafe.Pointer
	SetPropertiesToFetch(value unsafe.Pointer)
	// methods:
}

// An object that describes an expression to include with a fetch request.
//
// An expression description describes a value that a fetch request returns, which doesn’t appear as an attribute or relationship on an entity. For example, expressions can aggregate data, or transform an attribute’s value. You add expression descriptions to a fetch request using the method.


// An object that describes an expression to include with a fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSExpressionDescription
type ExpressionDescription struct {
	PropertyDescription
}

// ExpressionDescriptionFrom constructs a [ExpressionDescription] from an unsafe.Pointer.
//
// An object that describes an expression to include with a fetch request.
func ExpressionDescriptionFrom(ptr unsafe.Pointer) ExpressionDescription {
	return ExpressionDescription{
		PropertyDescription: PropertyDescriptionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ec _ExpressionDescriptionClass) Alloc() ExpressionDescription {
	rv := objc.Send[ExpressionDescription](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _ExpressionDescriptionClass) New() ExpressionDescription {
	rv := objc.Send[ExpressionDescription](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ ExpressionDescription) Init() ExpressionDescription {
	rv := objc.Send[ExpressionDescription](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ ExpressionDescription) Autorelease() ExpressionDescription {
	rv := objc.Send[ExpressionDescription](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewExpressionDescription creates a new ExpressionDescription instance.
func NewExpressionDescription() ExpressionDescription {
	return getExpressionDescriptionClass().New()
}



// The expression to evaluate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSExpressionDescription/expression
func (e_ ExpressionDescription) Expression() objc.IObject /* cross-framework: Expression */ {
	rv := objc.Send[Expression](e_.ID, objc.Sel("expression"))
	return rv
}


// The expression to evaluate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSExpressionDescription/expression
func (e_ ExpressionDescription) SetExpression(value objc.IObject /* cross-framework: Expression */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setExpression:"), value)
}


// The attribute type of the expression’s result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSExpressionDescription/expressionResultType
func (e_ ExpressionDescription) ExpressionResultType() AttributeType {
	rv := objc.Send[AttributeType](e_.ID, objc.Sel("expressionResultType"))
	return rv
}


// The attribute type of the expression’s result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSExpressionDescription/expressionResultType
func (e_ ExpressionDescription) SetExpressionResultType(value AttributeType) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setExpressionResultType:"), value)
}


// An array containing the properties of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/properties
func (e_ ExpressionDescription) Properties() IPropertyDescription {
	rv := objc.Send[PropertyDescription](e_.ID, objc.Sel("properties"))
	return rv
}


// An array containing the properties of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/properties
func (e_ ExpressionDescription) SetProperties(value IPropertyDescription) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setProperties:"), value)
}


// The attribute type of the expression’s result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsexpressiondescription/resulttype
func (e_ ExpressionDescription) ResultType() AttributeType {
	rv := objc.Send[AttributeType](e_.ID, objc.Sel("resultType"))
	return rv
}


// The attribute type of the expression’s result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsexpressiondescription/resulttype
func (e_ ExpressionDescription) SetResultType(value AttributeType) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setResultType:"), value)
}


// An array of persistent stores specified for the fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/affectedstores
func (e_ ExpressionDescription) AffectedStores() IPersistentStore {
	rv := objc.Send[PersistentStore](e_.ID, objc.Sel("affectedStores"))
	return rv
}


// An array of persistent stores specified for the fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/affectedstores
func (e_ ExpressionDescription) SetAffectedStores(value IPersistentStore) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setAffectedStores:"), value)
}


// The batch size of the objects specified in the fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/fetchbatchsize
func (e_ ExpressionDescription) FetchBatchSize() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](e_.ID, objc.Sel("fetchBatchSize"))
	return rv
}


// The batch size of the objects specified in the fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/fetchbatchsize
func (e_ ExpressionDescription) SetFetchBatchSize(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setFetchBatchSize:"), value)
}


// The fetch limit of the fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/fetchlimit
func (e_ ExpressionDescription) FetchLimit() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](e_.ID, objc.Sel("fetchLimit"))
	return rv
}


// The fetch limit of the fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/fetchlimit
func (e_ ExpressionDescription) SetFetchLimit(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setFetchLimit:"), value)
}


// The fetch offset of the fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/fetchoffset
func (e_ ExpressionDescription) FetchOffset() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](e_.ID, objc.Sel("fetchOffset"))
	return rv
}


// The fetch offset of the fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/fetchoffset
func (e_ ExpressionDescription) SetFetchOffset(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setFetchOffset:"), value)
}


// The predicate of the fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/predicate
func (e_ ExpressionDescription) Predicate() objc.IObject /* cross-framework: Predicate */ {
	rv := objc.Send[Predicate](e_.ID, objc.Sel("predicate"))
	return rv
}


// The predicate of the fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/predicate
func (e_ ExpressionDescription) SetPredicate(value objc.IObject /* cross-framework: Predicate */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setPredicate:"), value)
}


// A collection of either property descriptions or string property names that specify which properties should be returned by the fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/propertiestofetch
func (e_ ExpressionDescription) PropertiesToFetch() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("propertiesToFetch"))
	return rv
}


// A collection of either property descriptions or string property names that specify which properties should be returned by the fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/propertiestofetch
func (e_ ExpressionDescription) SetPropertiesToFetch(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setPropertiesToFetch:"), value)
}



