// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewHashTable

// ExampleNewHashTableWithOptions demonstrates how to create a HashTable instance using NewHashTableWithOptions.
// Returns a hash table with given pointer functions options.
func ExampleNewHashTableWithOptions() {
	_ = foundation.NewHashTableWithOptions(
		foundation.PointerFunctionsOptions{}, // options PointerFunctionsOptions
	)
	// Output:
}
