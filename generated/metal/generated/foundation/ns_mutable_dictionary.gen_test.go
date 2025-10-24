// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewMutableDictionary

// ExampleNewMutableDictionary demonstrates how to create a MutableDictionary instance.
// Initializes a newly allocated mutable dictionary.
func ExampleNewMutableDictionary() {
	_ = foundation.NewMutableDictionary()
	// Output:
}
// ExampleNewMutableDictionaryWithCapacity demonstrates how to create a MutableDictionary instance using NewMutableDictionaryWithCapacity.
// Initializes a newly allocated mutable dictionary, allocating enough memory to hold   entries.
func ExampleNewMutableDictionaryWithCapacity() {
	_ = foundation.NewMutableDictionaryWithCapacity(
		0, // numItems uint
	)
	// Output:
}
// ExampleMutableDictionary_GetHeaderBytes demonstrates using GetHeaderBytes on a MutableDictionary instance.
func ExampleMutableDictionary_GetHeaderBytes() {
	obj := foundation.NewMutableDictionary()
	_ = obj.GetHeaderBytes()
	// Output:
	}

// ExampleMutableDictionary_RemoveAllObjects demonstrates using RemoveAllObjects on a MutableDictionary instance.
// Empties the dictionary of its entries.
func ExampleMutableDictionary_RemoveAllObjects() {
	obj := foundation.NewMutableDictionary()
	obj.RemoveAllObjects()
	// Output:
	}

