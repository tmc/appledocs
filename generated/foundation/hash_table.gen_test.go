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
// ExampleNewHashTableWithOptionsCapacity demonstrates how to create a HashTable instance using NewHashTableWithOptionsCapacity.
// Returns a hash table initialized with the given attributes.
func ExampleNewHashTableWithOptionsCapacity() {
	_ = foundation.NewHashTableWithOptionsCapacity(
		foundation.PointerFunctionsOptions{}, // options PointerFunctionsOptions
		0, // initialCapacity uint
	)
	// Output:
}
// ExampleNewHashTableWithPointerFunctionsCapacity demonstrates how to create a HashTable instance using NewHashTableWithPointerFunctionsCapacity.
// Returns a hash table initialized with the given functions and capacity.
func ExampleNewHashTableWithPointerFunctionsCapacity() {
	_ = foundation.NewHashTableWithPointerFunctionsCapacity(
		foundation.NSPointerFunctions{}, // functions NSPointerFunctions
		0, // initialCapacity uint
	)
	// Output:
}
