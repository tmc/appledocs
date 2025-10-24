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
	// properties:
	Arguments() []Expression /* primitive/slice/pointer. */
	Collection() objc.ID
	ConstantValue() objc.ID
	ExpressionBlock() unsafe.Pointer
	ExpressionType() ExpressionType
	FalseExpression() IExpression
	Function() IString
	KeyPath() IString
	LeftExpression() IExpression
	Operand() IExpression
	Predicate() IPredicate
	RightExpression() IExpression
	TrueExpression() IExpression
	Variable() IString
	False() IExpression
	SetFalse(value IExpression)
	Left() IExpression
	SetLeft(value IExpression)
	Right() IExpression
	SetRight(value IExpression)
	True() IExpression
	SetTrue(value IExpression)
	// methods:
	AllowEvaluation()
	ExpressionValueWithObjectContext(object objectivec.IObject, context IMutableDictionary) objc.ID
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



// Creates an aggregate expression for a specified collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forAggregate:)
func NewExpressionForAggregate(subexpressions []Expression /* primitive/slice/pointer. */) Expression {
	rv := objc.Send[Expression](objc.ID(getExpressionClass().class), objc.Sel("expressionForAggregate:"), subexpressions)
	return rv
}


// Creates an expression object that uses the block for evaluating objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(block:arguments:)
func NewExpressionForBlockArguments(block unsafe.Pointer, arguments []Expression /* primitive/slice/pointer. */) Expression {
	rv := objc.Send[Expression](objc.ID(getExpressionClass().class), objc.Sel("expressionForBlock:arguments:"), block, arguments)
	return rv
}


// Creates an expression that returns a result, depending on the value of predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forConditional:trueExpression:falseExpression:)
func NewExpressionForConditionalTrueExpressionFalseExpression(predicate IPredicate, trueExpression IExpression, falseExpression IExpression) Expression {
	rv := objc.Send[Expression](objc.ID(getExpressionClass().class), objc.Sel("expressionForConditional:trueExpression:falseExpression:"), predicate, trueExpression, falseExpression)
	return rv
}


// Creates an expression that represents a specified constant value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forConstantValue:)
func NewExpressionForConstantValue(obj objectivec.IObject) Expression {
	rv := objc.Send[Expression](objc.ID(getExpressionClass().class), objc.Sel("expressionForConstantValue:"), obj)
	return rv
}


// Creates an expression that invokes one of the predefined functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forFunction:arguments:)
func NewExpressionForFunctionArguments(name IString, parameters IArray) Expression {
	rv := objc.Send[Expression](objc.ID(getExpressionClass().class), objc.Sel("expressionForFunction:arguments:"), name, parameters)
	return rv
}


// Creates an expression that returns the result of invoking a selector with a specified name using specified arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forFunction:selectorName:arguments:)
func NewExpressionForFunctionSelectorNameArguments(target IExpression, name IString, parameters IArray) Expression {
	rv := objc.Send[Expression](objc.ID(getExpressionClass().class), objc.Sel("expressionForFunction:selectorName:arguments:"), target, name, parameters)
	return rv
}


// Creates an expression object that represents the intersection of a specified set and collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forIntersectSet:with:)
func NewExpressionForIntersectSetWith(left IExpression, right IExpression) Expression {
	rv := objc.Send[Expression](objc.ID(getExpressionClass().class), objc.Sel("expressionForIntersectSet:with:"), left, right)
	return rv
}


// Creates an expression that invokes the value function with a specified key path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forKeyPath:)-1aqf5
func NewExpressionForKeyPath(keyPath IString) Expression {
	rv := objc.Send[Expression](objc.ID(getExpressionClass().class), objc.Sel("expressionForKeyPath:"), keyPath)
	return rv
}


// Creates an expression object that represents the subtraction of a specified collection from a specified set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forMinusSet:with:)
func NewExpressionForMinusSetWith(left IExpression, right IExpression) Expression {
	rv := objc.Send[Expression](objc.ID(getExpressionClass().class), objc.Sel("expressionForMinusSet:with:"), left, right)
	return rv
}


// Creates an expression that filters a collection by storing elements in the collection in a specified variable and keeping the elements that the qualifier returns as true.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forSubquery:usingIteratorVariable:predicate:)
func NewExpressionForSubqueryUsingIteratorVariablePredicate(expression IExpression, variable IString, predicate IPredicate) Expression {
	rv := objc.Send[Expression](objc.ID(getExpressionClass().class), objc.Sel("expressionForSubquery:usingIteratorVariable:predicate:"), expression, variable, predicate)
	return rv
}


// Creates an expression object that represents the union of a specified set and collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forUnionSet:with:)
func NewExpressionForUnionSetWith(left IExpression, right IExpression) Expression {
	rv := objc.Send[Expression](objc.ID(getExpressionClass().class), objc.Sel("expressionForUnionSet:with:"), left, right)
	return rv
}


// Creates an expression that extracts a value from the variable bindings dictionary for a specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forVariable:)
func NewExpressionForVariable(string_ IString) Expression {
	rv := objc.Send[Expression](objc.ID(getExpressionClass().class), objc.Sel("expressionForVariable:"), string_)
	return rv
}


