// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Expression] class.
var (
	ExpressionClass     _ExpressionClass
	ExpressionClassOnce sync.Once
)

func getExpressionClass() _ExpressionClass {
	ExpressionClassOnce.Do(func() {
		ExpressionClass = _ExpressionClass{objc.GetClass("NSExpression")}
	})
	return ExpressionClass
}

type _ExpressionClass struct {
	class objc.Class
}

// An interface definition for the [Expression] class.
type IExpression interface {
	objectivec.IObject
	AllowEvaluation()
	ExpressionValueWithObjectContext(object objc.ID, context unsafe.Pointer) objc.ID
}

// An expression for use in a comparison predicate.
//
// Comparison operations in an derive from two expressions as instances of the class. You create expressions for constant values, key paths, and so on. Generally, anywhere in the class hierarchy where there’s a composite API and subtypes that may only reasonably respond to a subset of that API, invoking a method that doesn’t make sense for that subtype throws an exception.
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

// Alloc allocates a new instance without initialization.
func (ec _ExpressionClass) Alloc() Expression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _ExpressionClass) New() Expression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ Expression) Init() Expression {
	rv := objc.Send[Expression](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ Expression) Autorelease() Expression {
	rv := objc.Send[Expression](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewExpression creates a new Expression instance.
func NewExpression() Expression {
	return getExpressionClass().New()
}


// Forces a securely decoded expression to allow evaluation.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/allowEvaluation()
func (e_ Expression) AllowEvaluation() {
	objc.Send[objc.ID](e_.ID, objc.Sel("allowEvaluation"))
}

// Evaluates an expression using a specified object and context.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/expressionValue(with:context:)
func (e_ Expression) ExpressionValueWithObjectContext(object objc.ID, context unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](e_.ID, objc.Sel("expressionValueWithObject:context:"), object, context)
	return rv
}

// An expression to evalutate if a conditional expression’s predicate evaluates to false.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/false
func (e_ Expression) FalseExpression() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("falseExpression"))
	return rv
}

// The predicate of a subquery expression.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/predicate
func (e_ Expression) Predicate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("predicate"))
	return rv
}

// An expression to evalutate if a conditional expression’s predicate evaluates to true.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/true
func (e_ Expression) TrueExpression() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("trueExpression"))
	return rv
}



