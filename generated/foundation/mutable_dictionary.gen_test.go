// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewMutableDictionary


// ExampleNewMutableDictionaryWithCapacity demonstrates how to create a MutableDictionary instance using NewMutableDictionaryWithCapacity.
// Initializes a newly allocated mutable dictionary, allocating enough memory to hold   entries.
func ExampleNewMutableDictionaryWithCapacity() {
	_ = foundation.NewMutableDictionaryWithCapacity(
		0, // numItems uint
	)
	// Output:
}


// ExampleNewMutableDictionaryWithSharedKeySet demonstrates how to create a MutableDictionary instance using NewMutableDictionaryWithSharedKeySet.
// Creates a mutable dictionary which is optimized for dealing with a known set of keys.
func ExampleNewMutableDictionaryWithSharedKeySet() {
	_ = foundation.NewMutableDictionaryWithSharedKeySet(
		0, // keyset objc.ID
	)
	// Output:
}

// ExampleNewMutableDictionary demonstrates how to create a MutableDictionary instance.
// Initializes a newly allocated mutable dictionary.
func ExampleNewMutableDictionary() {
	_ = foundation.NewMutableDictionary()
	// Output:
}


// ExampleNewMutableDictionaryWithContentsOfFile demonstrates how to create a MutableDictionary instance using NewMutableDictionaryWithContentsOfFile.
func ExampleNewMutableDictionaryWithContentsOfFile() {
	_ = foundation.NewMutableDictionaryWithContentsOfFile(
		"/tmp/test", // path string
	)
	// Output:
}




