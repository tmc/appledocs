// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility_test

import (
	"github.com/tmc/appledocs/generated/accessibility"
)

// Suppress unused import errors
var _ = accessibility.NewAXMathExpressionTableCell

// ExampleNewAXMathExpressionTableCellWithExpressions demonstrates how to create a AXMathExpressionTableCell instance using NewAXMathExpressionTableCellWithExpressions.
func ExampleNewAXMathExpressionTableCellWithExpressions() {
	_ = accessibility.NewAXMathExpressionTableCellWithExpressions(
		[]accessibility.IAXMathExpression{}, // expressions []IAXMathExpression
	)
	// Output:
}
