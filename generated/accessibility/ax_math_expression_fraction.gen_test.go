// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility_test

import (
	"github.com/tmc/appledocs/generated/accessibility"
)

// Suppress unused import errors
var _ = accessibility.NewAXMathExpressionFraction

// ExampleNewAXMathExpressionFractionWithNumeratorExpressionDenimonatorExpression demonstrates how to create a AXMathExpressionFraction instance using NewAXMathExpressionFractionWithNumeratorExpressionDenimonatorExpression.
func ExampleNewAXMathExpressionFractionWithNumeratorExpressionDenimonatorExpression() {
	_ = accessibility.NewAXMathExpressionFractionWithNumeratorExpressionDenimonatorExpression(
		accessibility.AXMathExpression{}, // numeratorExpression AXMathExpression
		accessibility.AXMathExpression{}, // denimonatorExpression AXMathExpression
	)
	// Output:
}
