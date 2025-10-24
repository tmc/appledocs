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
