// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewURLQueryItem

// ExampleNewURLQueryItemWithNameValue demonstrates how to create a URLQueryItem instance using NewURLQueryItemWithNameValue.
// Initializes a newly allocated query item with the specified name and value.
func ExampleNewURLQueryItemWithNameValue() {
	_ = foundation.NewURLQueryItemWithNameValue(
		"name", // name string
		"value", // value string
	)
	// Output:
}
