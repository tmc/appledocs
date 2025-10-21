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

// The collection of expressions in an aggregate expression, or the collection element of a subquery expression.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/collection
func (e_ Expression) Collection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("collection"))
	return rv
}


// SetCollection sets the value of the collection property.
// The collection of expressions in an aggregate expression, or the collection element of a subquery expression.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/collection
func (e_ Expression) SetCollection(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setCollection:"), value)
}

// An expression to evalutate if a conditional expression’s predicate evaluates to false.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/false
func (e_ Expression) `false`() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("`false`"))
	return rv
}


// Set`false` sets the value of the `false` property.
// An expression to evalutate if a conditional expression’s predicate evaluates to false.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/false
func (e_ Expression) Set`false`(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("set`false`:"), value)
}

// The right expression of an aggregate expression.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/right
func (e_ Expression) Right() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("right"))
	return rv
}


// SetRight sets the value of the right property.
// The right expression of an aggregate expression.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/right
func (e_ Expression) SetRight(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setRight:"), value)
}

// The constant value of the expression.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/constantvalue
func (e_ Expression) ConstantValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("constantValue"))
	return rv
}


// SetConstantValue sets the value of the constantValue property.
// The constant value of the expression.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/constantvalue
func (e_ Expression) SetConstantValue(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setConstantValue:"), value)
}

// The operand for the expression.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/operand
func (e_ Expression) Operand() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("operand"))
	return rv
}


// SetOperand sets the value of the operand property.
// The operand for the expression.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/operand
func (e_ Expression) SetOperand(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setOperand:"), value)
}

// The function for the expression.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/function
func (e_ Expression) Function() string {
	rv := objc.Send[string](e_.ID, objc.Sel("function"))
	return rv
}


// SetFunction sets the value of the function property.
// The function for the expression.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/function
func (e_ Expression) SetFunction(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setFunction:"), objc.String(value))
}

// The left expression of an aggregate expression.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/left
func (e_ Expression) Left() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("left"))
	return rv
}


// SetLeft sets the value of the left property.
// The left expression of an aggregate expression.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/left
func (e_ Expression) SetLeft(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setLeft:"), value)
}

// The key path for the expression.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/keypath
func (e_ Expression) KeyPath() string {
	rv := objc.Send[string](e_.ID, objc.Sel("keyPath"))
	return rv
}


// SetKeyPath sets the value of the keyPath property.
// The key path for the expression.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/keypath
func (e_ Expression) SetKeyPath(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setKeyPath:"), objc.String(value))
}

// The expression type for the expression.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/expressiontype-swift.property
func (e_ Expression) ExpressionType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("expressionType"))
	return rv
}


// SetExpressionType sets the value of the expressionType property.
// The expression type for the expression.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/expressiontype-swift.property
func (e_ Expression) SetExpressionType(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setExpressionType:"), value)
}

// The variable for the expression.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/variable
func (e_ Expression) Variable() string {
	rv := objc.Send[string](e_.ID, objc.Sel("variable"))
	return rv
}


// SetVariable sets the value of the variable property.
// The variable for the expression.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/variable
func (e_ Expression) SetVariable(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setVariable:"), objc.String(value))
}

// The arguments for the expression.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/arguments
func (e_ Expression) Arguments() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("arguments"))
	return rv
}


// SetArguments sets the value of the arguments property.
// The arguments for the expression.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/arguments
func (e_ Expression) SetArguments(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setArguments:"), value)
}

// An expression to evalutate if a conditional expression’s predicate evaluates to true.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/true
func (e_ Expression) `true`() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("`true`"))
	return rv
}


// Set`true` sets the value of the `true` property.
// An expression to evalutate if a conditional expression’s predicate evaluates to true.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/true
func (e_ Expression) Set`true`(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("set`true`:"), value)
}

// The block that executes to evaluate the expression.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/expressionblock
func (e_ Expression) ExpressionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("expressionBlock"))
	return rv
}


// SetExpressionBlock sets the value of the expressionBlock property.
// The block that executes to evaluate the expression.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/expressionblock
func (e_ Expression) SetExpressionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setExpressionBlock:"), value)
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



