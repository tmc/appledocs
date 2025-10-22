// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewUnitConverterLinear

// ExampleNewUnitConverterLinearWithCoefficient demonstrates how to create a UnitConverterLinear instance using NewUnitConverterLinearWithCoefficient.
// Initializes the unit converter with the coefficient you specify.
func ExampleNewUnitConverterLinearWithCoefficient() {
	_ = foundation.NewUnitConverterLinearWithCoefficient(
		0.0, // coefficient float64
	)
	// Output:
}
// ExampleNewUnitConverterLinearWithCoefficientConstant demonstrates how to create a UnitConverterLinear instance using NewUnitConverterLinearWithCoefficientConstant.
// Creates a unit converter with the coefficient and constant you specify.
func ExampleNewUnitConverterLinearWithCoefficientConstant() {
	_ = foundation.NewUnitConverterLinearWithCoefficientConstant(
		0.0, // coefficient float64
		0.0, // constant float64
	)
	// Output:
}
