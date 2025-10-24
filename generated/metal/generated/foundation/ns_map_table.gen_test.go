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
// ExampleMapTable_DictionaryRepresentation demonstrates using DictionaryRepresentation on a MapTable instance.
// Returns a dictionary representation of the map table.
func ExampleMapTable_DictionaryRepresentation() {
	obj := foundation.NewMapTable()
	_ = obj.DictionaryRepresentation()
	// Output:
	}

// ExampleMapTable_KeyEnumerator demonstrates using KeyEnumerator on a MapTable instance.
// Returns an enumerator object that lets you access each key in the map table.
func ExampleMapTable_KeyEnumerator() {
	obj := foundation.NewMapTable()
	_ = obj.KeyEnumerator()
	// Output:
	}

// ExampleMapTable_ObjectEnumerator demonstrates using ObjectEnumerator on a MapTable instance.
// Returns an enumerator object that lets you access each value in the map table.
func ExampleMapTable_ObjectEnumerator() {
	obj := foundation.NewMapTable()
	_ = obj.ObjectEnumerator()
	// Output:
	}

// ExampleMapTable_RemoveAllObjects demonstrates using RemoveAllObjects on a MapTable instance.
// Empties the map table of its entries.
func ExampleMapTable_RemoveAllObjects() {
	obj := foundation.NewMapTable()
	obj.RemoveAllObjects()
	// Output:
	}

