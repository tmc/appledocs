// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [FetchRequestExpression] class.
var fetchRequestExpressionClass = _FetchRequestExpressionClass{objc.GetClass("NSFetchRequestExpression")}

type _FetchRequestExpressionClass struct {
	class objc.Class
}

// An interface definition for the [FetchRequestExpression] class.
type IFetchRequestExpression interface {
	IExpression
}

// An expression that evaluates the result of a fetch request on a managed object context. [Full Topic]
//
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

// New creates and returns a new instance with a +1 retain count.
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
	return fetchRequestExpressionClass.New()
}


// Returns an expression which will evaluate to the result of executing a fetch request on a context. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequestExpression/expression(forFetch:context:countOnly:)
func (fc _FetchRequestExpressionClass) ExpressionForFetchContextCountOnly(fetch unsafe.Pointer, context unsafe.Pointer, countFlag bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("expressionForFetch:context:countOnly:"), fetch, context, countFlag)
	return rv
}


