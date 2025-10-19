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

// Returns an expression which will evaluate to the result of executing a fetch request on a context. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequestExpression/expression(forFetch:context:countOnly:)
func (fc _FetchRequestExpressionClass) ExpressionForFetchContextCountOnly(fetch unsafe.Pointer, context unsafe.Pointer, countFlag bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("expressionForFetch:context:countOnly:"), fetch, context, countFlag)
	return rv
}


