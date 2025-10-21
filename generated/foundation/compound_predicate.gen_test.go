// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewCompoundPredicate

// ExampleNewCompoundPredicateAndPredicateWithSubpredicates demonstrates how to create a CompoundPredicate instance using NewCompoundPredicateAndPredicateWithSubpredicates.
// Returns a new predicate that you form using an AND operation on the predicates in a specified array.
func ExampleNewCompoundPredicateAndPredicateWithSubpredicates() {
	_ = foundation.NewCompoundPredicateAndPredicateWithSubpredicates(
		[]foundation.Predicate{}, // subpredicates []Predicate
	)
	// Output:
}
// ExampleNewCompoundPredicateNotPredicateWithSubpredicate demonstrates how to create a CompoundPredicate instance using NewCompoundPredicateNotPredicateWithSubpredicate.
// Returns a new predicate that you form using a NOT operation on a specified predicate.
func ExampleNewCompoundPredicateNotPredicateWithSubpredicate() {
	_ = foundation.NewCompoundPredicateNotPredicateWithSubpredicate(
		foundation.NSPredicate{}, // predicate NSPredicate
	)
	// Output:
}
// ExampleNewCompoundPredicateWithTypeSubpredicates demonstrates how to create a CompoundPredicate instance using NewCompoundPredicateWithTypeSubpredicates.
// Returns the receiver that a specified type initializes using predicates from a specified array.
func ExampleNewCompoundPredicateWithTypeSubpredicates() {
	_ = foundation.NewCompoundPredicateWithTypeSubpredicates(
		foundation.CompoundPredicateType{}, // type CompoundPredicateType
		[]foundation.Predicate{}, // subpredicates []Predicate
	)
	// Output:
}
