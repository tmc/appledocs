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
// ExampleHashTable_ObjectEnumerator demonstrates using ObjectEnumerator on a HashTable instance.
// Returns an enumerator object that lets you access each object in the hash table.
func ExampleHashTable_ObjectEnumerator() {
	obj := foundation.NewHashTable()
	_ = obj.ObjectEnumerator()
	// Output:
	}

// ExampleHashTable_RemoveAllObjects demonstrates using RemoveAllObjects on a HashTable instance.
// Removes all objects from the hash table.
func ExampleHashTable_RemoveAllObjects() {
	obj := foundation.NewHashTable()
	obj.RemoveAllObjects()
	// Output:
	}