// Creates an expression by decoding from the coder you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(coder:)
func NewExpressionWithCoder(coder ICoder) Expression {
	instance := getExpressionClass().Alloc()
	rv := objc.Send[Expression](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// Creates the expression with the specified expression type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(expressionType:)
func NewExpressionWithExpressionType(type_ ExpressionType) Expression {
	instance := getExpressionClass().Alloc()
	rv := objc.Send[Expression](instance.ID, objc.Sel("initWithExpressionType:"), type_)
	rv.Autorelease()
	return rv
}


// Creates the expression with the specified expression format and array of arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(format:argumentArray:)
func NewExpressionWithFormatArgumentArray(expressionFormat IString, arguments IArray) Expression {
	rv := objc.Send[Expression](objc.ID(getExpressionClass().class), objc.Sel("expressionWithFormat:argumentArray:"), expressionFormat, arguments)
	return rv
}


// Creates the expression with the specified expression format and arguments list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(format:arguments:)
func NewExpressionWithFormatArguments(expressionFormat IString, argList unsafe.Pointer) Expression {
	rv := objc.Send[Expression](objc.ID(getExpressionClass().class), objc.Sel("expressionWithFormat:arguments:"), expressionFormat, argList)
	return rv
}



// Creates an expression that represents any key for a Spotlight query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/expressionForAnyKey()
func (ec _ExpressionClass) ExpressionForAnyKey() IExpression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("expressionForAnyKey"))
	return rv
}


// Creates an expression that represents the object you’re evaluating.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/expressionForEvaluatedObject()
func (ec _ExpressionClass) ExpressionForEvaluatedObject() IExpression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("expressionForEvaluatedObject"))
	return rv
}


// Creates the expression with the specified expression arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/expressionWithFormat:
func (ec _ExpressionClass) ExpressionWithFormat(expressionFormat IString) IExpression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("expressionWithFormat:"), expressionFormat)
	return rv
}


// Creates an expression object that uses the block for evaluating objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(block:arguments:)
func (ec _ExpressionClass) ExpressionForBlockArguments(block unsafe.Pointer, arguments []Expression /* primitive/slice/pointer. */) IExpression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("expressionForBlock:arguments:"), block, arguments)
	return rv
}


// Creates an aggregate expression for a specified collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forAggregate:)
func (ec _ExpressionClass) ExpressionForAggregate(subexpressions []Expression /* primitive/slice/pointer. */) IExpression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("expressionForAggregate:"), subexpressions)
	return rv
}


// Creates an expression that returns a result, depending on the value of predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forConditional:trueExpression:falseExpression:)
func (ec _ExpressionClass) ExpressionForConditionalTrueExpressionFalseExpression(predicate IPredicate, trueExpression IExpression, falseExpression IExpression) IExpression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("expressionForConditional:trueExpression:falseExpression:"), predicate, trueExpression, falseExpression)
	return rv
}


// Creates an expression that represents a specified constant value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forConstantValue:)
func (ec _ExpressionClass) ExpressionForConstantValue(obj objectivec.IObject) IExpression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("expressionForConstantValue:"), obj)
	return rv
}


// Creates an expression that invokes one of the predefined functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forFunction:arguments:)
func (ec _ExpressionClass) ExpressionForFunctionArguments(name IString, parameters IArray) IExpression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("expressionForFunction:arguments:"), name, parameters)
	return rv
}


// Creates an expression that returns the result of invoking a selector with a specified name using specified arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forFunction:selectorName:arguments:)
func (ec _ExpressionClass) ExpressionForFunctionSelectorNameArguments(target IExpression, name IString, parameters IArray) IExpression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("expressionForFunction:selectorName:arguments:"), target, name, parameters)
	return rv
}


// Creates an expression object that represents the intersection of a specified set and collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forIntersectSet:with:)
func (ec _ExpressionClass) ExpressionForIntersectSetWith(left IExpression, right IExpression) IExpression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("expressionForIntersectSet:with:"), left, right)
	return rv
}


// Creates an expression that invokes the value function with a specified key path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forKeyPath:)-1aqf5
func (ec _ExpressionClass) ExpressionForKeyPath(keyPath IString) IExpression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("expressionForKeyPath:"), keyPath)
	return rv
}


// Creates an expression object that represents the subtraction of a specified collection from a specified set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forMinusSet:with:)
func (ec _ExpressionClass) ExpressionForMinusSetWith(left IExpression, right IExpression) IExpression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("expressionForMinusSet:with:"), left, right)
	return rv
}


// Creates an expression that filters a collection by storing elements in the collection in a specified variable and keeping the elements that the qualifier returns as true.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forSubquery:usingIteratorVariable:predicate:)
func (ec _ExpressionClass) ExpressionForSubqueryUsingIteratorVariablePredicate(expression IExpression, variable IString, predicate IPredicate) IExpression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("expressionForSubquery:usingIteratorVariable:predicate:"), expression, variable, predicate)
	return rv
}


