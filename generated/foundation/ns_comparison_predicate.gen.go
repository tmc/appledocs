// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [ComparisonPredicate] class.
type IComparisonPredicate interface {
	IPredicate
	ComparisonPredicateModifier() NSComparisonPredicateModifier
	CustomSelector() objc.SEL
	LeftExpression() IExpression
	Options() NSComparisonPredicateOptions
	PredicateOperatorType() NSPredicateOperatorType
	RightExpression() IExpression
}

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

// Alloc allocates a new instance without initialization.
func (cc _ComparisonPredicateClass) Alloc() ComparisonPredicate {
	rv := objc.Send[ComparisonPredicate](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Creates a predicate by decoding from the coder you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/init(coder:)
func NewComparisonPredicateWithCoder(coder ICoder) ComparisonPredicate {
	instance := getComparisonPredicateClass().Alloc()
	rv := objc.Send[ComparisonPredicate](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// Creates a predicate that you form by combining specified left and right expressions using a specified selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/init(leftExpression:rightExpression:customSelector:)
func NewComparisonPredicateWithLeftExpressionRightExpressionCustomSelector(lhs IExpression, rhs IExpression, selector objc.SEL) ComparisonPredicate {
	instance := getComparisonPredicateClass().Alloc()
	rv := objc.Send[ComparisonPredicate](instance.ID, objc.Sel("initWithLeftExpression:rightExpression:customSelector:"), lhs, rhs, selector)
	rv.Autorelease()
	return rv
}


// Creates a predicate to a specified type that you form by combining specified left and right expressions using a specified modifier and options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/init(leftExpression:rightExpression:modifier:type:options:)
func NewComparisonPredicateWithLeftExpressionRightExpressionModifierTypeOptions(lhs IExpression, rhs IExpression, modifier NSComparisonPredicateModifier, type_ NSPredicateOperatorType, options NSComparisonPredicateOptions) ComparisonPredicate {
	instance := getComparisonPredicateClass().Alloc()
	rv := objc.Send[ComparisonPredicate](instance.ID, objc.Sel("initWithLeftExpression:rightExpression:modifier:type:options:"), lhs, rhs, modifier, type_, options)
	rv.Autorelease()
	return rv
}



// Returns a new predicate formed by combining the left and right expressions using a given selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/predicateWithLeftExpression:rightExpression:customSelector:
func (cc _ComparisonPredicateClass) PredicateWithLeftExpressionRightExpressionCustomSelector(lhs IExpression, rhs IExpression, selector objc.SEL) IComparisonPredicate {
	rv := objc.Send[ComparisonPredicate](objc.ID(cc.class), objc.Sel("predicateWithLeftExpression:rightExpression:customSelector:"), lhs, rhs, selector)
	return rv
}


// Creates and returns a predicate of a given type formed by combining given left and right expressions using a given modifier and options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/predicateWithLeftExpression:rightExpression:modifier:type:options:
func (cc _ComparisonPredicateClass) PredicateWithLeftExpressionRightExpressionModifierTypeOptions(lhs IExpression, rhs IExpression, modifier NSComparisonPredicateModifier, type_ NSPredicateOperatorType, options NSComparisonPredicateOptions) IComparisonPredicate {
	rv := objc.Send[ComparisonPredicate](objc.ID(cc.class), objc.Sel("predicateWithLeftExpression:rightExpression:modifier:type:options:"), lhs, rhs, modifier, type_, options)
	return rv
}


// The comparison predicate modifier for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/comparisonPredicateModifier
func (c_ ComparisonPredicate) ComparisonPredicateModifier() NSComparisonPredicateModifier {
	rv := objc.Send[ComparisonPredicateModifier](c_.ID, objc.Sel("comparisonPredicateModifier"))
	return rv
}


// The selector for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/customSelector
func (c_ ComparisonPredicate) CustomSelector() objc.SEL {
	rv := objc.Send[objc.SEL](c_.ID, objc.Sel("customSelector"))
	return rv
}


// The left expression for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/leftExpression
func (c_ ComparisonPredicate) LeftExpression() IExpression {
	rv := objc.Send[Expression](c_.ID, objc.Sel("leftExpression"))
	return rv
}


// The options to use for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/options-swift.property
func (c_ ComparisonPredicate) Options() NSComparisonPredicateOptions {
	rv := objc.Send[ComparisonPredicateOptions](c_.ID, objc.Sel("options"))
	return rv
}


// The predicate type for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/predicateOperatorType
func (c_ ComparisonPredicate) PredicateOperatorType() NSPredicateOperatorType {
	rv := objc.Send[PredicateOperatorType](c_.ID, objc.Sel("predicateOperatorType"))
	return rv
}


// The right expression for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/rightExpression
func (c_ ComparisonPredicate) RightExpression() IExpression {
	rv := objc.Send[Expression](c_.ID, objc.Sel("rightExpression"))
	return rv
}


