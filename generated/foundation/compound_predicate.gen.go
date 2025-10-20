// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CompoundPredicate] class.
var (
	compoundPredicateClass     _CompoundPredicateClass
	compoundPredicateClassOnce sync.Once
)

func getCompoundPredicateClass() _CompoundPredicateClass {
	compoundPredicateClassOnce.Do(func() {
		compoundPredicateClass = _CompoundPredicateClass{objc.GetClass("NSCompoundPredicate")}
	})
	return compoundPredicateClass
}

type _CompoundPredicateClass struct {
	class objc.Class
}

// An interface definition for the [CompoundPredicate] class.
type ICompoundPredicate interface {
	IPredicate
}

// A specialized predicate that evaluates logical combinations of other predicates.
//
// Use to create an or compound predicate of one or more other predicates, or the of a single predicate. For the logical and operations: An predicate with no subpredicates evaluates to . An predicate with no subpredicates evaluates to . A compound predicate with one or more subpredicates evaluates to the truth of its subpredicates.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate
type CompoundPredicate struct {
	Predicate
}

// CompoundPredicateFrom constructs a [CompoundPredicate] from an unsafe.Pointer.
//
// A specialized predicate that evaluates logical combinations of other predicates.
func CompoundPredicateFrom(ptr unsafe.Pointer) CompoundPredicate {
	return CompoundPredicate{
		Predicate: PredicateFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CompoundPredicateClass) Alloc() CompoundPredicate {
	rv := objc.Send[CompoundPredicate](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CompoundPredicateClass) New() CompoundPredicate {
	rv := objc.Send[CompoundPredicate](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CompoundPredicate) Init() CompoundPredicate {
	rv := objc.Send[CompoundPredicate](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CompoundPredicate) Autorelease() CompoundPredicate {
	rv := objc.Send[CompoundPredicate](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCompoundPredicate creates a new CompoundPredicate instance.
func NewCompoundPredicate() CompoundPredicate {
	return getCompoundPredicateClass().New()
}




