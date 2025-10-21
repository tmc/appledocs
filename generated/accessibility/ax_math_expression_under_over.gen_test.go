// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility_test

import (
	"github.com/tmc/appledocs/generated/accessibility"
)

// Suppress unused import errors
var _ = accessibility.NewAXMathExpressionUnderOver

// ExampleNewAXMathExpressionUnderOverWithBaseExpressionUnderExpressionOverExpression demonstrates how to create a AXMathExpressionUnderOver instance using NewAXMathExpressionUnderOverWithBaseExpressionUnderExpressionOverExpression.
func ExampleNewAXMathExpressionUnderOverWithBaseExpressionUnderExpressionOverExpression() {
	_ = accessibility.NewAXMathExpressionUnderOverWithBaseExpressionUnderExpressionOverExpression(
		accessibility.AXMathExpression{}, // baseExpression AXMathExpression
		accessibility.AXMathExpression{}, // underExpression AXMathExpression
		accessibility.AXMathExpression{}, // overExpression AXMathExpression
	)
	// Output:
}
