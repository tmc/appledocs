// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Expression] class.
var ExpressionClass objc.Class

func init() {
	ExpressionClass = objc.GetClass("NSExpression")
}

type Expression struct {
	objc.ID
}

func ExpressionFrom(ptr unsafe.Pointer) Expression {
	return Expression{
		ID: objc.ID(ptr),
	}
}


// Evaluates an expression using a specified object and context. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSExpression/expressionValue(with:context:)
func (e_ Expression) ExpressionValueWithObjectContext(object objc.ID, context unsafe.Pointer) objc.ID {
	sel := objc.RegisterName("expressionValueWithObject:context:")
	ret := e_.ID.Send(sel, object, context)
	return ret
}