// Creates an expression object that represents the union of a specified set and collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forUnionSet:with:)
func (ec _ExpressionClass) ExpressionForUnionSetWith(left IExpression, right IExpression) IExpression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("expressionForUnionSet:with:"), left, right)
	return rv
}


// Creates an expression that extracts a value from the variable bindings dictionary for a specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forVariable:)
func (ec _ExpressionClass) ExpressionForVariable(string_ IString) IExpression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("expressionForVariable:"), string_)
	return rv
}


// Creates the expression with the specified expression format and array of arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(format:argumentArray:)
func (ec _ExpressionClass) ExpressionWithFormatArgumentArray(expressionFormat IString, arguments IArray) IExpression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("expressionWithFormat:argumentArray:"), expressionFormat, arguments)
	return rv
}


// Creates the expression with the specified expression format and arguments list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(format:arguments:)
func (ec _ExpressionClass) ExpressionWithFormatArguments(expressionFormat IString, argList unsafe.Pointer) IExpression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("expressionWithFormat:arguments:"), expressionFormat, argList)
	return rv
}


// Forces a securely decoded expression to allow evaluation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/allowEvaluation()
func (e_ Expression) AllowEvaluation() {
	objc.Send[objc.ID](e_.ID, objc.Sel("allowEvaluation"))
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
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/arguments
func (e_ Expression) Arguments() []Expression /* primitive/slice/pointer. */ {
	rv := objc.Send[[]Expression](e_.ID, objc.Sel("arguments"))
	return rv
}


// The collection of expressions in an aggregate expression, or the collection element of a subquery expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/collection
func (e_ Expression) Collection() objc.ID {
	rv := objc.Send[objc.ID](e_.ID, objc.Sel("collection"))
	return rv
}


// The constant value of the expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/constantValue
func (e_ Expression) ConstantValue() objc.ID {
	rv := objc.Send[objc.ID](e_.ID, objc.Sel("constantValue"))
	return rv
}


// The block that executes to evaluate the expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/expressionBlock
func (e_ Expression) ExpressionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("expressionBlock"))
	return rv
}


// The expression type for the expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/expressionType-swift.property
func (e_ Expression) ExpressionType() ExpressionType {
	rv := objc.Send[ExpressionType](e_.ID, objc.Sel("expressionType"))
	return rv
}


// An expression to evalutate if a conditional expression’s predicate evaluates to false.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/false
func (e_ Expression) FalseExpression() IExpression {
	rv := objc.Send[Expression](e_.ID, objc.Sel("falseExpression"))
	return rv
}


// The function for the expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/function
func (e_ Expression) Function() IString {
	rv := objc.Send[String](e_.ID, objc.Sel("function"))
	return rv
}


// The key path for the expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/keyPath
func (e_ Expression) KeyPath() IString {
	rv := objc.Send[String](e_.ID, objc.Sel("keyPath"))
	return rv
}


// The left expression of an aggregate expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/left
func (e_ Expression) LeftExpression() IExpression {
	rv := objc.Send[Expression](e_.ID, objc.Sel("leftExpression"))
	return rv
}


// The operand for the expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/operand
func (e_ Expression) Operand() IExpression {
	rv := objc.Send[Expression](e_.ID, objc.Sel("operand"))
	return rv
}


// The predicate of a subquery expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/predicate
func (e_ Expression) Predicate() IPredicate {
	rv := objc.Send[Predicate](e_.ID, objc.Sel("predicate"))
	return rv
}


// The right expression of an aggregate expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/right
func (e_ Expression) RightExpression() IExpression {
	rv := objc.Send[Expression](e_.ID, objc.Sel("rightExpression"))
	return rv
}


// An expression to evalutate if a conditional expression’s predicate evaluates to true.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/true
func (e_ Expression) TrueExpression() IExpression {
	rv := objc.Send[Expression](e_.ID, objc.Sel("trueExpression"))
	return rv
}


// The variable for the expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/variable
func (e_ Expression) Variable() IString {
	rv := objc.Send[String](e_.ID, objc.Sel("variable"))
	return rv
}


// An expression to evalutate if a conditional expression’s predicate evaluates to false.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/false
func (e_ Expression) False() IExpression {
	rv := objc.Send[Expression](e_.ID, objc.Sel("false"))
	return rv
}


// An expression to evalutate if a conditional expression’s predicate evaluates to false.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/false
func (e_ Expression) SetFalse(value IExpression) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setFalse:"), value)
}


// The left expression of an aggregate expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/left
func (e_ Expression) Left() IExpression {
	rv := objc.Send[Expression](e_.ID, objc.Sel("left"))
	return rv
}


// The left expression of an aggregate expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/left
func (e_ Expression) SetLeft(value IExpression) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setLeft:"), value)
}


// The right expression of an aggregate expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/right
func (e_ Expression) Right() IExpression {
	rv := objc.Send[Expression](e_.ID, objc.Sel("right"))
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
	rv := objc.Send[Expression](e_.ID, objc.Sel("true"))
	return rv
}


// An expression to evalutate if a conditional expression’s predicate evaluates to true.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/true
func (e_ Expression) SetTrue(value IExpression) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setTrue:"), value)
}


