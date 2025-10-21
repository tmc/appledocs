// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility_test

import (
	"github.com/tmc/appledocs/generated/accessibility"
)

// Suppress unused import errors
var _ = accessibility.NewAXMathExpressionRoot

// ExampleNewAXMathExpressionRootWithRadicandExpressionsRootIndexExpression demonstrates how to create a AXMathExpressionRoot instance using NewAXMathExpressionRootWithRadicandExpressionsRootIndexExpression.
func ExampleNewAXMathExpressionRootWithRadicandExpressionsRootIndexExpression() {
	_ = accessibility.NewAXMathExpressionRootWithRadicandExpressionsRootIndexExpression(
		[]accessibility.AXMathExpression{}, // radicandExpressions []AXMathExpression
		accessibility.AXMathExpression{}, // rootIndexExpression AXMathExpression
	)
	// Output:
}
