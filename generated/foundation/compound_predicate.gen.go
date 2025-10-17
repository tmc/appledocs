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



