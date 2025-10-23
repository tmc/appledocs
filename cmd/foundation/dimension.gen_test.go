// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewDimension

// ExampleNewDimensionWithSymbolConverter demonstrates how to create a Dimension instance using NewDimensionWithSymbolConverter.
// Initializes a dimensional unit with the symbol and unit converter you specify.
func ExampleNewDimensionWithSymbolConverter() {
	_ = foundation.NewDimensionWithSymbolConverter(
		"symbol", // symbol string
		foundation.NSUnitConverter{}, // converter NSUnitConverter
	)
	// Output:
}
