// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSExpression */


/* debug [class_header]: Header for NSExpression */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Expression */
// An interface definition for the [Expression] class.
type IExpression interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Expression */
	// properties:
	Arguments() []Expression
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Expression */
	// methods:
	AllowEvaluation()
	ExpressionValueWithObjectContext(object objc.IObject, context IMutableDictionary) objc.ID
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Expression */
// Alloc allocates a new instance without initialization.
func (ec _ExpressionClass) Alloc() Expression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Expression */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Expression */

// Creates an aggregate expression for a specified collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forAggregate:)
func NewExpressionForAggregate(subexpressions []Expression) Expression {
	rv := objc.Send[Expression](objc.ID(getExpressionClass().class), objc.Sel("expressionForAggregate:"), subexpressions)
	return rv
}/* debug [class_init_methods/constructor]: NewExpressionForAggregate */


// Creates an expression object that uses the block for evaluating objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(block:arguments:)
func NewExpressionForBlockArguments(block unsafe.Pointer, arguments []Expression) Expression {
	rv := objc.Send[Expression](objc.ID(getExpressionClass().class), objc.Sel("expressionForBlock:arguments:"), block, arguments)
	return rv
}/* debug [class_init_methods/constructor]: NewExpressionForBlockArguments */


// Creates an expression that returns a result, depending on the value of predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forConditional:trueExpression:falseExpression:)
func NewExpressionForConditionalTrueExpressionFalseExpression(predicate IPredicate, trueExpression IExpression, falseExpression IExpression) Expression {
	rv := objc.Send[Expression](objc.ID(getExpressionClass().class), objc.Sel("expressionForConditional:trueExpression:falseExpression:"), predicate, trueExpression, falseExpression)
	return rv
}/* debug [class_init_methods/constructor]: NewExpressionForConditionalTrueExpressionFalseExpression */


// Creates an expression that represents a specified constant value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forConstantValue:)
func NewExpressionForConstantValue(obj objc.IObject) Expression {
	rv := objc.Send[Expression](objc.ID(getExpressionClass().class), objc.Sel("expressionForConstantValue:"), obj)
	return rv
}/* debug [class_init_methods/constructor]: NewExpressionForConstantValue */


// Creates an expression that invokes one of the predefined functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forFunction:arguments:)
func NewExpressionForFunctionArguments(name IString, parameters IArray) Expression {
	rv := objc.Send[Expression](objc.ID(getExpressionClass().class), objc.Sel("expressionForFunction:arguments:"), name, parameters)
	return rv
}/* debug [class_init_methods/constructor]: NewExpressionForFunctionArguments */


// Creates an expression that returns the result of invoking a selector with a specified name using specified arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forFunction:selectorName:arguments:)
func NewExpressionForFunctionSelectorNameArguments(target IExpression, name IString, parameters IArray) Expression {
	rv := objc.Send[Expression](objc.ID(getExpressionClass().class), objc.Sel("expressionForFunction:selectorName:arguments:"), target, name, parameters)
	return rv
}/* debug [class_init_methods/constructor]: NewExpressionForFunctionSelectorNameArguments */


// Creates an expression object that represents the intersection of a specified set and collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forIntersectSet:with:)
func NewExpressionForIntersectSetWith(left IExpression, right IExpression) Expression {
	rv := objc.Send[Expression](objc.ID(getExpressionClass().class), objc.Sel("expressionForIntersectSet:with:"), left, right)
	return rv
}/* debug [class_init_methods/constructor]: NewExpressionForIntersectSetWith */


// Creates an expression that invokes the value function with a specified key path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forKeyPath:)-1aqf5
func NewExpressionForKeyPath(keyPath IString) Expression {
	rv := objc.Send[Expression](objc.ID(getExpressionClass().class), objc.Sel("expressionForKeyPath:"), keyPath)
	return rv
}/* debug [class_init_methods/constructor]: NewExpressionForKeyPath */


// Creates an expression object that represents the subtraction of a specified collection from a specified set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forMinusSet:with:)
func NewExpressionForMinusSetWith(left IExpression, right IExpression) Expression {
	rv := objc.Send[Expression](objc.ID(getExpressionClass().class), objc.Sel("expressionForMinusSet:with:"), left, right)
	return rv
}/* debug [class_init_methods/constructor]: NewExpressionForMinusSetWith */


