// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ComparisonPredicate] class.
var comparisonPredicateClass = _ComparisonPredicateClass{objc.GetClass("NSComparisonPredicate")}

type _ComparisonPredicateClass struct {
	class objc.Class
}

// An interface definition for the [ComparisonPredicate] class.
type IComparisonPredicate interface {
	IPredicate
}

// A specialized predicate for comparing expressions. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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
	return comparisonPredicateClass.New()
}




