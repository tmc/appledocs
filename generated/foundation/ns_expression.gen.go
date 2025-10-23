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
	ExpressionValueWithObjectContext(object objectivec.IObject, context IMutableDictionary) objc.ID
	Arguments() IExpression
	SetArguments(value IExpression)
	Collection() unsafe.Pointer
	SetCollection(value unsafe.Pointer)
	ConstantValue() unsafe.Pointer
	SetConstantValue(value unsafe.Pointer)
	ExpressionBlock() IMutableDictionary
	SetExpressionBlock(value IMutableDictionary)
	ExpressionType() unsafe.Pointer
	SetExpressionType(value unsafe.Pointer)
	False() IExpression
	SetFalse(value IExpression)
	Function() string
	SetFunction(value string)
	KeyPath() string
	SetKeyPath(value string)
	Left() IExpression
	SetLeft(value IExpression)
	Operand() IExpression
	SetOperand(value IExpression)
	Predicate() IPredicate
	SetPredicate(value IPredicate)
	Right() IExpression
	SetRight(value IExpression)
	True() IExpression
	SetTrue(value IExpression)
	Variable() string
	SetVariable(value string)
}

// An expression for use in a comparison predicate.
//
// Comparison operations in an derive from two expressions as instances of the class. You create expressions for constant values, key paths, and so on. Generally, anywhere in the class hierarchy where there’s a composite API and subtypes that may only reasonably respond to a subset of that API, invoking a method that doesn’t make sense for that subtype throws an exception.


// An expression for use in a comparison predicate.
//
// [Full Topic]
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



// Evaluates an expression using a specified object and context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/expressionValue(with:context:)
func (e_ Expression) ExpressionValueWithObjectContext(object objectivec.IObject, context IMutableDictionary) objc.ID {
	rv := objc.Send[objc.ID](e_.ID, objc.Sel("expressionValueWithObject:context:"), object, context)
	return rv
}


// The arguments for the expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/arguments
func (e_ Expression) Arguments() IExpression {
	rv := objc.Send[NSExpression](e_.ID, objc.Sel("arguments"))
	return rv
}


// The arguments for the expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/arguments
func (e_ Expression) SetArguments(value IExpression) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setArguments:"), value)
}


// The collection of expressions in an aggregate expression, or the collection element of a subquery expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/collection
func (e_ Expression) Collection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("collection"))
	return rv
}


// The collection of expressions in an aggregate expression, or the collection element of a subquery expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/collection
func (e_ Expression) SetCollection(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setCollection:"), value)
}


// The constant value of the expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/constantvalue
func (e_ Expression) ConstantValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("constantValue"))
	return rv
}


// The constant value of the expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/constantvalue
func (e_ Expression) SetConstantValue(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setConstantValue:"), value)
}


// The block that executes to evaluate the expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/expressionblock
func (e_ Expression) ExpressionBlock() IMutableDictionary {
	rv := objc.Send[NSMutableDictionary](e_.ID, objc.Sel("expressionBlock"))
	return rv
}


// The block that executes to evaluate the expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/expressionblock
func (e_ Expression) SetExpressionBlock(value IMutableDictionary) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setExpressionBlock:"), value)
}


// The expression type for the expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/expressiontype-swift.property
func (e_ Expression) ExpressionType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("expressionType"))
	return rv
}


// The expression type for the expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/expressiontype-swift.property
func (e_ Expression) SetExpressionType(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setExpressionType:"), value)
}


// An expression to evalutate if a conditional expression’s predicate evaluates to false.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/false
func (e_ Expression) False() IExpression {
	rv := objc.Send[NSExpression](e_.ID, objc.Sel("false"))
	return rv
}


// An expression to evalutate if a conditional expression’s predicate evaluates to false.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/false
func (e_ Expression) SetFalse(value IExpression) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setFalse:"), value)
}


// The function for the expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/function
func (e_ Expression) Function() string {
	rv := objc.Send[string](e_.ID, objc.Sel("function"))
	return rv
}


// The function for the expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/function
func (e_ Expression) SetFunction(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setFunction:"), objc.String(value))
}


// The key path for the expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/keypath
func (e_ Expression) KeyPath() string {
	rv := objc.Send[string](e_.ID, objc.Sel("keyPath"))
	return rv
}


// The key path for the expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/keypath
func (e_ Expression) SetKeyPath(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setKeyPath:"), objc.String(value))
}


// The left expression of an aggregate expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/left
func (e_ Expression) Left() IExpression {
	rv := objc.Send[NSExpression](e_.ID, objc.Sel("left"))
	return rv
}


// The left expression of an aggregate expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/left
func (e_ Expression) SetLeft(value IExpression) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setLeft:"), value)
}


// The operand for the expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/operand
func (e_ Expression) Operand() IExpression {
	rv := objc.Send[NSExpression](e_.ID, objc.Sel("operand"))
	return rv
}


// The operand for the expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/operand
func (e_ Expression) SetOperand(value IExpression) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setOperand:"), value)
}


// The predicate of a subquery expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/predicate
func (e_ Expression) Predicate() IPredicate {
	rv := objc.Send[NSPredicate](e_.ID, objc.Sel("predicate"))
	return rv
}


// The predicate of a subquery expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/predicate
func (e_ Expression) SetPredicate(value IPredicate) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setPredicate:"), value)
}


// The right expression of an aggregate expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/right
func (e_ Expression) Right() IExpression {
	rv := objc.Send[NSExpression](e_.ID, objc.Sel("right"))
	return rv
}


// The right expression of an aggregate expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/right
func (e_ Expression) SetRight(value IExpression) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setRight:"), value)
}


// An expression to evalutate if a conditional expression’s predicate evaluates to true.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/true
func (e_ Expression) True() IExpression {
	rv := objc.Send[NSExpression](e_.ID, objc.Sel("true"))
	return rv
}


// An expression to evalutate if a conditional expression’s predicate evaluates to true.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/true
func (e_ Expression) SetTrue(value IExpression) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setTrue:"), value)
}


// The variable for the expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/variable
func (e_ Expression) Variable() string {
	rv := objc.Send[string](e_.ID, objc.Sel("variable"))
	return rv
}


// The variable for the expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/variable
func (e_ Expression) SetVariable(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setVariable:"), objc.String(value))
}



