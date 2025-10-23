// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [FetchRequestExpression] class.
var (
	FetchRequestExpressionClass     _FetchRequestExpressionClass
	FetchRequestExpressionClassOnce sync.Once
)

func getFetchRequestExpressionClass() _FetchRequestExpressionClass {
	FetchRequestExpressionClassOnce.Do(func() {
		FetchRequestExpressionClass = _FetchRequestExpressionClass{objc.GetClass("NSFetchRequestExpression")}
	})
	return FetchRequestExpressionClass
}

type _FetchRequestExpressionClass struct {
	class objc.Class
}

// An interface definition for the [FetchRequestExpression] class.
type IFetchRequestExpression interface {
	IExpression
	// properties:
	ContextExpression() objc.IObject /* cross-framework: Expression */
	CountOnlyRequest() bool /* primitive/slice/pointer. */
	RequestExpression() objc.IObject /* cross-framework: Expression */
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
	IsCountOnlyRequest() bool /* primitive/slice/pointer. */
	SetIsCountOnlyRequest(value bool /* primitive/slice/pointer. */)
	NSFetchRequestExpressionType() unsafe.Pointer
	// methods:
}

// An expression that evaluates the result of a fetch request on a managed object context.
//
// inherits from , which provides most of the basic behavior. The first argument must be an expression which evaluates to an object, and the second must be an expression which evaluates to an object. If you simply want the count for the request, the argument should be .


// An expression that evaluates the result of a fetch request on a managed object context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequestExpression
type FetchRequestExpression struct {
	Expression
}

// FetchRequestExpressionFrom constructs a [FetchRequestExpression] from an unsafe.Pointer.
//
// An expression that evaluates the result of a fetch request on a managed object context.
func FetchRequestExpressionFrom(ptr unsafe.Pointer) FetchRequestExpression {
	return FetchRequestExpression{
		Expression: ExpressionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (fc _FetchRequestExpressionClass) Alloc() FetchRequestExpression {
	rv := objc.Send[FetchRequestExpression](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FetchRequestExpressionClass) New() FetchRequestExpression {
	rv := objc.Send[FetchRequestExpression](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FetchRequestExpression) Init() FetchRequestExpression {
	rv := objc.Send[FetchRequestExpression](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FetchRequestExpression) Autorelease() FetchRequestExpression {
	rv := objc.Send[FetchRequestExpression](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFetchRequestExpression creates a new FetchRequestExpression instance.
func NewFetchRequestExpression() FetchRequestExpression {
	return getFetchRequestExpressionClass().New()
}



// Returns an expression which will evaluate to the result of executing a fetch request on a context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequestExpression/expression(forFetch:context:countOnly:)
func (fc _FetchRequestExpressionClass) ExpressionForFetchContextCountOnly(fetch objc.IObject /* cross-framework Expression */, context objc.IObject /* cross-framework Expression */, countFlag bool /* primitive/slice/pointer. */) objc.IObject /* cross-framework: Expression */ {
	rv := objc.Send[Expression](objc.ID(fc.class), objc.Sel("expressionForFetch:context:countOnly:"), fetch, context, countFlag)
	return rv
}


// The expression for the receiver’s managed object context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequestExpression/contextExpression
func (f_ FetchRequestExpression) ContextExpression() objc.IObject /* cross-framework: Expression */ {
	rv := objc.Send[Expression](f_.ID, objc.Sel("contextExpression"))
	return rv
}


// Returns a Boolean value that indicates whether the receiver represents a count-only fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequestExpression/isCountOnlyRequest
func (f_ FetchRequestExpression) CountOnlyRequest() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](f_.ID, objc.Sel("countOnlyRequest"))
	return rv
}


// The expression for the receiver’s fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequestExpression/requestExpression
func (f_ FetchRequestExpression) RequestExpression() objc.IObject /* cross-framework: Expression */ {
	rv := objc.Send[Expression](f_.ID, objc.Sel("requestExpression"))
	return rv
}


// An array of persistent stores specified for the fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/affectedstores
func (f_ FetchRequestExpression) AffectedStores() IPersistentStore {
	rv := objc.Send[PersistentStore](f_.ID, objc.Sel("affectedStores"))
	return rv
}


// An array of persistent stores specified for the fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/affectedstores
func (f_ FetchRequestExpression) SetAffectedStores(value IPersistentStore) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setAffectedStores:"), value)
}


// The batch size of the objects specified in the fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/fetchbatchsize
func (f_ FetchRequestExpression) FetchBatchSize() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](f_.ID, objc.Sel("fetchBatchSize"))
	return rv
}


// The batch size of the objects specified in the fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/fetchbatchsize
func (f_ FetchRequestExpression) SetFetchBatchSize(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFetchBatchSize:"), value)
}


// The fetch limit of the fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/fetchlimit
func (f_ FetchRequestExpression) FetchLimit() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](f_.ID, objc.Sel("fetchLimit"))
	return rv
}


// The fetch limit of the fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/fetchlimit
func (f_ FetchRequestExpression) SetFetchLimit(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFetchLimit:"), value)
}


// The fetch offset of the fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/fetchoffset
func (f_ FetchRequestExpression) FetchOffset() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](f_.ID, objc.Sel("fetchOffset"))
	return rv
}


// The fetch offset of the fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/fetchoffset
func (f_ FetchRequestExpression) SetFetchOffset(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFetchOffset:"), value)
}


// The predicate of the fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/predicate
func (f_ FetchRequestExpression) Predicate() objc.IObject /* cross-framework: Predicate */ {
	rv := objc.Send[Predicate](f_.ID, objc.Sel("predicate"))
	return rv
}


// The predicate of the fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/predicate
func (f_ FetchRequestExpression) SetPredicate(value objc.IObject /* cross-framework: Predicate */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setPredicate:"), value)
}


// Returns a Boolean value that indicates whether the receiver represents a count-only fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequestexpression/iscountonlyrequest
func (f_ FetchRequestExpression) IsCountOnlyRequest() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](f_.ID, objc.Sel("isCountOnlyRequest"))
	return rv
}


// Returns a Boolean value that indicates whether the receiver represents a count-only fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequestexpression/iscountonlyrequest
func (f_ FetchRequestExpression) SetIsCountOnlyRequest(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsCountOnlyRequest:"), value)
}


// This constant specifies the fetch request expression type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequestexpressiontype
func (f_ FetchRequestExpression) NSFetchRequestExpressionType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("NSFetchRequestExpressionType"))
	return rv
}



