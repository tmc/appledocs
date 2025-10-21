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

// ExampleNewMutableDictionaryWithContentsOfFile demonstrates how to create a MutableDictionary instance using NewMutableDictionaryWithContentsOfFile.
func ExampleNewMutableDictionaryWithContentsOfFile() {
	_ = foundation.NewMutableDictionaryWithContentsOfFile(
		"/tmp/test", // path string
	)
	// Output:
}


