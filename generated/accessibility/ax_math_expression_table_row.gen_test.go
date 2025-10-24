// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility_test

import (
	"github.com/tmc/appledocs/generated/accessibility"
)

// Suppress unused import errors
var _ = accessibility.NewAXMathExpressionTableRow

// ExampleNewAXMathExpressionTableRowWithExpressions demonstrates how to create a AXMathExpressionTableRow instance using NewAXMathExpressionTableRowWithExpressions.
func ExampleNewAXMathExpressionTableRowWithExpressions() {
	_ = accessibility.NewAXMathExpressionTableRowWithExpressions(
		[]accessibility.IAXMathExpression{}, // expressions []IAXMathExpression
	)
	// Output:
}
