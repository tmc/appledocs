// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSComparisonPredicate */


/* debug [class_header]: Header for NSComparisonPredicate */
// The class instance for the [ComparisonPredicate] class.
var (
	ComparisonPredicateClass     _ComparisonPredicateClass
	ComparisonPredicateClassOnce sync.Once
)

func getComparisonPredicateClass() _ComparisonPredicateClass {
	ComparisonPredicateClassOnce.Do(func() {
		ComparisonPredicateClass = _ComparisonPredicateClass{objc.GetClass("NSComparisonPredicate")}
	})
	return ComparisonPredicateClass
}

type _ComparisonPredicateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ComparisonPredicate */
// An interface definition for the [ComparisonPredicate] class.
type IComparisonPredicate interface {
	IPredicate
	
/* debug [class_interface_properties]: Properties for ComparisonPredicate */
	// properties:
	ComparisonPredicateModifier() ComparisonPredicateModifier
	CustomSelector() objc.SEL
	LeftExpression() IExpression
	Options() ComparisonPredicateOptions
	PredicateOperatorType() PredicateOperatorType
	RightExpression() IExpression
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ComparisonPredicate */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ComparisonPredicate */
// Alloc allocates a new instance without initialization.
func (cc _ComparisonPredicateClass) Alloc() ComparisonPredicate {
	rv := objc.Send[ComparisonPredicate](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ComparisonPredicateClass) New() ComparisonPredicate {
	rv := objc.Send[ComparisonPredicate](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ComparisonPredicate) Init() ComparisonPredicate {
	rv := objc.Send[ComparisonPredicate](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ComparisonPredicate) Autorelease() ComparisonPredicate {
	rv := objc.Send[ComparisonPredicate](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewComparisonPredicate creates a new ComparisonPredicate instance.
func NewComparisonPredicate() ComparisonPredicate {
	return getComparisonPredicateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ComparisonPredicate */
// A specialized predicate for comparing expressions.
//
// Use comparison predicates to compare the results of two expressions. You create a comparison predicate with an operator, a left expression, and a right expression, and use instances of the class to represent those expressions. When you evaluate the predicate, it returns a value as the result of invoking the operator with the results of evaluating the expressions.


// A specialized predicate for comparing expressions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate
type ComparisonPredicate struct {
	Predicate
}

// ComparisonPredicateFrom constructs a [ComparisonPredicate] from an unsafe.Pointer.
//
// A specialized predicate for comparing expressions.
func ComparisonPredicateFrom(ptr unsafe.Pointer) ComparisonPredicate {
	return ComparisonPredicate{
		Predicate: PredicateFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ComparisonPredicate */

// Creates a predicate by decoding from the coder you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/init(coder:)
func NewComparisonPredicateWithCoder(coder ICoder) ComparisonPredicate {
	instance := getComparisonPredicateClass().Alloc()
	rv := objc.Send[ComparisonPredicate](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewComparisonPredicateWithCoder */


// Creates a predicate that you form by combining specified left and right expressions using a specified selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/init(leftExpression:rightExpression:customSelector:)
func NewComparisonPredicateWithLeftExpressionRightExpressionCustomSelector(lhs IExpression, rhs IExpression, selector objc.SEL) ComparisonPredicate {
	instance := getComparisonPredicateClass().Alloc()
	rv := objc.Send[ComparisonPredicate](instance.ID, objc.Sel("initWithLeftExpression:rightExpression:customSelector:"), lhs, rhs, selector)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewComparisonPredicateWithLeftExpressionRightExpressionCustomSelector */


// Creates a predicate to a specified type that you form by combining specified left and right expressions using a specified modifier and options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/init(leftExpression:rightExpression:modifier:type:options:)
func NewComparisonPredicateWithLeftExpressionRightExpressionModifierTypeOptions(lhs IExpression, rhs IExpression, modifier ComparisonPredicateModifier, type_ PredicateOperatorType, options ComparisonPredicateOptions) ComparisonPredicate {
	instance := getComparisonPredicateClass().Alloc()
	rv := objc.Send[ComparisonPredicate](instance.ID, objc.Sel("initWithLeftExpression:rightExpression:modifier:type:options:"), lhs, rhs, modifier, type_, options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewComparisonPredicateWithLeftExpressionRightExpressionModifierTypeOptions */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ComparisonPredicate */

// Returns a new predicate formed by combining the left and right expressions using a given selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/predicateWithLeftExpression:rightExpression:customSelector:
func (cc _ComparisonPredicateClass) PredicateWithLeftExpressionRightExpressionCustomSelector(lhs IExpression, rhs IExpression, selector objc.SEL) IComparisonPredicate {
	rv := objc.Send[ComparisonPredicate](objc.ID(cc.class), objc.Sel("predicateWithLeftExpression:rightExpression:customSelector:"), lhs, rhs, selector)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateWithLeftExpressionRightExpressionCustomSelector) */


// Creates and returns a predicate of a given type formed by combining given left and right expressions using a given modifier and options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/predicateWithLeftExpression:rightExpression:modifier:type:options:
func (cc _ComparisonPredicateClass) PredicateWithLeftExpressionRightExpressionModifierTypeOptions(lhs IExpression, rhs IExpression, modifier ComparisonPredicateModifier, type_ PredicateOperatorType, options ComparisonPredicateOptions) IComparisonPredicate {
	rv := objc.Send[ComparisonPredicate](objc.ID(cc.class), objc.Sel("predicateWithLeftExpression:rightExpression:modifier:type:options:"), lhs, rhs, modifier, type_, options)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateWithLeftExpressionRightExpressionModifierTypeOptions) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ComparisonPredicate */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ComparisonPredicate */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ComparisonPredicate */

// The comparison predicate modifier for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/comparisonPredicateModifier
func (c_ ComparisonPredicate) ComparisonPredicateModifier() ComparisonPredicateModifier {
	rv := objc.Send[ComparisonPredicateModifier](c_.ID, objc.Sel("comparisonPredicateModifier"))
	return rv
}/* debug [instance_properties/getter]: comparisonPredicateModifier */


// The selector for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/customSelector
func (c_ ComparisonPredicate) CustomSelector() objc.SEL {
	rv := objc.Send[objc.SEL](c_.ID, objc.Sel("customSelector"))
	return rv
}/* debug [instance_properties/getter]: customSelector */


// The left expression for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/leftExpression
func (c_ ComparisonPredicate) LeftExpression() IExpression {
	rv := objc.Send[Expression](c_.ID, objc.Sel("leftExpression"))
	return rv
}/* debug [instance_properties/getter]: leftExpression */


// The options to use for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/options-swift.property
func (c_ ComparisonPredicate) Options() ComparisonPredicateOptions {
	rv := objc.Send[ComparisonPredicateOptions](c_.ID, objc.Sel("options"))
	return rv
}/* debug [instance_properties/getter]: options */


// The predicate type for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/predicateOperatorType
func (c_ ComparisonPredicate) PredicateOperatorType() PredicateOperatorType {
	rv := objc.Send[PredicateOperatorType](c_.ID, objc.Sel("predicateOperatorType"))
	return rv
}/* debug [instance_properties/getter]: predicateOperatorType */


// The right expression for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/rightExpression
func (c_ ComparisonPredicate) RightExpression() IExpression {
	rv := objc.Send[Expression](c_.ID, objc.Sel("rightExpression"))
	return rv
}/* debug [instance_properties/getter]: rightExpression */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSComparisonPredicate */


