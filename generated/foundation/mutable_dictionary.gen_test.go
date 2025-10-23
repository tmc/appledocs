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
// ExampleNewMutableDictionaryWithCoder demonstrates how to create a MutableDictionary instance using NewMutableDictionaryWithCoder.
func ExampleNewMutableDictionaryWithCoder() {
	_ = foundation.NewMutableDictionaryWithCoder(
		foundation.NSCoder{}, // coder NSCoder
	)
	// Output:
}
