// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility_test

import (
	"github.com/tmc/appledocs/generated/accessibility"
)

// Suppress unused import errors
var _ = accessibility.NewAXMathExpressionTable

// ExampleNewAXMathExpressionTableWithExpressions demonstrates how to create a AXMathExpressionTable instance using NewAXMathExpressionTableWithExpressions.
func ExampleNewAXMathExpressionTableWithExpressions() {
	_ = accessibility.NewAXMathExpressionTableWithExpressions(
		[]accessibility.IAXMathExpression{}, // expressions []IAXMathExpression
	)
	// Output:
}
