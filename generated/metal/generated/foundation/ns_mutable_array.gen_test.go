// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewMutableArray

// ExampleNewMutableArray demonstrates how to create a MutableArray instance.
// Initializes a newly allocated array.
func ExampleNewMutableArray() {
	_ = foundation.NewMutableArray()
	// Output:
}
// ExampleNewMutableArrayWithCapacity demonstrates how to create a MutableArray instance using NewMutableArrayWithCapacity.
// Returns an array, initialized with enough memory to initially hold a given number of objects.
func ExampleNewMutableArrayWithCapacity() {
	_ = foundation.NewMutableArrayWithCapacity(
		0, // numItems uint
	)
	// Output:
}
// ExampleMutableArray_RemoveAllObjects demonstrates using RemoveAllObjects on a MutableArray instance.
// Empties the array of all its elements.
func ExampleMutableArray_RemoveAllObjects() {
	obj := foundation.NewMutableArray()
	obj.RemoveAllObjects()
	// Output:
	}

// ExampleMutableArray_RemoveLastObject demonstrates using RemoveLastObject on a MutableArray instance.
// Removes the object with the highest-valued index in the array
func ExampleMutableArray_RemoveLastObject() {
	obj := foundation.NewMutableArray()
	obj.RemoveLastObject()
	// Output:
	}

