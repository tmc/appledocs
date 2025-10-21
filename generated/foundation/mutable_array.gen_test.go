// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewMutableArray

// ExampleNewMutableArrayWithContentsOfURL demonstrates how to create a MutableArray instance using NewMutableArrayWithContentsOfURL.
// Initialized a newly allocated mutable array with the contents of the location specified by a given URL.
func ExampleNewMutableArrayWithContentsOfURL() {
	_ = foundation.NewMutableArrayWithContentsOfURL(
		foundation.URL{}, // url URL
	)
	// Output:
}