// Creates an expression that filters a collection by storing elements in the collection in a specified variable and keeping the elements that the qualifier returns as true.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forSubquery:usingIteratorVariable:predicate:)
func NewExpressionForSubqueryUsingIteratorVariablePredicate(expression IExpression, variable IString, predicate IPredicate) Expression {
	rv := objc.Send[Expression](objc.ID(getExpressionClass().class), objc.Sel("expressionForSubquery:usingIteratorVariable:predicate:"), expression, variable, predicate)
	return rv
}/* debug [class_init_methods/constructor]: NewExpressionForSubqueryUsingIteratorVariablePredicate */


// Creates an expression object that represents the union of a specified set and collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forUnionSet:with:)
func NewExpressionForUnionSetWith(left IExpression, right IExpression) Expression {
	rv := objc.Send[Expression](objc.ID(getExpressionClass().class), objc.Sel("expressionForUnionSet:with:"), left, right)
	return rv
}/* debug [class_init_methods/constructor]: NewExpressionForUnionSetWith */


// Creates an expression that extracts a value from the variable bindings dictionary for a specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forVariable:)
func NewExpressionForVariable(string_ IString) Expression {
	rv := objc.Send[Expression](objc.ID(getExpressionClass().class), objc.Sel("expressionForVariable:"), string_)
	return rv
}/* debug [class_init_methods/constructor]: NewExpressionForVariable */


// Creates an expression by decoding from the coder you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(coder:)
func NewExpressionWithCoder(coder ICoder) Expression {
	instance := getExpressionClass().Alloc()
	rv := objc.Send[Expression](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewExpressionWithCoder */


// Creates the expression with the specified expression type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(expressionType:)
func NewExpressionWithExpressionType(type_ ExpressionType) Expression {
	instance := getExpressionClass().Alloc()
	rv := objc.Send[Expression](instance.ID, objc.Sel("initWithExpressionType:"), type_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewExpressionWithExpressionType */


// Creates the expression with the specified expression format and array of arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(format:argumentArray:)
func NewExpressionWithFormatArgumentArray(expressionFormat IString, arguments IArray) Expression {
	rv := objc.Send[Expression](objc.ID(getExpressionClass().class), objc.Sel("expressionWithFormat:argumentArray:"), expressionFormat, arguments)
	return rv
}/* debug [class_init_methods/constructor]: NewExpressionWithFormatArgumentArray */


// Creates the expression with the specified expression format and arguments list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(format:arguments:)
func NewExpressionWithFormatArguments(expressionFormat IString, argList objectivec.IObject) Expression {
	rv := objc.Send[Expression](objc.ID(getExpressionClass().class), objc.Sel("expressionWithFormat:arguments:"), expressionFormat, argList)
	return rv
}/* debug [class_init_methods/constructor]: NewExpressionWithFormatArguments */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Expression */

// Creates an expression that represents any key for a Spotlight query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/expressionForAnyKey()
func (ec _ExpressionClass) ExpressionForAnyKey() IExpression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("expressionForAnyKey"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ExpressionForAnyKey) */


// Creates an expression that represents the object you’re evaluating.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/expressionForEvaluatedObject()
func (ec _ExpressionClass) ExpressionForEvaluatedObject() IExpression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("expressionForEvaluatedObject"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ExpressionForEvaluatedObject) */


// Creates the expression with the specified expression arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/expressionWithFormat:
func (ec _ExpressionClass) ExpressionWithFormat(expressionFormat IString) IExpression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("expressionWithFormat:"), expressionFormat)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ExpressionWithFormat) */


// Creates an expression object that uses the block for evaluating objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(block:arguments:)
func (ec _ExpressionClass) ExpressionForBlockArguments(block unsafe.Pointer, arguments []Expression) IExpression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("expressionForBlock:arguments:"), block, arguments)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ExpressionForBlockArguments) */


// Creates an aggregate expression for a specified collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forAggregate:)
func (ec _ExpressionClass) ExpressionForAggregate(subexpressions []Expression) IExpression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("expressionForAggregate:"), subexpressions)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ExpressionForAggregate) */


// Creates an expression that returns a result, depending on the value of predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forConditional:trueExpression:falseExpression:)
func (ec _ExpressionClass) ExpressionForConditionalTrueExpressionFalseExpression(predicate IPredicate, trueExpression IExpression, falseExpression IExpression) IExpression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("expressionForConditional:trueExpression:falseExpression:"), predicate, trueExpression, falseExpression)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ExpressionForConditionalTrueExpressionFalseExpression) */


// Creates an expression that represents a specified constant value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forConstantValue:)
func (ec _ExpressionClass) ExpressionForConstantValue(obj objc.IObject) IExpression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("expressionForConstantValue:"), obj)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ExpressionForConstantValue) */


