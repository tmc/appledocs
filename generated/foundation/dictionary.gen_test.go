// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewDictionary


// ExampleNewDictionary demonstrates how to create a Dictionary instance.
// Initializes a newly allocated dictionary.
func ExampleNewDictionary() {
	_ = foundation.NewDictionary()
	// Output:
}



// ExampleNewDictionaryWithContentsOfFile demonstrates how to create a Dictionary instance using NewDictionaryWithContentsOfFile.
// Initializes a newly allocated dictionary using the keys and values found in a file at a given path.
func ExampleNewDictionaryWithContentsOfFile() {
	_ = foundation.NewDictionaryWithContentsOfFile(
		"/tmp/test", // path string
	)
	// Output:
}










