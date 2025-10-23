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
// ExampleNewMutableDictionaryWithCoder demonstrates how to create a MutableDictionary instance using NewMutableDictionaryWithCoder.
func ExampleNewMutableDictionaryWithCoder() {
	_ = foundation.NewMutableDictionaryWithCoder(
		foundation.NSCoder{}, // coder NSCoder
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
// ExampleNewMutableDictionaryWithContentsOfURL demonstrates how to create a MutableDictionary instance using NewMutableDictionaryWithContentsOfURL.
func ExampleNewMutableDictionaryWithContentsOfURL() {
	_ = foundation.NewMutableDictionaryWithContentsOfURL(
		foundation.URL{}, // url URL
	)
	// Output:
}
// ExampleNewMutableDictionaryWithOBEXHeadersData demonstrates how to create a MutableDictionary instance using NewMutableDictionaryWithOBEXHeadersData.
func ExampleNewMutableDictionaryWithOBEXHeadersData() {
	_ = foundation.NewMutableDictionaryWithOBEXHeadersData(
		foundation.NSData{}, // inHeadersData NSData
	)
	// Output:
}
