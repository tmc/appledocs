// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewDecimalNumber

// ExampleNewDecimalNumberWithDecimal demonstrates how to create a DecimalNumber instance using NewDecimalNumberWithDecimal.
// Initializes a decimal number to represent a given decimal.
func ExampleNewDecimalNumberWithDecimal() {
	_ = foundation.NewDecimalNumberWithDecimal(
		foundation.Decimal{}, // dcm Decimal
	)
	// Output:
}
// ExampleNewDecimalNumberWithString demonstrates how to create a DecimalNumber instance using NewDecimalNumberWithString.
// Initializes a decimal number so that its value is equivalent to that in a given numeric string.
func ExampleNewDecimalNumberWithString() {
	_ = foundation.NewDecimalNumberWithString(
		"numberValue", // numberValue string
	)
	// Output:
}
