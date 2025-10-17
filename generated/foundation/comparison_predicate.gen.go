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



