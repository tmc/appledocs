// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewArray

// ExampleNewArray demonstrates how to create a Array instance.
// Initializes a newly allocated array.
func ExampleNewArray() {
	_ = foundation.NewArray()
	// Output:
}
// ExampleArray_ObjectEnumerator demonstrates using ObjectEnumerator on a Array instance.
// Returns an enumerator object that lets you access each object in the array.
func ExampleArray_ObjectEnumerator() {
	obj := foundation.NewArray()
	_ = obj.ObjectEnumerator()
	// Output:
	}

// ExampleArray_ReverseObjectEnumerator demonstrates using ReverseObjectEnumerator on a Array instance.
// Returns an enumerator object that lets you access each object in the array, in reverse order.
func ExampleArray_ReverseObjectEnumerator() {
	obj := foundation.NewArray()
	_ = obj.ReverseObjectEnumerator()
	// Output:
	}

// ExampleArray_ShuffledArray demonstrates using ShuffledArray on a Array instance.
// Returns a new array that lists this array’s elements in a random order.
func ExampleArray_ShuffledArray() {
	obj := foundation.NewArray()
	_ = obj.ShuffledArray()
	// Output:
	}

