// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility_test

import (
	"github.com/tmc/appledocs/generated/accessibility"
)

// Suppress unused import errors
var _ = accessibility.NewAXMathExpressionSubSuperscript

// ExampleNewAXMathExpressionSubSuperscriptWithBaseExpressionSubscriptExpressionsSuperscriptExpressions demonstrates how to create a AXMathExpressionSubSuperscript instance using NewAXMathExpressionSubSuperscriptWithBaseExpressionSubscriptExpressionsSuperscriptExpressions.
func ExampleNewAXMathExpressionSubSuperscriptWithBaseExpressionSubscriptExpressionsSuperscriptExpressions() {
	_ = accessibility.NewAXMathExpressionSubSuperscriptWithBaseExpressionSubscriptExpressionsSuperscriptExpressions(
		[]accessibility.AXMathExpression{}, // baseExpression []AXMathExpression
		[]accessibility.AXMathExpression{}, // subscriptExpressions []AXMathExpression
		[]accessibility.AXMathExpression{}, // superscriptExpressions []AXMathExpression
	)
	// Output:
}