// Creates an expression that invokes one of the predefined functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forFunction:arguments:)
func (ec _ExpressionClass) ExpressionForFunctionArguments(name IString, parameters IArray) IExpression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("expressionForFunction:arguments:"), name, parameters)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ExpressionForFunctionArguments) */


// Creates an expression that returns the result of invoking a selector with a specified name using specified arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forFunction:selectorName:arguments:)
func (ec _ExpressionClass) ExpressionForFunctionSelectorNameArguments(target IExpression, name IString, parameters IArray) IExpression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("expressionForFunction:selectorName:arguments:"), target, name, parameters)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ExpressionForFunctionSelectorNameArguments) */


// Creates an expression object that represents the intersection of a specified set and collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forIntersectSet:with:)
func (ec _ExpressionClass) ExpressionForIntersectSetWith(left IExpression, right IExpression) IExpression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("expressionForIntersectSet:with:"), left, right)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ExpressionForIntersectSetWith) */


// Creates an expression that invokes the value function with a specified key path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forKeyPath:)-1aqf5
func (ec _ExpressionClass) ExpressionForKeyPath(keyPath IString) IExpression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("expressionForKeyPath:"), keyPath)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ExpressionForKeyPath) */


// Creates an expression object that represents the subtraction of a specified collection from a specified set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forMinusSet:with:)
func (ec _ExpressionClass) ExpressionForMinusSetWith(left IExpression, right IExpression) IExpression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("expressionForMinusSet:with:"), left, right)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ExpressionForMinusSetWith) */


// Creates an expression that filters a collection by storing elements in the collection in a specified variable and keeping the elements that the qualifier returns as true.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forSubquery:usingIteratorVariable:predicate:)
func (ec _ExpressionClass) ExpressionForSubqueryUsingIteratorVariablePredicate(expression IExpression, variable IString, predicate IPredicate) IExpression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("expressionForSubquery:usingIteratorVariable:predicate:"), expression, variable, predicate)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ExpressionForSubqueryUsingIteratorVariablePredicate) */


// Creates an expression object that represents the union of a specified set and collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forUnionSet:with:)
func (ec _ExpressionClass) ExpressionForUnionSetWith(left IExpression, right IExpression) IExpression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("expressionForUnionSet:with:"), left, right)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ExpressionForUnionSetWith) */


// Creates an expression that extracts a value from the variable bindings dictionary for a specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(forVariable:)
func (ec _ExpressionClass) ExpressionForVariable(string_ IString) IExpression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("expressionForVariable:"), string_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ExpressionForVariable) */


// Creates the expression with the specified expression format and array of arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(format:argumentArray:)
func (ec _ExpressionClass) ExpressionWithFormatArgumentArray(expressionFormat IString, arguments IArray) IExpression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("expressionWithFormat:argumentArray:"), expressionFormat, arguments)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ExpressionWithFormatArgumentArray) */


// Creates the expression with the specified expression format and arguments list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/init(format:arguments:)
func (ec _ExpressionClass) ExpressionWithFormatArguments(expressionFormat IString, argList objectivec.IObject) IExpression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("expressionWithFormat:arguments:"), expressionFormat, argList)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ExpressionWithFormatArguments) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Expression */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Expression */

// Forces a securely decoded expression to allow evaluation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/allowEvaluation()
func (e_ Expression) AllowEvaluation() {
	objc.Send[objc.ID](e_.ID, objc.Sel("allowEvaluation"))
}/* debug [instance_methods/method]: AllowEvaluation */


// Evaluates an expression using a specified object and context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/expressionValue(with:context:)
func (e_ Expression) ExpressionValueWithObjectContext(object objc.IObject, context IMutableDictionary) objc.ID {
	rv := objc.Send[objc.ID](e_.ID, objc.Sel("expressionValueWithObject:context:"), object, context)
	return rv
}/* debug [instance_methods/method]: ExpressionValueWithObjectContext */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Expression */

// The arguments for the expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/arguments
func (e_ Expression) Arguments() []Expression {
	rv := objc.Send[[]Expression](e_.ID, objc.Sel("arguments"))
	return rv
}/* debug [instance_properties/getter]: arguments */


// The collection of expressions in an aggregate expression, or the collection element of a subquery expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/collection
func (e_ Expression) Collection() objc.ID {
	rv := objc.Send[objc.ID](e_.ID, objc.Sel("collection"))
	return rv
}/* debug [instance_properties/getter]: collection */


