// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewItemProvider

// ExampleNewItemProvider demonstrates how to create a ItemProvider instance.
// Creates an empty item provider to which you can later register a data or file representation.
func ExampleNewItemProvider() {
	_ = foundation.NewItemProvider()
	// Output:
}
// ExampleNewItemProviderWithContentsOfURL demonstrates how to create a ItemProvider instance using NewItemProviderWithContentsOfURL.
// Provides data-backed content from an existing file.
func ExampleNewItemProviderWithContentsOfURL() {
	_ = foundation.NewItemProviderWithContentsOfURL(
		foundation.URL{}, // fileURL URL
	)
	// Output:
}
