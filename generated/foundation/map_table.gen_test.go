// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewMapTable

// ExampleNewMapTableWithKeyOptionsValueOptions demonstrates how to create a MapTable instance using NewMapTableWithKeyOptionsValueOptions.
// Returns a new map table, initialized with the given options
func ExampleNewMapTableWithKeyOptionsValueOptions() {
	_ = foundation.NewMapTableWithKeyOptionsValueOptions(
		foundation.PointerFunctionsOptions{}, // keyOptions PointerFunctionsOptions
		foundation.PointerFunctionsOptions{}, // valueOptions PointerFunctionsOptions
	)
	// Output:
}
// ExampleNewMapTableWithKeyOptionsValueOptionsCapacity demonstrates how to create a MapTable instance using NewMapTableWithKeyOptionsValueOptionsCapacity.
// Returns a map table, initialized with the given options.
func ExampleNewMapTableWithKeyOptionsValueOptionsCapacity() {
	_ = foundation.NewMapTableWithKeyOptionsValueOptionsCapacity(
		foundation.PointerFunctionsOptions{}, // keyOptions PointerFunctionsOptions
		foundation.PointerFunctionsOptions{}, // valueOptions PointerFunctionsOptions
		0, // initialCapacity uint
	)
	// Output:
}
// ExampleNewMapTableWithKeyPointerFunctionsValuePointerFunctionsCapacity demonstrates how to create a MapTable instance using NewMapTableWithKeyPointerFunctionsValuePointerFunctionsCapacity.
// Returns a map table, initialized with the given functions.
func ExampleNewMapTableWithKeyPointerFunctionsValuePointerFunctionsCapacity() {
	_ = foundation.NewMapTableWithKeyPointerFunctionsValuePointerFunctionsCapacity(
		foundation.NSPointerFunctions{}, // keyFunctions NSPointerFunctions
		foundation.NSPointerFunctions{}, // valueFunctions NSPointerFunctions
		0, // initialCapacity uint
	)
	// Output:
}
