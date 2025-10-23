// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewExpression

// ExampleNewExpressionForAggregate demonstrates how to create a Expression instance using NewExpressionForAggregate.
// Creates an aggregate expression for a specified collection.
func ExampleNewExpressionForAggregate() {
	_ = foundation.NewExpressionForAggregate(
		[]foundation.Expression{}, // subexpressions []Expression
	)
	// Output:
}
// ExampleNewExpressionForConditionalTrueExpressionFalseExpression demonstrates how to create a Expression instance using NewExpressionForConditionalTrueExpressionFalseExpression.
// Creates an expression that returns a result, depending on the value of predicate.
func ExampleNewExpressionForConditionalTrueExpressionFalseExpression() {
	_ = foundation.NewExpressionForConditionalTrueExpressionFalseExpression(
		foundation.NSPredicate{}, // predicate NSPredicate
		foundation.NSExpression{}, // trueExpression NSExpression
		foundation.NSExpression{}, // falseExpression NSExpression
	)
	// Output:
}
// ExampleNewExpressionForIntersectSetWith demonstrates how to create a Expression instance using NewExpressionForIntersectSetWith.
// Creates an expression object that represents the intersection of a specified set and collection.
func ExampleNewExpressionForIntersectSetWith() {
	_ = foundation.NewExpressionForIntersectSetWith(
		foundation.NSExpression{}, // left NSExpression
		foundation.NSExpression{}, // right NSExpression
	)
	// Output:
}
// ExampleNewExpressionForKeyPath demonstrates how to create a Expression instance using NewExpressionForKeyPath.
// Creates an expression that invokes the value function with a specified key path.
func ExampleNewExpressionForKeyPath() {
	_ = foundation.NewExpressionForKeyPath(
		"/tmp/test", // keyPath string
	)
	// Output:
}
// ExampleNewExpressionForMinusSetWith demonstrates how to create a Expression instance using NewExpressionForMinusSetWith.
// Creates an expression object that represents the subtraction of a specified collection from a specified set.
func ExampleNewExpressionForMinusSetWith() {
	_ = foundation.NewExpressionForMinusSetWith(
		foundation.NSExpression{}, // left NSExpression
		foundation.NSExpression{}, // right NSExpression
	)
	// Output:
}
// ExampleNewExpressionForSubqueryUsingIteratorVariablePredicate demonstrates how to create a Expression instance using NewExpressionForSubqueryUsingIteratorVariablePredicate.
// Creates an expression that filters a collection by storing elements in the collection in a specified variable and keeping the elements that the qualifier returns as true.
func ExampleNewExpressionForSubqueryUsingIteratorVariablePredicate() {
	_ = foundation.NewExpressionForSubqueryUsingIteratorVariablePredicate(
		foundation.NSExpression{}, // expression NSExpression
		"variable", // variable string
		foundation.NSPredicate{}, // predicate NSPredicate
	)
	// Output:
}
// ExampleNewExpressionForUnionSetWith demonstrates how to create a Expression instance using NewExpressionForUnionSetWith.
// Creates an expression object that represents the union of a specified set and collection.
func ExampleNewExpressionForUnionSetWith() {
	_ = foundation.NewExpressionForUnionSetWith(
		foundation.NSExpression{}, // left NSExpression
		foundation.NSExpression{}, // right NSExpression
	)
	// Output:
}
// ExampleNewExpressionForVariable demonstrates how to create a Expression instance using NewExpressionForVariable.
// Creates an expression that extracts a value from the variable bindings dictionary for a specified key.
func ExampleNewExpressionForVariable() {
	_ = foundation.NewExpressionForVariable(
		"string", // string string
	)
	// Output:
}
// ExampleNewExpressionWithCoder demonstrates how to create a Expression instance using NewExpressionWithCoder.
// Creates an expression by decoding from the coder you specify.
func ExampleNewExpressionWithCoder() {
	_ = foundation.NewExpressionWithCoder(
		foundation.NSCoder{}, // coder NSCoder
	)
	// Output:
}
// ExampleNewExpressionWithExpressionType demonstrates how to create a Expression instance using NewExpressionWithExpressionType.
// Creates the expression with the specified expression type.
func ExampleNewExpressionWithExpressionType() {
	_ = foundation.NewExpressionWithExpressionType(
		foundation.ExpressionType{}, // type ExpressionType
	)
	// Output:
}
