// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewExpression

// ExampleNewExpressionWithExpressionType demonstrates how to create a Expression instance using NewExpressionWithExpressionType.
// Creates the expression with the specified expression type.
func ExampleNewExpressionWithExpressionType() {
	_ = foundation.NewExpressionWithExpressionType(
		foundation.ExpressionType{}, // type ExpressionType
	)
	// Output:
}
