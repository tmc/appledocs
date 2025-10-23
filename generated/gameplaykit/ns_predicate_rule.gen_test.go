// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit_test

import (
	"github.com/tmc/appledocs/generated/gameplaykit"
)

// Suppress unused import errors
var _ = gameplaykit.NewNSPredicateRule

// ExampleNewNSPredicateRuleWithPredicate demonstrates how to create a NSPredicateRule instance using NewNSPredicateRuleWithPredicate.
// Initializes a rule with the specified predicate.
func ExampleNewNSPredicateRuleWithPredicate() {
	_ = gameplaykit.NewNSPredicateRuleWithPredicate(
		gameplaykit.Predicate{}, // predicate Predicate
	)
	// Output:
}
