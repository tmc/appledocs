// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [FetchRequestExpression] class.
var (
	fetchRequestExpressionClass     _FetchRequestExpressionClass
	fetchRequestExpressionClassOnce sync.Once
)

func getFetchRequestExpressionClass() _FetchRequestExpressionClass {
	fetchRequestExpressionClassOnce.Do(func() {
		fetchRequestExpressionClass = _FetchRequestExpressionClass{objc.GetClass("NSFetchRequestExpression")}
	})
	return fetchRequestExpressionClass
}

type _FetchRequestExpressionClass struct {
	class objc.Class
}

// An interface definition for the [FetchRequestExpression] class.
type IFetchRequestExpression interface {
	foundation.IExpression
}

// An expression that evaluates the result of a fetch request on a managed object context.
//
// inherits from , which provides most of the basic behavior. The first argument must be an expression which evaluates to an object, and the second must be an expression which evaluates to an object. If you simply want the count for the request, the argument should be .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequestExpression
type FetchRequestExpression struct {
	foundation.Expression
}

// FetchRequestExpressionFrom constructs a [FetchRequestExpression] from an unsafe.Pointer.
//
// An expression that evaluates the result of a fetch request on a managed object context.
func FetchRequestExpressionFrom(ptr unsafe.Pointer) FetchRequestExpression {
	return FetchRequestExpression{
		Expression: foundation.ExpressionFrom(ptr),
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
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequestExpression/expression(forFetch:context:countOnly:)
func (fc _FetchRequestExpressionClass) ExpressionForFetchContextCountOnly(fetch unsafe.Pointer, context unsafe.Pointer, countFlag bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("expressionForFetch:context:countOnly:"), fetch, context, countFlag)
	return rv
}

// The expression for the receiver’s managed object context.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequestExpression/contextExpression
func (f_ FetchRequestExpression) ContextExpression() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("contextExpression"))
	return rv
}

// Returns a Boolean value that indicates whether the receiver represents a count-only fetch request.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequestExpression/isCountOnlyRequest
func (f_ FetchRequestExpression) CountOnlyRequest() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("countOnlyRequest"))
	return rv
}

// The expression for the receiver’s fetch request.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequestExpression/requestExpression
func (f_ FetchRequestExpression) RequestExpression() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("requestExpression"))
	return rv
}



