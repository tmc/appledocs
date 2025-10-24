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
// ExampleNewExpressionWithExpressionType demonstrates how to create a Expression instance using NewExpressionWithExpressionType.
// Creates the expression with the specified expression type.
func ExampleNewExpressionWithExpressionType() {
	_ = foundation.NewExpressionWithExpressionType(
		foundation.ExpressionType{}, // type ExpressionType
	)
	// Output:
}
// ExampleExpression_AllowEvaluation demonstrates using AllowEvaluation on a Expression instance.
// Forces a securely decoded expression to allow evaluation.
func ExampleExpression_AllowEvaluation() {
	obj := foundation.NewExpression()
	obj.AllowEvaluation()
	// Output:
	}

