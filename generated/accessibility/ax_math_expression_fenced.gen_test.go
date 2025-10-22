// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility_test

import (
	"github.com/tmc/appledocs/generated/accessibility"
)

// Suppress unused import errors
var _ = accessibility.NewAXMathExpressionFenced

// ExampleNewAXMathExpressionFencedWithExpressionsOpenStringCloseString demonstrates how to create a AXMathExpressionFenced instance using NewAXMathExpressionFencedWithExpressionsOpenStringCloseString.
func ExampleNewAXMathExpressionFencedWithExpressionsOpenStringCloseString() {
	_ = accessibility.NewAXMathExpressionFencedWithExpressionsOpenStringCloseString(
		[]accessibility.AXMathExpression{}, // expressions []AXMathExpression
		"openString", // openString string
		"closeString", // closeString string
	)
	// Output:
}
