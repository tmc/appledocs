// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewDictionary

// ExampleNewDictionaryWithContentsOfFile demonstrates how to create a Dictionary instance using NewDictionaryWithContentsOfFile.
// Initializes a newly allocated dictionary using the keys and values found in a file at a given path.
func ExampleNewDictionaryWithContentsOfFile() {
	_ = foundation.NewDictionaryWithContentsOfFile(
		foundation.NSString{}, // path NSString
	)
	// Output:
}
// ExampleNewDictionaryWithContentsOfURL demonstrates how to create a Dictionary instance using NewDictionaryWithContentsOfURL.
// Initializes a newly allocated dictionary using the keys and values found at a given URL.
func ExampleNewDictionaryWithContentsOfURL() {
	_ = foundation.NewDictionaryWithContentsOfURL(
		foundation.NSURL{}, // url NSURL
	)
	// Output:
}
