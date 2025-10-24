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
		[]foundation.IPredicate{}, // subpredicates []IPredicate
	)
	// Output:
}
// ExampleNewCompoundPredicateOrPredicateWithSubpredicates demonstrates how to create a CompoundPredicate instance using NewCompoundPredicateOrPredicateWithSubpredicates.
// Returns a new predicate that you form using an OR operation on the predicates in a specified array.
func ExampleNewCompoundPredicateOrPredicateWithSubpredicates() {
	_ = foundation.NewCompoundPredicateOrPredicateWithSubpredicates(
		[]foundation.IPredicate{}, // subpredicates []IPredicate
	)
	// Output:
}
// ExampleNewCompoundPredicateWithTypeSubpredicates demonstrates how to create a CompoundPredicate instance using NewCompoundPredicateWithTypeSubpredicates.
// Returns the receiver that a specified type initializes using predicates from a specified array.
func ExampleNewCompoundPredicateWithTypeSubpredicates() {
	_ = foundation.NewCompoundPredicateWithTypeSubpredicates(
		foundation.CompoundPredicateType{}, // type CompoundPredicateType
		[]foundation.IPredicate{}, // subpredicates []IPredicate
	)
	// Output:
}
