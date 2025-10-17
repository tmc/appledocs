// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Expression] class.
var expressionClass = _ExpressionClass{objc.GetClass("NSExpression")}

type _ExpressionClass struct {
	class objc.Class
}

// An expression for use in a comparison predicate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression

type Expression struct {
	objectivec.Object
}

// ExpressionFrom constructs a [Expression] from an unsafe.Pointer.
//
// An expression for use in a comparison predicate.
func ExpressionFrom(ptr unsafe.Pointer) Expression {
	return Expression{objectivec.Object{objc.ID(ptr)}}
}

// Evaluates an expression using a specified object and context. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/expressionValue(with:context:)
func (e_ Expression) ExpressionValueWithObjectContext(object objc.ID, context unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](e_.ID, objc.Sel("expressionValueWithObject:context:"), object, context)
	return rv
}


