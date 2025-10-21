// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewSortDescriptor

// ExampleNewSortDescriptorWithCoder demonstrates how to create a SortDescriptor instance using NewSortDescriptorWithCoder.
// Creates a sort descriptor by decoding from the coder you specify.
func ExampleNewSortDescriptorWithCoder() {
	_ = foundation.NewSortDescriptorWithCoder(
		foundation.NSCoder{}, // coder NSCoder
	)
	// Output:
}
