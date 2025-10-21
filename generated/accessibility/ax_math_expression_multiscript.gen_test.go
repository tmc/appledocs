// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility_test

import (
	"github.com/tmc/appledocs/generated/accessibility"
)

// Suppress unused import errors
var _ = accessibility.NewAXMathExpressionMultiscript

// ExampleNewAXMathExpressionMultiscriptWithBaseExpressionPrescriptExpressionsPostscriptExpressions demonstrates how to create a AXMathExpressionMultiscript instance using NewAXMathExpressionMultiscriptWithBaseExpressionPrescriptExpressionsPostscriptExpressions.
func ExampleNewAXMathExpressionMultiscriptWithBaseExpressionPrescriptExpressionsPostscriptExpressions() {
	_ = accessibility.NewAXMathExpressionMultiscriptWithBaseExpressionPrescriptExpressionsPostscriptExpressions(
		accessibility.AXMathExpression{}, // baseExpression AXMathExpression
		[]accessibility.AXMathExpressionSubSuperscript{}, // prescriptExpressions []AXMathExpressionSubSuperscript
		[]accessibility.AXMathExpressionSubSuperscript{}, // postscriptExpressions []AXMathExpressionSubSuperscript
	)
	// Output:
}
