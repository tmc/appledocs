// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CompoundPredicate] class.
var compoundPredicateClass = _CompoundPredicateClass{objc.GetClass("NSCompoundPredicate")}

type _CompoundPredicateClass struct {
	class objc.Class
}

// An interface definition for the [CompoundPredicate] class.
type ICompoundPredicate interface {
	IPredicate
}

// A specialized predicate that evaluates logical combinations of other predicates. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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
	return compoundPredicateClass.New()
}




