// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewComparisonPredicate

// ExampleNewComparisonPredicateWithCoder demonstrates how to create a ComparisonPredicate instance using NewComparisonPredicateWithCoder.
// Creates a predicate by decoding from the coder you specify.
func ExampleNewComparisonPredicateWithCoder() {
	_ = foundation.NewComparisonPredicateWithCoder(
		foundation.NSCoder{}, // coder NSCoder
	)
	// Output:
}
// ExampleNewComparisonPredicateWithLeftExpressionRightExpressionCustomSelector demonstrates how to create a ComparisonPredicate instance using NewComparisonPredicateWithLeftExpressionRightExpressionCustomSelector.
// Creates a predicate that you form by combining specified left and right expressions using a specified selector.
func ExampleNewComparisonPredicateWithLeftExpressionRightExpressionCustomSelector() {
	_ = foundation.NewComparisonPredicateWithLeftExpressionRightExpressionCustomSelector(
		foundation.NSExpression{}, // lhs NSExpression
		foundation.NSExpression{}, // rhs NSExpression
		0, // selector objc.SEL
	)
	// Output:
}
// ExampleNewComparisonPredicateWithLeftExpressionRightExpressionModifierTypeOptions demonstrates how to create a ComparisonPredicate instance using NewComparisonPredicateWithLeftExpressionRightExpressionModifierTypeOptions.
// Creates a predicate to a specified type that you form by combining specified left and right expressions using a specified modifier and options.
func ExampleNewComparisonPredicateWithLeftExpressionRightExpressionModifierTypeOptions() {
	_ = foundation.NewComparisonPredicateWithLeftExpressionRightExpressionModifierTypeOptions(
		foundation.NSExpression{}, // lhs NSExpression
		foundation.NSExpression{}, // rhs NSExpression
		foundation.ComparisonPredicateModifier{}, // modifier ComparisonPredicateModifier
		foundation.PredicateOperatorType{}, // type PredicateOperatorType
		foundation.ComparisonPredicateOptions{}, // options ComparisonPredicateOptions
	)
	// Output:
}
