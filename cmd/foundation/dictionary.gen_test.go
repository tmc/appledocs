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
// ExampleNewDictionaryWithCoder demonstrates how to create a Dictionary instance using NewDictionaryWithCoder.
// Creates a dictionary initialized from data in the provided unarchiver.
func ExampleNewDictionaryWithCoder() {
	_ = foundation.NewDictionaryWithCoder(
		foundation.NSCoder{}, // coder NSCoder
	)
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
// ExampleNewDictionaryWithContentsOfURL demonstrates how to create a Dictionary instance using NewDictionaryWithContentsOfURL.
// Initializes a newly allocated dictionary using the keys and values found at a given URL.
func ExampleNewDictionaryWithContentsOfURL() {
	_ = foundation.NewDictionaryWithContentsOfURL(
		foundation.URL{}, // url URL
	)
	// Output:
}
// ExampleNewDictionaryWithContentsOfURLError demonstrates how to create a Dictionary instance using NewDictionaryWithContentsOfURLError.
// Initializes a newly allocated dictionary using the keys and values found at a given URL.
func ExampleNewDictionaryWithContentsOfURLError() {
	_ = foundation.NewDictionaryWithContentsOfURLError(
		foundation.URL{}, // url URL
		foundation.NSError{}, // error NSError
	)
	// Output:
}
// ExampleNewDictionaryWithDictionary demonstrates how to create a Dictionary instance using NewDictionaryWithDictionary.
// Initializes a newly allocated dictionary by placing in it the keys and values contained in another given dictionary.
func ExampleNewDictionaryWithDictionary() {
	_ = foundation.NewDictionaryWithDictionary(
		foundation.IDictionary{}, // otherDictionary IDictionary
	)
	// Output:
}
// ExampleNewDictionaryWithDictionaryCopyItems demonstrates how to create a Dictionary instance using NewDictionaryWithDictionaryCopyItems.
// Initializes a newly allocated dictionary using the objects contained in another given dictionary.
func ExampleNewDictionaryWithDictionaryCopyItems() {
	_ = foundation.NewDictionaryWithDictionaryCopyItems(
		foundation.IDictionary{}, // otherDictionary IDictionary
		false, // flag bool
	)
	// Output:
}