// The constant value of the expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/constantValue
func (e_ Expression) ConstantValue() objc.ID {
	rv := objc.Send[objc.ID](e_.ID, objc.Sel("constantValue"))
	return rv
}/* debug [instance_properties/getter]: constantValue */


// The block that executes to evaluate the expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/expressionBlock
func (e_ Expression) ExpressionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("expressionBlock"))
	return rv
}/* debug [instance_properties/getter]: expressionBlock */


// The expression type for the expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/expressionType-swift.property
func (e_ Expression) ExpressionType() ExpressionType {
	rv := objc.Send[ExpressionType](e_.ID, objc.Sel("expressionType"))
	return rv
}/* debug [instance_properties/getter]: expressionType */


// An expression to evalutate if a conditional expression’s predicate evaluates to false.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/false
func (e_ Expression) FalseExpression() IExpression {
	rv := objc.Send[Expression](e_.ID, objc.Sel("falseExpression"))
	return rv
}/* debug [instance_properties/getter]: falseExpression */


// The function for the expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/function
func (e_ Expression) Function() IString {
	rv := objc.Send[String](e_.ID, objc.Sel("function"))
	return rv
}/* debug [instance_properties/getter]: function */


// The key path for the expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/keyPath
func (e_ Expression) KeyPath() IString {
	rv := objc.Send[String](e_.ID, objc.Sel("keyPath"))
	return rv
}/* debug [instance_properties/getter]: keyPath */


// The left expression of an aggregate expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/left
func (e_ Expression) LeftExpression() IExpression {
	rv := objc.Send[Expression](e_.ID, objc.Sel("leftExpression"))
	return rv
}/* debug [instance_properties/getter]: leftExpression */


// The operand for the expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/operand
func (e_ Expression) Operand() IExpression {
	rv := objc.Send[Expression](e_.ID, objc.Sel("operand"))
	return rv
}/* debug [instance_properties/getter]: operand */


// The predicate of a subquery expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/predicate
func (e_ Expression) Predicate() IPredicate {
	rv := objc.Send[Predicate](e_.ID, objc.Sel("predicate"))
	return rv
}/* debug [instance_properties/getter]: predicate */


// The right expression of an aggregate expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/right
func (e_ Expression) RightExpression() IExpression {
	rv := objc.Send[Expression](e_.ID, objc.Sel("rightExpression"))
	return rv
}/* debug [instance_properties/getter]: rightExpression */


// An expression to evalutate if a conditional expression’s predicate evaluates to true.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/true
func (e_ Expression) TrueExpression() IExpression {
	rv := objc.Send[Expression](e_.ID, objc.Sel("trueExpression"))
	return rv
}/* debug [instance_properties/getter]: trueExpression */


// The variable for the expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExpression/variable
func (e_ Expression) Variable() IString {
	rv := objc.Send[String](e_.ID, objc.Sel("variable"))
	return rv
}/* debug [instance_properties/getter]: variable */


// An expression to evalutate if a conditional expression’s predicate evaluates to false.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/false
func (e_ Expression) False() IExpression {
	rv := objc.Send[Expression](e_.ID, objc.Sel("false"))
	return rv
}/* debug [instance_properties/getter]: false */


// An expression to evalutate if a conditional expression’s predicate evaluates to false.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/false
func (e_ Expression) SetFalse(value IExpression) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setFalse:"), value)
}/* debug [instance_properties/setter]: false */


// The left expression of an aggregate expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/left
func (e_ Expression) Left() IExpression {
	rv := objc.Send[Expression](e_.ID, objc.Sel("left"))
	return rv
}/* debug [instance_properties/getter]: left */


// The left expression of an aggregate expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/left
func (e_ Expression) SetLeft(value IExpression) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setLeft:"), value)
}/* debug [instance_properties/setter]: left */


// The right expression of an aggregate expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/right
func (e_ Expression) Right() IExpression {
	rv := objc.Send[Expression](e_.ID, objc.Sel("right"))
	return rv
}/* debug [instance_properties/getter]: right */


// The right expression of an aggregate expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/right
func (e_ Expression) SetRight(value IExpression) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setRight:"), value)
}/* debug [instance_properties/setter]: right */


// An expression to evalutate if a conditional expression’s predicate evaluates to true.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/true
func (e_ Expression) True() IExpression {
	rv := objc.Send[Expression](e_.ID, objc.Sel("true"))
	return rv
}/* debug [instance_properties/getter]: true */


// An expression to evalutate if a conditional expression’s predicate evaluates to true.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/true
func (e_ Expression) SetTrue(value IExpression) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setTrue:"), value)
}/* debug [instance_properties/setter]: true */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSExpression */


