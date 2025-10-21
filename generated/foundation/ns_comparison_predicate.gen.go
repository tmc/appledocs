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
}

// A specialized predicate for comparing expressions.
//
// Use comparison predicates to compare the results of two expressions. You create a comparison predicate with an operator, a left expression, and a right expression, and use instances of the class to represent those expressions. When you evaluate the predicate, it returns a value as the result of invoking the operator with the results of evaluating the expressions.
//
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


// The right expression for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/rightExpression
func (c_ ComparisonPredicate) RightExpression() NSExpression {
	rv := objc.Send[NSExpression](c_.ID, objc.Sel("rightExpression"))
	return rv
}

// The comparison predicate modifier for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscomparisonpredicate/comparisonpredicatemodifier
func (c_ ComparisonPredicate) ComparisonPredicateModifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("comparisonPredicateModifier"))
	return rv
}


// SetComparisonPredicateModifier sets the value of the comparisonPredicateModifier property.
// The comparison predicate modifier for the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscomparisonpredicate/comparisonpredicatemodifier
func (c_ ComparisonPredicate) SetComparisonPredicateModifier(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setComparisonPredicateModifier:"), value)
}

// The selector for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscomparisonpredicate/customselector
func (c_ ComparisonPredicate) CustomSelector() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("customSelector"))
	return rv
}


// SetCustomSelector sets the value of the customSelector property.
// The selector for the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscomparisonpredicate/customselector
func (c_ ComparisonPredicate) SetCustomSelector(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCustomSelector:"), value)
}

// The left expression for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscomparisonpredicate/leftexpression
func (c_ ComparisonPredicate) LeftExpression() NSExpression {
	rv := objc.Send[NSExpression](c_.ID, objc.Sel("leftExpression"))
	return rv
}


// SetLeftExpression sets the value of the leftExpression property.
// The left expression for the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscomparisonpredicate/leftexpression
func (c_ ComparisonPredicate) SetLeftExpression(value IExpression) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLeftExpression:"), value)
}

// The options to use for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscomparisonpredicate/options-swift.property
func (c_ ComparisonPredicate) Options() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("options"))
	return rv
}


// SetOptions sets the value of the options property.
// The options to use for the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscomparisonpredicate/options-swift.property
func (c_ ComparisonPredicate) SetOptions(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOptions:"), value)
}

// The predicate type for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscomparisonpredicate/predicateoperatortype
func (c_ ComparisonPredicate) PredicateOperatorType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("predicateOperatorType"))
	return rv
}


// SetPredicateOperatorType sets the value of the predicateOperatorType property.
// The predicate type for the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscomparisonpredicate/predicateoperatortype
func (c_ ComparisonPredicate) SetPredicateOperatorType(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPredicateOperatorType:"), value)
}



